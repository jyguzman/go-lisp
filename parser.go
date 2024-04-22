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

func (p *Parser) parseAtom() *Cons {
	t := p.peek()
	if t.tokenType == STRING || t.tokenType == FLOAT ||
		t.tokenType == INTEGER || t.tokenType == IDENT || t.tokenType == DEF || t.tokenType == PLUS {
		return cons(t.literal, nil)
	}
	return &Cons{nil, nil}
}

func (p *Parser) parseList() *Cons {
	t := p.advance()
	cell := p.parseExpression()
	next := cell
	for t.tokenType != RPAREN && t.tokenType != EOF {
		t = p.advance()
		if p.peek().tokenType == RPAREN {
			break
		}
		next.cdr = p.parseExpression()
		for next.cdr != nil {
			next = next.cdr
		}
	}
	return cell
}

func (p *Parser) parseExpression() *Cons {
	if p.peek().tokenType == LPAREN {
		return p.parseList()
	}
	return p.parseAtom()
}

func (p *Parser) parse() []*Cons {
	expressions := []*Cons{}
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
