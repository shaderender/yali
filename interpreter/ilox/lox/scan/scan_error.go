package scan

import (
	"fmt"
	"ilox/lox/util"
	"os"
)

func (s *Scanner) reportError(offset int, message string) {
	if offset < 0 || offset > len(s.source) {
		fmt.Fprintf(os.Stderr, "Scanner error called with out-of-bounds offset")
		return
	}

	point := util.GetLinePoint(s.source, offset)
	s.err.Report(point, message)
}
