package main

import (
	"strconv"
	"strings"
)

//EPSILON is the error tolerance used for practical comparisons.
const EPSILON = 0.00001

//floatEqual determines if two floats are equal within a tolerance Epsilon.
func floatEqual(a, b float64) bool {
	return abs(a-b) < EPSILON
}

//Abs returns absolute value
func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

// floatToString converts a float to a String
func floatToString(n float64, cut int) string {
	// to convert a float number to a string
	s := strconv.FormatFloat(n, 'f', 6, 64)
	if cut > len(s) {
		return s[:]
	}
	return s[:cut]
}

func floatToUint8String(f float64) string {
	if f < 0.0 {
		return "0"
	}
	f *= 255.0
	if f > 255.0 {
		return "255"
	}
	return strconv.Itoa(int(f))
}

// split the string into a second item of the slice when the original string surpasses
// the character limit.
func split(s string, lim int) []string {
	l := []string{}
	for len(s) > lim {
		idx := strings.LastIndex(s[:lim+1], " ")
		l = append(l, s[:idx])
		s = s[idx+1:]
	}
	l = append(l, s)
	return l
}
