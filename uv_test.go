package main

import (
	"math"
	"testing"
)

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

func TestSphericalMappingOn3DPoint(t *testing.T) {
	// Using a spherical mapping on a 3D point.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(0, 0, -1), expectedU: 0.0, expectedV: 0.5},
		{point: Point(1, 0, 0), expectedU: 0.25, expectedV: 0.5},
		{point: Point(0, 0, 1), expectedU: 0.5, expectedV: 0.5},
		{point: Point(-1, 0, 0), expectedU: 0.75, expectedV: 0.5},
		{point: Point(0, 1, 0), expectedU: 0.5, expectedV: 1},
		{point: Point(0, -1, 0), expectedU: 0.5, expectedV: 0},
		{point: Point((math.Sqrt(2) / 2), (math.Sqrt(2) / 2), 0), expectedU: 0.25, expectedV: 0.75},
	}

	for _, val := range expectedTest {

		u, v := sphericalMap(val.point)

		if !floatEqual(val.expectedU, u) {
			t.Errorf("Using a spherical mapping on a 3D point, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !floatEqual(val.expectedV, v) {
			t.Errorf("Using a spherical mapping on a 3D point, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}
