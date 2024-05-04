package parser

import (
	"fmt"
	"log"
)

type Tokenizer interface {
	NextToken()
	PeekToken() Token
	CurrentToken() Token
}

type AnnotationDef struct {
	Token
	Name  Identifier
	Value AnnotationValue
}

func (a AnnotationDef) node()       {}
func (a AnnotationDef) definition() {}

type AnnotationValue struct {
	String *string
	Number *string
	Bool   *string
}

func ParseAnnotation(p Tokenizer) (Def, error) {
	return parseAnnotation(p)
}

func parseAnnotation(tok Tokenizer) (*AnnotationDef, error) {
	ann := &AnnotationDef{
		Token: tok.CurrentToken(),
	}

	log.Println("Parsing annotation")
	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	ann.Name = Identifier{
		Token: tok.CurrentToken(),
	}

	if tok.PeekToken().Type == LParen {
		tok.NextToken()

		switch tok.PeekToken().Type {
		case String:
			tok.NextToken()

			ann.Value.String = StringPtr(tok.CurrentToken().Literal)
		case Number:
			tok.NextToken()

			ann.Value.Number = StringPtr(tok.CurrentToken().Literal)
		default:
			return nil, fmt.Errorf("expected (<ident>|<number>|<string>) but found: %v", tok.PeekToken())
		}
		tok.NextToken()
	}

	tok.NextToken()

	return ann, nil
}

func StringPtr(s string) *string {
	return &s
}
