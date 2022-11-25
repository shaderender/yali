package scan

import (
	"ilox/lox/loxerr"
	"ilox/lox/token"
)

type Scanner struct {
	errRep *loxerr.Reporter
	source string
	tokens []token.Token

	start   int
	current int
	line    int
	column  int
}

func NewScanner(src string, errRep *loxerr.Reporter) Scanner {
	return Scanner{
		errRep: errRep,
		source: src,
		tokens: []token.Token{},
		line:   1,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.New(token.Eof, "", nil, s.start))
	return s.tokens
}
