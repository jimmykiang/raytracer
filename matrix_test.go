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

func TestMatrixMulMatrix(t *testing.T) {
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

func TestMatrixMulTuple(t *testing.T) {
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
