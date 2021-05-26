package contrib

/*
#include <stdlib.h>
#include "xphoto.h"
*/
import "C"
import (
	"gocv.io/x/gocv"
	"unsafe"
)

// GrayworldWB is a wrapper around the cv::xphoto::GrayworldWB.
type GrayworldWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// LearningBasedWB is a wrapper around the cv::xphoto::LearningBasedWB.
type LearningBasedWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// SimpleWB is a wrapper around the cv::xphoto::SimpleWB.
type SimpleWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// TonemapDurand is a wrapper around the cv::xphoto::TonemapDurand.
type TonemapDurand struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// Bm3dSteps is the type for the various BM3D algorithm steps
type Bm3dSteps int

const (
	Bm3dAlgoStepAll Bm3dSteps = 0
	Bm3dAlgoSte1    Bm3dSteps = 1
	Bm3dAlgoSte2    Bm3dSteps = 2
)

type TransformTypes int

const (
	Bm3dTypeHaar TransformTypes = 0
)

type InpaintTypes int

const (
	ShitMap InpaintTypes = 0
	FsrBest InpaintTypes = 1
	FsrFast InpaintTypes = 2
)

// ----------------------- ---------------------------------------
// ----------------------- Bm3dDenoising -------------------------
// ----------------------- ---------------------------------------

func ApplyChannelGains(src gocv.Mat, dst *gocv.Mat, gainB float32, gainG float32, gainR float32) {
	C.Xphoto_ApplyChannelGains(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.float(gainB), C.float(gainG), C.float(gainR))
	return
}

// src = Input 8-bit or 16-bit 1-channel image.
func Bm3dDenoisingStep(src gocv.Mat, dststep1 *gocv.Mat, dststep2 *gocv.Mat) {
	C.Xphoto_Bm3dDenoising_Step(C.Mat(src.Ptr()), C.Mat(dststep1.Ptr()), C.Mat(dststep2.Ptr()))
	return
}

// src = Input 8-bit or 16-bit 1-channel image.
func Bm3dDenoisingStepWithParams(src gocv.Mat, dststep1 *gocv.Mat, dststep2 *gocv.Mat,
	h float32, templateWindowSize int,
	searchWindowSize int, blockMatchingStep1 int,
	blockMatchingStep2 int, groupSize int,
	slidingStep int, beta float32,
	normType gocv.NormType, step Bm3dSteps,
	transformType TransformTypes) {
	C.Xphoto_Bm3dDenoising_Step_WithParams(C.Mat(src.Ptr()), C.Mat(dststep1.Ptr()), C.Mat(dststep2.Ptr()),
		C.float(h), C.int(templateWindowSize),
		C.int(searchWindowSize), C.int(blockMatchingStep1),
		C.int(blockMatchingStep2), C.int(groupSize),
		C.int(slidingStep), C.float(beta),
		C.int(normType), C.int(step),
		C.int(transformType))
	return
}

// src = Input 8-bit or 16-bit 1-channel image.
func Bm3dDenoising(src gocv.Mat, dst *gocv.Mat) {
	C.Xphoto_Bm3dDenoising(C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
}

// src = Input 8-bit or 16-bit 1-channel image.
func Bm3dDenoisingWithParams(src gocv.Mat, dst *gocv.Mat,
	h float32, templateWindowSize int,
	searchWindowSize int, blockMatchingStep1 int,
	blockMatchingStep2 int, groupSize int,
	slidingStep int, beta float32,
	normType gocv.NormType, step Bm3dSteps,
	transformType TransformTypes) {

	C.Xphoto_Bm3dDenoising_WithParams(C.Mat(src.Ptr()), C.Mat(dst.Ptr()),
		C.float(h), C.int(templateWindowSize),
		C.int(searchWindowSize), C.int(blockMatchingStep1),
		C.int(blockMatchingStep2), C.int(groupSize),
		C.int(slidingStep), C.float(beta),
		C.int(normType), C.int(step),
		C.int(transformType))
	return
}

// ----------------------- ---------------------------------------
// ----------------------- GrayworldWB ---------------------------
// ----------------------- ---------------------------------------

// NewGrayworldWBWithParams returns a new Gray-world white balance algorithm.
// of type GrayworldWB with customized parameters. GrayworldWB algorithm scales the values
// of pixels based on a gray-world assumption which states that the average of all
// channels should result in a gray image.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
//
func NewGrayworldWB() GrayworldWB {
	return GrayworldWB{p: unsafe.Pointer(C.GrayworldWB_Create())}
}

// SetSaturationThreshold set a Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) SetSaturationThreshold(saturationThreshold float32) {
	C.GrayworldWB_SetSaturationThreshold((C.GrayworldWB)(b.p), C.float(saturationThreshold))
	return
}

