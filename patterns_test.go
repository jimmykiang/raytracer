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
