package main

import (
	"math"
	"math/rand"
)

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
	GetParent() Shape
	SetParent(shape Shape)
	getId() int
}

// Sphere object
type Sphere struct {
	origin           *Tuple
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	material         *Material
	parent           Shape
	id               int
}

// NewSphere creates a new default sphere centered at the origin with Identity matrix as transform and default material.
func NewSphere() *Sphere {
	return &Sphere{origin: Point(0, 0, 0),
		transform:        IdentityMatrix,
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		material:         DefaultMaterial(),
		id:               rand.Int(),
	}
}

// GlassSphere returns a sphere with transparency and refractiveIndex values simulating glass.
func GlassSphere() *Sphere {
	m := DefaultMaterial()
	m.transparency = 1.0
	m.refractiveIndex = 1.5
	return &Sphere{
		origin:    Point(0, 0, 0),
		transform: IdentityMatrix,
		inverse:   IdentityMatrix,
		material:  m,
	}
}

// getId returns the id of the shape.
func (sphere *Sphere) getId() int {
	return sphere.id
}

// GetParent returns the parent shape from this current shape.
func (sphere *Sphere) GetParent() Shape {
	return sphere.parent
}

// SetParent sets the parent shape from this current shape.
func (sphere *Sphere) SetParent(shape Shape) {
	sphere.parent = shape
}

// Material returns the material of a Sphere.
func (sphere *Sphere) Material() *Material {
	return sphere.material
}

// SetTransform sets the spheres transformation.
func (sphere *Sphere) SetTransform(transformation Matrix) {
	sphere.transform = sphere.transform.MultiplyMatrix(transformation)
	sphere.inverse = sphere.transform.Inverse()
}

// SetMaterial sets the spheres material.
func (sphere *Sphere) SetMaterial(material *Material) {
	sphere.material = material
}

// Transform returns the transformation.
func (sphere *Sphere) Transform() Matrix {
	return sphere.transform
}

func (sphere *Sphere) localNormalAt(localPoint *Tuple) (localNormal *Tuple) {

	localNormal = localPoint.Substract(sphere.origin)
	return
}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (sphere *Sphere) NormalAt(worldPoint *Tuple) *Tuple {
	localPoint := sphere.inverse.MultiplyMatrixByTuple(worldPoint)

	localNormal := sphere.localNormalAt(localPoint)

	worldNormal := sphere.inverse.Transpose().MultiplyMatrixByTuple(localNormal)

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
	localRay := worldRay.Transform(sphere.inverse)
	return sphere.localIntersect(localRay)
}

// Plane Shape
type Plane struct {
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	material         *Material
	parent           Shape
	id               int
}

// NewPlane creates a new default Plane centered at the origin with Identity matrix as transform and default material.
func NewPlane() *Plane {
	return &Plane{
		transform:        IdentityMatrix,
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		material:         DefaultMaterial(),
		id:               rand.Int(),
	}
}

// getId returns the id of the shape.
func (plane *Plane) getId() int {
	return plane.id
}

// GetParent returns the parent shape from this current shape.
func (plane *Plane) GetParent() Shape {
	return plane.parent
}

// SetParent sets the parent shape from this current shape.
func (plane *Plane) SetParent(shape Shape) {
	plane.parent = shape
}

func (plane *Plane) localNormalAt(localPoint *Tuple) (localNormal *Tuple) {

	localNormal = Vector(0, 1, 0)
	return
}

// NormalAt calculates the normal(vector perpendicular to the surface) at a given point.
func (plane *Plane) NormalAt(worldPoint *Tuple) *Tuple {
	localPoint := plane.inverse.MultiplyMatrixByTuple(worldPoint)
	localNormal := plane.localNormalAt(localPoint)
	worldNormal := plane.inverse.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()

}

func (plane *Plane) localIntersect(localRay *Ray) []*Intersection {

	if math.Abs(localRay.direction.y) < EPSILON {
		return []*Intersection{}
	}

	t := -localRay.origin.y / localRay.direction.y
	return []*Intersection{NewIntersection(t, plane)}
}

// Intersect calculates the local intersections between a ray and a plane.
func (plane *Plane) Intersect(worldRay *Ray) []*Intersection {

	localRay := worldRay.Transform(plane.inverse)
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
	plane.transform = plane.transform.MultiplyMatrix(transform)
	plane.inverse = plane.transform.Inverse()
}

