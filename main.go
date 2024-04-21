package main

import (
	"fmt"
	"os"
)

func readSourceFile(path string) string {
	b, err := os.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return ""
	}
	str := string(b) // convert content to a 'string'
	return str
}

func main() {
	source := readSourceFile("test.glisp")
	lexer := NewLexer(source)
	lexer.lex()
	lexer.print()
}
