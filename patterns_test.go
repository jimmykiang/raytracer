package main

import "testing"

// Lighting with a pattern applied
func TestLightingWithPattern(t *testing.T) {
	m := DefaultMaterial()
	m.ambient = 1
	m.diffuse = 0
	m.specular = 0
	m.pattern = StripePattern(White, Black)
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)

	light := NewPointLight(Point(0, 0, -10), White)

	c1 := Lighting(m, NewSphere(), light, Point(0.9, 0, 0), eyev, normalv, false)
	c2 := Lighting(m, NewSphere(), light, Point(1.1, 0, 0), eyev, normalv, false)

	if !c1.Equals(White) {
		t.Errorf("LightingWithPattern(stripe): expected %v to be %v", c1, White)
	}
	if !c2.Equals(Black) {
		t.Errorf("LightingWithPattern(stripe): expected %v to be %v", c2, Black)
	}
}
func TestCheckersPattern(t *testing.T) {
	pattern := CheckersPattern(White, Black)

	points := []*Tuple{Point(0, 0, 0), Point(.99, 0, 0), Point(1.01, 0, 0), Point(0, .99, 0), Point(0, 1.01, 0), Point(0, 0, .99), Point(0, 0, 1.01)}
	expected := []*Color{White, White, Black, White, Black, White, Black}

	for i := 0; i < len(expected); i++ {
		r := pattern.ColorAt(points[i])
		if !r.Equals(expected[i]) {
			t.Errorf("CheckersPattern: expected %v to be %v", r, expected[i])
		}
	}
}

func TestGradientPattern(t *testing.T) {
	pattern := GradientPattern(White, Black)
	expected := []*Color{White, NewColor(.75, .75, .75), NewColor(0.5, 0.5, 0.5), NewColor(.25, .25, .25)}

	i := 0

	for x := 0.0; x < 1.0; x += .25 {
		c := pattern.ColorAt(Point(x, 0, 0))
		if !c.Equals(expected[i]) {
			t.Errorf("GradientPattern: expected %v to be %v", c, expected[i])
		}
		i++
	}
}
