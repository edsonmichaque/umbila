package parser

import "fmt"

type EnumDefinition struct {
	Token
	Name     *Identifier
	Variants []*Identifier
}

func (e *EnumDefinition) node() {}

func (e *EnumDefinition) definition() {}

func (p *Parser) parseEnumDefinition() (*EnumDefinition, error) {
	enumDef := &EnumDefinition{
		Token: p.curToken,
	}
	if p.peekToken.Type != TypeIdent {
		return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	enumDef.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("expected <lbrace> but found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		p.NextToken()

		if p.curToken.Type != TypeIdent {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
		}

		enumDef.Variants = append(
			enumDef.Variants,
			&Identifier{Token: p.curToken, Value: p.curToken.Literal},
		)
	}
	p.NextToken()
	p.NextToken()

	return enumDef, nil
}
