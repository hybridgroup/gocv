package pvl

/*
#include <stdlib.h>
#include "pvl.h"
*/
import "C"

// Face is a wrapper around `cv::pvl::Face`.
type Face struct {
	p C.Face
}

// NewFace returns a new PVL Face
func NewFace() Face {
	return Face{p: C.Face_New()}
}

// Delete Face.
func (f *Face) Delete() {
	C.Face_Delete(f.p)
	f.p = nil
}

// Rect returns the Rect for this Face
func (f *Face) Rect() Rect {
	C.struct_Rect(C.Face_GetRect(f.p))
}

// FaceDetector is a bind of `cv::pvl::FaceDetector`.
type FaceDetector struct {
	p C.FaceDetector
}

// NewFaceDetector returns a new PVL FaceDetector.
func NewFaceDetector() FaceDetector {
	return FaceDetector{p: C.FaceDetector_New()}
}

// Delete object.
func (f *FaceDetector) Delete() {
	C.FaceDetector_Delete(f.p)
	f.p = nil
}

// SetTrackingModeEnabled sets if the PVL FaceDetector tracking mode is enabled.
func (f *FaceDetector) SetTrackingModeEnabled(enabled bool) {
	C.FaceDetector_SetTrackingModeEnabled(f.p, C.bool(enabled))
}
