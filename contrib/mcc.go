//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_mcc)

package contrib

/*
#include <stdlib.h>
#include "mcc.h"
#include "../core.h"
*/
import "C"
import (
	"image"
	"unsafe"

	"gocv.io/x/gocv"
)

type TYPECHART int

const (
	MCC24 TYPECHART = iota
	SG140
	VINYL18
)

type MccCCheckerDetector struct {
	p C.MccCCheckerDetector
}

type MccDnnNet struct {
	p C.MccDnnNet
}

// NewMccCCheckerDetector returns a new MccCCheckerDetector.
func NewMccCCheckerDetector() MccCCheckerDetector {
	return MccCCheckerDetector{p: C.MccCCheckerDetector_New()}
}

// Close deletes the MccCCheckerDetector's pointer.
func (md *MccCCheckerDetector) Close() error {
	C.MccCCheckerDetector_Close(md.p)
	md.p = nil
	return nil
}

// Get the best color checker. By the best it means the one detected with the highest confidence.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#a4bb27724447e01b613f5bc8fb9bb3bfc
func (md *MccCCheckerDetector) GetBestColorChecker() MccCChecker {
	res := C.MccCCheckerDetector_GetBestColorChecker(md.p)
	return MccCChecker{p: res}
}

// Get the list of all detected colorcheckers.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#a67bd0b2271e93b550c0d4d058d438cb6
func (md *MccCCheckerDetector) GetListColorChecker() []MccCChecker {
	list := C.MccCCheckerDetector_GetListColorChecker(md.p)
	defer C.MccCCheckerVector_Close(list)
	size := int(C.MccCCheckerVector_Size(list))
	res := make([]MccCChecker, size)
	for i := 0; i < size; i++ {
		res[i] = MccCChecker{p: C.MccCCheckerVector_At(list, C.int(i))}
	}
	return res
}

// Process does basic macbeth chart detection.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#aa07092a6bc9f0a2b75738bc76f9b2d8b
func (md *MccCCheckerDetector) Process(input gocv.Mat, chartType TYPECHART) bool {
	res := C.MccCCheckerDetector_Process(md.p, C.Mat(input.Ptr()), C.int(chartType), C.int(1), C.bool(false))
	return bool(res)
}

// Process does basic macbeth chart detection.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#aa07092a6bc9f0a2b75738bc76f9b2d8b
func (md *MccCCheckerDetector) ProcessWithRegionsOfInterest(input gocv.Mat, chartType TYPECHART, regionsOfInterest []image.Rectangle) bool {
	vec := NewMccRectVector(regionsOfInterest)
	defer vec.Close()
	res := C.MccCCheckerDetector_ProcessWithRegionsOfInterest(md.p, C.Mat(input.Ptr()), C.int(chartType), vec.p, C.int(1), C.bool(false))
	return bool(res)
}

// Process does basic macbeth chart detection.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#aa07092a6bc9f0a2b75738bc76f9b2d8b
func (md *MccCCheckerDetector) ProcessWithParams(input gocv.Mat, chartType TYPECHART, nc int, useNet bool) bool {
	res := C.MccCCheckerDetector_Process(md.p, C.Mat(input.Ptr()), C.int(chartType), C.int(nc), C.bool(useNet))
	return bool(res)
}

// Process does basic macbeth chart detection.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#aa07092a6bc9f0a2b75738bc76f9b2d8b
func (md *MccCCheckerDetector) ProcessWithRegionsOfInterestWithParams(input gocv.Mat, chartType TYPECHART, regionsOfInterest []image.Rectangle, nc int, useNet bool) bool {
	vec := NewMccRectVector(regionsOfInterest)
	defer vec.Close()
	res := C.MccCCheckerDetector_ProcessWithRegionsOfInterest(md.p, C.Mat(input.Ptr()), C.int(chartType), vec.p, C.int(nc), C.bool(useNet))
	return bool(res)
}

// Set the net which will be used to find the approximate bounding boxes for the color charts.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d9/d53/classcv_1_1mcc_1_1CCheckerDetector.html#a1ad73dda58c3a8c6cb822397e8ef76a1
func (md *MccCCheckerDetector) SetNet(net MccDnnNet) bool {
	res := C.MccCCheckerDetector_SetNet(md.p, net.p)
	return bool(res)
}

