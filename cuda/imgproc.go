package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "core.h"
#include "imgproc.h"
*/
import "C"
import (
	"unsafe"

	"gocv.io/x/gocv"
)

// CannyEdgeDetector
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html
//
type CannyEdgeDetector struct {
	p unsafe.Pointer
}

// NewCannyEdgeDetector returns a new CannyEdgeDetector.
func NewCannyEdgeDetector(lowThresh, highThresh float64) CannyEdgeDetector {
	return CannyEdgeDetector{p: unsafe.Pointer(C.CreateCannyEdgeDetector(C.double(lowThresh), C.double(highThresh)))}
}

// NewCannyEdgeDetectorWithParams returns a new CannyEdgeDetector.
func NewCannyEdgeDetectorWithParams(lowThresh, highThresh float64, appertureSize int, L2gradient bool) CannyEdgeDetector {
	return CannyEdgeDetector{p: unsafe.Pointer(C.CreateCannyEdgeDetectorWithParams(C.double(lowThresh), C.double(highThresh), C.int(appertureSize), C.bool(L2gradient)))}
}

// Close CannyEdgeDetector
func (h *CannyEdgeDetector) Close() error {
	C.CannyEdgeDetector_Close((C.CannyEdgeDetector)(h.p))
	h.p = nil
	return nil
}

