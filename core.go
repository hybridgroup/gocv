package opencv3

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"

// CMat is an alias for C pointer.
type CMat C.Mat

// Mat is a bind of `cv::Mat
type Mat struct {
	p C.Mat
}

// NewMat returns a new Mat.
func NewMat() Mat {
	return Mat{p: C.Mat_New()}
}

// Close the Mat object.
func (m *Mat) Close() {
	C.Mat_Close(m.p)
	m.p = nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return m.p
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Point represents a single point with X and Y coordinates.
type Point struct {
	X int
	Y int
}

// Rect represents rectangle. X and Y is a start point of Width and Height.
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Size represents a size of something with Width and Height.
type Size struct {
	Width  int
	Height int
}

// Scalar represents a Scalar set of 4 float64 values.
type Scalar struct {
	Val1 float64
	Val2 float64
	Val3 float64
	Val4 float64
}

// NewScalar returns a new Scalar. These are usually colors typically being in BGR order.
func NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
	s := Scalar{Val1: v1, Val2: v2, Val3: v3, Val4: v4}
	return s
}
