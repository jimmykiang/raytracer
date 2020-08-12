package main

import (
	"math"
	"testing"
)

func TestSphereNormal(t *testing.T) {

	// The normal on a sphere at a point on the x axis.
	s := NewSphere()
	n := s.localNormalAt(Point(1, 0, 0), nil)
	expected := Vector(1, 0, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the y axis.
	n = s.localNormalAt(Point(0, 1, 0), nil)
	expected = Vector(0, 1, 0)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a point on the z axis.
	n = s.localNormalAt(Point(0, 0, 1), nil)
	expected = Vector(0, 0, 1)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// The normal on a sphere at a nonaxial point.
	v := math.Sqrt(3) / 3
	n = s.localNormalAt(Point(v, v, v), nil)
	expected = Vector(v, v, v)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	// Computing the normal on a translated sphere.
	s.SetTransform(Translation(0, 1, 0))
	n = s.NormalAt(Point(0, 1.70711, -0.70711), nil)
	expected = Vector(0, 0.70711, -0.70711)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}

	v = math.Sqrt(2) / 2
	s = NewSphere()
	transformMatrix := Scaling(1, 0.5, 1).MultiplyMatrix(RotationZ(PI / 5))
	s.SetTransform(transformMatrix)
	n = s.NormalAt(Point(0, v, -v), nil)
	expected = Vector(0, 0.97014, -0.24254)
	if !n.Equals(expected) {
		t.Errorf("SphereNormal: expected %v to be %v", n, expected)
	}
}

// The normal of a plane is constant everywhere
func TestPlaneNormal(t *testing.T) {
	p := NewPlane()
	n1 := p.localNormalAt(Point(0, 0, 0), nil)
	n2 := p.localNormalAt(Point(10, 0, -10), nil)
	n3 := p.localNormalAt(Point(-5, 0, 150), nil)
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
		n := c.localNormalAt(v.point, nil)

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
		n := c.localNormalAt(v.point, nil)

		if !n.Equals(v.normal) {
			t.Errorf("Normal vector on a cylinder, got: %v and expected to be %v", n, v.normal)
		}
	}
}

func TestCylinderMinMax(t *testing.T) {
	// The default minimum and maximum for a cylinder.

	c := NewCylinder()

	if c.minimum != math.Inf(-1) {
		t.Errorf("The default minimum for a cylinder, got: %v and expected to be %v", c.minimum, math.Inf(-1))
	}

	if c.maximum != math.Inf(1) {
		t.Errorf("The default maximum for a cylinder, got: %v and expected to be %v", c.maximum, math.Inf(1))
	}
}

func TestCylinderClosedValue(t *testing.T) {
	// The default closed value for a cylinder.

	c := NewCylinder()

	if c.closed != false {
		t.Errorf("The default closed value for a cylinder, got: %v and expected to be %v", c.closed, false)
	}
}

func TestCylinderEndCapsNormal(t *testing.T) {
	// The normal vector on a cylinder's end caps.

	type cylindertest struct {
		point, normal *Tuple
	}

	c := NewCylinder()
	c.minimum = 1
	c.maximum = 2
	c.closed = true

	expectedNormals := []*cylindertest{
		{point: Point(0, 1, 0), normal: Vector(0, -1, 0)},
		{point: Point(0.5, 1, 0), normal: Vector(0, -1, 0)},
		{point: Point(0, 1, 0.5), normal: Vector(0, -1, 0)},
		{point: Point(0, 2, 0), normal: Vector(0, 1, 0)},
		{point: Point(0.5, 2, 0), normal: Vector(0, 1, 0)},
		{point: Point(0, 2, 0.5), normal: Vector(0, 1, 0)},
	}

	for _, v := range expectedNormals {
		n := c.localNormalAt(v.point, nil)

		if !n.Equals(v.normal) {
			t.Errorf("The normal vector on a cylinder's end caps, got: %v and expected to be %v", n, v.normal)
		}
	}
}

func TestConeNormal(t *testing.T) {
	// Computing the normal vector on a cone.

	type cylindertest struct {
		point, normal *Tuple
	}

	c := NewCone()

	expectedNormals := []*cylindertest{
		{point: Point(0, 0, 0), normal: Vector(0, 0, 0)},
		{point: Point(1, 1, 1), normal: Vector(1, -math.Sqrt(2), 1)},
		{point: Point(-1, -1, 0), normal: Vector(-1, 1, 0)},
	}

	for _, v := range expectedNormals {
		n := c.localNormalAt(v.point, nil)

		if !n.Equals(v.normal) {
			t.Errorf("Computing the normal vector on a cone, got: %v and expected to be %v", n, v.normal)
		}
	}
}

func TestConstructTriangle(t *testing.T) {
	// Constructing a triangle.

	p1 := Point(0, 1, 0)
	p2 := Point(-1, 0, 0)
	p3 := Point(1, 0, 0)
	triangle := NewTriangle(p1, p2, p3)

	expectedE1 := Vector(-1, -1, 0)
	expectedE2 := Vector(1, -1, 0)
	expectedNormal := Vector(0, 0, -1)

	if !(triangle.p1.Equals(p1)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.p1, p1)
	}
	if !(triangle.p2.Equals(p2)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.p2, p2)
	}
	if !(triangle.p3.Equals(p3)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.p3, p3)
	}
	if !(triangle.e1.Equals(expectedE1)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.e1, expectedE1)
	}
	if !(triangle.e2.Equals(expectedE2)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.e2, expectedE2)
	}
	if !(triangle.normal.Equals(expectedNormal)) {
		t.Errorf("Constructing a triangle, got: %v and expected to be %v", triangle.normal, expectedNormal)
	}
}

