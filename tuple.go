package main

import (
	"math"
)

// Tuple of four floating points.
// Point when w == 1
// vector when w == 0
type Tuple struct {
	x, y, z, w float64
}

// Point creates a tuple representing a point (w == 1)
func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

// Vector creates a tuple representing a vector (w == 0)
func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

// Equals returns true if x, y, z, w from tuples t and o are within the error margin Epsilon.
func (t *Tuple) Equals(o *Tuple) bool {

	return floatEqual(t.x, o.x) && floatEqual(t.y, o.y) && floatEqual(t.z, o.z) && floatEqual(t.w, o.w)
}

// Add tuples.
func (t *Tuple) Add(o *Tuple) *Tuple {
	return &Tuple{
		x: t.x + o.x,
		y: t.y + o.y,
		z: t.z + o.z,
		w: t.w + o.w,
	}
}

// Substract tuples.
func (t *Tuple) Substract(o *Tuple) *Tuple {
	return &Tuple{
		t.x - o.x,
		t.y - o.y,
		t.z - o.z,
		t.w - o.w,
	}
}

// Negate values contained in tuple.
func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}

// Multiply tuples.
func (t *Tuple) Multiply(o float64) *Tuple {
	return &Tuple{
		x: t.x * o,
		y: t.y * o,
		z: t.z * o,
		w: t.w * o,
	}
}

// Divide tuples.
func (t *Tuple) Divide(o float64) *Tuple {
	return &Tuple{
		x: t.x / o,
		y: t.y / o,
		z: t.z / o,
		w: t.w / o,
	}
}

// square the value
func square(v float64) float64 {
	return math.Pow(v, 2.0)
}

// Magnitude of a vector
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(square(t.x) + square(t.y) + square(t.z) + square(t.w))
}

// Normalize a vector (tuple with w == 0)
func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	if mag == 0.0 {
		return t
	}
	return Vector(t.x/mag, t.y/mag, t.z/mag)
}
