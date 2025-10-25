package main

import (
	"os"

	"github.com/starillume/plang/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.p")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		lexer.DebugToken(token)
	}
}
