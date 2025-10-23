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

			if pos != nil && pos[0] == 0 {
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

func createPattern(regex *regexp.Regexp, kind TokenKind, ignore ...bool) regexPattern {
	ig := false
	if len(ignore) > 0 {
		ig = ignore[0]
	}

	handler := func (lexer *lexer) {
		match := regex.FindString(lexer.remainder())
		if !ig {
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
			createPattern(regexp.MustCompile(`^->`), ASSIGNMENT),
			createPattern(regexp.MustCompile(`^[0-9]+(\.[0-9]+)?`), NUMBER),
			createPattern(regexp.MustCompile(`^\bint\b`), INT),
			createPattern(regexp.MustCompile(`^\bfloat\b`), FLOAT),
			createPattern(regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`), IDENTIFIER),
			createPattern(regexp.MustCompile(`^;`), SEMICOLON),
			createPattern(regexp.MustCompile(`^\s+`), WHITESPACE, true),
		},
	}
}
