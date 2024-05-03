package main

import "fmt"

type TokenType int

const (
	Illegal TokenType = iota
	LParen
	Ident
	RParen
	EOF
	At
	LBrace
	RBrace
	Number
	Eq
	Assign
	Neq
	GThan
	LThan
	LTEq
	GTEq
	Comment
	Collon
	Interface
	Struct
	String
	Comma
	Enum
)

func (t TokenType) String() string {
	types := map[TokenType]string{
		Illegal:   "ILLEGAL",
		LParen:    "L_PAREN",
		Ident:     "IDENT",
		RParen:    "R_PAREN",
		EOF:       "EOF",
		At:        "AT",
		LBrace:    "L_BRACE",
		RBrace:    "R_BRACE",
		Number:    "Number",
		Eq:        "EQ",
		Assign:    "ASSIGN",
		Neq:       "NEQ",
		GThan:     "GT",
		LThan:     "LT",
		LTEq:      "LTE",
		GTEq:      "GTE",
		Comment:   "COMMENT",
		Collon:    "COLLON",
		Interface: "INTERFACE",
		Struct:    "STRUCT",
		String:    "STRING",
		Comma:     "COMMA",
		Enum:      "ENUM",
	}

	return types[t]
}

type Token struct {
	Type    TokenType
	Literal string
	start   Position
	end     Position
}

func (t Token) String() string {
	f := `Token{Type => %s, Literal => %s}`

	return fmt.Sprintf(f, t.Type, t.Literal)
}

var keywords = map[string]TokenType{
	"interface": Interface,
	"struct":    Struct,
	"enum":      Enum,
}

func checkKeyword(ident string) TokenType {
	tt, ok := keywords[ident]
	if !ok {
		return Ident
	}

	return tt
}
