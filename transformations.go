package main

// Translation Returns a translation matrix
func Translation(x, y, z float64) Matrix {
	matrix := NewIdentityMatrix()

	matrix.Set(0, 3, x)
	matrix.Set(1, 3, y)
	matrix.Set(2, 3, z)

	return matrix
}
