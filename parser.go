package main

type Parser struct {
	pos    int
	tokens []Token
}

func NewParser(tokens []Token) *Parser {
	return &Parser{0, tokens}
}

func (p *Parser) isEof() bool {
	return p.peek().tokenType == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.pos]
}

func (p *Parser) advance() Token {
	p.pos += 1
	return p.tokens[p.pos]
}

func (p *Parser) parseAtom() LispValue {
	t := p.peek()
	switch {
	case t.tokenType == FLOAT:
		return LispValue{LFloat, t.literal}
	case t.tokenType == INTEGER:
		return LispValue{LInteger, t.literal}
	case t.tokenType == STRING:
		return LispValue{LString, t.literal}
	case t.tokenType == LAMBDA:
		return LispValue{LLambda, t.literal}
	case t.tokenType == TRUE || t.tokenType == FALSE:
		return LispValue{LBoolean, t.literal}
	case t.tokenType == IDENT:
		return LispValue{LSymbol, t.literal}
	case isOperator(t.tokenType):
		return LispValue{LSymbol, t.literal}
	case isSpecialForm(t.tokenType):
		return LispValue{LSpecial, t.literal}
	default:
		return LispValue{LNil, nil}
	}
}

func (p *Parser) parseList() LispValue {
	t := p.advance()
	list := []LispValue{p.parseExpression()}

	for t.tokenType != RPAREN && t.tokenType != EOF {
		t = p.advance()
		if t.tokenType == RPAREN {
			break
		}
		list = append(list, p.parseExpression())
	}
	return LispValue{LList, list}
}

func (p *Parser) parseExpression() LispValue {
	if p.peek().tokenType == LPAREN {
		return p.parseList()
	}
	return p.parseAtom()
}

func (p *Parser) parse() []LispValue {
	expressions := []LispValue{}
	for !p.isEof() {
		t := p.peek()
		for t.tokenType == RPAREN {
			t = p.advance()
		}
		if t.tokenType == EOF {
			return expressions
		}
		exp := p.parseExpression()
		expressions = append(expressions, exp)
		p.advance()
	}
	return expressions
}
