package main

import (
	"testing"
)

func TestConstrucAndInspect4x4Matrix(t *testing.T) {

	// NewMatrix will initialize all values by default with 0.
	matrix := NewMatrix(4, 4)
	var value float64
	expectedNewValue := 0.0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value = matrix.Get(i, j)
			pass := value == expectedNewValue
			if !pass {
				t.Errorf("NewMatrix: %f should be %f", value, expectedNewValue)
			}
		}
	}

	matrix.Set(0, 0, 1)
	matrix.Set(0, 1, 2)
	matrix.Set(0, 2, 3)
	matrix.Set(0, 3, 4)
	matrix.Set(1, 0, 5.5)
	matrix.Set(1, 1, 6.5)
	matrix.Set(1, 2, 7.5)
	matrix.Set(1, 3, 8.5)
	matrix.Set(2, 0, 9)
	matrix.Set(2, 1, 10)
	matrix.Set(2, 2, 11)
	matrix.Set(2, 3, 12)
	matrix.Set(3, 0, 13.5)
	matrix.Set(3, 1, 14.5)
	matrix.Set(3, 2, 15.5)
	matrix.Set(3, 3, 16.5)

	matrix.Set(0, 0, 1)
	matrix.Set(0, 3, 4)
	matrix.Set(1, 0, 5.5)
	matrix.Set(1, 2, 7.5)
	matrix.Set(2, 2, 11)
	matrix.Set(3, 0, 13.5)
	matrix.Set(3, 2, 15.5)

	passSingleValues :=
		floatEqual(matrix.Get(0, 0), 1) &&
			floatEqual(matrix.Get(0, 3), 4) &&
			floatEqual(matrix.Get(1, 0), 5.5) &&
			floatEqual(matrix.Get(1, 2), 7.5) &&
			floatEqual(matrix.Get(2, 2), 11) &&
			floatEqual(matrix.Get(3, 0), 13.5) &&
			floatEqual(matrix.Get(3, 2), 15.5)

	expectedMatrix := Matrix([][]float64{
		[]float64{1, 2, 3, 4},
		[]float64{5.5, 6.5, 7.5, 8.5},
		[]float64{9, 10, 11, 12},
		[]float64{13.5, 14.5, 15.5, 16.5},
	},
	)

	passMatrix := matrix.Equals(expectedMatrix)

	if !(passMatrix && passSingleValues) {
		t.Errorf("Matrix: %v to be %v", matrix, expectedMatrix)
	}
}

func Test2x2Matrix(t *testing.T) {

	matrix := Matrix(
		[][]float64{
			[]float64{-3, 5},
			[]float64{1, 2},
		},
	)

	passSingleValues :=
		floatEqual(matrix.Get(0, 0), -3) &&
			floatEqual(matrix.Get(0, 1), 5) &&
			floatEqual(matrix.Get(1, 0), 1) &&
			floatEqual(matrix.Get(1, 1), 2)

	if !(passSingleValues) {
		t.Errorf("Problem in matrix: %v", matrix)
	}
}

func TestMatrixEquality(t *testing.T) {

	matrix1 := Matrix([][]float64{
		[]float64{1, 2, 3, 4},
		[]float64{5, 6, 7, 8},
		[]float64{9, 8, 7, 6},
		[]float64{5, 4, 3, 2},
	},
	)

	matrix2 := Matrix([][]float64{
		[]float64{1, 2, 3, 4},
		[]float64{5, 6, 7, 8},
		[]float64{9, 8, 7, 6},
		[]float64{5, 4, 3, 2},
	},
	)

	pass := matrix1.Equals(matrix2)

	if !pass {
		t.Errorf("Matrix1: %v not equals to matrix2: %v", matrix1, matrix2)
	}
}

