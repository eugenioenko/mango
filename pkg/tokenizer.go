package mango

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// Tokenizer
// Converts a source file into an array of tokens
type Tokenizer struct {
	source  []byte
	current int
	start   int
	tokens  []Token
}

func Scan(source string) ([]Token, error) {
	tokenizer := MakeTokenizer()
	tokenizer.LoadFromString(source)
	return tokenizer.Tokenize()
}

func MakeTokenizer() Tokenizer {
	return Tokenizer{}
}

func (tokenizer *Tokenizer) LoadFromString(source string) {
	tokenizer.source = []byte(source)
	tokenizer.tokens = make([]Token, 0)
}

func (tokenizer *Tokenizer) Eof() bool {
	return tokenizer.current >= len(tokenizer.source)
}

func (tokenizer *Tokenizer) Advance() byte {
	var current byte = tokenizer.source[tokenizer.current]
	tokenizer.current += 1
	return current
}

func (tokenizer *Tokenizer) Match(expected byte) bool {
	if tokenizer.Eof() {
		return false
	}
	if tokenizer.source[tokenizer.current] != expected {
		return false
	}
	tokenizer.current += 1
	return true
}

func (tokenizer *Tokenizer) Peek() rune {
	if tokenizer.Eof() {
		return 0
	}
	return rune(tokenizer.source[tokenizer.current])
}

func (tokenizer *Tokenizer) PeekNext() rune {
	if tokenizer.current+1 >= len(tokenizer.source) {
		return 0
	}
	return rune(tokenizer.source[tokenizer.current+1])
}

func (tokenizer *Tokenizer) AddToken(Type TokenType, literal string) {
	tokenizer.tokens = append(tokenizer.tokens, MakeToken(Type, literal))
}

func (tokenizer *Tokenizer) Error(errorMessage string) {
	fmt.Println("[Scan Error] " + errorMessage)
	os.Exit(1)
}

func (tokenizer *Tokenizer) Tokenize() ([]Token, error) {
	tokenizer.current = 0
	tokenizer.start = 0

	for !tokenizer.Eof() {
		tokenizer.start = tokenizer.current
		tokenizer.ScanToken()
	}
	tokenizer.AddToken(TokenTypeEof, "")
	return tokenizer.tokens, nil
}

func (tokenizer *Tokenizer) Identifier() {
	for unicode.IsLetter(tokenizer.Peek()) ||
		unicode.IsDigit(tokenizer.Peek()) ||
		tokenizer.Peek() == '-' ||
		tokenizer.Peek() == '_' {
		tokenizer.Advance()
	}
	token := string(tokenizer.source[tokenizer.start:tokenizer.current])
	tokenizer.AddToken(TokenTypeIdentifier, token)
}

func (tokenizer *Tokenizer) Number() {

	for unicode.IsDigit(tokenizer.Peek()) {
		tokenizer.Advance()
	}

	if tokenizer.Match('.') && unicode.IsDigit(tokenizer.Peek()) {
		for unicode.IsDigit(tokenizer.Peek()) {
			tokenizer.Advance()
		}
	}

	tokenizer.AddToken(TokenTypeNumber, string(tokenizer.source[tokenizer.start:tokenizer.current]))
}

func (tokenizer *Tokenizer) twoChar(char rune) bool {
	return (char == '!' && tokenizer.Match('=')) ||
		(char == '=' && tokenizer.Match('=')) ||
		(char == ':' && tokenizer.Match('=')) ||
		(char == '/' && tokenizer.Match('='))
}

func (tokenizer *Tokenizer) oneChar(char rune) bool {
	return char == '*' || char == '+' || char == '-' || char == '!' || char == '=' || char == '/'
}

func (tokenizer *Tokenizer) ignoreChar(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func (tokenizer *Tokenizer) ScanToken() (bool, error) {
	var err string
	char := rune(tokenizer.Advance())

	switch {
	case char == '(':
		tokenizer.AddToken(TokenTypeLeftParen, "(")
	case char == ')':
		tokenizer.AddToken(TokenTypeRightParen, ")")
	case tokenizer.twoChar(char):
		tokenizer.AddToken(TokenTypeOperator, string(char)+"=")
	case tokenizer.oneChar(char):
		tokenizer.AddToken(TokenTypeOperator, string(char))
	case unicode.IsDigit(char):
		tokenizer.Number()
	case unicode.IsLetter(char):
		tokenizer.Identifier()
	case tokenizer.ignoreChar(char):
		break
	default:
		err = fmt.Sprintf("[Tokenizer] Unexpected character:  %q", char)
	}
	if err != "" {
		return false, errors.New(err)
	}
	return true, nil
}
