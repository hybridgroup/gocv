package gocv

/*
#include <stdlib.h>
#include "features2d.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// SimpleBlobDetector is a wrapper around the cv::SimpleBlobDetector.
type SimpleBlobDetector struct {
	// C.SimpleBlobDetector
	p unsafe.Pointer
}

// NewSimpleBlobDetector returns a new SimpleBlobDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d7a/classcv_1_1SimpleBlobDetector.html
//
func NewSimpleBlobDetector() SimpleBlobDetector {
	return SimpleBlobDetector{p: unsafe.Pointer(C.SimpleBlobDetector_Create())}
}

// Close SimpleBlobDetector.
func (b *SimpleBlobDetector) Close() error {
	C.SimpleBlobDetector_Close((C.SimpleBlobDetector)(b.p))
	b.p = nil
	return nil
}

// Detect keypoints in an image using SimpleBlobDetector.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (b *SimpleBlobDetector) Detect(src Mat) []KeyPoint {
	ret := C.SimpleBlobDetector_Detect((C.SimpleBlobDetector)(b.p), src.p)
	defer C.KeyPoints_Close(ret)

	cArray := ret.keypoints
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.KeyPoint)(unsafe.Pointer(&hdr))

	keys := make([]KeyPoint, length)
	for i, r := range s {
		keys[i] = KeyPoint{float64(r.x), float64(r.y), float64(r.size), float64(r.angle), float64(r.response),
			int(r.octave), int(r.classID)}
	}
	return keys
}
