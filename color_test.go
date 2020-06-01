package main

import "testing"

func TestAddColor(t *testing.T) {
	a := NewColor(0.9, 0.6, 0.75)
	b := NewColor(0.7, 0.1, 0.25)
	result := a.Add(b)
	expected := NewColor(1.6, 0.7, 1.0)

	pass := result.Equals(expected)

	if !pass {
		t.Errorf("AddColor: result %v should equal %v", result, expected)
	}
}

func TestSubColor(t *testing.T) {
	a := NewColor(0.9, 0.6, 0.75)
	b := NewColor(0.7, 0.1, 0.25)
	result := a.Subtract(b)
	expected := NewColor(0.2, 0.5, 0.5)

	pass := result.Equals(expected)

	if !pass {
		t.Errorf("SubColor: result %v should equal %v", result, expected)
	}
}

func TestMultiplybyScalarColor(t *testing.T) {
	a := NewColor(0.2, 0.3, 0.4)
	scalar := 2.0
	result := a.MultiplyByScalar(scalar)
	expected := NewColor(0.4, 0.6, 0.8)

	pass := result.Equals(expected)

	if !pass {
		t.Errorf("MulScalarColor: result %v should equal %v", result, expected)
	}
}
