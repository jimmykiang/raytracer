package main

import "math"

// Translation Returns a translation matrix
func Translation(x, y, z float64) Matrix {
	matrix := NewIdentityMatrix()

	matrix.Set(0, 3, x)
	matrix.Set(1, 3, y)
	matrix.Set(2, 3, z)

	return matrix
}

// Scaling returns a scale matrix
func Scaling(x, y, z float64) Matrix {
	matrix := NewIdentityMatrix()
	matrix.Set(0, 0, x)
	matrix.Set(1, 1, y)
	matrix.Set(2, 2, z)
	return matrix
}

// RotationX returns a rotation matrix of the given radians
func RotationX(r float64) Matrix {
	matrix := NewIdentityMatrix()
	matrix.Set(1, 1, math.Cos(r))
	matrix.Set(1, 2, -math.Sin(r))
	matrix.Set(2, 1, math.Sin(r))
	matrix.Set(2, 2, math.Cos(r))
	return matrix
}

// RotationY returns a rotation matrix of the given radians
func RotationY(r float64) Matrix {
	matrix := NewIdentityMatrix()
	matrix.Set(0, 0, math.Cos(r))
	matrix.Set(0, 2, math.Sin(r))
	matrix.Set(2, 0, -math.Sin(r))
	matrix.Set(2, 2, math.Cos(r))
	return matrix
}

// RotationZ returns a rotation matrix of the given radians
func RotationZ(r float64) Matrix {
	matrix := NewIdentityMatrix()
	matrix.Set(0, 0, math.Cos(r))
	matrix.Set(0, 1, -math.Sin(r))
	matrix.Set(1, 0, math.Sin(r))
	matrix.Set(1, 1, math.Cos(r))
	return matrix
}

// Shearing returns a shearing(for skewing) matrix
func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	matrix := NewIdentityMatrix()

	matrix.Set(0, 1, xy)
	matrix.Set(0, 2, xz)
	matrix.Set(1, 0, yx)
	matrix.Set(1, 2, yz)
	matrix.Set(2, 0, zx)
	matrix.Set(2, 1, zy)

	return matrix
}
