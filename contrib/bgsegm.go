package contrib

/*
#include <stdlib.h>
#include "bgsegm.h"
*/
import "C"

import (
	"unsafe"

	"gocv.io/x/gocv"
)

// BackgroundSubtractorCNT is a wrapper around the cv::BackgroundSubtractorCNT.
type BackgroundSubtractorCNT struct {
	// C.BackgroundSubtractorCNT
	p unsafe.Pointer
}

// NewBackgroundSubtractorCNT returns a new BackgroundSubtractor algorithm
// of type CNT. CNT is Background subtraction algorithm based on counting.
// About as fast as MOG2 on a high end system. More than twice faster than MOG2 on cheap hardware (benchmarked on Raspberry Pi3).
// Algorithm by Sagi Zeevi
//
// For further details, please see:
// https://docs.opencv.org/3.4/de/dca/classcv_1_1bgsegm_1_1BackgroundSubtractorCNT.html
//
func NewBackgroundSubtractorCNT() BackgroundSubtractorCNT {
	return BackgroundSubtractorCNT{p: unsafe.Pointer(C.BackgroundSubtractorCNT_Create())}
}

// Close BackgroundSubtractorCNT.
func (b *BackgroundSubtractorCNT) Close() error {
	C.BackgroundSubtractorCNT_Close((C.BackgroundSubtractorCNT)(b.p))
	b.p = nil

	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorCNT.
//
// For further details, please see:
// https://docs.opencv.org/3.4/de/dca/classcv_1_1bgsegm_1_1BackgroundSubtractorCNT.html
//
func (b *BackgroundSubtractorCNT) Apply(src gocv.Mat, dst *gocv.Mat) {
	C.BackgroundSubtractorCNT_Apply((C.BackgroundSubtractorCNT)(b.p), (C.Mat)(src.Ptr()), (C.Mat)(dst.Ptr()))

	return
}
