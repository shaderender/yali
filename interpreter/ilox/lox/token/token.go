package token

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Offset  int
}

func New(t TokenType, lexeme string, literal any, offset int) Token {
	return Token{
		Type:    t,
		Lexeme:  lexeme,
		Literal: literal,
		Offset:  offset,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%d | %12s | %s", t.Offset, t.Type, t.Lexeme)
}
