package main

import (
	"fmt"
	"time"
)

const defaultRecursionDepth = 4

func sixPhongSpheres() *Canvas {
	start := time.Now()

	// The floor is an extremely flattened sphere with a matte texture.
	floor := NewSphere()
	floor.SetTransform(Scaling(10, 0.01, 10))
	floor.material = DefaultMaterial()
	floor.material.color = NewColor(1, 0.9, 0.9)
	floor.material.specular = 0

	// The wall on the left has the same scale and color as the floor, but is also
	// rotated and translated into place.
	// The wall needs to be scaled, then rotated in x, then rotated in y, and lastly translated, so
	// the transformations are multiplied in the reverse order.
	leftWall := NewSphere()
	leftWall.SetTransform(Translation(0, 0, 5).
		MultiplyMatrix(RotationY(-PI / 4)).
		MultiplyMatrix(RotationX(PI / 2)).
		MultiplyMatrix(Scaling(10, 0.01, 10)))

	leftWall.material = floor.material

	// The wall on the right is identical to the left wall, but is rotated the opposite	direction in y.
	rightWall := NewSphere()
	rightWall.SetTransform(Translation(0, 0, 5).
		MultiplyMatrix(RotationY(PI / 4)).
		MultiplyMatrix(RotationX(PI / 2)).
		MultiplyMatrix(Scaling(10, 0.01, 10)))

	rightWall.material = floor.material

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

	// The light source is white, shining from above and to the left.
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(1.0, 1.0, 1.0)),
	}
	world := NewWorld(lights, []Shape{floor, leftWall, rightWall, middle, right, left})

	// And the camera is configured like so:
	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)

	fmt.Println("Render time: ", time.Now().Sub(start))
	return canvas
}
