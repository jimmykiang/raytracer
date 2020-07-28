package main

import (
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	// Creating a new group.
	g := NewGroup()
	g.Transform = NewIdentityMatrix()

	if !(len(g.Children) == 0) {
		t.Errorf("Creating a new group, does not contain children shapes: got: %v, expected: %v", len(g.Children), 0)
	}
	if !(g.Transform != nil) {
		t.Errorf("Creating a new group, contains a default transformation matrix: got: %v, expected: %v", reflect.TypeOf(g.Transform), reflect.TypeOf(NewIdentityMatrix()))
	}
}
