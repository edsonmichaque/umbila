package parser

import (
	"fmt"
)

type InterfaceDef struct {
	Token
	Name *Identifier
	Ops  []*OperationDef
}

func (i *InterfaceDef) node() {}

func (i *InterfaceDef) def() {}

func (p *Parser) parseInterface() (*InterfaceDef, error) {
	interfaceDef := InterfaceDef{
		Token: p.curToken,
		Ops:   []*OperationDef{},
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
