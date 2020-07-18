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

//Transform returns the transformation.
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

// Plane Shape
type Plane struct {
	transform Matrix
	material  *Material
}

// NewPlane creates a new default Plane centered at the origin with Identity matrix as transform and default material.
func NewPlane() *Plane {
	return &Plane{IdentityMatrix, DefaultMaterial()}

}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (plane *Plane) NormalAt(point *Tuple) *Tuple {
	localNormal := Vector(0, 1, 0)
	worldNormal := plane.transform.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()

}

// Intersect calculates the local intersections between a ray and a plane.
func (plane *Plane) Intersect(ray *Ray) []*Intersection {
	if abs(ray.direction.y) < EPSILON {
		return []*Intersection{}
	}
	ray = ray.Transform(plane.transform)
	t := -ray.origin.y / ray.direction.y

	return []*Intersection{NewIntersection(t, plane)}
}

// Transform returns the transformation.
func (plane *Plane) Transform() Matrix {
	return plane.transform
}

// Material returns the material of a Plane.
func (plane *Plane) Material() *Material {
	return plane.material
}

// SetTransform sets the Plane's transformation.
func (plane *Plane) SetTransform(transform Matrix) {
	plane.transform = transform.Inverse()
}

// SetMaterial returns the material of a Plane.
func (plane *Plane) SetMaterial(material *Material) {
	plane.material = material
}

// GlassSphere returns a sphere with transparency and refractiveIndex values simulating glass.
func GlassSphere() *Sphere {
	m := DefaultMaterial()
	m.transparency = 1.0
	m.refractiveIndex = 1.5
	return &Sphere{Point(0, 0, 0), IdentityMatrix, m}
}

// Cube struct.
type Cube struct {
	origin    *Tuple
	transform Matrix
	material  *Material
}

// NewCube creates a new default NewCube centered at the origin with Identity matrix as transform and default material.
func NewCube() *Cube {
	return &Cube{Point(0, 0, 0), IdentityMatrix, DefaultMaterial()}
}

// Intersect computes the local intersection between a cube and a ray.
func (cube *Cube) Intersect(ray *Ray) []*Intersection {

	xTMin, xTMax := checkAxis(ray.origin.x, ray.direction.x)
	yTMin, yTMax := checkAxis(ray.origin.y, ray.direction.y)
	zTMin, zTMax := checkAxis(ray.origin.z, ray.direction.z)

	tMin := max(xTMin, yTMin, zTMin)
	tMax := min(xTMax, yTMax, zTMax)

	if tMin > tMax {
		return nil
	}

	return []*Intersection{
		NewIntersection(tMin, cube),
		NewIntersection(tMax, cube)}
}

func checkAxis(origin float64, direction float64) (tMin float64, tMax float64) {

	tMinNumerator := -1 - origin
	tMaxNumerator := 1 - origin

	if abs(direction) >= EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}

	if tMin > tMax {
		tMin, tMax = tMax, tMin
	}
	return
}

// Material returns the material of a Cube.
func (cube *Cube) Material() *Material {
	return cube.material
}

// NormalAt calculates the local normal (vector perpendicular to the surface) at a given point of the object.
func (cube *Cube) NormalAt(point *Tuple) *Tuple {

	maxc := max(abs(point.x), abs(point.y), abs(point.z))

	if maxc == abs(point.x) {
		return Vector(point.x, 0, 0)
	} else if maxc == abs(point.y) {
		return Vector(0, point.y, 0)
	} else {
		return Vector(0, 0, point.z)
	}
}

// SetMaterial returns the material of a Cube.
func (cube *Cube) SetMaterial(material *Material) {
	cube.material = material
}

// SetTransform sets the Cube's transformation.
func (cube *Cube) SetTransform(transform Matrix) {
	cube.transform = transform.Inverse()
}

// Transform returns the transformation.
func (cube *Cube) Transform() Matrix {
	return cube.transform
}
