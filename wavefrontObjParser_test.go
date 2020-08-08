package main

import "testing"

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
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t1.p1, parser.vertices[1])
	}
	if !(t1.p2.Equals(parser.vertices[2])) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t1.p2, parser.vertices[2])
	}
	if !(t1.p3.Equals(parser.vertices[3])) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t1.p3, parser.vertices[3])
	}
	if !(t2.p1.Equals(parser.vertices[1])) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t2.p1, parser.vertices[1])
	}
	if !(t2.p2.Equals(parser.vertices[3])) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t2.p2, parser.vertices[3])
	}
	if !(t2.p3.Equals(parser.vertices[4])) {
		t.Errorf("Ignoring unrecognized lines, got: %v and expected to be %v", t2.p3, parser.vertices[4])
	}
}
