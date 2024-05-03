package main

import (
	"fmt"
	"log"
)

type structDef struct {
	Token
	Name *Identifier
	Args []*argDef
}

func (s *structDef) node() {}

func (s *structDef) def() {}

func (p *Parser) parseStruct() (*structDef, error) {
	log.Println("Parsing struct")
	if p.peekToken.Type != Ident {
		return nil, fmt.Errorf("expected <ident> found: %v", p.peekToken.Literal)
	}
	p.NextToken()

	def := structDef{
		Token: p.curToken,
		Args:  make([]*argDef, 0),
	}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("expected '{' but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		log.Println("Parsing struct arg")
		argDef := &argDef{
			Token: p.curToken,
		}

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found: %v", p.peekToken.Literal)
		}
		p.NextToken()

		log.Println("Parsing struct arg name")
		argDef.Var = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if p.peekToken.Type != Collon {
			return nil, fmt.Errorf("expected <:> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		log.Println("Parsing struct arg type")
		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("expected <ident> but found %v", p.peekToken.Literal)
		}
		p.NextToken()

		argDef.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		def.Args = append(def.Args, argDef)
	}
	p.NextToken()
	p.NextToken()

	return &def, nil
}
