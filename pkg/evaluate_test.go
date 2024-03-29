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
		var := 5 + 90
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}
}

func TestItGroups(t *testing.T) {
	source := `
		(5 + 7)
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}

	if result[0].ToInteger() != 12 {
		t.Fail()
	}
}

func TestItPrints(t *testing.T) {
	source := `
		print 2
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}

	if result[0].ToInteger() != 2 {
		t.Fail()
	}
}
func TestItAssignVars(t *testing.T) {
	source := `
		a := 5 + 90
		b := a + 5
		print b
	`
	result, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}

	if result[2].ToInteger() != 100 {
		t.Fail()
	}
}

func TestItDoesBlock(t *testing.T) {
	source := `
		{
			a := 5 + 90
			b := a + 5
		}
	`
	_, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}
}

func TestItDoesIf(t *testing.T) {
	source := `
		if 2 == 2 {
			print "yes"
		} else {
			print "what"
		}
		if 2 == 1 {
			print "no"
		} else {
			print "yes"
		}
		if 3 != 1 {
			print "yes"
		}
	`
	_, err := mango.Eval(source)

	if err != nil {
		t.Fail()
	}
}
