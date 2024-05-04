package parser

import (
	"fmt"
	"log"
)

type AnnotationDef struct {
	Token
	Name  Identifier
	Value AnnotationValue
}

func (a *AnnotationDef) node()       {}
func (a *AnnotationDef) definition() {}

type AnnotationValue struct {
	String *string
	Number *string
	Bool   *string
	Params []*assignmentDef
}

type AnnotationParam struct {
	Token
	Name  string
	Value string
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
		case TypeIdent:

			if ann.Value.Params == nil {
				ann.Value.Params = make([]*assignmentDef, 0)
			}

			for tok.PeekToken().Type != TypeRParen {
				tok.NextToken()

				log.Println("Parsing assignment")
				def, err := parseAssignment(tok)
				if err != nil {
					return nil, err
				}

				ann.Value.Params = append(ann.Value.Params, def)

				if tok.PeekToken().Type == TypeComma {
					tok.NextToken()
				}
			}

		case String:
			tok.NextToken()

			ann.Value.String = StringPtr(tok.CurrentToken().Literal)
		case Number:
			tok.NextToken()

			ann.Value.Number = StringPtr(tok.CurrentToken().Literal)
		case Bool:
			tok.NextToken()

			ann.Value.Bool = StringPtr(tok.CurrentToken().Literal)
		default:
			return nil, fmt.Errorf("expected (<ident>|<number>|<string>) but found: %v", tok.PeekToken())
		}
		tok.NextToken()
	}
	tok.NextToken()

	return ann, nil
}

type assignmentDef struct {
	Token
	Var   string
	Value AnnotationValue
}

func (a *assignmentDef) node() {}

func (a *assignmentDef) definition() {}

func parseAssignment(tok Tokenizer) (*assignmentDef, error) {
	log.Println("Parse assignment var")
	if tok.CurrentToken().Type != TypeIdent {
		return nil, fmt.Errorf("expected <ident> but found %v", tok.CurrentToken())
	}

	def := &assignmentDef{
		Token: tok.CurrentToken(),
	}

	log.Println("Parse assignment sign")
	if err := advanceToken(tok, Assign); err != nil {
		return nil, err
	}

	switch tok.PeekToken().Type {
	case String:
		tok.NextToken()

		log.Println("Parse assignment string")
		def.Value.String = StringPtr(tok.CurrentToken().Literal)
	case Bool:
		tok.NextToken()

		log.Println("Parse assignment bool")
		def.Value.Bool = StringPtr(tok.CurrentToken().Literal)
	case Number:
		tok.NextToken()

		log.Println("Parse assignment number")
		def.Value.Number = StringPtr(tok.CurrentToken().Literal)
	default:
		return nil, fmt.Errorf("expected one of <string|bool|number> %v", tok.CurrentToken())
	}

	return def, nil
}

func StringPtr(s string) *string {
	return &s
}
