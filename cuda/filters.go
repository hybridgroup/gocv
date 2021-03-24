package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "core.h"
#include "filters.h"
*/
import "C"
import (
	"image"
	"unsafe"

	"gocv.io/x/gocv"
)

// GaussianFilter
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gaa4df286369114cfd4b144ae211f6a6c8
//
type GaussianFilter struct {
	p unsafe.Pointer
}

// NewGaussianFilter returns a new GaussianFilter.
func NewGaussianFilter(srcType gocv.MatType, dstType gocv.MatType, ksize image.Point, sigma1 float64) GaussianFilter {
	pSize := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	return GaussianFilter{p: unsafe.Pointer(C.CreateGaussianFilter(C.int(srcType), C.int(dstType), pSize, C.double(sigma1)))}
}

// Close GaussianFilter
func (gf *GaussianFilter) Close() error {
	C.GaussianFilter_Close((C.GaussianFilter)(gf.p))
	gf.p = nil
	return nil
}

// Apply applies the Gaussian filter.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
//
func (gf *GaussianFilter) Apply(img GpuMat) GpuMat {
	return newGpuMat(C.GaussianFilter_Apply(C.GaussianFilter(gf.p), img.p))
}
