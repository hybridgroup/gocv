package contrib

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"gocv.io/x/gocv"
)

const (
	arucoImage6X6_250         = "../images/aruco_6X6_250_6.png"
	arucoImage6X6_250_contour = "../images/aruco_6X6_250_6_contour.png"
	arucoImage6X6_250_1       = "../images/aruco_6X6_250_1.png"
)

func TestArucoDetectorParams(t *testing.T) {

	adaptiveThreshWinSizeMin := 4
	adaptiveThreshWinSizeMax := 22
	adaptiveThreshWinSizeStep := 1
	adaptiveThreshConstant := 1.0
	minMarkerPerimeterRate := 0.2
	maxMarkerPerimeterRate := 0.5
	polygonalApproxAccuracyRate := 1.0
	minCornerDistanceRate := 0.1
	minDistanceToBorder := 0
	minMarkerDistanceRate := 1.0
	cornerRefinementMethod := 1
	cornerRefinementWinSize := 1
	cornerRefinementMaxIterations := 1
	cornerRefinementMinAccuracy := 0.5
	markerBorderBits := 1
	perspectiveRemovePixelPerCell := 1
	perspectiveRemoveIgnoredMarginPerCell := 1.0
	maxErroneousBitsInBorderRate := 0.5
	minOtsuStdDev := .5
	errorCorrectionRate := 0.2
	aprilTagQuadDecimate := float32(0.5)
	aprilTagQuadSigma := float32(1)
	aprilTagMinClusterPixels := 1
	aprilTagMaxNmaxima := 1
	aprilTagCriticalRad := float32(0.2)
	aprilTagMaxLineFitMse := float32(0.2)
	aprilTagMinWhiteBlackDiff := 1
	aprilTagDeglitch := 1
	detectInvertedMarker := false

	params := NewArucoDetectorParameters()
	params.SetAdaptiveThreshWinSizeMin(adaptiveThreshWinSizeMin)
	params.SetAdaptiveThreshWinSizeMax(adaptiveThreshWinSizeMax)
	params.SetAdaptiveThreshWinSizeStep(adaptiveThreshWinSizeStep)
	params.SetAdaptiveThreshConstant(adaptiveThreshConstant)
	params.SetMinMarkerPerimeterRate(minMarkerPerimeterRate)
	params.SetMaxMarkerPerimeterRate(maxMarkerPerimeterRate)
	params.SetPolygonalApproxAccuracyRate(polygonalApproxAccuracyRate)
	params.SetMinCornerDistanceRate(minCornerDistanceRate)
	params.SetMinDistanceToBorder(minDistanceToBorder)
	params.SetMinMarkerDistanceRate(minMarkerDistanceRate)
	params.SetCornerRefinementMethod(cornerRefinementMethod)
	params.SetCornerRefinementWinSize(cornerRefinementWinSize)
	params.SetCornerRefinementMaxIterations(cornerRefinementMaxIterations)
	params.SetCornerRefinementMinAccuracy(cornerRefinementMinAccuracy)
	params.SetMarkerBorderBits(markerBorderBits)
	params.SetPerspectiveRemovePixelPerCell(perspectiveRemovePixelPerCell)
	params.SetPerspectiveRemoveIgnoredMarginPerCell(perspectiveRemoveIgnoredMarginPerCell)
	params.SetMaxErroneousBitsInBorderRate(maxErroneousBitsInBorderRate)
	params.SetMinOtsuStdDev(minOtsuStdDev)
	params.SetErrorCorrectionRate(errorCorrectionRate)
	params.SetAprilTagQuadDecimate(aprilTagQuadDecimate)
	params.SetAprilTagQuadSigma(aprilTagQuadSigma)
	params.SetAprilTagMinClusterPixels(aprilTagMinClusterPixels)
	params.SetAprilTagMaxNmaxima(aprilTagMaxNmaxima)
	params.SetAprilTagCriticalRad(aprilTagCriticalRad)
	params.SetAprilTagMaxLineFitMse(aprilTagMaxLineFitMse)
	params.SetAprilTagMinWhiteBlackDiff(aprilTagMinWhiteBlackDiff)
	params.SetAprilTagDeglitch(aprilTagDeglitch)
	params.SetDetectInvertedMarker(detectInvertedMarker)
	if params.GetAdaptiveThreshWinSizeMin() != adaptiveThreshWinSizeMin {
		t.Error(fmt.Sprintf("AdaptiveThreshWinSizeMin expected %v got %v", adaptiveThreshWinSizeMin, params.GetAdaptiveThreshWinSizeMin()))
	}
	if params.GetAdaptiveThreshWinSizeMax() != adaptiveThreshWinSizeMax {
		t.Error(fmt.Sprintf("AdaptiveThreshWinSizeMax expected %v got %v", adaptiveThreshWinSizeMax, params.GetAdaptiveThreshWinSizeMax()))
	}
	if params.GetAdaptiveThreshWinSizeStep() != adaptiveThreshWinSizeStep {
		t.Error(fmt.Sprintf("AdaptiveThreshWinSizeStep expected %v got %v", adaptiveThreshWinSizeStep, params.GetAdaptiveThreshWinSizeStep()))
	}
	if params.GetAdaptiveThreshConstant() != adaptiveThreshConstant {
		t.Error(fmt.Sprintf("AdaptiveThreshConstant expected %v got %v", adaptiveThreshConstant, params.GetAdaptiveThreshConstant()))
	}
	if params.GetMinMarkerPerimeterRate() != minMarkerPerimeterRate {
		t.Error(fmt.Sprintf("MinMarkerPerimeterRate expected %v got %v", minMarkerPerimeterRate, params.GetMinMarkerPerimeterRate()))
	}
	if params.GetMaxMarkerPerimeterRate() != maxMarkerPerimeterRate {
		t.Error(fmt.Sprintf("MaxMarkerPerimeterRate expected %v got %v", maxMarkerPerimeterRate, params.GetMaxMarkerPerimeterRate()))
	}
	if params.GetPolygonalApproxAccuracyRate() != polygonalApproxAccuracyRate {
		t.Error(fmt.Sprintf("PolygonalApproxAccuracyRate expected %v got %v", polygonalApproxAccuracyRate, params.GetPolygonalApproxAccuracyRate()))
	}
	if params.GetMinCornerDistanceRate() != minCornerDistanceRate {
		t.Error(fmt.Sprintf("MinCornerDistanceRate expected %v got %v", minCornerDistanceRate, params.GetMinCornerDistanceRate()))
	}
	if params.GetMinDistanceToBorder() != minDistanceToBorder {
		t.Error(fmt.Sprintf("MinDistanceToBorder expected %v got %v", minDistanceToBorder, params.GetMinDistanceToBorder()))
	}
	if params.GetMinMarkerDistanceRate() != minMarkerDistanceRate {
		t.Error(fmt.Sprintf("MinMarkerDistanceRate expected %v got %v", minMarkerDistanceRate, params.GetMinMarkerDistanceRate()))
	}
	if params.GetCornerRefinementMethod() != cornerRefinementMethod {
		t.Error(fmt.Sprintf("CornerRefinementMethod expected %v got %v", cornerRefinementMethod, params.GetCornerRefinementMethod()))
	}
	if params.GetCornerRefinementWinSize() != cornerRefinementWinSize {
		t.Error(fmt.Sprintf("CornerRefinementWinSize expected %v got %v", cornerRefinementWinSize, params.GetCornerRefinementWinSize()))
	}
	if params.GetCornerRefinementMaxIterations() != cornerRefinementMaxIterations {
		t.Error(fmt.Sprintf("CornerRefinementMaxIterations expected %v got %v", cornerRefinementMaxIterations, params.GetCornerRefinementMaxIterations()))
	}
	if params.GetCornerRefinementMinAccuracy() != cornerRefinementMinAccuracy {
		t.Error(fmt.Sprintf("CornerRefinementMinAccuracy expected %v got %v", cornerRefinementMinAccuracy, params.GetCornerRefinementMinAccuracy()))
	}
	if params.GetMarkerBorderBits() != markerBorderBits {
		t.Error(fmt.Sprintf("MarkerBorderBits expected %v got %v", markerBorderBits, params.GetMarkerBorderBits()))
	}
	if params.GetPerspectiveRemovePixelPerCell() != perspectiveRemovePixelPerCell {
		t.Error(fmt.Sprintf("PerspectiveRemovePixelPerCell expected %v got %v", perspectiveRemovePixelPerCell, params.GetPerspectiveRemovePixelPerCell()))
	}
	if params.GetPerspectiveRemoveIgnoredMarginPerCell() != perspectiveRemoveIgnoredMarginPerCell {
		t.Error(fmt.Sprintf("PerspectiveRemoveIgnoredMarginPerCell expected %v got %v", perspectiveRemoveIgnoredMarginPerCell, params.GetPerspectiveRemoveIgnoredMarginPerCell()))
	}
	if params.GetMaxErroneousBitsInBorderRate() != maxErroneousBitsInBorderRate {
		t.Error(fmt.Sprintf("MaxErroneousBitsInBorderRate expected %v got %v", maxErroneousBitsInBorderRate, params.GetMaxErroneousBitsInBorderRate()))
	}
	if params.GetMinOtsuStdDev() != minOtsuStdDev {
		t.Error(fmt.Sprintf("MinOtsuStdDev expected %v got %v", minOtsuStdDev, params.GetMinOtsuStdDev()))
	}
	if params.GetErrorCorrectionRate() != errorCorrectionRate {
		t.Error(fmt.Sprintf("ErrorCorrectionRate expected %v got %v", errorCorrectionRate, params.GetErrorCorrectionRate()))
	}
	if params.GetAprilTagQuadDecimate() != aprilTagQuadDecimate {
		t.Error(fmt.Sprintf("AprilTagQuadDecimate expected %v got %v", aprilTagQuadDecimate, params.GetAprilTagQuadDecimate()))
	}
	if params.GetAprilTagQuadSigma() != aprilTagQuadSigma {
		t.Error(fmt.Sprintf("AprilTagQuadSigma expected %v got %v", aprilTagQuadSigma, params.GetAprilTagQuadSigma()))
	}
	if params.GetAprilTagMinClusterPixels() != aprilTagMinClusterPixels {
		t.Error(fmt.Sprintf("AprilTagMinClusterPixels expected %v got %v", aprilTagMinClusterPixels, params.GetAprilTagMinClusterPixels()))
	}
	if params.GetAprilTagMaxNmaxima() != aprilTagMaxNmaxima {
		t.Error(fmt.Sprintf("AprilTagMaxNmaxima expected %v got %v", aprilTagMaxNmaxima, params.GetAprilTagMaxNmaxima()))
	}
	if params.GetAprilTagCriticalRad() != aprilTagCriticalRad {
		t.Error(fmt.Sprintf("AprilTagCriticalRad expected %v got %v", aprilTagCriticalRad, params.GetAprilTagCriticalRad()))
	}
	if params.GetAprilTagMaxLineFitMse() != aprilTagMaxLineFitMse {
		t.Error(fmt.Sprintf("AprilTagMaxLineFitMse expected %v got %v", aprilTagMaxLineFitMse, params.GetAprilTagMaxLineFitMse()))
	}
	if params.GetAprilTagMinWhiteBlackDiff() != aprilTagMinWhiteBlackDiff {
		t.Error(fmt.Sprintf("AprilTagMinWhiteBlackDiff expected %v got %v", aprilTagMinWhiteBlackDiff, params.GetAprilTagMinWhiteBlackDiff()))
	}
	if params.GetAprilTagDeglitch() != aprilTagDeglitch {
		t.Error(fmt.Sprintf("AprilTagDeglitch expected %v got %v", aprilTagDeglitch, params.GetAprilTagDeglitch()))
	}
	if params.GetDetectInvertedMarker() != detectInvertedMarker {
		t.Error(fmt.Sprintf("DetectInvertedMarker expected %v got %v", detectInvertedMarker, params.GetDetectInvertedMarker()))
	}

}

