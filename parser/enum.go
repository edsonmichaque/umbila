package parser

import (
	"fmt"
	"log"
)

func ParseEnum(tok Tokenizer) (Def, error) {
	return parseEnumn(tok)
}

type EnumDef struct {
	Token
	Name   *Identifier
	Values []*Identifier
}

func (e *EnumDef) node() {}

func (e *EnumDef) definition() {}

func parseEnumn(tok Tokenizer) (*EnumDef, error) {
	log.Println("Parsing enum")
	if tok.CurrentToken().Type != TypeEnum {
		return nil, fmt.Errorf("expected <enum> but found %v", tok.CurrentToken())
	}

	log.Println("Parsing enum name")
	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	def := &EnumDef{
		Token: tok.CurrentToken(),
	}

	def.Name = &Identifier{
		Token: tok.CurrentToken(),
		Value: tok.CurrentToken().Literal,
	}

	log.Println("Parsing enum <left_brace>")
	if err := advanceToken(tok, TypeLeftBrace); err != nil {
		return nil, err
	}

	log.Println("Parsing enum values")
	for tok.PeekToken().Type != TypeRightBrace {
		log.Println("Parsing enum value")

		if err := advanceToken(tok, TypeIdent); err != nil {
			return nil, err
		}

		def.Values = append(
			def.Values,
			&Identifier{
				Token: tok.CurrentToken(),
				Value: tok.CurrentToken().Literal,
			},
		)
	}
	log.Println("Parsing <rigth_brace>", tok.CurrentToken())
	tok.NextToken()
	tok.NextToken()

	return def, nil
}
