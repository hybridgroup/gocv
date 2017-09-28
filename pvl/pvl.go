package pvl

/*
#include <stdlib.h>
#include "pvl.h"
*/
import "C"

import (
	"reflect"
	"unsafe"

	opencv3 ".."
)

// Face is a wrapper around `cv::pvl::Face`.
type Face struct {
	p C.Face
}

// NewFace returns a new PVL Face
func NewFace() Face {
	return Face{p: C.Face_New()}
}

// Close Face.
func (f *Face) Close() {
	C.Face_Close(f.p)
	f.p = nil
}

// Ptr returns the Face's underlying object pointer.
func (f *Face) Ptr() C.Face {
	return f.p
}

// Rect returns the Rect for this Face
func (f *Face) Rect() opencv3.Rect {
	rect := C.Face_GetRect(f.p)
	return opencv3.Rect{X: int(rect.x),
		Y:      int(rect.y),
		Width:  int(rect.width),
		Height: int(rect.height)}
}

// IsSmiling Face? :)
// You must call FaceDetector's DetectEye() and DetectSmile() with this Face
// first, or this function will throw an exception
func (f *Face) IsSmiling() bool {
	return bool(C.Face_IsSmiling(f.p))
}

// FaceDetector is a bind of `cv::pvl::FaceDetector`.
type FaceDetector struct {
	p C.FaceDetector
}

// NewFaceDetector returns a new PVL FaceDetector.
func NewFaceDetector() FaceDetector {
	return FaceDetector{p: C.FaceDetector_New()}
}

// Close FaceDetector.
func (f *FaceDetector) Close() {
	C.FaceDetector_Close(f.p)
	f.p = nil
}

// SetTrackingModeEnabled sets if the PVL FaceDetector tracking mode is enabled.
func (f *FaceDetector) SetTrackingModeEnabled(enabled bool) {
	C.FaceDetector_SetTrackingModeEnabled(f.p, C.bool(enabled))
}

// DetectFaceRect tries to detect Faces from the image Mat passed in as the param.
// The Mat must be a grayed image that has only one channel and 8-bit depth.
func (f *FaceDetector) DetectFaceRect(img opencv3.Mat) []Face {
	ret := C.FaceDetector_DetectFaceRect(f.p, C.Mat(img.Ptr()))
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

// DetectEyes uses PVL FaceDetector to detect eyes on a Face
func (f *FaceDetector) DetectEye(img opencv3.Mat, face Face) {
	C.FaceDetector_DetectEye(f.p, C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}

// DetectSmile uses PVL FaceDetector to detect eyes on a Face
func (f *FaceDetector) DetectSmile(img opencv3.Mat, face Face) {
	C.FaceDetector_DetectSmile(f.p, C.Mat(img.Ptr()), C.Face(face.Ptr()))
	return
}