func TestDetectMarkersWithDictID(t *testing.T) {
	path := arucoImage6X6_250
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		t.Error(errors.New("Invalid input"))
	}
	defer img.Close()

	params := NewArucoDetectorParameters()

	_, markerIds, _ := DetectMarkersWithDictID(img, ArucoDict6x6_250, params)
	expected := []int{40, 98, 62, 23, 124, 203}
	if !reflect.DeepEqual(markerIds, expected) {
		t.Error(fmt.Sprintf("Marker id expected %v got %v", expected, markerIds))
	}
}

func TestDetectMarkers(t *testing.T) {
	path := arucoImage6X6_250
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		t.Error(errors.New("Invalid input"))
	}
	defer img.Close()

	params := NewArucoDetectorParameters()
	dict := GetPredefinedDictionary(ArucoDict6x6_250)

	_, markerIds, _ := DetectMarkers(img, dict, params)
	expected := []int{40, 98, 62, 23, 124, 203}
	if !reflect.DeepEqual(markerIds, expected) {
		t.Error(fmt.Sprintf("Marker id expected %v got %v", expected, markerIds))
	}
}

func TestDrawDetectedMarkers(t *testing.T) {

	borderColor := gocv.NewScalar(200, 0, 0, 0)

	img := gocv.IMRead(arucoImage6X6_250, gocv.IMReadColor)
	defer img.Close()
	if img.Empty() {
		t.Error(errors.New("Invalid input"))
	}
	defer img.Close()
	imgExpected := gocv.IMRead(arucoImage6X6_250_contour, gocv.IMReadColor)
	if imgExpected.Empty() {
		t.Error(errors.New("Invalid input"))
	}
	defer imgExpected.Close()

	params := NewArucoDetectorParameters()
	markerCorners, markerIds, _ := DetectMarkersWithDictID(img, ArucoDict6x6_250, params)

	DrawDetectedMarkers(img, markerCorners, markerIds, borderColor)
	diff := gocv.NewMat()
	defer diff.Close()
	gocv.AbsDiff(img, imgExpected, &diff)

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(diff, &gray, gocv.ColorBGRToGray)
	if gocv.CountNonZero(gray) > 0 {
		t.Errorf("expected output to match %s", arucoImage6X6_250_contour)
	}
}

