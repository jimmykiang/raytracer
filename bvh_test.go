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
