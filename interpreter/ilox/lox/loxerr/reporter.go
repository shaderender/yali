package loxerr

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
