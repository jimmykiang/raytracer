package main

import "testing"

func Test2DCheckerPattern(t *testing.T) {
	// Checker pattern in 2D.
	checkers := uvCheckers(2, 2, Black, White)

	type testStruct struct {
		u             float64
		v             float64
		expectedColor *Color
	}

	expectedTest := []testStruct{

		{u: 0.0, v: 0.0, expectedColor: Black},
		{u: 0.5, v: 0.0, expectedColor: White},
		{u: 0.0, v: 0.5, expectedColor: White},
		{u: 0.5, v: 0.5, expectedColor: Black},
		{u: 1.0, v: 1.0, expectedColor: Black},
	}

	for _, val := range expectedTest {

		if !(uvPatternAt(checkers, val.u, val.v) == val.expectedColor) {
			t.Errorf("Checker pattern in 2D, got: %v and expected to be %v",
				uvPatternAt(checkers, val.u, val.v), val.expectedColor)
		}
	}
}
