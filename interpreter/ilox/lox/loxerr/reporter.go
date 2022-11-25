package loxerr

import (
	"fmt"
	"os"
)

type Reporter struct {
	hasError bool
}

func (r *Reporter) Reset() {
	r.hasError = false
}

func (r *Reporter) HasError() bool {
	return r.hasError
}

func (r *Reporter) Error(s string, offset int, message string) {
	r.ErrorWhere(s, offset, "", message)
}

func (r *Reporter) ErrorWhere(s string, offset int, where, message string) {
	line, ok := getSourceLine(s, offset)
	if !ok {
		return
	}

	fmt.Fprintf(os.Stderr, "Error: %s\n\n", message)
	marker := fmt.Sprintf("    %d | ", line.line)
	fmt.Fprintf(os.Stderr, "%s%s\n", marker, line.text)
	fmt.Fprintf(os.Stderr, "%s^-- Here\n", nSpaces(len(marker)+line.column-1))

	r.hasError = true
}
