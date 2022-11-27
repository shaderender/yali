package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Prints the generated file to standard output.
func main() {
	text, err := defineAst()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Println(strings.TrimSpace(text))
}

func defineAst() (string, error) {
	var b strings.Builder
	// Package.
	b.WriteString(fmt.Sprintf("package ast\n\n"))

	// Imports.
	b.WriteString(fmt.Sprintf("import \"ilox/lox/token\"\n\n"))

	// Expression interface.
	b.WriteString(fmt.Sprintf("type Expr interface {\n"))
	b.WriteString(fmt.Sprintf("\tacceptPrinter(AstPrinter) string\n"))
	b.WriteString(fmt.Sprintf("}\n\n"))

	// Type definitions.
	types := []string{
		"Binary : Left Expr, Operator token.Token, Right Expr",
		"Grouping : Expression Expr",
		"Literal : Value any",
		"Unary : Operator token.Token, Right Expr",
	}
	for _, s := range types {
		t, err := NewType(s)
		if err != nil {
			return "", err
		}
		b.WriteString(fmt.Sprintf("%s", t.Printable()))
	}

	return b.String(), nil
}

type Type struct {
	Name   string
	Fields []string
}

func NewType(s string) (Type, error) {
	components := strings.Split(s, ":")
	if len(components) != 2 {
		return Type{}, fmt.Errorf("Invalid type definition: %q\n", s)
	}

	fields := []string{}
	for _, field := range strings.Split(strings.TrimSpace(components[1]), ",") {
		fields = append(fields, strings.TrimSpace(field))
	}

	t := Type{
		Name:   strings.TrimSpace(components[0]),
		Fields: fields,
	}
	return t, nil
}

func (t Type) Printable() string {
	var b strings.Builder

	// Type definition.
	b.WriteString(fmt.Sprintf("type %s struct {\n", t.Name))
	for _, field := range t.Fields {
		b.WriteString(fmt.Sprintf("\t%s\n", field))
	}
	b.WriteString("}\n\n")

	// AST printer.
	firstChar, _ := utf8.DecodeRuneInString(t.Name)
	firstChar = unicode.ToLower(firstChar)
	b.WriteString(fmt.Sprintf("func (%c %s) acceptPrinter(a AstPrinter) string {\n", firstChar, t.Name))
	b.WriteString(fmt.Sprintf("\treturn a.visit%s(%c)\n", t.Name, firstChar))
	b.WriteString("}\n\n")

	return b.String()
}