type MccCChecker struct {
	p C.MccCChecker
}

func (mc *MccCChecker) SetTarget(target TYPECHART) {
	C.MccCChecker_SetTarget(mc.p, C.int(target))
}

func (mc *MccCChecker) GetTarget() TYPECHART {
	return TYPECHART(C.MccCChecker_GetTarget(mc.p))
}

func (mc *MccCChecker) SetBox(box []gocv.Point2f) {
	n := len(box)
	if n == 0 {
		return
	}
	cBox := make([]C.Point2f, n)
	for i := range box {
		cBox[i] = C.Point2f{x: C.float(box[i].X), y: C.float(box[i].Y)}
	}
	C.MccCChecker_SetBox(mc.p, (*C.Point2f)(unsafe.Pointer(&cBox[0])), C.int(n))
}

func (mc *MccCChecker) GetBox() []gocv.Point2f {
	res := C.MccCChecker_GetBox(mc.p)
	defer C.Points2f_Close(res)
	n := int(res.length)
	if n == 0 || res.points == nil {
		return nil
	}
	pts := (*[1 << 28]C.Point2f)(unsafe.Pointer(res.points))[:n:n]
	out := make([]gocv.Point2f, n)
	for i := 0; i < n; i++ {
		out[i] = gocv.Point2f{X: float32(pts[i].x), Y: float32(pts[i].y)}
	}
	return out
}

func (mc *MccCChecker) SetChartsRGB(mat gocv.Mat) {
	C.MccCChecker_SetChartsRGB(mc.p, C.Mat(mat.Ptr()))
}

func (mc *MccCChecker) GetChartsRGB() gocv.Mat {
	cmat := C.MccCChecker_GetChartsRGB(mc.p)
	defer C.Mat_Close(cmat)
	return gocv.NewMatFromCMat(unsafe.Pointer(cmat))
}

func (mc *MccCChecker) SetChartsYCbCr(mat gocv.Mat) {
	C.MccCChecker_SetChartsYCbCr(mc.p, C.Mat(mat.Ptr()))
}

func (mc *MccCChecker) GetChartsYCbCr() gocv.Mat {
	cmat := C.MccCChecker_GetChartsYCbCr(mc.p)
	defer C.Mat_Close(cmat)
	return gocv.NewMatFromCMat(unsafe.Pointer(cmat))
}

func (mc *MccCChecker) SetCost(cost float32) {
	C.MccCChecker_SetCost(mc.p, C.float(cost))
}

func (mc *MccCChecker) GetCost() float32 {
	return float32(C.MccCChecker_GetCost(mc.p))
}

func (mc *MccCChecker) SetCenter(center gocv.Point2f) {
	C.MccCChecker_SetCenter(mc.p, C.Point2f{x: C.float(center.X), y: C.float(center.Y)})
}

func (mc *MccCChecker) GetCenter() gocv.Point2f {
	res := C.MccCChecker_GetCenter(mc.p)
	return gocv.Point2f{X: float32(res.x), Y: float32(res.y)}
}

// Computes and returns the coordinates of the central parts of the charts modules.
//
// This method computes transformation matrix from the checkers's coordinates (cv::mcc::CChecker::getBox()) and
// find by this the coordinates of the central parts of the charts modules. It is used in cv::mcc::CCheckerDraw::draw()
// and in ChartsRGB calculation.
func (mc *MccCChecker) GetColorCharts() []gocv.Point2f {
	res := C.MccCChecker_GetColorCharts(mc.p)
	n := int(res.length)
	if n == 0 || res.points == nil {
		return nil
	}
	pts := (*[1 << 28]C.Point2f)(unsafe.Pointer(res.points))[:n:n]
	out := make([]gocv.Point2f, n)
	for i := 0; i < n; i++ {
		out[i] = gocv.Point2f{X: float32(pts[i].x), Y: float32(pts[i].y)}
	}

	return out
}

type MccDetectorParameters struct {
	p C.MccDetectorParameters
}

// NewMccDetectorParameters returns the default parameters for the MccDetector
func NewMccDetectorParameters() MccDetectorParameters {
	return MccDetectorParameters{p: C.MccDetectorParameters_Create()}
}

