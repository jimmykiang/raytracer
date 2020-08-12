package main

import (
	"math"
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	// Creating a new group.
	g := NewGroup()
	g.transform = NewIdentityMatrix()

	if !(len(g.children) == 0) {
		t.Errorf("Creating a new group, does not contain children shapes: got: %v, expected: %v", len(g.children), 0)
	}
	if !(g.transform != nil) {
		t.Errorf("Creating a new group, contains a default transformation matrix: got: %v, expected: %v", reflect.TypeOf(g.Transform), reflect.TypeOf(NewIdentityMatrix()))
	}
}

func TestGroup_AddChild(t *testing.T) {
	// Adding a child to a group.
	g := NewGroup()
	s := NewSphere()

	g.AddChild(s)

	if !(len(g.children) == 1) {
		t.Errorf("Adding a child to a group, should contain 1 child shapes: got: %v, expected: %v", len(g.children), 0)
	}

	if !(g.children[0] == s) {
		t.Errorf("Adding a child to a group, should contain the child shape: got: %v, expected: %v", reflect.TypeOf(g.children[0]), reflect.TypeOf(s))
	}

	if !(s.GetParent() == g) {
		t.Errorf("Adding a child to a group, child shape should contain (reference) parent shape (group): got: %v, expected: %v", reflect.TypeOf(s.GetParent), reflect.TypeOf(g))
	}
}

func TestIntersectEmptyGroup(t *testing.T) {
	// Intersecting a ray with an empty group.
	g := NewGroup()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs := g.localIntersect(r)

	if !(len(xs) == 0) {
		t.Errorf("Intersecting a ray with an empty group: got: %v, expected: %v", len(xs), 0)
	}
}

func TestIntersectGroup(t *testing.T) {
	// Intersecting a ray with a nonempty group.
	// The spheres are arranged inside the group so that the ray will intersect two of the
	// spheres but miss the third. The resulting collection of intersections should
	// include those of the two spheres.

	g := NewGroup()
	s1 := NewSphere()

	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, -3))

	s3 := NewSphere()
	s3.SetTransform(Translation(5, 0, 0))

	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	xs := g.localIntersect(r)

	if !(len(xs) == 4) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", len(xs), 4)
	}
	if !(s2.GetID() == xs[0].object.GetID()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.GetID(), s2.GetID())
	}
	if !(s2.GetID() == xs[1].object.GetID()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.GetID(), s2.GetID())
	}
	if !(s1.GetID() == xs[2].object.GetID()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.GetID(), s2.GetID())
	}
	if !(s1.GetID() == xs[3].object.GetID()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.GetID(), s2.GetID())
	}

	xs3 := s3.Intersect(r)
	if !(len(xs3) == 0) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", len(xs3), 0)
	}
}

func TestGroupTransform(t *testing.T) {
	// Intersecting a transformed group.

	g := NewGroup()
	g.SetTransform(Scaling(2, 2, 2))
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g.AddChild(s)

	r := NewRay(Point(10, 0, -10), Vector(0, 0, 1))
	xs := g.Intersect(r)

	if !(len(xs) == 2) {
		t.Errorf("Intersecting a transformed group: got: %v, expected: %v", len(xs), 2)
	}
}

func TestConvertPointFromWorldToObjectSpace(t *testing.T) {
	// Converting a point from world to object space.

	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(2, 2, 2))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)

	p := WorldToObject(s, Point(-2, 0, -10))

	expectedPoint := Point(0, 0, -1)
	if !(expectedPoint.Equals(p)) {
		t.Errorf("Converting a point from world to object space: got: %v, expected: %v", p, expectedPoint)
	}
}

func TestConvertNormalFromObjectToWorldSpace(t *testing.T) {
	// Converting a normal from object to world space.

	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)

	v := NormalToWorld(s, Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	expectedVector := Vector(0.28571428571428575, 0.42857142857142855, -0.8571428571428571)
	if !(expectedVector.Equals(v)) {
		t.Errorf("Converting a normal from object to world space: got: %v, expected: %v", v, expectedVector)
	}
}

func TestNormalOnChildObject(t *testing.T) {
	// Finding the normal on a child object.

	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.AddChild(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.AddChild(s)

	v := NormalAt(s, Point(1.7321, 1.1547, -5.5774), nil)

	expectedVector := Vector(0.28570368184140726, 0.42854315178114105, -0.8571605294481017)
	if !(expectedVector.Equals(v)) {
		t.Errorf("Finding the normal on a child object: %v, expected: %v", v, expectedVector)
	}
}
