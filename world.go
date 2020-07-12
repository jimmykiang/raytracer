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

// Computation is a struct for storing some precomputed values.

type Computation struct {
	t, n1, n2                                             float64
	object                                                Shape
	point, eyev, normalv, reflectv, overPoint, underPoint *Tuple
	inside                                                bool
}

// PrepareComputations precomputes the point (in world space)
// where the intersection occurred, the eye vector (pointing
// back toward the eye, or camera), and the normal vector.
func PrepareComputations(hit *Intersection, ray *Ray, xs Intersections) *Computation {
	point := ray.Position(hit.t)
	comps := &Computation{
		t:       hit.t,
		object:  hit.object,
		point:   point,
		eyev:    ray.direction.Negate(),
		normalv: hit.object.NormalAt(point),
		inside:  false,
	}
	if comps.normalv.DotProduct(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Negate()
	}

	comps.reflectv = ray.direction.Reflect(comps.normalv)
	comps.overPoint = comps.point.Add(comps.normalv.Multiply(EPSILON))
	comps.underPoint = comps.point.Substract(comps.normalv.Multiply(EPSILON))

	containers := []Shape{}

	for _, inters := range xs {
		if inters == hit {
			if len(containers) == 0 {
				comps.n1 = 1
			} else {
				comps.n1 = containers[len(containers)-1].Material().refractiveIndex
			}
		}
		if !removeIfContains(&containers, inters.object) {
			containers = append(containers, inters.object)
		}
		if inters == hit {
			if len(containers) == 0 {
				comps.n2 = 1
			} else {
				comps.n2 = containers[len(containers)-1].Material().refractiveIndex
			}
			break
		}
	}

	return comps
}

func removeIfContains(containers *[]Shape, obj Shape) bool {
	C := *containers
	for i, shape := range C {
		if shape == obj {
			for i < len(C)-1 {
				C[i] = C[i+1]
				i++
			}
			*containers = C[:len(C)-1]
			return true
		}
	}
	return false
}
