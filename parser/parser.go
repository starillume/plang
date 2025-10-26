package parser

import (
	"fmt"

	"github.com/starillume/plang/ast"
	"github.com/starillume/plang/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos int
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	return &parser{tokens, 0}
}

func Parse(tokens []lexer.Token) ast.BlockStatement {
	body := make([]ast.Statement, 0)
	p := createParser(tokens)
	for p.hasTokens() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStatement{Body: body}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentToken().Kind != lexer.EOF
}

func (p *parser) expectError(expected lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expected {
		if err == nil {
			err = fmt.Sprintf("expected %s, but got %s", expected, kind)
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expected lexer.TokenKind) lexer.Token {
	return p.expectError(expected, nil)
}

