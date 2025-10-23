package lexer

import (
	"fmt"
)

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	ASSIGNMENT
	WHITESPACE
	IDENTIFIER
	INT
	FLOAT
	SEMICOLON
	// OPEN_PAREN
	// CLOSE_PAREN
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
	default:
		return ""
	}
}
