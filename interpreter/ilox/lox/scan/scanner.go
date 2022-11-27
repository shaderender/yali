package scan

import (
	"fmt"
	"ilox/lox/loxerr"
	"ilox/lox/token"
	"log"
	"strconv"
)

type Scanner struct {
	err    *loxerr.Reporter
	source string
	tokens []token.Token

	start   int
	current int
	line    int
	column  int
}

func NewScanner(src string, errRep *loxerr.Reporter) Scanner {
	return Scanner{
		err:    errRep,
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
		s.addStringToken()
	default:
		// TODO: Group multiple unexpected characters.
		if isDigit(c) {
			s.addNumberToken()
		} else if isAlpha(c) {
			s.addIdentifierToken()
		} else {
			s.reportError(s.current, fmt.Sprintf("Unexpected character %q.", c))
		}
	}
}

func (s *Scanner) addToken(t token.TokenType) {
	s.addTokenLiteral(t, nil)
}

func (s *Scanner) addTokenLiteral(t token.TokenType, literal any) {
	text := s.source[s.start:s.current] // TODO: of by one?
	s.tokens = append(s.tokens, token.New(t, text, literal, s.start))
}

func (s *Scanner) addStringToken() {
	for c, ok := s.peek(); ok && c != '"'; c, ok = s.peek() {
		if c == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.reportError(s.start, "Unterminated string.")
		return
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	text := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(token.String, text)
}

func (s *Scanner) addNumberToken() {
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

func (s *Scanner) addIdentifierToken() {
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
