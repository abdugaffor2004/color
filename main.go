package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Attr int

type CacheKey interface {
	Attr | string
}

var (
	NoColor    bool
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

func main() {
	// Простые функции для цветов
	fmt.Println(Red("Ошибка: файл не найден"))
	fmt.Println(Green("Успешно завершено"))
	fmt.Println(Yellow("Предупреждение: устаревший метод"))
	fmt.Println(Blue("Информация: обработано 100 файлов"))

	// Яркие варианты цветов
	fmt.Println(BrightRed("КРИТИЧЕСКАЯ ОШИБКА"))
	fmt.Println(BrightGreen("ТЕСТЫ ПРОЙДЕНЫ"))
	fmt.Println(BrightYellow("ВНИМАНИЕ"))

	// Комбинирование стилей
	fmt.Println(Style("Важная ошибка", AttrFgRed, AttrBold))
	fmt.Println(Style("Успех!", AttrFgGreen, AttrBold, AttrUnderline))
	fmt.Println(Style("Предупреждение", AttrFgYellow, AttrItalic))
	fmt.Println(Style("Критично", AttrFgWhite, AttrBgRed, AttrBold))
	fmt.Println(Style("Информация", AttrFgBlack, AttrBgBrightCyan))

	arr := []string{"lol", "qwe"}
	fmt.Print(slices.Contains(arr, "qwe"))
}

func Style(text string, attrs ...Attr) string {
	if !allowColor() {
		return text
	}

	if len(attrs) == 0 {
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

func IsTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()

	return (fileInfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

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

	separator := ";"
	strAttrs := make([]string, 0, len(attrs))

	for _, attr := range attrs {
		strAttrs = append(strAttrs, strconv.Itoa(ansiCodes[attr]))
	}

	sort.Strings(strAttrs)

	return strings.Join(strAttrs, separator)
}

func checkCache(cache []string, key string) (string, bool) {
	idx := slices.Index(cache, key)

	if idx == -1 {
		return "", false
	} else {

		return cache[idx], true
	}
}
