package parser

import (
	"github.com/starillume/plang/ast"
	"github.com/starillume/plang/lexer"
)

type bindingPower int

const (
	defaultBinding bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type statementHandler func(*parser) ast.Statement
type nudHandler func(*parser) ast.Expr
type ledHandler func(*parser, ast.Expr, bindingPower) ast.Expr

type statementLookup map[lexer.TokenKind]statementHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

var statements = statementLookup{}
var nuds = nudLookup{}
var leds = ledLookup{}
var bps = bpLookup{}

func led(kind lexer.TokenKind, bp bindingPower, handler ledHandler) {
	bps[kind] = bp
	leds[kind] = handler
}

func nud(kind lexer.TokenKind, bp bindingPower, handler nudHandler) {
	bps[kind] = bp
	nuds[kind] = handler
}

func statement(kind lexer.TokenKind, handler statementHandler) {
	bps[kind] = defaultBinding
	statements[kind] = handler
}

func createTokenLookups() {
	led(lexer.AND, logical, parseBinaryExpr)
	led(lexer.OR, logical, parseBinaryExpr)

	led(lexer.LESS_THAN, relational, parseBinaryExpr)
	led(lexer.LESS_OR_EQUAL, relational, parseBinaryExpr)
	led(lexer.GREATER_THAN, relational, parseBinaryExpr)
	led(lexer.GREATER_OR_EQUAL, relational, parseBinaryExpr)
	led(lexer.EQUALS, relational, parseBinaryExpr)
	led(lexer.NOT_EQUALS, relational, parseBinaryExpr)

	led(lexer.PLUS, additive, parseBinaryExpr)
	led(lexer.MINUS, additive, parseBinaryExpr)

	led(lexer.STAR, multiplicative, parseBinaryExpr)
	led(lexer.SLASH, multiplicative, parseBinaryExpr)

	nud(lexer.NUMBER, primary, parsePrimaryExpr)
	nud(lexer.STRING_LITERAL, primary, parsePrimaryExpr)
	nud(lexer.IDENTIFIER, primary, parsePrimaryExpr)
}

