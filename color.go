package color

import (
	"slices"
	"strconv"
	"strings"
)

const (
	startSeqPref = "\x1b["
	endSeq       = "\x1b[0m"
)

// Style applies the given attributes to the text and returns the styled string.
func Style(text string, attrs ...Attr) string {
	if !allowColor() {
		return text
	}

	if len(attrs) == 0 || text == "" {
		return text
	}

	startSeq, ok := ac.get(attrs)
	if !ok {
		startSeq = makeStartSeq(attrs)
		ac.set(attrs, startSeq)
	}

	var sb strings.Builder
	sb.WriteString(startSeq)
	sb.WriteString(text)
	sb.WriteString(endSeq)

	return sb.String()
}

func makeStartSeq(attrs []Attr) string {
	var sb strings.Builder
	sb.WriteString(startSeqPref)
	sb.WriteString(makeAttrSeq(attrs))
	sb.WriteByte('m')

	return sb.String()
}

func makeAttrSeq(attrs []Attr) string {
	if len(attrs) == 0 {
		return ""
	}

	compacted := slices.Compact(attrs)
	strAttrs := make([]string, len(compacted))

	for i, attr := range compacted {
		strAttrs[i] = strconv.Itoa(ansiCodes[attr])
	}

	return strings.Join(strAttrs, ";")
}
