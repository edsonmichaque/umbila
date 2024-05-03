package parser

type Node interface {
	node()
}

type Definition interface {
	Node
	definition()
}

type AST struct {
	Definitions []Definition
}

type Identifier struct {
	Token
	Value string
}
