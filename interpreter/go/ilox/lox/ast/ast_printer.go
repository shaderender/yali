package ast

import (
	"fmt"
	"ilox/lox/token"
	"strings"
)

func TestAstPrinter() {
	expr := Binary{
		Left: Unary{
			Operator: token.New(token.Minus, "-", nil, 1),
			Right:    Literal{123},
		},
		Operator: token.New(token.Star, "*", nil, 1),
		Right:    Grouping{Literal{45.67}},
	}

	printer := AstPrinter{}
	fmt.Println(printer.String(expr))
}

type AstPrinter struct {
}

func (a AstPrinter) String(expr Expr) string {
	return expr.acceptPrinter(a)
}

// visitors
func (a AstPrinter) visitBinary(expr Binary) string {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a AstPrinter) visitGrouping(expr Grouping) string {
	return a.parenthesize("group", expr.Expression)
}

func (a AstPrinter) visitLiteral(expr Literal) string {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (a AstPrinter) visitUnary(expr Unary) string {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

// Utility

func (a AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(name)

	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.acceptPrinter(a))
	}

	builder.WriteString(")")

	return builder.String()
}
