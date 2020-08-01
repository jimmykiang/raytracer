package main

import (
	"os"
	"testing"
)

func TestGroupScene(t *testing.T) {

	canvas := sceneGroup()
	file, err := os.Create("sceneGroup.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
