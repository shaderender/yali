package token

import "fmt"

type Token struct {
	t       TokenType
	lexeme  string
	literal any
	offset  int
}

// TODO: Do we really need New()? We'll have to see how it gets used.
func New(t TokenType, lexeme string, literal any, offset int) Token {
	return Token{
		t:       t,
		lexeme:  lexeme,
		literal: literal,
		offset:  offset,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %s", t.t, t.lexeme)
}
