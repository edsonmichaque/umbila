package parser

type ParamDefinition struct {
	Token
	Var  *Identifier
	Type *Identifier
}
