package parser

import (
	"fmt"
)

func New(l *Lexer) *Parser {
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
	defs := []Def{}

	for {
		def, err := p.parseDefinition()
		if err != nil {
			return nil, err
		}

		if _, ok := def.(*End); ok {
			break
		}

		defs = append(defs, def)
	}

	return &AST{
		Defs: defs,
	}, nil
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.readToken()
}

func (p *Parser) PeekToken() Token {
	return p.peekToken
}

func (p *Parser) CurrentToken() Token {
	return p.curToken
}

type ParserFunc func(Tokenizer) (Def, error)

func (p *Parser) parseDefinition() (Def, error) {
	if p.peekToken.Type == EOF {
		return &End{}, nil
	}

	parsers := map[TokenType]ParserFunc{
		TypeInterface: ParseInterface,
		TypeAt:        ParseAnnotation,
		TypeEnum:      ParseEnum,
	}

	parseFunc, ok := parsers[p.CurrentToken().Type]
	if !ok {
		return nil, fmt.Errorf("kaboom")
	}

	return parseFunc(p)
}

type End struct{}

func (e *End) node() {}

func (e *End) definition() {}
