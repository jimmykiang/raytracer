package main

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := Point(-3, 4, 5)

	result := transform.MultiplyMatrixByTuple(p)

	expected := Point(2, 1, 7)

	if !result.Equals(expected) {
		t.Errorf("Translation:Point expected %v to equal %v", result, expected)
	}

	inv := transform.Inverse()

	result = inv.MultiplyMatrixByTuple(p)
	expected = Point(-8, 7, 3)

	if !result.Equals(expected) {
		t.Errorf("Translation:Point expected %v to equal %v", result, expected)
	}

	v := Vector(-3, 4, 5)

	if !transform.MultiplyMatrixByTuple(v).Equals(v) {
		t.Errorf("Translation:Vector vector was changed by translation.")
	}

}
