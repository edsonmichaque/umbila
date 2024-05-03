package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if len(os.Args) < 2 {
		panic("I need args")
	}

	input := os.Args[1]

	fmt.Println("Source:", input)

	lexer := newLexer(os.Args[1])

	parser := newParser(lexer)

	spec, err := parser.parse()
	if err != nil {
		panic(err)
	}

	spew.Printf("Spec: %v\n", spec)
}

// AST
