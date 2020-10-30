// Package cuda is the GoCV wrapper around OpenCV cuda.
//
// For further details, please see:
// https://github.com/opencv/c
//
// import "gocv.io/x/gocv/cuda"
package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "objdetect.h"
#include "core.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

type DescriptorStorageFormat int

const (
	DESCR_FORMAT_COL_BY_COL DescriptorStorageFormat = 0

	DESCR_FORMAT_COL_BY_ROW DescriptorStorageFormat = 1
)

// CascadeClassifier_GPU is a cascade classifier class for object detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d80/classcv_1_1cuda_1_1CascadeClassifier.html
//
type CascadeClassifier struct {
	p unsafe.Pointer
}

// NewCascadeClassifier_GPU returns a new CascadeClassifier.
func NewCascadeClassifier(name string) CascadeClassifier {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return CascadeClassifier{p: unsafe.Pointer(C.CascadeClassifier_GPU_Create(cName))}
}

// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d80/classcv_1_1cuda_1_1CascadeClassifier.html#a182656b878046eb3f0e9c0f0ee327f08
//
func (c *CascadeClassifier) DetectMultiScale(img GpuMat) []image.Rectangle {
	ret := C.CascadeClassifier_GPU_DetectMultiScale(C.CascadeClassifier_GPU(c.p), img.p)
	defer C.GpuRects_Close(ret)

	return toRectangles(ret)
}

// HOG is a Histogram Of Gradiants (HOG) for object detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOG.html#a723b95b709cfd3f95cf9e616de988fc8
//
type HOG struct {
	p unsafe.Pointer
}

// NewHOG returns a new HOG.
func CreateHOG() HOG {
	return HOG{p: unsafe.Pointer(C.HOG_Create())}
}

func CreateHOGWithParams(winSize, blockSize, blockStride, cellSize image.Point, nbins int) HOG {
	wSz := C.struct_Size{
		width:  C.int(winSize.X),
		height: C.int(winSize.Y),
	}

	bSz := C.struct_Size{
		width:  C.int(blockSize.X),
		height: C.int(blockSize.Y),
	}

	bSt := C.struct_Size{
		width:  C.int(blockStride.X),
		height: C.int(blockStride.Y),
	}

	cSz := C.struct_Size{
		width:  C.int(cellSize.X),
		height: C.int(cellSize.Y),
	}

	return HOG{p: unsafe.Pointer(C.HOG_CreateWithParams(wSz, bSz, bSt, cSz, C.int(nbins)))}
}

// Compute returns block descriptors computed for the whole image.
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#ab4287267081959ec77c01269dbfcd373
//
func (h *HOG) Compute(img GpuMat) GpuMat {
	return newGpuMat(C.HOG_Compute(C.HOG(h.p), img.p))
}

// DetectMultiScale detects objects in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOG.html#a660e5cd036fd5ddf0f5767b352acd948
//
func (h *HOG) DetectMultiScale(img GpuMat) []image.Rectangle {
	ret := C.HOG_DetectMultiScale(C.HOG(h.p), img.p)
	defer C.GpuRects_Close(ret)

	return toRectangles(ret)
}

// GetDefaultPeopleDetector returns a new Mat with the HOG DefaultPeopleDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a016f9ffced8b2f4b20bdd06a775017d1
//
func (h *HOG) GetDefaultPeopleDetector() C.Mat {
	return C.Mat(C.HOG_GetPeopleDetector(C.HOG(h.p)))
}

// SetSVMDetector sets the data for the HOG.
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a5d12db2277b7c3c849d75258eec8d1d4
//
func (h *HOG) SetSVMDetector(det C.Mat) error {
	C.HOG_SetSVMDetector(C.HOG(h.p), det)
	return nil
}

// GetDescriptorFormat
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#adad29ed960a953aa13dc59c410683620
//
func (h *HOG) GetDescriptorFormat() DescriptorStorageFormat {
	return DescriptorStorageFormat(C.HOG_GetDescriptorFormat(C.HOG(h.p)))
}

// GetBlockHistogramSize returns the block histogram size.
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a016f9ffced8b2f4b20bdd06a775017d1
//
func (h *HOG) GetBlockHistogramSize() int {
	return int(C.HOG_GetBlockHistogramSize(C.HOG(h.p)))
}

