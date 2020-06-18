package main

// Matrix is a new type defined by a double slice of float64.
type Matrix [][]float64

// IdentityMatrix holds a copy if the Identity Matrix.
var IdentityMatrix Matrix

func init() {

	IdentityMatrix = NewIdentityMatrix()
}

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

// Get returns a specific value from a matrix.
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

// MultiplyMatrix returns the multiplication of two 4x4 matrices.
func (matrix Matrix) MultiplyMatrix(matrix2 Matrix) Matrix {
	resultMatrix := NewMatrix(4, 4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			product := dotProducOfMatricesRowColumn(matrix.Row(row), matrix2.Column(col))
			resultMatrix.Set(row, col, product)
		}
	}

	return resultMatrix

}

// Row returns the slice from the elements of the entire row from the current matrix.
func (matrix Matrix) Row(r int) []float64 {
	return matrix[r]
}

// Column returns the slice from the elements of the entire column from the current the matrix.
func (matrix Matrix) Column(c int) []float64 {
	h, _ := matrix.Size()
	col := make([]float64, h, h)
	for i, row := range matrix {
		col[i] = row[c]
	}
	return col
}

// Size returns the height and width of the current matrix.
func (matrix Matrix) Size() (int, int) {
	height := len(matrix)
	width := 0
	if height > 0 {
		width = len(matrix[0])
	}
	return height, width
}

// MultiplyMatrixByTuple returns the multiplication of a Matrix by a Tuple.
func (matrix Matrix) MultiplyMatrixByTuple(tuple *Tuple) *Tuple {
	tupleAsMatrix := []float64{tuple.x, tuple.y, tuple.z, tuple.w}
	newTup := &Tuple{
		dotProducOfMatricesRowColumn(matrix.Row(0), tupleAsMatrix),
		dotProducOfMatricesRowColumn(matrix.Row(1), tupleAsMatrix),
		dotProducOfMatricesRowColumn(matrix.Row(2), tupleAsMatrix),
		dotProducOfMatricesRowColumn(matrix.Row(3), tupleAsMatrix),
	}

	return newTup
}

// IdentityMatrix returns a copy of that matrix.
func NewIdentityMatrix() Matrix {
	return Matrix(
		[][]float64{
			[]float64{1, 0, 0, 0},
			[]float64{0, 1, 0, 0},
			[]float64{0, 0, 1, 0},
			[]float64{0, 0, 0, 1},
		},
	)
}

// Transpose returns a copy of the transposed matrix.
func (matrix Matrix) Transpose() Matrix {
	height, width := matrix.Size()
	transposedMatrix := NewMatrix(width, height)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			transposedMatrix.Set(i, j, matrix.Get(j, i))
		}
	}
	return transposedMatrix
}

// Determinant calculates the determinant of a matrix.
func (matrix Matrix) Determinant() float64 {
	h, w := matrix.Size()
	if h == 2 && w == 2 {
		return matrix.Get(0, 0)*matrix.Get(1, 1) - matrix.Get(0, 1)*matrix.Get(1, 0)
	} else if h > 2 && w > 2 {
		det := 0.0

		for col := 0; col < w; col++ {
			det += matrix.Get(0, col) * matrix.Cofactor(0, col)
		}
		return det
	}
	return 0.0
}

// Cofactor caculates the cofactor of a submatrix.
func (matrix Matrix) Cofactor(row, col int) float64 {
	minor := matrix.Minor(row, col)

	// bitwise "AND" checking for odd number.
	if ((row + col) & 1) == 1 {
		return -minor
	}
	return minor
}

// Minor returns the determinant of a submatrix of a 3x3 matrix
func (matrix Matrix) Minor(row, col int) float64 {
	return matrix.SubMatrix(row, col).Determinant()
}

// SubMatrix returns a matrix with the given row and column removed.
func (matrix Matrix) SubMatrix(row, col int) Matrix {
	height, width := matrix.Size()
	subMatrix := NewMatrix(height-1, width-1)

	i := 0
	j := 0

	for r := 0; r < height; r++ {
		if r == row {
			continue
		}
		j = 0
		for c := 0; c < width; c++ {
			if c == col {
				continue
			}
			subMatrix.Set(i, j, matrix.Get(r, c))
			j++

		}
		i++
	}
	return subMatrix
}
