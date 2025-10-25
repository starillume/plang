package lexer

import (
	"fmt"
	"slices"
)

const (
	INT TokenKind = "int"
	FLOAT TokenKind = "float"
	STRING TokenKind = "string"
	BOOLEAN TokenKind = "boolean"
	TRUE TokenKind = "true"
	FALSE TokenKind = "false"
	FUNC TokenKind = "func"
	RETURN TokenKind = "return"
	IF TokenKind = "if"
	ELSE TokenKind = "else"
	IMPORT TokenKind = "import"
)

var keywords []TokenKind = []TokenKind{
	BOOLEAN, ELSE, IF, IMPORT, INT, FLOAT, FUNC, RETURN, STRING,
}


const (
	ASSIGNMENT TokenKind = "assignment"
	PLUS TokenKind = "plus"
	MINUS TokenKind = "minus"
	STAR TokenKind = "star"
	SLASH TokenKind = "slash"
)

var operations []TokenKind = []TokenKind{
	ASSIGNMENT, MINUS, PLUS, SLASH, STAR,
}


const (
	EQUALS TokenKind = "equals"
	NOT_EQUALS TokenKind = "not_equals"
	GREATER_THAN TokenKind = "greater_than"
	GREATER_OR_EQUAL TokenKind = "greater_or_equal"
	LESS_THAN TokenKind = "less_than"
	LESS_OR_EQUAL TokenKind = "less_or_equal"
	AND TokenKind = "and"
	OR TokenKind = "or"
	NOT TokenKind = "not"
)

var conditionals []TokenKind = []TokenKind{
	AND, EQUALS, GREATER_OR_EQUAL, GREATER_THAN, LESS_OR_EQUAL, LESS_THAN, NOT, NOT_EQUALS, OR,
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
	COMMENT TokenKind = "comment"
)

type TokenResolver func (k TokenKind, v string) (Token, bool)

var resolvers []TokenResolver = []TokenResolver{
	func (k TokenKind, v string) (Token, bool) { return tryTokenType[KeywordToken](k, v, keywords) },
	func (k TokenKind, v string) (Token, bool) { return tryTokenType[OperationToken](k, v, operations) },
	func (k TokenKind, v string) (Token, bool) { return tryTokenType[ConditionalToken](k, v, conditionals) },
}

type Token interface {
	new(TokenKind, string) Token
}

type GenericToken struct {
	Kind TokenKind
	Value string
}

func (t GenericToken) new(kind TokenKind, value string) Token {
	return GenericToken{
		kind, value,
	}
}

type KeywordToken struct {
	Kind TokenKind
	Value string
}

func (t KeywordToken) new(kind TokenKind, value string) Token {
	return KeywordToken{
		kind, value,
	}
}

type OperationToken struct {
	Kind TokenKind
	Value string
}

func (t OperationToken) new(kind TokenKind, value string) Token {
	return OperationToken{
		kind, value,
	}
}

type ConditionalToken struct {
	Kind TokenKind
	Value string
}

func (t ConditionalToken) new(kind TokenKind, value string) Token {
	return ConditionalToken{
		kind, value,
	}
}


func tryTokenType[T Token](kind TokenKind, value string, kinds []TokenKind) (Token, bool) {
	var t T = *new(T)
	if _, ok := slices.BinarySearch(kinds, kind); ok {
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

func DebugToken(token Token) {
	fmt.Printf("Token%+v\n", token)
}
