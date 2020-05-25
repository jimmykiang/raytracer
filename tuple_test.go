package main

import "testing"

func TestTupleEqual(t *testing.T) {
	a := &Tuple{1.000001, 2.000002, 3.000000, 0.0}
	b := &Tuple{1.000000, 2.000001, 2.999999, 0.0}

	pass := a.Equals(b)
	if !pass {
		t.Errorf("TupleEqual: %v should equal %v", a, b)
	}
	a.w = 1.0

	pass = !a.Equals(b)
	if !pass {
		t.Errorf("TupleEqual: %v should not equal %v", a, b)
	}
}
