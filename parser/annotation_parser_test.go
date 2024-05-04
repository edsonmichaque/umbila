package parser_test

import "github.com/edsonmichaque/umbila/parser"

type MockParser struct {
	tokens []parser.Token
	cursor int
}

func (m *MockParser) NextToken() {
	if m.cursor < len(m.tokens) {
		m.cursor++
	}
}

func (m *MockParser) PeekToken() parser.Token {
	if m.cursor+1 >= len(m.tokens) {
		return parser.Token{Type: parser.EOF}
	}

	return m.tokens[m.cursor+1]
}

func (m *MockParser) CurrentToken() parser.Token {
	if m.cursor >= len(m.tokens) {
		return parser.Token{Type: parser.EOF}
	}

	return m.tokens[m.cursor]
}
