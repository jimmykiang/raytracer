package main

import "fmt"

// Canvas contains the color information for every displayable pixel.
type Canvas struct {
	width, height    int
	pixels           [][]*Color
	originX, originY int
}

// NewCanvas creates and returns a new canvas reference, all pixels are black.
func NewCanvas(width int, height int) *Canvas {
	canvas := &Canvas{width: width,
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