func (mp *MccDetectorParameters) SetAdaptiveThreshWinSizeMin(adaptiveThreshWinSizeMin int) {
	C.MccDetectorParameters_SetAdaptiveThreshWinSizeMin(mp.p, C.int(adaptiveThreshWinSizeMin))
}

func (mp *MccDetectorParameters) GetAdaptiveThreshWinSizeMin() int {
	return int(C.MccDetectorParameters_GetAdaptiveThreshWinSizeMin(mp.p))
}

func (mp *MccDetectorParameters) SetAdaptiveThreshWinSizeMax(adaptiveThreshWinSizeMax int) {
	C.MccDetectorParameters_SetAdaptiveThreshWinSizeMax(mp.p, C.int(adaptiveThreshWinSizeMax))
}

func (mp *MccDetectorParameters) GetAdaptiveThreshWinSizeMax() int {
	return int(C.MccDetectorParameters_GetAdaptiveThreshWinSizeMax(mp.p))
}

func (mp *MccDetectorParameters) SetAdaptiveThreshWinSizeStep(adaptiveThreshWinSizeStep int) {
	C.MccDetectorParameters_SetAdaptiveThreshWinSizeStep(mp.p, C.int(adaptiveThreshWinSizeStep))
}

func (mp *MccDetectorParameters) GetAdaptiveThreshWinSizeStep() int {
	return int(C.MccDetectorParameters_GetAdaptiveThreshWinSizeStep(mp.p))
}

func (mp *MccDetectorParameters) SetBorderWidth(borderWidth int) {
	C.MccDetectorParameters_SetBorderWidth(mp.p, C.int(borderWidth))
}

func (mp *MccDetectorParameters) GetBorderWidth() int {
	return int(C.MccDetectorParameters_GetBorderWidth(mp.p))
}

func (mp *MccDetectorParameters) SetMinContourLengthAllowed(minContourLengthAllowed int) {
	C.MccDetectorParameters_SetMinContourLengthAllowed(mp.p, C.int(minContourLengthAllowed))
}

func (mp *MccDetectorParameters) GetMinContourLengthAllowed() int {
	return int(C.MccDetectorParameters_GetMinContourLengthAllowed(mp.p))
}

func (mp *MccDetectorParameters) SetMinContourPointsAllowed(minContourPointsAllowed int) {
	C.MccDetectorParameters_SetMinContourPointsAllowed(mp.p, C.int(minContourPointsAllowed))
}
func (mp *MccDetectorParameters) GetMinContourPointsAllowed() int {
	return int(C.MccDetectorParameters_GetMinContourPointsAllowed(mp.p))
}

func (mp *MccDetectorParameters) SetMinImageSize(minImageSize int) {
	C.MccDetectorParameters_SetMinImageSize(mp.p, C.int(minImageSize))
}
func (mp *MccDetectorParameters) GetMinImageSize() int {
	return int(C.MccDetectorParameters_GetMinImageSize(mp.p))
}
func (mp *MccDetectorParameters) SetMinInterCheckerDistance(minInterCheckerDistance int) {
	C.MccDetectorParameters_SetMinInterCheckerDistance(mp.p, C.int(minInterCheckerDistance))
}
func (mp *MccDetectorParameters) GetMinInterCheckerDistance() int {
	return int(C.MccDetectorParameters_GetMinInterCheckerDistance(mp.p))
}
func (mp *MccDetectorParameters) SetMinInterContourDistance(minInterContourDistance int) {
	C.MccDetectorParameters_SetMinInterContourDistance(mp.p, C.int(minInterContourDistance))
}
func (mp *MccDetectorParameters) GetMinInterContourDistance() int {
	return int(C.MccDetectorParameters_GetMinInterContourDistance(mp.p))
}

func (mp *MccDetectorParameters) SetAdaptiveThreshConstant(adaptiveThreshConstant float64) {
	C.MccDetectorParameters_SetAdaptiveThreshConstant(mp.p, C.double(adaptiveThreshConstant))
}

func (mp *MccDetectorParameters) GetAdaptiveThreshConstant() float64 {
	return float64(C.MccDetectorParameters_GetAdaptiveThreshConstant(mp.p))
}

