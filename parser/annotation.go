package parser

import (
	"fmt"
)

type AnnotationDefinition struct {
	Token
	Name        *Identifier
	StringValue string
	NumberValue string
}

func (a *AnnotationDefinition) node() {}

func (a *AnnotationDefinition) definition() {}

func (p *Parser) parseAnnotationDefinition() (*AnnotationDefinition, error) {
	def := &AnnotationDefinition{
		Token: p.curToken,
	}

	if p.peekToken.Type != TypeIdent {
		return nil, fmt.Errorf(
			"expected <ident> but found %v: %v",
			p.curToken.Type,
			p.curToken.Literal,
		)
	}
	p.NextToken()

	def.Name = &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if p.peekToken.Type == LParen {
		p.NextToken()

		if p.peekToken.Type != String {
			return nil, fmt.Errorf(
				"expected <string> but found %v: %v",
				p.peekToken.Type,
				p.peekToken.Literal,
			)
		}
		p.NextToken()

		def.StringValue = p.curToken.Literal

		if p.peekToken.Type != TypeRParen {
			return nil, fmt.Errorf("expected <rparen> but found %v", p.peekToken.Literal)
		}
		p.NextToken()
	}
	p.NextToken()

	return def, nil
}
