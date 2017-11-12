package simplemath

import (
	"testing"
	"git.tvblack.com/pratice/calc/simplemath"
)

func TestSqrt(t *testing.T) {
	r := simplemath.Sqrt(16)
	expected := float64(4)
	if r != expected {
		t.Errorf("Sqrt(16) failed. Got %d, expected %d", r, expected)
	}
}
