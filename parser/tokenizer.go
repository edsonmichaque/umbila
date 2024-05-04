package parser

type Tokenizer interface {
	NextToken()
	PeekToken() Token
	CurrentToken() Token
}
