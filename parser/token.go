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
	TypeLeftBrace
	TypeRightBrace
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
	TypeEnum
	Union
	Bool
)

func (t TokenType) String() string {
	types := map[TokenType]string{
		Illegal:        "ILLEGAL",
		LParen:         "L_PAREN",
		TypeIdent:      "IDENT",
		TypeRParen:     "R_PAREN",
		EOF:            "EOF",
		TypeAt:         "AT",
		TypeLeftBrace:  "L_BRACE",
		TypeRightBrace: "R_BRACE",
		Number:         "Number",
		Equal:          "EQ",
		Assign:         "ASSIGN",
		Neq:            "NEQ",
		GreaterThan:    "GT",
		LessThan:       "LT",
		NotEqual:       "NOT_EQUAL",
		GTEq:           "GTE",
		Comment:        "COMMENT",
		TypeCollon:     "COLLON",
		TypeInterface:  "INTERFACE",
		Struct:         "STRUCT",
		String:         "STRING",
		TypeComma:      "COMMA",
		TypeEnum:       "ENUM",
		Union:          "UNION",
		Bool:           "BOOL",
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
	"enum":      TypeEnum,
	"union":     Union,
	"false":     Bool,
	"true":      Bool,
}

func lookupKeyword(ident string) TokenType {
	tokenType, ok := keywords[ident]
	if !ok {
		return TypeIdent
	}

	return tokenType
}
