package cuda

import "math"

var eps = 0.00000001

func floatEquals(a, b float32) bool {
	if math.Abs(float64(a-b)) < eps {
		return true
	}
	return false
}

// round helper from https://stackoverflow.com/questions/39544571/golang-round-to-nearest-0-05
func round(x, unit float32) float32 {
	return float32(int32(x/unit+0.5)) * unit
}
