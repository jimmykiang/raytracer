package main

import (
	"os"
	"testing"
)

func TestPatternScene(t *testing.T) {

	canvas := scenePattern()
	file, err := os.Create("scenePattern.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
