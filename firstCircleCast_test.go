package main

import (
	"os"
	"testing"
)

func TestFirstCircleCast(t *testing.T) {

	canvas := CircleCast()

	file, err := os.Create("firstCircleCast.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
