package main

import (
	"os"
	"testing"
)

func TestPhongSphere(t *testing.T) {

	canvas := PhongSphere()

	file, err := os.Create("phongSphere.ppm")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	file.WriteString(canvas.ToPPM())
	file.Close()
}
