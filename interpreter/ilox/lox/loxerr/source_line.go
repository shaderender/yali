package loxerr

import "strings"

type SourceLine struct {
	text   string
	line   int
	column int
}

func getSourceLine(s string, offset int) (SourceLine, bool) {
	if offset >= len(s) {
		return SourceLine{}, false
	}

	prevNewline := strings.LastIndex(s[:offset], "\n")

	nextNewline := strings.Index(s[offset:], "\n")
	if nextNewline < 0 {
		nextNewline = len(s)
	} else {
		nextNewline += offset
	}

	line := SourceLine{
		text:   s[prevNewline+1 : nextNewline],
		line:   strings.Count(s[:prevNewline+1], "\n") + 1,
		column: offset - (prevNewline + 1),
	}

	return line, true
}

func nSpaces(n int) string {
	if n < 1 {
		return ""
	}

	spaces := make([]rune, n)
	for i := 0; i < n; i++ {
		spaces[i] = ' '
	}

	return string(spaces)
}
