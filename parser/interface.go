package parser

import (
	"fmt"
)

type InterfaceDefinition struct {
	Token
	Name *Identifier
	Ops  []*OperationDefinition
}

func (i *InterfaceDefinition) node() {}

func (i *InterfaceDefinition) definition() {}

func (p *Parser) parseInterface() (*InterfaceDefinition, error) {
	interfaceDef := InterfaceDefinition{
		Token: p.curToken,
		Ops:   []*OperationDefinition{},
	}

	if p.peekToken.Type != Ident {
		return nil, fmt.Errorf("6: expected <ident> but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	interfaceDef.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("5: expected ( but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		p.NextToken()
		opDef, err := p.parseOpDef()
		if err != nil {
			return nil, err
		}

		if opDef != nil {
			interfaceDef.Ops = append(interfaceDef.Ops, opDef)
		}
	}
	p.NextToken()
	p.NextToken()

	return &interfaceDef, nil
}
