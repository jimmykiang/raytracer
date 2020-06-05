package main

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	h := 10
	w := 20
	canvas := NewCanvas(w, h)

	if len(canvas.pixels) != h {
		t.Errorf("NewCanvas: height of canvas should be %v but got %v", h, len(canvas.pixels))
	}
	defaultBlackColor := NewColor(0, 0, 0)

	for y, row := range canvas.pixels {
		if len(row) != w {
			t.Errorf("NewCanvas: width of canvas should be %v but got %v", w, len(row))
		}
		for x, xPixel := range row {
			if !xPixel.Equals(defaultBlackColor) {
				t.Errorf("NewCanvas: pixel at %v,%v is not of default color", x, y)
			}
		}
	}
}

func TestCanvasPixels(t *testing.T) {
	h := 10
	w := 20
	canvas := NewCanvas(h, w)

	red := NewColor(1, 0, 0)
	x, y := 2, 3
	canvas.WritePixel(x, y, red)
	pixelColor := canvas.PixelAt(x, y)
	if !pixelColor.Equals(red) {
		t.Errorf("CanvasPixels: pixel %v at %v,%v should be %v", pixelColor.String(), x, y, red.String())
	}
}

func TestPPMHeader(t *testing.T) {
	h := 5
	w := 3
	canvas := NewCanvas(h, w)

	result := canvas.ToPPM()
	expected := "P3\n5 3\n255\n"

	if !strings.Contains(result, expected) {
		t.Errorf("TestPPMHeader: result %v should contain %v", result, expected)
	}
}
