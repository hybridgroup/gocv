// Package cuda is the GoCV wrapper around OpenCV cuda.
//
// For further details, please see:
// https://github.com/opencv/opencv
//
// import "gocv.io/x/gocv/cuda"
package cuda

/*
#include <stdlib.h>
#include "cuda.h"
*/
import "C"
import "gocv.io/x/gocv"

// GpuMat is the GPU version of a Mat
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html
type GpuMat struct {
	p C.GpuMat
}

// Upload performs data upload to GpuMat (Blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a00ef5bfe18d14623dcf578a35e40a46b
//
func (g *GpuMat) Upload(data gocv.Mat) {
	C.GpuMat_Upload(g.p, C.Mat(data.Ptr()))
}

// Download performs data download from GpuMat (Blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a027e74e4364ddfd9687b58aa5db8d4e8
func (g *GpuMat) Download(dst *gocv.Mat) {
	C.GpuMat_Download(g.p, C.Mat(dst.Ptr()))
}

// Empty returns true if GpuMat is empty
func (g *GpuMat) Empty() bool {
	return C.GpuMat_Empty(g.p) != 0
}

// Close the GpuMat object
func (g *GpuMat) Close() error {
	C.GpuMat_Close(g.p)
	g.p = nil
	return nil
}

// NewGpuMat returns a new empty GpuMat
func NewGpuMat() GpuMat {
	return newGpuMat(C.GpuMat_New())
}

// NewGpuMatFromMat returns a new GpuMat based on a Mat
func NewGpuMatFromMat(mat gocv.Mat) GpuMat {
	return newGpuMat(C.GpuMat_NewFromMat(C.Mat(mat.Ptr())))
}

func newGpuMat(p C.GpuMat) GpuMat {
	return GpuMat{p: p}
}

// PrintCudaDeviceInfo prints extensive cuda device information
func PrintCudaDeviceInfo(device int) {
	C.PrintCudaDeviceInfo(C.int(device))
}

// PrintShortCudaDeviceInfo prints a small amount of cuda device information
func PrintShortCudaDeviceInfo(device int) {
	C.PrintShortCudaDeviceInfo(C.int(device))
}

// GetCudaEnabledDeviceCount returns the number of cuda enabled devices on the
// system
func GetCudaEnabledDeviceCount() int {
	return int(C.GetCudaEnabledDeviceCount())
}

// ConvertTo converts GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a3a1b076e54d8a8503014e27a5440d98a
//
func (m *GpuMat) ConvertTo(dst *GpuMat, mt gocv.MatType) {
	C.GpuMat_ConvertTo(m.p, dst.p, C.int(mt))
	return
}
