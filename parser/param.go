package parser

type ParamDef struct {
	Token
	Var  *Identifier
	Type *Identifier
}
