package main

import (
	"fmt"
	"io/ioutil"
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
	// p1.material.pattern = uvPlanarCheckersPattern(NewColor(1, 0, 1), NewColor(0, 1, 0))
	p1.material.pattern = uvAlignCheckPattern()
	p1.material.pattern.SetTransform((RotationY(PI / 2).
		MultiplyMatrix(Scaling(1.1, 1.1, 1.1))))

	// The large sphere in the middle is a unit sphere.

	middle := NewSphere()
	middle.SetTransform(Scaling(1.1, 1.1, 1.1).
		MultiplyMatrix(Translation(0, 1, 0.5)).
		MultiplyMatrix(RotationX(PI / 6)))

	middle.material = DefaultMaterial()
	// middle.material.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	// middle.material.pattern = uvSphericalCheckersPattern(NewColor(0, 0, 0), NewColor(1, 1, 1))

	// earth texture.
	ppmFileByteSlice, _ := ioutil.ReadFile("earthmap1kYolo.ppm")
	var textureCanvas *Canvas
	var err error
	if textureCanvas, err = canvasFromPPM(string(ppmFileByteSlice)); err != nil {
		panic(err)
	}
	middle.material.pattern = uvSphericalCanvasPattern(textureCanvas)

	secondSphere := NewSphere()
	secondSphere.SetTransform(Scaling(1.1, 1.1, 1.1).
		MultiplyMatrix(Translation(-0.5, 1, 0.5)).
		MultiplyMatrix(RotationX(PI / 6)))

	secondSphere.material = DefaultMaterial()
	secondSphere.material.pattern = StripePattern(NewColor(0, 0, 0), NewColor(1, 1, 1), NewColor(0.5, 0.4, 0.7))
	secondSphere.material.pattern.SetTransform(RotationZ(PI / 6).
		MultiplyMatrix(RotationY(-PI / 3)).
		MultiplyMatrix(Scaling(0.03, 1, 1)))

	middle.material.pattern.SetTransform(RotationZ(PI / 10).
		MultiplyMatrix(RotationY(-4 * PI / 4)).
		MultiplyMatrix(RotationX(PI / 4)).
		MultiplyMatrix(Scaling(1, 1, 1)))

	middle.material.pattern = PatternChain(secondSphere.material.pattern, middle.material.pattern)

	// middle.material.color = NewColor(0.1, 1, 0.5)
	// middle.material.diffuse = 0.7
	// middle.material.specular = 0.3

	right := NewCylinder()
	right.material.pattern = uvCylindricalCheckersPattern(NewColor(0, 0, 0), NewColor(1, 1, 1))
	right.material.pattern.SetTransform(Scaling(1, 1, 1))

	right.SetTransform(Translation(2, 1, 3).
		MultiplyMatrix(RotationX(-PI / 4)))

	right.closed = true
	right.minimum = 1
	right.maximum = 1.5

	left := NewCube()
	left.material.pattern = uvCubeMapAlignPattern()
	left.SetTransform(Translation(-2.5, 1.5, 3).
		MultiplyMatrix(RotationX(-PI / 4)).
		MultiplyMatrix(RotationY(PI / 4)),
	)

	g := NewGroup()
	g.AddChild(middle, right, left)

	g.Bounds()
	Divide(g, 1)

	world := NewWorld(lights, []Shape{p1, g})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0)))

	canvas := camera.Render(world, defaultRecursionDepth)
	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
