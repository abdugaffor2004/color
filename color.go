package color

import (
	"slices"
	"strconv"
	"strings"
)

// Style applies the given attributes to the text and returns the styled string.
func Style(text string, attrs ...Attr) string {
	if !allowColor() {
		return text
	}

	if len(attrs) == 0 || text == "" {
		return text
	}

	startSeq, ok := ansiCache.Get(attrs)
	if !ok {
		startSeq = makeStartSeq(attrs)
		ansiCache.Set(attrs, startSeq)
	}

	var sb strings.Builder
	endSeq := "\x1b[0m"

	sb.WriteString(startSeq)
	sb.WriteString(text)
	sb.WriteString(endSeq)

	return sb.String()
}

func makeStartSeq(attrs []Attr) string {
	var sb strings.Builder

	sb.WriteString("\x1b[")
	sb.WriteString(makeAttrSeq(attrs))
	sb.WriteByte('m')

	return sb.String()
}

func makeAttrSeq(attrs []Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	strAttrs := make([]string, len(attrs))

	for i, attr := range attrs {
		strAttrs[i] = strconv.Itoa(ansiCodes[attr])
	}

	filtered := slices.Compact(strAttrs)

	return strings.Join(filtered, ";")
}
