package main

import (
	"math"
	"testing"
)

func TestCameraPixelSize(t *testing.T) {
	// The pixel size for a horizontal canvas.
	c := NewCamera(200, 125, PI/2)

	expected := 0.01

	if !floatEqual(c.pixelSize, expected) {
		t.Errorf("CameraPixelSize: expected %v to be %v", c.pixelSize, expected)
	}

	// The pixel size for a vertical canvas.
	c = NewCamera(125, 200, PI/2)
	expected = 0.01

	if !floatEqual(c.pixelSize, expected) {
		t.Errorf("CameraPixelSize: expected %v to be %v", c.pixelSize, expected)
	}
}

func TestCameraRayForPixel(t *testing.T) {
	c := NewCamera(201, 101, PI/2)

	// Constructing a ray through the center of the canvas.
	r := c.RayForPixel(100, 50)
	expected := NewRay(Point(0, 0, 0), Vector(0, 0, -1))

	if !r.Equals(expected) {
		t.Errorf("CameraRayForPixel(center): expected %v to be %v", r, expected)
	}

	// Constructing a ray through a corner of the canvas.
	r = c.RayForPixel(0, 0)
	expected = NewRay(Point(0, 0, 0), Vector(0.66519, 0.33259, -0.66851))

	if !r.Equals(expected) {
		t.Errorf("CameraRayForPixel(corner): expected %v to be %v", r, expected)
	}

	// Constructing a ray when the camera is transformed.
	c.SetTransform(RotationY(PI / 4).MultiplyMatrix(Translation(0, -2, 5)))
	r = c.RayForPixel(100, 50)

	expected = NewRay(Point(0, 2, -5), Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2))

	if !r.Equals(expected) {
		t.Errorf("CameraRayForPixel(tranformed camera): expected %v to be %v", r, expected)
	}

}
