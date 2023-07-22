package lexer

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestTokenizeSingleCharacterLexeme(t *testing.T) {
	lexemesToTokenType := map[string]TokenType{
		"(": LeftParen,
		")": RightParen,
		"{": LeftBrace,
		"}": RightBrace,
		",": Comma,
		".": Dot,
		";": Semicolon,
		"-": Minus,
		"+": Plus,
		"/": Slash,
		"*": Star,
	}

	for lexeme, tokenType := range lexemesToTokenType {
		t.Run(fmt.Sprintf("%s matches %s", lexeme, tokenType), func(t *testing.T) {
			token := NewLexer(lexeme).Tokenize()[0]
			if token.TokenType() != tokenType {
				t.Errorf("Expected lexeme %s to produce %s but got %s instead", lexeme, tokenType, token.TokenType())
			}
		})
	}
}

func TestTokenizeOneOrTwoCharactersLexeme(t *testing.T) {
	testCases := []struct {
		lexeme    string
		tokenType TokenType
	}{
		{"!", Bang},
		{"!=", BangEqual},
		{"=", Equal},
		{"==", EqualEqual},
		{">", Greater},
		{">=", GreaterEqual},
		{"<", Less},
		{"<=", LessEqual},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s matches %s", testCase.lexeme, testCase.tokenType), func(t *testing.T) {
			token := NewLexer(testCase.lexeme).Tokenize()[0]
			if token.TokenType() != testCase.tokenType {
				t.Errorf("Expected lexeme %s to produce %s but got %s instead", testCase.lexeme, testCase.tokenType, token.TokenType())
			}
		})
	}
}

func TestTokenizeString(t *testing.T) {
	source := "\"a string\""
	lexer := NewLexer(source)
	tokens := lexer.Tokenize()
	if len(tokens) != 1 {
		t.Errorf("Expected one token but got %d instead", len(tokens))
	}

	token := tokens[0]
	tokenType := token.TokenType()
	if tokenType != String {
		t.Errorf("Expected token to be of type String but got %s instead", tokenType)
	}

	literal := token.Literal()
	if literal != strings.ReplaceAll(source, "\"", "") {
		t.Errorf("Expected token literal to contain value %s but got %s instead", source, literal)
	}
}

func TestTokenizeMultilineString(t *testing.T) {
	source := "\"a string\nanother string\""
	lexer := NewLexer(source)
	tokens := lexer.Tokenize()
	if len(tokens) != 1 {
		t.Errorf("Expected one token but got %d instead", len(tokens))
	}

	token := tokens[0]
	tokenType := token.TokenType()
	if tokenType != String {
		t.Errorf("Expected token to be of type String but got %s instead", tokenType)
	}

	literal := token.Literal()
	if literal != strings.ReplaceAll(source, "\"", "") {
		t.Errorf("Expected token literal to contain value %s but got %s instead", source, literal)
	}
}

func TestVariableDeclaration(t *testing.T) {
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
			t.Errorf("Expected token [%d] to be '%s' but found '%s' instead", i, expectedToken, token)
		}
	}
}
