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

// ColorAt returns a reference to Color at a specific point in world space of the pattern.
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

// CheckersPattern creates a new checker pattern using the checkersFunc().
func CheckersPattern(a, b *Color) *Pattern {
	return NewPattern([][]*Color{[]*Color{a, b}}, checkersFunc)
}

// checkersFunc defines the checkers pattern.
var checkersFunc = func(colors []*Color, p *Tuple) *Color {
	if (int(p.x)+int(p.y)+int(p.z))%2 == 0 {
		return colors[0]
	}
	return colors[1]
}

// gradientFunc defines a gradient pattern.
func gradientFunc(colors []*Color, p *Tuple) *Color {
	dist := colors[1].Subtract(colors[0])
	frac := p.x - math.Floor(p.x)

	return colors[0].Add(dist.MultiplyByScalar(frac))
}

// GradientPattern creates a new gradient pattern using the gradientFunc().
func GradientPattern(a, b *Color) *Pattern {
	return NewPattern([][]*Color{[]*Color{a, b}}, gradientFunc)
}

// PatternChain chains patterns together.
func PatternChain(patterns ...*Pattern) *Pattern {
	colors := [][]*Color{}
	funcs := []getColorFunc{}
	transform := NewIdentityMatrix()
	for _, p := range patterns {
		for _, cs := range p.colors {
			colors = append(colors, cs)
		}
		for _, f := range p.funcs {
			funcs = append(funcs, f)
		}
		transform = transform.MultiplyMatrix(p.transform)
	}
	resultPatterns := NewPattern(colors, funcs...)
	resultPatterns.SetTransform(transform)
	return resultPatterns
}
