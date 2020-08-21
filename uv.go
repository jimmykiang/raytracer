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
	} else {
		return uvCheckers.colorB
	}
}
