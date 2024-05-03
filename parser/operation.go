package parser

import (
	"fmt"
)

type OperationDef struct {
	Token
	Value  string
	Params []*ParamDefinition
	Return *Token
}

func (o *OperationDef) node() {}

func (o *OperationDef) def() {}

func (p *Parser) parseOpDef() (*OperationDef, error) {
	operationDef := &OperationDef{
		Token:  p.curToken,
		Value:  p.curToken.Literal,
		Params: []*ParamDefinition{},
	}

	if p.peekToken.Type != LParen {
		return nil, fmt.Errorf("1: expected ( but found %v", p.peekToken.Literal)
	}
	p.NextToken()

	for p.peekToken.Type != RParen {
		p.NextToken()

		if p.curToken.Type != Ident {
			return nil, fmt.Errorf("2: expected <ident> but found %v", p.curToken)
		}

		argDef := ParamDefinition{
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

		argDef.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
		operationDef.Params = append(operationDef.Params, &argDef)

		if p.peekToken.Type == Comma {
			p.NextToken()
		}
	}
	p.NextToken()

	if p.peekToken.Type == Collon {
		p.NextToken()

		if p.peekToken.Type != Ident {
			return nil, fmt.Errorf("6: expected <ident> but found: %v", p.peekToken.Literal)
		}
		p.NextToken()

		operationDef.Return = &p.curToken
	}

	return operationDef, nil
}
