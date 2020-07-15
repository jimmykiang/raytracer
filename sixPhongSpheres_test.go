package main

import (
	"os"
	"testing"
)

// Constructed from six spheres.
func TestSixPhongSpheres(t *testing.T) {

	canvas := sixPhongSpheres()

	file, err := os.Create("sixPhongSpheres.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
