package main

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
