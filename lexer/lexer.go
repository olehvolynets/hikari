package lexer

import (
	"errors"
	"io"

	"github.com/olehvolynets/sylphy/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

var ErrEOF = errors.New("end of format-string")

func NewLexer(r io.Reader) *Lexer {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return &Lexer{
		input: string(b),
	}
}

func Parse(r io.Reader) []token.Token {
	fp := NewLexer(r)
	fp.readChar()

	tokens := make([]token.Token, 0)
	for {
		t, err := fp.NextToken()
		if err != nil {
			if errors.Is(err, ErrEOF) {
				break
			}

			panic(err)
		}

		tokens = append(tokens, t)
	}

	return tokens
}

func (p *Lexer) NextToken() (token.Token, error) {
	var tok token.Token

	switch p.ch {
	case '#':
		p.readChar()
		attrName := p.readIdentifier()
		tok = token.Token{Type: token.ATTRIBUTE, Literal: attrName}
	case 0:
		return tok, ErrEOF
	default:
		lit := p.readLiteral()
		tok = token.Token{Type: token.LITERAL, Literal: lit}
	}

	return tok, nil
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
