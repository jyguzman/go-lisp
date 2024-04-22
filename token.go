package main

import "fmt"

type TokenType string

const (
	// punctuation
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	SEMI   TokenType = "SEMICOLON"

	// keywords
	DEF    TokenType = "DEF"
	LET    TokenType = "LET"
	LAMBDA TokenType = "LAMBDA"
	LIST   TokenType = "LIST"
	IF     TokenType = "IF"

	IDENT TokenType = "IDENT"

	// types
	TRUE    TokenType = "TRUE"
	FALSE   TokenType = "FALSE"
	INTEGER TokenType = "INTEGER"
	FLOAT   TokenType = "FLOAT"
	STRING  TokenType = "STRING"
	NIL     TokenType = "NIL"

	// operators
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	EQUALS   TokenType = "EQUALS"
	GREATER  TokenType = "GREATER"
	LESS     TokenType = "LESS"
	GEQ      TokenType = "GEQ" // >=
	LEQ      TokenType = "LEQ" // <=

	EOF TokenType = "EOF"
)

var keywords = map[string]TokenType{
	"def":    DEF,
	"let":    LET,
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
