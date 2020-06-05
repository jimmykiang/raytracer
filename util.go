package main

import "strconv"

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

// FloatToString converts a float to a String
func floatToString(n float64, cut int) string {
	// to convert a float number to a string
	s := strconv.FormatFloat(n, 'f', 6, 64)
	if cut > len(s) {
		return s[:]
	}
	return s[:cut]
}
