package loxerr

import (
	"fmt"
)

type LineError struct {
	Message      string
	Line         string
	LineNumber   int
	ColumnNumber int
}

func (e LineError) String() string {
	s := fmt.Sprintf("Error: %s\n\n", e.Message)

	marker := fmt.Sprintf("    %d | ", e.LineNumber)
	s += fmt.Sprintf("%s%s\n", marker, e.Line)

	s += fmt.Sprintf("%s^-- Here", nSpaces(len(marker)+e.ColumnNumber-1))

	return s
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
