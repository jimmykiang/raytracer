package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Canvas contains the color information for every displayable pixel.
type Canvas struct {
	width, height    int
	pixels           [][]*Color
	originX, originY int
}

// NewCanvas creates and returns a new canvas reference, all pixels are black.
func NewCanvas(width int, height int) *Canvas {
	canvas := &Canvas{
		width:   width,
		height:  height,
		pixels:  make([][]*Color, height, height),
		originX: 0,
		originY: 0,
	}
	for y := 0; y < height; y++ {
		canvas.pixels[y] = make([]*Color, width, width)
		for x := 0; x < width; x++ {
			canvas.pixels[y][x] = NewColor(0, 0, 0)
		}
	}
	return canvas
}

// WritePixel assign a color to a pixel at a specified position.
func (canvas *Canvas) WritePixel(x, y int, c *Color) {
	if !canvas.checkBounds(x, y) {
		return
	}
	canvas.pixels[y][x] = NewColor(c.r, c.g, c.b)
}

func (canvas *Canvas) checkBounds(x, y int) (check bool) {
	check = true
	if y < 0 || y >= canvas.height {
		check = false
	}
	if x < 0 || x >= canvas.width {
		check = false
	}
	if !check {
		fmt.Println(x, y)
	}
	return
}

// PixelAt returns a color reference of a specific pixel.
func (canvas *Canvas) PixelAt(x, y int) *Color {
	if !canvas.checkBounds(x, y) {
		return nil
	}
	return canvas.pixels[y][x]
}

// ToPPM returns the canvas information into string based fo.
func (canvas *Canvas) ToPPM() string {

	// slice representing with pixel color information in this format [[r g b r g b ...], [r g b r g b ...], ...}
	// each r, g, b values ranges from 0 to 255
	// where the the outer slice contains the "y" (row) inner slices
	// which in turn contains the actual "x" (column) pixel rgb color information.
	// [row][column]
	lines := [][]rune{}

	for y := 0; y < canvas.height; y++ {

		// insert a new "inner" slice for the next "y" or "row".
		lines = append(lines, []rune{})

		for x := 0; x < canvas.width; x++ {
			pixelColorStringFormat := canvas.pixels[y][x].colorToStringFormat()

			// if the current length of the inner "x" or "column" cannot hold
			// the contents of pixelColorStringFormat without surpassing 69 characters,
			// then "insert" a new slice (new row) that will contain the pixelColor of the next line.
			if len(pixelColorStringFormat) > 69-len(lines[len(lines)-1]) {
				lines = append(lines, []rune{})
			}
			for _, ch := range pixelColorStringFormat {
				lines[len(lines)-1] = append(lines[len(lines)-1], ch)
			}
		}
	}

	ppm := []string{"P3", strconv.Itoa(canvas.width) + " " + strconv.Itoa(canvas.height), "255"}

	for _, line := range lines {
		ppm = append(ppm, string(line))
	}
	ppm = append(ppm, "\n")

	return strings.Join(ppm, "\n")
}
