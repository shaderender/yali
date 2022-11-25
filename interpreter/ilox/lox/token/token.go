package token

import "fmt"

type Token struct {
	t       TokenType
	lexeme  string
	literal any
	offset  int
}

func New(t TokenType, lexeme string, literal any, offset int) Token {
	return Token{
		t:       t,
		lexeme:  lexeme,
		literal: literal,
		offset:  offset,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%d | %v %s", t.offset, t.t, t.lexeme)
}
