package main

import (
	"os"
	"testing"
)

func TestClock(t *testing.T) {

	canvas := Clock()

	file, err := os.Create("clock.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
