package lexer

import (
	"fmt"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	ASSIGNMENT
	ARROW
	WHITESPACE
	IDENTIFIER
	INT
	FLOAT
	SEMICOLON
	PLUS
	MINUS
	STAR
	SLASH
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	COMMA
	COLON
	FUNC
	RETURN
	// STRING
	// ...
)

type Token struct {
	Kind TokenKind
	Value string
}

func (t Token) Debug() {
	fmt.Printf("token kind: %s, token value: %s\n", TokenKindString(t.Kind), t.Value)
}

func NewToken(kind TokenKind, value string) Token {
	return Token{Kind: kind, Value: value}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	case NUMBER:
		return "number"
	case ASSIGNMENT:
		return "assignment"
	case WHITESPACE:
		return "whitespace"
	case IDENTIFIER:
		return "identifier"
	case INT:
		return "int"
	case FLOAT:
		return "float"
	case SEMICOLON:
		return "semicolon"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case STAR:
		return "star"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case COMMA:
		return "comma"
	case COLON:
		return "colon"
	case FUNC:
		return "func"
	case RETURN:
		return "return"
	case ARROW:
		return "arrow"
	case SLASH:
		return "slash"
	default:
		return ""
	}
}
