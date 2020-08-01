package main

import (
	"math/rand"
	"sort"
)

// Group will implement all the methods defined in the interface Shape becoming a Shape itself.
type Group struct {
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	children         []Shape
	id               int
	parent           Shape
	BoundingBox      *BoundingBox
}

func NewGroup() *Group {

	return &Group{
		transform:        IdentityMatrix,
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		children:         make([]Shape, 0),
		id:               rand.Int(),
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

	if g.BoundingBox != nil && !IntersectRayWithBox(r, g.BoundingBox) {
		return nil
	}

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
	localGroupRay := worldRay.Transform(g.inverse)
	return g.localIntersect(localGroupRay)
}

func (g *Group) SetTransform(transformation Matrix) {
	g.transform = g.transform.MultiplyMatrix(transformation)
	g.inverse = g.transform.Inverse()
	g.inverseTranspose = g.inverse.Transpose()
}

func (g *Group) Transform() Matrix {
	return g.transform
}

func (g *Group) GetInverse() Matrix {

	return g.inverse
}

func (g *Group) GetInverseTranspose() Matrix {

	return g.inverseTranspose
}

// WorldToObject converts a Point from world space to the defined (shape) object space,
// recursively taking into consideration any parent object(s) between the two spaces.
func WorldToObject(shape Shape, point *Tuple) *Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.GetInverse().MultiplyMatrixByTuple(point)
}

// NormalToWorld receives a normal vector in object space and transform it to world space,
// taking into consideration any parent objects between the two spaces.
func NormalToWorld(shape Shape, normal *Tuple) *Tuple {

	normal = shape.GetInverseTranspose().MultiplyMatrixByTuple(normal)
	normal.w = 0
	normal = normal.Normalize()

	if shape.GetParent() != nil {
		normal = NormalToWorld(shape.GetParent(), normal)
	}

	return normal
}

// Find the normal on a child object of a group, taking into account transformations
// on both the child object and the parent(s).
func NormalAt(s Shape, worldPoint *Tuple) *Tuple {

	// Transform point from world to object space, including recursively traversing any parent object
	// transforms.
	localPoint := WorldToObject(s, worldPoint)

	// Normal in local space given the shape's implementation.
	objectNormal := s.localNormalAt(localPoint)

	// Convert normal from object space back into world space, again recursively applying any
	// parent transforms.
	return NormalToWorld(s, objectNormal)
}

func (g *Group) GetParent() Shape {
	return g.parent
}

func (g *Group) SetParent(shape Shape) {
	g.parent = shape
}

func (g *Group) NormalAt(*Tuple) *Tuple {
	panic("not applicable to a group. Use NormalAt() instead")
}
func (g *Group) localNormalAt(*Tuple) *Tuple {
	panic("not applicable to a group. normals are always computed by calling the concrete shape’s local_normal_at()")
}

// SetMaterial will propagate the material to the child shapes.
func (g *Group) SetMaterial(material *Material) {
	for _, c := range g.children {
		c.SetMaterial(material)
	}
}
func (g *Group) Material() *Material { panic("not applicable to a group.") }
