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
	width := 5
	height := 3
	canvas := NewCanvas(width, height)

	result := canvas.ToPPM()
	expected := "P3\n5 3\n255\n"

	if !strings.Contains(result, expected) {
		t.Errorf("TestPPMHeader: result %v should contain %v", result, expected)
	}
}

func TestPPMPixelData(t *testing.T) {
	width := 5
	height := 3
	canvas := NewCanvas(width, height)
	c1 := NewColor(1.5, 0, 0)
	c2 := NewColor(0, 0.5, 0)
	c3 := NewColor(-0.5, 0, 1)

	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	result := canvas.ToPPM()
	expected :=
		`
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 127 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	if !strings.Contains(result, expected) {
		t.Errorf("TestPPMPixelData: result %v should contain %v", result, expected)
	}
}

// Ensure that pixel data lines do not exceed 70 characters.
func TestPPMPixelDataSplitLines(t *testing.T) {
	width := 10
	height := 2
	canvas := NewCanvas(width, height)
	c := NewColor(1, 0.8, 0.6)

	for i := 0; i < height; i++ {

		for j := 0; j < width; j++ {
			canvas.WritePixel(j, i, c)
		}
	}

	canvas.ToPPM()

	result := canvas.ToPPM()
	expected := `
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	if !strings.Contains(result, expected) {
		t.Errorf("TestPPMPixelDataSplitLines: result %v should contain %v", result, expected)
	}
}

func TestTerminateByNewLineCharacter(t *testing.T) {
	width := 5
	height := 3
	canvas := NewCanvas(width, height)

	ppm := canvas.ToPPM()
	newlineIndex := strings.LastIndex(ppm, "\n")
	expectedIndex := len(ppm) - 1

	if newlineIndex != expectedIndex {
		t.Errorf("TestTerminateByNewLineCharacter: result %v should contain %v", newlineIndex, expectedIndex)
	}
}

func TestReadfileWithWrongMagicNumber(t *testing.T) {
	// Reading a file with the wrong magic number.

	ppm :=
		`P32
1 1
255
0 0 0`

	expectedError := "Incorrect magic number at line 1: expected P3"
	_, err := canvasFromPPM(ppm)

	if err == nil {
		t.Errorf("Reading a file with the wrong magic number: result %v should contain %v",
			err, expectedError)
	}
}

func TestReadPPMReturnsCanvasRightSize(t *testing.T) {
	// Reading a PPM returns a canvas of the right size.

	ppm :=
		`P3
10 2
255
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`

	canvas, err := canvasFromPPM(ppm)

	if !(err == nil) {
		t.Errorf("Reading a PPM returns a canvas of the right size: result %v should contain %v",
			err, nil)
	}

	if !(canvas.width == 10) {
		t.Errorf("Reading a PPM returns a canvas of the right size: result %v should contain %v",
			canvas, canvas.width == 10)
	}
	if !(canvas.height == 2) {
		t.Errorf("Reading a PPM returns a canvas of the right size: result %v should contain %v",
			canvas, canvas.height == 2)
	}
}

func TestReadPixelDataFromPPM(t *testing.T) {
	// Reading pixel data from a PPM file.

	ppm :=
		`P3
4 3
255
255 127 0 0 127 255 127 255 0 255 255 255
0 0 0 255 0 0 0 255 0 0 0 255
255 255 0 0 255 255 255 0 255 127 127 127`

	canvas, err := canvasFromPPM(ppm)

	type testStruct struct {
		x             int
		y             int
		expectedColor *Color
	}

	expectedTest := []testStruct{

		{x: 0, y: 0, expectedColor: NewColor(1, 0.498039, 0)},
		{x: 1, y: 0, expectedColor: NewColor(0, 0.498039, 1)},
		{x: 2, y: 0, expectedColor: NewColor(0.498039, 1, 0)},
		{x: 3, y: 0, expectedColor: NewColor(1, 1, 1)},
		{x: 0, y: 1, expectedColor: NewColor(0, 0, 0)},
		{x: 1, y: 1, expectedColor: NewColor(1, 0, 0)},
		{x: 2, y: 1, expectedColor: NewColor(0, 1, 0)},
		{x: 3, y: 1, expectedColor: NewColor(0, 0, 1)},
		{x: 0, y: 2, expectedColor: NewColor(1, 1, 0)},
		{x: 1, y: 2, expectedColor: NewColor(0, 1, 1)},
		{x: 2, y: 2, expectedColor: NewColor(1, 0, 1)},
		{x: 3, y: 2, expectedColor: NewColor(0.498039, 0.498039, 0.498039)},
	}

	if !(err == nil) {
		t.Errorf("Reading pixel data from a PPM file: result %v should contain %v",
			err, nil)
	}

	for _, val := range expectedTest {

		if !((canvas).PixelAt(val.x, val.y).Equals(val.expectedColor)) {
			t.Errorf("Reading pixel data from a PPM file, got: %v and expected to be %v",
				(canvas).PixelAt(val.x, val.y), val.expectedColor)
		}
	}
}

func TestPPMParseIgnoreCommentLines(t *testing.T) {
	// PPM parsing ignores comment lines.

	ppm :=
		`P3
# this is a comment
2 1
# this, too
255
# another comment
255 255 255
# oh, no, comments in the pixel data!
255 0 255`

	canvas, err := canvasFromPPM(ppm)

	type testStruct struct {
		x             int
		y             int
		expectedColor *Color
	}

	expectedTest := []testStruct{

		{x: 0, y: 0, expectedColor: NewColor(1, 1, 1)},
		{x: 1, y: 0, expectedColor: NewColor(1, 0, 1)},
	}

	if !(err == nil) {
		t.Errorf("Reading pixel data from a PPM file: result %v should contain %v",
			err, nil)
	}

	for _, val := range expectedTest {

		if !((canvas).PixelAt(val.x, val.y).Equals(val.expectedColor)) {
			t.Errorf("PPM parsing ignores comment lines, got: %v and expected to be %v",
				(canvas).PixelAt(val.x, val.y), val.expectedColor)
		}
	}
}

func TestPPMParseAllowRGBTripletSpanLine(t *testing.T) {
	// PPM parsing allows an RGB triple to span lines.

	ppm :=
		`P3
1 1
255
51
153

204`

	canvas, err := canvasFromPPM(ppm)

	type testStruct struct {
		x             int
		y             int
		expectedColor *Color
	}

	expectedTest := []testStruct{

		{x: 0, y: 0, expectedColor: NewColor(0.2, 0.6, 0.8)},
	}

	if !(err == nil) {
		t.Errorf("PPM parsing allows an RGB triple to span lines: result %v should contain %v",
			err, nil)
	}

	for _, val := range expectedTest {

		if !((canvas).PixelAt(val.x, val.y).Equals(val.expectedColor)) {
			t.Errorf("PPM parsing allows an RGB triple to span lines, got: %v and expected to be %v",
				(canvas).PixelAt(val.x, val.y), val.expectedColor)
		}
	}
}
