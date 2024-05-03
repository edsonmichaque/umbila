package main

import (
	"fmt"
)

type Node interface {
	node()
}

type Definition interface {
	Node
	def()
}

type Spec struct {
	Defs []Definition
}

type Identifier struct {
	Token
	Value string
}

func newParser(l *Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}

	p.NextToken()
	p.NextToken()

	return p
}

type Parser struct {
	curToken  Token
	peekToken Token
	lexer     *Lexer
}

func (p *Parser) parse() (*Spec, error) {
	defs := []Definition{}

	for {
		def, err := p.parseDef()
		if err != nil {
			return nil, err
		}

		if _, ok := def.(*End); ok {
			break
		}

		defs = append(defs, def)
	}

	return &Spec{
		Defs: defs,
	}, nil
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.readToken()
}

func (p *Parser) PeekToken() Token {
	return Token{}
}

func (p *Parser) parseDef() (Definition, error) {
	for p.peekToken.Type != EOF {
		switch p.curToken.Type {
		case Interface:
			return p.parseInterface()
		case Struct:
			return p.parseStruct()
		case Comment:
			p.NextToken()
		}
	}

	if p.curToken.Type == EOF {
		return &End{}, nil
	}

	return nil, fmt.Errorf("7: Invalid %v", p.curToken)
}

type End struct{}

func (e *End) node() {}

func (e *End) def() {}
