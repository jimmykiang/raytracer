package main

import (
	"fmt"
	"time"
)

// sceneGroup builds a model of a hexagon using cylinders and spheres as a group.
func sceneGroup() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(-10, 10, -10), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 5, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()
	p1.material.pattern = CheckersPattern(NewColor(0.5, 0.5, 0.5), NewColor(0, 0, 0))
	p1.material.pattern.SetTransform((RotationY(PI / 6)))

	g := hexagon()
	// gMaterial := DefaultMaterial()
	// gMaterial.color = NewColor(0.6, 0.4, 0.3)
	// gMaterial.diffuse = 0.7
	// gMaterial.specular = 0.3
	// gMaterial.reflective = 1
	// gMaterial.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	// gMaterial.pattern.SetTransform(RotationZ(PI / 6).
	// gMaterial.refractiveIndex = 1.52
	// gMaterial.transparency = 1.0
	// gMaterial.shininess = 300

	// g.SetMaterial(gMaterial)

	g.SetTransform(Translation(0.3, 0.8, -0.4).
		MultiplyMatrix(RotationX(-PI / 6)).
		MultiplyMatrix(RotationZ(-PI / 6)),
	)

	// The large sphere in the middle is a unit sphere, translated upward slightly and colored green.

	// middle := NewSphere()
	// middle.SetTransform(Scaling(1.1, 1.1, 1.1).
	// 	MultiplyMatrix(Translation(-0.5, 1, 0.5)).
	// 	MultiplyMatrix(RotationX(PI / 6)))

	// middle.material = DefaultMaterial()
	// middle.material.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	// middle.material.pattern.SetTransform(RotationZ(PI / 6).
	// 	MultiplyMatrix(RotationY(-PI / 3)).
	// 	MultiplyMatrix(Scaling(0.03, 1, 1)))

	// middle.material.color = NewColor(0.1, 1, 0.5)
	// middle.material.diffuse = 0.7
	// middle.material.specular = 0.3

	world := NewWorld(lights, []Shape{p1, g})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}

// hexagonCorner returns a *Sphere scaled by 25% and translated bt -1 unit in z.
func hexagonCorner() *Sphere {

	s := NewSphere()

	sMaterial := DefaultMaterial()
	sMaterial.color = NewColor(0.2, 0.4, 1)
	sMaterial.diffuse = 0.7
	sMaterial.specular = 0.3
	sMaterial.reflective = 1
	sMaterial.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	sMaterial.pattern.SetTransform(RotationZ(PI / 6))

	sMaterial.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	sMaterial.pattern.SetTransform(RotationZ(PI / 6).
		MultiplyMatrix(RotationY(-PI / 3)).
		MultiplyMatrix(Scaling(0.03, 1, 1)),
	)
	sMaterial.refractiveIndex = 1.52
	sMaterial.transparency = 1.0
	sMaterial.shininess = 300

	s.SetMaterial(sMaterial)

	s.SetTransform(Translation(0, 0, -1).MultiplyMatrix(Scaling(0.25, 0.25, 0.25)))
	return s
}

// hexagonSide will scale the Cylinder by 25% in x and z, Rotate it -π⁄2 radians in z (to tip it over) and -π⁄6
// radians in y (to orient it as an edge).
func hexagonEdge() *Cylinder {

	c := NewCylinder()
	c.minimum = 0
	c.maximum = 1

	cMaterial := DefaultMaterial()
	cMaterial.color = NewColor(1, 0.2, 0.4)
	cMaterial.diffuse = 0.2
	cMaterial.specular = 0.8

	cMaterial.pattern = PatternChain(
		StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(1, 0, 0)),
		CheckersPattern(NewColor(1, 0, 1), NewColor(0, 1, 0)),
	)
	cMaterial.pattern.SetTransform(RotationZ(PI / 6).
		MultiplyMatrix(RotationY(-PI / 3)).
		MultiplyMatrix(Scaling(0.03, 1, 1)))

	cMaterial.color = NewColor(0.1, 1, 0.5)
	cMaterial.diffuse = 0.7
	cMaterial.specular = 0.3

	c.SetMaterial(cMaterial)

	c.SetTransform(
		Translation(0, 0, -1).
			MultiplyMatrix(RotationY(-PI / 6)).
			MultiplyMatrix(RotationZ(-PI / 2)).
			MultiplyMatrix(Scaling(0.25, 1, 0.25)),
	)

	return c
}

func hexagonSide() *Group {

	g := NewGroup()
	g.AddChild(hexagonCorner())
	g.AddChild(hexagonEdge())

	return g
}

func hexagon() *Group {

	g := NewGroup()
	for i := 0; i <= 5; i++ {
		h := hexagonSide()
		h.SetTransform(RotationY(float64(i) * (-PI / 3)))
		g.AddChild(h)
	}
	return g
}