// GetSaturationThreshold return the Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) GetSaturationThreshold() float32 {
	return float32(C.GrayworldWB_GetSaturationThreshold((C.GrayworldWB)(b.p)))
}

// Close GrayworldWB.
func (b *GrayworldWB) Close() error {
	C.GrayworldWB_Close((C.GrayworldWB)(b.p))
	b.p = nil
	return nil
}

// BalanceWhite computes a Gray-world white balance using the current GrayworldWB.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#details
//
func (b *GrayworldWB) BalanceWhite(src gocv.Mat, dst *gocv.Mat) {
	C.GrayworldWB_BalanceWhite((C.GrayworldWB)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}

// ----------------------- ---------------------------------------
// ----------------------- LearningBasedWB -----------------------
// ----------------------- ---------------------------------------

// NewLearningBasedWB returns more sophisticated learning-based
// automatic white balance algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gac8fb5636b27eac575f4a4c9c54dd1c7c
//
func NewLearningBasedWB() LearningBasedWB {
	return LearningBasedWB{p: unsafe.Pointer(C.LearningBasedWB_Create())}
}

// NewLearningBasedWBWithParams returns more sophisticated learning-based
// automatic white balance algorithm.
// A type LearningBasedWB algorithm with path model parameters.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gac8fb5636b27eac575f4a4c9c54dd1c7c
//
func NewLearningBasedWBWithParams(pathmodel string) LearningBasedWB {
	cpath := C.CString(pathmodel)
	defer C.free(unsafe.Pointer(cpath))
	return LearningBasedWB{p: unsafe.Pointer(C.LearningBasedWB_CreateWithParams(cpath))}
}

// Close LearningBasedWB.
func (b *LearningBasedWB) Close() error {
	C.LearningBasedWB_Close((C.LearningBasedWB)(b.p))
	b.p = nil
	return nil
}

// ExtractSimpleFeatures
// Implements the feature extraction part of the algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#aeeaca052262a01d0feed6312ccb9a76e
//
func (b *LearningBasedWB) ExtractSimpleFeatures(src gocv.Mat, dst *gocv.Mat) {
	C.LearningBasedWB_ExtractSimpleFeatures((C.LearningBasedWB)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}

// GetHistBinNum
// Defines the size of one dimension of a three-dimensional RGB histogram that is used internally by the algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#abfe7d3983f8245a7eba0a7f9de40e3e1
//
func (b *LearningBasedWB) GetHistBinNum() int {
	return int(C.LearningBasedWB_GetHistBinNum((C.LearningBasedWB)(b.p)))
}

// GetRangeMaxVal
// Maximum possible value of the input image (e.g. 255 for 8 bit images, 4095 for 12 bit images)
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#a764b51265b5a1bd7bd11ce9d14d6f75f
//
func (b *LearningBasedWB) GetRangeMaxVal() int {
	return int(C.LearningBasedWB_GetRangeMaxVal((C.LearningBasedWB)(b.p)))
}

// GetSaturationThreshold
// Threshold that is used to determine saturated pixels
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#ae7eb310249709c2aef41d6399ebd7660
//
func (b *LearningBasedWB) GetSaturationThreshold() float32 {
	return float32(C.LearningBasedWB_GetSaturationThreshold((C.LearningBasedWB)(b.p)))
}

// SetHistBinNum
// Defines the size of one dimension of a three-dimensional RGB histogram that is used internally by the algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#a3381bd425bc4201133c9669071908e7f
//
func (b *LearningBasedWB) SetHistBinNum(val int) {
	C.LearningBasedWB_SetHistBinNum((C.LearningBasedWB)(b.p), C.int(val))
	return
}

// SetRangeMaxVal
// Maximum possible value of the input image (e.g. 255 for 8 bit images, 4095 for 12 bit images)
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#a3d9395274be8053b2f09e46d11a24a65
//
func (b *LearningBasedWB) SetRangeMaxVal(val int) {
	C.LearningBasedWB_SetRangeMaxVal((C.LearningBasedWB)(b.p), C.int(val))
	return
}

// SetSaturationThreshold
// Threshold that is used to determine saturated pixels
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d3b/classcv_1_1xphoto_1_1LearningBasedWB.html#a9bff5a507d4dffc58e16d85b1d07f35f
//
func (b *LearningBasedWB) SetSaturationThreshold(val float32) {
	C.LearningBasedWB_SetSaturationThreshold((C.LearningBasedWB)(b.p), C.float(val))
	return
}

// BalanceWhite computes a learning-based white balance using the current LearningBasedWB.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d7a/classcv_1_1xphoto_1_1WhiteBalancer.html#ae23838a1a54f101b255bca1a97418aa3
//
func (b *LearningBasedWB) BalanceWhite(src gocv.Mat, dst *gocv.Mat) {
	C.LearningBasedWB_BalanceWhite((C.LearningBasedWB)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}

// ----------------------- ---------------------------------------
// ----------------------- SimpleWB ------------------------------
// ----------------------- ---------------------------------------

// NewSimpleWBWithParams returns more sophisticated learning-based
// automatic white balance algorithm.
// A type SimpleWB algorithm with path model parameters.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga2b48b3b384b5c5ee1b15a2a01c26d5f1
//
func NewSimpleWB() SimpleWB {
	return SimpleWB{p: unsafe.Pointer(C.SimpleWB_Create())}
}

// Close SimpleWB.
func (b *SimpleWB) Close() error {
	C.SimpleWB_Close((C.SimpleWB)(b.p))
	b.p = nil
	return nil
}

// GetInputMax
// Input image range maximum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a45fef780842168cba868212c71ad8318
//
func (b *SimpleWB) GetInputMax() float32 {
	return float32(C.SimpleWB_GetInputMax((C.SimpleWB)(b.p)))
}

// GetInputMin
// Input image range minimum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a69ee7c05e5ca45cac60040371a4a648c
//
func (b *SimpleWB) GetInputMin() float32 {
	return float32(C.SimpleWB_GetInputMin((C.SimpleWB)(b.p)))
}

// GetOutputMax
// Output image range maximum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#af8b051a0b2d3e8ad0aa323c195d966c1
//
func (b *SimpleWB) GetOutputMax() float32 {
	return float32(C.SimpleWB_GetOutputMax((C.SimpleWB)(b.p)))
}

// GetOutputMin
// Output image range minimum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#ada11ef4159bd6354a98d371f5be68b44
//
func (b *SimpleWB) GetOutputMin() float32 {
	return float32(C.SimpleWB_GetOutputMin((C.SimpleWB)(b.p)))
}

// GetP
// Percent of top/bottom values to ignore.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a49d63fd73572fc88c944f5ffbcb085a3
//
func (b *SimpleWB) GetP() float32 {
	return float32(C.SimpleWB_GetP((C.SimpleWB)(b.p)))
}

// SetInputMax
// Input image range maximum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a6b2523a8740b353ef50e136d0399bf3a
//
func (b *SimpleWB) SetInputMax(val float32) {
	C.SimpleWB_SetInputMax((C.SimpleWB)(b.p), C.float(val))
	return
}

// SetInputMin
// Input image range minimum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a1b08a24a8589aae886fbf96ba27691a0
//
func (b *SimpleWB) SetInputMin(val float32) {
	C.SimpleWB_SetInputMin((C.SimpleWB)(b.p), C.float(val))
	return
}

// SetOutputMax
// Output image range maximum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a06962b042d9089366bc6347e608481b6
//
func (b *SimpleWB) SetOutputMax(val float32) {
	C.SimpleWB_SetOutputMax((C.SimpleWB)(b.p), C.float(val))
	return
}

// SetOutputMin
// Output image range minimum value.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a60d1e06b122747d416c4f1563167c740
//
func (b *SimpleWB) SetOutputMin(val float32) {
	C.SimpleWB_SetOutputMin((C.SimpleWB)(b.p), C.float(val))
	return
}

// SetP
// Percent of top/bottom values to ignore.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d8b/classcv_1_1xphoto_1_1SimpleWB.html#a31b6bb5452afdb5a444920013417f018
//
func (b *SimpleWB) SetP(val float32) {
	C.SimpleWB_SetP((C.SimpleWB)(b.p), C.float(val))
	return
}

// BalanceWhite computes a learning-based white balance using the current LearningBasedWB.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d7a/classcv_1_1xphoto_1_1WhiteBalancer.html#ae23838a1a54f101b255bca1a97418aa3
//
func (b *SimpleWB) BalanceWhite(src gocv.Mat, dst *gocv.Mat) {
	C.SimpleWB_BalanceWhite((C.SimpleWB)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}

// ----------------------- ---------------------------------------
// ----------------------- TonemapDurand -------------------------
// ----------------------- ---------------------------------------

// NewTonemapDurand returns more sophisticated learning-based
// automatic white balance algorithm.
// A type TonemapDurand algorithm with path model parameters.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d1/db3/classcv_1_1xphoto_1_1TonemapDurand.html
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga51a091aa54e26b3546316ce2c1df190b
//
func NewTonemapDurand() TonemapDurand {
	return TonemapDurand{p: unsafe.Pointer(C.TonemapDurand_Create())}
}

// NewTonemapDurandWithParams returns more sophisticated learning-based
// automatic white balance algorithm.
// A type TonemapDurand algorithm with path model parameters.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d1/db3/classcv_1_1xphoto_1_1TonemapDurand.html
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga51a091aa54e26b3546316ce2c1df190b
//
func NewTonemapDurandWithParams(gamma float32, contrast float32, saturation float32,
	sigma_color float32, sigma_space float32) TonemapDurand {
	return TonemapDurand{p: unsafe.Pointer(C.TonemapDurand_CreateWithParams(C.float(gamma), C.float(contrast),
		C.float(saturation), C.float(sigma_color), C.float(sigma_space)))}
}

// Close TonemapDurand.
func (b *TonemapDurand) Close() error {
	C.TonemapDurand_Close((C.TonemapDurand)(b.p))
	b.p = nil
	return nil
}

// GetContrast
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga2df693dd285a7e7fd3b4fc8a8a750cce
//
func (b *TonemapDurand) GetContrast() float32 {
	return float32(C.TonemapDurand_GetContrast((C.TonemapDurand)(b.p)))
}

// GetSaturation
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gad8ab5850af6fdb3f6bc51d5c9371bbfe
//
func (b *TonemapDurand) GetSaturation() float32 {
	return float32(C.TonemapDurand_GetSaturation((C.TonemapDurand)(b.p)))
}

// GetSigmaColor
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga31d7588db7e47fb81890cba7ff014edb
//
func (b *TonemapDurand) GetSigmaColor() float32 {
	return float32(C.TonemapDurand_GetSigmaColor((C.TonemapDurand)(b.p)))
}

// GetSigmaSpace
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga078b11014d8a41920f50cab57fce9515
//
func (b *TonemapDurand) GetSigmaSpace() float32 {
	return float32(C.TonemapDurand_GetSigmaSpace((C.TonemapDurand)(b.p)))
}

// SetContrast
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga99eaaa24dc25ba387093e957bfca1cad
//
func (b *TonemapDurand) SetContrast(val float32) {
	C.TonemapDurand_SetContrast((C.TonemapDurand)(b.p), C.float(val))
	return
}

// SetSaturation
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga9c7de9517f95fd046910fc818e256d55
//
func (b *TonemapDurand) SetSaturation(val float32) {
	C.TonemapDurand_SetSaturation((C.TonemapDurand)(b.p), C.float(val))
	return
}

// SetSigmaColor
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga7dec866735ecbae9e05224958d4585fd
//
func (b *TonemapDurand) SetSigmaColor(val float32) {
	C.TonemapDurand_SetSigmaColor((C.TonemapDurand)(b.p), C.float(val))
	return
}

// SetSigmaSpace
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#ga8869068912fae078699ce931bdc17fc4
//
func (b *TonemapDurand) SetSigmaSpace(val float32) {
	C.TonemapDurand_SetSigmaSpace((C.TonemapDurand)(b.p), C.float(val))
	return
}

// GetGamma
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d5e/classcv_1_1Tonemap.html#a147c2b57ed5a5a0566001f4db2ddc0dd
//
func (b *TonemapDurand) GetGamma() float32 {
	return float32(C.TonemapDurand_GetGamma((C.TonemapDurand)(b.p)))
}

// SetGamma
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d5e/classcv_1_1Tonemap.html#ac809d2967f942b038b4bf21c97f8b1b7
//
func (b *TonemapDurand) SetGamma(val float32) {
	C.TonemapDurand_SetGamma((C.TonemapDurand)(b.p), C.float(val))
	return
}

// Process
// Tonemaps image with gocv.MatTypeCV32FC3 type image
// For further details, please see:
// https://docs.opencv.org/master/d8/d5e/classcv_1_1Tonemap.html#aa705c3b7226f7028a5c117dffab60fe4
//
func (b *TonemapDurand) Process(src gocv.Mat, dst *gocv.Mat) {
	C.TonemapDurand_Process((C.TonemapDurand)(b.p), C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
	return
}

// ----------------------- ---------------------------------------
// -------------------------- Inpaint ----------------------------
// ----------------------- ---------------------------------------

// The function implements different single-image inpainting algorithms.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gab4febba6be53e5fddc480b8cedf51eee
//
func Inpaint(src *gocv.Mat, mask *gocv.Mat, dst *gocv.Mat, algorithmType InpaintTypes) {
	C.Inpaint(C.Mat(src.Ptr()), C.Mat(mask.Ptr()), C.Mat(dst.Ptr()), C.int(algorithmType))
}

// oilPainting, See the book for details :
// GerPublished by ard J. Holzmann. Beyond Photography: The Digital Darkroom. Prentice Hall in 1988.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gac050a6e876298cb9713cd2c09db9a027
//
func OilPaintingWithParams(src gocv.Mat, dst gocv.Mat, size int, dynRatio int, code gocv.ColorConversionCode) {
	C.OilPaintingWithParams(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.int(size), C.int(dynRatio), C.int(code))
}

// oilPainting, See the book for details :
// GerPublished by ard J. Holzmann. Beyond Photography: The Digital Darkroom. Prentice Hall in 1988.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html#gac18ef93a7b1e65f703f7dc3b1e8e5235
//
func OilPainting(src gocv.Mat, dst *gocv.Mat, size int, dynRatio int) {
	C.OilPainting(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.int(size), C.int(dynRatio))
}
