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

// Delete object.
func (m *Mat) Delete() {
	C.Mat_Delete(m.p)
	m.p = nil
}

// Delete object.
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

// DrawRectsToImage draws rectangle information to target image Mat.
func DrawRectsToImage(img Mat, rects []Rect) {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.X),
			y:      C.int(r.Y),
			width:  C.int(r.Width),
			height: C.int(r.Height),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}
	C.DrawRectsToImage(img.p, cRects)
}
