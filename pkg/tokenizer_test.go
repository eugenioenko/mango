package mango_test

import (
	mango "mango/pkg"
	"testing"
)

func TestScan(t *testing.T) {
	source := "(+ 10 1 20)"
	tokens, err := mango.Scan(source)

	if err != nil {
		t.Fail()
	}
	if len(tokens) != 7 {
		t.Fail()
	}
}
