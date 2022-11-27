package scan

import (
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

// match deviates from the book in that here we do not
// implicitly advance the scanner, to avoid surpises.
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	if r != expected {
		return false
	}

	return true
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
