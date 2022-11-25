package scan

import (
	"ilox/lox/token"
	"log"
	"strconv"
)

// Methods for adding tokens

func (s *Scanner) addToken(t token.TokenType) {
	s.addTokenLiteral(t, nil)
}

func (s *Scanner) addTokenLiteral(t token.TokenType, literal any) {
	text := s.source[s.start:s.current] // TODO: of by one?
	s.tokens = append(s.tokens, token.New(t, text, literal, s.start))
}

func (s *Scanner) string() {
	for c, ok := s.peek(); ok && c != '"'; c, ok = s.peek() {
		if c == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errRep.Error(s.source, s.current, "Unterminated string.")
		return
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	text := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(token.String, text)
}

func (s *Scanner) number() {
	for c, ok := s.peek(); ok && isDigit(c); c, ok = s.peek() {
		s.advance()
	}

	// Look for a fractional part.
	if c, ok := s.peek(); ok && c == '.' {
		if _, ok = s.peekTwo(); ok {
			// Consume the "."
			s.advance()

			for c, ok := s.peek(); ok && isDigit(c); c, ok = s.peek() {
				s.advance()
			}
		}
	}

	text := s.source[s.start:s.current]
	v, err := strconv.ParseFloat(text, 64)
	if err != nil {
		log.Fatal(err) // Note: in theory, we should never reach this.
	}
	s.addTokenLiteral(token.Number, v)
}

func (s *Scanner) identifier() {
	for c, ok := s.peek(); ok && isAlphaNumeric(c); c, ok = s.peek() {
		s.advance()
	}

	text := s.source[s.start:s.current]
	if t, ok := keywords[text]; ok {
		s.addToken(t)
	} else {
		s.addToken(token.Identifier)
	}
}
