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

func TestTextureMapPatterWithSphericalMap(t *testing.T) {
	// Using a texture map pattern with a spherical map.

	checkers := uvCheckers(16, 8, Black, White)
	pattern := textureMap(checkers, sphericalMap)

	type testStruct struct {
		expectedColor *Color
		point         *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(0.4315, 0.4670, 0.7719), expectedColor: White},
		{point: Point(-0.9654, 0.2552, -0.0534), expectedColor: Black},
		{point: Point(0.1039, 0.7090, 0.6975), expectedColor: White},
		{point: Point(-0.4986, -0.7856, -0.3663), expectedColor: Black},
		{point: Point(-0.0317, -0.9395, 0.3411), expectedColor: Black},
		{point: Point(0.4809, -0.7721, 0.4154), expectedColor: Black},
		{point: Point(0.0285, -0.9612, -0.2745), expectedColor: Black},
		{point: Point(-0.5734, -0.2162, -0.7903), expectedColor: White},
		{point: Point(0.7688, -0.1470, 0.6223), expectedColor: Black},
		{point: Point(-0.7652, 0.2175, 0.6060), expectedColor: Black},
	}

	for _, val := range expectedTest {

		if !(patternAt(pattern, val.point) == val.expectedColor) {
			t.Errorf("Using a texture map pattern with a spherical map, got: %v and expected to be %v",
				patternAt(pattern, val.point), val.expectedColor)
		}
	}
}
