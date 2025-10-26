package parser

import (
	"fmt"
	"strconv"

	"github.com/starillume/plang/ast"
	"github.com/starillume/plang/lexer"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	tokenKind := p.currentToken().Kind
	nHandler, ok := nuds[tokenKind]
	if !ok {
		panic(fmt.Sprintf("nud handler expected for token %s\n", tokenKind))
	}

	left := nHandler(p)
	for bps[p.currentToken().Kind] > bp {
		tokenKind = p.currentToken().Kind
		lHandler, ok := leds[tokenKind]
		if !ok {
			panic(fmt.Sprintf("led handler expected for token %s\n", tokenKind))
		}

		left = lHandler(p, left, bp)
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentToken().Kind {
	case lexer.NUMBER:
		n, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{Value: n}
	case lexer.STRING_LITERAL:
		return ast.StringExpr{Value: p.advance().Value}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{Value: p.advance().Value}
	default:
		panic(fmt.Sprintf("cannot create primary expr from %s\n", p.currentToken().Kind))
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operatorToken := p.advance()
	right := parseExpr(p, bp)

	return ast.BinaryExpr{
		Left: left,
		Operator: operatorToken,
		Right: right,
	}
}