// SetMaterial returns the material of a Plane.
func (plane *Plane) SetMaterial(material *Material) {
	plane.material = material
}

// Cube struct.
type Cube struct {
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	material         *Material
	parent           Shape
	id               int
}

// NewCube creates a new default NewCube centered at the origin with Identity matrix as transform and default material.
func NewCube() *Cube {
	return &Cube{
		transform:        IdentityMatrix,
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		material:         DefaultMaterial(),
		id:               rand.Int(),
	}
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

// getId returns the id of the shape.
func (cube *Cube) getId() int {
	return cube.id
}

// GetParent returns the parent shape from this current shape.
func (cube *Cube) GetParent() Shape {
	return cube.parent
}

// SetParent sets the parent shape from this current shape.
func (cube *Cube) SetParent(shape Shape) {
	cube.parent = shape
}

// Intersect computes the local intersection between a cube and a ray.
func (cube *Cube) Intersect(worldRay *Ray) []*Intersection {

	ray := worldRay.Transform(cube.inverse)
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
		tmin, tmax = tmax, tmin
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

	localPoint := cube.inverse.MultiplyMatrixByTuple(worldPoint)
	localNormal := cube.localNormalAt(localPoint)
	worldNormal := cube.inverse.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

// SetMaterial returns the material of a Cube.
func (cube *Cube) SetMaterial(material *Material) {
	cube.material = material
}

// SetTransform sets the Cube's transformation.
func (cube *Cube) SetTransform(transform Matrix) {
	cube.transform = cube.transform.MultiplyMatrix(transform)
	cube.inverse = cube.transform.Inverse()
}

// Transform returns the transformation.
func (cube *Cube) Transform() Matrix {
	return cube.transform
}

// Cylinder struct.
type Cylinder struct {
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	material         *Material
	minimum, maximum float64
	closed           bool
	parent           Shape
	id               int
}

// NewCylinder creates a new default Cylinder centered at the origin with Identity matrix as transform and default material.
func NewCylinder() *Cylinder {
	return &Cylinder{
		transform:        NewIdentityMatrix(),
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		material:         DefaultMaterial(),
		minimum:          math.Inf(-1),
		maximum:          math.Inf(1),
		id:               rand.Int(),
	}
}

// getId returns the id of the shape.
func (cylinder *Cylinder) getId() int {
	return cylinder.id
}

// GetParent returns the parent shape from this current shape.
func (cylinder *Cylinder) GetParent() Shape {
	return cylinder.parent
}

// SetParent sets the parent shape from this current shape.
func (cylinder *Cylinder) SetParent(shape Shape) {
	cylinder.parent = shape
}

// Intersect calculates the local intersections between a ray and a cylinder.
func (cylinder *Cylinder) Intersect(worldRay *Ray) []*Intersection {

	localRay := worldRay.Transform(cylinder.inverse)
	return cylinder.localIntersect(localRay)
}

func (cylinder *Cylinder) localIntersect(localRay *Ray) []*Intersection {

	a := math.Pow(localRay.direction.x, 2) + math.Pow(localRay.direction.z, 2)

	// localRay is parallel to the y axis.
	if math.Abs(a) < EPSILON {
		return cylinder.intersectCaps(localRay, Intersections{})
	}

	b := 2*localRay.origin.x*localRay.direction.x +
		2*localRay.origin.z*localRay.direction.z

	c := math.Pow(localRay.origin.x, 2) + math.Pow(localRay.origin.z, 2) - 1

	disc := b*b - 4*a*c

	// localRay does not intersect the cylinder.
	if disc < 0 {
		return []*Intersection{}
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	xs := Intersections{}

	y0 := localRay.origin.y + t0*localRay.direction.y

	if cylinder.minimum < y0 && y0 < cylinder.maximum {
		xs = append(xs, NewIntersection(t0, cylinder))
	}

	y1 := localRay.origin.y + t1*localRay.direction.y

	if cylinder.minimum < y1 && y1 < cylinder.maximum {
		xs = append(xs, NewIntersection(t1, cylinder))
	}

	return cylinder.intersectCaps(localRay, xs)
}

// Material returns the material of a Sphere.
func (cylinder *Cylinder) Material() *Material {
	return cylinder.material
}

// SetTransform sets the spheres transformation.
func (cylinder *Cylinder) SetTransform(transformation Matrix) {
	cylinder.transform = cylinder.transform.MultiplyMatrix(transformation)
	cylinder.inverse = cylinder.transform.Inverse()
}

// SetMaterial sets the spheres material.
func (cylinder *Cylinder) SetMaterial(material *Material) {
	cylinder.material = material
}

//Transform returns the transformation.
func (cylinder *Cylinder) Transform() Matrix {
	return cylinder.transform
}

func (cylinder *Cylinder) localNormalAt(localPoint *Tuple) *Tuple {

	// Compute the square of the distance from the y axis.
	dist := math.Pow(localPoint.x, 2) + math.Pow(localPoint.z, 2)

	if dist < 1 && localPoint.y >= cylinder.maximum-EPSILON {
		return Vector(0, 1, 0)

	} else if dist < 1 && localPoint.y <= cylinder.minimum+EPSILON {
		return Vector(0, -1, 0)

	} else {
		return Vector(localPoint.x, 0, localPoint.z)
	}
}

// NormalAt calculates the local normal (vector perpendicular to the surface) at a given point of the object.
func (cylinder *Cylinder) NormalAt(worldPoint *Tuple) *Tuple {

	localPoint := cylinder.inverse.MultiplyMatrixByTuple(worldPoint)
	localNormal := cylinder.localNormalAt(localPoint)
	worldNormal := cylinder.inverse.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

// Checks to see if the intersection at `t` is within a radius of 1 (the radius of your cylinders) from the y axis.
func (cylinder *Cylinder) checkCap(ray *Ray, t float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z
	return math.Pow(x, 2)+math.Pow(z, 2) <= 1.0
}

func (cylinder *Cylinder) intersectCaps(ray *Ray, xs Intersections) Intersections {

	// Caps only matter if the cylinder is closed, and might possibly be intersected by the ray.
	if !cylinder.closed || math.Abs(ray.direction.y) < EPSILON {
		return xs
	}

	// check for an intersection with the lower end cap by intersecting
	// the ray with the plane at y=cyl.minimum.
	t := (cylinder.minimum - ray.origin.y) / ray.direction.y
	if cylinder.checkCap(ray, t) {
		xs = append(xs, NewIntersection(t, cylinder))
	}

	// check for an intersection with the upper end cap by intersecting
	// the ray with the plane at y=cyl.maximum.
	t = (cylinder.maximum - ray.origin.y) / ray.direction.y
	if cylinder.checkCap(ray, t) {
		xs = append(xs, NewIntersection(t, cylinder))
	}
	return xs
}

// Cone struct.
type Cone struct {
	transform        Matrix
	inverse          Matrix
	inverseTranspose Matrix
	material         *Material
	minimum, maximum float64
	closed           bool
	parent           Shape
	id               int
}

// NewCone creates a new default Cone centered at the origin with Identity matrix as transform and default material.
func NewCone() *Cone {
	return &Cone{
		transform:        NewIdentityMatrix(),
		inverse:          IdentityMatrix,
		inverseTranspose: IdentityMatrix,
		material:         DefaultMaterial(),
		minimum:          math.Inf(-1),
		maximum:          math.Inf(1),
		id:               rand.Int(),
	}
}

// getId returns the id of the shape.
func (cone *Cone) getId() int {
	return cone.id
}

func (cone *Cone) GetParent() Shape {
	return cone.parent
}

// SetParent sets the parent shape from this current shape.
func (cone *Cone) SetParent(shape Shape) {
	cone.parent = shape
}

// Intersect calculates the local intersections between a ray and a Cone.
func (cone *Cone) Intersect(worldRay *Ray) []*Intersection {

	localRay := worldRay.Transform(cone.inverse)
	return cone.localIntersect(localRay)
}

func (cone *Cone) localIntersect(localRay *Ray) []*Intersection {

	xs := Intersections{}

	a := math.Pow(localRay.direction.x, 2) -
		math.Pow(localRay.direction.y, 2) +
		math.Pow(localRay.direction.z, 2)

	b := 2*localRay.origin.x*localRay.direction.x -
		2*localRay.origin.y*localRay.direction.y +
		2*localRay.origin.z*localRay.direction.z

	if math.Abs(a) < EPSILON && math.Abs(b) < EPSILON {

		return xs
	}

	c := math.Pow(localRay.origin.x, 2) -
		math.Pow(localRay.origin.y, 2) +
		math.Pow(localRay.origin.z, 2)

	disc := b*b - 4*a*c

	// localRay does not intersect the cone.
	if disc < 0 {
		return xs
	}

	if math.Abs(a) < EPSILON && math.Abs(b) > EPSILON {
		t0 := -c / (2.0 * b)
		xs = append(xs, NewIntersection(t0, cone))
	} else {

		t0 := (-b - math.Sqrt(disc)) / (2 * a)
		t1 := (-b + math.Sqrt(disc)) / (2 * a)

		y0 := localRay.origin.y + t0*localRay.direction.y

		if cone.minimum < y0 && y0 < cone.maximum {
			xs = append(xs, NewIntersection(t0, cone))
		}

		y1 := localRay.origin.y + t1*localRay.direction.y

		if cone.minimum < y1 && y1 < cone.maximum {
			xs = append(xs, NewIntersection(t1, cone))
		}
	}

	return cone.intersectCaps(localRay, xs)
}

// Material returns the material of a Sphere.
func (cone *Cone) Material() *Material {
	return cone.material
}

// SetTransform sets the shape's transformation.
func (cone *Cone) SetTransform(transformation Matrix) {
	cone.transform = cone.transform.MultiplyMatrix(transformation)
	cone.inverse = cone.transform.Inverse()
}

// SetMaterial sets the shape's material.
func (cone *Cone) SetMaterial(material *Material) {
	cone.material = material
}

//Transform returns the transformation.
func (cone *Cone) Transform() Matrix {
	return cone.transform
}

func (cone *Cone) localNormalAt(localPoint *Tuple) *Tuple {

	// Compute the square of the distance from the y axis.
	dist := math.Pow(localPoint.x, 2) + math.Pow(localPoint.z, 2)

	if dist < 1 && localPoint.y >= cone.maximum-EPSILON {
		return Vector(0, 1, 0)

	} else if dist < 1 && localPoint.y <= cone.minimum+EPSILON {
		return Vector(0, -1, 0)

	} else {
		y := math.Sqrt(math.Pow(localPoint.x, 2) + math.Pow(localPoint.z, 2))

		if localPoint.y > 0 {

			y = -y
		}
		return Vector(localPoint.x, y, localPoint.z)
	}
}

// NormalAt calculates the local normal (vector perpendicular to the surface) at a given point of the object.
func (cone *Cone) NormalAt(worldPoint *Tuple) *Tuple {

	localPoint := cone.inverse.MultiplyMatrixByTuple(worldPoint)
	localNormal := cone.localNormalAt(localPoint)
	worldNormal := cone.inverse.Transpose().MultiplyMatrixByTuple(localNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

func (cone *Cone) intersectCaps(localRay *Ray, xs Intersections) Intersections {

	// Caps only matter if the cone is closed, and might possibly be intersected by the ray.
	if !cone.closed || math.Abs(localRay.direction.y) < EPSILON {
		return xs
	}

	// check for an intersection with the lower end cap by intersecting
	// the ray with the plane at y=cyl.minimum.
	t := (cone.minimum - localRay.origin.y) / localRay.direction.y
	if cone.checkCap(localRay, t, cone.minimum) {
		xs = append(xs, NewIntersection(t, cone))
	}

	// check for an intersection with the upper end cap by intersecting
	// the ray with the plane at y=cyl.maximum.
	t = (cone.maximum - localRay.origin.y) / localRay.direction.y
	if cone.checkCap(localRay, t, cone.maximum) {
		xs = append(xs, NewIntersection(t, cone))
	}
	return xs
}

// checkCap for cone: the radius of a cone will change with y.
// In fact, a cone’s radius at any given y will be the absolute value of that y.
func (cone *Cone) checkCap(localRay *Ray, t float64, minMaxY float64) bool {
	x := localRay.origin.x + t*localRay.direction.x
	z := localRay.origin.z + t*localRay.direction.z
	return math.Pow(x, 2)+math.Pow(z, 2) <= math.Abs(minMaxY)
}
