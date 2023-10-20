package lexer

import (
	"strings"
	"testing"

	"github.com/olehvolynets/sylphy/token"
)

func TestNextToken(t *testing.T) {
	input := "#foo\t #bar"
	l := New(strings.NewReader(input))

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ATTRIBUTE, "foo"},
		{token.LITERAL, "\t "},
		{token.ATTRIBUTE, "bar"},
		{token.EOF, ""},
		{token.EOF, ""}, // to make sure it keeps returning it once reached
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong token type, expected=%d, got=%d", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong token literal, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
