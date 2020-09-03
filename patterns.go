package main

import "math"

type getColorFunc func(*Canvas, []*Color, *Tuple) *Color

// Pattern struct.
type Pattern struct {
	colors     [][]*Color
	funcs      []getColorFunc
	transforms []Matrix
	canvas     *Canvas
}

// ColorAtObject calculates the end color based on a worldPoint of the object.
// Allows patterns to be transformed by converting points from world space to object space,
// and from there to pattern space, before computing the color. Added support for parent grouping.
func (pattern *Pattern) ColorAtObject(object Shape, worldPoint *Tuple) *Color {
	// objectPoint := object.Transform().MultiplyMatrixByTuple(worldPoint)
	objectPoint := WorldToObject(object, worldPoint)
	// patternPoint := pattern.transform.MultiplyMatrixByTuple(objectPoint)

	// return pattern.ColorAt(patternPoint)
	return pattern.ColorAt(objectPoint)
}

// ColorAt returns a reference to Color at a specific point in world space of the pattern.
func (pattern *Pattern) ColorAt(p *Tuple) *Color {
	color := Black

	// Final color at a specific point by applying individual transformations and patterns.
	for i := 0; i < len(pattern.funcs); i++ {
		patternPoint := pattern.transforms[i].MultiplyMatrixByTuple(p)
		color = color.Add(pattern.funcs[i](pattern.canvas, pattern.colors[i], patternPoint))
	}
	return color
}

// StripePattern creates a new stripe patter using the stripeFunc().
func StripePattern(colors ...*Color) *Pattern {

	return NewPattern(NewCanvas(0, 0), [][]*Color{colors}, stripeFunc)
}

// NewPattern returns a reference to a Pattern struct with a pattern generating function.
func NewPattern(canvas *Canvas, colors [][]*Color, getColor ...getColorFunc) *Pattern {
	return &Pattern{colors, getColor, []Matrix{NewIdentityMatrix()}, canvas}
}

// stripeFunc defines the stripe pattern.
func stripeFunc(_ *Canvas, colors []*Color, p *Tuple) *Color {
	return colors[(int(math.Abs(p.x)))%len(colors)]
}

// SetTransform sets the transform for the pattern accordingly.
func (pattern *Pattern) SetTransform(transform Matrix) {
	pattern.transforms[0] = transform.Inverse()
}

// CheckersPattern creates a new checker pattern using the checkersFunc().
func CheckersPattern(a, b *Color) *Pattern {
	return NewPattern(NewCanvas(0, 0), [][]*Color{[]*Color{a, b}}, checkersFunc)
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
	return NewPattern(NewCanvas(0, 0), [][]*Color{[]*Color{a, b}}, gradientFunc)
}

// PatternChain chains patterns together. For now it will apply the canvas from the last pattern as valid image mapping.
func PatternChain(patterns ...*Pattern) *Pattern {
	colors := [][]*Color{}
	funcs := []getColorFunc{}
	transforms := []Matrix{}
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

		// transform = transform.MultiplyMatrix(p.transform)
		transforms = append(transforms, p.transforms...)
	}
	resultPatterns := NewPattern(canvas, colors, funcs...)
	// resultPatterns.SetTransform(transform)
	resultPatterns.transforms = transforms
	return resultPatterns
}
