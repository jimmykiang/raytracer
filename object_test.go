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

// The normal of a plane is constant everywhere
func TestPlaneNormal(t *testing.T) {
	p := NewPlane()
	n1 := p.NormalAt(Point(0, 0, 0))
	n2 := p.NormalAt(Point(10, 0, -10))
	n3 := p.NormalAt(Point(-5, 0, 150))
	expected := Vector(0, 1, 0)

	if !n1.Equals(expected) {
		t.Errorf("PlaneNormal: expected %v to equal %v", n1, expected)
	}
	if !n2.Equals(expected) {
		t.Errorf("PlaneNormal: expected %v to equal %v", n2, expected)
	}
	if !n3.Equals(expected) {
		t.Errorf("PlaneNormal: expected %v to equal %v", n3, expected)
	}
}

func TestPlaneIntersect(t *testing.T) {
	p := NewPlane()
	r := NewRay(Point(0, 10, 0), Vector(0, 0, 1))

	xs := p.Intersect(r)

	if len(xs) != 0 {
		t.Errorf("PlaneIntersect(parallel): expected no intersections")
	}

	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs = p.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("PlaneIntersect(coplanar): expected no intersections")
	}

	r = NewRay(Point(0, 1, 0), Vector(0, -1, 0))
	xs = p.Intersect(r)

	if len(xs) != 1 {
		t.Errorf("PlaneIntersect(above): expected one intersection")
	}

	if !floatEqual(xs[0].t, 1) {
		t.Errorf("PlaneIntersect(above): expected intersection at %v to be %v", xs[0].t, 1)
	}

}
