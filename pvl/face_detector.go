package pvl

/*
#include <stdlib.h>
#include "face_detector.h"
*/
import "C"

import (
	"reflect"
	"unsafe"

	"gocv.io/x/gocv"
)

// FaceDetector is a wrapper around the cv::pvl::FaceDetector.
type FaceDetector struct {
	// C.FaceDetector
	p unsafe.Pointer
}

// NewFaceDetector returns a new PVL FaceDetector.
func NewFaceDetector() FaceDetector {
	return FaceDetector{p: unsafe.Pointer(C.FaceDetector_New())}
}

// Close FaceDetector.
func (f *FaceDetector) Close() error {
	C.FaceDetector_Close((C.FaceDetector)(f.p))
	f.p = nil
	return nil
}

// SetTrackingModeEnabled sets if the PVL FaceDetector tracking mode is enabled.
func (f *FaceDetector) SetTrackingModeEnabled(enabled bool) {
	C.FaceDetector_SetTrackingModeEnabled((C.FaceDetector)(f.p), C.bool(enabled))
}

// DetectFaceRect tries to detect Faces from the image Mat passed in as the param.
// The Mat must be a grayed image that has only one channel and 8-bit depth.
func (f *FaceDetector) DetectFaceRect(img gocv.Mat) []Face {
	ret := C.FaceDetector_DetectFaceRect((C.FaceDetector)(f.p), C.Mat(img.Ptr()))
	fArray := ret.faces
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(fArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.Face)(unsafe.Pointer(&hdr))

	faces := make([]Face, length)
	for i, r := range s {
		faces[i] = Face{p: r}
	}
	return faces
}

// DetectEye uses PVL FaceDetector to detect eyes on a Face.
func (f *FaceDetector) DetectEye(img gocv.Mat, face Face) {
	C.FaceDetector_DetectEye((C.FaceDetector)(f.p), C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}

// DetectMouth uses PVL FaceDetector to detect mouth on a Face.
func (f *FaceDetector) DetectMouth(img gocv.Mat, face Face) {
	C.FaceDetector_DetectMouth((C.FaceDetector)(f.p), C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}

// DetectSmile uses PVL FaceDetector to detect smile on a Face.
func (f *FaceDetector) DetectSmile(img gocv.Mat, face Face) {
	C.FaceDetector_DetectSmile((C.FaceDetector)(f.p), C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}

// DetectBlink uses PVL FaceDetector to detect blink on a Face.
func (f *FaceDetector) DetectBlink(img gocv.Mat, face Face) {
	C.FaceDetector_DetectBlink((C.FaceDetector)(f.p), C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}
