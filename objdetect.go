package gocv

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

// CascadeClassifier is a cascade classifier class for object detection.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d1/de5/classcv_1_1CascadeClassifier.html
//
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

// Load cascade classifier from a file.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d1/de5/classcv_1_1CascadeClassifier.html#a1a5884c8cc749422f9eb77c2471958bc
//
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
//
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
