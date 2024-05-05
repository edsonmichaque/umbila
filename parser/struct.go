package parser

import (
	"fmt"
)

func ParseStruct(tok Tokenizer) (Def, error) {
	return parseStruct(tok)
}

type StructDefinition struct {
	Token
	Name   *Identifier
	Params []*ParamDef
}

func (s *StructDefinition) node() {}

func (s *StructDefinition) definition() {}

func parseStruct(p Tokenizer) (*StructDefinition, error) {
	if p.PeekToken().Type != TypeIdent {
		return nil, fmt.Errorf("expected <ident> found: %v", p.PeekToken().Literal)
	}
	p.NextToken()

	structDef := StructDefinition{
		Token:  p.CurrentToken(),
		Params: make([]*ParamDef, 0),
	}

	if p.PeekToken().Type != TypeLeftBrace {
		return nil, fmt.Errorf("expected '{' but found %v", p.PeekToken().Literal)
	}
	p.NextToken()

	for p.PeekToken().Type != TypeRightBrace {
		if p.PeekToken().Type != TypeIdent {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.PeekToken().Literal)
		}
		p.NextToken()

		paramDef := &ParamDef{
			Token: p.CurrentToken(),
		}
		paramDef.Var = &Identifier{
			Token: p.CurrentToken(),
			Value: p.CurrentToken().Literal,
		}

		if p.PeekToken().Type != TypeCollon {
			return nil, fmt.Errorf("expected <:> but found %v", p.PeekToken().Literal)
		}
		p.NextToken()

		if p.PeekToken().Type != TypeIdent {
			return nil, fmt.Errorf("expected <ident> but found %v", p.PeekToken().Literal)
		}
		p.NextToken()

		paramDef.Type = &Identifier{
			Token: p.CurrentToken(),
			Value: p.CurrentToken().Literal,
		}

		structDef.Params = append(structDef.Params, paramDef)
	}
	p.NextToken()
	p.NextToken()

	return &structDef, nil
}
