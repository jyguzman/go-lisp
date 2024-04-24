package main

import "fmt"

type TokenType string

const (
	// punctuation
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	SEMI   TokenType = "SEMICOLON"

	// special forms
	DEF    TokenType = "DEF"
	IF     TokenType = "IF"
	LAMBDA TokenType = "LAMBDA"
	LIST   TokenType = "LIST"

	// types
	TRUE    TokenType = "TRUE"
	FALSE   TokenType = "FALSE"
	INTEGER TokenType = "INTEGER"
	FLOAT   TokenType = "FLOAT"
	STRING  TokenType = "STRING"
	NIL     TokenType = "NIL"

	// symbols
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	EQUALS   TokenType = "EQUALS"
	GREATER  TokenType = "GREATER"
	LESS     TokenType = "LESS"
	GEQ      TokenType = "GEQ" // >=
	LEQ      TokenType = "LEQ" // <=
	IDENT    TokenType = "IDENT"

	EOF TokenType = "EOF"
)

var keywords = map[string]TokenType{
	"def":    DEF,
	"lambda": LAMBDA,
	"list":   LIST,
	"if":     IF,
}

type Token struct {
	tokenType TokenType
	pos       int
	literal   any
}

func (t *Token) str() string {
	return fmt.Sprintf("Token(%s, %d, %v)", t.tokenType, t.pos, t.literal)
}
