package scan

import "ilox/lox/token"

var keywords map[string]token.TokenType

func init() {
	keywords = make(map[string]token.TokenType)

	keywords["and"] = token.And
	keywords["class"] = token.Class
	keywords["else"] = token.Else
	keywords["false"] = token.False
	keywords["fun"] = token.Fun
	keywords["for"] = token.For
	keywords["if"] = token.If
	keywords["nil"] = token.Nil
	keywords["or"] = token.Or
	keywords["print"] = token.Print
	keywords["return"] = token.Return
	keywords["super"] = token.Super
	keywords["this"] = token.This
	keywords["true"] = token.True
	keywords["var"] = token.Var
	keywords["while"] = token.While
}
