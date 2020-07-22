package main

import (
	"math"
	"testing"
)

func TestSphereNormal(t *testing.T) {

	// The normal on a sphere at a point on the x axis.
	s := NewSphere()
	n := s.localNormalAt(Point(1, 0, 0))
	expected := Vector(1, 0, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the y axis.
	n = s.localNormalAt(Point(0, 1, 0))
	expected = Vector(0, 1, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the z axis.
	n = s.localNormalAt(Point(0, 0, 1))
	expected = Vector(0, 0, 1)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a nonaxial point.
	v := math.Sqrt(3) / 3
	n = s.localNormalAt(Point(v, v, v))
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
	n1 := p.localNormalAt(Point(0, 0, 0))
	n2 := p.localNormalAt(Point(10, 0, -10))
	n3 := p.localNormalAt(Point(-5, 0, 150))
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

func TestCubeNormal(t *testing.T) {
	// The normal on the surface of a cube
	type cubeTest struct {
		point, normal *Tuple
	}
	c := NewCube()
	expectedNormals := []*cubeTest{
		{point: Point(1, 0.5, -0.8), normal: Vector(1, 0, 0)},
		{point: Point(-1, -0.2, 0.9), normal: Vector(-1, 0, 0)},
		{point: Point(-0.4, 1, -0.1), normal: Vector(0, 1, 0)},
		{point: Point(0.3, -1, -0.7), normal: Vector(0, -1, 0)},
		{point: Point(-0.6, 0.3, 1), normal: Vector(0, 0, 1)},
		{point: Point(0.4, 0.4, -1), normal: Vector(0, 0, -1)},
		{point: Point(1, 1, 1), normal: Vector(1, 0, 0)},
		{point: Point(-1, -1, -1), normal: Vector(-1, 0, 0)},
	}

	for _, v := range expectedNormals {
		n := c.localNormalAt(v.point)

		if !n.Equals(v.normal) {
			t.Errorf("The normal on the surface of a cube, got: %v and expected to be %v", n, v.normal)
		}
	}
}

func TestCylinderNormal(t *testing.T) {
	// Normal vector on a cylinder.

	type cylindertest struct {
		point, normal *Tuple
	}

	c := NewCylinder()

	expectedNormals := []*cylindertest{
		{point: Point(1, 0, 0), normal: Vector(1, 0, 0)},
		{point: Point(0, 5, -1), normal: Vector(0, 0, -1)},
		{point: Point(0, -2, 1), normal: Vector(0, 0, 1)},
		{point: Point(-1, 1, 0), normal: Vector(-1, 0, 0)},
	}

	for _, v := range expectedNormals {
		n := c.localNormalAt(v.point)

		if !n.Equals(v.normal) {
			t.Errorf("Normal vector on a cylinder, got: %v and expected to be %v", n, v.normal)
		}
	}
}
