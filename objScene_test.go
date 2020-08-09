package main

import (
	"os"
	"testing"
)

func TestObjWorld(t *testing.T) {

	canvas := objWorld()

	file, err := os.Create("objWorld.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
