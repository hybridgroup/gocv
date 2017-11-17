package gocv

/*
#include <stdlib.h>
#include "features2d.h"
*/
import "C"

// SimpleBlobDetector is a wrapper around the cv::SimpleBlobDetector.
type SimpleBlobDetector struct {
	p C.SimpleBlobDetector
}

// NewSimpleBlobDetector returns a new SimpleBlobDetector algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d7a/classcv_1_1SimpleBlobDetector.html
//
func NewSimpleBlobDetector() SimpleBlobDetector {
	return SimpleBlobDetector{p: C.SimpleBlobDetector_Create()}
}

// Close SimpleBlobDetector.
func (b *SimpleBlobDetector) Close() error {
	C.SimpleBlobDetector_Close(b.p)
	b.p = nil
	return nil
}

// Detect keypoints in an image
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
//
func (b *SimpleBlobDetector) Detect(src Mat) {
	C.SimpleBlobDetector_Detect(b.p, src.p)
	return
}
