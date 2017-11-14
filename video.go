package gocv

/*
#include <stdlib.h>
#include "video.h"
*/
import "C"

// BackgroundSubtractorMOG2 is a wrapper around the cv::BackgroundSubtractorMOG2.
type BackgroundSubtractorMOG2 struct {
	p C.BackgroundSubtractorMOG2
}

// NewBackgroundSubtractorMOG2 returns a new BackgroundSubtractor algorithm
// of type MOG2. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/d7b/classcv_1_1BackgroundSubtractorMOG2.html
//
func NewBackgroundSubtractorMOG2() BackgroundSubtractorMOG2 {
	return BackgroundSubtractorMOG2{p: C.BackgroundSubtractorMOG2_Create()}
}

// Close BackgroundSubtractorMOG2.
func (b *BackgroundSubtractorMOG2) Close() error {
	C.BackgroundSubtractorMOG2_Close(b.p)
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG2.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/df6/classcv_1_1BackgroundSubtractor.html#aa735e76f7069b3fa9c3f32395f9ccd21
//
func (b *BackgroundSubtractorMOG2) Apply(src Mat, dst Mat) {
	C.BackgroundSubtractorMOG2_Apply(b.p, src.p, dst.p)
	return
}

// BackgroundSubtractorKNN is a wrapper around the cv::BackgroundSubtractorKNN.
type BackgroundSubtractorKNN struct {
	p C.BackgroundSubtractorKNN
}

// NewBackgroundSubtractorKNN returns a new BackgroundSubtractor algorithm
// of type KNN. K-Nearest Neighbors (KNN) uses a Background/Foreground
// Segmentation Algorithm
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/db/d88/classcv_1_1BackgroundSubtractorKNN.html
//
func NewBackgroundSubtractorKNN() BackgroundSubtractorKNN {
	return BackgroundSubtractorKNN{p: C.BackgroundSubtractorKNN_Create()}
}

// Close BackgroundSubtractorKNN.
func (k *BackgroundSubtractorKNN) Close() error {
	C.BackgroundSubtractorKNN_Close(k.p)
	k.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorKNN.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d7/df6/classcv_1_1BackgroundSubtractor.html#aa735e76f7069b3fa9c3f32395f9ccd21
//
func (k *BackgroundSubtractorKNN) Apply(src Mat, dst Mat) {
	C.BackgroundSubtractorKNN_Apply(k.p, src.p, dst.p)
	return
}
