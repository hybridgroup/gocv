//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_cuda_optflow)

package cuda

import "testing"

func TestSparsePyrLKOpticalFlow_Calc(t *testing.T) {
	prevImg := NewGpuMat()
	defer prevImg.Close()

	nextImg := NewGpuMat()
	defer nextImg.Close()

	prevPts := NewGpuMat()
	defer prevPts.Close()

	nextPts := NewGpuMat()
	defer nextPts.Close()

	status := NewGpuMat()
	defer status.Close()

	pyrLk := NewSparsePyrLKOpticalFlow()
	pyrLk.Calc(prevImg, nextImg, prevPts, nextPts, status)
}
