package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestItShouldParseBinary(t *testing.T) {
	source := `
		1 + 2 - 5 / 2
	`
	tokens, _ := mango.Tokenize(source)
	_, err := mango.Parse(tokens)

	if err != nil {
		t.Fail()
	}

}

func TestItShouldReturnEmptyOnNoTokens(t *testing.T) {
	_, err := mango.Parse(nil)

	if err != nil {
		t.Fail()
	}

}

func TestItShouldErrorOnUnknown(t *testing.T) {
	source := `
		1 $ 2
	`
	tokens, _ := mango.Tokenize(source)
	_, err := mango.Parse(tokens)

	if err != nil {
		t.Fail()
	}

}

func TestItShouldErrorOnEof(t *testing.T) {
	source := `
		1 +
	`
	tokens, _ := mango.Tokenize(source)
	_, err := mango.Parse(tokens)

	if err != nil {
		t.Fail()
	}

}
