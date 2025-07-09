package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Attr int

var (
	NoColor    bool
	ForceColor bool

	singleAttrCache = make(map[Attr]string)
	singleAttrMu    sync.RWMutex

	comboAttrCache = make(map[string]string)
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

	singleAttrMu.RLock()
	ansiSeq, ok := singleAttrCache[Attr(ansi)]
	singleAttrMu.RUnlock()

	if ok {
		return ansiSeq
	}

	var sb strings.Builder

	sb.WriteString(start)
	sb.WriteString(strconv.Itoa(ansi))
	sb.WriteByte('m')
	sb.WriteString(text)
	sb.WriteString(end)
	seq := sb.String()

	singleAttrMu.Lock()
	singleAttrCache[Attr(ansi)] = seq
	singleAttrMu.Unlock()

	return seq
}

func getComplexAnsi(text string, attrs ...Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	key := makeAttrSeq(",", attrs...)

	comboAttrMu.RLock()
	comboAnsiSeq, ok := comboAttrCache[key]
	comboAttrMu.RUnlock()

	if ok {
		return comboAnsiSeq
	}

	var sb strings.Builder

	sb.WriteString(start)
	sb.WriteString(makeAttrSeq(";", attrs...))
	sb.WriteByte('m')
	sb.WriteString(text)
	sb.WriteString(end)
	seq := sb.String()

	comboAttrMu.Lock()
	comboAttrCache[key] = seq
	comboAttrMu.Unlock()

	return seq
}

func makeAttrSeq(separator string, attrs ...Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	if separator == "" {
		separator = ","
	}

	strAttrs := make([]string, 0, len(attrs))

	for _, attr := range attrs {
		strAttrs = append(strAttrs, strconv.Itoa(ansiCodes[attr]))
	}

	sort.Strings(strAttrs)

	return strings.Join(strAttrs, separator)
}
