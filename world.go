package main

import (
	"math"
	"sort"
)

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

	if len(xs) > 1 {
		sort.Slice(xs, func(i, j int) bool { return xs[i].t < xs[j].t })
	}

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

// If the intersection’s object is already in the containers list,
// then this intersection must be exiting the object. Remove the object from the containers
// list in this case. Otherwise, the intersection is entering the object, and
// the object should be added to the end of the list.
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

// ShadeHit returns the color encapsulated by the Computation struct of the world.
func (world *World) ShadeHit(comps *Computation, remaining int) *Color {
	light := Black
	material := comps.object.Material()
	reflectance := 1.0
	refractance := 1.0
	if material.reflective > 0 && material.transparency > 0 {
		reflectance = comps.Schlick()
		refractance = 1 - reflectance
	}
	for i := 0; i < len(world.lights); i++ {

		light = light.Add(
			Lighting(
				comps.object.Material(),
				comps.object,
				world.lights[i],
				comps.overPoint,
				comps.eyev,
				comps.normalv,
				world.IsShadowed(comps.overPoint, i)),
		).Add(
			world.ReflectedColor(comps, remaining).MultiplyByScalar(reflectance),
		).Add(
			world.RefractedColor(comps, remaining).MultiplyByScalar(refractance),
		)
	}
	return light
}

// ReflectedColor creates a new ray, originating at the hit’s location and pointing in the direction of reflectv.
func (world *World) ReflectedColor(comps *Computation, remaining int) *Color {
	if comps.object.Material().reflective == 0.0 || remaining < 1 {
		return Black
	}
	reflectRay := NewRay(comps.overPoint, comps.reflectv)
	color := world.ColorAt(reflectRay, remaining-1)

	return color.MultiplyByScalar(comps.object.Material().reflective)
}

// ColorAt will combine intersect(), prepare_computations() and shade_hit() functions and will
// intersect the world with the given ray and then return the color at the resulting intersection.
func (world *World) ColorAt(ray *Ray, remaining int) *Color {
	xs := world.Intersect(ray)
	hit := xs.Hit()
	if hit == nil {
		return Black
	}
	comps := PrepareComputations(hit, ray, xs)
	return world.ShadeHit(comps, remaining)
}

// RefractedColor calculates the resulting color from a simulated refraction mathematical model.
func (world *World) RefractedColor(comps *Computation, remaining int) *Color {
	if comps.object.Material().transparency == 0.0 || remaining < 1 {
		return Black
	}
	nRatio := comps.n1 / comps.n2

	cosI := comps.eyev.DotProduct(comps.normalv)

	sin2t := square(nRatio) * (1 - square(cosI))

	if sin2t > 1.0 {
		return Black
	}

	cosT := math.Sqrt(1.0 - sin2t)

	direction := comps.normalv.Multiply(nRatio*cosI - cosT).Substract(comps.eyev.Multiply(nRatio))

	refractRay := NewRay(comps.underPoint, direction)

	color := world.ColorAt(refractRay, remaining-1).MultiplyByScalar(comps.object.Material().transparency)

	return color
}

// Schlick returns the reflectance, represents what fraction of light is reflected given surface info and the Intersections.Hit().
// This is a faster approximation to Fresnel’s equation and with enough accuracy.
func (comps *Computation) Schlick() float64 {
	cos := comps.eyev.DotProduct(comps.normalv)
	if comps.n1 > comps.n2 {
		n := comps.n1 / comps.n2
		sin2T := square(n) * (1.0 - square(cos))
		if sin2T > 1.0 {
			return 1.0
		}
		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}
	ro := square((comps.n1 - comps.n2) / (comps.n1 + comps.n2))

	return ro + (1-ro)*math.Pow(1-cos, 5)
}

// IsShadowed returns whether a point is considered to be under a shadow.
func (world *World) IsShadowed(point *Tuple, light int) bool {
	v := world.lights[light].position.Substract(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	ray := NewRay(point, direction)

	intersections := world.Intersect(ray)

	hit := intersections.Hit()

	return hit != nil && hit.t < distance
}