// Detect finds edges in an image using the Canny algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a6438cf8453f2dfd6703ceb50056de309
//
func (h *CannyEdgeDetector) Detect(img GpuMat, dst *GpuMat) {
	C.CannyEdgeDetector_Detect(C.CannyEdgeDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds edges in an image using the Canny algorithm
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a6438cf8453f2dfd6703ceb50056de309
//
func (h *CannyEdgeDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.CannyEdgeDetector_Detect(C.CannyEdgeDetector(h.p), img.p, dst.p, s.p)
	return
}

// GetAppertureSize
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a19c2963ff255b0c18387594a704439d3
//
func (h *CannyEdgeDetector) GetAppertureSize() int {
	return int(C.CannyEdgeDetector_GetAppertureSize(C.CannyEdgeDetector(h.p)))
}

// GetHighThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a8366296a57059487dcfd7b30f4a9e3b1
//
func (h *CannyEdgeDetector) GetHighThreshold() float64 {
	return float64(C.CannyEdgeDetector_GetHighThreshold(C.CannyEdgeDetector(h.p)))
}

// GetL2Gradient
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a8fe4ed887c226b12ab44084789b4c6dd
//
func (h *CannyEdgeDetector) GetL2Gradient() bool {
	return bool(C.CannyEdgeDetector_GetL2Gradient(C.CannyEdgeDetector(h.p)))
}

// GetLowThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#aaf5a8944a8ac11093cf7a093b45cd3a8
//
func (h *CannyEdgeDetector) GetLowThreshold() float64 {
	return float64(C.CannyEdgeDetector_GetLowThreshold(C.CannyEdgeDetector(h.p)))
}

// SetAppertureSize
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#aac7d0602338e1a2a783811a929967714
//
func (h *CannyEdgeDetector) SetAppertureSize(appertureSize int) {
	C.CannyEdgeDetector_SetAppertureSize(C.CannyEdgeDetector(h.p), C.int(appertureSize))
}

// SetHighThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a63d352fe7f3bad640e63f4e394619235
//
func (h *CannyEdgeDetector) SetHighThreshold(highThresh float64) {
	C.CannyEdgeDetector_SetHighThreshold(C.CannyEdgeDetector(h.p), C.double(highThresh))
}

// SetL2Gradient
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#ac2e8a675cc30cb3e621ac684e22f89d1
//
func (h *CannyEdgeDetector) SetL2Gradient(L2gradient bool) {
	C.CannyEdgeDetector_SetL2Gradient(C.CannyEdgeDetector(h.p), C.bool(L2gradient))
}

// SetLowThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a6bdc1479c1557288a69c6314c61d1548
//
func (h *CannyEdgeDetector) SetLowThreshold(lowThresh float64) {
	C.CannyEdgeDetector_SetLowThreshold(C.CannyEdgeDetector(h.p), C.double(lowThresh))
}

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga48d0f208181d5ca370d8ff6b62cbe826
//
func CvtColor(src GpuMat, dst *GpuMat, code gocv.ColorConversionCode) {
	C.GpuCvtColor(src.p, dst.p, C.int(code), nil)
}

// CvtColorWithStream converts an image from one color space to another
// using a Stream for concurrency.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga48d0f208181d5ca370d8ff6b62cbe826
//
func CvtColorWithStream(src GpuMat, dst *GpuMat, code gocv.ColorConversionCode, s Stream) {
	C.GpuCvtColor(src.p, dst.p, C.int(code), s.p)
}

// HoughLinesDetector
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html
//
type HoughLinesDetector struct {
	p unsafe.Pointer
}

// NewHoughLinesDetector returns a new HoughLinesDetector.
func NewHoughLinesDetector(rho float32, theta float32, threshold int) HoughLinesDetector {
	return HoughLinesDetector{p: unsafe.Pointer(C.HoughLinesDetector_Create(C.double(rho), C.double(theta), C.int(threshold)))}
}

// NewHoughLinesDetectorWithParams returns a new HoughLinesDetector.
func NewHoughLinesDetectorWithParams(rho float32, theta float32, threshold int, sort bool, maxlines int) HoughLinesDetector {
	return HoughLinesDetector{p: unsafe.Pointer(C.HoughLinesDetector_CreateWithParams(C.double(rho), C.double(theta), C.int(threshold), C.bool(sort), C.int(maxlines)))}
}

// Close HoughLinesDetector
func (h *HoughLinesDetector) Close() error {
	C.HoughLinesDetector_Close((C.HoughLinesDetector)(h.p))
	h.p = nil
	return nil
}

// Detect finds lines in a binary image using the classical Hough transform.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html#a18ff6d0886833ac6215054e191ae2520
//
func (h *HoughLinesDetector) Detect(img GpuMat, dst *GpuMat) {
	C.HoughLinesDetector_Detect(C.HoughLinesDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds lines in a binary image using the classical Hough transform
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html#a18ff6d0886833ac6215054e191ae2520
//
func (h *HoughLinesDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.HoughLinesDetector_Detect(C.HoughLinesDetector(h.p), img.p, dst.p, s.p)
	return
}

// HoughSegmentDetector
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html
//
type HoughSegmentDetector struct {
	p unsafe.Pointer
}

// NewHoughSegmentDetector returns a new HoughSegmentDetector.
func NewHoughSegmentDetector(rho float32, theta float32, minLineLength int, maxLineGap int) HoughSegmentDetector {
	return HoughSegmentDetector{p: unsafe.Pointer(C.HoughSegmentDetector_Create(C.double(rho), C.double(theta), C.int(minLineLength), C.int(maxLineGap)))}
}

// Close HoughSegmentDetector
func (h *HoughSegmentDetector) Close() error {
	C.HoughSegmentDetector_Close((C.HoughSegmentDetector)(h.p))
	h.p = nil
	return nil
}

// Detect finds lines in a binary image using the Hough probabilistic transform.
// For further details, please see:
// https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html#a739bf84825ca455966d69dd75ca0ea6e
//
func (h *HoughSegmentDetector) Detect(img GpuMat, dst *GpuMat) {
	C.HoughSegmentDetector_Detect(C.HoughSegmentDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds lines in a binary image using the Hough probabilistic transform
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html#a739bf84825ca455966d69dd75ca0ea6e
//
func (h *HoughSegmentDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.HoughSegmentDetector_Detect(C.HoughSegmentDetector(h.p), img.p, dst.p, s.p)
	return
}
