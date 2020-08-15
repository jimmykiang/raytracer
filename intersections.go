package main

import "sort"

// Intersection struct.
type Intersection struct {
	t, u, v float64
	object  Shape
}

// Intersections contains a slice of Intersection pointers.
type Intersections []*Intersection

// NewIntersection returns a reference of the intersection struct.
func NewIntersection(t float64, object Shape) *Intersection {
	return &Intersection{
		t:      t,
		object: object,
	}
}

// NewIntersectionUV adds u and v Properties to the intersection struct.
func NewIntersectionUV(t float64, s Shape, u, v float64) *Intersection {
	return &Intersection{
		t:      t,
		object: s,
		u:      u,
		v:      v,
	}
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

// IntersectionAllowed evaluates the Rules for a CSG operations.
func IntersectionAllowed(op string, lhit, inl, inr bool) bool {
	if op == "union" {
		return (lhit && !inr) || (!lhit && !inl)
	} else if op == "intersection" {
		return (lhit && inr) || (!lhit && inl)
	} else if op == "difference" {
		return (lhit && !inr) || (!lhit && inl)
	}

	return false
}

// FilterIntersections will produce a subset of only those intersections that
// conform to the operation of the current CSG object.
func FilterIntersections(csg *CSG, xs []*Intersection) []*Intersection {
	// begin outside of both children
	inl := false
	inr := false
	// prepare a list to receive the filtered intersections
	result := make([]*Intersection, 0)

	for i, v := range xs {
		// if i.object is part of the "left" child, then lhit is true
		lhit := includes(csg.left, v.object)

		if IntersectionAllowed(csg.operation, lhit, inl, inr) {
			result = append(result, xs[i])
		}

		// depending on which object was hit, toggle either inl or inr
		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}
	return result
}

func includes(left Shape, object Shape) bool {
	switch t := left.(type) {
	case *Group:
		for _, child := range t.children {
			if child.GetID() == object.GetID() {
				return true
			}
			return includes(child, object)
		}
		return false
	case *CSG:
		a := includes(t.left, object)
		b := includes(t.right, object)
		return a || b
	default:
		return left.GetID() == object.GetID()
	}
}