func (mp *MccDetectorParameters) SetConfidenceThreshold(confidenceThreshold float64) {
	C.MccDetectorParameters_SetConfidenceThreshold(mp.p, C.double(confidenceThreshold))
}
func (mp *MccDetectorParameters) GetConfidenceThreshold() float64 {
	return float64(C.MccDetectorParameters_GetConfidenceThreshold(mp.p))
}
func (mp *MccDetectorParameters) SetFindCandidatesApproxPolyDPEpsMultiplier(findCandidatesApproxPolyDPEpsMultiplier float64) {
	C.MccDetectorParameters_SetFindCandidatesApproxPolyDPEpsMultiplier(mp.p, C.double(findCandidatesApproxPolyDPEpsMultiplier))
}
func (mp *MccDetectorParameters) GetFindCandidatesApproxPolyDPEpsMultiplier() float64 {
	return float64(C.MccDetectorParameters_GetFindCandidatesApproxPolyDPEpsMultiplier(mp.p))
}
func (mp *MccDetectorParameters) SetMinContourSolidity(minContourSolidity float64) {
	C.MccDetectorParameters_SetMinContourSolidity(mp.p, C.double(minContourSolidity))
}
func (mp *MccDetectorParameters) GetMinContourSolidity() float64 {
	return float64(C.MccDetectorParameters_GetMinContourSolidity(mp.p))
}
func (mp *MccDetectorParameters) SetMinContoursArea(minContoursArea float64) {
	C.MccDetectorParameters_SetMinContoursArea(mp.p, C.double(minContoursArea))
}
func (mp *MccDetectorParameters) GetMinContoursArea() float64 {
	return float64(C.MccDetectorParameters_GetMinContoursArea(mp.p))
}
func (mp *MccDetectorParameters) SetMinContoursAreaRate(minContoursAreaRate float64) {
	C.MccDetectorParameters_SetMinContoursAreaRate(mp.p, C.double(minContoursAreaRate))
}
func (mp *MccDetectorParameters) GetMinContoursAreaRate() float64 {
	return float64(C.MccDetectorParameters_GetMinContoursAreaRate(mp.p))
}

func (mp *MccDetectorParameters) SetB0factor(B0factor float32) {
	C.MccDetectorParameters_SetB0factor(mp.p, C.float(B0factor))
}
func (mp *MccDetectorParameters) GetB0factor() float32 {
	return float32(C.MccDetectorParameters_GetB0factor(mp.p))
}
func (mp *MccDetectorParameters) SetMaxError(maxError float32) {
	C.MccDetectorParameters_SetMaxError(mp.p, C.float(maxError))
}
func (mp *MccDetectorParameters) GetMaxError() float32 {
	return float32(C.MccDetectorParameters_GetMaxError(mp.p))
}
func (mp *MccDetectorParameters) SetMinGroupSize(minGroupSize uint) {
	C.MccDetectorParameters_SetMinGroupSize(mp.p, C.uint(minGroupSize))
}
func (mp *MccDetectorParameters) GetMinGroupSize() uint {
	return uint(C.MccDetectorParameters_GetMinGroupSize(mp.p))
}

type MccCCheckerDraw struct {
	p C.MccCCheckerDraw
}

// NewMccCCheckerDraw creates a new CCheckerDraw with the given color and thickness.
func NewMccCCheckerDraw(cc MccCChecker, color gocv.Scalar, thickness int) MccCCheckerDraw {
	p := C.MccCCheckerDraw_Create(cc.p, C.double(color.Val1), C.double(color.Val2), C.double(color.Val3), C.double(color.Val4), C.int(thickness))
	return MccCCheckerDraw{p: p}
}

// Draw draws the checker on the given image.
func (md *MccCCheckerDraw) Draw(img gocv.Mat) {
	C.MccCCheckerDraw_Draw(md.p, C.Mat(img.Ptr()))
}

func (md *MccCCheckerDraw) Close() error {
	C.MccCCheckerDraw_Close(md.p)
	md.p = nil
	return nil
}

type MccRectVector struct {
	p C.MccRectVector
}

func NewMccRectVector(rects []image.Rectangle) MccRectVector {
	vec := C.MccRectVector_New()
	for _, r := range rects {
		C.MccRectVector_PushBack(vec, C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()))
	}
	return MccRectVector{p: vec}
}

func (mv *MccRectVector) Close() error {
	C.MccRectVector_Close(mv.p)
	mv.p = nil
	return nil
}
