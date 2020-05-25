package main

//Tuple of four floating points.
type Tuple struct {
	x, y, z, w float64
}

//Equals returns true if x, y, z, w from tuples t and o are within the error margin Epsilon.
func (t *Tuple) Equals(o *Tuple) bool {

	return floatEqual(t.x, o.x) && floatEqual(t.y, o.y) && floatEqual(t.z, o.z) && floatEqual(t.w, o.w)
}
