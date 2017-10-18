package gocv

/*
#include <stdlib.h>
#include "video.h"
*/
import "C"

// BackgroundSubtractor is a wrapper around the cv::BackgroundSubtractor.
type BackgroundSubtractor struct {
	p C.BackgroundSubtractor
}

// NewBackgroundSubtractorMOG2 returns a new BackgroundSubtractor algorithm
// of type MOG2. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d7/d7b/classcv_1_1BackgroundSubtractorMOG2.html
//
func NewBackgroundSubtractorMOG2() BackgroundSubtractor {
	return BackgroundSubtractor{p: C.BackgroundSubtractor_CreateMOG2()}
}

// Close BackgroundSubtractor.
func (b *BackgroundSubtractor) Close() error {
	C.BackgroundSubtractor_Close(b.p)
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractor.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d7/df6/classcv_1_1BackgroundSubtractor.html#aa735e76f7069b3fa9c3f32395f9ccd21
//
func (b *BackgroundSubtractor) Apply(src Mat, dst Mat) {
	C.BackgroundSubtractor_Apply(b.p, src.p, dst.p)
	return
}
