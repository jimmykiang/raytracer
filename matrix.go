package main

// Matrix is a new type defined by a double slice of float64.
type Matrix [][]float64

// NewMatrix creates a rows x cols matrix
func NewMatrix(rows, columns int) Matrix {
	matrix := make([][]float64, rows, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, columns, columns)
	}
	return matrix
}

// Set a specific value in a matrix.
func (matrix Matrix) Set(row, column int, val float64) float64 {
	matrix[row][column] = val
	return val
}

// Get returns the values of a matrix.
func (matrix Matrix) Get(row, column int) float64 {
	return matrix[row][column]
}

// Equals will compare each value between 2 matrices.
func (matrix Matrix) Equals(other Matrix) bool {

	if len(matrix) != len(other) {
		return false
	}
	if len(matrix[0]) != len(other[0]) {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !floatEqual(matrix[i][j], other[i][j]) {
				return false
			}
		}
	}
	return true
}
