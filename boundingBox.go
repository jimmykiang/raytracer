package main

import "math"

type BoundingBox struct {
	min *Tuple
	max *Tuple
}

func NewBoundingBoxFloat(x1, y1, z1, x2, y2, z2 float64) *BoundingBox {
	return &BoundingBox{
		min: Point(x1, y1, z1),
		max: Point(x2, y2, z2),
	}
}

func NewBoundingBox(pointA *Tuple, pointB *Tuple) *BoundingBox {
	return &BoundingBox{
		min: pointA,
		max: pointB,
	}
}

func NewEmptyBoundingBox() *BoundingBox {
	return &BoundingBox{
		min: Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		max: Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
}

func ParentSpaceBounds(shape Shape) *BoundingBox {
	BoundingBox := Bounds(shape)
	return TransformBoundingBox(BoundingBox, shape.Transform())
}

func (b *BoundingBox) ContainsPoint(p *Tuple) bool {
	return b.min.x <= p.x && b.min.y <= p.y && b.min.z <= p.z &&
		b.max.x >= p.x && b.max.y >= p.y && b.max.z >= p.z
}

func (b *BoundingBox) ContainsBox(b2 *BoundingBox) bool {
	return b.ContainsPoint(b2.min) && b.ContainsPoint(b2.max)
}

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

// Bounds returns the bounds for a given shape.
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

	default:
		return NewBoundingBoxFloat(-1, -1, -1, 1, 1, 1)
	}
}

func (b *BoundingBox) Merge(b2 *BoundingBox) {
	b.Add(b2.min)
	b.Add(b2.max)
}
