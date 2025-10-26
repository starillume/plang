package main

import (
	"fmt"
	"os"

	"github.com/starillume/plang/lexer"
	"github.com/starillume/plang/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/03.p")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	fmt.Printf("+%v\n", ast)
}
