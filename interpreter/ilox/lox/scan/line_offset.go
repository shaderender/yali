package scan

import "strings"

type lineOffset struct {
	text   string
	line   int
	column int
}

func (s *Scanner) getLineOffset(offset int) (lineOffset, bool) {
	if offset >= len(s.source) {
		return lineOffset{}, false
	}

	prevNewline := strings.LastIndex(s.source[:offset], "\n")

	nextNewline := strings.Index(s.source[offset:], "\n")
	if nextNewline < 0 {
		nextNewline = len(s.source)
	} else {
		nextNewline += offset
	}

	line := lineOffset{
		text:   s.source[prevNewline+1 : nextNewline],
		line:   strings.Count(s.source[:prevNewline+1], "\n") + 1,
		column: offset - (prevNewline + 1),
	}

	return line, true
}
