package ast

import "github.com/starillume/plang/lexer"

type BlockStatement struct {
	Body []Statement
}

func (b BlockStatement) statement() {}

type ExprStatement struct {
	Expr
}

func (e ExprStatement) statement() {}

type VarDeclarationStatement struct {
	Name string
	Assigned Expr
	Type lexer.TokenKind
}

func (v VarDeclarationStatement) statement() {}

type FuncDeclarationStatement struct {
	Name string
	Params []FuncParam
	Body BlockStatement
	ReturnType lexer.Token
}

func (v FuncDeclarationStatement) statement() {}

type ResturnStatement struct {
	Value Expr
}

func (v ResturnStatement) statement() {}

type FuncParam struct {
	Name string
	Type lexer.Token
}
