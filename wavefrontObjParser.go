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
	groups       map[string]*Group
}

func handlepanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVERED FROM:", r)
	}
}

// DefaultGroup returns a DefaultGroup key in the Group map for Obj.
func (o *Obj) defaultGroup() *Group {
	return o.groups["defaultGroup"]
}

// parseObjData parses the data in wavefront OBJ file
func parseObjData(data string) *Obj {

	defer handlepanic()

	result := &Obj{
		vertices:     make([]*Tuple, 0),
		ignoredLines: 0,
		groups:       make(map[string]*Group),
	}

	lines := strings.Split(data, "\n")

	var x, y, z float64
	var index1, index2, index3 int
	var err error
	currentGroup := "defaultGroup"
	result.groups[currentGroup] = NewGroup()

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

			case "f":

				for i := 2; i < len(tokenSlice)-1; i++ {
					if index1, err = strconv.Atoi(tokenSlice[1]); err != nil {
						panic(err)
					}
					if index2, err = strconv.Atoi(tokenSlice[i]); err != nil {
						panic(err)
					}
					if index3, err = strconv.Atoi(tokenSlice[i+1]); err != nil {
						panic(err)
					}
					triangle := NewTriangle(
						result.vertices[index1],
						result.vertices[index2],
						result.vertices[index3])
					result.groups[currentGroup].AddChild(triangle)
				}

			case "g":
				fallthrough
			case "o":
				currentGroup = strings.Fields(strings.TrimSpace(line))[1]
				if _, exists := result.groups[currentGroup]; !exists {

					result.groups[currentGroup] = NewGroup()
				}

			default:
				result.ignoredLines++
			}
		} else {
			result.ignoredLines++
		}
	}

	triangles := 0
	for i := range result.groups {
		triangles += len(result.groups[i].children)
	}

	fmt.Println("Wavefront OBJ loaded:")
	fmt.Printf("Groups:    %d\n", len(result.groups))
	fmt.Printf("Vertices: %d\n", len(result.vertices)-1)
	fmt.Printf("Triangles: %d\n", triangles)

	return result
}

func (o *Obj) objToGroup() *Group {
	g := NewGroup()
	for _, v := range o.groups {
		g.AddChild(v)
	}
	return g
}
