package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestItShouldEvaluate(t *testing.T) {
	source := `
		2 + 2 - 1
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 31 {
		t.Fail()
	}

}
