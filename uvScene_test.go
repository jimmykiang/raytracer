package main

import (
	"os"
	"testing"
)

func TestUVScene(t *testing.T) {

	canvas := sceneUV()
	file, err := os.Create("sceneUV.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
