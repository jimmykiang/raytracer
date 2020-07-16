package main

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

// StripePattern ...
func StripePattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, stripeFunc)
}

// NewPattern ...
func NewPattern(colors [][]*Color, getColor ...getColorFunc) *Pattern {
	return &Pattern{colors, getColor, NewIdentityMatrix()}
}

// stripeFunc defines the stripe pattern.
func stripeFunc(colors []*Color, p *Tuple) *Color {
	return colors[(int(abs(p.x)))%len(colors)]
}

// SetTransform sets the transform for the pattern accordingly.
func (pattern *Pattern) SetTransform(transform Matrix) {
	pattern.transform = transform.Inverse()
}
