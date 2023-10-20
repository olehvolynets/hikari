package token

type TokenType int

const (
	LITERAL = iota
	ATTRIBUTE
)

type Token struct {
	Type    TokenType
	Literal string
}
