package main

import "testing"

func TestSplitPerfectCube(t *testing.T) {
	// Splitting a perfect cube.
	box := NewBoundingBoxFloat(-1, -4, -5, 9, 6, 5)
	left, right := SplitBounds(box)

	expectedLeftMinPoint := Point(-1, -4, -5)
	if !(expectedLeftMinPoint.Equals(left.min)) {
		t.Errorf("Splitting a perfect cube: got %v, expected: %v", left.min, expectedLeftMinPoint)
	}
	expectedLeftMaxPoint := Point(4, 6, 5)
	if !(expectedLeftMaxPoint.Equals(left.max)) {
		t.Errorf("Splitting a perfect cube: got %v, expected: %v", left.max, expectedLeftMaxPoint)
	}
	expectedRightMinPoint := Point(4, -4, -5)
	if !(expectedRightMinPoint.Equals(right.min)) {
		t.Errorf("Splitting a perfect cube: got %v, expected: %v", right.min, expectedRightMinPoint)
	}
	expectedRightMaxPoint := Point(9, 6, 5)
	if !(expectedRightMaxPoint.Equals(right.max)) {
		t.Errorf("Splitting a perfect cube: got %v, expected: %v", right.max, expectedRightMaxPoint)
	}
}

func TestSplitXWideBoundingBox(t *testing.T) {
	// Splitting an x-wide box.

	box := NewBoundingBoxFloat(-1, -2, -3, 9, 5.5, 3)
	left, right := SplitBounds(box)

	expectedLeftMinPoint := Point(-1, -2, -3)
	if !(expectedLeftMinPoint.Equals(left.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.min, expectedLeftMinPoint)
	}
	expectedLeftMaxPoint := Point(4, 5.5, 3)
	if !(expectedLeftMaxPoint.Equals(left.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.max, expectedLeftMaxPoint)
	}
	expectedRightMinPoint := Point(4, -2, -3)
	if !(expectedRightMinPoint.Equals(right.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.min, expectedRightMinPoint)
	}
	expectedRightMaxPoint := Point(9, 5.5, 3)
	if !(expectedRightMaxPoint.Equals(right.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.max, expectedRightMaxPoint)
	}
}

func TestSplitYWideBoundingBox(t *testing.T) {
	// Splitting an x-wide box.

	box := NewBoundingBoxFloat(-1, -2, -3, 5, 8, 3)
	left, right := SplitBounds(box)

	expectedLeftMinPoint := Point(-1, -2, -3)
	if !(expectedLeftMinPoint.Equals(left.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.min, expectedLeftMinPoint)
	}
	expectedLeftMaxPoint := Point(5, 3, 3)
	if !(expectedLeftMaxPoint.Equals(left.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.max, expectedLeftMaxPoint)
	}
	expectedRightMinPoint := Point(-1, 3, -3)
	if !(expectedRightMinPoint.Equals(right.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.min, expectedRightMinPoint)
	}
	expectedRightMaxPoint := Point(5, 8, 3)
	if !(expectedRightMaxPoint.Equals(right.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.max, expectedRightMaxPoint)
	}
}

func TestSplitZWideBoundingBox(t *testing.T) {
	// Splitting an x-wide box.

	box := NewBoundingBoxFloat(-1, -2, -3, 5, 3, 7)
	left, right := SplitBounds(box)

	expectedLeftMinPoint := Point(-1, -2, -3)
	if !(expectedLeftMinPoint.Equals(left.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.min, expectedLeftMinPoint)
	}
	expectedLeftMaxPoint := Point(5, 3, 2)
	if !(expectedLeftMaxPoint.Equals(left.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", left.max, expectedLeftMaxPoint)
	}
	expectedRightMinPoint := Point(-1, -2, 2)
	if !(expectedRightMinPoint.Equals(right.min)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.min, expectedRightMinPoint)
	}
	expectedRightMaxPoint := Point(5, 3, 7)
	if !(expectedRightMaxPoint.Equals(right.max)) {
		t.Errorf("Splitting an x-wide box: got %v, expected: %v", right.max, expectedRightMaxPoint)
	}
}
