package parser

type Node interface {
	node()
}

type Definition interface {
	Node
	def()
}

type AST struct {
	Definitions []Definition
}

type Identifier struct {
	Token
	Value string
}
