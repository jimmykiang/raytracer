package main

import (
	"os"
	"testing"
)

func TestConeScene(t *testing.T) {

	canvas := coneScene()
	file, err := os.Create("coneScene.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
