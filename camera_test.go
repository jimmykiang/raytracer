package main

import (
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
