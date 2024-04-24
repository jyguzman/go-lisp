package main

import (
	"fmt"
	"os"
)

func readSourceFile(path string) string {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return string(sourceBytes)
}

func main() {
	source := readSourceFile("test.glisp")
	lexer := NewLexer(source)
	lexer.lex()
	// lexer.print()
	parser := NewParser(lexer.tokens)
	exps := parser.parse()
	fmt.Println("\nin main")
	for _, list := range exps {
		fmt.Println(list)
	}
}
