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
			token := New(lexeme).Tokenize()[0]
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
			token := New(testCase.lexeme).Tokenize()[0]
			if token.TokenType() != testCase.tokenType {
				t.Errorf("Expected lexeme %s to produce %s but got %s instead", testCase.lexeme, testCase.tokenType, token.TokenType())
			}
		})
	}
}

func TestTokenizeString(t *testing.T) {
	sources := []string{"\"a string\"", "\"a string\nanother string\""}
	for _, source := range sources {
		t.Run(fmt.Sprintf("%s matches String", source), func(t *testing.T) {
			lexer := New(source)
			tokens := lexer.Tokenize()
			if len(tokens) != 2 {
				t.Errorf("Expected two tokens but got %d instead: %v", len(tokens), tokens)
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
		})
	}
}

func TestTokenizeNumber(t *testing.T) {
	testCases := []struct {
		lexeme  string
		literal float32
	}{
		{"1", 1.0},
		{"135488", 135488.0},
		{"123.188", 123.188},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("'%s' matches Number %f", testCase.lexeme, testCase.literal), func(t *testing.T) {
			lexer := New(testCase.lexeme)
			tokens := lexer.Tokenize()
			if len(tokens) != 2 {
				t.Errorf("Expected two tokens but got %d instead: %v", len(tokens), tokens)
			}

			token := tokens[0]
			tokenType := token.TokenType()
			if tokenType != Number {
				t.Errorf("Expected token to be of type String but got %s instead", tokenType)
			}

			literal := token.Literal()
			if literal != testCase.literal {
				t.Errorf("Expected token literal to contain value %f but got %f instead", testCase.literal, literal)
			}
		})
	}
}

func TestVariableDeclaration(t *testing.T) {
	source := "// This is a comment\nvar hello = \"world\";"
	lexer := New(source)
	tokens := lexer.Tokenize()
	expectedTokens := []Token{
		WithTokenType(Var, 21),
		NewToken(Identifier, "hello", nil, 25),
		WithTokenType(Equal, 31),
		NewToken(String, "\"world\"", "world", 33),
		WithTokenType(Semicolon, 40),
		WithTokenType(EOF, 41),
	}
	for i, expectedToken := range expectedTokens {
		token := tokens[i]
		if !reflect.DeepEqual(token, expectedToken) {
			t.Errorf("Expected token [%d] to be '%s' but found '%s' instead", i, expectedToken, token)
		}
	}
}
