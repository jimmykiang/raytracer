package main

import (
	"reflect"
	"testing"
)

func TestIntersections(t *testing.T) {

	//  An intersection encapsulates t and object.
	s := NewSphere()
	i := NewIntersection(3.5, s)

	if i.t != 3.5 {
		t.Errorf("TestIntersections: expected t from intersection to be %v but got %v", 3.5, i.t)
	}

	if i.object != s {
		t.Errorf("TestIntersections: expected object from intersection to be %T but got %T", s, i.object)
	}

	// Aggregating intersections.
	s = NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	xs := Intersections{i1, i2}

	if len(xs) != 2 {
		t.Errorf("TestIntersections: expected number of intersections to be %v but got %v", 2, len(xs))
	}

	expected := []float64{1, 2}

	for i, intersection := range xs {
		if expected[i] != intersection.t {
			t.Errorf("TestIntersections: expected %v to be %v", intersection.t, expected[i])
		}
	}

	// Intersect sets the object on the intersection.
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s = NewSphere()
	xs = s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("TestIntersections: expected number of intersections to be %v but got %v", 2, len(xs))
	}

	expectedShape := []*Sphere{s, s}

	for i, intersection := range xs {
		if reflect.TypeOf(expectedShape[i]) != reflect.TypeOf(intersection.object) {
			t.Errorf("TestIntersections: expected %T to be %T", intersection.object, expectedShape[i])
		}
	}
}

func TestHit(t *testing.T) {
	// The hit, when all intersections have positive t.
	s := NewSphere()
	i1 := &Intersection{1, s, -1}
	i2 := &Intersection{2, s, -1}
	xs := NewIntersections([]*Intersection{i1, i2})
	i := xs.Hit()
	if i != i1 {
		t.Errorf("Hit: expected %v to be %v", i, i1)
	}

	// The hit, when some intersections have negative t.
	i1 = &Intersection{-1, s, -1}
	i2 = &Intersection{2, s, -1}
	xs = NewIntersections([]*Intersection{i1, i2})
	i = xs.Hit()
	if i != i2 {
		t.Errorf("Hit: expected %v to be %v", i, i2)
	}

	// The hit, when all intersections have negative t.
	i1 = &Intersection{-1, s, -1}
	i2 = &Intersection{-2, s, -1}
	xs = NewIntersections([]*Intersection{i1, i2})
	i = xs.Hit()
	if i != nil {
		t.Errorf("Hit: expected %v to be %v", i, nil)
	}

	// The hit is always the lowest nonnegative intersection.
	i1 = &Intersection{5, s, -1}
	i2 = &Intersection{7, s, -1}
	i3 := &Intersection{-3, s, -1}
	i4 := &Intersection{2, s, -1}
	xs = NewIntersections([]*Intersection{i1, i2, i3, i4})
	i = xs.Hit()
	if i != i4 {
		t.Errorf("Hit: expected %v to be %v", i, i4)
	}

}
