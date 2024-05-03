package main

import (
	"fmt"
	"log"
)

type interfaceDef struct {
	Token
	Name *Identifier
	Ops  []*opDef
}

func (i *interfaceDef) String() string {
	f := `Interface{
    Token => %v
    Ident => %v
    Ops   => %v
  }`

	return fmt.Sprintf(f, i.Token, i.Name, i.Ops)
}

func (i *interfaceDef) node() {}

func (i *interfaceDef) def() {}

func (i *opDef) String() string {
	f := `Operation{
    Name   => %v
    Args   => %v
  }`

	return fmt.Sprintf(f, i.Value, i.Args)
}

func (p *Parser) parseInterface() (*interfaceDef, error) {
	log.Println("Parsing interface")
	interfaceDef := interfaceDef{
		Token: p.curToken,
		Ops:   []*opDef{},
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

	// parse operations
	for p.peekToken.Type != RBrace {
		log.Println("Iterate ops")
		p.NextToken()
		// parse operation
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
