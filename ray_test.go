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
