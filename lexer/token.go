package lexer

import (
	"fmt"
)

type TokenKind string

// keywords
const (
	INT     TokenKind = "int"
	FLOAT   TokenKind = "float"
	STRING  TokenKind = "string"
	BOOLEAN TokenKind = "boolean"
	TRUE    TokenKind = "true"
	FALSE   TokenKind = "false"
	FUNC    TokenKind = "func"
	RETURN  TokenKind = "return"
	IF      TokenKind = "if"
	ELSE    TokenKind = "else"
	IMPORT  TokenKind = "import"
)

// operators
const (
	ASSIGNMENT TokenKind = "assignment"
	PLUS       TokenKind = "plus"
	MINUS      TokenKind = "minus"
	STAR       TokenKind = "star"
	SLASH      TokenKind = "slash"
)

// conditionals
const (
	EQUALS           TokenKind = "equals"
	NOT_EQUALS       TokenKind = "not_equals"
	GREATER_THAN     TokenKind = "greater_than"
	GREATER_OR_EQUAL TokenKind = "greater_or_equal"
	LESS_THAN        TokenKind = "less_than"
	LESS_OR_EQUAL    TokenKind = "less_or_equal"
	AND              TokenKind = "and"
	OR               TokenKind = "or"
	NOT              TokenKind = "not"
)

// others
const (
	EOF            TokenKind = "eof"
	NUMBER         TokenKind = "number"
	FAT_ARROW      TokenKind = "fat_arrow"
	WHITESPACE     TokenKind = "whitespace"
	IDENTIFIER     TokenKind = "identifier"
	STRING_LITERAL TokenKind = "string_literal"
	SEMICOLON      TokenKind = "semicolon"
	OPEN_BRACKET   TokenKind = "open_bracket"
	CLOSE_BRACKET  TokenKind = "close_bracket"
	OPEN_PAREN     TokenKind = "open_paren"
	CLOSE_PAREN    TokenKind = "close_paren"
	COMMA          TokenKind = "comma"
	COLON          TokenKind = "colon"
	COMMENT        TokenKind = "comment"
)

type Token struct {
	Kind TokenKind
	Value  string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}

func DebugToken(token Token) {
	fmt.Printf("Token{Kind: %s, Value: %s}\n", token.Kind, token.Value)
}
