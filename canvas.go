package main

import (
	"errors"
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

// WriteTuple assign a color to a pixel at a specified tuple (point) position relative from the canvas X, Y origin.
func (canvas *Canvas) WriteTuple(point *Tuple, c *Color) {
	canvas.WritePixel(int(point.x)+canvas.originX, (int(point.y) + canvas.originY), c)
}

// SetOrigin sets the origin of the canvas.
func (canvas *Canvas) SetOrigin(x, y int) {
	if !canvas.checkBounds(x, y) {
		return
	}

	canvas.originX = x
	canvas.originY = y
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
	lines := [][]string{}

	for y := 0; y < canvas.height; y++ {

		lines = append(lines, []string{})
		for x := 0; x < canvas.width; x++ {
			pixelColorStringFormat := canvas.pixels[y][x].colorToStringFormat()

			lines[len(lines)-1] = append(lines[len(lines)-1], pixelColorStringFormat)
		}
	}

	var resultString strings.Builder
	resultString.WriteString("P3\n")
	resultString.WriteString(fmt.Sprintf("%d %d\n", canvas.width, canvas.height))
	resultString.WriteString("255\n")

	for _, stringArray := range lines {

		for _, splittedLineString := range split(strings.Join(stringArray, " "), 70) {
			resultString.WriteString(splittedLineString + "\n")
		}
	}

	return resultString.String()
}

func canvasFromPPM(data string) (*Canvas, error) {

	var canvas *Canvas
	var width, height int
	var err error

	lines := strings.Split(data, "\n")

	if strings.TrimSpace(lines[0]) != "" {
		if lines[0] != "P3" {

			return nil, errors.New("Incorrect magic number at line 1: expected P3")
		}
	}
	if strings.TrimSpace(lines[1]) != "" {
		tokenSlice := strings.Fields(strings.TrimSpace(lines[1]))

		if width, err = strconv.Atoi(tokenSlice[0]); err != nil {

			return nil, err
		}
		if height, err = strconv.Atoi(tokenSlice[1]); err != nil {

			return nil, err
		}

		canvas = NewCanvas(width, height)

	}
	return canvas, nil
}
