package loxerr

import (
	"fmt"
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
