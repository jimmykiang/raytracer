package main

// World creates an struct containing slices of Shape and PointLight.
type World struct {
	lights  []*PointLight
	objects []Shape
}

// DefaultWorld creates a world with default values.
func DefaultWorld() *World {
	light := NewPointLight(Point(-10, 10, -10), NewColor(1, 1, 1))
	s1 := NewSphere()
	s1.material.color = NewColor(0.8, 1.0, .6)
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2

	s2 := NewSphere()
	s2.SetTransform(Scaling(.5, .5, .5))
	return NewWorld([]*PointLight{light}, []Shape{s1, s2})
}

// NewWorld returns a World pointer.
func NewWorld(lights []*PointLight, objects []Shape) *World {
	return &World{lights, objects}
}

// Intersect returns the intersections between a ray and the objects of the world struct.
func (world *World) Intersect(ray *Ray) Intersections {
	intersections := []*Intersection{}

	for _, object := range world.objects {
		intersections = append(intersections, object.Intersect(ray)...)
	}

	xs := NewIntersections(intersections)

	return xs

}
