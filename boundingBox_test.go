package main

import (
	"math"
	"testing"
)

func TestNewEmptyBoundingBox(t *testing.T) {
	// Creating an empty bounding box.
	box := NewEmptyBoundingBox()

	expectedMinPoint := Point(math.Inf(1), math.Inf(1), math.Inf(1))
	if !(expectedMinPoint.x == box.min.x) ||
		!(expectedMinPoint.y == box.min.y) ||
		!(expectedMinPoint.z == box.min.z) {
		t.Errorf("Bounding box min value: %v, expected: %v", box.min, expectedMinPoint)
	}

	expectedMaxPoint := Point(math.Inf(-1), math.Inf(-1), math.Inf(-1))
	if !(expectedMaxPoint.x == box.max.x) ||
		!(expectedMaxPoint.y == box.max.y) ||
		!(expectedMaxPoint.z == box.max.z) {
		t.Errorf("Bounding box max value: %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestNewBoundingBoxWithVolume(t *testing.T) {
	// Test BoundingBox with initial volume.
	box := NewBoundingBoxFloat(-1, -2, -3, 3, 2, 1)

	expectedMinPoint := Point(-1, -2, -3)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("Test BoundingBox with initial volume min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(3, 2, 1)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("Test BoundingBox with initial volume max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestAddPointToBoundingBox(t *testing.T) {
	// Adding points to an empty bounding box.
	box := NewEmptyBoundingBox()
	p1 := Point(-5, 2, 0)
	p2 := Point(7, 0, -3)
	box.Add(p1)
	box.Add(p2)

	expectedMinPoint := Point(-5, 0, -3)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestAddPointToBoundingBox min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(7, 2, 0)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestAddPointToBoundingBox max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfSphere(t *testing.T) {
	// A sphere has a bounding box.
	s := NewSphere()
	box := Bounds(s)

	expectedMinPoint := Point(-1, -1, -1)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestBoundsOfSphere min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(1, 1, 1)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestBoundsOfSphere max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfPlane(t *testing.T) {
	// A plane has a bounding box.
	p := NewPlane()
	box := Bounds(p)

	expectedMinPoint := Point(math.Inf(-1), 0, math.Inf(-1))
	if !(expectedMinPoint.x == box.min.x) ||
		!(expectedMinPoint.y == box.min.y) ||
		!(expectedMinPoint.z == box.min.z) {
		t.Errorf("TestBoundsOfPlane min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(math.Inf(1), 0, math.Inf(1))
	if !(expectedMaxPoint.x == box.max.x) ||
		!(expectedMaxPoint.y == box.max.y) ||
		!(expectedMaxPoint.z == box.max.z) {
		t.Errorf("TestBoundsOfPlane max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfCube(t *testing.T) {
	// A cube has a bounding box.
	c := NewCube()
	box := Bounds(c)

	expectedMinPoint := Point(-1, -1, -1)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestBoundsOfCube min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(1, 1, 1)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestBoundsOfCube max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfInfiniteCylinder(t *testing.T) {
	// An unbounded cylinder has a bounding box.
	c := NewCylinder()
	box := Bounds(c)

	expectedMinPoint := Point(-1, math.Inf(-1), -1)
	if !(expectedMinPoint.x == box.min.x) ||
		!(expectedMinPoint.y == box.min.y) ||
		!(expectedMinPoint.z == box.min.z) {
		t.Errorf("TestBoundsOfInfiniteCylinder min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(1, math.Inf(1), 1)
	if !(expectedMaxPoint.x == box.max.x) ||
		!(expectedMaxPoint.y == box.max.y) ||
		!(expectedMaxPoint.z == box.max.z) {
		t.Errorf("TestBoundsOfInfiniteCylinder max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfFiniteCylinder(t *testing.T) {
	// A bounded cylinder has a bounding box.
	c := NewCylinder()
	c.minimum = -5
	c.maximum = 3
	box := Bounds(c)

	expectedMinPoint := Point(-1, -5, -1)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestBoundsOfFiniteCylinder min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(1, 3, 1)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestBoundsOfFiniteCylinder max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfInfiniteCone(t *testing.T) {
	// An unbounded cone has a bounding box.
	c := NewCone()
	box := Bounds(c)

	expectedMinPoint := Point(math.Inf(-1), math.Inf(-1), math.Inf(-1))
	if !(expectedMinPoint.x == box.min.x) ||
		!(expectedMinPoint.y == box.min.y) ||
		!(expectedMinPoint.z == box.min.z) {
		t.Errorf("TestBoundsOfInfiniteCylinder min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(math.Inf(1), math.Inf(1), math.Inf(1))
	if !(expectedMaxPoint.x == box.max.x) ||
		!(expectedMaxPoint.y == box.max.y) ||
		!(expectedMaxPoint.z == box.max.z) {
		t.Errorf("TestBoundsOfInfiniteCylinder max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfFiniteCone(t *testing.T) {
	// A bounded cone has a bounding box.
	c := NewCone()
	c.minimum = -5
	c.maximum = 3
	box := Bounds(c)

	expectedMinPoint := Point(-5, -5, -5)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestBoundsOfFiniteCone min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(5, 3, 5)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestBoundsOfFiniteCone max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundsOfTriangle(t *testing.T) {
	// A triangle has a bounding box.
	p1 := Point(-3, 7, 2)
	p2 := Point(6, 2, -4)
	p3 := Point(2, -1, -1)
	shape := NewTriangle(p1, p2, p3)
	box := Bounds(shape)

	expectedMinPoint := Point(-3, -1, -4)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestBoundsOfFiniteCone min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(6, 7, 2)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestBoundsOfFiniteCone max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestBoundingBoxMerge(t *testing.T) {
	// Adding one bounding box to another.
	b1 := NewBoundingBoxFloat(-5, -2, 0, 7, 4, 4)
	b2 := NewBoundingBoxFloat(8, -7, -2, 14, 2, 8)
	b1.Merge(b2)

	expectedMinPoint := Point(-5, -7, -2)
	if !(expectedMinPoint.Equals(b1.min)) {
		t.Errorf("TestBoundingBoxMerge min: got %v, expected: %v", b1.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(14, 4, 8)
	if !(expectedMaxPoint.Equals(b1.max)) {
		t.Errorf("TestBoundingBoxMerge max: got %v, expected: %v", b1.max, expectedMaxPoint)
	}
}

func TestBoundingBoxContainsPoint(t *testing.T) {
	// Checking to see if a box contains a given point.
	BoundingBox := NewBoundingBoxFloat(5, -2, 0, 11, 4, 7)

	tests := []struct {
		point  *Tuple
		result bool
	}{
		{Point(5, -2, 0), true},
		{Point(11, 4, 7), true},
		{Point(8, 1, 3), true},
		{Point(3, 0, 3), false},
		{Point(8, -4, 3), false},
		{Point(8, 1, -1), false},
		{Point(13, 1, 3), false},
		{Point(8, 5, 3), false},
		{Point(8, 1, 8), false},
	}

	for _, expectedPointInBox := range tests {
		result := BoundingBox.ContainsPoint(expectedPointInBox.point)

		if !(expectedPointInBox.result == result) {
			t.Errorf("TestBoundingBoxContainsPoint: Point %v, got %v, expected: %v",
				expectedPointInBox.point, result, expectedPointInBox.result)
		}
	}
}

func TestBoxContainsBox(t *testing.T) {
	// Checking to see if a box contains a given box.
	BoundingBox := NewBoundingBoxFloat(5, -2, 0, 11, 4, 7)

	tests := []struct {
		min    *Tuple
		max    *Tuple
		result bool
	}{
		{Point(5, -2, 0), Point(11, 4, 7), true},
		{Point(6, -1, 1), Point(10, 3, 6), true},
		{Point(4, -3, -1), Point(10, 3, 6), false},
		{Point(6, -1, 1), Point(12, 5, 8), false},
	}

	for _, expectedBoundInBox := range tests {
		result := BoundingBox.ContainsBox(NewBoundingBox(expectedBoundInBox.min, expectedBoundInBox.max))

		if !(expectedBoundInBox.result == result) {
			t.Errorf("TestBoxContainsBox: got %v, expected: %v", result, expectedBoundInBox.result)
		}
	}
}

func TestTransformBoundingBox(t *testing.T) {
	// Transforming a bounding box.
	box := NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)

	m1 := RotationX(math.Pi / 4).MultiplyMatrix(RotationY(math.Pi / 4))

	box2 := TransformBoundingBox(box, m1)

	tests := []struct {
		boxFloatValue float64
		expectedValue float64
	}{
		{box2.min.x, -1.414213562373095},
		{box2.min.y, -1.7071067811865475},
		{box2.min.z, -1.7071067811865475},
		{box2.max.x, 1.414213562373095},
		{box2.max.y, 1.7071067811865475},
		{box2.max.z, 1.7071067811865475},
	}

	for _, expected := range tests {

		if !(expected.boxFloatValue == expected.expectedValue) {
			t.Errorf("TestTransformBoundingBox: got %v, expected: %v", expected.expectedValue, expected.boxFloatValue)
		}
	}
}

func TestQueryBBTransformInParentSpace(t *testing.T) {

	// Querying a shape's bounding box in its parent's space.
	shape := NewSphere()
	shape.SetTransform(Translation(1, -3, 5))
	shape.SetTransform(Scaling(0.5, 2, 4))
	box := ParentSpaceBounds(shape)

	expectedMinPoint := Point(0.5, -5, 1)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestQueryBBTransformInParentSpace min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(1.5, -1, 9)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestQueryBBTransformInParentSpace max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestGroupBoundingBoxContainsAllItsChildren(t *testing.T) {
	// A group has a bounding box that contains its children.
	s := NewSphere()
	s.SetTransform(Translation(2, 5, -3))
	s.SetTransform(Scaling(2, 2, 2))

	c := NewCylinder()
	c.minimum = -2
	c.maximum = 2
	c.SetTransform(Translation(-4, -1, 4))
	c.SetTransform(Scaling(0.5, 1, 0.5))
	g := NewGroup()
	g.AddChild(s)
	g.AddChild(c)
	box := Bounds(g)

	expectedMinPoint := Point(-4.5, -3, -5)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestQueryBBTransformInParentSpace min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(4, 7, 4.5)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestQueryBBTransformInParentSpace max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestCSGBoundingBoxContainsAllItsChildren(t *testing.T) {
	// A CSG shape has a bounding box that contains its children.
	left := NewSphere()
	right := NewSphere()
	right.SetTransform(Translation(2, 3, 4))
	csg := NewCSG("difference", left, right)
	box := Bounds(csg)

	expectedMinPoint := Point(-1, -1, -1)
	if !(expectedMinPoint.Equals(box.min)) {
		t.Errorf("TestQueryBBTransformInParentSpace min: got %v, expected: %v", box.min, expectedMinPoint)
	}
	expectedMaxPoint := Point(3, 4, 5)
	if !(expectedMaxPoint.Equals(box.max)) {
		t.Errorf("TestQueryBBTransformInParentSpace max: got %v, expected: %v", box.max, expectedMaxPoint)
	}
}

func TestIntersectBoundingBoxWithRayAtOrigin(t *testing.T) {
	// Intersecting a ray with a bounding box at the origin.
	box := NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)

	testcases := []struct {
		origin    *Tuple
		direction *Tuple
		result    bool
	}{
		{Point(5, 0.5, 0), Vector(-1, 0, 0), true},
		{Point(-5, 0.5, 0), Vector(1, 0, 0), true},
		{Point(0.5, 5, 0), Vector(0, -1, 0), true},
		{Point(0.5, -5, 0), Vector(0, 1, 0), true},
		{Point(0.5, 0, 5), Vector(0, 0, -1), true},
		{Point(0.5, 0, -5), Vector(0, 0, 1), true},
		{Point(0, 0.5, 0), Vector(0, 0, 1), true},
		{Point(-2, 0, 0), Vector(2, 4, 6), false},
		{Point(0, -2, 0), Vector(6, 2, 4), false},
		{Point(0, 0, -2), Vector(4, 6, 2), false},
		{Point(2, 0, 2), Vector(0, 0, -1), false},
		{Point(0, 2, 2), Vector(0, -1, 0), false},
		{Point(2, 2, 0), Vector(-1, 0, 0), false},
	}

	for _, tc := range testcases {
		direction := tc.direction.Normalize()
		r := NewRay(tc.origin, direction)
		if !(tc.result == IntersectRayWithBox(r, box)) {
			t.Errorf("TestIntersectBoundingBoxWithRayAtOrigin: got %v, expected: %v",
				IntersectRayWithBox(r, box), tc.result)
		}
	}
}

func TestIntersectNonCubicBoundingBoxWithRay(t *testing.T) {
	// Intersecting a ray with a non-cubic bounding box.
	box := NewBoundingBoxFloat(5, -2, 0, 11, 4, 7)

	testcases := []struct {
		origin    *Tuple
		direction *Tuple
		result    bool
	}{
		{Point(15, 1, 2), Vector(-1, 0, 0), true},
		{Point(-5, -1, 4), Vector(1, 0, 0), true},
		{Point(7, 6, 5), Vector(0, -1, 0), true},
		{Point(9, -5, 6), Vector(0, 1, 0), true},
		{Point(8, 2, 12), Vector(0, 0, -1), true},
		{Point(6, 0, -5), Vector(0, 0, 1), true},
		{Point(8, 1, 3.5), Vector(0, 0, 1), true},
		{Point(9, -1, -8), Vector(2, 4, 6), false},
		{Point(8, 3, -4), Vector(6, 2, 4), false},
		{Point(9, -1, -2), Vector(4, 6, 2), false},
		{Point(4, 0, 9), Vector(0, 0, -1), false},
		{Point(8, 6, -1), Vector(0, -1, 0), false},
		{Point(12, 5, 4), Vector(-1, 0, 0), false},
	}

	for _, tc := range testcases {
		direction := tc.direction.Normalize()
		r := NewRay(tc.origin, direction)
		if !(tc.result == IntersectRayWithBox(r, box)) {
			t.Errorf("TestIntersectNonCubicBoundingBoxWithRay: got %v, expected: %v", IntersectRayWithBox(r, box), tc.result)
		}
	}
}

func TestIntersectRayGroupWithMiss(t *testing.T) {
	// Intersecting ray+group doesn't test children if box is missed.
	s := NewSphere()
	g := NewGroup()
	g.AddChild(s)
	g.Bounds()
	r := NewRay(Point(0, 0, -5), Point(0, 1, 0))
	g.Intersect(r)

	if !(s.savedRay.origin.x == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.x, 0)
	}
	if !(s.savedRay.origin.y == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.y, 0)
	}
	if !(s.savedRay.origin.z == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.z, 0)
	}
	if !(s.savedRay.direction.x == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.x, 0)
	}
	if !(s.savedRay.direction.y == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.y, 0)
	}
	if !(s.savedRay.direction.z == 0) {
		t.Errorf("Intersecting ray+group doesn't test children if box is missed: got %v, expected: %v",
			s.savedRay.origin.z, 0)
	}
}

func TestIntersectRayGroupWithHit(t *testing.T) {
	// Intersecting ray+group tests children if box is hit.
	s := NewSphere()
	g := NewGroup()
	g.AddChild(s)
	g.Bounds()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	g.Intersect(r)

	if !(s.savedRay.origin.x == 0) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.x, 0)
	}
	if !(s.savedRay.origin.y == 0) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.y, 0)
	}
	if !(s.savedRay.origin.z == -5) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.y, -5)
	}
	if !(s.savedRay.direction.x == 0) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.x, 0)
	}
	if !(s.savedRay.direction.y == 0) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.y, 0)
	}
	if !(s.savedRay.direction.z == 1) {
		t.Errorf("Intersecting ray+group tests children if box is hit: got %v, expected: %v",
			s.savedRay.origin.z, 1)
	}
}

func TestIntersectRayWithCSGMissesBox(t *testing.T) {
	// Intersecting ray+csg doesn't test children if box is missed.
	left := NewSphere()
	right := NewSphere()
	csg := NewCSG("difference", left, right)
	csg.Bounds()
	r := NewRay(Point(0, 0, -5), Point(0, 1, 0))
	csg.Intersect(r)

	if !(left.savedRay.direction.x == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			left.savedRay.direction.x, 0)
	}
	if !(right.savedRay.direction.x == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			right.savedRay.direction.x, 0)
	}
	if !(left.savedRay.direction.y == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			left.savedRay.direction.y, 0)
	}
	if !(right.savedRay.direction.y == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			right.savedRay.direction.y, 0)
	}
	if !(left.savedRay.direction.z == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			left.savedRay.direction.z, 0)
	}
	if !(right.savedRay.direction.z == 0) {
		t.Errorf("Intersecting ray+csg doesn't test children if box is missed: got %v, expected: %v",
			right.savedRay.direction.z, 0)
	}
}

func TestIntersectRayWithCSGHitsBox(t *testing.T) {
	// Intersecting ray+csg tests children if box is hit.
	left := NewSphere()
	right := NewSphere()
	csg := NewCSG("difference", left, right)
	csg.Bounds()
	r := NewRay(Point(0, 0, -5), Point(0, 0, 1))
	csg.Intersect(r)

	if !(left.savedRay.direction.z == 1) {
		t.Errorf("Intersecting ray+csg tests children if box is hit: got %v, expected: %v",
			left.savedRay.direction.z, 1)
	}
	if !(right.savedRay.direction.z == 1) {
		t.Errorf("Intersecting ray+csg tests children if box is hit: got %v, expected: %v",
			right.savedRay.direction.z, 1)
	}
}
