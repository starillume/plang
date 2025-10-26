package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func (lexer *lexer)

type regexPattern struct {
	regex *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens []Token
	source string
	pos    int
}

func Tokenize (source string) []Token {
	lexer := newLexer(source)

	for !lexer.atEOF() {
		matched := false
		for _, pattern := range lexer.patterns {
			pos := pattern.regex.FindStringIndex(lexer.remainder())

			if pos != nil {
				pattern.handler(lexer)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("not recognized token at %s", lexer.remainder()))
		}
	}
	
	lexer.push(NewToken(EOF, "EOF"))
	return lexer.Tokens
}

func (l *lexer) advance(n int) {
	l.pos += n
}

func (l *lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *lexer) at() byte {
	return l.source[l.pos]
}

func (l *lexer) remainder() string {
	return l.source[l.pos:]
}

func (l *lexer) atEOF() bool {
	return l.pos >= len(l.source)
}

func createPattern(regex *regexp.Regexp, kind TokenKind, ignore bool) regexPattern {
	handler := func (lexer *lexer) {
		match := regex.FindString(lexer.remainder())
		if !ignore {
			lexer.push(NewToken(kind, match))
		}

		lexer.advance(len(match))
	}

	return regexPattern{handler: handler, regex: regex}
}

func newLexer(source string) *lexer {
	return &lexer{
		pos: 0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			createPattern(regexp.MustCompile(`^//.*\n`), COMMENT, true),
			createPattern(regexp.MustCompile(`^\s+`), WHITESPACE, true),
			createPattern(regexp.MustCompile(`^!=`), NOT_EQUALS, false),
			createPattern(regexp.MustCompile(`^!`), NOT, false),
			createPattern(regexp.MustCompile(`^&&`), AND, false),
			createPattern(regexp.MustCompile(`^\|\|`), OR, false),
			createPattern(regexp.MustCompile(`^==`), EQUALS, false),
			createPattern(regexp.MustCompile(`^>=`), GREATER_OR_EQUAL, false),
			createPattern(regexp.MustCompile(`^<=`), LESS_OR_EQUAL, false),
			createPattern(regexp.MustCompile(`^->`), ASSIGNMENT, false),
			createPattern(regexp.MustCompile(`^=>`), FAT_ARROW, false),
			createPattern(regexp.MustCompile(`^<`), LESS_THAN, false),
			createPattern(regexp.MustCompile(`^>`), GREATER_THAN, false),
			createPattern(regexp.MustCompile(`^\b[0-9]+(\.[0-9]+)?\b`), NUMBER, false),
			createPattern(regexp.MustCompile(`^\bint\b`), INT, false),
			createPattern(regexp.MustCompile(`^\bfloat\b`), FLOAT, false),
			createPattern(regexp.MustCompile(`^\bstring\b`), STRING, false),
			createPattern(regexp.MustCompile(`^\bboolean\b`), BOOLEAN, false),
			createPattern(regexp.MustCompile(`^\bfunc\b`), FUNC, false),
			createPattern(regexp.MustCompile(`^\breturn\b`), RETURN, false),
			createPattern(regexp.MustCompile(`^\btrue\b`), TRUE, false),
			createPattern(regexp.MustCompile(`^\bfalse\b`), FALSE, false),
			createPattern(regexp.MustCompile(`^\bif\b`), IF, false),
			createPattern(regexp.MustCompile(`^\belse\b`), ELSE, false),
			createPattern(regexp.MustCompile(`^\bimport\b`), IMPORT, false),
			createPattern(regexp.MustCompile(`^"[^"]*"`), STRING_LITERAL, false),
			createPattern(regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`), IDENTIFIER, false),
			createPattern(regexp.MustCompile(`^-`), MINUS, false),
			createPattern(regexp.MustCompile(`^\+`), PLUS, false),
			createPattern(regexp.MustCompile(`^\*`), STAR, false),
			createPattern(regexp.MustCompile(`^/`), SLASH, false),
			createPattern(regexp.MustCompile(`^;`), SEMICOLON, false),
			createPattern(regexp.MustCompile(`^{`), OPEN_BRACKET, false),
			createPattern(regexp.MustCompile(`^}`), CLOSE_BRACKET, false),
			createPattern(regexp.MustCompile(`^\(`), OPEN_PAREN, false),
			createPattern(regexp.MustCompile(`^\)`), CLOSE_PAREN, false),
			createPattern(regexp.MustCompile(`^,`), COMMA, false),
			createPattern(regexp.MustCompile(`^:`), COLON, false),
		},
	}
}
