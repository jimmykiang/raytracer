package main

// Ray is a struct used for raycasting purposes.
// It contains the representation of a origin point and a direction vector.
type Ray struct {
	origin, direction *Tuple
}

// NewRay creates a new ray.
func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{origin, direction}
}

// Position calculates the point at the given distance t along the ray
func (ray *Ray) Position(t float64) *Tuple {
	return ray.origin.Add(ray.direction.Multiply(t))
}

// Transform will return a new ray with its origin and direction transformed.
func (ray *Ray) Transform(transformations ...Matrix) *Ray {
	return NewRay(
		ray.origin.Transform(transformations...),
		ray.direction.Transform(transformations...),
	)
}

// Equals checks ray equality
func (ray *Ray) Equals(other *Ray) bool {
	return ray.origin.Equals(other.origin) && ray.direction.Equals(other.direction)
}
