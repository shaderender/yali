package ast

import "ilox/lox/token"

type Expr interface {
	acceptPrinter(AstPrinter) string
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b Binary) acceptPrinter(a AstPrinter) string {
	return a.visitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (g Grouping) acceptPrinter(a AstPrinter) string {
	return a.visitGrouping(g)
}

type Literal struct {
	Value any
}

func (l Literal) acceptPrinter(a AstPrinter) string {
	return a.visitLiteral(l)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (u Unary) acceptPrinter(a AstPrinter) string {
	return a.visitUnary(u)
}
