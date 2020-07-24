package main

import (
	"fmt"
	"time"
)

// cylinderScene tests cylinders with Phong shading and patterns.
func cylinderScene() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()
	p1.material.pattern = CheckersPattern(NewColor(1, 0, 1), NewColor(0, 1, 0))
	p1.material.pattern.SetTransform((RotationY(PI / 6)))

	// The large sphere in the middle is a unit sphere, translated upward slightly and colored green.

	middle := NewCylinder()
	middle.closed = false
	middle.minimum = 1
	middle.maximum = 1.5
	middle.SetTransform(Scaling(1.1, 1.1, 1.1).
		MultiplyMatrix(Translation(-0.5, 0.3, 0.5)).
		MultiplyMatrix(RotationX(PI / 6)))

	middle.material = DefaultMaterial()
	middle.material.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	middle.material.pattern.SetTransform(RotationZ(PI / 6).
		MultiplyMatrix(RotationY(-PI / 3)).
		MultiplyMatrix(Scaling(0.03, 1, 1)))

	middle.material.color = NewColor(0.1, 1, 0.5)
	middle.material.diffuse = 0.7
	middle.material.specular = 0.3

	// The smaller green sphere on the right is scaled in half.

	right := NewCylinder()
	right.closed = true
	right.minimum = 1
	right.maximum = 1.5
	right.SetTransform(Translation(1.5, 0.5, -0.5).
		MultiplyMatrix(Scaling(0.5, 0.5, 0.5)).
		MultiplyMatrix(RotationX(-PI / 8)),
	)
	right.material = DefaultMaterial()
	right.material.color = NewColor(0.5, 1, 0.1)
	right.material.diffuse = 0.7
	right.material.specular = 0.3

	// The smallest sphere is scaled by a third, before being translated.

	left := NewSphere()
	left.SetTransform(Translation(-1.5, 0.33, -0.75).
		MultiplyMatrix(Scaling(0.33, 0.33, 0.33)).
		MultiplyMatrix(RotationZ(PI / 7)),
	)
	left.material = DefaultMaterial()
	left.material.pattern = GradientPattern(NewColor(1, 1, 0), NewColor(0, 0, 1))
	left.material.pattern.SetTransform(
		Scaling(2, 1, 1).
			MultiplyMatrix(Translation(0.5, 0, 0)),
	)
	left.material.pattern = PatternChain(middle.material.pattern, left.material.pattern)

	left.material.color = NewColor(1, 0.8, 0.1)
	left.material.diffuse = 0.7
	left.material.specular = 0.3
	world := NewWorld(lights, []Shape{p1, right, left, middle})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
