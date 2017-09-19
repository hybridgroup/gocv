package opencv3

/*
#include <stdlib.h>
#include "objdetect.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// CascadeClassifier is a bind of `cv::CascadeClassifier`
type CascadeClassifier struct {
	p C.CascadeClassifier
}

// NewCascadeClassifier returns a new CascadeClassifier.
func NewCascadeClassifier() CascadeClassifier {
	return CascadeClassifier{p: C.CascadeClassifier_New()}
}

// Delete CascadeClassifier's pointer.
func (c *CascadeClassifier) Delete() {
	C.CascadeClassifier_Delete(c.p)
	c.p = nil
}

// Load cascade configuration file to classifier.
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// DetectMultiScale detects something which is decided by loaded file. Returns
// multi results addressed with rectangle.
func (c *CascadeClassifier) DetectMultiScale(img MatVec3b) []Rect {
	ret := C.CascadeClassifier_DetectMultiScale(c.p, img.p)
	defer C.Rects_Delete(ret)

	cArray := ret.rects
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	goSlice := *(*[]C.Rect)(unsafe.Pointer(&hdr))

	rects := make([]Rect, length)
	for i, r := range goSlice {
		rects[i] = Rect{
			X:      int(r.x),
			Y:      int(r.y),
			Width:  int(r.width),
			Height: int(r.height),
		}
	}
	return rects
}
