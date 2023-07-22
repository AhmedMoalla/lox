package lexer

import (
	"unicode"
)

type Lexer interface {
	Tokenize() []Token
}

func NewLexer(source string) Lexer {
	return &loxLexer{source: source}
}

var typeByKeyword = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"fun":    Fun,
	"for":    For,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}

type loxLexer struct {
	source  string
	start   int     // Start of a token
	current int     // Current character
	tokens  []Token // Produced tokens
	line    int     // Current line
}

func (l *loxLexer) Tokenize() []Token {
	for l.current < len(l.source) {
		char := l.currentRune()
		l.start = l.current
		l.current++

		if char == '\n' {
			l.line++
		}

		switch char {
		case '(':
			l.addTokenType(LeftParen)
		case ')':
			l.addTokenType(RightParen)
		case '{':
			l.addTokenType(LeftBrace)
		case '}':
			l.addTokenType(RightBrace)
		case ',':
			l.addTokenType(Comma)
		case '.':
			l.addTokenType(Dot)
		case ';':
			l.addTokenType(Semicolon)
		case '-':
			l.addTokenType(Minus)
		case '+':
			l.addTokenType(Plus)
		case '*':
			l.addTokenType(Star)
		case '=':
			l.addTokenIfNextChar('=', EqualEqual, Equal)
		case '!':
			l.addTokenIfNextChar('=', BangEqual, Bang)
		case '>':
			l.addTokenIfNextChar('=', GreaterEqual, Greater)
		case '<':
			l.addTokenIfNextChar('=', LessEqual, Less)
		case '"':

		default:
			if unicode.IsSpace(char) {
				continue
			} else if isIdentifierCharacter(char) {
				l.handleIdentifier()
			}
		}
	}

	l.tokens = append(l.tokens, WithTokenType(EOF, len(l.source)))
	return l.tokens
}

func (l *loxLexer) addTokenIfNextChar(char rune, typeIf TokenType, typeElse TokenType) {
	if l.nextRune() == char {
		l.addTokenType(typeIf)
		l.current++
	} else {
		l.addTokenType(typeElse)
	}
}

func (l *loxLexer) addTokenType(token TokenType) {
	l.tokens = append(l.tokens, WithTokenType(token, l.start))
}

func (l *loxLexer) addToken(token Token) {
	l.tokens = append(l.tokens, token)
}

func (l *loxLexer) handleIdentifier() {
	for isIdentifierCharacter(l.currentRune()) {
		l.current++
	}
	lexeme := l.source[l.start:l.current]
	if tokenType, ok := typeByKeyword[lexeme]; ok {
		l.addTokenType(tokenType)
	} else {
		l.addToken(NewToken(Identifier, lexeme, nil, l.start))
	}
}

func (l *loxLexer) currentRune() rune {
	return rune(l.source[l.current])
}

func (l *loxLexer) nextRune() rune {
	if l.current+1 >= len(l.source) {
		return 0
	}
	return rune(l.source[l.current+1])
}

func isIdentifierCharacter(char rune) bool {
	return char == '_' || unicode.IsLetter(char)
}
