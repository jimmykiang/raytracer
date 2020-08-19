package main

import (
	"reflect"
	"testing"
)

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

func TestPartitionChildrenOfGroup(t *testing.T) {
	// Partitioning a group's children.

	s1 := NewSphere()
	s1.SetTransform(Translation(-2, 0, 0))
	s2 := NewSphere()
	s2.SetTransform(Translation(2, 0, 0))
	s3 := NewSphere()

	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)
	g.Bounds()

	left, right := PartitionChildren(g)

	if !(len(g.children) == 1) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", len(g.children), 1)
	}
	if !(len(left.children) == 1) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", len(left.children), 1)
	}
	if !(len(right.children) == 1) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", len(right.children), 1)
	}
	if !(g.children[0].GetID() == s3.GetID()) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", g.children[0].GetID(), s3.GetID())
	}
	if !(left.children[0].GetID() == s1.GetID()) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", left.children[0].GetID(), s1.GetID())
	}
	if !(right.children[0].GetID() == s2.GetID()) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", right.children[0].GetID(), s2.GetID())
	}
}

func TestCreateSubGroupFromListOfChildren(t *testing.T) {
	// Creating a sub-group from a list of children.
	s1 := NewSphere()
	s2 := NewSphere()
	g := NewGroup()
	MakeSubGroup(g, s1, s2)

	if !(len(g.children) == 1) {
		t.Errorf("Creating a sub-group from a list of children: got %v, expected: %v", len(g.children), 1)
	}
	subGroup := g.children[0].(*Group)
	if !(subGroup.children[0].GetID() == s1.GetID()) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", subGroup.children[0].GetID(), s1.GetID())
	}
	if !(subGroup.children[1].GetID() == s2.GetID()) {
		t.Errorf("Partitioning a group's children: got %v, expected: %v", subGroup.children[1].GetID(), s2.GetID())
	}
}

func TestSubDividePrimitiveDoesNothing(t *testing.T) {
	// Subdividing a primitive does nothing.
	s := NewSphere()
	Divide(s, 1)

	if !(reflect.TypeOf(s) == reflect.TypeOf(NewSphere())) {

		t.Errorf("Subdividing a primitive does nothing: got %v, expected: %v", reflect.TypeOf(s), reflect.TypeOf(NewSphere()))
	}
}
