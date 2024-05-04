package parser

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
	Equal
	Assign
	Neq
	GreaterThan
	LessThan
	NotEqual
	GTEq
	Comment
	Collon
	Interface
	Struct
	String
	Comma
	Enum
	Union
)

func (t TokenType) String() string {
	types := map[TokenType]string{
		Illegal:     "ILLEGAL",
		LParen:      "L_PAREN",
		Ident:       "IDENT",
		RParen:      "R_PAREN",
		EOF:         "EOF",
		At:          "AT",
		LBrace:      "L_BRACE",
		RBrace:      "R_BRACE",
		Number:      "Number",
		Equal:       "EQ",
		Assign:      "ASSIGN",
		Neq:         "NEQ",
		GreaterThan: "GT",
		LessThan:    "LT",
		NotEqual:    "NOT_EQUAL",
		GTEq:        "GTE",
		Comment:     "COMMENT",
		Collon:      "COLLON",
		Interface:   "INTERFACE",
		Struct:      "STRUCT",
		String:      "STRING",
		Comma:       "COMMA",
		Enum:        "ENUM",
		Union:       "UNION",
	}

	return types[t]
}

func (t TokenType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v"`, t.String())), nil
}

type Token struct {
	Type    TokenType
	Literal string
	start   Position
	end     Position
}

var keywords = map[string]TokenType{
	"interface": Interface,
	"struct":    Struct,
	"enum":      Enum,
	"union":     Union,
}

func lookupKeyword(ident string) TokenType {
	tokenType, ok := keywords[ident]
	if !ok {
		return Ident
	}

	return tokenType
}
