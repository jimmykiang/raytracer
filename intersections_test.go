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

	// localIntersect sets the object on the intersection.
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s = NewSphere()
	xs = s.localIntersect(r)

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
	i1 := &Intersection{t: 1, object: s}
	i2 := &Intersection{t: 2, object: s}
	xs := NewIntersections([]*Intersection{i1, i2})
	i := xs.Hit()
	if i != i1 {
		t.Errorf("Hit: expected %v to be %v", i, i1)
	}

	// The hit, when some intersections have negative t.
	i1 = &Intersection{t: -1, object: s}
	i2 = &Intersection{t: 2, object: s}
	xs = NewIntersections([]*Intersection{i1, i2})
	i = xs.Hit()
	if i != i2 {
		t.Errorf("Hit: expected %v to be %v", i, i2)
	}

	// The hit, when all intersections have negative t.
	i1 = &Intersection{t: -1, object: s}
	i2 = &Intersection{t: -2, object: s}
	xs = NewIntersections([]*Intersection{i1, i2})
	i = xs.Hit()
	if i != nil {
		t.Errorf("Hit: expected %v to be %v", i, nil)
	}

	// The hit is always the lowest nonnegative intersection.
	i1 = &Intersection{t: 5, object: s}
	i2 = &Intersection{t: 7, object: s}
	i3 := &Intersection{t: -3, object: s}
	i4 := &Intersection{t: 2, object: s}
	xs = NewIntersections([]*Intersection{i1, i2, i3, i4})
	i = xs.Hit()
	if i != i4 {
		t.Errorf("Hit: expected %v to be %v", i, i4)
	}
}

// The hit should offset the point.
func TestOverPoint(t *testing.T) {

	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	shape.SetTransform(Translation(0, 0, 1))

	i := NewIntersection(5, shape)
	xs := NewIntersections([]*Intersection{i})
	comps := PrepareComputations(i, r, xs)

	if !(comps.underPoint.z > EPSILON/2.0 && comps.point.z < comps.underPoint.z) {
		t.Errorf("PrepareComputationWithRefraction: underPoint %v not valid", comps.underPoint)
	}
}

func TestPlaneIntersect(t *testing.T) {
	p := NewPlane()
	r := NewRay(Point(0, 10, 0), Vector(0, 0, 1))

	xs := p.localIntersect(r)

	if len(xs) != 0 {
		t.Errorf("PlaneIntersect(parallel): expected no intersections")
	}

	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	xs = p.localIntersect(r)
	if len(xs) != 0 {
		t.Errorf("PlaneIntersect(coplanar): expected no intersections")
	}

	r = NewRay(Point(0, 1, 0), Vector(0, -1, 0))
	xs = p.localIntersect(r)

	if len(xs) != 1 {
		t.Errorf("PlaneIntersect(above): expected one intersection")
	}

	if !floatEqual(xs[0].t, 1) {
		t.Errorf("PlaneIntersect(above): expected intersection at %v to be %v", xs[0].t, 1)
	}
}

func TestCubeIntersect(t *testing.T) {
	// A ray intersects a cube.

	c := NewCube()

	expectedIntersectionMap := map[string][]interface{}{
		"x":      []interface{}{NewRay(Point(5, 0.5, 0), Vector(-1, 0, 0)), 4, 6},
		"-x":     []interface{}{NewRay(Point(-5, 0.5, 0), Vector(1, 0, 0)), 4, 6},
		"y":      []interface{}{NewRay(Point(0.5, 5, 0), Vector(0, -1, 0)), 4, 6},
		"-y":     []interface{}{NewRay(Point(0.5, -5, 0), Vector(0, 1, 0)), 4, 6},
		"z":      []interface{}{NewRay(Point(0.5, 0, 5), Vector(0, 0, -1)), 4, 6},
		"-z":     []interface{}{NewRay(Point(0.5, 0, -5), Vector(0, 0, 1)), 4, 6},
		"inside": []interface{}{NewRay(Point(0, 0.5, 0), Vector(0, 0, 1)), -1, 1},
	}

	for k, v := range expectedIntersectionMap {
		xs := c.localIntersect(v[0].(*Ray))

		if len(xs) != 2 {
			t.Errorf("A ray intersects a cube count: %v expected to be %v", len(xs), 2)
		}

		if !floatEqual(xs[0].t, float64(v[1].(int))) || !floatEqual(xs[1].t, float64(v[2].(int))) {
			t.Errorf("A ray intersects a cube: expected %v intersection xs[0].t = %v to be %v and xs[1].t = %v to be %v", k, xs[0].t, v[1], xs[1].t, v[2])
		}
	}
}

