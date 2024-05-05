package parser

import (
	"fmt"
	"log"
)

type InterfaceDef struct {
	Token
	Operations []*OperationDef
}

func (i *InterfaceDef) node() {}

func (i *InterfaceDef) definition() {}

func ParseInterface(tok Tokenizer) (Def, error) {
	return parseInterface(tok)
}

func parseInterface(tok Tokenizer) (*InterfaceDef, error) {
	def := &InterfaceDef{
		Token:      tok.CurrentToken(),
		Operations: make([]*OperationDef, 0),
	}

	if tok.CurrentToken().Type != TypeInterface {
		return nil, fmt.Errorf("expected interface but got: %v", tok.CurrentToken())
	}

	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	if err := advanceToken(tok, TypeLeftBrace); err != nil {
		return nil, err
	}

	for tok.PeekToken().Type != TypeRightBrace {
		log.Println("Parsing operation")

		operationDef, err := parseOperation(tok)
		if err != nil {
			return nil, err
		}

		def.Operations = append(def.Operations, operationDef)
	}
	tok.NextToken()

	log.Println("Parsing return type")
	if tok.CurrentToken().Type == TypeCollon {
		if err := advanceToken(tok, TypeIdent); err != nil {
			return nil, err
		}
	}
	tok.NextToken()

	return def, nil
}

func parseOperation(tok Tokenizer) (*OperationDef, error) {
	def := &OperationDef{
		Token:  tok.PeekToken(),
		Params: make([]*ParamDef, 0),
	}

	log.Println("Parsing operation name")
	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	def.Name = &Identifier{
		Token: tok.CurrentToken(),
		Value: tok.CurrentToken().Literal,
	}

	log.Println("Parsing operation params")
	if err := advanceToken(tok, LParen); err != nil {
		return nil, err
	}

	for tok.PeekToken().Type != TypeRParen {
		log.Println("Parsing operation param")

		paramDef, err := parseParam(tok)
		if err != nil {
			return nil, err
		}

		def.Params = append(def.Params, paramDef)

		if tok.PeekToken().Type == TypeComma {
			tok.NextToken()
			log.Printf("TOKEN: %v", tok.CurrentToken())
		}
	}
	tok.NextToken()

	log.Printf("TOKEN: %v", tok.CurrentToken())

	if tok.PeekToken().Type == TypeCollon {
		log.Println("Found return type")
		tok.NextToken()

		if err := advanceToken(tok, TypeIdent); err != nil {
			return nil, err
		}

		log.Printf("Found <ident> %v", tok.CurrentToken())
		def.Return = &ReturnType{
			Token: tok.CurrentToken(),
			Value: tok.CurrentToken().Literal,
		}
	}

	log.Printf("TOKEN: %v", tok.CurrentToken())

	return def, nil
}

func parseParam(tok Tokenizer) (*ParamDef, error) {
	param := &ParamDef{
		Token: tok.PeekToken(),
	}

	log.Println("Parsing operations param name")
	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	param.Var = &Identifier{
		Token: tok.CurrentToken(),
		Value: tok.CurrentToken().Literal,
	}

	log.Println("Parsing <comma>")
	if err := advanceToken(tok, TypeCollon); err != nil {
		return nil, err
	}

	log.Println("Parsing operations param type")
	if err := advanceToken(tok, TypeIdent); err != nil {
		return nil, err
	}

	param.Type = &Identifier{
		Token: tok.CurrentToken(),
		Value: tok.CurrentToken().Literal,
	}

	return param, nil
}

func advanceToken(tokenizer Tokenizer, tt TokenType) error {
	if tok := tokenizer.PeekToken(); tok.Type != tt {
		return fmt.Errorf("expected %v but found %v", tt, tok.Literal)
	}
	tokenizer.NextToken()

	return nil
}
