package color

import (
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Attr represents a text style attribute such as color, bold, italic, etc.
type Attr int

var (
	// NoColor disables color output if set to true.
	// This can be controlled via the NO_COLOR environment variable.
	NoColor bool

	// ForceColor forces color output, its ignore terminal support check.
	// This can be controlled via the FORCE_COLOR environment variable.
	ForceColor bool

	singleAttrCache = []string{}
	singleAttrMu    sync.RWMutex

	comboAttrCache = []string{}
	comboAttrMu    sync.RWMutex

	start = "\x1b["
	end   = "\x1b[0m"
)

func init() {
	if os.Getenv("NO_COLOR") != "" {
		NoColor = true
	}

	if os.Getenv("FORCE_COLOR") != "" {
		ForceColor = true
	}
}

// Style applies the given attributes to the text and returns the styled string.
func Style(text string, attrs ...Attr) string {
	if !allowColor() {
		return text
	}

	if len(attrs) == 0 || text == "" {
		return text
	}

	if len(attrs) == 1 {
		return getAnsi(text, attrs[0])
	}

	return getComplexAnsi(text, attrs...)
}

func allowColor() bool {
	if NoColor {
		return false
	}

	if ForceColor {
		return true
	}

	return IsTerminal() && SupportsColor()
}

// IsTerminal reports whether the standard output is connected to a terminal (TTY).
func IsTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()

	return (fileInfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

// SupportsColor reports whether the current environment supports color output,
// based on the TERM and COLORTERM environment variables.
func SupportsColor() bool {
	term := os.Getenv("TERM")
	if term == "" {
		return false
	}

	if os.Getenv("COLORTERM") != "" {
		return true
	}

	return false
}

func getAnsi(text string, attr Attr) string {
	ansi := ansiCodes[attr]
	cacheKey := strconv.Itoa(ansi)

	singleAttrMu.RLock()
	cached, ok := checkCache(singleAttrCache, cacheKey)
	singleAttrMu.RUnlock()

	if !ok {
		singleAttrMu.Lock()
		singleAttrCache = append(singleAttrCache, cacheKey)
		singleAttrMu.Unlock()
	}

	return buildAnsi(text, cached, attr)
}

func getComplexAnsi(text string, attrs ...Attr) string {
	if len(attrs) == 0 {
		return text
	}

	cacheKey := makeAttrSeq(attrs...)
	comboAttrMu.RLock()
	cached, ok := checkCache(comboAttrCache, cacheKey)
	comboAttrMu.RUnlock()

	if !ok {
		comboAttrMu.Lock()
		comboAttrCache = append(comboAttrCache, cacheKey)
		comboAttrMu.Unlock()
	}

	return buildAnsi(text, cached, attrs...)
}

func buildAnsi(text string, cachedAttrs string, attrs ...Attr) string {
	var sb strings.Builder

	sb.WriteString(start)

	if cachedAttrs != "" {
		sb.WriteString(cachedAttrs)
	} else {
		sb.WriteString(makeAttrSeq(attrs...))
	}

	sb.WriteByte('m')
	sb.WriteString(text)
	sb.WriteString(end)

	return sb.String()
}

func makeAttrSeq(attrs ...Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	strAttrs := make([]string, 0, len(attrs))

	for _, attr := range attrs {
		strAttrs = append(strAttrs, strconv.Itoa(ansiCodes[attr]))
	}

	sort.Slice(strAttrs, func(i, j int) bool {
		a, _ := strconv.Atoi(strAttrs[i])
		b, _ := strconv.Atoi(strAttrs[j])
		return a < b
	})

	filtered := slices.Compact(strAttrs)

	return strings.Join(filtered, ";")
}

func checkCache(cache []string, key string) (string, bool) {
	idx := slices.Index(cache, key)

	if idx == -1 {
		return "", false
	} else {
		return cache[idx], true
	}
}
