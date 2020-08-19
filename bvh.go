package main

import "fmt"

// SplitBounds splits the passed bounding box perpendicular of its longest axis.
func SplitBounds(b1 *BoundingBox) (*BoundingBox, *BoundingBox) {
	// find the box's largest dimension
	dx := b1.max.x - b1.min.x
	dy := b1.max.y - b1.min.y
	dz := b1.max.z - b1.min.z

	greatest := max(dx, dy, dz)

	// variables to help construct the points on
	// the dividing plane
	x0 := b1.min.x
	y0 := b1.min.y
	z0 := b1.min.z

	x1 := b1.max.x
	y1 := b1.max.y
	z1 := b1.max.z

	// adjust the points so that they lie on the
	// dividing plane
	if greatest == dx {
		x0 = x0 + dx/2.0
		x1 = x0
	} else if greatest == dy {
		y0 = y0 + dy/2.0
		y1 = y0
	} else {
		z0 = z0 + dz/2.0
		z1 = z0
	}

	midMin := Point(x0, y0, z0)
	midMax := Point(x1, y1, z1)

	// construct and return the two halves of
	// the bounding box
	left := NewBoundingBox(b1.min, midMax)
	right := NewBoundingBox(midMin, b1.max)

	return left, right
}

// PartitionChildren will further separate the group's children shapes,
// if a children lays in the middle of the boundingBox partition, then it will remain in the original group.
func PartitionChildren(g *Group) (*Group, *Group) {
	left := NewGroup()
	right := NewGroup()
	bbound := Bounds(g)
	leftBounds, rightBounds := SplitBounds(bbound)

	remain := make([]Shape, 0)
	for i := range g.children {
		childBound := ParentSpaceBounds(g.children[i])
		if leftBounds.ContainsBox(childBound) {
			left.AddChild(g.children[i])
		} else if rightBounds.ContainsBox(childBound) {
			right.AddChild(g.children[i])
		} else {
			remain = append(remain, g.children[i])
		}
	}
	// copy over the remaining ones
	g.children = g.children[:0]
	g.children = append(g.children, remain...)

	// we should really automate bounds-recalc whenever a group is mutated...
	g.Bounds()
	left.Bounds()
	right.Bounds()
	return left, right
}

var subGroupCounter int = 0

func MakeSubGroup(g *Group, shapes ...Shape) {

	subGroupCounter++
	subgroup := NewGroup()
	subgroup.label = fmt.Sprintf("Subgroup %v", subGroupCounter)
	for i := range shapes {
		subgroup.AddChild(shapes[i])
	}
	g.AddChild(subgroup)
}