func TestMatrixInequality(t *testing.T) {

	matrix1 := Matrix([][]float64{
		[]float64{1, 2, 3, 4},
		[]float64{5, 6, 7, 8},
		[]float64{9, 8, 7, 6},
		[]float64{5, 4, 3, 2},
	},
	)

	matrix2 := Matrix([][]float64{
		[]float64{2, 3, 4, 5},
		[]float64{6, 7, 8, 9},
		[]float64{8, 7, 6, 5},
		[]float64{4, 3, 2, 1},
	},
	)

	pass := !(matrix1.Equals(matrix2))

	if !pass {
		t.Errorf("Matrix1: %v not equals to matrix2: %v", matrix1, matrix2)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := Matrix(
		[][]float64{
			[]float64{1, 2, 3, 4},
			[]float64{5, 6, 7, 8},
			[]float64{9, 8, 7, 6},
			[]float64{5, 4, 3, 2},
		},
	)
	m2 := Matrix(
		[][]float64{
			[]float64{-2, 1, 2, 3},
			[]float64{3, 2, 1, -1},
			[]float64{4, 3, 6, 5},
			[]float64{1, 2, 7, 8},
		},
	)

	expected := Matrix(
		[][]float64{
			[]float64{20, 22, 50, 48},
			[]float64{44, 54, 114, 108},
			[]float64{40, 58, 110, 102},
			[]float64{16, 26, 46, 42},
		},
	)

	result := m1.MultiplyMatrix(m2)
	pass := result.Equals(expected)

	if !pass {
		t.Errorf("MulMatrix: expected %v to be %v", result, expected)
	}

}

func TestMatrixMultiplyByTuple(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{1, 2, 3, 4},
			[]float64{2, 4, 4, 2},
			[]float64{8, 6, 4, 1},
			[]float64{0, 0, 0, 1},
		},
	)
	tuple := &Tuple{1, 2, 3, 1}

	expected := &Tuple{18, 24, 33, 1}

	result := m.MultiplyMatrixByTuple(tuple)

	if !result.Equals(expected) {
		t.Errorf("MatrixMulTuple: expected %v to be %v", result, expected)
	}
}

func TestMultiplyMatrixByIdentityMatrix(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{0, 1, 2, 4},
			[]float64{1, 2, 4, 8},
			[]float64{2, 4, 8, 16},
			[]float64{4, 8, 16, 32},
		},
	)

	if !m.Equals(m.MultiplyMatrix(IdentityMatrix)) {
		t.Errorf("IdentityMatrix invalid.")
	}

}

func TestMultiplyIdentityMatrixByTuple(t *testing.T) {
	tuple := &Tuple{1, 2, 3, 1}

	if !tuple.Equals(IdentityMatrix.MultiplyMatrixByTuple(tuple)) {
		t.Errorf("IdentityMatrix invalid.")
	}
}

func TestTransposeMatrix(t *testing.T) {
	if !IdentityMatrix.Transpose().Equals(IdentityMatrix) {
		t.Errorf("MatrixTranspose on IdentityMatrix failed")
	}

	m := Matrix(
		[][]float64{
			[]float64{0, 9, 3, 0},
			[]float64{9, 8, 0, 8},
			[]float64{1, 8, 5, 3},
			[]float64{0, 0, 5, 8},
		},
	)
	expected := Matrix(
		[][]float64{
			[]float64{0, 9, 1, 0},
			[]float64{9, 8, 8, 0},
			[]float64{3, 0, 5, 5},
			[]float64{0, 8, 3, 8},
		},
	)
	result := m.Transpose()

	if !result.Equals(expected) {
		t.Errorf("MatrixTranspose: expected %v to equal %v", result, expected)

	}
}

func TestMatrixDeterminant(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{1, 5},
			[]float64{-3, 2},
		},
	)

	result := m.Determinant()
	expected := 17.0

	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}

	m = Matrix(
		[][]float64{
			[]float64{1, 2, 6},
			[]float64{-5, 8, -4},
			[]float64{2, 6, 4},
		},
	)
	result = m.Determinant()
	expected = -196.0

	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}
	m = Matrix(
		[][]float64{
			[]float64{-2, -8, 3, 5},
			[]float64{-3, 1, 7, 3},
			[]float64{1, 2, -9, 6},
			[]float64{-6, 7, 7, -9},
		},
	)
	result = m.Determinant()
	expected = -4071.0
	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}
}

