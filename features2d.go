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

// AgastFeatureDetector is a wrapper around the cv::AgastFeatureDetector.
type AgastFeatureDetector struct {
	// C.AgastFeatureDetector
	p unsafe.Pointer
}

// NewAgastFeatureDetector returns a new AgastFeatureDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/d19/classcv_1_1AgastFeatureDetector.html
//
func NewAgastFeatureDetector() AgastFeatureDetector {
	return AgastFeatureDetector{p: unsafe.Pointer(C.AgastFeatureDetector_Create())}
}

// Close AgastFeatureDetector.
func (a *AgastFeatureDetector) Close() error {
	C.AgastFeatureDetector_Close((C.AgastFeatureDetector)(a.p))
	a.p = nil
	return nil
}

// Detect keypoints in an image using AgastFeatureDetector.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (a *AgastFeatureDetector) Detect(src Mat) []KeyPoint {
	ret := C.AgastFeatureDetector_Detect((C.AgastFeatureDetector)(a.p), src.p)
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

// FastFeatureDetector is a wrapper around the cv::FastFeatureDetector.
type FastFeatureDetector struct {
	// C.FastFeatureDetector
	p unsafe.Pointer
}

// NewFastFeatureDetector returns a new FastFeatureDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/df/d74/classcv_1_1FastFeatureDetector.html
//
func NewFastFeatureDetector() FastFeatureDetector {
	return FastFeatureDetector{p: unsafe.Pointer(C.FastFeatureDetector_Create())}
}

// Close FastFeatureDetector.
func (f *FastFeatureDetector) Close() error {
	C.FastFeatureDetector_Close((C.FastFeatureDetector)(f.p))
	f.p = nil
	return nil
}

// Detect keypoints in an image using FastFeatureDetector.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (f *FastFeatureDetector) Detect(src Mat) []KeyPoint {
	ret := C.FastFeatureDetector_Detect((C.FastFeatureDetector)(f.p), src.p)
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

// ORB is a wrapper around the cv::ORB.
type ORB struct {
	// C.ORB
	p unsafe.Pointer
}

// NewORB returns a new ORB algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/d19/classcv_1_1AgastFeatureDetector.html
//
func NewORB() ORB {
	return ORB{p: unsafe.Pointer(C.ORB_Create())}
}

// Close ORB.
func (o *ORB) Close() error {
	C.ORB_Close((C.ORB)(o.p))
	o.p = nil
	return nil
}

// Detect keypoints in an image using ORB.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (o *ORB) Detect(src Mat) []KeyPoint {
	ret := C.ORB_Detect((C.ORB)(o.p), src.p)
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
