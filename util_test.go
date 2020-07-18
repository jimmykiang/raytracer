package main

import (
	"math"
	"testing"
)

func TestFloatEqual(t *testing.T) {
	a := 0.0
	b := 1.0

	pass := !floatEqual(a, b)

	if !pass {
		t.Errorf("floatEqual: %f should not equal %f", a, b)
	}

	a = 2.000001
	b = 2.000000
	pass = floatEqual(a, b)

	if !pass {
		t.Errorf("floatEqual: %f should equal %f", a, b)
	}
}

func TestAbs(t *testing.T) {
	a := 0.0
	b := 1.0

	pass := math.Abs(a-b) == math.Abs(b-a)

	if !pass {
		t.Error("Abs test failed")
	}
}
