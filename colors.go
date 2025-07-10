package main

import (
	"fmt"
)

func Black(text string) string {
	return Style(text, AttrFgBlack)
}
func BrightBlack(text string) string {
	return Style(text, AttrFgBrightBlack)
}
func Blackf(format string, a ...interface{}) string {
	return Black(fmt.Sprintf(format, a...))
}

func Red(text string) string {
	return Style(text, AttrFgRed)
}
func BrightRed(text string) string {
	return Style(text, AttrFgBrightRed)
}
func Redf(format string, a ...interface{}) string {
	return Red(fmt.Sprintf(format, a...))
}

func Green(text string) string {
	return Style(text, AttrFgGreen)
}
func BrightGreen(text string) string {
	return Style(text, AttrFgBrightGreen)
}
func Greenf(format string, a ...interface{}) string {
	return Green(fmt.Sprintf(format, a...))
}

func Yellow(text string) string {
	return Style(text, AttrFgYellow)
}
func BrightYellow(text string) string {
	return Style(text, AttrFgBrightYellow)
}
func Yellowf(format string, a ...interface{}) string {
	return Yellow(fmt.Sprintf(format, a...))
}

func Blue(text string) string {
	return Style(text, AttrFgBlue)
}
func BrightBlue(text string) string {
	return Style(text, AttrFgBrightBlue)
}
func Bluef(format string, a ...interface{}) string {
	return Blue(fmt.Sprintf(format, a...))
}

func Magenta(text string) string {
	return Style(text, AttrFgMagenta)
}
func BrightMagenta(text string) string {
	return Style(text, AttrFgBrightMagenta)
}
func Magentaf(format string, a ...interface{}) string {
	return Magenta(fmt.Sprintf(format, a...))
}

func Cyan(text string) string {
	return Style(text, AttrFgCyan)
}
func BrightCyan(text string) string {
	return Style(text, AttrFgBrightCyan)
}
func Cyanf(format string, a ...interface{}) string {
	return Cyan(fmt.Sprintf(format, a...))
}

func White(text string) string {
	return Style(text, AttrFgWhite)
}
func BrightWhite(text string) string {
	return Style(text, AttrFgBrightWhite)
}
func Whitef(format string, a ...interface{}) string {
	return White(fmt.Sprintf(format, a...))
}
