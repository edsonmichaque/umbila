package main

type Position struct {
	offset int
	line   int
	column int
}

type Lexer struct {
	src      string
	ch       byte
	offset   int
	position Position
}

func newLexer(input string) *Lexer {
	return &Lexer{
		src: input,
	}
}

func (l *Lexer) peek() byte {
	if l.offset < len(l.src) {
		return l.src[l.offset]
	}

	return 0
}

func (l *Lexer) scan() {
	if l.offset >= len(l.src) {
		l.ch = 0
		return
	}

	l.ch = l.src[l.offset]
	l.offset++
}

func (l *Lexer) readToken() Token {
	l.scan()

	l.ignoreSpace()

	switch l.ch {
	case '=':
		ch := l.peek()
		switch ch {
		case '=':
			l.scan()
			return Token{Type: Eq, Literal: "=="}
		default:
			return Token{Type: Assign, Literal: "="}
		}

	case '#':
		lit := l.readComment()
		return Token{
			Type:    Comment,
			Literal: lit,
		}
	case '@':
		return Token{Type: At, Literal: "@"}
	case ',':
		return Token{Type: Comma, Literal: ","}
	case ':':
		return Token{Type: Collon, Literal: ":"}
	case '!':
		if ch := l.peek(); ch == '=' {
			l.scan()
			return Token{Type: LTEq, Literal: "!="}
		}

		return Token{Type: Illegal}
	case '>':
		if ch := l.peek(); ch == '=' {
			l.scan()
			return Token{Type: GTEq, Literal: "!="}
		}

		return Token{Type: GThan, Literal: ">"}
	case '<':
		if ch := l.peek(); ch == '=' {
			l.scan()
			return Token{Type: Neq, Literal: "!="}
		}

		return Token{Type: LThan, Literal: "<"}
	case '(':
		return Token{Type: LParen, Literal: "("}
	case ')':
		return Token{Type: RParen, Literal: ")"}
	case '{':
		return Token{Type: LBrace, Literal: "{"}
	case '}':
		return Token{Type: RBrace, Literal: "}"}
	case '"':
		lit := l.readString()
		return Token{Type: String, Literal: lit}
	case 0:
		return Token{Type: EOF}
	default:
		if isLetter(l.ch) {
			lit := l.readIdent()
			return Token{Type: lookupKeyword(lit), Literal: lit}
		}

		if isDigit(l.ch) {
			lit := l.readNumber()
			return Token{Type: Number, Literal: lit}
		}

		return Token{Type: Illegal, Literal: string(l.ch)}
	}
}

func (l *Lexer) readString() string {
	start := l.offset - 1

	for {
		l.scan()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.src[start:l.offset]
}

func (l *Lexer) readComment() string {
	start := l.offset - 1

	for {
		ch := l.peek()

		if ch == '\n' || ch == 0 {
			break
		}

		l.scan()
	}

	return l.src[start:l.offset]
}

func (l *Lexer) ignoreSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.scan()
	}
}

func (l *Lexer) readNumber() string {
	start := l.offset - 1

	for {
		ch := l.peek()

		if !isDigit(ch) && ch != '_' {
			break
		}

		l.scan()
	}

	return l.src[start:l.offset]
}

func (l *Lexer) readIdent() string {
	start := l.offset - 1

	for {
		ch := l.peek()

		if !isLetter(ch) && !isDigit(ch) {
			break
		}

		l.scan()
	}

	return l.src[start:l.offset]
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_'
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
