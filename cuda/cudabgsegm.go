package cuda

/*
#include <stdlib.h>
#include "cudabgsegm.h"
*/
import "C"
import "unsafe"

// BackgroundSubtractorMOG2 is a wrapper around the cv::cuda::BackgroundSubtractorMOG2.
type BackgroundSubtractorMOG2 struct {
	// C.BackgroundSubtractorMOG2
	p unsafe.Pointer
}

// BackgroundSubtractorMOG is a wrapper around the cv::cuda::BackgroundSubtractorMOG.
type BackgroundSubtractorMOG struct {
	// C.BackgroundSubtractorMOG
	p unsafe.Pointer
}

// NewBackgroundSubtractorMOG2 returns a new BackgroundSubtractor algorithm
// of type MOG2. MOG2 is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d3d/cudabgsegm_8hpp.html
//
func NewBackgroundSubtractorMOG2() BackgroundSubtractorMOG2 {
	return BackgroundSubtractorMOG2{p: unsafe.Pointer(C.CudaBackgroundSubtractorMOG2_Create())}
}

// Close BackgroundSubtractorMOG2.
func (b *BackgroundSubtractorMOG2) Close() error {
	C.CudaBackgroundSubtractorMOG2_Close((C.CudaBackgroundSubtractorMOG2)(b.p))
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG2.
//
// For further details, please see:
// https://docs.opencv.org/master/df/d23/classcv_1_1cuda_1_1BackgroundSubtractorMOG2.html#a92408f07bf1268c1b778cb186b3113b0
//
func (b *BackgroundSubtractorMOG2) Apply(src GpuMat, dst *GpuMat) {
	C.CudaBackgroundSubtractorMOG2_Apply((C.CudaBackgroundSubtractorMOG2)(b.p), src.p, dst.p)
	return
}

// NewBackgroundSubtractorMOG returns a new BackgroundSubtractor algorithm
// of type MOG. MOG is a Gaussian Mixture-based Background/Foreground
// Segmentation Algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d3d/cudabgsegm_8hpp.html
//
func NewBackgroundSubtractorMOG() BackgroundSubtractorMOG {
	return BackgroundSubtractorMOG{p: unsafe.Pointer(C.CudaBackgroundSubtractorMOG_Create())}
}

// Close BackgroundSubtractorMOG.
func (b *BackgroundSubtractorMOG) Close() error {
	C.CudaBackgroundSubtractorMOG_Close((C.CudaBackgroundSubtractorMOG)(b.p))
	b.p = nil
	return nil
}

// Apply computes a foreground mask using the current BackgroundSubtractorMOG.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dfe/classcv_1_1cuda_1_1BackgroundSubtractorMOG.html#a8f52d2f7abd1c77c84243efc53972cbf
//
func (b *BackgroundSubtractorMOG) Apply(src GpuMat, dst *GpuMat) {
	C.CudaBackgroundSubtractorMOG_Apply((C.CudaBackgroundSubtractorMOG)(b.p), src.p, dst.p)
	return
}
