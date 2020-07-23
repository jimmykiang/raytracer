package main

import (
	"os"
	"testing"
)

func TestCylinderScene(t *testing.T) {

	canvas := cylinderScene()
	file, err := os.Create("cylinderScene.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
