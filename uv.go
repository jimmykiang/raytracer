package main

import "math"

// UVCheckers encapsulates the parameters for uvcheckers.
type UVCheckers struct {
	colorA *Color
	colorB *Color
	width  float64
	height float64
}

type patternType interface {
	isPattern() bool
}

func (pattern *UVCheckers) isPattern() bool {

	return true
}
func (pattern *UVAlignCheck) isPattern() bool {

	return true
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

// uvPatternAt will return the pattern's color at the given u and v coordinates,
// where both u and v are floating point numbers between 0 and 1, inclusive.
func uvPatternAt(pattern patternType, u, v float64) *Color {

	switch p := pattern.(type) {

	case *UVCheckers:
		u2 := int(math.Floor(u * p.width))
		v2 := int(math.Floor(v * p.height))

		if (u2+v2)%2 == 0 {
			return p.colorA
		}
		return p.colorB

	case *UVAlignCheck:
		// remember: v=0 at the bottom, v=1 at the top
		if v > 0.8 {

			if u < 0.2 {
				return p.ul
			}
			if u > 0.8 {
				return p.ur
			}
		} else if v < 0.2 {
			if u < 0.2 {
				return p.bl
			}
			if u > 0.8 {
				return p.br
			}
		}
		return p.main

	default:
		return nil
	}
}

// sphericalMap maps a 3D point (x, y, z) on the surface of sphere to a 2D point (u, v) on the flattened surface.
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

// TextureMap encapsulates the given uv_pattern (like uv_checkers() ) and uv_map (like spherical_map() ).
type TextureMap struct {
	uvPattern patternType
	uvMap     func(point *Tuple) (u, v float64)
}

// textureMap returns a *TextureMap struct.
func textureMap(uvPattern patternType, uvMap func(point *Tuple) (u, v float64)) *TextureMap {
	return &TextureMap{
		uvPattern: uvPattern,
		uvMap:     uvMap,
	}
}

func patternAt(textureMap *TextureMap, point *Tuple) *Color {

	u, v := textureMap.uvMap(point)
	return uvPatternAt(textureMap.uvPattern, u, v)
}

// uvSphericalCheckersFunc adapts the uvCheckers and textureMap to be set as a func to the *Pattern struct.
// only the 2 first colors from the parameter slice are processed.
func uvSphericalCheckersFunc(colors []*Color, p *Tuple) *Color {

	checkers := uvCheckers(16, 8, colors[0], colors[1])
	pattern := textureMap(checkers, sphericalMap)
	return patternAt(pattern, p)
}

// uvSphericalCheckersPattern returns the appropiate *Pattern struct.
func uvSphericalCheckersPattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, uvSphericalCheckersFunc)
}

// planarMap returns the u,v coordinates for a flattened surface.
func planarMap(point *Tuple) (u, v float64) {

	// Working Implementation different from:
	/*
	   function planar_map(p)
	    let u ← p.x mod 1
	    let v ← p.z mod 1
	    return (u, v)
	   end function
	*/

	if point.x < 0 {
		u = math.Abs(math.Floor(point.x)) + point.x
	} else {
		u = point.x - math.Floor(point.x)
	}

	if point.z < 0 {
		v = math.Abs(math.Floor(point.z)) + point.z
	} else {
		v = point.z - math.Floor(point.z)
	}

	return
}

// uvPlanarCheckersFunc adapts the uvCheckers and textureMap to be set as a func to the *Pattern struct.
// only the 2 first colors from the parameter slice are processed.
func uvPlanarCheckersFunc(colors []*Color, p *Tuple) *Color {

	checkers := uvCheckers(16, 8, colors[0], colors[1])
	pattern := textureMap(checkers, planarMap)
	return patternAt(pattern, p)
}

// uvPlanarCheckersPattern returns the appropiate *Pattern struct.
func uvPlanarCheckersPattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, uvPlanarCheckersFunc)
}

// cylindricalMap maps a 3D point (x, y, z) on the surface of cylindricalMap to a 2D point (u, v) on the flattened surface.
func cylindricalMap(point *Tuple) (u, v float64) {

	// Compute the azimuthal angle (-π < theta <= π) same as with spherical_map().
	theta := math.Atan2(point.x, point.z)
	// -0.5 < raw_u <= 0.5
	rawU := theta / (2 * PI)
	// 0 <= u < 1
	// here's also where we fix the direction of u. Subtract it from 1,
	// so that it increases counterclockwise as viewed from above.
	u = 1 - (rawU + 0.5)

	// let v go from 0 to 1 between whole units of y
	// original: let v ← p.y mod 1

	if point.y < 0 {
		v = math.Abs(math.Floor(point.y)) + point.y
	} else {
		v = point.y - math.Floor(point.y)
	}

	return
}

// uvCylindricalCheckersFunc adapts the uvCheckers and textureMap to be set as a func to the *Pattern struct.
// only the 2 first colors from the parameter slice are processed.
func uvCylindricalCheckersFunc(colors []*Color, p *Tuple) *Color {

	checkers := uvCheckers(16, 8, colors[0], colors[1])
	pattern := textureMap(checkers, cylindricalMap)
	return patternAt(pattern, p)
}

// uvCylindricalCheckersPattern returns the appropiate *Pattern struct.
func uvCylindricalCheckersPattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, uvCylindricalCheckersFunc)
}

// UVAlignCheck defines a struct for an align pattern.
type UVAlignCheck struct {
	main *Color
	ul   *Color
	ur   *Color
	bl   *Color
	br   *Color
}

// uvAlignCheck returns a *UVAlignCheck.
func uvAlignCheck(main, ul, ur, bl, br *Color) *UVAlignCheck {

	return &UVAlignCheck{main, ul, ur, bl, br}
}

// uvAlignCheckFunc adapts the patternType method and textureMap to be set as a func to the *Pattern struct.
// only the 2 first colors from the parameter slice are processed.
func uvAlignCheckFunc(_ []*Color, p *Tuple) *Color {

	// Predefined colors for UVAlignCheck.
	alignCheck := uvAlignCheck(White, Red, Yellow, Green, Cyan)
	pattern := textureMap(alignCheck, planarMap)
	return patternAt(pattern, p)
}

// uvCylindricalCheckersPattern returns the appropiate *Pattern struct.
func uvAlignCheckPattern() *Pattern {

	// Predefined colors for UVAlignCheck.
	return NewPattern([][]*Color{{White}, {Red}, {Yellow}, {Green}, {Cyan}}, uvAlignCheckFunc)
}
