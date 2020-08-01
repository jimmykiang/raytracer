package main

import (
	"math"
	"testing"
)

func TestNewEmptyBoundingBox(t *testing.T) {
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

func TestBoundingBoxMerge(t *testing.T) {

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
