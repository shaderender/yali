package scan

import (
	"ilox/lox/token"
	"unicode"
	"unicode/utf8"
)

// Methods for querying around current rune.

func isDigit(c rune) bool {
	// TODO: Are there other valid numbers in UTF-8?
	return c >= '0' && c <= '9'
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

func isAlphaNumeric(c rune) bool {
	return isDigit(c) || isAlpha(c)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	r, size := utf8.DecodeRuneInString(s.source[s.current:])
	if r != expected {
		return false
	}

	s.current += size
	return true
}

// If the next rune matches expected, return a, else return b.
func (s *Scanner) matchElse(expected rune, a, b token.TokenType) token.TokenType {
	if s.match(expected) {
		return a
	}
	return b
}

func (s *Scanner) advance() (rune, bool) {
	if s.isAtEnd() {
		return 0, false
	}
	// Handles UTF-8.
	r, size := utf8.DecodeRuneInString(s.source[s.current:])
	s.current += size
	return r, true
}

func (s *Scanner) peek() (rune, bool) {
	if s.isAtEnd() {
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return r, true
}

func (s *Scanner) peekTwo() (rune, bool) {
	if s.isAtEnd() {
		return 0, false
	}

	for i, r := range s.source[s.current:] {
		if i == 1 {
			return r, true
		}
	}
	return 0, false
}
