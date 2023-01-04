package gocv

/*
#include <stdlib.h>
#include "aruco.h"
#include "core.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type ArucoDetector struct {
	p C.ArucoDetector
}

// NewArucoDetector returns a new ArucoDetector.
func NewArucoDetector() ArucoDetector {
	return ArucoDetector{p: C.ArucoDetector_New()}
}

// NewArucoDetectorWithParams returns a new ArucoDetector.
func NewArucoDetectorWithParams(dictionary ArucoDictionary, params ArucoDetectorParameters) ArucoDetector {
	return ArucoDetector{p: C.ArucoDetector_NewWithParams(dictionary.p, params.p)}
}

// Close deletes the ArucoDetector's pointer.
func (a *ArucoDetector) Close() error {
	C.ArucoDetector_Close(a.p)
	a.p = nil
	return nil
}

// DetectMarkers does basic marker detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d6a/group__aruco.html#gab9159aa69250d8d3642593e508cb6baa
func (a *ArucoDetector) DetectMarkers(input Mat) (
	markerCorners [][]Point2f, markerIds []int, rejectedCandidates [][]Point2f,
) {

	pvsCorners := NewPoints2fVector()
	defer pvsCorners.Close()
	pvsRejected := NewPoints2fVector()
	defer pvsRejected.Close()
	cmarkerIds := C.IntVector{}
	defer C.free(unsafe.Pointer(cmarkerIds.val))

	C.ArucoDetector_DetectMarkers(a.p, C.Mat(input.Ptr()), C.Points2fVector(pvsCorners.P()),
		&cmarkerIds, C.Points2fVector(pvsRejected.P()))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cmarkerIds.val)),
		Len:  int(cmarkerIds.length),
		Cap:  int(cmarkerIds.length),
	}
	pcids := *(*[]C.int)(unsafe.Pointer(h))
	markerIds = []int{}
	for i := 0; i < int(cmarkerIds.length); i++ {
		markerIds = append(markerIds, int(pcids[i]))
	}

	return pvsCorners.ToPoints(), markerIds, pvsRejected.ToPoints()
}

func ArucoDrawDetectedMarkers(img Mat, markerCorners [][]Point2f, markerIds []int, borderColor Scalar) {
	cMarkerIds := make([]C.int, len(markerIds))
	for i, s := range markerIds {
		cMarkerIds[i] = C.int(s)
	}
	markerIdsIntVec := C.IntVector{
		val:    (*C.int)(&cMarkerIds[0]),
		length: C.int(len(cMarkerIds)),
	}
	_markerCorners := NewPoints2fVectorFromPoints(markerCorners)

	cBorderColor := C.struct_Scalar{
		val1: C.double(borderColor.Val1),
		val2: C.double(borderColor.Val2),
		val3: C.double(borderColor.Val3),
		val4: C.double(borderColor.Val4),
	}

	C.ArucoDrawDetectedMarkers(
		C.Mat(img.Ptr()),
		C.Points2fVector(_markerCorners.P()),
		markerIdsIntVec,
		cBorderColor,
	)
}

func ArucoGenerateImageMarker(dictionaryId ArucoDictionaryCode, id int, sidePixels int, img Mat, borderBits int) {
	C.ArucoGenerateImageMarker(C.int(dictionaryId), C.int(id), C.int(sidePixels), C.Mat(img.Ptr()), C.int(borderBits))
}

type ArucoDetectorParameters struct {
	p C.ArucoDetectorParameters
}

// NewArucoDetectorParameters returns the default parameters for the SimpleBobDetector
func NewArucoDetectorParameters() ArucoDetectorParameters {
	return ArucoDetectorParameters{p: C.ArucoDetectorParameters_Create()}
}

func (ap *ArucoDetectorParameters) SetAdaptiveThreshWinSizeMin(adaptiveThreshWinSizeMin int) {
	C.ArucoDetectorParameters_SetAdaptiveThreshWinSizeMin(ap.p, C.int(adaptiveThreshWinSizeMin))
}

func (ap *ArucoDetectorParameters) GetAdaptiveThreshWinSizeMin() int {
	return int(C.ArucoDetectorParameters_GetAdaptiveThreshWinSizeMin(ap.p))
}

func (ap *ArucoDetectorParameters) SetAdaptiveThreshWinSizeMax(adaptiveThreshWinSizeMax int) {
	C.ArucoDetectorParameters_SetAdaptiveThreshWinSizeMax(ap.p, C.int(adaptiveThreshWinSizeMax))
}

func (ap *ArucoDetectorParameters) GetAdaptiveThreshWinSizeMax() int {
	return int(C.ArucoDetectorParameters_GetAdaptiveThreshWinSizeMax(ap.p))
}

func (ap *ArucoDetectorParameters) SetAdaptiveThreshWinSizeStep(adaptiveThreshWinSizeStep int) {
	C.ArucoDetectorParameters_SetAdaptiveThreshWinSizeStep(ap.p, C.int(adaptiveThreshWinSizeStep))
}

func (ap *ArucoDetectorParameters) GetAdaptiveThreshWinSizeStep() int {
	return int(C.ArucoDetectorParameters_GetAdaptiveThreshWinSizeStep(ap.p))
}

func (ap *ArucoDetectorParameters) SetAdaptiveThreshConstant(adaptiveThreshConstant float64) {
	C.ArucoDetectorParameters_SetAdaptiveThreshConstant(ap.p, C.double(adaptiveThreshConstant))
}

func (ap *ArucoDetectorParameters) GetAdaptiveThreshConstant() float64 {
	return float64(C.ArucoDetectorParameters_GetAdaptiveThreshConstant(ap.p))
}

func (ap *ArucoDetectorParameters) SetMinMarkerPerimeterRate(minMarkerPerimeterRate float64) {
	C.ArucoDetectorParameters_SetMinMarkerPerimeterRate(ap.p, C.double(minMarkerPerimeterRate))
}

func (ap *ArucoDetectorParameters) GetMinMarkerPerimeterRate() float64 {
	return float64(C.ArucoDetectorParameters_GetMinMarkerPerimeterRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetMaxMarkerPerimeterRate(maxMarkerPerimeterRate float64) {
	C.ArucoDetectorParameters_SetMaxMarkerPerimeterRate(ap.p, C.double(maxMarkerPerimeterRate))
}

func (ap *ArucoDetectorParameters) GetMaxMarkerPerimeterRate() float64 {
	return float64(C.ArucoDetectorParameters_GetMaxMarkerPerimeterRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetPolygonalApproxAccuracyRate(polygonalApproxAccuracyRate float64) {
	C.ArucoDetectorParameters_SetPolygonalApproxAccuracyRate(ap.p, C.double(polygonalApproxAccuracyRate))
}

func (ap *ArucoDetectorParameters) GetPolygonalApproxAccuracyRate() float64 {
	return float64(C.ArucoDetectorParameters_GetPolygonalApproxAccuracyRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetMinCornerDistanceRate(minCornerDistanceRate float64) {
	C.ArucoDetectorParameters_SetMinCornerDistanceRate(ap.p, C.double(minCornerDistanceRate))
}

func (ap *ArucoDetectorParameters) GetMinCornerDistanceRate() float64 {
	return float64(C.ArucoDetectorParameters_GetMinCornerDistanceRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetMinDistanceToBorder(minDistanceToBorder int) {
	C.ArucoDetectorParameters_SetMinDistanceToBorder(ap.p, C.int(minDistanceToBorder))
}

func (ap *ArucoDetectorParameters) GetMinDistanceToBorder() int {
	return int(C.ArucoDetectorParameters_GetMinDistanceToBorder(ap.p))
}

func (ap *ArucoDetectorParameters) SetMinMarkerDistanceRate(minMarkerDistanceRate float64) {
	C.ArucoDetectorParameters_SetMinMarkerDistanceRate(ap.p, C.double(minMarkerDistanceRate))
}

func (ap *ArucoDetectorParameters) GetMinMarkerDistanceRate() float64 {
	return float64(C.ArucoDetectorParameters_GetMinMarkerDistanceRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetCornerRefinementMethod(cornerRefinementMethod int) {
	C.ArucoDetectorParameters_SetCornerRefinementMethod(ap.p, C.int(cornerRefinementMethod))
}

func (ap *ArucoDetectorParameters) GetCornerRefinementMethod() int {
	return int(C.ArucoDetectorParameters_GetCornerRefinementMethod(ap.p))
}

func (ap *ArucoDetectorParameters) SetCornerRefinementWinSize(cornerRefinementWinSize int) {
	C.ArucoDetectorParameters_SetCornerRefinementWinSize(ap.p, C.int(cornerRefinementWinSize))
}

func (ap *ArucoDetectorParameters) GetCornerRefinementWinSize() int {
	return int(C.ArucoDetectorParameters_GetCornerRefinementWinSize(ap.p))
}

func (ap *ArucoDetectorParameters) SetCornerRefinementMaxIterations(cornerRefinementMaxIterations int) {
	C.ArucoDetectorParameters_SetCornerRefinementMaxIterations(ap.p, C.int(cornerRefinementMaxIterations))
}

func (ap *ArucoDetectorParameters) GetCornerRefinementMaxIterations() int {
	return int(C.ArucoDetectorParameters_GetCornerRefinementMaxIterations(ap.p))
}

func (ap *ArucoDetectorParameters) SetCornerRefinementMinAccuracy(cornerRefinementMinAccuracy float64) {
	C.ArucoDetectorParameters_SetCornerRefinementMinAccuracy(ap.p, C.double(cornerRefinementMinAccuracy))
}

func (ap *ArucoDetectorParameters) GetCornerRefinementMinAccuracy() float64 {
	return float64(C.ArucoDetectorParameters_GetCornerRefinementMinAccuracy(ap.p))
}

func (ap *ArucoDetectorParameters) SetMarkerBorderBits(markerBorderBits int) {
	C.ArucoDetectorParameters_SetMarkerBorderBits(ap.p, C.int(markerBorderBits))
}

func (ap *ArucoDetectorParameters) GetMarkerBorderBits() int {
	return int(C.ArucoDetectorParameters_GetMarkerBorderBits(ap.p))
}

func (ap *ArucoDetectorParameters) SetPerspectiveRemovePixelPerCell(perspectiveRemovePixelPerCell int) {
	C.ArucoDetectorParameters_SetPerspectiveRemovePixelPerCell(ap.p, C.int(perspectiveRemovePixelPerCell))
}

func (ap *ArucoDetectorParameters) GetPerspectiveRemovePixelPerCell() int {
	return int(C.ArucoDetectorParameters_GetPerspectiveRemovePixelPerCell(ap.p))
}

func (ap *ArucoDetectorParameters) SetPerspectiveRemoveIgnoredMarginPerCell(perspectiveRemoveIgnoredMarginPerCell float64) {
	C.ArucoDetectorParameters_SetPerspectiveRemoveIgnoredMarginPerCell(ap.p, C.double(perspectiveRemoveIgnoredMarginPerCell))
}

func (ap *ArucoDetectorParameters) GetPerspectiveRemoveIgnoredMarginPerCell() float64 {
	return float64(C.ArucoDetectorParameters_GetPerspectiveRemoveIgnoredMarginPerCell(ap.p))
}

func (ap *ArucoDetectorParameters) SetMaxErroneousBitsInBorderRate(maxErroneousBitsInBorderRate float64) {
	C.ArucoDetectorParameters_SetMaxErroneousBitsInBorderRate(ap.p, C.double(maxErroneousBitsInBorderRate))
}

func (ap *ArucoDetectorParameters) GetMaxErroneousBitsInBorderRate() float64 {
	return float64(C.ArucoDetectorParameters_GetMaxErroneousBitsInBorderRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetMinOtsuStdDev(minOtsuStdDev float64) {
	C.ArucoDetectorParameters_SetMinOtsuStdDev(ap.p, C.double(minOtsuStdDev))
}

func (ap *ArucoDetectorParameters) GetMinOtsuStdDev() float64 {
	return float64(C.ArucoDetectorParameters_GetMinOtsuStdDev(ap.p))
}

func (ap *ArucoDetectorParameters) SetErrorCorrectionRate(errorCorrectionRate float64) {
	C.ArucoDetectorParameters_SetErrorCorrectionRate(ap.p, C.double(errorCorrectionRate))
}

func (ap *ArucoDetectorParameters) GetErrorCorrectionRate() float64 {
	return float64(C.ArucoDetectorParameters_GetErrorCorrectionRate(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagQuadDecimate(aprilTagQuadDecimate float32) {
	C.ArucoDetectorParameters_SetAprilTagQuadDecimate(ap.p, C.float(aprilTagQuadDecimate))
}

func (ap *ArucoDetectorParameters) GetAprilTagQuadDecimate() float32 {
	return float32(C.ArucoDetectorParameters_GetAprilTagQuadDecimate(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagQuadSigma(aprilTagQuadSigma float32) {
	C.ArucoDetectorParameters_SetAprilTagQuadSigma(ap.p, C.float(aprilTagQuadSigma))
}

func (ap *ArucoDetectorParameters) GetAprilTagQuadSigma() float32 {
	return float32(C.ArucoDetectorParameters_GetAprilTagQuadSigma(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagMinClusterPixels(aprilTagMinClusterPixels int) {
	C.ArucoDetectorParameters_SetAprilTagMinClusterPixels(ap.p, C.int(aprilTagMinClusterPixels))
}

func (ap *ArucoDetectorParameters) GetAprilTagMinClusterPixels() int {
	return int(C.ArucoDetectorParameters_GetAprilTagMinClusterPixels(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagMaxNmaxima(aprilTagMaxNmaxima int) {
	C.ArucoDetectorParameters_SetAprilTagMaxNmaxima(ap.p, C.int(aprilTagMaxNmaxima))
}

func (ap *ArucoDetectorParameters) GetAprilTagMaxNmaxima() int {
	return int(C.ArucoDetectorParameters_GetAprilTagMaxNmaxima(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagCriticalRad(aprilTagCriticalRad float32) {
	C.ArucoDetectorParameters_SetAprilTagCriticalRad(ap.p, C.float(aprilTagCriticalRad))
}

func (ap *ArucoDetectorParameters) GetAprilTagCriticalRad() float32 {
	return float32(C.ArucoDetectorParameters_GetAprilTagCriticalRad(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagMaxLineFitMse(aprilTagMaxLineFitMse float32) {
	C.ArucoDetectorParameters_SetAprilTagMaxLineFitMse(ap.p, C.float(aprilTagMaxLineFitMse))
}

func (ap *ArucoDetectorParameters) GetAprilTagMaxLineFitMse() float32 {
	return float32(C.ArucoDetectorParameters_GetAprilTagMaxLineFitMse(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagMinWhiteBlackDiff(aprilTagMinWhiteBlackDiff int) {
	C.ArucoDetectorParameters_SetAprilTagMinWhiteBlackDiff(ap.p, C.int(aprilTagMinWhiteBlackDiff))
}

func (ap *ArucoDetectorParameters) GetAprilTagMinWhiteBlackDiff() int {
	return int(C.ArucoDetectorParameters_GetAprilTagMinWhiteBlackDiff(ap.p))
}

func (ap *ArucoDetectorParameters) SetAprilTagDeglitch(aprilTagDeglitch int) {
	C.ArucoDetectorParameters_SetAprilTagDeglitch(ap.p, C.int(aprilTagDeglitch))
}

func (ap *ArucoDetectorParameters) GetAprilTagDeglitch() int {
	return int(C.ArucoDetectorParameters_GetAprilTagDeglitch(ap.p))
}

func (ap *ArucoDetectorParameters) SetDetectInvertedMarker(detectInvertedMarker bool) {
	C.ArucoDetectorParameters_SetDetectInvertedMarker(ap.p, C.bool(detectInvertedMarker))
}

func (ap *ArucoDetectorParameters) GetDetectInvertedMarker() bool {
	return bool(C.ArucoDetectorParameters_GetDetectInvertedMarker(ap.p))
}
