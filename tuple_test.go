package main

import (
	"math"
	"testing"
)

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

func TestSubstractTuple(t *testing.T) {
	a := Point(3.0, 2.0, 1.0)
	b := Point(5.0, 6.0, 7.0)

	result := a.Substract(b)
	expected := Vector(-2.0, -4.0, -6.0)

	pass := result.Equals(expected)

	if !pass {
		t.Errorf("Substract Tuple: result %v should equal %v", result, expected)
	}

	a = Point(3.0, 2.0, 1.0)
	b = Vector(5.0, 6.0, 7.0)

	result = a.Substract(b)
	expected = Point(-2.0, -4.0, -6.0)

	pass = result.Equals(expected)

	if !pass {
		t.Errorf("Substract Tuple: result %v should equal %v", result, expected)
	}

	a = Vector(3.0, 2.0, 1.0)
	b = Vector(5.0, 6.0, 7.0)

	result = a.Substract(b)
	expected = Vector(-2.0, -4.0, -6.0)

	pass = result.Equals(expected)

	if !pass {
		t.Errorf("Substract Tuple: result %v should equal %v", result, expected)
	}
}

func TestNegateTuple(t *testing.T) {
	vector := &Tuple{1, -2, 3, 4}
	result := vector.Negate()
	expected := &Tuple{-1, 2, -3, -4}

	pass := result.Equals(expected)
	if !pass {
		t.Errorf("NegateTuple: result %v should equal %v", result, expected)
	}
}

func TestMultiplyTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	result := a.Multiply(3.5)
	expected := &Tuple{3.5, -7, 10.5, -14}

	pass := result.Equals(expected)
	if !pass {
		t.Errorf("Muliply Tuple: result %v should equal %v", result, expected)
	}

	result = a.Multiply(0.5)
	expected = &Tuple{0.5, -1, 1.5, -2}

	pass = result.Equals(expected)
	if !pass {
		t.Errorf("Muliply Tuple: result %v should equal %v", result, expected)
	}
}

func TestDivideTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	result := a.Divide(2)
	expected := &Tuple{0.5, -1, 1.5, -2}

	pass := result.Equals(expected)
	if !pass {
		t.Errorf("DivTuple: result %v should equal %v", result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	vector := Vector(0, 1, 0)
	result := vector.Magnitude()
	expected := 1.0
	pass := floatEqual(result, expected)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", result, expected)
	}

	vector = Vector(1, 0, 0)
	result = vector.Magnitude()
	pass = floatEqual(result, expected)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", result, expected)
	}

	vector = Vector(0, 0, 1)
	result = vector.Magnitude()
	pass = floatEqual(result, expected)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", result, expected)
	}

	vector = Vector(-1, -2, -3)
	result = vector.Magnitude()
	expected = math.Sqrt(14.0)
	pass = floatEqual(result, expected)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", result, expected)
	}
}

func TestNormalize(t *testing.T) {
	vector := Vector(4, 0, 0)
	result := vector.Normalize()
	expected := Vector(1, 0, 0)

	pass := result.Equals(expected)
	if !pass {
		t.Errorf("Normalize: result %v should equal %v", result, expected)
	}

	vector = Vector(1, 2, 3)
	result = vector.Normalize()
	expected = Vector(0.26726, 0.53452, 0.80178)
	pass = result.Equals(expected)
	if !pass {
		t.Errorf("Normalize: result %v should equal %v", result, expected)
	}

	vector = Vector(1, 2, 3)
	result = vector.Normalize()
	mag := result.Magnitude()

	pass = floatEqual(mag, 1.0)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", mag, 1.0)
	}
}

func TestDotProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	result := a.DotProduct(b)
	expected := 20.0

	pass := floatEqual(result, expected)
	if !pass {
		t.Errorf("Magnitude: result %f should equal %f", result, expected)
	}
}

func TestCross(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	result1 := a.CrossProduct(b)
	expected1 := Vector(-1, 2, -1)
	result2 := b.CrossProduct(a)
	expected2 := Vector(1, -2, 1)

	pass := result1.Equals(expected1)

	if !pass {
		t.Errorf("Cross: result %v should equal %v", result1, expected1)
	}

	pass = result2.Equals(expected2)
	if !pass {
		t.Errorf("Cross: result %v should equal %v", result2, expected2)
	}
}
