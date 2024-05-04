package parser

import (
	"fmt"
)

type InterfaceDefinition struct {
	Token
	Name       *Identifier
	Operations []*OperationDefinition
}

func (i *InterfaceDefinition) node() {}

func (i *InterfaceDefinition) definition() {}

func (p *Parser) parseInterface() (*InterfaceDefinition, error) {
	interfaceDefinition := InterfaceDefinition{
		Token:      p.curToken,
		Operations: []*OperationDefinition{},
	}

	if p.peekToken.Type != Ident {
		return nil, fmt.Errorf("6: expected <ident> but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	interfaceDefinition.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != LBrace {
		return nil, fmt.Errorf("5: expected ( but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RBrace {
		var (
			annotation *AnnotationDefinition
			err        error
		)

		if p.peekToken.Type == At {
			p.NextToken()

			annotation, err = p.parseAnnotationDefinition()
			if err != nil {
				return nil, err
			}
		}
		p.NextToken()

		operationDefinition, err := p.parseOperationDefinition()
		if err != nil {
			return nil, err
		}

		if annotation != nil {
			operationDefinition.Annotation = annotation
		}

		if operationDefinition != nil {
			interfaceDefinition.Operations = append(
				interfaceDefinition.Operations,
				operationDefinition,
			)
		}
	}
	p.NextToken()
	p.NextToken()

	return &interfaceDefinition, nil
}
