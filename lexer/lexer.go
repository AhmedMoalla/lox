package lexer

import (
	"github.com/AhmedMoalla/lox/errors"
	"strconv"
	"unicode"
)

type Lexer interface {
	Tokenize() []Token
}

func New(source string) Lexer {
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
	for !l.endOfSource() {
		char := l.currentRune()
		l.start = l.current

		switch char {
		case '\n':
			l.line++
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
		case '/':
			if l.current+1 < len(l.source) && l.source[l.current+1] == '/' {
				l.handleComments()
			} else {
				l.addTokenType(Slash)
			}
		case '"':
			l.handleString()
		default:
			if unicode.IsSpace(char) {
				l.current++
				continue
			} else if unicode.IsDigit(char) {
				l.handleNumber()
			} else if isIdentifierCharacter(char) {
				l.handleIdentifier()
			} else {
				errors.Error(l.line, "Unexpected character "+string(char))
			}
		}
		l.current++
	}
	l.tokens = append(l.tokens, WithTokenType(EOF, len(l.source)))
	return l.tokens
}

func (l *loxLexer) addTokenIfNextChar(char rune, typeIf TokenType, typeElse TokenType) {
	if l.current+1 < len(l.source) && rune(l.source[l.current+1]) == char {
		l.addTokenType(typeIf)
		l.current++
	} else {
		l.addTokenType(typeElse)
	}
}

func (l *loxLexer) handleIdentifier() {
	for !l.endOfSource() && isIdentifierCharacter(l.currentRune()) {
		l.current++
	}
	lexeme := l.source[l.start:l.current]
	if tokenType, ok := typeByKeyword[lexeme]; ok {
		l.addTokenType(tokenType)
	} else {
		l.addToken(NewToken(Identifier, lexeme, nil, l.start))
	}
}

func (l *loxLexer) handleString() {
	l.current++ // Advance cursor to take into account the opening "
	for !l.endOfSource() && l.currentRune() != '"' {
		if l.currentRune() == '\n' {
			l.line++
		}
		l.current++
	}

	if l.endOfSource() {
		errors.Error(l.line, "unterminated string")
		return
	}

	l.addToken(NewToken(String, l.source[l.start:l.current+1], l.source[l.start+1:l.current], l.start))
}

func (l *loxLexer) handleNumber() {
	for !l.endOfSource() && unicode.IsDigit(l.currentRune()) {
		l.current++
	}

	if !l.endOfSource() {
		// Floating point number
		if l.currentRune() == '.' {
			l.current++
			for !l.endOfSource() && unicode.IsDigit(l.currentRune()) {
				l.current++
			}
		}
	}

	literal, err := strconv.ParseFloat(l.source[l.start:l.current], 32)
	if err != nil {
		errors.Error(l.line, "unable to parse number "+l.source[l.start:l.current])
	}
	l.addToken(NewToken(Number, l.source[l.start:l.current], float32(literal), l.start))

	l.current-- // Go back to the last non digit character
}

func (l *loxLexer) handleComments() {
	l.current++ // Advance cursor to take into account the second slash
	for !l.endOfSource() && rune(l.source[l.current+1]) != '\n' {
		l.current++
	}
}

func (l *loxLexer) addTokenType(token TokenType) {
	l.tokens = append(l.tokens, WithTokenType(token, l.start))
}

func (l *loxLexer) addToken(token Token) {
	l.tokens = append(l.tokens, token)
}

func (l *loxLexer) currentRune() rune {
	return rune(l.source[l.current])
}

func isIdentifierCharacter(char rune) bool {
	return char == '_' || unicode.IsLetter(char)
}

func (l *loxLexer) endOfSource() bool {
	return l.current >= len(l.source)
}
