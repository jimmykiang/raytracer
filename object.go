package main

import "math"

// Shape interface defining any object in the scene.
type Shape interface {
	SetMaterial(*Material)
	SetTransform(Matrix)
	Transform() Matrix
	Material() *Material
	Intersect(*Ray) []*Intersection
	localIntersect(*Ray) []*Intersection
	NormalAt(*Tuple) *Tuple
	localNormalAt(*Tuple) *Tuple
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

func (sphere *Sphere) localNormalAt(localPoint *Tuple) (localNormal *Tuple) {

	localNormal = localPoint.Substract(sphere.origin)
	return
}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (sphere *Sphere) NormalAt(worldPoint *Tuple) *Tuple {
	localPoint := sphere.transform.MultiplyMatrixByTuple(worldPoint)

	localNormal := sphere.localNormalAt(localPoint)

	worldNormal := sphere.transform.Transpose().MultiplyMatrixByTuple(localNormal)

	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

func (sphere *Sphere) localIntersect(localRay *Ray) []*Intersection {
	sphereToRay := localRay.origin.Substract(sphere.origin)
	a := localRay.direction.DotProduct(localRay.direction)
	b := 2 * localRay.direction.DotProduct(sphereToRay)
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

// Intersect computes the intersection between a sphere and a ray
func (sphere *Sphere) Intersect(worldRay *Ray) []*Intersection {
	localRay := worldRay.Transform(sphere.transform)
	return sphere.localIntersect(localRay)
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

func (plane *Plane) localNormalAt(localPoint *Tuple) (localNormal *Tuple) {

	localNormal = Vector(0, 1, 0)
	return
}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (plane *Plane) NormalAt(worldPoint *Tuple) *Tuple {

	localPoint := plane.transform.MultiplyMatrixByTuple(worldPoint)

	localNormal := plane.localNormalAt(localPoint)

	worldNormal := plane.transform.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()

}

func (plane *Plane) localIntersect(localRay *Ray) []*Intersection {

	t := -localRay.origin.y / localRay.direction.y
	return []*Intersection{NewIntersection(t, plane)}
}

// Intersect calculates the local intersections between a ray and a plane.
func (plane *Plane) Intersect(worldRay *Ray) []*Intersection {
	if math.Abs(worldRay.direction.y) < EPSILON {
		return []*Intersection{}
	}
	localRay := worldRay.Transform(plane.transform)
	return plane.localIntersect(localRay)
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
	transform Matrix
	material  *Material
}

// NewCube creates a new default NewCube centered at the origin with Identity matrix as transform and default material.
func NewCube() *Cube {
	return &Cube{NewIdentityMatrix(), DefaultMaterial()}
}

func (cube *Cube) localIntersect(localRay *Ray) []*Intersection {

	xTMin, xTMax := checkAxis(localRay.origin.x, localRay.direction.x)
	yTMin, yTMax := checkAxis(localRay.origin.y, localRay.direction.y)
	zTMin, zTMax := checkAxis(localRay.origin.z, localRay.direction.z)

	tMin := max(xTMin, yTMin, zTMin)
	tMax := min(xTMax, yTMax, zTMax)

	if tMin > tMax {
		return []*Intersection{}
	}

	return []*Intersection{
		NewIntersection(tMin, cube),
		NewIntersection(tMax, cube)}
}

// Intersect computes the local intersection between a cube and a ray.
func (cube *Cube) Intersect(worldRay *Ray) []*Intersection {

	ray := worldRay.Transform(cube.transform)
	return cube.localIntersect(ray)

}

func checkAxis(origin float64, direction float64) (min float64, max float64) {
	tminNumerator := -1 - origin
	tmaxNumerator := 1 - origin
	var tmin, tmax float64
	if math.Abs(direction) >= EPSILON {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}
	if tmin > tmax {
		// swap
		temp := tmin
		tmin = tmax
		tmax = temp
	}
	return tmin, tmax
}

// Material returns the material of a Cube.
func (cube *Cube) Material() *Material {
	return cube.material
}

func (cube *Cube) localNormalAt(localPoint *Tuple) (localNormal *Tuple) {

	maxc := max(math.Abs(localPoint.x), math.Abs(localPoint.y), math.Abs(localPoint.z))

	if maxc == math.Abs(localPoint.x) {

		localNormal = Vector(localPoint.x, 0, 0)
	} else if maxc == math.Abs(localPoint.y) {

		localNormal = Vector(0, localPoint.y, 0)
	} else {

		localNormal = Vector(0, 0, localPoint.z)
	}
	return
}

// NormalAt calculates the local normal (vector perpendicular to the surface) at a given point of the object.
func (cube *Cube) NormalAt(worldPoint *Tuple) *Tuple {

	localPoint := cube.transform.MultiplyMatrixByTuple(worldPoint)

	localNormal := cube.localNormalAt(localPoint)
	worldNormal := cube.transform.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0
	return worldNormal.Normalize()
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
