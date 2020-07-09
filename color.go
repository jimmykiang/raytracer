package main

// White color.
var White = NewColor(1, 1, 1)

// Black ...
var Black = NewColor(0, 0, 0)

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

// Multiply operation for 2 colors (resulting in a blend of colors).
func (c *Color) Multiply(o *Color) *Color {
	return NewColor(c.r*o.r, c.g*o.g, c.b*o.b)
}

// Equals returns true if the r, g, b from tuples t and o are within the error margin Epsilon.
func (c *Color) Equals(o *Color) bool {
	return floatEqual(c.r, o.r) && floatEqual(c.g, o.g) && floatEqual(c.b, o.b)
}

// String formats a color as a string limit to 8 characters.
func (c *Color) String() string {
	return "c(" + floatToString(c.r, 8) + "," + floatToString(c.g, 8) + "," + floatToString(c.b, 8) + ")"
}

//  colorToStringFormat converts the pixel color (range from 0.0 to 1.0 float64) r,g,b information
//  scaled into a range from (0 to 255) in a specific string format, for example: "255 128 13"
func (c *Color) colorToStringFormat() string {
	return floatToUint8String(c.r) + " " + floatToUint8String(c.g) + " " + floatToUint8String(c.b)
}
