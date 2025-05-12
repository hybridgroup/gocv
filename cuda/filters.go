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
func (gf *GaussianFilter) Apply(img GpuMat, dst *GpuMat) error {
	return OpenCVResult(C.GaussianFilter_Apply(C.GaussianFilter(gf.p), img.p, dst.p, nil))
}

// ApplyWithStream applies the Gaussian filter
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
func (gf *GaussianFilter) ApplyWithStream(img GpuMat, dst *GpuMat, s Stream) error {
	return OpenCVResult(C.GaussianFilter_Apply(C.GaussianFilter(gf.p), img.p, dst.p, s.p))
}

// MorphologyFilter
//
// For further details, please see:
// https://docs.opencv.org/4.x/dc/d66/group__cudafilters.html#gae58694e07be6bdbae126f36c75c08ee6
type MorphologyFilter struct {
	p unsafe.Pointer
}

// NewMorphologyFilter returns a new MorphologyFilter.
func NewMorphologyFilter(op gocv.MorphType, srcType gocv.MatType, kernel gocv.Mat) MorphologyFilter {
	return MorphologyFilter{p: unsafe.Pointer(C.CreateMorphologyFilter(C.int(op), C.int(srcType), C.Mat(kernel.Ptr())))}
}

// NewMorphologyFilterWithParams returns a new MorphologyFilter.
func NewMorphologyFilterWithParams(op gocv.MorphType, srcType gocv.MatType, kernel gocv.Mat, anchor image.Point, iterations int) MorphologyFilter {
	pt := C.struct_Point{
		x: C.int(anchor.X),
		y: C.int(anchor.Y),
	}
	return MorphologyFilter{p: unsafe.Pointer(C.CreateMorphologyFilterWithParams(C.int(op), C.int(srcType), C.Mat(kernel.Ptr()), pt, C.int(iterations)))}
}

// Close MorphologyFilter
func (mf *MorphologyFilter) Close() error {
	C.MorphologyFilter_Close((C.MorphologyFilter)(mf.p))
	mf.p = nil
	return nil
}

// Apply applies the Morphology filter.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
func (mf *MorphologyFilter) Apply(img GpuMat, dst *GpuMat) error {
	return OpenCVResult(C.MorphologyFilter_Apply(C.MorphologyFilter(mf.p), img.p, dst.p, nil))
}

// ApplyWithStream applies the Morphology filter
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
func (mf *MorphologyFilter) ApplyWithStream(img GpuMat, dst *GpuMat, s Stream) error {
	return OpenCVResult(C.MorphologyFilter_Apply(C.MorphologyFilter(mf.p), img.p, dst.p, s.p))
}

// SobelFilter
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gabf85fe61958bb21e93211a6fcc7c5c3b
type SobelFilter struct {
	p unsafe.Pointer
}

// NewSobelFilter returns a new SobelFilter.
func NewSobelFilter(srcType gocv.MatType, dstType gocv.MatType, dx int, dy int) SobelFilter {
	return SobelFilter{p: unsafe.Pointer(C.CreateSobelFilter(C.int(srcType), C.int(dstType), C.int(dx), C.int(dy)))}
}

// NewSobelFilterWithParams returns a new SobelFilter.
func NewSobelFilterWithParams(srcType gocv.MatType, dstType gocv.MatType, dx int, dy int, ksize int, scale float64, rowBorderMode int, columnBorderMode int) SobelFilter {
	return SobelFilter{p: unsafe.Pointer(C.CreateSobelFilterWithParams(C.int(srcType), C.int(dstType), C.int(dx), C.int(dy), C.int(ksize), C.double(scale), C.int(rowBorderMode), C.int(columnBorderMode)))}
}

// Close SobelFilter
func (sf *SobelFilter) Close() error {
	C.SobelFilter_Close((C.SobelFilter)(sf.p))
	sf.p = nil
	return nil
}

// Apply applies the Sobel filter.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
func (sf *SobelFilter) Apply(img GpuMat, dst *GpuMat) error {
	return OpenCVResult(C.SobelFilter_Apply(C.SobelFilter(sf.p), img.p, dst.p, nil))
}

// ApplyWithStream applies the Sobel filter
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/dc/d2b/classcv_1_1cuda_1_1Filter.html#a20b58d13871027473b4c39cc698cf80f
func (sf *SobelFilter) ApplyWithStream(img GpuMat, dst *GpuMat, s Stream) error {
	return OpenCVResult(C.SobelFilter_Apply(C.SobelFilter(sf.p), img.p, dst.p, s.p))
}
