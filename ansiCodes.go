package main

const (
	AttrFgBlack Attr = iota
	AttrFgRed
	AttrFgGreen
	AttrFgYellow
	AttrFgBlue
	AttrFgMagenta
	AttrFgCyan
	AttrFgWhite
	AttrFgBrightBlack
	AttrFgBrightRed
	AttrFgBrightGreen
	AttrFgBrightYellow
	AttrFgBrightBlue
	AttrFgBrightMagenta
	AttrFgBrightCyan
	AttrFgBrightWhite

	AttrBgBlack
	AttrBgRed
	AttrBgGreen
	AttrBgYellow
	AttrBgBlue
	AttrBgMagenta
	AttrBgCyan
	AttrBgWhite
	AttrBgBrightBlack
	AttrBgBrightRed
	AttrBgBrightGreen
	AttrBgBrightYellow
	AttrBgBrightBluecomboAttrCache
	AttrBgBrightMagenta
	AttrBgBrightCyan
	AttrBgBrightWhite

	AttrBold
	AttrDim
	AttrItalic
	AttrUnderline
)

var ansiCodes = map[Attr]int{
	AttrBold:      1,
	AttrDim:       2,
	AttrItalic:    3,
	AttrUnderline: 4,

	AttrFgBlack:         30,
	AttrFgRed:           31,
	AttrFgGreen:         32,
	AttrFgYellow:        33,
	AttrFgBlue:          34,
	AttrFgMagenta:       35,
	AttrFgCyan:          36,
	AttrFgWhite:         37,
	AttrFgBrightBlack:   90,
	AttrFgBrightRed:     91,
	AttrFgBrightGreen:   92,
	AttrFgBrightYellow:  93,
	AttrFgBrightBlue:    94,
	AttrFgBrightMagenta: 95,
	AttrFgBrightCyan:    96,
	AttrFgBrightWhite:   97,

	AttrBgBlack:                    40,
	AttrBgRed:                      41,
	AttrBgGreen:                    42,
	AttrBgYellow:                   43,
	AttrBgBlue:                     44,
	AttrBgMagenta:                  45,
	AttrBgCyan:                     46,
	AttrBgWhite:                    47,
	AttrBgBrightBlack:              100,
	AttrBgBrightRed:                101,
	AttrBgBrightGreen:              102,
	AttrBgBrightYellow:             103,
	AttrBgBrightBluecomboAttrCache: 104,
	AttrBgBrightMagenta:            105,
	AttrBgBrightCyan:               106,
	AttrBgBrightWhite:              107,
}
