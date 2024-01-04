package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestScan(t *testing.T) {
	source := `
		x = 12 * 4
		y = x + (x * 2)

		func name(a, b) {
			res = a + b
			return res
		}
	`
	tokens, err := mango.Tokenize(source)

	if err != nil {
		t.Fail()
	}
	if len(tokens) != 28 {
		t.Fail()
	}
}
