package loxerr

import (
	"fmt"
	"ilox/lox/util"
	"os"
)

type Reporter struct {
	UseStdout bool
	hasError  bool
}

func (r *Reporter) Reset() {
	r.hasError = false
}

func (r *Reporter) HasError() bool {
	return r.hasError
}

func (r *Reporter) Report(p util.LinePoint, message string) {
	r.ReportWhere(p, "", message)
}

func (r *Reporter) ReportWhere(p util.LinePoint, where, message string) {
	if r.UseStdout {
		fmt.Println(errorToString(p, where, message))
	} else {
		fmt.Fprintln(os.Stderr, errorToString(p, where, message))
	}

	r.hasError = true
}

func errorToString(p util.LinePoint, where, message string) string {
	s := fmt.Sprintf("Error: %s\n\n", message)

	marker := fmt.Sprintf("    %d | ", p.Line)
	s += fmt.Sprintf("%s%s\n", marker, p.Text)

	s += fmt.Sprintf("%s^-- Here", nSpaces(len(marker)+p.Column-1))
	s += "\n"

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
