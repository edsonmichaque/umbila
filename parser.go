package main

import (
	"fmt"
)

func newParser(l *Lexer) *Parser {
	parser := &Parser{
		lexer: l,
	}

	parser.NextToken()
	parser.NextToken()

	return parser
}

type Parser struct {
	curToken  Token
	peekToken Token
	lexer     *Lexer
}

func (p *Parser) Parse() (*AST, error) {
	definitions := []Definition{}

	for {
		definition, err := p.parseDef()
		if err != nil {
			return nil, err
		}

		if _, ok := definition.(*End); ok {
			break
		}

		definitions = append(definitions, definition)
	}

	return &AST{
		Definitions: definitions,
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
