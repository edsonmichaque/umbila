package main

import (
	"fmt"
	"log"
)

type opDef struct {
	Token
	Value  string
	Args   []*argDef
	Return *argDef
}

func (o *opDef) node() {}

func (o *opDef) def() {}

func (p *Parser) parseOpDef() (*opDef, error) {
	log.Println("Parsing operation")
	opDef := &opDef{
		Token: p.curToken,
		Value: p.curToken.Literal,
		Args:  []*argDef{},
	}

	if p.peekToken.Type != LParen {
		return nil, fmt.Errorf("1: expected ( but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RParen {
		log.Printf("Parsing op arg")
		p.NextToken()

		if p.curToken.Type != Ident {
			return nil, fmt.Errorf("2: expected <ident> but found %v", p.curToken)
		}

		log.Printf("Found var %v", p.curToken)
		argDef := argDef{
			Token: p.curToken,
			Var:   &Identifier{Token: p.curToken, Value: p.curToken.Literal},
		}

		if p.peekToken.Type != Collon {
			return nil, fmt.Errorf("3: expected : but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("4: expected ident but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		log.Printf("Found type %v", p.curToken)

		argDef.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
		opDef.Args = append(opDef.Args, &argDef)

		if p.peekToken.Type == Comma {
			log.Println("Found <comma>")
			p.NextToken()
		}
	}
	p.NextToken()

	return opDef, nil
}
