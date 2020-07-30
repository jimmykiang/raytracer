package main

import (
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
	if !(s2.getId() == xs[0].object.getId()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.getId(), s2.getId())
	}
	if !(s2.getId() == xs[1].object.getId()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.getId(), s2.getId())
	}
	if !(s1.getId() == xs[2].object.getId()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.getId(), s2.getId())
	}
	if !(s1.getId() == xs[3].object.getId()) {
		t.Errorf("Intersecting a ray with a nonempty group: got: %v, expected: %v", xs[0].object.getId(), s2.getId())
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