// GetDescriptorFormat returns the number of coefficients required for the classification.
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#adb8c714cba1a025b8869d5a0e152f824
//
func (h *HOG) GetDescriptorSize() int {
	return int(C.HOG_GetDescriptorSize(C.HOG(h.p)))
}

// GetGammaCorrection
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a7032eed27cf7a004b727a6e522c2404e
//
func (h *HOG) GetGammaCorrection() bool {
	return bool(C.HOG_GetGammaCorrection(C.HOG(h.p)))
}

// GetGroupThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a7032eed27cf7a004b727a6e522c2404e
//
func (h *HOG) GetGroupThreshold() int {
	return int(C.HOG_GetGroupThreshold(C.HOG(h.p)))
}

// GetHitThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#ae0de149980ea47fbd39b7766df565b27
//
func (h *HOG) GetHitThreshold() float64 {
	return float64(C.HOG_GetHitThreshold(C.HOG(h.p)))
}

// GetL2HysThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a6853c9a66889fed996678f7972df9660
//
func (h *HOG) GetL2HysThreshold() float64 {
	return float64(C.HOG_GetL2HysThreshold(C.HOG(h.p)))
}

// GetNumLevels
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a15238eb6f52a1ddeedd015773c46efd8
//
func (h *HOG) GetNumLevels() int {
	return int(C.HOG_GetNumLevels(C.HOG(h.p)))
}

// GetScaleFactor
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a89c59564625bb2c691af8c2cf49aab9e
//
func (h *HOG) GetScaleFactor() float64 {
	return float64(C.HOG_GetScaleFactor(C.HOG(h.p)))
}

// GetWinSigma
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a22d03fa05b251b4f19cfa1fab36e754e
//
func (h *HOG) GetWinSigma() float64 {
	return float64(C.HOG_GetWinSigma(C.HOG(h.p)))
}

// GetWinStride
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a6c63504790b51963ca33496a0b039b48
//
func (h *HOG) GetWinStride() image.Point {
	sz := C.HOG_GetWinStride(C.HOG(h.p))
	return image.Pt(int(sz.width), int(sz.height))
}

// SetDescriptorFormat
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a6e3e1075a567268f2dfb2151b1c99cb6
//
func (h *HOG) SetDescriptorFormat(descrFormat DescriptorStorageFormat) {
	C.HOG_SetDescriptorFormat(C.HOG(h.p), C.int(descrFormat))
}

// SetGammaCorrection
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a0eb2f1ecf59ccc599bffac3a0a55562f
//
func (h *HOG) SetGammaCorrection(gammaCorrection bool) {
	C.HOG_SetGammaCorrection(C.HOG(h.p), C.bool(gammaCorrection))
}

// SetGroupThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#adad9af4e4ed0e0a045a70cd44520eefd
//
func (h *HOG) SetGroupThreshold(groupThreshold int) {
	C.HOG_SetGroupThreshold(C.HOG(h.p), C.int(groupThreshold))
}

// SetHitThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a8b623393c11d18b89fa373269b97aea4
//
func (h *HOG) SetHitThreshold(hitThreshold float64) {
	C.HOG_SetHitThreshold(C.HOG(h.p), C.double(hitThreshold))
}

// SetL2HysThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a30e5c88864fff774f403313993947d62
//
func (h *HOG) SetL2HysThreshold(thresholdL2hys float64) {
	C.HOG_SetL2HysThreshold(C.HOG(h.p), C.double(thresholdL2hys))
}

// SetNumLevels
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a7602088f3e792de196f8f7efcd9bd448
//
func (h *HOG) SetNumLevels(nlevels int) {
	C.HOG_SetNumLevels(C.HOG(h.p), C.int(nlevels))
}

// SetScaleFactor
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a21dc5e3dc6272030694d52e83352b337
//
func (h *HOG) SetScaleFactor(scale0 float64) {
	C.HOG_SetScaleFactor(C.HOG(h.p), C.double(scale0))
}

// SetWinSigma
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#ab291779ff8ac649174b102f64c5f9012
//
func (h *HOG) SetWinSigma(winSigma float64) {
	C.HOG_SetWinSigma(C.HOG(h.p), C.double(winSigma))
}

// SetWinStride
//
// For further details, please see:
// https://docs.opencv.org/master/de/da6/classcv_1_1cuda_1_1HOG.html#a5e74646651209ae13f1b3dd18179773f
//
func (h *HOG) SetWinStride(sz image.Point) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.HOG_SetWinStride(C.HOG(h.p), pSize)
	return
}
