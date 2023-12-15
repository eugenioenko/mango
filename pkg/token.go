package mango

// Token and TokenTypes

type TokenType int

// Token definitions
const (
	// parser tokens
	TokenTypeEof  = -1
	TokenTypeNull = 0

	// single character tokens
	TokenTypeLeftBrace    = 1
	TokenTypeLeftBracket  = 2
	TokenTypeLeftParen    = 3
	TokenTypeRightBrace   = 5
	TokenTypeRightBracket = 6
	TokenTypeRightParen   = 7

	// literals
	TokenTypeReserved   = 30
	TokenTypeIdentifier = 31
	TokenTypeBoolean    = 34
	TokenTypeTrue       = 35
	TokenTypeFalse      = 36
	TokenTypeNumber     = 37
	TokenTypeOperator   = 39
)

type Token struct {
	Type    TokenType
	Literal string
}

func MakeToken(Type TokenType, literal string) Token {
	return Token{Type, literal}
}
