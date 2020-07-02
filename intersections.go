package main

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
