package opencv3

/*
#include <stdlib.h>
#include "objdetect.h"
*/
import "C"
import (
	"image"
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

// Close deletes the CascadeClassifier's pointer.
func (c *CascadeClassifier) Close() error {
	C.CascadeClassifier_Close(c.p)
	c.p = nil
	return nil
}

// Load cascade configuration file to classifier.
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// DetectMultiScale detects something which is decided by loaded file. Returns
// multi results addressed with rectangle.
func (c *CascadeClassifier) DetectMultiScale(img Mat) []image.Rectangle {
	ret := C.CascadeClassifier_DetectMultiScale(c.p, img.p)
	defer C.Rects_Close(ret)

	cArray := ret.rects
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.Rect)(unsafe.Pointer(&hdr))

	rects := make([]image.Rectangle, length)
	for i, r := range s {
		rects[i] = image.Rect(int(r.x), int(r.y), int(r.x+r.width), int(r.y+r.height))
	}
	return rects
}
