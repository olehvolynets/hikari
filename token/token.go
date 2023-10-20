package token

type TokenType int

const (
	EOF = iota
	LITERAL
	ATTRIBUTE
)

type Token struct {
	Type    TokenType
	Literal string
}
