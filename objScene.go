package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

// objWorld tests triangles from wavefront obj data.
func objWorld() *Canvas {
	start := time.Now()
	lights := []*PointLight{
		NewPointLight(Point(0, 10, -15), NewColor(0, 1, 1)),
		NewPointLight(Point(0, 10, 0), NewColor(1, 0.5, 0.5)),
	}
	p1 := NewPlane()
	p1.material.pattern = CheckersPattern(NewColor(0.6, 0.6, 0.6), NewColor(0.1, 0.1, 0.1))
	p1.material.pattern.SetTransform((RotationY(PI / 6)))
	p1.material.reflective = 0.5

	// objBytes, _ := ioutil.ReadFile("gopher.obj")
	objBytes, _ := ioutil.ReadFile("MKIII.obj")
	obj := parseObjData(string(objBytes)).objToGroup()

	objMaterial := DefaultMaterial()

	objMaterial.color = NewColor(0.77, 0.62, 0.24)
	// objMaterial.pattern = CheckersPattern(NewColor(0.8, 0.8, 0.8), NewColor(0.2, 0.2, 0.2))
	// objMaterial.pattern.SetTransform((RotationX(PI / 6)).
	// 	MultiplyMatrix(Scaling(0.3, 0.3, 0.3)),
	// )
	objMaterial.ambient = 0.25
	objMaterial.diffuse = 0.7
	objMaterial.specular = 0.6
	objMaterial.shininess = 51.2
	obj.SetMaterial(objMaterial)

	obj.SetTransform(Translation(5, 1, -12).
		MultiplyMatrix(RotationY(0)),
	)

	// obj.SetTransform(
	// 	RotationY(PI / 3),
	// )

	world := NewWorld(lights, []Shape{p1, obj})

	camera := NewCamera(1000, 500, PI/3)
	camera.SetTransform(ViewTransform(Point(0, 1.5, -15), Point(0, 3, 0), Vector(0, 1, 0)))

	// canvas := camera.Render(world, defaultRecursionDepth)
	canvas := camera.RenderWithThreadPool(world, defaultRecursionDepth)

	fmt.Println("Render time: ", time.Now().Sub(start))

	return canvas
}
