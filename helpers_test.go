package gocv

import "math"

var eps = 0.00000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}
