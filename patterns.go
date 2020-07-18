package main

import "math"

type getColorFunc func([]*Color, *Tuple) *Color

// Pattern struct.
type Pattern struct {
	colors    [][]*Color
	funcs     []getColorFunc
	transform Matrix
}

// ColorAtObject calculates the end color based on a worldPoint of the object.
func (pattern *Pattern) ColorAtObject(object Shape, worldPoint *Tuple) *Color {
	objectPoint := object.Transform().MultiplyMatrixByTuple(worldPoint)
	patternPoint := pattern.transform.MultiplyMatrixByTuple(objectPoint)

	return pattern.ColorAt(patternPoint)
}

// ColorAt returns a reference to Color at a specific point in the pattern.
func (pattern *Pattern) ColorAt(p *Tuple) *Color {
	color := Black
	for i := 0; i < len(pattern.funcs); i++ {
		color = color.Add(pattern.funcs[i](pattern.colors[i], p))
	}
	return color
}

// StripePattern creates a new stripe patter using the stripeFunc().
func StripePattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, stripeFunc)
}

// NewPattern returns a reference to a Pattern struct with a pattern generating function.
func NewPattern(colors [][]*Color, getColor ...getColorFunc) *Pattern {
	return &Pattern{colors, getColor, NewIdentityMatrix()}
}

// stripeFunc defines the stripe pattern.
func stripeFunc(colors []*Color, p *Tuple) *Color {
	return colors[(int(math.Abs(p.x)))%len(colors)]
}

// SetTransform sets the transform for the pattern accordingly.
func (pattern *Pattern) SetTransform(transform Matrix) {
	pattern.transform = transform.Inverse()
}

// CheckersPattern creates a new checker patter using the checkersFunc().
func CheckersPattern(a, b *Color) *Pattern {
	return NewPattern([][]*Color{[]*Color{a, b}}, checkersFunc)
}

// // checkersFunc defines the checkers pattern.
var checkersFunc = func(colors []*Color, p *Tuple) *Color {
	if (int(p.x)+int(p.y)+int(p.z))%2 == 0 {
		return colors[0]
	}
	return colors[1]
}
