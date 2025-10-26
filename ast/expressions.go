package ast

import "github.com/starillume/plang/lexer"

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (s StringExpr) expr() {}


type SymbolExpr struct {
	Value string
}

func (s SymbolExpr) expr() {}


type BinaryExpr struct {
	Left Expr
	Operator lexer.Token
	Right Expr
}

func (b BinaryExpr) expr() {}


