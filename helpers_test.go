package gocv

import "math"

var eps = 0.00000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}
