package main

import (
	"fmt"
	"os"

	"github.com/starillume/plang/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.p")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		switch token := token.(type) {
		case lexer.OperationToken:
			fmt.Println("operation: ", token.Value)
		case lexer.KeywordToken:
			fmt.Println("keyword: ", token.Value)
		}
	}
}
