package parser

type Node interface {
	node()
}

type Def interface {
	Node
	definition()
}

type AST struct {
	Defs []Def
}

type Identifier struct {
	Token
	Value string
}
