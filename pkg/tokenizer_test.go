package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestItShouldScanTokens(t *testing.T) {
	source := `
		x := 12 * 4
		y := x + (x * 2)

		func name(a, b) {
			res := a + b
			return res
		}
	`
	tokens, err := mango.Tokenize(source)

	if err != nil {
		t.Fail()
	}
	if len(tokens) != 31 {
		t.Fail()
	}
}

func TestItShouldScanStrings(t *testing.T) {
	source := `
		x := "string"
	`
	tokens, err := mango.Tokenize(source)

	if err != nil {
		t.Fail()
	}
	if len(tokens) != 4 {
		t.Fail()
	}
}

func TestItShouldReturnErrorForUnknownTokens(t *testing.T) {
	source := `
		$
	`
	_, err := mango.Tokenize(source)

	if err == nil || err.Error() != "[Scanner Error] Unexpected character: '$'" {
		t.Fail()
	}
}
