package main

import "math"

// BoundingBox struct.
type BoundingBox struct {
	min *Tuple
	max *Tuple
}

// NewBoundingBoxFloat receives min and max xyz values to returns a *BoundingBox.
func NewBoundingBoxFloat(x1, y1, z1, x2, y2, z2 float64) *BoundingBox {
	return &BoundingBox{
		min: Point(x1, y1, z1),
		max: Point(x2, y2, z2),
	}
}

// NewBoundingBox receives min and max point values to returns a *BoundingBox.
func NewBoundingBox(pointA *Tuple, pointB *Tuple) *BoundingBox {
	return &BoundingBox{
		min: pointA,
		max: pointB,
	}
}

// NewEmptyBoundingBox returns a default empty *BoundingBox.
func NewEmptyBoundingBox() *BoundingBox {
	return &BoundingBox{
		min: Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		max: Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
}

// ParentSpaceBounds transforms the shape's bounding box by the shape's transformation matrix.
func ParentSpaceBounds(shape Shape) *BoundingBox {
	BoundingBox := Bounds(shape)
	return TransformBoundingBox(BoundingBox, shape.Transform())
}

func (b *BoundingBox) ContainsPoint(p *Tuple) bool {
	return b.min.x <= p.x && b.min.y <= p.y && b.min.z <= p.z &&
		b.max.x >= p.x && b.max.y >= p.y && b.max.z >= p.z
}

// ContainsBox contains a point if each of the point's components lie between the corresponding
// min and max components.
func (b *BoundingBox) ContainsBox(b2 *BoundingBox) bool {
	return b.ContainsPoint(b2.min) && b.ContainsPoint(b2.max)
}

// TransformBoundingBox transforms the points at all eight corners of the
// cube, and then find a new bounding box that contains all eight transformed points.
func TransformBoundingBox(bbox *BoundingBox, m1 Matrix) *BoundingBox {
	p1 := bbox.min
	p2 := Point(bbox.min.x, bbox.min.y, bbox.max.z)
	p3 := Point(bbox.min.x, bbox.max.y, bbox.min.z)
	p4 := Point(bbox.min.x, bbox.max.y, bbox.max.z)
	p5 := Point(bbox.max.x, bbox.min.y, bbox.min.z)
	p6 := Point(bbox.max.x, bbox.min.y, bbox.max.z)
	p7 := Point(bbox.max.x, bbox.max.y, bbox.min.z)
	p8 := bbox.max

	// find a single bounding box that fits all of the children.
	out := NewEmptyBoundingBox()
	out.Add(m1.MultiplyMatrixByTuple(p1))
	out.Add(m1.MultiplyMatrixByTuple(p2))
	out.Add(m1.MultiplyMatrixByTuple(p3))
	out.Add(m1.MultiplyMatrixByTuple(p4))
	out.Add(m1.MultiplyMatrixByTuple(p5))
	out.Add(m1.MultiplyMatrixByTuple(p6))
	out.Add(m1.MultiplyMatrixByTuple(p7))
	out.Add(m1.MultiplyMatrixByTuple(p8))
	return out
}

// Add operation to resize a *BoundingBox.
func (b *BoundingBox) Add(p *Tuple) {
	if b.min.x > p.x {
		b.min.x = p.x
	}
	if b.min.y > p.y {
		b.min.y = p.y
	}
	if b.min.z > p.z {
		b.min.z = p.z
	}

	if b.max.x < p.x {
		b.max.x = p.x
	}
	if b.max.y < p.y {
		b.max.y = p.y
	}
	if b.max.z < p.z {
		b.max.z = p.z
	}
}

// Bounds returns the bounding box in object space for a given shape.
// For groups it will convert the bounds of all the group’s children into “group space,”
// and then combines them into a single bounding box.
func Bounds(shape Shape) *BoundingBox {
	switch val := shape.(type) {
	case *Group:
		box := NewEmptyBoundingBox()
		for i := 0; i < len(val.children); i++ {
			cbox := ParentSpaceBounds(val.children[i])
			box.Merge(cbox)
		}
		return box
	case *CSG:
		box := NewEmptyBoundingBox()
		box.Merge(ParentSpaceBounds(val.left))
		box.Merge(ParentSpaceBounds(val.right))
		return box
	case *Cube:
		return NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)
	case *Sphere:
		return NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)
	case *Plane:
		return NewBoundingBoxFloat(math.Inf(-1), 0, math.Inf(-1), math.Inf(1), 0, math.Inf(1))
	case *Cylinder:
		return NewBoundingBoxFloat(-1, val.minimum, -1, 1, val.maximum, 1)
	case *Cone:
		xzMin := math.Abs(val.minimum)
		xzMax := math.Abs(val.maximum)
		limit := xzMin
		if xzMax > limit {
			limit = xzMax
		}
		return NewBoundingBoxFloat(-limit, val.minimum, -limit, limit, val.maximum, limit)
	case *Triangle:
		BoundingBox := NewEmptyBoundingBox()
		BoundingBox.Add(val.p1)
		BoundingBox.Add(val.p2)
		BoundingBox.Add(val.p3)
		return BoundingBox
	case *smoothTriangle:
		BoundingBox := NewEmptyBoundingBox()
		BoundingBox.Add(val.p1)
		BoundingBox.Add(val.p2)
		BoundingBox.Add(val.p3)
		return BoundingBox

	default:
		return NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)
	}
}

// Merge will cause the BoundingBox to resize until it contains both points.
func (b *BoundingBox) Merge(b2 *BoundingBox) {
	b.Add(b2.min)
	b.Add(b2.max)
}

// IntersectRayWithBox test the intersection between a ray and a cubeshaped AABB at the origin.
func IntersectRayWithBox(ray *Ray, boundingBox *BoundingBox) bool {

	xtmin, xtmax := checkAxisForBB(ray.origin.x, ray.direction.x, boundingBox.min.x, boundingBox.max.x)
	ytmin, ytmax := checkAxisForBB(ray.origin.y, ray.direction.y, boundingBox.min.y, boundingBox.max.y)
	ztmin, ztmax := checkAxisForBB(ray.origin.z, ray.direction.z, boundingBox.min.z, boundingBox.max.z)

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

		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}
