package main

import "math"

type getColorFunc func(*Canvas, []*Color, *Tuple) *Color

// Pattern struct.
type Pattern struct {
	colors    [][]*Color
	funcs     []getColorFunc
	transform Matrix
	canvas    *Canvas
}

// ColorAtObject calculates the end color based on a worldPoint of the object.
// Allows patterns to be transformed by converting points from world space to object space,
// and from there to pattern space, before computing the color. Added support for parent grouping.
func (pattern *Pattern) ColorAtObject(object Shape, worldPoint *Tuple) *Color {
	// objectPoint := object.Transform().MultiplyMatrixByTuple(worldPoint)
	objectPoint := WorldToObject(object, worldPoint)
	patternPoint := pattern.transform.MultiplyMatrixByTuple(objectPoint)

	return pattern.ColorAt(patternPoint)
}

// ColorAt returns a reference to Color at a specific point in world space of the pattern.
func (pattern *Pattern) ColorAt(p *Tuple) *Color {
	color := Black
	for i := 0; i < len(pattern.funcs); i++ {
		color = color.Add(pattern.funcs[i](pattern.canvas, pattern.colors[i], p))
	}
	return color
}

// StripePattern creates a new stripe patter using the stripeFunc().
func StripePattern(colors ...*Color) *Pattern {

	return NewPattern(nil, [][]*Color{colors}, stripeFunc)
}

// NewPattern returns a reference to a Pattern struct with a pattern generating function.
func NewPattern(canvas *Canvas, colors [][]*Color, getColor ...getColorFunc) *Pattern {
	return &Pattern{colors, getColor, NewIdentityMatrix(), canvas}
}

// stripeFunc defines the stripe pattern.
func stripeFunc(_ *Canvas, colors []*Color, p *Tuple) *Color {
	return colors[(int(math.Abs(p.x)))%len(colors)]
}

// SetTransform sets the transform for the pattern accordingly.
func (pattern *Pattern) SetTransform(transform Matrix) {
	pattern.transform = transform.Inverse()
}

// CheckersPattern creates a new checker pattern using the checkersFunc().
func CheckersPattern(a, b *Color) *Pattern {
	return NewPattern(nil, [][]*Color{[]*Color{a, b}}, checkersFunc)
}

// checkersFunc defines the checkers pattern.
var checkersFunc = func(_ *Canvas, colors []*Color, p *Tuple) *Color {
	if (int(p.x)+int(p.y)+int(p.z))%2 == 0 {
		return colors[0]
	}
	return colors[1]
}

// gradientFunc defines a gradient pattern.
func gradientFunc(_ *Canvas, colors []*Color, p *Tuple) *Color {
	dist := colors[1].Subtract(colors[0])
	frac := p.x - math.Floor(p.x)

	return colors[0].Add(dist.MultiplyByScalar(frac))
}

// GradientPattern creates a new gradient pattern using the gradientFunc().
func GradientPattern(a, b *Color) *Pattern {
	return NewPattern(nil, [][]*Color{[]*Color{a, b}}, gradientFunc)
}

// PatternChain chains patterns together.
func PatternChain(patterns ...*Pattern) *Pattern {
	colors := [][]*Color{}
	funcs := []getColorFunc{}
	transform := NewIdentityMatrix()
	canvas := NewCanvas(0, 0)
	for _, p := range patterns {
		for _, cs := range p.colors {
			colors = append(colors, cs)
		}
		for _, f := range p.funcs {
			funcs = append(funcs, f)
		}
		if canvas != nil {
			canvas = p.canvas
		}

		transform = transform.MultiplyMatrix(p.transform)
	}
	resultPatterns := NewPattern(canvas, colors, funcs...)
	resultPatterns.SetTransform(transform)
	return resultPatterns
}
