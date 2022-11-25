package scan

import (
	"fmt"
	"ilox/lox/token"
)

func (s *Scanner) scanToken() {
	c, ok := s.advance()
	if !ok {
		return
	}

	switch c {
	// Single-character tokens. (Excluding Slash.)
	case '(':
		s.addToken(token.LeftParen)
	case ')':
		s.addToken(token.RightParen)
	case '{':
		s.addToken(token.LeftBrace)
	case '}':
		s.addToken(token.RightBrace)
	case ',':
		s.addToken(token.Comma)
	case '.':
		s.addToken(token.Dot)
	case '-':
		s.addToken(token.Minus)
	case '+':
		s.addToken(token.Plus)
	case ';':
		s.addToken(token.Semicolon)
	case '*':
		s.addToken(token.Star)
	// Operators
	case '!':
		s.addToken(s.matchElse('=', token.BangEqual, token.Bang))
	case '=':
		s.addToken(s.matchElse('=', token.EqualEqual, token.Equal))
	case '<':
		s.addToken(s.matchElse('=', token.LessEqual, token.Less))
	case '>':
		s.addToken(s.matchElse('=', token.GreaterEqual, token.Greater))
	// Division
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for r, ok := s.peek(); ok && r != '\n'; r, ok = s.peek() {
				s.advance()
			}
			s.addToken(token.Comment)
		} else {
			s.addToken(token.Slash)
		}
	// Whitespace
	case ' ', '\r', '\t':
		// Ignore.
	case '\n':
		// Preserve newlines to support formatting tool.
		s.addToken(token.Newline)
		s.line++
	// String literals
	case '"':
		s.string()
	default:
		// TODO: Group multiple unexpected characters.
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.errRep.Error(s.source, s.current, fmt.Sprintf("Unexpected character %q.", c))
		}
	}
}
