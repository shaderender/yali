package util

import "strings"

type LinePoint struct {
	Text   string
	Line   int
	Column int
}

func GetLinePoint(s string, offset int) LinePoint {
	if offset >= len(s) {
		return LinePoint{}
	}

	prevNewline := strings.LastIndex(s[:offset], "\n")

	nextNewline := strings.Index(s[offset:], "\n")
	if nextNewline < 0 {
		nextNewline = len(s)
	} else {
		nextNewline += offset
	}

	line := LinePoint{
		Text:   s[prevNewline+1 : nextNewline],
		Line:   strings.Count(s[:prevNewline+1], "\n") + 1,
		Column: offset - prevNewline,
	}

	return line
}
