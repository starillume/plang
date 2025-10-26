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

func parseFuncDeclarationStatement(p *parser) ast.Statement {
	p.expect(lexer.FUNC)

	name := p.expectError(lexer.IDENTIFIER, "invalid function declaration").Value
	p.expect(lexer.ASSIGNMENT)

	params := make([]ast.FuncParam, 0)

	current := p.currentToken().Kind
	if current == lexer.IDENTIFIER {
		params = parseFuncParams(p)
	}

	// NOTE: como limitar?
	returnType := p.advance()

	body := ast.BlockStatement{Body: make([]ast.Statement, 0)}

	isInline := p.currentToken().Kind == lexer.FAT_ARROW
	if isInline {
		p.advance()
		expr := parseExpr(p, assignment)
		exprStatment := ast.ExprStatement{Expr: expr}
		body.Body = append(body.Body, exprStatment)
		p.expect(lexer.SEMICOLON)
	} else {
		// NOTE: parse function statement?
		funcBody := parseFuncBody(p)
		body.Body = append(body.Body, funcBody...)
	}

	return ast.FuncDeclarationStatement{Name: name, Params: params, Body: body, ReturnType: returnType}
}

func parseFuncParams(p *parser) []ast.FuncParam {
	params := make([]ast.FuncParam, 0)
	current := p.currentToken()
	for current.Kind != lexer.COLON {
		name := p.expectError(lexer.IDENTIFIER, "invalid parameter declaration").Value
		// NOTE: como limitar?
		typ := p.advance()

		param := ast.FuncParam{Name: name, Type: typ}
		params = append(params, param)

		current = p.currentToken()
		if current.Kind != lexer.COLON {
			p.expect(lexer.COMMA)
		}
	}

	p.advance()

	return params
}

func parseFuncBody(p *parser) []ast.Statement {
	p.expect(lexer.OPEN_BRACKET)
	body := make([]ast.Statement, 0)
	current := p.currentToken()
	for current.Kind != lexer.CLOSE_BRACKET {
		sfunc := parseStatement(p)
		body = append(body, sfunc)
		current = p.currentToken()
	}

	p.expect(lexer.CLOSE_BRACKET)

	return body
}

func parseReturnStatement(p *parser) ast.Statement {
	p.expect(lexer.RETURN)

	returnExpr := parseExpr(p, defaultBinding)

	p.expect(lexer.SEMICOLON)

	return ast.ResturnStatement{Value: returnExpr}
}
