package lexer

import (
	"fmt"
	"slices"
)

type Keyword TokenKind

const (
	INT Keyword = "int"
	FLOAT Keyword = "float"
	STRING Keyword = "string"
	FUNC Keyword = "func"
	RETURN Keyword = "return"
)

var keywords []Keyword = []Keyword{
	INT, FLOAT, FUNC, RETURN, STRING,
}

type Operation TokenKind

const (
	ASSIGNMENT Operation = "assignment"
	PLUS Operation = "plus"
	MINUS Operation = "minus"
	STAR Operation = "star"
	SLASH Operation = "slash"
)

var operations []Operation = []Operation{
	ASSIGNMENT, MINUS, PLUS, SLASH, STAR,
}

type TokenKind string

const (
	EOF TokenKind = "eof"
	NUMBER TokenKind = "number"
	FAT_ARROW TokenKind = "fat_arrow"
	WHITESPACE TokenKind = "whitespace"
	IDENTIFIER TokenKind = "identifier"
	STRING_LITERAL TokenKind = "string_literal"
	SEMICOLON TokenKind = "semicolon"
	OPEN_BRACKET TokenKind = "open_bracket"
	CLOSE_BRACKET TokenKind = "close_bracket"
	OPEN_PAREN TokenKind = "open_paren"
	CLOSE_PAREN TokenKind = "close_paren"
	COMMA TokenKind = "comma"
	COLON TokenKind = "colon"
)

type TokenResolver func (k TokenKind, v string) (Token, bool)

var resolvers []TokenResolver = []TokenResolver{
	func (k TokenKind, v string) (Token, bool) { return tryTokenType[Keyword, KeywordToken](k, v, keywords) },
	func (k TokenKind, v string) (Token, bool) { return tryTokenType[Operation, OperationToken](k, v, operations) },
}

type Token interface {
	Debug()
	new(TokenKind, string) Token
}

type GenericToken struct {
	Kind TokenKind
	Value string
}

func (t GenericToken) Debug() {
	fmt.Printf("token kind: %s, token value: %s\n", t.Kind, t.Value)
}

func (t GenericToken) new(kind TokenKind, value string) Token {
	return GenericToken{
		kind, value,
	}
}

type KeywordToken struct {
	Kind Keyword
	Value string
}

func (t KeywordToken) Debug() {
	fmt.Printf("token kind: %s, token value: %s\n", t.Kind, t.Value)
}

func (t KeywordToken) new(kind TokenKind, value string) Token {
	return KeywordToken{
		Keyword(kind), value,
	}
}

type OperationToken struct {
	Kind Operation
	Value string
}

func (t OperationToken) new(kind TokenKind, value string) Token {
	return OperationToken{
		Operation(kind), value,
	}
}

func (t OperationToken) Debug() {
	fmt.Printf("token kind: %s, token value: %s\n", t.Kind, t.Value)
}

func tryTokenType[K ~string, T Token](kind TokenKind, value string, kinds []K) (Token, bool) {
	var t T = *new(T)
	if _, ok := slices.BinarySearch(kinds, K(kind)); ok {
		return t.new(kind, value), true
	}

	return nil, false
}

func NewToken(kind TokenKind, value string) Token {
	for _, resolver := range resolvers {
		if val, ok := resolver(kind, value); ok {
			return val
		}
	}

	return GenericToken{kind, value}
}
