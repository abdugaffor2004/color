package color

import "fmt"

// Black returns the given text in black color.
func Black(text string) string {
	return Style(text, AttrFgBlack)
}

// BrightBlack returns the given text in bright black color.
func BrightBlack(text string) string {
	return Style(text, AttrFgBrightBlack)
}

// BgBlack returns the given text with black background.
func BgBlack(text string) string {
	return Style(text, AttrBgBlack)
}

// BgBrightBlack returns the given text with bright black background.
func BgBrightBlack(text string) string {
	return Style(text, AttrBgBrightBlack)
}

// Blackf formats the given string using fmt.Sprintf and returns it in black color.
func Blackf(format string, a ...interface{}) string {
	return Black(fmt.Sprintf(format, a...))
}

// Red returns the given text in red color.
func Red(text string) string {
	return Style(text, AttrFgRed)
}

// BrightRed returns the given text in bright red color.
func BrightRed(text string) string {
	return Style(text, AttrFgBrightRed)
}

// BgRed returns the given text with black background.
func BgRed(text string) string {
	return Style(text, AttrBgRed)
}

// BgBrightRed returns the given text with bright black background.
func BgBrightRed(text string) string {
	return Style(text, AttrBgBrightRed)
}

// Redf formats the given string using fmt.Sprintf and returns it in red color.
func Redf(format string, a ...interface{}) string {
	return Red(fmt.Sprintf(format, a...))
}

// Green returns the given text in green color.
func Green(text string) string {
	return Style(text, AttrFgGreen)
}

// BrightGreen returns the given text in bright green color.
func BrightGreen(text string) string {
	return Style(text, AttrFgBrightGreen)
}

// BgGreen returns the given text with black background.
func BgGreen(text string) string {
	return Style(text, AttrBgGreen)
}

// BgBrightGreen returns the given text with bright black background.
func BgBrightGreen(text string) string {
	return Style(text, AttrBgBrightGreen)
}

// Greenf formats the given string using fmt.Sprintf and returns it in green color.
func Greenf(format string, a ...interface{}) string {
	return Green(fmt.Sprintf(format, a...))
}

// Yellow returns the given text in yellow color.
func Yellow(text string) string {
	return Style(text, AttrFgYellow)
}

// BrightYellow returns the given text in bright yellow color.
func BrightYellow(text string) string {
	return Style(text, AttrFgBrightYellow)
}

// BgYellow returns the given text with black background.
func BgYellow(text string) string {
	return Style(text, AttrBgYellow)
}

// BgBrightYellow returns the given text with bright black background.
func BgBrightYellow(text string) string {
	return Style(text, AttrBgBrightYellow)
}

// Yellowf formats the given string using fmt.Sprintf and returns it in yellow color.
func Yellowf(format string, a ...interface{}) string {
	return Yellow(fmt.Sprintf(format, a...))
}

// Blue returns the given text in blue color.
func Blue(text string) string {
	return Style(text, AttrFgBlue)
}

// BrightBlue returns the given text in bright blue color.
func BrightBlue(text string) string {
	return Style(text, AttrFgBrightBlue)
}

// BgBlue returns the given text with black background.
func BgBlue(text string) string {
	return Style(text, AttrBgBlue)
}

// BgBrightBlue returns the given text with bright black background.
func BgBrightBlue(text string) string {
	return Style(text, AttrBgBrightBlue)
}

// Bluef formats the given string using fmt.Sprintf and returns it in blue color.
func Bluef(format string, a ...interface{}) string {
	return Blue(fmt.Sprintf(format, a...))
}

// Magenta returns the given text in magenta color.
func Magenta(text string) string {
	return Style(text, AttrFgMagenta)
}

// BrightMagenta returns the given text in bright magenta color.
func BrightMagenta(text string) string {
	return Style(text, AttrFgBrightMagenta)
}

// BgMagenta returns the given text with black background.
func BgMagenta(text string) string {
	return Style(text, AttrBgMagenta)
}

// BgBrightMagenta returns the given text with bright black background.
func BgBrightMagenta(text string) string {
	return Style(text, AttrBgBrightMagenta)
}

// Magentaf formats the given string using fmt.Sprintf and returns it in magenta color.
func Magentaf(format string, a ...interface{}) string {
	return Magenta(fmt.Sprintf(format, a...))
}

// Cyan returns the given text in cyan color.
func Cyan(text string) string {
	return Style(text, AttrFgCyan)
}

// BrightCyan returns the given text in bright cyan color.
func BrightCyan(text string) string {
	return Style(text, AttrFgBrightCyan)
}

// BgCyan returns the given text with black background.
func BgCyan(text string) string {
	return Style(text, AttrBgCyan)
}

// BgBrightCyan returns the given text with bright black background.
func BgBrightCyan(text string) string {
	return Style(text, AttrBgBrightCyan)
}

// Cyanf formats the given string using fmt.Sprintf and returns it in cyan color.
func Cyanf(format string, a ...interface{}) string {
	return Cyan(fmt.Sprintf(format, a...))
}

// White returns the given text in white color.
func White(text string) string {
	return Style(text, AttrFgWhite)
}

// BrightWhite returns the given text in bright white color.
func BrightWhite(text string) string {
	return Style(text, AttrFgBrightWhite)
}

// BgWhite returns the given text with black background.
func BgWhite(text string) string {
	return Style(text, AttrBgWhite)
}

// BgBrightWhite returns the given text with bright black background.
func BgBrightWhite(text string) string {
	return Style(text, AttrBgBrightWhite)
}

// Whitef formats the given string using fmt.Sprintf and returns it in white color.
func Whitef(format string, a ...interface{}) string {
	return White(fmt.Sprintf(format, a...))
}

// Bold returns the given text in bold style.
func Bold(text string) string {
	return Style(text, AttrBold)
}

// Italic returns the given text in italic style.
func Italic(text string) string {
	return Style(text, AttrItalic)
}

// Underline returns the given text underlined.
func Underline(text string) string {
	return Style(text, AttrUnderline)
}

// Dim returns the given text in dimmed style.
func Dim(text string) string {
	return Style(text, AttrDim)
}
