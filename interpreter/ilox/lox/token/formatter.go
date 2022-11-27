package token

import (
	"strings"
)

func Format(tokens []Token) string {
	lines := []string{}

	line := ""
	prev := Newline
	indent := 0
	for _, token := range tokens {
		switch token.t {
		case LeftBrace:
			line += " " + token.lexeme
			// TODO: Consolodate below
			line = nTabs(indent) + strings.TrimSpace(line)
			lines = append(lines, line)
			line = ""
			// TODO: Consolodate above
			indent++
		case Newline:
			line = nTabs(indent) + strings.TrimSpace(line)
			lines = append(lines, line)
			line = ""
		case RightBrace:
			indent--
			fallthrough
		default:
			line += prefix(prev, token.t) + token.lexeme + suffix(prev, token.t)
		}
		prev = token.t
	}

	return strings.Join(lines, "\n")
}

func prefix(prev, t TokenType) string {
	switch t {
	case Plus, Slash:
		return " "
	case Star:
		return " "
	case Equal, EqualEqual, Greater, GreaterEqual, Less, LessEqual:
		return " "
	case And, Or, If, Else:
		return " "
	case Comment:
		if prev != Newline {
			return " "
		}
	}
	return ""
}

func suffix(prev, t TokenType) string {
	switch t {
	case Plus, Slash:
		return " "
	case Star:
		return " "
	case Equal, EqualEqual, Greater, GreaterEqual, Less, LessEqual:
		return " "
	case And, Or, If, Else:
		return " "
	case For, While:
		return " "
	case Var, Fun, Class, Print:
		return " "
	}
	return ""
}

// TODO: Find a more efficient method.
func nTabs(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "\t"
	}
	return s
}
