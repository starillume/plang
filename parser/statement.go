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

func parseVarDeclarationStatement(p *parser) ast.Statement {
	var assigned ast.Expr

	typ := p.advance().Kind
	name := p.expectError(lexer.IDENTIFIER, "invalid variable declaration").Value
	if p.currentToken().Kind != lexer.SEMICOLON {
		p.expect(lexer.ASSIGNMENT)
		assigned = parseExpr(p, assignment)
	}
	p.expect(lexer.SEMICOLON)

	return ast.VarDeclarationStatement{Name: name, Assigned: assigned, Type: typ}
}
