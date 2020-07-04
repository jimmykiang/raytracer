package main

import "math"

// Shape interface defining any object in the scene.
type Shape interface {
	SetMaterial(*Material)
	SetTransform(Matrix)
	Transform() Matrix
	Material() *Material
	Intersect(*Ray) []*Intersection
	NormalAt(*Tuple) *Tuple
}

// Sphere object
type Sphere struct {
	origin    *Tuple
	transform Matrix
	material  *Material
}

// NewSphere creates a new default sphere centered at the origin with Identity matrix as transform and default material.
func NewSphere() *Sphere {
	return &Sphere{Point(0, 0, 0), IdentityMatrix, DefaultMaterial()}
}

// Material returns the material of a Sphere.
func (sphere *Sphere) Material() *Material {
	return sphere.material
}

// SetTransform sets the spheres transformation.
func (sphere *Sphere) SetTransform(transformation Matrix) {
	sphere.transform = transformation.Inverse()
}

// SetMaterial sets the spheres material.
func (sphere *Sphere) SetMaterial(material *Material) {
	sphere.material = material
}

//Transform returns the transformation
func (sphere *Sphere) Transform() Matrix {
	return sphere.transform
}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (sphere *Sphere) NormalAt(point *Tuple) *Tuple {
	objectPoint := sphere.transform.MultiplyMatrixByTuple(point)
	objectNormal := objectPoint.Substract(sphere.origin)
	worldNormal := sphere.transform.Transpose().MultiplyMatrixByTuple(objectNormal)

	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

// Intersect computes the intersection between a sphere and a ray
func (sphere *Sphere) Intersect(ray *Ray) []*Intersection {
	ray = ray.Transform(sphere.transform)
	sphereToRay := ray.origin.Substract(sphere.origin)
	a := ray.direction.DotProduct(ray.direction)
	b := 2 * ray.direction.DotProduct(sphereToRay)
	c := sphereToRay.DotProduct(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return []*Intersection{}
	}
	sqrtDisc := math.Sqrt(discriminant)
	div := (2 * a)
	t1 := (-b - sqrtDisc) / div
	t2 := (-b + sqrtDisc) / div
	return []*Intersection{NewIntersection(t1, sphere), NewIntersection(t2, sphere)}

}
