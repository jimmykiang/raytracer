package main

import (
	"fmt"
	"time"
)

// planePhong tests a plane with Phong shading.
func planePhong() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()

	// The large sphere in the middle is a unit sphere, translated upward slightly and colored green.

	middle := NewSphere()
	middle.SetTransform(Translation(-0.5, 1, 0.5))
	middle.material = DefaultMaterial()
	middle.material.color = NewColor(0.1, 1, 0.5)
	middle.material.diffuse = 0.7
	middle.material.specular = 0.3

	// The smaller green sphere on the right is scaled in half.

	right := NewSphere()
	right.SetTransform(Translation(1.5, 0.5, -0.5).
		MultiplyMatrix(Scaling(0.5, 0.5, 0.5)))
	right.material = DefaultMaterial()
	right.material.color = NewColor(0.5, 1, 0.1)
	right.material.diffuse = 0.7
	right.material.specular = 0.3

	// The smallest sphere is scaled by a third, before being translated.

	left := NewSphere()
	left.SetTransform(Translation(-1.5, 0.33, -0.75).
		MultiplyMatrix(Scaling(0.33, 0.33, 0.33)))
	left.material = DefaultMaterial()
	left.material.color = NewColor(1, 0.8, 0.1)
	left.material.diffuse = 0.7
	left.material.specular = 0.3
	world := NewWorld(lights, []Shape{p1, right, left})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
