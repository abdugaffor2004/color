package color

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
