package parser

import (
	"github.com/starillume/plang/ast"
	"github.com/starillume/plang/lexer"
)

func parseStatement(p *parser) ast.Statement {
	sHandler, ok := statements[p.currentToken().Kind]
	if ok {
		return sHandler(p)
	}

	expr := parseExpr(p, defaultBinding)
	p.expect(lexer.SEMICOLON)

	return ast.ExprStatement{Expr: expr}
}
