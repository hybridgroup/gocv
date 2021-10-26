package contrib

/*
#include <stdlib.h>
#include "aruco.h"
#include "../core.h"
*/
import "C"

import (
	"reflect"
	"unsafe"

	"gocv.io/x/gocv"
)

// DetectMarkers does basic marker detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d6a/group__aruco.html#gab9159aa69250d8d3642593e508cb6baa
//
func DetectMarkersWithDictID(input gocv.Mat, dictionaryId ArucoDictionaryCode, params ArucoDetectorParameters) (
	markerCorners [][]gocv.Point2f, markerIds []int, rejectedCandidates [][]gocv.Point2f,
) {
	pvsCorners := gocv.NewPoints2fVector()
	defer pvsCorners.Close()
	pvsRejected := gocv.NewPoints2fVector()
	defer pvsRejected.Close()
	cmarkerIds := C.IntVector{}
	defer C.free(unsafe.Pointer(cmarkerIds.val))

	C.detectMarkersWithDictId(
		C.Mat(input.Ptr()),
		C.int(dictionaryId),
		C.Points2fVector(pvsCorners.P()),
		&cmarkerIds,
		C.ArucoDetectorParameters(params.p),
		C.Points2fVector(pvsRejected.P()),
	)

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

func DetectMarkers(input gocv.Mat, dictionary ArucoDictionary, params ArucoDetectorParameters) (
	markerCorners [][]gocv.Point2f, markerIds []int, rejectedCandidates [][]gocv.Point2f,
) {

	pvsCorners := gocv.NewPoints2fVector()
	defer pvsCorners.Close()
	pvsRejected := gocv.NewPoints2fVector()
	defer pvsRejected.Close()
	cmarkerIds := C.IntVector{}
	defer C.free(unsafe.Pointer(cmarkerIds.val))

	C.detectMarkers(
		C.Mat(input.Ptr()),
		C.ArucoDictionary(dictionary.p),
		C.Points2fVector(pvsCorners.P()),
		&cmarkerIds,
		C.ArucoDetectorParameters(params.p),
		C.Points2fVector(pvsRejected.P()),
	)

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

func DrawDetectedMarkers(img gocv.Mat, markerCorners [][]gocv.Point2f, markerIds []int, borderColor gocv.Scalar) {
	cMarkerIds := make([]C.int, len(markerIds))
	for i, s := range markerIds {
		cMarkerIds[i] = C.int(s)
	}
	markerIdsIntVec := C.IntVector{
		val:    (*C.int)(&cMarkerIds[0]),
		length: C.int(len(cMarkerIds)),
	}
	_markerCorners := gocv.NewPoints2fVectorFromPoints(markerCorners)

	cBorderColor := C.struct_Scalar{
		val1: C.double(borderColor.Val1),
		val2: C.double(borderColor.Val2),
		val3: C.double(borderColor.Val3),
		val4: C.double(borderColor.Val4),
	}

	C.drawDetectedMarkers(
		C.Mat(img.Ptr()),
		C.Points2fVector(_markerCorners.P()),
		markerIdsIntVec,
		cBorderColor,
	)
}

func DrawMarker(dictionaryId ArucoDictionaryCode, id int, sidePixels int, img gocv.Mat, borderBits int) {

	C.drawMarker(C.int(dictionaryId), C.int(id), C.int(sidePixels), C.Mat(img.Ptr()), C.int(borderBits))
}

func GetPredefinedDictionary(dictionaryId ArucoDictionaryCode) ArucoDictionary {
	var p C.ArucoDictionary = C.getPredefinedDictionary(C.int(dictionaryId))
	return ArucoDictionary{p}
}

type ArucoDetectorParameters struct {
	p C.ArucoDetectorParameters
}

// NewArucoDetectorParameters returns the default parameters for the SimpleBobDetector
func NewArucoDetectorParameters() ArucoDetectorParameters {
	return ArucoDetectorParameters{p: C.ArucoDetectorParameters_Create()}
}

func (p *ArucoDetectorParameters) SetAdaptiveThreshWinSizeMin(adaptiveThreshWinSizeMin int) {
	p.p.adaptiveThreshWinSizeMin = C.int(adaptiveThreshWinSizeMin)
}
func (p *ArucoDetectorParameters) GetAdaptiveThreshWinSizeMin() int {
	return int(p.p.adaptiveThreshWinSizeMin)
}
func (p *ArucoDetectorParameters) SetAdaptiveThreshWinSizeMax(adaptiveThreshWinSizeMax int) {
	p.p.adaptiveThreshWinSizeMax = C.int(adaptiveThreshWinSizeMax)
}
func (p *ArucoDetectorParameters) GetAdaptiveThreshWinSizeMax() int {
	return int(p.p.adaptiveThreshWinSizeMax)
}
func (p *ArucoDetectorParameters) SetAdaptiveThreshWinSizeStep(adaptiveThreshWinSizeStep int) {
	p.p.adaptiveThreshWinSizeStep = C.int(adaptiveThreshWinSizeStep)
}
func (p *ArucoDetectorParameters) GetAdaptiveThreshWinSizeStep() int {
	return int(p.p.adaptiveThreshWinSizeStep)
}
func (p *ArucoDetectorParameters) SetAdaptiveThreshConstant(adaptiveThreshConstant float64) {
	p.p.adaptiveThreshConstant = C.double(adaptiveThreshConstant)
}
func (p *ArucoDetectorParameters) GetAdaptiveThreshConstant() float64 {
	return float64(p.p.adaptiveThreshConstant)
}
func (p *ArucoDetectorParameters) SetMinMarkerPerimeterRate(minMarkerPerimeterRate float64) {
	p.p.minMarkerPerimeterRate = C.double(minMarkerPerimeterRate)
}
func (p *ArucoDetectorParameters) GetMinMarkerPerimeterRate() float64 {
	return float64(p.p.minMarkerPerimeterRate)
}
func (p *ArucoDetectorParameters) SetMaxMarkerPerimeterRate(maxMarkerPerimeterRate float64) {
	p.p.maxMarkerPerimeterRate = C.double(maxMarkerPerimeterRate)
}
func (p *ArucoDetectorParameters) GetMaxMarkerPerimeterRate() float64 {
	return float64(p.p.maxMarkerPerimeterRate)
}
func (p *ArucoDetectorParameters) SetPolygonalApproxAccuracyRate(polygonalApproxAccuracyRate float64) {
	p.p.polygonalApproxAccuracyRate = C.double(polygonalApproxAccuracyRate)
}
func (p *ArucoDetectorParameters) GetPolygonalApproxAccuracyRate() float64 {
	return float64(p.p.polygonalApproxAccuracyRate)
}
func (p *ArucoDetectorParameters) SetMinCornerDistanceRate(minCornerDistanceRate float64) {
	p.p.minCornerDistanceRate = C.double(minCornerDistanceRate)
}
func (p *ArucoDetectorParameters) GetMinCornerDistanceRate() float64 {
	return float64(p.p.minCornerDistanceRate)
}
func (p *ArucoDetectorParameters) SetMinDistanceToBorder(minDistanceToBorder int) {
	p.p.minDistanceToBorder = C.int(minDistanceToBorder)
}
func (p *ArucoDetectorParameters) GetMinDistanceToBorder() int {
	return int(p.p.minDistanceToBorder)
}
func (p *ArucoDetectorParameters) SetMinMarkerDistanceRate(minMarkerDistanceRate float64) {
	p.p.minMarkerDistanceRate = C.double(minMarkerDistanceRate)
}
func (p *ArucoDetectorParameters) GetMinMarkerDistanceRate() float64 {
	return float64(p.p.minMarkerDistanceRate)
}
func (p *ArucoDetectorParameters) SetCornerRefinementMethod(cornerRefinementMethod int) {
	p.p.cornerRefinementMethod = C.int(cornerRefinementMethod)
}
func (p *ArucoDetectorParameters) GetCornerRefinementMethod() int {
	return int(p.p.cornerRefinementMethod)
}
func (p *ArucoDetectorParameters) SetCornerRefinementWinSize(cornerRefinementWinSize int) {
	p.p.cornerRefinementWinSize = C.int(cornerRefinementWinSize)
}
func (p *ArucoDetectorParameters) GetCornerRefinementWinSize() int {
	return int(p.p.cornerRefinementWinSize)
}
func (p *ArucoDetectorParameters) SetCornerRefinementMaxIterations(cornerRefinementMaxIterations int) {
	p.p.cornerRefinementMaxIterations = C.int(cornerRefinementMaxIterations)
}
func (p *ArucoDetectorParameters) GetCornerRefinementMaxIterations() int {
	return int(p.p.cornerRefinementMaxIterations)
}
func (p *ArucoDetectorParameters) SetCornerRefinementMinAccuracy(cornerRefinementMinAccuracy float64) {
	p.p.cornerRefinementMinAccuracy = C.double(cornerRefinementMinAccuracy)
}
func (p *ArucoDetectorParameters) GetCornerRefinementMinAccuracy() float64 {
	return float64(p.p.cornerRefinementMinAccuracy)
}
func (p *ArucoDetectorParameters) SetMarkerBorderBits(markerBorderBits int) {
	p.p.markerBorderBits = C.int(markerBorderBits)
}
func (p *ArucoDetectorParameters) GetMarkerBorderBits() int {
	return int(p.p.markerBorderBits)
}
func (p *ArucoDetectorParameters) SetPerspectiveRemovePixelPerCell(perspectiveRemovePixelPerCell int) {
	p.p.perspectiveRemovePixelPerCell = C.int(perspectiveRemovePixelPerCell)
}
func (p *ArucoDetectorParameters) GetPerspectiveRemovePixelPerCell() int {
	return int(p.p.perspectiveRemovePixelPerCell)
}
func (p *ArucoDetectorParameters) SetPerspectiveRemoveIgnoredMarginPerCell(perspectiveRemoveIgnoredMarginPerCell float64) {
	p.p.perspectiveRemoveIgnoredMarginPerCell = C.double(perspectiveRemoveIgnoredMarginPerCell)
}
func (p *ArucoDetectorParameters) GetPerspectiveRemoveIgnoredMarginPerCell() float64 {
	return float64(p.p.perspectiveRemoveIgnoredMarginPerCell)
}
func (p *ArucoDetectorParameters) SetMaxErroneousBitsInBorderRate(maxErroneousBitsInBorderRate float64) {
	p.p.maxErroneousBitsInBorderRate = C.double(maxErroneousBitsInBorderRate)
}
func (p *ArucoDetectorParameters) GetMaxErroneousBitsInBorderRate() float64 {
	return float64(p.p.maxErroneousBitsInBorderRate)
}
func (p *ArucoDetectorParameters) SetMinOtsuStdDev(minOtsuStdDev float64) {
	p.p.minOtsuStdDev = C.double(minOtsuStdDev)
}
func (p *ArucoDetectorParameters) GetMinOtsuStdDev() float64 {
	return float64(p.p.minOtsuStdDev)
}
func (p *ArucoDetectorParameters) SetErrorCorrectionRate(errorCorrectionRate float64) {
	p.p.errorCorrectionRate = C.double(errorCorrectionRate)
}
func (p *ArucoDetectorParameters) GetErrorCorrectionRate() float64 {
	return float64(p.p.errorCorrectionRate)
}

func (p *ArucoDetectorParameters) SetAprilTagQuadDecimate(aprilTagQuadDecimate float32) {
	p.p.aprilTagQuadDecimate = C.float(aprilTagQuadDecimate)
}
func (p *ArucoDetectorParameters) GetAprilTagQuadDecimate() float32 {
	return float32(p.p.aprilTagQuadDecimate)
}
func (p *ArucoDetectorParameters) SetAprilTagQuadSigma(aprilTagQuadSigma float32) {
	p.p.aprilTagQuadSigma = C.float(aprilTagQuadSigma)
}
func (p *ArucoDetectorParameters) GetAprilTagQuadSigma() float32 {
	return float32(p.p.aprilTagQuadSigma)
}

func (p *ArucoDetectorParameters) SetAprilTagMinClusterPixels(aprilTagMinClusterPixels int) {
	p.p.aprilTagMinClusterPixels = C.int(aprilTagMinClusterPixels)
}
func (p *ArucoDetectorParameters) GetAprilTagMinClusterPixels() int {
	return int(p.p.aprilTagMinClusterPixels)
}
func (p *ArucoDetectorParameters) SetAprilTagMaxNmaxima(aprilTagMaxNmaxima int) {
	p.p.aprilTagMaxNmaxima = C.int(aprilTagMaxNmaxima)
}
func (p *ArucoDetectorParameters) GetAprilTagMaxNmaxima() int {
	return int(p.p.aprilTagMaxNmaxima)
}
func (p *ArucoDetectorParameters) SetAprilTagCriticalRad(aprilTagCriticalRad float32) {
	p.p.aprilTagCriticalRad = C.float(aprilTagCriticalRad)
}
func (p *ArucoDetectorParameters) GetAprilTagCriticalRad() float32 {
	return float32(p.p.aprilTagCriticalRad)
}
func (p *ArucoDetectorParameters) SetAprilTagMaxLineFitMse(aprilTagMaxLineFitMse float32) {
	p.p.aprilTagMaxLineFitMse = C.float(aprilTagMaxLineFitMse)
}
func (p *ArucoDetectorParameters) GetAprilTagMaxLineFitMse() float32 {
	return float32(p.p.aprilTagMaxLineFitMse)
}
func (p *ArucoDetectorParameters) SetAprilTagMinWhiteBlackDiff(aprilTagMinWhiteBlackDiff int) {
	p.p.aprilTagMinWhiteBlackDiff = C.int(aprilTagMinWhiteBlackDiff)
}
func (p *ArucoDetectorParameters) GetAprilTagMinWhiteBlackDiff() int {
	return int(p.p.aprilTagMinWhiteBlackDiff)
}
func (p *ArucoDetectorParameters) SetAprilTagDeglitch(aprilTagDeglitch int) {
	p.p.aprilTagDeglitch = C.int(aprilTagDeglitch)
}
func (p *ArucoDetectorParameters) GetAprilTagDeglitch() int {
	return int(p.p.aprilTagDeglitch)
}

func (p *ArucoDetectorParameters) SetDetectInvertedMarker(detectInvertedMarker bool) {
	p.p.detectInvertedMarker = C.bool(detectInvertedMarker)
}
func (p *ArucoDetectorParameters) GetDetectInvertedMarker() bool {
	return bool(p.p.detectInvertedMarker)
}
