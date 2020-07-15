package main

import (
	"fmt"
	"sync"
	"time"
)

// PhongSphere will return a *Canvas with the circle casted onto a wall. Transformations on the sphere tested too.
func PhongSphere() *Canvas {
	start := time.Now()
	canvas := NewCanvas(500, 500)

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.width)

	half := wallSize / 2

	shape := NewSphere()
	shape.material = DefaultMaterial()
	shape.material.color = NewColor(1, 0.2, 1)
	shape.SetTransform(Scaling(1, 1, 1))
	light := NewPointLight(Point(-10, 10, -10), NewColor(1.0, 1.0, 1.0))
	var wg sync.WaitGroup

	for y := 0; y < (canvas.height); y++ {
		wg.Add(1)
		go func(y int) {
			worldY := half - pixelSize*float64(y)
			for x := 0; x < (canvas.width); x++ {
				worldX := -half + pixelSize*float64(x)

				position := Point(worldX, worldY, wallZ)
				r := NewRay(rayOrigin, position.Substract(rayOrigin).Normalize())

				xs := r.Intersect(shape)
				hit := xs.Hit()
				if hit != nil {

					point := r.Position(hit.t)
					normal := hit.object.NormalAt(point)
					eye := r.direction.Negate()

					color := Lighting(hit.object.Material(), hit.object, light, point, eye, normal, false)

					canvas.WritePixel(x, y, color)
				}

			}
			wg.Done()
		}(y)
	}

	wg.Wait()
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
