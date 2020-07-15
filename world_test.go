package main

import "testing"

// Intersect a world with a ray.
func TestWorldIntersections(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	xs := w.Intersect(r)
	xs.Hit()
	count := xs.Count()
	if count != 4 {
		t.Errorf("WorldIntersections: %v should be %v", count, 4)
	}

	expected := []float64{4, 4.5, 5.5, 6}

	for i, v := range expected {
		top := xs[i]
		if top.t != v {
			t.Errorf("WorldIntersections: expected hit %v to be %v", top.t, v)
		}
	}
}

func TestPrepareComputation(t *testing.T) {
	//  Precomputing the state of an intersection.
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := &Intersection{4, shape, -1}
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	if !floatEqual(comps.t, i.t) {
		t.Errorf("PrepareComputations failed")
	}

	if comps.object != i.object {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.point.Equals(Point(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.eyev.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.normalv.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if comps.inside {
		t.Errorf("PrepareComputations failed")
	}

	// The hit, when an intersection occurs on the inside.
	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = NewSphere()
	i = &Intersection{1, shape, -1}
	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	if !comps.point.Equals(Point(0, 0, 1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.eyev.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.normalv.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.inside {
		t.Errorf("PrepareComputations failed")
	}
}

func TestShadeHit(t *testing.T) {
	// Shading an intersection.
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.objects[0]
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result := w.ShadeHit(comps, 10)
	expected := NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	// Shading an intersection from the inside.
	w.lights[0] = NewPointLight(Point(0, .25, 0), NewColor(1, 1, 1))
	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = w.objects[1]
	i = NewIntersection(0.5, shape)
	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result = w.ShadeHit(comps, 10)
	expected = NewColor(0.90498, 0.90498, 0.90498)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	// shade_hit() is given an intersection in shadow.
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, 10))

	w = NewWorld(
		[]*PointLight{NewPointLight(Point(0, 0, -10), NewColor(1, 1, 1))},
		[]Shape{s1, s2},
	)
	r = NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	i = NewIntersection(4, s2)

	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result = w.ShadeHit(comps, 10)
	expected = NewColor(0.1, 0.1, 0.1)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}
}

func TestWorldColorAt(t *testing.T) {

	// The color when a ray misses.
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 1, 0))
	result := w.ColorAt(r, 10)
	expected := Black

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (no hit): expected %v to be %v", result, expected)
	}

	// The color when a ray hits.
	r = NewRay(r.origin, Vector(0, 0, 1))
	result = w.ColorAt(r, 10)
	expected = NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit): expected %v to be %v", result, expected)
	}

	// The color with an intersection behind the ray.
	outer := w.objects[0]
	outer.Material().ambient = 1
	inner := w.objects[1]

	inner.Material().ambient = 1

	r = NewRay(Point(0, 0, .75), Vector(0, 0, -1))

	result = w.ColorAt(r, 10)
	expected = inner.Material().color

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit inner): expected %v to be %v", result, expected)
	}
}

func TestIsShadowed(t *testing.T) {
	w := DefaultWorld()

	p := Point(0, 10, 0)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: expected no shadow when nothing is collinear point and light")
	}

	p = Point(10, -10, 10)
	if !w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: expected object between point and light to create shadow")
	}

	p = Point(-20, 20, -20)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: There should be no shadow when an object is behind the light")
	}

	p = Point(-2, 2, -2)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: There is no shadow when an object is behind the point ")
	}
}
