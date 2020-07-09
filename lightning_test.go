package main

import (
	"math"
	"testing"
)

func TestLighting(t *testing.T) {

	material := DefaultMaterial()
	position := Point(0, 0, 0)

	// Lighting with the eye between the light and the surface.
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), NewColor(1, 1, 1))
	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)
	expected := NewColor(1.9, 1.9, 1.9)
	if !result.Equals(expected) {
		t.Errorf("Lighting: (eye between light and surface) expected %v to be %v", result, expected)
	}

	// Lighting with the eye between light and surface, eye offset 45°.
	eyev = Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv = Vector(0, 0, -1)
	result = Lighting(material, NewSphere(), light, position, eyev, normalv, false)
	expected = NewColor(1, 1, 1)
	if !result.Equals(expected) {
		t.Errorf("Lighting: (eye between light and surface, eye offset 45deg) expected %v to be %v", result, expected)
	}

	// Lighting with eye opposite surface, light offset 45°.
	eyev = Vector(0, 0, -1)
	normalv = Vector(0, 0, -1)
	light = NewPointLight(Point(0, 10, -10), NewColor(1, 1, 1))
	result = Lighting(material, NewSphere(), light, position, eyev, normalv, false)
	expected = NewColor(0.7364, 0.7364, 0.7364)
	if !result.Equals(expected) {
		t.Errorf("Lighting: (eye opposite surface, light offset 45deg) expected %v to be %v", result, expected)
	}

	// Lighting with eye in the path of the reflection vector.
	eyev = Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv = Vector(0, 0, -1)
	light = NewPointLight(Point(0, 10, -10), NewColor(1, 1, 1))
	result = Lighting(material, NewSphere(), light, position, eyev, normalv, false)
	expected = NewColor(1.6364, 1.6364, 1.6364)
	if !result.Equals(expected) {
		t.Errorf("Lighting: (eye in path of reflection) expected %v to be %v", result, expected)
	}

	// Lighting with the light behind the surface.
	eyev = Vector(0, 0, -1)
	normalv = Vector(0, 0, -1)
	light = NewPointLight(Point(0, 0, 10), NewColor(1, 1, 1))
	result = Lighting(material, NewSphere(), light, position, eyev, normalv, false)
	expected = NewColor(0.1, 0.1, 0.1)
	if !result.Equals(expected) {
		t.Errorf("Lighting: (light behind surface) expected %v to be %v", result, expected)
	}

	result = Lighting(material, NewSphere(), light, position, eyev, normalv, true)
	expected = NewColor(0.1, 0.1, 0.1)

	if !result.Equals(expected) {
		t.Errorf("Lighting: (in shadow) expected %v to be %v", result, expected)
	}
}
