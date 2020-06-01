package main

// Color represents R,G,B values between 0 and 1.
type Color struct {
	r, g, b float64
}

// NewColor returns a *Color.
func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

// Add 2 colors.
func (c *Color) Add(o *Color) *Color {
	return NewColor(c.r+o.r, c.g+o.g, c.b+o.b)
}

// Subtract operation for 2 colors.
func (c *Color) Subtract(o *Color) *Color {
	return NewColor(c.r-o.r, c.g-o.g, c.b-o.b)
}

// MultiplyByScalar operation by a scalar.
func (c *Color) MultiplyByScalar(scalar float64) *Color {
	return NewColor(c.r*scalar, c.g*scalar, c.b*scalar)
}

// Equals returns true if the r, g, b from tuples t and o are within the error margin Epsilon.
func (c *Color) Equals(o *Color) bool {
	return floatEqual(c.r, o.r) && floatEqual(c.g, o.g) && floatEqual(c.b, o.b)
}
