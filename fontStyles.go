package main

func Bold(text string) string {
	return Style(text, AttrBold)
}

func Italic(text string) string {
	return Style(text, AttrItalic)
}

func Underline(text string) string {
	return Style(text, AttrUnderline)
}

func Dim(text string) string {
	return Style(text, AttrDim)
}
