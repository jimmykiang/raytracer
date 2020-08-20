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
	label            string
	parent           Shape
	BoundingBox      *BoundingBox
	savedRay         *Ray
}

// NewGroup returns a *Group that can contain children Shapes. A group will implement the Shape interface behaviour.
func NewGroup() *Group {

	return &Group{
		transform:        IdentityMatrix,
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		BoundingBox:      NewEmptyBoundingBox(),
		children:         make([]Shape, 0),
		savedRay:         NewRay(Point(0, 0, 0), Vector(0, 0, 0)),
		id:               rand.Int(),
	}
}

// GetID returns the id of the group (shape).
func (g *Group) GetID() int {
	return g.id
}

// AddChild will add the shape as a child to the group and establish its parent relationship from the shape itself.
func (g *Group) AddChild(shapes ...Shape) {

	for i := 0; i < len(shapes); i++ {
		g.children = append(g.children, shapes[i])
		shapes[i].SetParent(g)

		// adjust boundingBox for additional shape.
		g.BoundingBox.Merge(Bounds(shapes[i]))
	}
}

func (g *Group) localIntersect(r *Ray) []*Intersection {

	if g.BoundingBox != nil && !IntersectRayWithBox(r, g.BoundingBox) {
		return nil
	}
	g.savedRay = r
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

// SetTransform applies the transformation matrix to the Shape.
func (g *Group) SetTransform(transformation Matrix) {
	g.transform = g.transform.MultiplyMatrix(transformation)
	g.inverse = g.transform.Inverse()
	g.inverseTranspose = g.inverse.Transpose()
}

// Transform returns the transformation.
func (g *Group) Transform() Matrix {
	return g.transform
}

// GetInverse returns the cached inverse matrix of the current Shape.
func (g *Group) GetInverse() Matrix {

	return g.inverse
}

// GetInverseTranspose returns the cached inverseTranspose matrix of the current Shape.
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

// NormalAt will find the normal on a child object of a group, taking into account transformations
// on both the child object and the parent(s).
func NormalAt(s Shape, worldPoint *Tuple, intersection *Intersection) *Tuple {

	// Transform point from world to object space, including recursively traversing any parent object
	// transforms.
	localPoint := WorldToObject(s, worldPoint)

	// Normal in local space given the shape's implementation.
	objectNormal := s.localNormalAt(localPoint, intersection)

	// Convert normal from object space back into world space, again recursively applying any
	// parent transforms.
	return NormalToWorld(s, objectNormal)
}

// GetParent returns the parent shape from this current shape.
func (g *Group) GetParent() Shape {
	return g.parent
}

// SetParent sets the parent shape from this current shape.
func (g *Group) SetParent(shape Shape) {
	g.parent = shape
}

// NormalAt is not applicable to a group. use the global NormalAt() instead.
func (g *Group) NormalAt(*Tuple, *Intersection) *Tuple {
	panic("not applicable to a group. Use NormalAt() instead.")
}
func (g *Group) localNormalAt(*Tuple, *Intersection) *Tuple {
	panic("not applicable to a group. normals are always computed by calling the concrete shape’s local_normal_at()")
}

// SetMaterial will propagate the material to the child shapes.
func (g *Group) SetMaterial(material *Material) {
	for _, c := range g.children {
		c.SetMaterial(material)
	}
}

// Material not applicable to a group.
func (g *Group) Material() *Material { panic("not applicable to a group.") }

// Bounds calculates de boundingBox of the group taking in considerantion of the group's children.
func (g *Group) Bounds() {
	g.BoundingBox = Bounds(g)
}