func TestSubMatrix(t *testing.T) {
	m1 := Matrix(
		[][]float64{
			[]float64{1, 5, 0},
			[]float64{-3, 2, 7},
			[]float64{0, 6, -3},
		},
	)
	m1ResultSub := m1.SubMatrix(0, 2)
	m1ExpectedSub := Matrix(
		[][]float64{
			[]float64{-3, 2},
			[]float64{0, 6},
		},
	)

	if !m1ResultSub.Equals(m1ExpectedSub) {
		t.Errorf("MatrixSubMatrix: expected %v to equal %v", m1ResultSub, m1ExpectedSub)
	}

	m2 := Matrix(
		[][]float64{
			[]float64{-6, 1, 1, 6},
			[]float64{-8, 5, 8, 6},
			[]float64{-1, 0, 8, 2},
			[]float64{-7, 1, -1, 1},
		},
	)
	m2ResultSub := m2.SubMatrix(2, 1)
	m2ExpectedSub := Matrix(
		[][]float64{
			[]float64{-6, 1, 6},
			[]float64{-8, 8, 6},
			[]float64{-7, -1, 1},
		},
	)
	if !m2ExpectedSub.Equals(m2ResultSub) {
		t.Errorf("MatrixSubMatrix: expected %v to equal %v", m2ResultSub, m2ExpectedSub)
	}

}

func TestMatrixCofactor(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{3, 5, 0},
			[]float64{2, -1, -7},
			[]float64{6, -1, 5},
		},
	)
	minor1 := m.Minor(0, 0)
	cofactor1 := m.Cofactor(0, 0)
	minor2 := m.Minor(1, 0)
	cofactor2 := m.Cofactor(1, 0)

	if !floatEqual(minor1, -12) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", minor1, -12.0)
	}
	if !floatEqual(cofactor1, -12) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", cofactor1, -12.0)
	}
	if !floatEqual(minor2, 25) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", minor2, 25.0)
	}
	if !floatEqual(cofactor2, -25) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", cofactor2, -25.0)
	}
}

func TestMatrixInverse(t *testing.T) {
	matrix1 := Matrix(
		[][]float64{
			[]float64{8, -5, 9, 2},
			[]float64{7, 5, 6, 1},
			[]float64{-6, 0, 9, 6},
			[]float64{-3, 0, -9, -4},
		},
	)
	expected1 := Matrix(
		[][]float64{
			[]float64{-0.15385, -0.15385, -0.28205, -0.53846},
			[]float64{-0.07692, 0.12308, 0.025641, 0.03077},
			[]float64{0.35897, 0.35897, 0.43590, 0.92308},
			[]float64{-0.69230, -0.69231, -0.76923, -1.92308},
		},
	)

	result1 := matrix1.Inverse()
	if !result1.Equals(expected1) {

		t.Errorf("MatrixInverse: result %v does not equal %v", result1, expected1)

	}
}

func TestMultiplyMatrixProductByInverse(t *testing.T) {
	a := Matrix(
		[][]float64{
			[]float64{3, -9, 7, 3},
			[]float64{3, -8, 2, -9},
			[]float64{-4, 4, 4, 1},
			[]float64{-6, 5, -1, 1},
		},
	)
	b := Matrix(
		[][]float64{
			[]float64{8, 2, 2, 2},
			[]float64{3, -1, 7, 0},
			[]float64{7, 0, 5, 4},
			[]float64{6, -2, 0, 5},
		},
	)

	c := a.MultiplyMatrix(b)
	expected := c.MultiplyMatrix(b.Inverse())

	if !a.Equals(expected) {
		t.Errorf("TestMultiplyMatrixProductByInverse: result %v does not equal %v", a, expected)
	}
}
