package cuda

/*
#include <stdlib.h>
#include "cuda.h"
#include "cudaoptflow.h"
*/
import "C"
import "unsafe"

// SparsePyrLKOpticalFlow is a wrapper around the cv::cuda::SparsePyrLKOpticalFlow.
type SparsePyrLKOpticalFlow struct {
	// C.SparsePyrLKOpticalFlow
	p unsafe.Pointer
}

// NewSparsePyrLKOpticalFlow returns a new SparsePyrLKOpticalFlow
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d05/classcv_1_1cuda_1_1SparsePyrLKOpticalFlow.html#a6bcd2d457532d7db76c3e7f11b60063b
//
func NewSparsePyrLKOpticalFlow() SparsePyrLKOpticalFlow {
	return SparsePyrLKOpticalFlow{p: unsafe.Pointer(C.CudaSparsePyrLKOpticalFlow_Create())}
}

// Calc calculates a sparse optical flow.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/dcf/classcv_1_1cuda_1_1SparseOpticalFlow.html#a80d5efbb7788e3dc4c49e6226ba34347
func (s SparsePyrLKOpticalFlow) Calc(prevImg, nextImg, prevPts, nextPts, status GpuMat) {
	C.CudaSparsePyrLKOpticalFlow_Calc(C.CudaSparsePyrLKOpticalFlow(s.p), prevImg.p, nextImg.p, prevPts.p, nextPts.p, status.p)
}
