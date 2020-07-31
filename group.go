package main

import (
	"math"
	"math/rand"
	"sort"
)

// Group will implement all the methods defined in the interface Shape becoming a Shape itself.
type Group struct {
	transform   Matrix
	children    []Shape
	id          int
	parent      Shape
	BoundingBox *BoundingBox
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

	if g.BoundingBox != nil && !g.IntersectRayWithBox(r) {
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

func (g *Group) IntersectRayWithBox(ray *Ray) bool {

	xtmin, xtmax := checkAxisForBB(ray.origin.x, ray.direction.x, g.BoundingBox.min.x, g.BoundingBox.max.x)
	ytmin, ytmax := checkAxisForBB(ray.origin.y, ray.direction.x, g.BoundingBox.min.x, g.BoundingBox.max.x)
	ztmin, ztmax := checkAxisForBB(ray.origin.z, ray.direction.x, g.BoundingBox.min.x, g.BoundingBox.max.x)

	tmin := max(xtmin, ytmin, ztmin)
	tmax := min(xtmax, ytmax, ztmax)
	return tmin < tmax
}
func checkAxisForBB(origin, direction, minBB, maxBB float64) (min float64, max float64) {
	tminNumerator := minBB - origin
	tmaxNumerator := maxBB - origin
	var tmin, tmax float64
	if math.Abs(direction) >= EPSILON {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}
	if tmin > tmax {
		// swap
		temp := tmin
		tmin = tmax
		tmax = temp
	}
	return tmin, tmax
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

// WorldToObject converts a Point from world space to the defined (shape) object space,
// recursively taking into consideration any parent object(s) between the two spaces.
func WorldToObject(shape Shape, point *Tuple) *Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.Transform().MultiplyMatrixByTuple(point)
}

// NormalToWorld receives a normal vector in object space and transform it to world space,
// taking into consideration any parent objects between the two spaces.
func NormalToWorld(shape Shape, normal *Tuple) *Tuple {

	normal = shape.Transform().Transpose().MultiplyMatrixByTuple(normal)
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

func (g *Group) SetMaterial(material *Material) {}
func (g *Group) Material() *Material            { return nil }
