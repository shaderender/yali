package scan

import (
	"ilox/lox/loxerr"
)

// From current offset
func (s *Scanner) error(offset int, message string) {
	if point, ok := s.newErrorLine(offset, message); ok {
		s.errRep.Error(point, message)
	}
}

func (s *Scanner) newErrorLine(offset int, message string) (loxerr.LineError, bool) {
	lineOffset, ok := s.getLineOffset(offset)
	if !ok {
		return loxerr.LineError{}, false
	}

	line := loxerr.LineError{
		Message:      message,
		Line:         lineOffset.text,
		LineNumber:   lineOffset.line,
		ColumnNumber: lineOffset.column,
	}

	return line, true
}
