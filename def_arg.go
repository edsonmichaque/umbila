package main

import "fmt"

type argDef struct {
	Token
	Var  *Identifier
	Type *Identifier
}

func (a *argDef) String() string {
	f := `{
    Var   => %v
    Type  => %v
  }`

	return fmt.Sprintf(f, a.Var, a.Type)
}
