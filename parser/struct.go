package parser

import (
	"fmt"
)

type StructDefinition struct {
	Token
	Name   *Identifier
	Params []*ParamDef
}

func (s *StructDefinition) node() {}

func (s *StructDefinition) definition() {}

func (p *Parser) parseStruct() (*StructDefinition, error) {
	if p.peekToken.Type != TypeIdent {
		return nil, fmt.Errorf("expected <ident> found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	structDefinition := StructDefinition{
		Token:  p.curToken,
		Params: make([]*ParamDef, 0),
	}

	if p.peekToken.Type != TypeLeftBrace {
		return nil, fmt.Errorf("expected '{' but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != TypeRightBrace {
		paramDefinition := &ParamDef{
			Token: p.curToken,
		}

		if p.peekToken.Type != TypeIdent {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
		}
		p.NextToken()

		paramDefinition.Var = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekToken.Type != TypeCollon {
			return nil, fmt.Errorf("expected <:> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		if p.peekToken.Type != TypeIdent {
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