func TestCubeRayMisses(t *testing.T) {
	//  A ray misses a cube.
	c := NewCube()

	expectedIntersections := []*Ray{
		NewRay(Point(-2, 0, 0), Point(0.2673, 0.5345, 0.8018)),
		NewRay(Point(0, -2, 0), Point(0.8018, 0.2673, 0.5345)),
		NewRay(Point(0, 0, -2), Point(0.5345, 0.8018, 0.2673)),
		NewRay(Point(2, 0, 2), Point(0, 0, -1)),
		NewRay(Point(0, 2, 2), Point(0, -1, 0)),
		NewRay(Point(2, 2, 0), Point(-1, 0, 0)),
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(v)

		if len(xs) != 0 {
			t.Errorf("A ray misses a cube: expected Ray intersection count xs= %v to be %v", len(xs), 0)
		}
	}
}

func TestCylinderRayMisses(t *testing.T) {
	//  A ray misses a cylinder.

	c := NewCylinder()

	expectedIntersections := []*Ray{
		NewRay(Point(1, 0, 0), Point(0, 1, 0)),
		NewRay(Point(0, 0, 0), Point(0, 1, 0)),
		NewRay(Point(0, 0, -5), Point(1, 1, 1)),
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(v)

		if len(xs) != 0 {
			t.Errorf("A ray misses a cylinder: expected Ray intersection count to be xs= %v, got %v", 0, len(xs))
		}
	}
}

func TestCylinderRayStrike(t *testing.T) {
	// A ray strikes a cylinder.

	type cylindertest struct {
		point, direction *Tuple
		t0               float64
		t1               float64
	}

	c := NewCylinder()

	expectedIntersections := []*cylindertest{
		{point: Point(1, 0, -5), direction: Vector(0, 0, 1), t0: 5, t1: 5},
		{point: Point(1, 0, -5), direction: Vector(0, 0, 1), t0: 4, t1: 6},
		{point: Point(0.5, 0, -5), direction: Vector(0.1, 1, 1), t0: 6.80798, t1: 7.08872},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction))

		if len(xs) != 2 {
			t.Errorf("A ray strikes a cylinder: expected Ray intersection count to be xs= %v, got %v", 2, len(xs))
		}
	}
}

func TestCylinderConstraints(t *testing.T) {
	// Intersecting a constrained cylinder.

	type cylindertest struct {
		point, direction  *Tuple
		intersectionCount int
	}

	c := NewCylinder()
	c.minimum = 1
	c.maximum = 2

	expectedIntersections := []*cylindertest{
		{point: Point(0, 1.5, 0), direction: Vector(0.1, 1, 0), intersectionCount: 0},
		{point: Point(0, 3, -5), direction: Vector(0, 0, 1), intersectionCount: 0},
		{point: Point(0, 0, -5), direction: Vector(0, 0, 1), intersectionCount: 0},
		{point: Point(0, 2, -5), direction: Vector(0, 0, 1), intersectionCount: 0},
		{point: Point(0, 1, -5), direction: Vector(0, 0, 1), intersectionCount: 0},
		{point: Point(0, 1.5, -2), direction: Vector(0, 0, 1), intersectionCount: 2},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction))

		if len(xs) != v.intersectionCount {
			t.Errorf("A ray strikes a cylinder: expected Ray intersection count to be xs= %v, got %v", v.intersectionCount, len(xs))
		}
	}
}

func TestCylinderCapsIntersection(t *testing.T) {
	// Intersecting the caps of a closed cylinder

	type cylindertest struct {
		point, direction  *Tuple
		intersectionCount int
	}

	c := NewCylinder()
	c.minimum = 1
	c.maximum = 2
	c.closed = true

	expectedIntersections := []*cylindertest{
		{point: Point(0, 3, 0), direction: Vector(0, -1, 0), intersectionCount: 2},
		{point: Point(0, 3, -2), direction: Vector(0, -1, 2), intersectionCount: 2},
		{point: Point(0, 4, -2), direction: Vector(0, -1, 1), intersectionCount: 2},
		{point: Point(0, 0, -2), direction: Vector(0, 1, 2), intersectionCount: 2},
		{point: Point(0, -1, -2), direction: Vector(0, 1, 1), intersectionCount: 2},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction))

		if len(xs) != v.intersectionCount {
			t.Errorf("A ray strikes a cylinder: expected Ray intersection count to be xs= %v, got %v", v.intersectionCount, len(xs))
		}
	}
}

