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

func TestItShouldDoAssignment(t *testing.T) {
	source := `
		variable = 100
	`
	tokens, _ := mango.Tokenize(source)
	expressions, err := mango.Parse(tokens)

	if err != nil {
		t.Fail()
	}

	if len(expressions) != 1 {
		t.Fail()
	}

	assignment := expressions[0]
	_, ok := assignment.(*mango.ExpressionAssign)
	if ok != true {
		t.Fail()
	}

}
