package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/edsonmichaque/umbila/parser"
)

func main() {
	if len(os.Args) < 2 {
		panic("I need args")
	}

	input := os.Args[1]

	fmt.Println("Source:", input)

	lexer := parser.NewLexer(os.Args[1])

	parser := parser.New(lexer)

	spec, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

// AST
