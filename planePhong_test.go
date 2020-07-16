package main

import (
	"os"
	"testing"
)

func TestPlanePhong(t *testing.T) {

	canvas := planePhong()

	file, err := os.Create("planePhong.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
