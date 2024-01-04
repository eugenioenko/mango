package mango

// Token and TokenTypes

type TokenType int

// Token definitions
const (
	// parser tokens
	TokenTypeEof = -1

	// literals
	TokenTypeSymbol     = 0
	TokenTypeReserved   = 1
	TokenTypeIdentifier = 2
	TokenTypeNumber     = 3
)

type Token struct {
	Type    TokenType
	Literal string
}

func MakeToken(Type TokenType, literal string) Token {
	return Token{Type, literal}
}

var SingleCharSymbols []rune = []rune{'*', '+', '-', '!', '=', '/', '(', ')', '{', '}', '[', ']', ',', '$'}
var WhitespaceCharSymbols []rune = []rune{' ', '\t', '\r', '\n'}
var ReservedIdentifiers []string = []string{"null", "false", "true", "func", "return", "and", "or", "if", "while"}
