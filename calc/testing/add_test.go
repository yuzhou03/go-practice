package simplemath

import (
	"testing"
	"git.tvblack.com/pratice/calc/simplemath"
)

func TestAdd(t *testing.T) {
	r := simplemath.Add(2, 4)
	expected := 6
	if r != expected {
		t.Errorf("Add(2,4) failed. Got %d, expected %d", r, expected)
	}
}
