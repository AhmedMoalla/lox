package lexer

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	source := "var hello = \"world\";"
	lexer := NewLexer(source)
	tokens := lexer.Tokenize()
	expectedTokens := []Token{
		WithTokenType(Var, 0),
		NewToken(Identifier, "hello", nil, 4),
		WithTokenType(Equal, 10),
		NewToken(String, "\"world\"", "world", 12),
		WithTokenType(Comma, 19),
		WithTokenType(EOF, 20),
	}
	for i, expectedToken := range expectedTokens {
		token := tokens[i]
		if !reflect.DeepEqual(token, expectedToken) {
			t.Errorf("Expected token [%d] to be '%+v' but found '%+v' instead", i, expectedToken, token)
		}
	}
}
