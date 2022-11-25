package scan

import "unicode"

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
