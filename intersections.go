package main

import "sort"

// Intersection struct.
type Intersection struct {
	t      float64
	object Shape
	index  int
}

// Intersections contains a slice of Intersection pointers.
type Intersections []*Intersection

// NewIntersection returns a reference of the intersection struct.
func NewIntersection(t float64, object Shape) *Intersection {
	return &Intersection{t, object, -1}
}

// NewIntersections returns an Intersections
func NewIntersections(intersections Intersections) Intersections {
	return Intersections(intersections)
}

// Hit returns the closest object with positive intersection.
func (xs Intersections) Hit() *Intersection {

	sort.Slice(xs, func(i, j int) bool { return xs[i].t < xs[j].t })

	for _, i := range xs {
		if i.t >= 0.0 {
			return i
		}
	}
	return nil
}

// Count returns the number of the Intersections.
func (xs *Intersections) Count() int {
	return len(*xs)
}
