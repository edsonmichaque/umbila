package parser

import "fmt"

type TokenType int

const (
	Illegal TokenType = iota
	LParen
	TypeIdent
	TypeRParen
	EOF
	TypeAt
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
	TypeCollon
	TypeInterface
	Struct
	String
	TypeComma
	Enum
	Union
)

func (t TokenType) String() string {
	types := map[TokenType]string{
		Illegal:       "ILLEGAL",
		LParen:        "L_PAREN",
		TypeIdent:     "IDENT",
		TypeRParen:    "R_PAREN",
		EOF:           "EOF",
		TypeAt:        "AT",
		LBrace:        "L_BRACE",
		RBrace:        "R_BRACE",
		Number:        "Number",
		Equal:         "EQ",
		Assign:        "ASSIGN",
		Neq:           "NEQ",
		GreaterThan:   "GT",
		LessThan:      "LT",
		NotEqual:      "NOT_EQUAL",
		GTEq:          "GTE",
		Comment:       "COMMENT",
		TypeCollon:    "COLLON",
		TypeInterface: "INTERFACE",
		Struct:        "STRUCT",
		String:        "STRING",
		TypeComma:     "COMMA",
		Enum:          "ENUM",
		Union:         "UNION",
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
	"interface": TypeInterface,
	"struct":    Struct,
	"enum":      Enum,
	"union":     Union,
}

func lookupKeyword(ident string) TokenType {
	tokenType, ok := keywords[ident]
	if !ok {
		return TypeIdent
	}

	return tokenType
}
