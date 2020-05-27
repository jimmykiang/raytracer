package main

//Tuple of four floating points.
// Point when w == 1
// vector when w == 0
type Tuple struct {
	x, y, z, w float64
}

//Point creates a tuple representing a point (w == 1)
func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

// Vector creates a tuple representing a vector (w == 0)
func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

//Equals returns true if x, y, z, w from tuples t and o are within the error margin Epsilon.
func (t *Tuple) Equals(o *Tuple) bool {

	return floatEqual(t.x, o.x) && floatEqual(t.y, o.y) && floatEqual(t.z, o.z) && floatEqual(t.w, o.w)
}

//Add tuples.
func (t *Tuple) Add(o *Tuple) *Tuple {
	return &Tuple{
		x: t.x + o.x,
		y: t.y + o.y,
		z: t.z + o.z,
		w: t.w + o.w,
	}
}
