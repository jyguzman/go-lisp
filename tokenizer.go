package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Lexer struct {
	source []rune
	pos    int
	tokens []Token
}

func NewLexer(source string) *Lexer {
	runeArray := []rune(source)
	return &Lexer{runeArray, -1, []Token{}}
}

func (l *Lexer) isEof() bool {
	return l.pos >= len(l.source)
}

func (l *Lexer) advance() rune {
	l.pos += 1
	if l.pos+1 >= len(l.source) {
		return '\x00'
	}
	return l.source[l.pos]
}

func (l *Lexer) peek() rune {
	if l.isEof() {
		return '\x00'
	}
	return l.source[l.pos]
}

func (l *Lexer) peekAhead() rune {
	if l.pos+1 >= len(l.source) {
		return '\x00'
	}
	return l.source[l.pos+1]
}

func (l *Lexer) lexIdentOrKeyword() Token {
	startPos, ident := l.pos, ""
	c := l.peek()
	for unicode.IsLetter(c) {
		ident += string(c)
		c = l.advance()
	}
	var tokenType TokenType
	keywordType, isKeyword := keywords[ident]
	if isKeyword {
		tokenType = keywordType
	} else {
		tokenType = IDENT
	}
	return Token{tokenType, startPos, ident}
}

func (l *Lexer) lexString() Token {
	start_pos := l.pos
	c, str := l.peek(), ""
	for c != '"' {
		str += string(c)
		c = l.advance()
	}
	if c != '"' {
		fmt.Println("Unclosed string.")
		os.Exit(1)
	}
	l.advance()
	return Token{STRING, start_pos, str}
}

func (l *Lexer) lexNumber() Token {
	start_pos := l.pos
	numStr, isFloat := "", false
	for c := l.peek(); unicode.IsDigit(c) || c == '.'; {
		if c == '.' {
			isFloat = true
		}
		numStr += string(c)
		c = l.advance()
	}
	var (
		tokenType TokenType
		number    any
		err       error
	)
	if isFloat {
		tokenType = FLOAT
		number, err = strconv.ParseFloat(numStr, 64)
	} else {
		tokenType = INTEGER
		number, err = strconv.ParseInt(numStr, 10, 32)
	}
	if err != nil {
		fmt.Println("Error parsing number.")
		os.Exit(1)
	}
	return Token{tokenType, start_pos, number}
}

func (l *Lexer) match() Token {
	c := l.peek()
	for c == ' ' || c == '\n' || c == '\r' {
		c = l.advance()
	}
	var token Token
	switch {
	case unicode.IsLetter(c):
		return l.lexIdentOrKeyword()
	case unicode.IsDigit(c):
		return l.lexNumber()
	case c == '"':
		return l.lexString()
	case c == '(':
		token = Token{LPAREN, l.pos, "("}
	case c == ')':
		token = Token{RPAREN, l.pos, ")"}
	case c == ';':
		token = Token{SEMI, l.pos, ";"}
	case c == '+':
		token = Token{PLUS, l.pos, "+"}
	case c == '-':
		token = Token{MINUS, l.pos, "-"}
	case c == '*':
		token = Token{MINUS, l.pos, "*"}
	case c == '/':
		token = Token{DIVIDE, l.pos, "/"}
	case c == '=':
		token = Token{EQUALS, l.pos, "="}
	case c == '>':
		if l.peekAhead() == '=' {
			token = Token{GEQ, l.pos, ">="}
			l.advance()
		} else {
			token = Token{GREATER, l.pos, ">"}
		}
	case c == '<':
		if l.peekAhead() == '=' {
			token = Token{LEQ, l.pos, "<="}
			l.advance()
		} else {
			token = Token{LESS, l.pos, "<"}
		}
	default:
		token = Token{EOF, l.pos, '\x00'}
	}
	l.advance()
	return token
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}

func (l *Lexer) lex() []Token {
	for !l.isEof() {
		l.addToken(l.match())
	}
	l.addToken(Token{EOF, l.pos, '\x00'})
	return l.tokens
}

func (l *Lexer) print() {
	for i := 0; i < len(l.tokens); i++ {
		fmt.Println(l.tokens[i].str())
	}
}
