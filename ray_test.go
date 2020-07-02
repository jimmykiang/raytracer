package main

import (
	"testing"
)

// Computing a point from a distance.
func TestRayPosition(t *testing.T) {
	ray := NewRay(Point(2, 3, 4), Vector(1, 0, 0))

	results := []*Tuple{
		ray.Position(0),
		ray.Position(1),
		ray.Position(-1),
		ray.Position(2.5),
	}
	expected := []*Tuple{
		Point(2, 3, 4),
		Point(3, 3, 4),
		Point(1, 3, 4),
		Point(4.5, 3, 4),
	}
	for i := 0; i < 4; i++ {
		if !results[i].Equals(expected[i]) {
			t.Errorf("RayPosition: expected %v to be %v", results[i], expected[i])
		}
	}
}

func TestIntersectSphere(t *testing.T) {

	// A ray intersects a sphere at two points.
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	s := NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}

	expected := []float64{4.0, 6.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

	// A ray intersects a sphere at a tangent.
	r = NewRay(Point(0, 1, -5), Vector(0, 0, 1))
	xs = s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}
	expected = []float64{5.0, 5.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

	// A ray misses a sphere.
	r = NewRay(Point(0, 2, -5), Vector(0, 0, 1))
	xs = s.Intersect(r)

	if len(xs) != 0 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 0, len(xs))
	}

	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs = s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}

	// A ray originates inside a sphere.
	expected = []float64{-1.0, 1.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

	// A sphere is behind a ray.
	r = NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	xs = s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}
	expected = []float64{-6.0, -4.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

}
