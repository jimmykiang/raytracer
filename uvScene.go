package main

import (
	"fmt"
	"time"
)

// sceneUV tests UV mappings.
func sceneUV() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()
	p1.material.pattern = CheckersPattern(NewColor(1, 0, 1), NewColor(0, 1, 0))
	p1.material.pattern.SetTransform((RotationY(PI / 6)))

	// The large sphere in the middle is a unit sphere.

	middle := NewSphere()
	middle.SetTransform(Scaling(1.1, 1.1, 1.1).
		MultiplyMatrix(Translation(-0.5, 1, 0.5)).
		MultiplyMatrix(RotationX(PI / 6)))

	middle.material = DefaultMaterial()
	middle.material.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	middle.material.pattern.SetTransform(RotationZ(PI / 6).
		MultiplyMatrix(RotationY(-PI / 3)).
		MultiplyMatrix(Scaling(0.03, 1, 1)))

	middle.material.color = NewColor(0.1, 1, 0.5)
	middle.material.diffuse = 0.7
	middle.material.specular = 0.3

	world := NewWorld(lights, []Shape{p1, middle})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
