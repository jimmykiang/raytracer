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
