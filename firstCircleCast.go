package main

import (
	"fmt"
	"sync"
	"time"
)

// CircleCast will return a *Canvas with the circle casted onto a wall. Transformations on the sphere tested too.
func CircleCast() *Canvas {
	start := time.Now()
	canvas := NewCanvas(100, 100)

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.width)

	half := wallSize / 2
	color := NewColor(1, 0, 0)
	shape := NewSphere()
	var wg sync.WaitGroup

	shape.SetTransform(Shearing(1, 0, 0, 0, 0, 0).MultiplyMatrix(Scaling(0.5, 1, 1)))

	for y := 0; y < (canvas.height); y++ {
		wg.Add(1)
		go func(y int) {
			worldY := half - pixelSize*float64(y)
			for x := 0; x < (canvas.width); x++ {
				worldX := -half + pixelSize*float64(x)

				position := Point(worldX, worldY, wallZ)
				r := NewRay(rayOrigin, position.Substract(rayOrigin).Normalize())

				xs := r.Intersect(shape)

				if xs.Hit() != nil {
					canvas.WritePixel(x, y, color)
				}

			}
			wg.Done()
		}(y)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))

	return canvas
}
