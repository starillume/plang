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
			createPattern(regexp.MustCompile(`^->`), TokenKind(ASSIGNMENT), false),
			createPattern(regexp.MustCompile(`^=>`), FAT_ARROW, false),
			createPattern(regexp.MustCompile(`^[0-9]+(\.[0-9]+)?`), NUMBER, false),
			createPattern(regexp.MustCompile(`^\bint\b`), TokenKind(INT), false),
			createPattern(regexp.MustCompile(`^\bfloat\b`), TokenKind(FLOAT), false),
			createPattern(regexp.MustCompile(`^\bstring\b`), TokenKind(STRING), false),
			createPattern(regexp.MustCompile(`^\bfunc\b`), TokenKind(FUNC), false),
			createPattern(regexp.MustCompile(`^\breturn\b`), TokenKind(RETURN), false),
			createPattern(regexp.MustCompile(`^"[^"]*"`), STRING_LITERAL, false),
			createPattern(regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`), IDENTIFIER, false),
			createPattern(regexp.MustCompile(`^-`), TokenKind(MINUS), false),
			createPattern(regexp.MustCompile(`^\+`), TokenKind(PLUS), false),
			createPattern(regexp.MustCompile(`^\*`), TokenKind(STAR), false),
			createPattern(regexp.MustCompile(`^/`), TokenKind(SLASH), false),
			createPattern(regexp.MustCompile(`^;`), SEMICOLON, false),
			createPattern(regexp.MustCompile(`^{`), OPEN_BRACKET, false),
			createPattern(regexp.MustCompile(`^}`), CLOSE_BRACKET, false),
			createPattern(regexp.MustCompile(`^\(`), OPEN_PAREN, false),
			createPattern(regexp.MustCompile(`^\)`), CLOSE_PAREN, false),
			createPattern(regexp.MustCompile(`^,`), COMMA, false),
			createPattern(regexp.MustCompile(`^:`), COLON, false),
			createPattern(regexp.MustCompile(`^\s+`), WHITESPACE, true),
		},
	}
}
