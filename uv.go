package main

import "math"

// UVCheckers encapsulates the parameters for uvcheckers.
type UVCheckers struct {
	colorA *Color
	colorB *Color
	width  float64
	height float64
}

// uvCheckers will return a data structure that encapsulates the function's parameters.
func uvCheckers(width, height float64, colorA, colorB *Color) *UVCheckers {

	return &UVCheckers{
		colorA: colorA,
		colorB: colorB,
		width:  width,
		height: height,
	}
}

func uvPatternAt(uvCheckers *UVCheckers, u, v float64) *Color {

	u2 := int(math.Floor(u * uvCheckers.width))
	v2 := int(math.Floor(v * uvCheckers.height))

	if (u2+v2)%2 == 0 {
		return uvCheckers.colorA
	}
	return uvCheckers.colorB
}

func sphericalMap(point *Tuple) (u, v float64) {

	// Compute the azimuthal angle (-π < theta <= π).
	// Angle increases clockwise as viewed from above, which is opposite of what we want, but we'll fix it later.
	theta := math.Atan2(point.x, point.z)

	// vec is the vector pointing from the sphere's origin (the world origin)
	// to the point, which will also happen to be exactly equal to the sphere's radius.
	vec := Vector(point.x, point.y, point.z)
	radius := vec.Magnitude()

	// Compute the polar angle
	// 0 <= phi <= π
	phi := math.Acos(point.y / radius)

	// -0.5 < raw_u <= 0.5
	rawU := theta / (2 * PI)

	// 0 <= u < 1
	// here's also where we fix the direction of u. Subtract it from 1,
	// so that it increases counterclockwise as viewed from above.
	u = 1 - (rawU + 0.5)

	// We want v to be 0 at the south pole of the sphere,
	// and 1 at the north pole, so we have to "flip it over"
	// by subtracting it from 1.
	v = 1 - phi/PI

	return
}
