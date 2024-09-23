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
func (h *CannyEdgeDetector) Detect(img GpuMat, dst *GpuMat) {
	C.CannyEdgeDetector_Detect(C.CannyEdgeDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds edges in an image using the Canny algorithm
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a6438cf8453f2dfd6703ceb50056de309
func (h *CannyEdgeDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.CannyEdgeDetector_Detect(C.CannyEdgeDetector(h.p), img.p, dst.p, s.p)
	return
}

// GetAppertureSize
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a19c2963ff255b0c18387594a704439d3
func (h *CannyEdgeDetector) GetAppertureSize() int {
	return int(C.CannyEdgeDetector_GetAppertureSize(C.CannyEdgeDetector(h.p)))
}

// GetHighThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a8366296a57059487dcfd7b30f4a9e3b1
func (h *CannyEdgeDetector) GetHighThreshold() float64 {
	return float64(C.CannyEdgeDetector_GetHighThreshold(C.CannyEdgeDetector(h.p)))
}

// GetL2Gradient
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a8fe4ed887c226b12ab44084789b4c6dd
func (h *CannyEdgeDetector) GetL2Gradient() bool {
	return bool(C.CannyEdgeDetector_GetL2Gradient(C.CannyEdgeDetector(h.p)))
}

// GetLowThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#aaf5a8944a8ac11093cf7a093b45cd3a8
func (h *CannyEdgeDetector) GetLowThreshold() float64 {
	return float64(C.CannyEdgeDetector_GetLowThreshold(C.CannyEdgeDetector(h.p)))
}

// SetAppertureSize
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#aac7d0602338e1a2a783811a929967714
func (h *CannyEdgeDetector) SetAppertureSize(appertureSize int) {
	C.CannyEdgeDetector_SetAppertureSize(C.CannyEdgeDetector(h.p), C.int(appertureSize))
}

// SetHighThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a63d352fe7f3bad640e63f4e394619235
func (h *CannyEdgeDetector) SetHighThreshold(highThresh float64) {
	C.CannyEdgeDetector_SetHighThreshold(C.CannyEdgeDetector(h.p), C.double(highThresh))
}

// SetL2Gradient
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#ac2e8a675cc30cb3e621ac684e22f89d1
func (h *CannyEdgeDetector) SetL2Gradient(L2gradient bool) {
	C.CannyEdgeDetector_SetL2Gradient(C.CannyEdgeDetector(h.p), C.bool(L2gradient))
}

// SetLowThreshold
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html#a6bdc1479c1557288a69c6314c61d1548
func (h *CannyEdgeDetector) SetLowThreshold(lowThresh float64) {
	C.CannyEdgeDetector_SetLowThreshold(C.CannyEdgeDetector(h.p), C.double(lowThresh))
}

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga48d0f208181d5ca370d8ff6b62cbe826
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
func CvtColorWithStream(src GpuMat, dst *GpuMat, code gocv.ColorConversionCode, s Stream) {
	C.GpuCvtColor(src.p, dst.p, C.int(code), s.p)
}

// Demosaicing converts an image from Bayer pattern to RGB or grayscale.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga7fb153572b573ebd2d7610fcbe64166e
func Demosaicing(src GpuMat, dst *GpuMat, code gocv.ColorConversionCode) {
	C.GpuDemosaicing(src.p, dst.p, C.int(code), nil)
}

// DemosaicingWithStream converts an image from Bayer pattern to RGB or grayscale
// using a Stream for concurrency.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga7fb153572b573ebd2d7610fcbe64166e
func DemosaicingWithStream(src GpuMat, dst *GpuMat, code gocv.ColorConversionCode, s Stream) {
	C.GpuDemosaicing(src.p, dst.p, C.int(code), s.p)
}

// HoughLinesDetector
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html
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
func (h *HoughLinesDetector) Detect(img GpuMat, dst *GpuMat) {
	C.HoughLinesDetector_Detect(C.HoughLinesDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds lines in a binary image using the classical Hough transform
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html#a18ff6d0886833ac6215054e191ae2520
func (h *HoughLinesDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.HoughLinesDetector_Detect(C.HoughLinesDetector(h.p), img.p, dst.p, s.p)
	return
}

// HoughSegmentDetector
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html
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
func (h *HoughSegmentDetector) Detect(img GpuMat, dst *GpuMat) {
	C.HoughSegmentDetector_Detect(C.HoughSegmentDetector(h.p), img.p, dst.p, nil)
	return
}

// DetectWithStream finds lines in a binary image using the Hough probabilistic transform
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html#a739bf84825ca455966d69dd75ca0ea6e
func (h *HoughSegmentDetector) DetectWithStream(img GpuMat, dst *GpuMat, s Stream) {
	C.HoughSegmentDetector_Detect(C.HoughSegmentDetector(h.p), img.p, dst.p, s.p)
	return
}

// TemplateMatching
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/d2/d58/classcv_1_1cuda_1_1TemplateMatching.html
type TemplateMatching struct {
	p unsafe.Pointer
}

// NewTemplateMatching returns a new TemplateMatching.
func NewTemplateMatching(srcType gocv.MatType, method gocv.TemplateMatchMode) TemplateMatching {
	return TemplateMatching{p: unsafe.Pointer(C.TemplateMatching_Create(C.int(srcType), C.int(method)))}
}

// Close TemplateMatching
func (tm *TemplateMatching) Close() error {
	C.TemplateMatching_Close((C.TemplateMatching)(tm.p))
	tm.p = nil
	return nil
}

// Match computes a proximity map for a raster template and an image where the template is searched for.
// For further details, please see:
// https://docs.opencv.org/4.6.0/d2/d58/classcv_1_1cuda_1_1TemplateMatching.html#a05a565a53461c916b3b10737cbe43a01
func (tm *TemplateMatching) Match(img GpuMat, tmpl GpuMat, dst *GpuMat) {
	C.TemplateMatching_Match(C.TemplateMatching(tm.p), img.p, tmpl.p, dst.p, nil)
	return
}

// MatchWithStream computes a proximity map for a raster template and an image where the template is searched for
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/4.6.0/d2/d58/classcv_1_1cuda_1_1TemplateMatching.html#a05a565a53461c916b3b10737cbe43a01
func (tm *TemplateMatching) MatchWithStream(img GpuMat, tmpl GpuMat, dst *GpuMat, s Stream) {
	C.TemplateMatching_Match(C.TemplateMatching(tm.p), img.p, tmpl.p, dst.p, s.p)
	return
}

// AlphaComp Composites two images using alpha opacity values contained in each image.
//
// img1: First image. Supports CV_8UC4 , CV_16UC4 , CV_32SC4 and CV_32FC4 types.
//
// img2: Second image. Must have the same size and the same type as img1.
//
// dst: Destination image.
//
// alpha_op: Flag specifying the alpha-blending operation
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#ga08a698700458d9311390997b57fbf8dc
func AlphaComp(img1 GpuMat, img2 GpuMat, dst *GpuMat, alphaOp AlphaCompTypes) {
	C.AlphaComp(img1.p, img2.p, dst.p, C.int(alphaOp), nil)
}

// AlphaCompWithStream Composites two images using alpha opacity values contained in each image.
//
// img1: First image. Supports CV_8UC4 , CV_16UC4 , CV_32SC4 and CV_32FC4 types.
//
// img2: Second image. Must have the same size and the same type as img1.
//
// dst: Destination image.
//
// alpha_op: Flag specifying the alpha-blending operation
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#ga08a698700458d9311390997b57fbf8dc
func AlphaCompWithStream(img1 GpuMat, img2 GpuMat, dst *GpuMat, alphaOp AlphaCompTypes, s Stream) {
	C.AlphaComp(img1.p, img2.p, dst.p, C.int(alphaOp), s.p)
}

// GammaCorrection Routines for correcting image color gamma.
//
// src: Source image (3- or 4-channel 8 bit).
//
// dst: Destination image.
//
// forward: true for forward gamma correction or false for inverse gamma correction.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#gaf4195a8409c3b8fbfa37295c2b2c4729
func GammaCorrection(src GpuMat, dst *GpuMat, forward bool) {
	C.GammaCorrection(src.p, dst.p, C.bool(forward), nil)
}

// GammaCorrectionWithStream Routines for correcting image color gamma.
//
// src: Source image (3- or 4-channel 8 bit).
//
// dst: Destination image.
//
// forward: true for forward gamma correction or false for inverse gamma correction.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#gaf4195a8409c3b8fbfa37295c2b2c4729
func GammaCorrectionWithStream(src GpuMat, dst *GpuMat, forward bool, s Stream) {
	C.GammaCorrection(src.p, dst.p, C.bool(forward), s.p)
}

// SwapChannels Exchanges the color channels of an image in-place.
//
// image: Source image. Supports only CV_8UC4 type.
//
// dstOrder: Integer array describing how channel values are permutated. The n-th entry of the array
// contains the number of the channel that is stored in the n-th channel of the output image.
// E.g. Given an RGBA image, aDstOrder = [3,2,1,0] converts this to ABGR channel order.
//
// supports arbitrary permutations of the original channels, including replication.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#ga75a29cc4a97cde0d43ea066b01de927e
func SwapChannels(image *GpuMat, dstOrder []int) {
	C.SwapChannels(image.p, (*C.int)(unsafe.Pointer(&dstOrder[0])), nil)
}

// SwapChannelsWithStream Exchanges the color channels of an image in-place.
//
// image: Source image. Supports only CV_8UC4 type.
//
// dstOrder: Integer array describing how channel values are permutated. The n-th entry of the array
// contains the number of the channel that is stored in the n-th channel of the output image.
// E.g. Given an RGBA image, aDstOrder = [3,2,1,0] converts this to ABGR channel order.
//
// stream: Stream for the asynchronous version.
//
// supports arbitrary permutations of the original channels, including replication.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d8c/group__cudaimgproc__color.html#ga75a29cc4a97cde0d43ea066b01de927e
func SwapChannelsWithStream(image *GpuMat, dstOrder []int, s Stream) {
	C.SwapChannels(image.p, (*C.int)(unsafe.Pointer(&dstOrder[0])), s.p)
}

// CalcHist Calculates histogram for one channel 8-bit image.
//
// src: Source image with CV_8UC1 type.
//
// hist: Destination histogram with one row, 256 columns, and the CV_32SC1 type.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#gaaf3944106890947020bb4522a7619c26
func CalcHist(src GpuMat, dst *GpuMat) {
	C.Cuda_CalcHist(src.p, dst.p, nil)
}

// CalcHistWithStream Calculates histogram for one channel 8-bit image.
//
// src: Source image with CV_8UC1 type.
//
// hist: Destination histogram with one row, 256 columns, and the CV_32SC1 type.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#gaaf3944106890947020bb4522a7619c26
func CalcHistWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.Cuda_CalcHist(src.p, dst.p, s.p)
}

// CalcHistWithParams Calculates histogram for one channel 8-bit image confined in given mask.
//
// src: Source image with CV_8UC1 type.
//
// mask: A mask image same size as src and of type CV_8UC1.
//
// hist: Destination histogram with one row, 256 columns, and the CV_32SC1 type.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga2d55b444ce776c8bbd3087cc90c47f32
func CalcHistWithParams(src GpuMat, mask GpuMat, dst *GpuMat, s Stream) {
	C.Cuda_CalcHist_WithParams(src.p, mask.p, dst.p, s.p)
}

// EqualizeHist Equalizes the histogram of a grayscale image.
//
// src: Source image with CV_8UC1 type.
//
// dst: Destination image.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga2384be74bd2feba7e6c46815513f0060
func EqualizeHist(src GpuMat, dst *GpuMat) {
	C.Cuda_EqualizeHist(src.p, dst.p, nil)
}

// EqualizeHistWithStream Equalizes the histogram of a grayscale image.
//
// src: Source image with CV_8UC1 type.
//
// dst: Destination image.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga2384be74bd2feba7e6c46815513f0060
func EqualizeHistWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.Cuda_EqualizeHist(src.p, dst.p, s.p)
}

// EvenLevels Computes levels with even distribution.
//
// levels: Destination array. levels has 1 row, nLevels columns, and the CV_32SC1 type.
//
// nLevels: Number of computed levels. nLevels must be at least 2.
//
// lowerLevel: Lower boundary value of the lowest level.
//
// upperLevel: Upper boundary value of the greatest level.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga2f2cbd21dc6d7367a7c4ee1a826f389dFor further details, please see:
func EvenLevels(levels *GpuMat, nLevels int, lowerLevel int, upperLevel int) {
	C.Cuda_EvenLevels(levels.p, C.int(nLevels), C.int(lowerLevel), C.int(upperLevel), nil)
}

// EvenLevelsWithStream Computes levels with even distribution.
//
// levels: Destination array. levels has 1 row, nLevels columns, and the CV_32SC1 type.
//
// nLevels: Number of computed levels. nLevels must be at least 2.
//
// lowerLevel: Lower boundary value of the lowest level.
//
// upperLevel: Upper boundary value of the greatest level.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga2f2cbd21dc6d7367a7c4ee1a826f389dFor further details, please see:
func EvenLevelsWithStream(levels *GpuMat, nLevels int, lowerLevel int, upperLevel int, s Stream) {
	C.Cuda_EvenLevels(levels.p, C.int(nLevels), C.int(lowerLevel), C.int(upperLevel), s.p)
}

// HistEven Calculates a histogram with evenly distributed bins.
//
// src: Source image. CV_8U, CV_16U, or CV_16S depth and 1 or 4 channels are supported.
// For a four-channel image, all channels are processed separately.
//
// hist: Destination histogram with one row, histSize columns, and the CV_32S type.
//
// histSize: Size of the histogram.
//
// lowerLevel: Lower boundary of lowest-level bin.
//
// upperLevel: Upper boundary of highest-level bin.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#gacd3b14279fb77a57a510cb8c89a1856f
func HistEven(src GpuMat, hist *GpuMat, histSize int, lowerLevel int, upperLevel int) {
	C.Cuda_HistEven(src.p, hist.p, C.int(histSize), C.int(lowerLevel), C.int(upperLevel), nil)
}

// HistEvenWithStream Calculates a histogram with evenly distributed bins.
//
// src: Source image. CV_8U, CV_16U, or CV_16S depth and 1 or 4 channels are supported.
// For a four-channel image, all channels are processed separately.
//
// hist: Destination histogram with one row, histSize columns, and the CV_32S type.
//
// histSize: Size of the histogram.
//
// lowerLevel: Lower boundary of lowest-level bin.
//
// upperLevel: Upper boundary of highest-level bin.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#gacd3b14279fb77a57a510cb8c89a1856f
func HistEvenWithStream(src GpuMat, hist *GpuMat, histSize int, lowerLevel int, upperLevel int, s Stream) {
	C.Cuda_HistEven(src.p, hist.p, C.int(histSize), C.int(lowerLevel), C.int(upperLevel), s.p)
}

// HistRange Calculates a histogram with bins determined by the levels array.
//
// src: Source image. CV_8U , CV_16U , or CV_16S depth and 1 or 4 channels are supported.
// For a four-channel image, all channels are processed separately.
//
// hist: Destination histogram with one row, (levels.cols-1) columns, and the CV_32SC1 type.
//
// levels: Number of levels in the histogram.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga87819085c1059186d9cdeacd92cea783
func HistRange(src GpuMat, hist *GpuMat, levels GpuMat) {
	C.Cuda_HistRange(src.p, hist.p, levels.p, nil)
}

// HistRangeWithStream Calculates a histogram with bins determined by the levels array.
//
// src: Source image. CV_8U , CV_16U , or CV_16S depth and 1 or 4 channels are supported.
// For a four-channel image, all channels are processed separately.
//
// hist: Destination histogram with one row, (levels.cols-1) columns, and the CV_32SC1 type.
//
// levels: Number of levels in the histogram.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d8/d0e/group__cudaimgproc__hist.html#ga87819085c1059186d9cdeacd92cea783
func HistRangeWithStream(src GpuMat, hist *GpuMat, levels GpuMat, s Stream) {
	C.Cuda_HistRange(src.p, hist.p, levels.p, s.p)
}

// BilateralFilter Performs bilateral filtering of passed image.
//
// src: Source image. Supports only (channels != 2 && depth() != CV_8S && depth() != CV_32S && depth() != CV_64F).
//
// dst: Destination imagwe.
//
// kernelSize: Kernel window size.
//
// sigmaColor: Filter sigma in the color space.
//
// sigmaSpatial: Filter sigma in the coordinate space.
//
// borderMode: Border type. See borderInterpolate for details.
// BorderReflect101 , BorderReplicate , BorderConstant , BorderReflect and BorderWrap are supported for now.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga6abeaecdd4e7edc0bd1393a04f4f20bd
func BilateralFilter(src GpuMat, dst *GpuMat, kernelSize int, sigmaColor float32, sigmaSpatial float32, borderMode BorderType) {
	C.Cuda_BilateralFilter(src.p, dst.p, C.int(kernelSize), C.float(sigmaColor), C.float(sigmaSpatial), C.int(borderMode), nil)
}

// BilateralFilterWithStream Performs bilateral filtering of passed image.
//
// src: Source image. Supports only (channels != 2 && depth() != CV_8S && depth() != CV_32S && depth() != CV_64F).
//
// dst: Destination imagwe.
//
// kernelSize: Kernel window size.
//
// sigmaColor: Filter sigma in the color space.
//
// sigmaSpatial: Filter sigma in the coordinate space.
//
// borderMode: Border type. See borderInterpolate for details.
// BorderReflect101 , BorderReplicate , BorderConstant , BorderReflect and BorderWrap are supported for now.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga6abeaecdd4e7edc0bd1393a04f4f20bd
func BilateralFilterWithStream(src GpuMat, dst *GpuMat, kernelSize int, sigmaColor float32, sigmaSpatial float32, borderMode BorderType, s Stream) {
	C.Cuda_BilateralFilter(src.p, dst.p, C.int(kernelSize), C.float(sigmaColor), C.float(sigmaSpatial), C.int(borderMode), s.p)
}

// BlendLinear Performs linear blending of two images.
//
// img1: First image. Supports only CV_8U and CV_32F depth.
//
// img2: Second image. Must have the same size and the same type as img1 .
//
// weights1: Weights for first image. Must have tha same size as img1 . Supports only CV_32F type.
//
// weights2: Weights for second image. Must have tha same size as img2 . Supports only CV_32F type.
//
// result: Destination image.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga4793607e5729bcc15b27ea33d9fe335e
func BlendLinear(img1 GpuMat, img2 GpuMat, weights1 GpuMat, weights2 GpuMat, result *GpuMat) {
	C.Cuda_BlendLinear(img1.p, img2.p, weights1.p, weights2.p, result.p, nil)
}

// BlendLinearWithStream Performs linear blending of two images.
//
// img1: First image. Supports only CV_8U and CV_32F depth.
//
// img2: Second image. Must have the same size and the same type as img1 .
//
// weights1: Weights for first image. Must have tha same size as img1 . Supports only CV_32F type.
//
// weights2: Weights for second image. Must have tha same size as img2 . Supports only CV_32F type.
//
// result: Destination image.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga4793607e5729bcc15b27ea33d9fe335e
func BlendLinearWithStream(img1 GpuMat, img2 GpuMat, weights1 GpuMat, weights2 GpuMat, result *GpuMat, s Stream) {
	C.Cuda_BlendLinear(img1.p, img2.p, weights1.p, weights2.p, result.p, s.p)
}

// MeanShiftFiltering Performs mean-shift filtering for each point of the source image.
// It maps each point of the source image into another point.
// As a result, you have a new color and new position of each point.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dst: Destination image containing the color of mapped points. It has the same size and type as src .
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// criteria: Termination criteria. See TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#gae13b3035bc6df0e512d876dbb8c00555
func MeanShiftFiltering(src GpuMat, dst *GpuMat, sp int, sr int, criteria gocv.TermCriteria) {
	C.Cuda_MeanShiftFiltering(src.p, dst.p, C.int(sp), C.int(sr), C.TermCriteria(criteria.Ptr()), nil)
}

// MeanShiftFilteringWithStream Performs mean-shift filtering for each point of the source image.
// It maps each point of the source image into another point.
// As a result, you have a new color and new position of each point.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dst: Destination image containing the color of mapped points. It has the same size and type as src .
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// criteria: Termination criteria. See TermCriteria.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#gae13b3035bc6df0e512d876dbb8c00555
func MeanShiftFilteringWithStream(src GpuMat, dst *GpuMat, sp int, sr int, criteria gocv.TermCriteria, s Stream) {
	C.Cuda_MeanShiftFiltering(src.p, dst.p, C.int(sp), C.int(sr), C.TermCriteria(criteria.Ptr()), s.p)
}

// MeanShiftProc Performs a mean-shift procedure and stores information
// about processed points (their colors and positions) in two images.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dstr: Destination image containing the color of mapped points. The size and type is the same as src .
//
// dstsp: Destination image containing the position of mapped points. The size is the same as src size. The type is CV_16SC2 .
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// criteria: Termination criteria. See TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga6039dc8ecbe2f912bc83fcc9b3bcca39
func MeanShiftProc(src GpuMat, dstr *GpuMat, dstsp *GpuMat, sp int, sr int, criteria gocv.TermCriteria) {
	C.Cuda_MeanShiftProc(src.p, dstr.p, dstsp.p, C.int(sp), C.int(sr), C.TermCriteria(criteria.Ptr()), nil)
}

// MeanShiftProcWithStream Performs a mean-shift procedure and stores information
// about processed points (their colors and positions) in two images.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dstr: Destination image containing the color of mapped points. The size and type is the same as src .
//
// dstsp: Destination image containing the position of mapped points. The size is the same as src size. The type is CV_16SC2 .
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// criteria: Termination criteria. See TermCriteria.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga6039dc8ecbe2f912bc83fcc9b3bcca39
func MeanShiftProcWithStream(src GpuMat, dstr *GpuMat, dstsp *GpuMat, sp int, sr int, criteria gocv.TermCriteria, s Stream) {
	C.Cuda_MeanShiftProc(src.p, dstr.p, dstsp.p, C.int(sp), C.int(sr), C.TermCriteria(criteria.Ptr()), s.p)
}

// MeanShiftSegmentation Performs a mean-shift segmentation of the source image and eliminates small segments.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dst: Segmented image with the same size and type as src.
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// minsize: Minimum segment size. Smaller segments are merged.
//
// criteria: Termination criteria. See TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga70ed80533a448829dc48cf22b1845c16
func MeanShiftSegmentation(src GpuMat, dst *GpuMat, sp int, sr int, minSize int, criteria gocv.TermCriteria) {
	C.Cuda_MeanShiftSegmentation(src.p, dst.p, C.int(sp), C.int(sr), C.int(minSize), C.TermCriteria(criteria.Ptr()), nil)
}

// MeanShiftSegmentationWithStream Performs a mean-shift segmentation of the source image and eliminates small segments.
//
// src: Source image. Only CV_8UC4 images are supported for now.
//
// dst: Segmented image with the same size and type as src.
//
// sp: Spatial window radius.
//
// sr: Color window radius.
//
// minsize: Minimum segment size. Smaller segments are merged.
//
// criteria: Termination criteria. See TermCriteria.
//
// stream: Stream for the asynchronous version.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d0/d05/group__cudaimgproc.html#ga70ed80533a448829dc48cf22b1845c16
func MeanShiftSegmentationWithStream(src GpuMat, dst *GpuMat, sp int, sr int, minSize int, criteria gocv.TermCriteria, s Stream) {
	C.Cuda_MeanShiftSegmentation(src.p, dst.p, C.int(sp), C.int(sr), C.int(minSize), C.TermCriteria(criteria.Ptr()), s.p)
}
