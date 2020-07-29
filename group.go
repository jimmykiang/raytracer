package main

import (
	"math/rand"
	"sort"
)

// Group will implement all the methods defined in the interface Shape becoming a Shape itself.
type Group struct {
	transform Matrix
	children  []Shape
	id        int
}

func NewGroup() *Group {

	return &Group{
		transform: IdentityMatrix,
		children:  make([]Shape, 0),
		id:        rand.Int(),
	}
}

// getId returns the id of the group (shape).
func (g *Group) getId() int {
	return g.id
}

func (g *Group) AddChild(shapes ...Shape) {

	for i := 0; i < len(shapes); i++ {
		g.children = append(g.children, shapes[i])
		shapes[i].SetParent(g)
	}
}

func (g *Group) localIntersect(r *Ray) []*Intersection {

	// intersections := []*Intersection{}
	intersections := Intersections{}
	for i := range g.children {

		xs := g.children[i].Intersect(r)
		if len(xs) > 0 {
			intersections = append(intersections, xs...)
		}
	}

	if len(intersections) > 1 {

		sort.Slice(
			intersections,
			func(i, j int) bool {
				return intersections[i].t < intersections[j].t
			},
		)
	}

	return intersections
}

func (g *Group) SetMaterial(material *Material)   {}
func (g *Group) SetTransform(Matrix)              {}
func (g *Group) Transform() Matrix                { return nil }
func (g *Group) Material() *Material              { return nil }
func (g *Group) Intersect(r *Ray) []*Intersection { return nil }
func (g *Group) NormalAt(*Tuple) *Tuple           { return nil }
func (g *Group) localNormalAt(*Tuple) *Tuple      { return nil }
func (g *Group) GetParent() Shape                 { return nil }
func (g *Group) SetParent(shape Shape)            {}
