package parser

type OperationDef struct {
	Token
	Value      string
	Name       *Identifier
	Params     []*ParamDef
	Return     *ReturnType
	Annotation *AnnotationDef
}

type ReturnType struct {
	Token
	Value string
}

func (o *OperationDef) node() {}

func (o *OperationDef) definition() {}
