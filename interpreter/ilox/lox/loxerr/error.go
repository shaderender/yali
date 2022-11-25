package loxerr

import (
	"fmt"
	"os"
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

func (r *Reporter) Error(e LineError, message string) {
	r.ErrorWhere(e, "", message)
}

func (r *Reporter) ErrorWhere(e LineError, where, message string) {
	if r.UseStdout {
		fmt.Println(e)
	} else {
		fmt.Fprintln(os.Stderr, e)
	}

	r.hasError = true
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
