package main

import (
	"fmt"
)

type StructDef struct {
	Token
	Name   *Identifier
	Params []*ParamDef
}

func (s *StructDef) node() {}

func (s *StructDef) def() {}

func (p *Parser) parseStruct() (*StructDef, error) {
	if p.peekToken.Type != Ident {
		return nil, fmt.Errorf("expected <ident> found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	def := StructDef{
		Token:  p.curToken,
		Params: make([]*ParamDef, 0),
	}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("expected '{' but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		argDef := &ParamDef{
			Token: p.curToken,
		}

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
		}
		p.NextToken()

		argDef.Var = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekToken.Type != Collon {
			return nil, fmt.Errorf("expected <:> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		argDef.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		def.Params = append(def.Params, argDef)
	}
	p.NextToken()
	p.NextToken()

	return &def, nil
}