func TestDrawMarker(t *testing.T) {
	imgExpected := gocv.IMRead(arucoImage6X6_250_1, gocv.IMReadGrayScale)
	if imgExpected.Empty() {
		t.Error(fmt.Errorf("Invalid marker image '%s'", arucoImage6X6_250_1))
	}
	defer imgExpected.Close()

	img := gocv.NewMat()
	defer img.Close()
	DrawMarker(ArucoDict6x6_250, 1, 200, img, 1)

	diff := gocv.NewMat()
	defer diff.Close()
	gocv.AbsDiff(img, imgExpected, &diff)

	if gocv.CountNonZero(diff) > 0 {
		t.Errorf("expected output to match %s", arucoImage6X6_250_1)
	}
}

func TestGetPredefinedDictionary(t *testing.T) {
	path := arucoImage6X6_250
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		t.Error(fmt.Errorf("Invalid arucoImage6X6_250 %s", path))
	}
	defer img.Close()

	dict := GetPredefinedDictionary(ArucoDict6x6_250)

	params := NewArucoDetectorParameters()
	_, markerIds, _ := DetectMarkers(img, dict, params)
	expected := []int{40, 98, 62, 23, 124, 203}
	if !reflect.DeepEqual(markerIds, expected) {
		t.Error(fmt.Sprintf("Marker id expected %v got %v", expected, markerIds))
	}
}
