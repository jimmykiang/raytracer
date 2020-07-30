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
	parent    Shape
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

// Intersect with the Shapes being transformed by both its own transformation and that of its parent (Group).
func (g *Group) Intersect(worldRay *Ray) []*Intersection {
	localGroupRay := worldRay.Transform(g.transform)
	return g.localIntersect(localGroupRay)
}

func (g *Group) SetTransform(transformation Matrix) {
	g.transform = transformation.Inverse()
}

func (g *Group) Transform() Matrix {
	return g.transform
}

// WorldToObject converts a point from world space to the defined (shape) object space,
// recursively taking into consideration any parent object(s) between the two spaces.
func WorldToObject(shape Shape, point *Tuple) *Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.Transform().MultiplyMatrixByTuple(point)
}

func (g *Group) GetParent() Shape {
	return g.parent
}

func (g *Group) SetParent(shape Shape) {
	g.parent = shape
}

func (g *Group) SetMaterial(material *Material) {}
func (g *Group) Material() *Material            { return nil }
func (g *Group) NormalAt(*Tuple) *Tuple         { return nil }
func (g *Group) localNormalAt(*Tuple) *Tuple    { return nil }