func TestTriangleNormal(t *testing.T) {
	// Finding the normal on a triangle.

	triangle := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	n1 := triangle.localNormalAt(Point(0, 0.5, 0), nil)
	n2 := triangle.localNormalAt(Point(-0.5, 0.75, 0), nil)
	n3 := triangle.localNormalAt(Point(0.5, 0.25, 0), nil)

	expectedNormal := triangle.normal

	if !(n1.Equals(expectedNormal)) {
		t.Errorf("Finding the normal on a triangle, got: %v and expected to be %v", n1, expectedNormal)
	}
	if !(n2.Equals(expectedNormal)) {
		t.Errorf("Finding the normal on a triangle, got: %v and expected to be %v", n2, expectedNormal)
	}
	if !(n3.Equals(expectedNormal)) {
		t.Errorf("Finding the normal on a triangle, got: %v and expected to be %v", n3, expectedNormal)
	}
}

func TestSmoothTriangleSetup(t *testing.T) {
	// Constructing a smooth triangle.
	smoothTriangle := defaultSmoothTriangle()

	expectedPoint1 := Point(0, 1, 0)
	expectedPoint2 := Point(-1, 0, 0)
	expectedPoint3 := Point(1, 0, 0)
	expectedNormal1 := Vector(0, 1, 0)
	expectedNormal2 := Vector(-1, 0, 0)
	expectedNormal3 := Vector(1, 0, 0)

	if !smoothTriangle.p1.Equals(expectedPoint1) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedPoint1)
	}
	if !smoothTriangle.p2.Equals(expectedPoint2) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedPoint2)
	}
	if !smoothTriangle.p3.Equals(expectedPoint3) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedPoint3)
	}
	if !smoothTriangle.n1.Equals(expectedNormal1) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedNormal1)
	}
	if !smoothTriangle.n2.Equals(expectedNormal2) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedNormal2)
	}
	if !smoothTriangle.n3.Equals(expectedNormal3) {
		t.Errorf("Constructing a smooth triangle: expected %v to be %v", smoothTriangle.p1, expectedNormal3)
	}
}

func TestSmoothTriangleWithUV(t *testing.T) {
	// An intersection can encapsulate `u` and `v`.
	smoothTriangle := defaultSmoothTriangle()
	xs := NewIntersectionUV(3.5, smoothTriangle, 0.2, 0.4)

	expectedU := 0.2
	expectedV := 0.4

	if !floatEqual(xs.u, expectedU) {
		t.Errorf("An intersection can encapsulate `u` and `v`: expected %v to be %v", xs.u, expectedU)
	}
	if !floatEqual(xs.v, expectedV) {
		t.Errorf("An intersection can encapsulate `u` and `v`: expected %v to be %v", xs.v, expectedV)
	}
}

func TestIntersectWithTriStoresUV(t *testing.T) {
	// An intersection with a smooth triangle stores u/v.
	smoothTriangle := defaultSmoothTriangle()
	r := NewRay(Point(-0.2, 0.3, -2), Vector(0, 0, 1))
	xs := smoothTriangle.localIntersect(r)

	expectedU := 0.45
	expectedV := 0.25

	if !floatEqual(xs[0].u, expectedU) {
		t.Errorf("An intersection with a smooth triangle stores u/v: expected %v to be %v", xs[0].u, expectedU)
	}
	if !floatEqual(xs[0].v, expectedV) {
		t.Errorf("An intersection with a smooth triangle stores u/v: expected %v to be %v", xs[0].v, expectedV)
	}
}

func TestInterpolatedNormal(t *testing.T) {
	// A smooth triangle uses u/v to interpolate the normal.
	smoothTriangle := defaultSmoothTriangle()
	i := NewIntersectionUV(1, smoothTriangle, 0.45, 0.25)
	n := NormalAt(smoothTriangle, Point(0, 0, 0), i)

	expectedVector := Vector(-0.5547, 0.83205, 0)

	if !n.Equals(expectedVector) {
		t.Errorf("A smooth triangle uses u/v to interpolate the normal: expected %v to be %v", n, expectedVector)
	}
}

func TestPrepareNormalOnSmoothTri(t *testing.T) {
	// Preparing the normal on a smooth triangle.
	smoothTriangle := defaultSmoothTriangle()
	i := NewIntersectionUV(1.0, smoothTriangle, 0.45, 0.25)
	r := NewRay(Point(-0.2, 0.3, -2), Vector(0, 0, 1))
	xs := []*Intersection{i}
	comps := PrepareComputations(i, r, xs)

	expectedVector := Vector(-0.5547, 0.83205, 0)

	if !comps.normalv.Equals(expectedVector) {
		t.Errorf("Preparing the normal on a smooth triangle: expected %v to be %v",
			comps.normalv, expectedVector)
	}
}
