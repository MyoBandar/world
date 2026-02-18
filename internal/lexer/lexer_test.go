package lexer

import "testing"

func TestNextToken_Display(t *testing.T) {
	code := "display hello"

	l := New(code)

	tests := []struct {
		expectedType  TokenType
		expectedValue string
	}{
		{DISPLAY, "display"},
		{STRING, "hello"},
		{EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("test[%d] - token value wrong. expected=%q, got=%q",
				i, tt.expectedValue, tok.Value)
		}
	}
}
