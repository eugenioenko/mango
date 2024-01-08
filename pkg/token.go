package mango

// Token and TokenTypes

// Token definitions
const (
	// parser tokens
	TokenTypeEof = -1

	// literals
	TokenTypeSymbol     = 0
	TokenTypeReserved   = 1
	TokenTypeIdentifier = 2
	TokenTypeNumber     = 3
	TokenTypeFloat      = 4
	TokenTypeString     = 5
)

type Token struct {
	Type    int
	Literal string
}

func MakeToken(Type int, literal string) Token {
	return Token{Type, literal}
}

var SingleCharSymbols []rune = []rune{'*', '+', '-', '!', '=', '/', '(', ')', '{', '}', '[', ']', ',', '$'}
var WhitespaceCharSymbols []rune = []rune{' ', '\t', '\r', '\n'}
var ReservedIdentifiers []string = []string{"null", "false", "true", "func", "return", "and", "or", "if", "while", "print"}
