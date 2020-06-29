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
