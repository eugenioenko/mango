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

	if err == nil {
		t.Fail()
	}

	if err.Error() != "[Syntax Error] Invalid or unexpected token: $" {
		t.Fail()
	}

}

func TestItShouldErrorOnEof(t *testing.T) {
	source := `
		1 +
	`
	tokens, _ := mango.Tokenize(source)
	_, err := mango.Parse(tokens)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "[Syntax Error] Unexpected end of file" {
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

func TestItShouldDoGrouping(t *testing.T) {
	source := `
		(100)
	`
	tokens, _ := mango.Tokenize(source)
	expressions, err := mango.Parse(tokens)

	if err != nil {
		t.Fail()
	}

	if len(expressions) != 1 {
		t.Fail()
	}

	grouping := expressions[0]
	_, ok := grouping.(*mango.ExpressionGrouping)
	if ok != true {
		t.Fail()
	}

}
