package main

import (
	"reflect"
	"testing"
)

func TestParseGibberish(t *testing.T) {
	// Ignoring unrecognized lines.

	gibberish := `There was a young lady named Bright
who traveled much faster than light.
She set out one day
in a relative way,
and came back the previous night.`
	parser := parseObjData(gibberish)
	if !(parser.ignoredLines == 5) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", parser.ignoredLines, 5)
	}
}

func TestParseVerticies(t *testing.T) {
	// Vertex records.

	data := `
v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
`
	expectedVertices := []*Tuple{
		Point(-1, 1, 0),
		Point(-1, 0.5, 0),
		Point(1, 0, 0),
		Point(1, 1, 0),
	}

	parser := parseObjData(data)
	// parser will refer to the vertices by their index, starting with 1.
	for i := range expectedVertices {
		if !(parser.vertices[i+1].Equals(expectedVertices[i])) {
			t.Errorf("Vertex records, got: %v and expected to be %v", parser.vertices[i+1], expectedVertices[i])
		}
	}
}

func TestParseTriangleFaces(t *testing.T) {
	// Parsing triangle faces.
	data := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
f 1 2 3
f 1 3 4
`
	parser := parseObjData(data)
	group := parser.defaultGroup()
	t1 := group.children[0].(*Triangle)
	t2 := group.children[1].(*Triangle)

	if !(t1.p1.Equals(parser.vertices[1])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t1.p1, parser.vertices[1])
	}
	if !(t1.p2.Equals(parser.vertices[2])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t1.p2, parser.vertices[2])
	}
	if !(t1.p3.Equals(parser.vertices[3])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t1.p3, parser.vertices[3])
	}
	if !(t2.p1.Equals(parser.vertices[1])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t2.p1, parser.vertices[1])
	}
	if !(t2.p2.Equals(parser.vertices[3])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t2.p2, parser.vertices[3])
	}
	if !(t2.p3.Equals(parser.vertices[4])) {
		t.Errorf("Parsing triangle faces, got: %v and expected to be %v", t2.p3, parser.vertices[4])
	}
}

func TestTriangulatePolygon(t *testing.T) {
	// Triangulating polygons.
	data := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
v 0 2 0
f 1 2 3 4 5`
	parser := parseObjData(data)
	group := parser.defaultGroup()
	t1 := group.children[0].(*Triangle)
	t2 := group.children[1].(*Triangle)
	t3 := group.children[2].(*Triangle)

	if !(t1.p1.Equals(parser.vertices[1])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p1, parser.vertices[1])
	}
	if !(t1.p2.Equals(parser.vertices[2])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p2, parser.vertices[2])
	}
	if !(t1.p3.Equals(parser.vertices[3])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p3, parser.vertices[3])
	}
	if !(t2.p1.Equals(parser.vertices[1])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p1, parser.vertices[1])
	}
	if !(t2.p2.Equals(parser.vertices[3])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p2, parser.vertices[3])
	}
	if !(t2.p3.Equals(parser.vertices[4])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p3, parser.vertices[4])
	}
	if !(t3.p1.Equals(parser.vertices[1])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t3.p1, parser.vertices[1])
	}
	if !(t3.p2.Equals(parser.vertices[4])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t3.p2, parser.vertices[4])
	}
	if !(t3.p3.Equals(parser.vertices[5])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p3, parser.vertices[5])
	}
}

func TestTrianglesInGroups(t *testing.T) {
	// Triangles in groups.
	data := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
g FirstGroup
f 1 2 3
g SecondGroup
f 1 3 4`

	parser := parseObjData(data)
	g1 := parser.groups["FirstGroup"]
	g2 := parser.groups["SecondGroup"]
	t1 := g1.children[0].(*Triangle)
	t2 := g2.children[0].(*Triangle)

	if !(t1.p1.Equals(parser.vertices[1])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p1, parser.vertices[1])
	}
	if !(t1.p2.Equals(parser.vertices[2])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p2, parser.vertices[2])
	}
	if !(t1.p3.Equals(parser.vertices[3])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t1.p3, parser.vertices[3])
	}
	if !(t2.p1.Equals(parser.vertices[1])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p1, parser.vertices[1])
	}
	if !(t2.p2.Equals(parser.vertices[3])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p2, parser.vertices[3])
	}
	if !(t2.p3.Equals(parser.vertices[4])) {
		t.Errorf("Triangulating polygons, got: %v and expected to be %v", t2.p3, parser.vertices[4])
	}
}

func TestNormalData(t *testing.T) {
	// Vertex normal records.
	data := `
vn 0 0 1
vn 0.707 0 -0.707
vn 1 2 3`

	parser := parseObjData(data)

	expectedNormals := []*Tuple{
		Vector(0, 0, 1),
		Vector(0.707, 0, -0.707),
		Vector(1, 2, 3),
	}

	// parser will refer to the normals by their index, starting with 1.
	for i := range expectedNormals {
		if !(parser.normals[i+1].Equals(expectedNormals[i])) {
			t.Errorf("Vertex normal records, got: %v and expected to be %v", parser.normals[i+1], expectedNormals[i])
		}
	}
}

func TestFacesWithNormals(t *testing.T) {
	// Faces with normals.
	data := `
v 0 1 0
v -1 0 0
v 1 0 0
vn -1 0 0
vn 1 0 0
vn 0 1 0
f 1//3 2//1 3//2
f 1/0/3 2/102/1 3/14/2`
	parser := parseObjData(data)

	g := parser.defaultGroup()
	t1 := g.children[0].(*smoothTriangle)
	t2 := g.children[1].(*smoothTriangle)

	if !(t1.p1.Equals(parser.vertices[1])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.p1, parser.vertices[1])
	}
	if !(t1.p2.Equals(parser.vertices[2])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.p2, parser.vertices[2])
	}
	if !(t1.p3.Equals(parser.vertices[3])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.p3, parser.vertices[3])
	}
	if !(t1.n1.Equals(parser.normals[3])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.n1, parser.normals[3])
	}
	if !(t1.n2.Equals(parser.normals[1])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.n2, parser.normals[1])
	}
	if !(t1.n3.Equals(parser.normals[2])) {
		t.Errorf("Faces with normals, got: %v and expected to be %v", t1.n3, parser.normals[2])
	}
	if !(reflect.DeepEqual(t1, t2)) {
		t.Errorf("Faces with normals, t1 is not equal to t2")
	}
}
