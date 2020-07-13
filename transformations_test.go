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

func TestRotationZ(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationZ(PI / 4)

	result := halfQuarter.MultiplyMatrixByTuple(p)

	expected := Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

	fullQuarter := RotationZ(PI / 2)

	result = fullQuarter.MultiplyMatrixByTuple(p)

	expected = Point(-1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

}

func TestShearing(t *testing.T) {
	p := Point(2, 3, 4)

	transform := Shearing(1, 0, 0, 0, 0, 0)
	result := transform.MultiplyMatrixByTuple(p)
	expected := Point(5, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: xy expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 1, 0, 0, 0, 0)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(6, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: xz expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 1, 0, 0, 0)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(2, 5, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: yx expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 1, 0, 0)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(2, 7, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: yz expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 0, 1, 0)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(2, 3, 6)

	if !result.Equals(expected) {
		t.Errorf("Shearing: zx expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 0, 0, 1)
	result = transform.MultiplyMatrixByTuple(p)
	expected = Point(2, 3, 7)

	if !result.Equals(expected) {
		t.Errorf("Shearing: zy expected %v to be %v", result, expected)
	}
}

func TestChainTransformations(t *testing.T) {
	p := Point(1, 0, 1)
	A := RotationX(PI / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)

	// Apply rotation first.
	// p2 ← A * p
	p2 := p.Transform(A)
	expected2 := Point(1, -1, 0)

	if !p2.Equals(expected2) {
		t.Errorf("p2 ← A * p: expected %v to be %v", p2, expected2)
	}

	// then apply scaling.
	// p3 ← B * p2
	p3 := p2.Transform(B)
	expected3 := Point(5, -5, 0)

	if !p3.Equals(expected3) {
		t.Errorf("p3 ← B * p2: expected %v to be %v", p3, expected3)
	}

	// then apply translation.
	// p4 ← C * p3

	p4 := p3.Transform(C)
	expected4 := Point(15, 0, 7)

	if !p4.Equals(expected4) {
		t.Errorf("p4 ← C * p3: expected %v to be %v", p4, expected4)
	}

	// T ← C * B * A
	// T * p
	result := p.Transform(A, B, C)
	expected := Point(15, 0, 7)

	if !result.Equals(expected) {
		t.Errorf("ChainTransformations: expected %v to be %v", result, expected)
	}
}

func TestSphereTransformation(t *testing.T) {

	// A sphere's default transformation.
	s := NewSphere()

	if !s.transform.Equals(IdentityMatrix) {
		t.Errorf("SphereTransformation: expected %v to be %v", IdentityMatrix, s.transform)
	}

	// Changing a sphere's transformation.

	s = NewSphere()
	transform := Translation(2, 3, 4)

	s.transform = transform

	if !s.transform.Equals(transform) {
		t.Errorf("SphereTransformation: expected %v to be %v", transform, s.transform)
	}
}

func TestViewTransform(t *testing.T) {
	// The transformation matrix for the default orientation.
	from := Point(0, 0, 0)
	to := Point(0, 0, -1)
	up := Vector(0, 1, 0)

	result := ViewTransform(from, to, up)
	expected := IdentityMatrix

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(default): expected %v to equal %v", result, expected)
	}

	// A view transformation matrix looking in positive z direction.
	from = Point(0, 0, 0)
	to = Point(0, 0, 1)
	up = Vector(0, 1, 0)
	result = ViewTransform(from, to, up)
	expected = Scaling(-1, 1, -1)

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(positive z): expected %v to equal %v", result, expected)
	}

	// The view transformation moves the world.
	from = Point(0, 0, 8)
	to = Point(0, 0, 0)
	up = Vector(0, 1, 0)
	result = ViewTransform(from, to, up)
	expected = Translation(0, 0, -8)

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(moves world): expected %v to equal %v", result, expected)
	}

	// An arbitrary view transformation.
	from = Point(1, 3, 2)
	to = Point(4, -2, 8)
	up = Vector(1, 1, 0)
	result = ViewTransform(from, to, up)
	expected = [][]float64{
		[]float64{-0.507093, 0.507093, 0.676123, -2.366432},
		[]float64{0.767716, 0.606092, 0.121218, -2.828427},
		[]float64{-0.358569, 0.597614, -0.717137, 0},
		[]float64{0, 0, 0, 1},
	}
	if !result.Equals(expected) {
		t.Errorf("ViewTransform(arbitary): expected %v to equal %v", result, expected)
	}
}
