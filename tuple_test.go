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

func TestAddTuple(t *testing.T) {
	a := Point(3.0, -2.0, 5.0)
	b := Vector(-2.0, 3.0, 1.0)

	result := a.Add(b)
	expected := Point(1.0, 1.0, 6.0)
	pass := result.Equals(expected)
	if !pass {
		t.Errorf("AddTuple: result %v should equal %v", result, expected)
	}
	a = Vector(3.0, -2.0, 5.0)
	b = Vector(-2.0, 3.0, 1.0)

	result = a.Add(b)
	expected = Vector(1.0, 1.0, 6.0)
	pass = result.Equals(expected)
	if !pass {
		t.Errorf("AddTuple: result %v should equal %v", result, expected)
	}
}
