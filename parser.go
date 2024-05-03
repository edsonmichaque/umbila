package main

import (
	"fmt"
	"log"
	"time"
)

type Node interface {
	node()
}

type Def interface {
	Node
	def()
}

type Spec struct {
	Defs []Def
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
	defs := []Def{}

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

func (p *Parser) parseDef() (Def, error) {
	for p.peekToken.Type != EOF {
		log.Printf("CUR TOKEN: %v", p.curToken)
		time.Sleep(500 * time.Millisecond)

		switch p.curToken.Type {
		case Interface:
			return p.parseInterface()
		case Struct:
			return p.parseStruct()
		case Comment:
			log.Println("Found a comment")
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
