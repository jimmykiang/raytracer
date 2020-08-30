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

func TestPlannarMappingOn3DPoint(t *testing.T) {
	// Using a planar mapping on a 3D point.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(0.25, 0, 0.5), expectedU: 0.25, expectedV: 0.5},
		{point: Point(0.25, 0, -0.25), expectedU: 0.25, expectedV: 0.75},
		{point: Point(0.25, 0.5, -0.25), expectedU: 0.25, expectedV: 0.75},
		{point: Point(1.25, 0, 0.5), expectedU: 0.25, expectedV: 0.5},
		{point: Point(0.25, 0, -1.75), expectedU: 0.25, expectedV: 0.25},
		{point: Point(1, 0, -1), expectedU: 0, expectedV: 0},
		{point: Point(0, 0, 0), expectedU: 0, expectedV: 0},
	}

	for _, val := range expectedTest {

		u, v := planarMap(val.point)
		if !(u == val.expectedU) {
			t.Errorf("Using a planar mapping on a 3D point, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("Using a planar mapping on a 3D point, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestTextureMapPatterWithCylindricalMap(t *testing.T) {
	// Using a cylindrical mapping on a 3D point.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(0, 0, -1), expectedU: 0.0, expectedV: 0.0},
		{point: Point(0, 0.5, -1), expectedU: 0.0, expectedV: 0.5},
		{point: Point(0, 1, -1), expectedU: 0.0, expectedV: 0.0},
		{point: Point(0.70711, 0.5, -0.70711), expectedU: 0.125, expectedV: 0.5},
		{point: Point(1, 0.5, 0), expectedU: 0.25, expectedV: 0.5},
		{point: Point(0.70711, 0.5, 0.70711), expectedU: 0.375, expectedV: 0.5},
		{point: Point(0, -0.25, 1), expectedU: 0.5, expectedV: 0.75},
		{point: Point(-0.70711, 0.5, 0.70711), expectedU: 0.625, expectedV: 0.5},
		{point: Point(-1, 1.25, 0), expectedU: 0.75, expectedV: 0.25},
		{point: Point(-0.70711, 0.5, -0.70711), expectedU: 0.875, expectedV: 0.5},
	}

	for _, val := range expectedTest {

		u, v := cylindricalMap(val.point)
		if !(u == val.expectedU) {
			t.Errorf("Using a cylindrical mapping on a 3D point, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("Using a cylindrical mapping on a 3D point, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestAlignCheckPattern(t *testing.T) {
	// Layout of the "align check" pattern.

	main := NewColor(1, 1, 1)
	ul := NewColor(1, 0, 0)
	ur := NewColor(1, 1, 0)
	bl := NewColor(0, 1, 0)
	br := NewColor(0, 1, 1)

	pattern := uvAlignCheck(main, ul, ur, bl, br)

	type testStruct struct {
		expectedU float64
		expectedV float64
		color     *Color
	}

	expectedTest := []testStruct{
		{color: main, expectedU: 0.5, expectedV: 0.5},
		{color: ul, expectedU: 0.1, expectedV: 0.9},
		{color: ur, expectedU: 0.9, expectedV: 0.9},
		{color: bl, expectedU: 0.1, expectedV: 0.1},
		{color: br, expectedU: 0.9, expectedV: 0.1},
	}

	for _, val := range expectedTest {

		c := uvPatternAt(pattern, val.expectedU, val.expectedV)
		if !(c == val.color) {
			t.Errorf("Layout of the align check pattern, got: %v and expected to be %v",
				c, val.color)
		}
	}
}

func TestIdentifyFaceOfCubeFromPoint(t *testing.T) {
	// Identifying the face of a cube from a point.

	type testStruct struct {
		point *Tuple
		face  string
	}

	expectedTest := []testStruct{
		{point: Point(-1, 0.5, -0.25), face: "left"},
		{point: Point(1.1, -0.75, 0.8), face: "right"},
		{point: Point(0.1, 0.6, 0.9), face: "front"},
		{point: Point(-0.7, 0, -2), face: "back"},
		{point: Point(0.5, 1, 0.9), face: "up"},
		{point: Point(-0.2, -1.3, 1.1), face: "down"},
	}

	for _, val := range expectedTest {

		face := faceFromPoint(val.point)
		if !(face == val.face) {
			t.Errorf("Identifying the face of a cube from a point, got: %v and expected to be %v",
				face, val.face)
		}
	}
}

func TestUVMappingFrontFaceCube(t *testing.T) {
	// UV mapping the front face of a cube.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(-0.5, 0.5, 1), expectedU: 0.25, expectedV: 0.75},
		{point: Point(0.5, -0.5, 1), expectedU: 0.75, expectedV: 0.25},
	}

	for _, val := range expectedTest {

		u, v := cubeUVFront(val.point)
		if !(u == val.expectedU) {
			t.Errorf("UV mapping the front face of a cube, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("UV mapping the front face of a cube, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestUVMappingBackFaceCube(t *testing.T) {
	// UV mapping the back face of a cube.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(0.5, 0.5, -1), expectedU: 0.25, expectedV: 0.75},
		{point: Point(-0.5, -0.5, -1), expectedU: 0.75, expectedV: 0.25},
	}

	for _, val := range expectedTest {

		u, v := cubeUVBack(val.point)
		if !(u == val.expectedU) {
			t.Errorf("UV mapping the back face of a cube, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("UV mapping the back face of a cube, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestUVMappingLeftFaceCube(t *testing.T) {
	// UV mapping the left face of a cube.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(-1, 0.5, -0.5), expectedU: 0.25, expectedV: 0.75},
		{point: Point(-1, -0.5, 0.5), expectedU: 0.75, expectedV: 0.25},
	}

	for _, val := range expectedTest {

		u, v := cubeUVLeft(val.point)
		if !(u == val.expectedU) {
			t.Errorf("UV mapping the back face of a cube, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("UV mapping the back face of a cube, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestUVMappingRightFaceCube(t *testing.T) {
	// UV mapping the right face of a cube.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(1, 0.5, 0.5), expectedU: 0.25, expectedV: 0.75},
		{point: Point(1, -0.5, -0.5), expectedU: 0.75, expectedV: 0.25},
	}

	for _, val := range expectedTest {

		u, v := cubeUVRight(val.point)
		if !(u == val.expectedU) {
			t.Errorf("UV mapping the right face of a cube, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("UV mapping the right face of a cube, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}

func TestUVMappingUpperFaceCube(t *testing.T) {
	// UV mapping the upper face of a cube.

	type testStruct struct {
		expectedU float64
		expectedV float64
		point     *Tuple
	}

	expectedTest := []testStruct{
		{point: Point(-0.5, 1, -0.5), expectedU: 0.25, expectedV: 0.75},
		{point: Point(0.5, 1, 0.5), expectedU: 0.75, expectedV: 0.25},
	}

	for _, val := range expectedTest {

		u, v := cubeUVUp(val.point)
		if !(u == val.expectedU) {
			t.Errorf("UV mapping the upper face of a cube, got: %v and expected to be %v",
				u, val.expectedU)
		}
		if !(v == val.expectedV) {
			t.Errorf("UV mapping the upper face of a cube, got: %v and expected to be %v",
				v, val.expectedV)
		}
	}
}
