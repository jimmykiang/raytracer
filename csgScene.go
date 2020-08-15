package main

import (
	"fmt"
	"math"
	"time"
)

// csgWorld tests CSGs.
func csgWorld() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()
	p1.material.pattern = CheckersPattern(NewColor(0.6, 0.6, 0.6), NewColor(0.1, 0.1, 0.1))
	p1.material.pattern.SetTransform((RotationY(PI / 6)))
	p1.material.reflective = 0.5

	s1 := NewSphere()
	m1 := DefaultMaterial()

	s1.SetMaterial(m1)
	c1 := NewCube()
	m2 := DefaultMaterial()

	m2.color = NewColor(0.6, 0.4, 0.3)

	c1.SetMaterial(m2)
	c1.SetTransform(Translation(-0.5, 0, 0))
	c1.SetTransform(Scaling(0.75, 0.5, 2.5))

	csg := NewCSG("difference", s1, c1)
	csg.SetTransform(Translation(0, 1, 0))
	csg.SetTransform(RotationY(-math.Pi / 1.7))
	csg.SetTransform(RotationZ(-math.Pi / 4))

	world := NewWorld(lights, []Shape{p1, csg})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 4.5, -5), Point(0, 0.7, 0), Vector(0, 1, 0)))

	// canvas := camera.Render(world, defaultRecursionDepth)
	canvas := camera.RenderWithThreadPool(world, defaultRecursionDepth)

	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
