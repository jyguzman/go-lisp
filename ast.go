package main

type ValType int

const (
	LInteger ValType = iota
	LFloat
	LString
	LBoolean
	LOperator
	LSymbol
	LSpecial
	LLambda
	LList
	LNil
)

type LispValue struct {
	Type ValType
	Val  interface{}
}

type Lambda struct {
	Params []LispValue
	Body   []LispValue
}

var operators []TokenType = []TokenType{
	PLUS, MINUS, MULTIPLY, DIVIDE, EQUALS, GREATER, GEQ, LESS, LEQ,
}

var specialForms []TokenType = []TokenType{
	IF, DEF, LAMBDA,
}

func contains(tokenTypes []TokenType, tokenType TokenType) bool {
	for _, tokType := range tokenTypes {
		if tokType == tokenType {
			return true
		}
	}
	return false
}

func isOperator(op TokenType) bool {
	return contains(operators, op)
}

func isSpecialForm(sf TokenType) bool {
	return contains(specialForms, sf)
}
