package lexer

import (
	"io"

	"github.com/olehvolynets/sylphy/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(r io.Reader) *Lexer {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	l := &Lexer{
		input: string(b),
	}

	l.readChar()

	return l
}

func (p *Lexer) NextToken() token.Token {
	var tok token.Token

	switch p.ch {
	case '#':
		p.readChar() // skip #
		tok.Type = token.ATTRIBUTE
		tok.Literal = p.readIdentifier()
	case 0:
		tok.Type = token.EOF
	default:
		tok.Type = token.LITERAL
		tok.Literal = p.readLiteral()
	}

	return tok
}

func (p *Lexer) readChar() {
	p.ch = p.peekChar()
	p.position = p.readPosition
	p.readPosition += 1
}

func (p *Lexer) peekChar() byte {
	if p.readPosition >= len(p.input) {
		return 0
	} else {
		return p.input[p.readPosition]
	}
}

func (p *Lexer) readIdentifier() string {
	start := p.position

	for isIdentifierSym(p.ch) {
		p.readChar()
	}

	return p.input[start:p.position]
}

func (p *Lexer) readLiteral() string {
	start := p.position

	for p.ch != '#' && p.ch != 0 {
		p.readChar()
	}

	return p.input[start:p.position]
}

func isIdentifierSym(ch byte) bool {
	isLetter := 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
	isNumber := '0' <= ch && ch <= '9'

	return isLetter || isNumber || ch == '-' || ch == '_'
}
