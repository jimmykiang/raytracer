package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Obj contains the information processed from the wavefront OBJ file.
type Obj struct {
	vertices     []*Tuple
	ignoredLines int
}

func handlepanic() {

	if r := recover(); r != nil {

		fmt.Println("RECOVERED FROM:", r)
	}
}

// parseObjData parses the data in wavefront OBJ file
func parseObjData(data string) *Obj {

	defer handlepanic()

	result := &Obj{
		vertices:     make([]*Tuple, 0),
		ignoredLines: 0,
	}

	lines := strings.Split(data, "\n")

	var x, y, z float64
	var err error

	// It is significant that the vertices slice is 1-based, and not 0-based,
	// refer to these vertices by their index, starting with 1.
	result.vertices = append(result.vertices, Point(0, 0, 0))
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			tokenSlice := strings.Fields(strings.TrimSpace(line))
			switch tokenSlice[0] {
			case "v":

				if x, err = strconv.ParseFloat(tokenSlice[1], 64); err != nil {

					panic(err)
				}
				if y, err = strconv.ParseFloat(tokenSlice[2], 64); err != nil {

					panic(err)
				}
				if z, err = strconv.ParseFloat(tokenSlice[3], 64); err != nil {

					panic(err)
				}
				result.vertices = append(result.vertices, Point(x, y, z))

			default:
				result.ignoredLines++
			}
		} else {
			result.ignoredLines++
		}
	}

	fmt.Println("Wavefront OBJ loaded:")
	fmt.Printf("Vertices: %d\n", len(result.vertices))

	return result
}
