package main

import (
	"math"
	"testing"
)

func TestSphereNormal(t *testing.T) {

	// The normal on a sphere at a point on the x axis.
	s := NewSphere()
	n := s.NormalAt(Point(1, 0, 0))
	expected := Vector(1, 0, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the y axis.
	n = s.NormalAt(Point(0, 1, 0))
	expected = Vector(0, 1, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the z axis.
	n = s.NormalAt(Point(0, 0, 1))
	expected = Vector(0, 0, 1)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a nonaxial point.
	v := math.Sqrt(3) / 3
	n = s.NormalAt(Point(v, v, v))
	expected = Vector(v, v, v)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// Computing the normal on a translated sphere.
	s.SetTransform(Translation(0, 1, 0))
	n = s.NormalAt(Point(0, 1.70711, -0.70711))
	expected = Vector(0, 0.70711, -0.70711)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	v = math.Sqrt(2) / 2
	s = NewSphere()
	transformMatrix := Scaling(1, 0.5, 1).MultiplyMatrix(RotationZ(PI / 5))
	s.SetTransform(transformMatrix)
	n = s.NormalAt(Point(0, v, -v))
	expected = Vector(0, 0.97014, -0.24254)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}
}
