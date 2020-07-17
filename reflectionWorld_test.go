package main

import (
	"os"
	"testing"
)

func TestReflectionWorld(t *testing.T) {

	canvas := reflectionWorld()

	file, err := os.Create("reflectionWorld.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
