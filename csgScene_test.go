package main

import (
	"os"
	"testing"
)

func TestCsgWorld(t *testing.T) {

	canvas := csgWorld()

	file, err := os.Create("csgWorld.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
