package pvl

/*
#include <stdlib.h>
#include "pvl.h"
*/
import "C"

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
