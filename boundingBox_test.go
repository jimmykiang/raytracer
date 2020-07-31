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
