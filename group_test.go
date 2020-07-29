package main

import (
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	// Creating a new group.
	g := NewGroup()
	g.transform = NewIdentityMatrix()

	if !(len(g.children) == 0) {
		t.Errorf("Creating a new group, does not contain children shapes: got: %v, expected: %v", len(g.children), 0)
	}
	if !(g.transform != nil) {
		t.Errorf("Creating a new group, contains a default transformation matrix: got: %v, expected: %v", reflect.TypeOf(g.Transform), reflect.TypeOf(NewIdentityMatrix()))
	}
}

func TestGroup_AddChild(t *testing.T) {
	// Adding a child to a group.
	g := NewGroup()
	s := NewSphere()

	g.AddChild(s)

	if !(len(g.children) == 1) {
		t.Errorf("Adding a child to a group, should contain 1 child shapes: got: %v, expected: %v", len(g.children), 0)
	}

	if !(g.children[0] == s) {
		t.Errorf("Adding a child to a group, should contain the child shape: got: %v, expected: %v", reflect.TypeOf(g.children[0]), reflect.TypeOf(s))
	}

	if !(s.GetParent() == g) {
		t.Errorf("Adding a child to a group, child shape should contain (reference) parent shape (group): got: %v, expected: %v", reflect.TypeOf(s.GetParent), reflect.TypeOf(g))
	}
}
