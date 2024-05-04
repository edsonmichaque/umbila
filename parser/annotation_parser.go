package parser

import "fmt"

type Tokenizer interface {
	NextToken()
	PeekToken() Token
	CurrentToken() Token
}

type annotationDefinition struct {
	Token
	Name  Identifier
	Value AnnotationValue
}

func (a annotationDefinition) node()       {}
func (a annotationDefinition) definition() {}

type AnnotationValue struct {
	String string
	Number string
	Bool   string
}

type annotationParserFunc func(*Tokenizer) (Definition, error)

func (a annotationParserFunc) Parse(t *Tokenizer) (Definition, error) {
	return a(t)
}

func ParseAnnotation(p Tokenizer) (Definition, error) {
	ann := &annotationDefinition{
		Token: p.CurrentToken(),
	}

	if p.CurrentToken().Type != Ident {
		return nil, fmt.Errorf("expected <ident> but found: %v", p.CurrentToken())
	}
	p.NextToken()

	if p.PeekToken().Type == LParen {
		p.NextToken()

		switch p.PeekToken().Type {
		case String:
			p.NextToken()

			ann.Value.String = p.CurrentToken().Literal
		case Number:
			p.NextToken()

			ann.Value.Number = p.CurrentToken().Literal
		default:
			return nil, fmt.Errorf("expected (<ident>|<number>|<string>) but found: %v", p.PeekToken())
		}
	}

	return ann, nil
}
