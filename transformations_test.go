package main

import (
	"math"
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

func TestScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)

	result := transform.MultiplyMatrixByTuple(p)
	expected := Point(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Point expected %v to equal %v", result, expected)
	}

	v := Vector(-4, 6, 8)

	result = transform.MultiplyMatrixByTuple(v)
	expected = Vector(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Vector expected %v to equal %v", result, expected)
	}

	inv := transform.Inverse()
	result = inv.MultiplyMatrixByTuple(v)
	expected = Vector(-2, 2, 2)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Vector (inverse) expected %v to equal %v", result, expected)
	}

	transform = Scaling(-1, 1, 1)
	p = Point(2, 3, 4)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(-2, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Point (reflection) expected %v to equal %v", result, expected)
	}
}

func TestRotationX(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationX(PI / 4)

	result := halfQuarter.MultiplyMatrixByTuple(p)

	expected := Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !result.Equals(expected) {

		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}

	fullQuarter := RotationX(PI / 2)

	expected = Point(0, 0, 1)

	result = fullQuarter.MultiplyMatrixByTuple(p)

	if !result.Equals(expected) {
		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}

	invHalfQuarter := halfQuarter.Inverse()

	expected = Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	result = invHalfQuarter.MultiplyMatrixByTuple(p)

	if !result.Equals(expected) {
		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}
}

func TestRotationY(t *testing.T) {
	p := Point(0, 0, 1)
	halfQuarter := RotationY(PI / 4)

	result := halfQuarter.MultiplyMatrixByTuple(p)

	expected := Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

	fullQuarter := RotationY(PI / 2)

	result = fullQuarter.MultiplyMatrixByTuple(p)

	expected = Point(1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}
}
