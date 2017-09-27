package opencv3

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"

const (
	// CvCapPropFrameWidth is OpenCV parameter of Frame Width
	CvCapPropFrameWidth = 3
	// CvCapPropFrameHeight is OpenCV parameter of Frame Height
	CvCapPropFrameHeight = 4
	// CvCapPropFps is OpenCV parameter of FPS
	CvCapPropFps = 5
)

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

// Rect represents rectangle. X and Y is a start point of Width and Height.
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}
