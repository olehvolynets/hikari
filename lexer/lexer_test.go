package lexer

import (
	"strings"
	"testing"

	"github.com/olehvolynets/sylphy/token"
)

func TestParse(t *testing.T) {
	input := "#foo\t #bar"

	results := Parse(strings.NewReader(input))

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ATTRIBUTE, "foo"},
		{token.LITERAL, "\t "},
		{token.ATTRIBUTE, "bar"},
	}

	if len(results) != len(tests) {
		t.Errorf("Result tokens: %d, expected %d\n", len(results), len(tests))
	}

	for i, tt := range tests {
		rr := results[i]

		if rr.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong token type, expected=%d, got=%d", i, tt.expectedType, rr.Type)
		}

		if rr.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong token literal, expected=%q, got=%q", i, tt.expectedLiteral, rr.Literal)
		}
	}
}