func TestConeRayIntersect(t *testing.T) {
	// Intersecting a cone with a ray.

	type conetest struct {
		point, direction *Tuple
		t0               float64
		t1               float64
	}

	c := NewCone()

	expectedIntersections := []*conetest{
		{point: Point(0, 0, -5), direction: Vector(0, 0, 1), t0: 5, t1: 5},
		{point: Point(0, 0, -5), direction: Vector(1, 1, 1), t0: 8.66025, t1: 8.66025},
		{point: Point(1, 1, -5), direction: Vector(-0.5, -1, 1), t0: 4.55006, t1: 49.44994},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction.Normalize()))

		if len(xs) != 2 {
			t.Errorf("Intersecting a cone with a ray: expected Ray intersection count to be xs= %v, got %v", 2, len(xs))
		}

		if !floatEqual(xs[0].t, v.t0) || !floatEqual(xs[1].t, v.t1) {
			t.Errorf("A ray intersects a cube: expected intersection xs[0].t = %v to be %v and xs[1].t = %v to be %v", xs[0].t, v.t0, xs[1].t, v.t1)
		}
	}
}

func TestConeRayIntersectParallelToHalves(t *testing.T) {
	// Intersecting a cone with a ray parallel to one of its halves.

	type conetest struct {
		point, direction *Tuple
		t0               float64
		t1               float64
	}

	c := NewCone()

	expectedIntersections := []*conetest{
		{point: Point(0, 0, -1), direction: Vector(0, 1, 1), t0: 0.35355, t1: 0.35355},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction.Normalize()))

		if len(xs) != 1 {
			t.Errorf("Intersecting a cone with a ray parallel to one of its halves: expected Ray intersection count to be xs= %v, got %v", 2, len(xs))
		}

		if !floatEqual(xs[0].t, v.t0) {
			t.Errorf("Intersecting a cone with a ray parallel to one of its halves: xs[0].t = %v to be %v", xs[0].t, v.t0)
		}
	}
}

func TestConeCapsIntersection(t *testing.T) {
	// Intersecting a cone's end caps.

	type conetest struct {
		point, direction  *Tuple
		intersectionCount int
	}

	c := NewCone()
	c.minimum = -0.5
	c.maximum = 0.5
	c.closed = true

	expectedIntersections := []*conetest{
		{point: Point(0, 0, -5), direction: Vector(0, 1, 0), intersectionCount: 0},
		{point: Point(0, 0, -0.25), direction: Vector(0, 1, 1), intersectionCount: 2},
		{point: Point(0, 0, -0.25), direction: Vector(0, 1, 0), intersectionCount: 4},
	}

	for _, v := range expectedIntersections {
		xs := c.localIntersect(NewRay(v.point, v.direction))

		if len(xs) != v.intersectionCount {
			t.Errorf("Intersecting a cone's end caps.: expected Ray intersection count to be xs= %v, got %v", v.intersectionCount, len(xs))
		}
	}
}

func TestIntersectTriangleParallel(t *testing.T) {
	// Intersecting a ray parallel to the triangle.

	triangle := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))
	ray := NewRay(Point(0, -1, -2), Point(0, 1, 0))
	xs := triangle.localIntersect(ray)

	if len(xs) != 0 {
		t.Errorf("Intersecting a ray parallel to the triangle: got %v expected be %v,", len(xs), 0)
	}

	// A ray misses the p1-p3 edge.
	ray = NewRay(Point(1, 1, -2), Point(0, 0, 1))
	xs = triangle.localIntersect(ray)

	if len(xs) != 0 {
		t.Errorf("A ray misses the p1-p3 edge: got %v expected be %v,", len(xs), 0)
	}

	// A ray misses the p1-p2 edge.
	ray = NewRay(Point(-1, 1, -2), Point(0, 0, 1))
	xs = triangle.localIntersect(ray)

	if len(xs) != 0 {
		t.Errorf("A ray misses the p1-p2 edge: got %v expected be %v,", len(xs), 0)
	}

	// A ray misses the p2-p3 edge.
	ray = NewRay(Point(0, -1, -2), Point(0, 0, 1))
	xs = triangle.localIntersect(ray)

	if len(xs) != 0 {
		t.Errorf("A ray misses the p2-p3 edge: got %v expected be %v,", len(xs), 0)
	}

	// A ray strikes a triangle.
	ray = NewRay(Point(0, 0.5, -2), Point(0, 0, 1))
	xs = triangle.localIntersect(ray)

	if !(len(xs) == 1 && xs[0].t == 2) {
		t.Errorf("A ray strikes a triangle: got %v expected be %v,", xs[0].t, 2)
	}
}

func TestIntersectionAllowedUnionCSG(t *testing.T) {
	// Evaluating the Rule for a CSG Union Operation.
	testStruct := []struct {
		op     string
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{"union", true, true, true, false},
		{"union", true, true, false, true},
		{"union", true, false, true, false},
		{"union", true, false, false, true},
		{"union", false, true, true, false},
		{"union", false, true, false, false},
		{"union", false, false, true, true},
		{"union", false, false, false, true},
	}

	for _, test := range testStruct {
		if !(IntersectionAllowed(test.op, test.lhit, test.inl, test.inr) == test.result) {
			t.Errorf("Evaluating the Rule for a CSG Union Operation: got %v expected be %v,",
				IntersectionAllowed(test.op, test.lhit, test.inl, test.inr), test.result)
		}
	}
}

