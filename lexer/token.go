package lexer

import (
	"fmt"
	"strconv"
)

type Token interface {
	TokenType() TokenType
}

func WithTokenType(tokenType TokenType, location int) Token {
	return NewToken(tokenType, "", nil, location)
}

func NewToken(tokenType TokenType, lexeme string, literal any, location int) Token {
	return loxToken{lexeme, tokenType, literal, location}
}

type loxToken struct {
	lexeme    string
	tokenType TokenType
	literal   any
	location  int
}

func (t loxToken) String() string {
	if t.tokenType == Identifier || t.tokenType == String || t.tokenType == Number {
		return fmt.Sprintf("@%d %s(%s)", t.location, t.tokenType, t.lexeme)
	}

	return fmt.Sprintf("@%d %s", t.location, t.tokenType)
}

func (t loxToken) TokenType() TokenType {
	return t.tokenType
}

type TokenType int

const (
	EOF TokenType = iota

	// Single-character tokens.
	LeftParen
	RightParen

	LeftBrace
	RightBrace

	Comma
	Dot
	Semicolon

	Minus
	Plus
	Slash
	Star

	// One or two character tokens.
	Bang
	BangEqual

	Equal
	EqualEqual

	Greater
	GreaterEqual

	Less
	LessEqual

	// Literals.
	Identifier
	String
	Number

	// Keywords.
	And
	Or

	True
	False

	If
	Else

	For
	While

	Class
	Super
	This

	Fun
	Nil
	Var

	Print
	Return
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"
	case Bang:
		return "Bang"
	case BangEqual:
		return "BangEqual"
	case Equal:
		return "Equal"
	case EqualEqual:
		return "EqualEqual"
	case Greater:
		return "Greater"
	case GreaterEqual:
		return "GreaterEqual"
	case Less:
		return "Less"
	case LessEqual:
		return "LessEqual"
	case Identifier:
		return "Identifier"
	case String:
		return "String"
	case Number:
		return "Number"
	case And:
		return "And"
	case Class:
		return "Class"
	case Else:
		return "Else"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case For:
		return "For"
	case If:
		return "If"
	case Nil:
		return "Nil"
	case Or:
		return "Or"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Super:
		return "Super"
	case This:
		return "This"
	case True:
		return "True"
	case Var:
		return "Var"
	case While:
		return "While"
	default:
		panic("Unhandled TokenType " + strconv.Itoa(int(t)))
	}
}
