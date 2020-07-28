package main

type Group struct {
	Transform Matrix
	Children  []Shape
}

func NewGroup() *Group {

	return &Group{
		Transform: IdentityMatrix,
		Children:  make([]Shape, 0),
	}
}