func TestIntersectionAllowedIntersectionCSG(t *testing.T) {
	// Evaluating the Rule for a CSG Intersect Operation.
	testStruct := []struct {
		op     string
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{"intersection", true, true, true, true},
		{"intersection", true, true, false, false},
		{"intersection", true, false, true, true},
		{"intersection", true, false, false, false},
		{"intersection", false, true, true, true},
		{"intersection", false, true, false, true},
		{"intersection", false, false, true, false},
		{"intersection", false, false, false, false},
	}

	for _, test := range testStruct {
		if !(IntersectionAllowed(test.op, test.lhit, test.inl, test.inr) == test.result) {
			t.Errorf("Evaluating the Rule for a CSG Intersect Operation: got %v expected be %v,",
				IntersectionAllowed(test.op, test.lhit, test.inl, test.inr), test.result)
		}
	}
}

func TestIntersectionAllowedDifferenceCSG(t *testing.T) {
	// Evaluating the Rule for a CSG Difference Operation.
	testStruct := []struct {
		op     string
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{"difference", true, true, true, false},
		{"difference", true, true, false, true},
		{"difference", true, false, true, false},
		{"difference", true, false, false, true},
		{"difference", false, true, true, true},
		{"difference", false, true, false, true},
		{"difference", false, false, true, false},
		{"difference", false, false, false, false},
	}

	for _, test := range testStruct {
		if !(IntersectionAllowed(test.op, test.lhit, test.inl, test.inr) == test.result) {
			t.Errorf("Evaluating the Rule for a CSG Difference Operation: got %v expected be %v,",
				IntersectionAllowed(test.op, test.lhit, test.inl, test.inr), test.result)
		}
	}
}

func TestFilterCSGIntersections(t *testing.T) {
	// Filtering a list of intersections.
	s1 := NewSphere()
	s2 := NewCube()
	xs := []*Intersection{
		NewIntersection(1, s1),
		NewIntersection(2, s2),
		NewIntersection(3, s1),
		NewIntersection(4, s2),
	}
	c1 := NewCSG("union", s1, s2)
	c2 := NewCSG("intersection", s1, s2)
	c3 := NewCSG("difference", s1, s2)

	rXs1 := FilterIntersections(c1, xs)
	rXs2 := FilterIntersections(c2, xs)
	rXs3 := FilterIntersections(c3, xs)

	if len(xs) != 2 && rXs1[0] != xs[0] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
	if len(xs) != 2 && rXs1[1] != xs[3] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
	if len(xs) != 2 && rXs2[0] != xs[1] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
	if len(xs) != 2 && rXs2[1] != xs[2] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
	if len(xs) != 2 && rXs3[0] != xs[0] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
	if len(xs) != 2 && rXs3[1] != xs[1] {
		t.Errorf("Filtering a list of intersections: got %v expected be %v,", rXs1[0], xs[0])
	}
}

func TestRayMissesCSG(t *testing.T) {
	// A ray misses a CSG object.

	c := NewCSG("union", NewSphere(), NewCube())
	r := NewRay(Point(0, 2, -5), Vector(0, 0, 1))
	xs := c.localIntersect(r)
	if !(len(xs) == 0) {
		t.Errorf("A ray misses a CSG object: got %v expected be %v,", len(xs), 0)
	}
}

func TestRayHitsCSG(t *testing.T) {
	//A ray hits a CSG object.
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, 0.5))
	c := NewCSG("union", s1, s2)
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := c.localIntersect(r)

	if !(len(xs) == 2) {
		t.Errorf("A ray hits a CSG object: got %v expected be %v,", len(xs), 2)
	}
	if !(xs[0].t == 4) {
		t.Errorf("A ray hits a CSG object: got %v expected be %v,", xs[0].t, 4)
	}
	if !(xs[0].object.GetID() == s1.GetID()) {
		t.Errorf("A ray hits a CSG object: got %v expected be %v,", xs[0].object.GetID(), s1.GetID())
	}
	if !(xs[1].t == 6.5) {
		t.Errorf("A ray hits a CSG object: got %v expected be %v,", xs[1].t, 6.5)
	}
	if !(xs[1].object.GetID() == s2.GetID()) {
		t.Errorf("A ray hits a CSG object: got %v expected be %v,", xs[1].object.GetID(), s2.GetID())
	}
}
