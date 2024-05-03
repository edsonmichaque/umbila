package parser

import (
	"fmt"
)

type StructDefinition struct {
	Token
	Name   *Identifier
	Params []*ParamDefinition
}

func (s *StructDefinition) node() {}

func (s *StructDefinition) def() {}

func (p *Parser) parseStruct() (*StructDefinition, error) {
	if p.peekToken.Type != Ident {
		return nil, fmt.Errorf("expected <ident> found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	structDefinition := StructDefinition{
		Token:  p.curToken,
		Params: make([]*ParamDefinition, 0),
	}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("expected '{' but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		paramDefinition := &ParamDefinition{
			Token: p.curToken,
		}

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
		}
		p.NextToken()

		paramDefinition.Var = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekToken.Type != Collon {
			return nil, fmt.Errorf("expected <:> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		paramDefinition.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		structDefinition.Params = append(structDefinition.Params, paramDefinition)
	}
	p.NextToken()
	p.NextToken()

	return &structDefinition, nil
}
