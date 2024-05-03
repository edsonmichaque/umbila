package parser

type Position struct {
	Offset int
	Line   int
	Column int
}

type Lexer struct {
	Src      string
	Ch       byte
	Offset   int
	position Position
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Src: input,
	}
}

func (l *Lexer) peek() byte {
	if l.Offset < len(l.Src) {
		return l.Src[l.Offset]
	}

	return 0
}

func (l *Lexer) scan() {
	if l.Offset >= len(l.Src) {
		l.Ch = 0
		return
	}

	l.Ch = l.Src[l.Offset]
	l.Offset++
}

func (l *Lexer) readToken() Token {
	l.scan()

	l.ignoreSpace()

	switch l.Ch {
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
		if isLetter(l.Ch) {
			lit := l.readIdent()
			return Token{Type: lookupKeyword(lit), Literal: lit}
		}

		if isDigit(l.Ch) {
			lit := l.readNumber()
			return Token{Type: Number, Literal: lit}
		}

		return Token{Type: Illegal, Literal: string(l.Ch)}
	}
}

func (l *Lexer) readString() string {
	start := l.Offset - 1

	for {
		l.scan()

		if l.Ch == '"' || l.Ch == 0 {
			break
		}
	}

	return l.Src[start:l.Offset]
}

func (l *Lexer) readComment() string {
	start := l.Offset - 1

	for {
		ch := l.peek()

		if ch == '\n' || ch == 0 {
			break
		}

		l.scan()
	}

	return l.Src[start:l.Offset]
}

func (l *Lexer) ignoreSpace() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\r' || l.Ch == '\n' {
		l.scan()
	}
}

func (l *Lexer) readNumber() string {
	start := l.Offset - 1

	for {
		ch := l.peek()

		if !isDigit(ch) && ch != '_' {
			break
		}

		l.scan()
	}

	return l.Src[start:l.Offset]
}

func (l *Lexer) readIdent() string {
	start := l.Offset - 1

	for {
		ch := l.peek()

		if !isLetter(ch) && !isDigit(ch) {
			break
		}

		l.scan()
	}

	return l.Src[start:l.Offset]
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_'
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
