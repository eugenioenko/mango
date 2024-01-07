package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestItShouldEvaluate(t *testing.T) {
	source := `
		5 + 90 * 120 + 200 / 2 * 1 + 100  + 3 * 3 + 191 - 3  / 3
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}

	if result[0].ToInteger() != 11204 {
		t.Fail()
	}

}

func TestItAssign(t *testing.T) {
	source := `
		var = 5 + 90
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}
}
