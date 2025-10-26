package ast

type BlockStatement struct {
	Body []Statement
}

func (b BlockStatement) statement() {}

type ExprStatement struct {
	Expr
}

func (e ExprStatement) statement() {}
