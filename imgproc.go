package gocv

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import (
	"errors"
	"image"
	"image/color"
	"reflect"
	"unsafe"
)

// ArcLength calculates a contour perimeter or a curve length.
//
// For further details, please see:
//
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8d26483c636be6b35c3ec6335798a47c
func ArcLength(curve PointVector, isClosed bool) float64 {
	return float64(C.ArcLength(curve.p, C.bool(isClosed)))
}

// ApproxPolyDP approximates a polygonal curve(s) with the specified precision.
//
// For further details, please see:
//
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga0012a5fdaea70b8a9970165d98722b4c
func ApproxPolyDP(curve PointVector, epsilon float64, closed bool) PointVector {
	return PointVector{p: C.ApproxPolyDP(curve.p, C.double(epsilon), C.bool(closed))}
}

// ConvexHull finds the convex hull of a point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga014b28e56cb8854c0de4a211cb2be656
func ConvexHull(points PointVector, hull *Mat, clockwise bool, returnPoints bool) {
	C.ConvexHull(points.p, hull.p, C.bool(clockwise), C.bool(returnPoints))
}

// ConvexityDefects finds the convexity defects of a contour.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gada4437098113fd8683c932e0567f47ba
func ConvexityDefects(contour PointVector, hull Mat, result *Mat) {
	C.ConvexityDefects(contour.p, hull.p, result.p)
}

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// http://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
func CvtColor(src Mat, dst *Mat, code ColorConversionCode) {
	C.CvtColor(src.p, dst.p, C.int(code))
}

// EqualizeHist normalizes the brightness and increases the contrast of the image.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga7e54091f0c937d49bf84152a16f76d6e
func EqualizeHist(src Mat, dst *Mat) {
	C.EqualizeHist(src.p, dst.p)
}

// CalcHist Calculates a histogram of a set of images
//
// For futher details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga6ca1876785483836f72a77ced8ea759a
func CalcHist(src []Mat, channels []int, mask Mat, hist *Mat, size []int, ranges []float64, acc bool) {
	cMatArray := make([]C.Mat, len(src))
	for i, r := range src {
		cMatArray[i] = r.p
	}

	cMats := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(src)),
	}

	chansInts := []C.int{}
	for _, v := range channels {
		chansInts = append(chansInts, C.int(v))
	}
	chansVector := C.struct_IntVector{}
	chansVector.val = (*C.int)(&chansInts[0])
	chansVector.length = (C.int)(len(chansInts))

	sizeInts := []C.int{}
	for _, v := range size {
		sizeInts = append(sizeInts, C.int(v))
	}
	sizeVector := C.struct_IntVector{}
	sizeVector.val = (*C.int)(&sizeInts[0])
	sizeVector.length = (C.int)(len(sizeInts))

	rangeFloats := []C.float{}
	for _, v := range ranges {
		rangeFloats = append(rangeFloats, C.float(v))
	}
	rangeVector := C.struct_FloatVector{}
	rangeVector.val = (*C.float)(&rangeFloats[0])
	rangeVector.length = (C.int)(len(rangeFloats))

	C.CalcHist(cMats, chansVector, mask.p, hist.p, sizeVector, rangeVector, C.bool(acc))
}

// CalcBackProject calculates the back projection of a histogram.
//
// For futher details, please see:
// https://docs.opencv.org/3.4/d6/dc7/group__imgproc__hist.html#ga3a0af640716b456c3d14af8aee12e3ca
func CalcBackProject(src []Mat, channels []int, hist Mat, backProject *Mat, ranges []float64, uniform bool) {
	cMatArray := make([]C.Mat, len(src))
	for i, r := range src {
		cMatArray[i] = r.p
	}

	cMats := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(src)),
	}

	chansInts := []C.int{}
	for _, v := range channels {
		chansInts = append(chansInts, C.int(v))
	}
	chansVector := C.struct_IntVector{}
	chansVector.val = (*C.int)(&chansInts[0])
	chansVector.length = (C.int)(len(chansInts))

	rangeFloats := []C.float{}
	for _, v := range ranges {
		rangeFloats = append(rangeFloats, C.float(v))
	}
	rangeVector := C.struct_FloatVector{}
	rangeVector.val = (*C.float)(&rangeFloats[0])
	rangeVector.length = (C.int)(len(rangeFloats))

	C.CalcBackProject(cMats, chansVector, hist.p, backProject.p, rangeVector, C.bool(uniform))
}

// HistCompMethod is the method for Histogram comparison
// For more information, see https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga994f53817d621e2e4228fc646342d386
type HistCompMethod int

const (
	// HistCmpCorrel calculates the Correlation
	HistCmpCorrel HistCompMethod = 0

	// HistCmpChiSqr calculates the Chi-Square
	HistCmpChiSqr HistCompMethod = 1

	// HistCmpIntersect calculates the Intersection
	HistCmpIntersect HistCompMethod = 2

	// HistCmpBhattacharya applies the HistCmpBhattacharya by calculating the Bhattacharya distance.
	HistCmpBhattacharya HistCompMethod = 3

	// HistCmpHellinger applies the HistCmpBhattacharya comparison. It is a synonym to HistCmpBhattacharya.
	HistCmpHellinger = HistCmpBhattacharya

	// HistCmpChiSqrAlt applies the Alternative Chi-Square (regularly used for texture comparsion).
	HistCmpChiSqrAlt HistCompMethod = 4

	// HistCmpKlDiv applies the Kullback-Liebler divergence comparison.
	HistCmpKlDiv HistCompMethod = 5
)

// CompareHist Compares two histograms.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#gaf4190090efa5c47cb367cf97a9a519bd
func CompareHist(hist1 Mat, hist2 Mat, method HistCompMethod) float32 {
	return float32(C.CompareHist(hist1.p, hist2.p, C.int(method)))
}

// ClipLine clips the line against the image rectangle.
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf483cb46ad6b049bc35ec67052ef1c2c
func ClipLine(imgSize image.Point, pt1 image.Point, pt2 image.Point) bool {
	pSize := C.struct_Size{
		width:  C.int(imgSize.X),
		height: C.int(imgSize.Y),
	}

	rPt1 := C.struct_Point{
		x: C.int(pt1.X),
		y: C.int(pt1.Y),
	}

	rPt2 := C.struct_Point{
		x: C.int(pt2.X),
		y: C.int(pt2.Y),
	}

	return bool(C.ClipLine(pSize, rPt1, rPt2))
}

// BilateralFilter applies a bilateral filter to an image.
//
// Bilateral filtering is described here:
// http://www.dai.ed.ac.uk/CVonline/LOCAL_COPIES/MANDUCHI1/Bilateral_Filtering.html
//
// BilateralFilter can reduce unwanted noise very well while keeping edges
// fairly sharp. However, it is very slow compared to most filters.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga9d7064d478c95d60003cf839430737ed
func BilateralFilter(src Mat, dst *Mat, diameter int, sigmaColor float64, sigmaSpace float64) {
	C.BilateralFilter(src.p, dst.p, C.int(diameter), C.double(sigmaColor), C.double(sigmaSpace))
}

// Blur blurs an image Mat using a normalized box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga8c45db9afe636703801b0b2e440fce37
func Blur(src Mat, dst *Mat, ksize image.Point) {
	pSize := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	C.Blur(src.p, dst.p, pSize)
}

// BoxFilter blurs an image using the box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad533230ebf2d42509547d514f7d3fbc3
func BoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}
	C.BoxFilter(src.p, dst.p, C.int(depth), pSize)
}

// SqBoxFilter calculates the normalized sum of squares of the pixel values overlapping the filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga045028184a9ef65d7d2579e5c4bff6c0
func SqBoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}
	C.SqBoxFilter(src.p, dst.p, C.int(depth), pSize)
}

// Dilate dilates an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga4ff0f3318642c4f469d0e11f242f3b6c
func Dilate(src Mat, dst *Mat, kernel Mat) {
	C.Dilate(src.p, dst.p, kernel.p)
}

// DilateWithParams dilates an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga4ff0f3318642c4f469d0e11f242f3b6c
func DilateWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType BorderType, borderValue color.RGBA) {
	cAnchor := C.struct_Point{
		x: C.int(anchor.X),
		y: C.int(anchor.Y),
	}

	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}

	C.DilateWithParams(src.p, dst.p, kernel.p, cAnchor, C.int(iterations), C.int(borderType), bv)
}

// DistanceTransformLabelTypes are the types of the DistanceTransform algorithm flag
type DistanceTransformLabelTypes int

const (
	// DistanceLabelCComp assigns the same label to each connected component of zeros in the source image
	// (as well as all the non-zero pixels closest to the connected component).
	DistanceLabelCComp DistanceTransformLabelTypes = 0

	// DistanceLabelPixel assigns its own label to each zero pixel (and all the non-zero pixels closest to it).
	DistanceLabelPixel
)

// DistanceTransformMasks are the marsk sizes for distance transform
type DistanceTransformMasks int

const (
	// DistanceMask3 is a mask of size 3
	DistanceMask3 DistanceTransformMasks = 0

	// DistanceMask5 is a mask of size 3
	DistanceMask5

	// DistanceMaskPrecise is not currently supported
	DistanceMaskPrecise
)

// DistanceTransform Calculates the distance to the closest zero pixel for each pixel of the source image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga8a0b7fdfcb7a13dde018988ba3a43042
func DistanceTransform(src Mat, dst *Mat, labels *Mat, distType DistanceTypes, maskSize DistanceTransformMasks, labelType DistanceTransformLabelTypes) {
	C.DistanceTransform(src.p, dst.p, labels.p, C.int(distType), C.int(maskSize), C.int(labelType))
}

// Erode erodes an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb
func Erode(src Mat, dst *Mat, kernel Mat) {
	C.Erode(src.p, dst.p, kernel.p)
}

// ErodeWithParams erodes an image by using a specific structuring element.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb
func ErodeWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType int) {
	cAnchor := C.struct_Point{
		x: C.int(anchor.X),
		y: C.int(anchor.Y),
	}

	C.ErodeWithParams(src.p, dst.p, kernel.p, cAnchor, C.int(iterations), C.int(borderType))
}

// RetrievalMode is the mode of the contour retrieval algorithm.
type RetrievalMode int

const (
	// RetrievalExternal retrieves only the extreme outer contours.
	// It sets `hierarchy[i][2]=hierarchy[i][3]=-1` for all the contours.
	RetrievalExternal RetrievalMode = 0

	// RetrievalList retrieves all of the contours without establishing
	// any hierarchical relationships.
	RetrievalList RetrievalMode = 1

	// RetrievalCComp retrieves all of the contours and organizes them into
	// a two-level hierarchy. At the top level, there are external boundaries
	// of the components. At the second level, there are boundaries of the holes.
	// If there is another contour inside a hole of a connected component, it
	// is still put at the top level.
	RetrievalCComp RetrievalMode = 2

	// RetrievalTree retrieves all of the contours and reconstructs a full
	// hierarchy of nested contours.
	RetrievalTree RetrievalMode = 3

	// RetrievalFloodfill lacks a description in the original header.
	RetrievalFloodfill RetrievalMode = 4
)

// ContourApproximationMode is the mode of the contour approximation algorithm.
type ContourApproximationMode int

const (
	// ChainApproxNone stores absolutely all the contour points. That is,
	// any 2 subsequent points (x1,y1) and (x2,y2) of the contour will be
	// either horizontal, vertical or diagonal neighbors, that is,
	// max(abs(x1-x2),abs(y2-y1))==1.
	ChainApproxNone ContourApproximationMode = 1

	// ChainApproxSimple compresses horizontal, vertical, and diagonal segments
	// and leaves only their end points.
	// For example, an up-right rectangular contour is encoded with 4 points.
	ChainApproxSimple ContourApproximationMode = 2

	// ChainApproxTC89L1 applies one of the flavors of the Teh-Chin chain
	// approximation algorithms.
	ChainApproxTC89L1 ContourApproximationMode = 3

	// ChainApproxTC89KCOS applies one of the flavors of the Teh-Chin chain
	// approximation algorithms.
	ChainApproxTC89KCOS ContourApproximationMode = 4
)

// BoundingRect calculates the up-right bounding rectangle of a point set.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#gacb413ddce8e48ff3ca61ed7cf626a366
func BoundingRect(contour PointVector) image.Rectangle {
	r := C.BoundingRect(contour.p)
	rect := image.Rect(int(r.x), int(r.y), int(r.x+r.width), int(r.y+r.height))
	return rect
}

// BoxPoints finds the four vertices of a rotated rect. Useful to draw the rotated rectangle.
//
// For further Details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#gaf78d467e024b4d7936cf9397185d2f5c
func BoxPoints(rect RotatedRect, pts *Mat) {
	rPoints := toCPoints(rect.Points)

	rRect := C.struct_Rect{
		x:      C.int(rect.BoundingRect.Min.X),
		y:      C.int(rect.BoundingRect.Min.Y),
		width:  C.int(rect.BoundingRect.Max.X - rect.BoundingRect.Min.X),
		height: C.int(rect.BoundingRect.Max.Y - rect.BoundingRect.Min.Y),
	}

	rCenter := C.struct_Point{
		x: C.int(rect.Center.X),
		y: C.int(rect.Center.Y),
	}

	rSize := C.struct_Size{
		width:  C.int(rect.Width),
		height: C.int(rect.Height),
	}

	r := C.struct_RotatedRect{
		pts:          rPoints,
		boundingRect: rRect,
		center:       rCenter,
		size:         rSize,
		angle:        C.double(rect.Angle),
	}

	C.BoxPoints(r, pts.p)
}

// ContourArea calculates a contour area.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d3/dc0/group__imgproc__shape.html#ga2c759ed9f497d4a618048a2f56dc97f1
func ContourArea(contour PointVector) float64 {
	result := C.ContourArea(contour.p)
	return float64(result)
}

type RotatedRect struct {
	Points       []image.Point
	BoundingRect image.Rectangle
	Center       image.Point
	Width        int
	Height       int
	Angle        float64
}

// toPoints converts C.Contour to []image.Points
func toPoints(points C.Contour) []image.Point {
	pArray := points.points
	pLength := int(points.length)

	pHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(pArray)),
		Len:  pLength,
		Cap:  pLength,
	}
	sPoints := *(*[]C.Point)(unsafe.Pointer(&pHdr))

	points4 := make([]image.Point, pLength)
	for j, pt := range sPoints {
		points4[j] = image.Pt(int(pt.x), int(pt.y))
	}
	return points4
}

// MinAreaRect finds a rotated rectangle of the minimum area enclosing the input 2D point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga3d476a3417130ae5154aea421ca7ead9
func MinAreaRect(points PointVector) RotatedRect {
	result := C.MinAreaRect(points.p)
	defer C.Points_Close(result.pts)

	return RotatedRect{
		Points:       toPoints(result.pts),
		BoundingRect: image.Rect(int(result.boundingRect.x), int(result.boundingRect.y), int(result.boundingRect.x)+int(result.boundingRect.width), int(result.boundingRect.y)+int(result.boundingRect.height)),
		Center:       image.Pt(int(result.center.x), int(result.center.y)),
		Width:        int(result.size.width),
		Height:       int(result.size.height),
		Angle:        float64(result.angle),
	}
}

// FitEllipse Fits an ellipse around a set of 2D points.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf259efaad93098103d6c27b9e4900ffa
func FitEllipse(pts PointVector) RotatedRect {
	cRect := C.FitEllipse(pts.p)
	defer C.Points_Close(cRect.pts)

	return RotatedRect{
		Points:       toPoints(cRect.pts),
		BoundingRect: image.Rect(int(cRect.boundingRect.x), int(cRect.boundingRect.y), int(cRect.boundingRect.x)+int(cRect.boundingRect.width), int(cRect.boundingRect.y)+int(cRect.boundingRect.height)),
		Center:       image.Pt(int(cRect.center.x), int(cRect.center.y)),
		Width:        int(cRect.size.width),
		Height:       int(cRect.size.height),
		Angle:        float64(cRect.angle),
	}

}

// MinEnclosingCircle finds a circle of the minimum area enclosing the input 2D point set.
//
// For further details, please see:
// https://docs.opencv.org/3.4/d3/dc0/group__imgproc__shape.html#ga8ce13c24081bbc7151e9326f412190f1
func MinEnclosingCircle(pts PointVector) (x, y, radius float32) {
	cCenterPoint := C.struct_Point2f{}
	var cRadius C.float
	C.MinEnclosingCircle(pts.p, &cCenterPoint, &cRadius)
	x, y = float32(cCenterPoint.x), float32(cCenterPoint.y)
	radius = float32(cRadius)
	return x, y, radius
}

// FindContours finds contours in a binary image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga95f5b48d01abc7c2e0732db24689837b
func FindContours(src Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
	hierarchy := NewMat()
	defer hierarchy.Close()
	return FindContoursWithParams(src, &hierarchy, mode, method)
}

// FindContoursWithParams finds contours in a binary image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga17ed9f5d79ae97bd4c7cf18403e1689a
func FindContoursWithParams(src Mat, hierarchy *Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
	return PointsVector{p: C.FindContours(src.p, hierarchy.p, C.int(mode), C.int(method))}
}

// PointPolygonTest performs a point-in-contour test.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1a539e8db2135af2566103705d7a5722
func PointPolygonTest(pts PointVector, pt image.Point, measureDist bool) float64 {
	cp := C.struct_Point{
		x: C.int(pt.X),
		y: C.int(pt.Y),
	}
	return float64(C.PointPolygonTest(pts.p, cp, C.bool(measureDist)))
}

// ConnectedComponentsAlgorithmType specifies the type for ConnectedComponents
type ConnectedComponentsAlgorithmType int

const (
	// SAUF algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity.
	CCL_WU ConnectedComponentsAlgorithmType = 0

	// BBDT algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity.
	CCL_DEFAULT ConnectedComponentsAlgorithmType = 1

	// BBDT algorithm for 8-way connectivity, SAUF algorithm for 4-way connectivity
	CCL_GRANA ConnectedComponentsAlgorithmType = 2
)

// ConnectedComponents computes the connected components labeled image of boolean image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaedef8c7340499ca391d459122e51bef5
func ConnectedComponents(src Mat, labels *Mat) int {
	return int(C.ConnectedComponents(src.p, labels.p, C.int(8), C.int(MatTypeCV32S), C.int(CCL_DEFAULT)))
}

// ConnectedComponents computes the connected components labeled image of boolean image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaedef8c7340499ca391d459122e51bef5
func ConnectedComponentsWithParams(src Mat, labels *Mat, conn int, ltype MatType,
	ccltype ConnectedComponentsAlgorithmType) int {
	return int(C.ConnectedComponents(src.p, labels.p, C.int(conn), C.int(ltype), C.int(ccltype)))
}

// ConnectedComponentsTypes are the connected components algorithm output formats
type ConnectedComponentsTypes int

const (
	//The leftmost (x) coordinate which is the inclusive start of the bounding box in the horizontal direction.
	CC_STAT_LEFT ConnectedComponentsTypes = 0

	//The topmost (y) coordinate which is the inclusive start of the bounding box in the vertical direction.
	CC_STAT_TOP ConnectedComponentsTypes = 1

	// The horizontal size of the bounding box.
	CC_STAT_WIDTH ConnectedComponentsTypes = 2

	// The vertical size of the bounding box.
	CC_STAT_HEIGHT ConnectedComponentsTypes = 3

	// The total area (in pixels) of the connected component.
	CC_STAT_AREA ConnectedComponentsTypes = 4

	CC_STAT_MAX ConnectedComponentsTypes = 5
)

// ConnectedComponentsWithStats computes the connected components labeled image of boolean
// image and also produces a statistics output for each label.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga107a78bf7cd25dec05fb4dfc5c9e765f
func ConnectedComponentsWithStats(src Mat, labels *Mat, stats *Mat, centroids *Mat) int {
	return int(C.ConnectedComponentsWithStats(src.p, labels.p, stats.p, centroids.p,
		C.int(8), C.int(MatTypeCV32S), C.int(CCL_DEFAULT)))
}

// ConnectedComponentsWithStats computes the connected components labeled image of boolean
// image and also produces a statistics output for each label.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga107a78bf7cd25dec05fb4dfc5c9e765f
func ConnectedComponentsWithStatsWithParams(src Mat, labels *Mat, stats *Mat, centroids *Mat,
	conn int, ltype MatType, ccltype ConnectedComponentsAlgorithmType) int {
	return int(C.ConnectedComponentsWithStats(src.p, labels.p, stats.p, centroids.p, C.int(conn),
		C.int(ltype), C.int(ccltype)))
}

// TemplateMatchMode is the type of the template matching operation.
type TemplateMatchMode int

const (
	// TmSqdiff maps to TM_SQDIFF
	TmSqdiff TemplateMatchMode = 0
	// TmSqdiffNormed maps to TM_SQDIFF_NORMED
	TmSqdiffNormed TemplateMatchMode = 1
	// TmCcorr maps to TM_CCORR
	TmCcorr TemplateMatchMode = 2
	// TmCcorrNormed maps to TM_CCORR_NORMED
	TmCcorrNormed TemplateMatchMode = 3
	// TmCcoeff maps to TM_CCOEFF
	TmCcoeff TemplateMatchMode = 4
	// TmCcoeffNormed maps to TM_CCOEFF_NORMED
	TmCcoeffNormed TemplateMatchMode = 5
)

// MatchTemplate compares a template against overlapped image regions.
//
// For further details, please see:
// https://docs.opencv.org/master/df/dfb/group__imgproc__object.html#ga586ebfb0a7fb604b35a23d85391329be
func MatchTemplate(image Mat, templ Mat, result *Mat, method TemplateMatchMode, mask Mat) {
	C.MatchTemplate(image.p, templ.p, result.p, C.int(method), mask.p)
}

// Moments calculates all of the moments up to the third order of a polygon
// or rasterized shape.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga556a180f43cab22649c23ada36a8a139
func Moments(src Mat, binaryImage bool) map[string]float64 {
	r := C.Moments(src.p, C.bool(binaryImage))

	result := make(map[string]float64)
	result["m00"] = float64(r.m00)
	result["m10"] = float64(r.m10)
	result["m01"] = float64(r.m01)
	result["m20"] = float64(r.m20)
	result["m11"] = float64(r.m11)
	result["m02"] = float64(r.m02)
	result["m30"] = float64(r.m30)
	result["m21"] = float64(r.m21)
	result["m12"] = float64(r.m12)
	result["m03"] = float64(r.m03)
	result["mu20"] = float64(r.mu20)
	result["mu11"] = float64(r.mu11)
	result["mu02"] = float64(r.mu02)
	result["mu30"] = float64(r.mu30)
	result["mu21"] = float64(r.mu21)
	result["mu12"] = float64(r.mu12)
	result["mu03"] = float64(r.mu03)
	result["nu20"] = float64(r.nu20)
	result["nu11"] = float64(r.nu11)
	result["nu02"] = float64(r.nu02)
	result["nu30"] = float64(r.nu30)
	result["nu21"] = float64(r.nu21)
	result["nu12"] = float64(r.nu12)
	result["nu03"] = float64(r.nu03)

	return result
}

// PyrDown blurs an image and downsamples it.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaf9bba239dfca11654cb7f50f889fc2ff
func PyrDown(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}
	C.PyrDown(src.p, dst.p, pSize, C.int(borderType))
}

// PyrUp upsamples an image and then blurs it.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gada75b59bdaaca411ed6fee10085eb784
func PyrUp(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}
	C.PyrUp(src.p, dst.p, pSize, C.int(borderType))
}

// MorphologyDefaultBorder returns "magic" border value for erosion and dilation.
// It is automatically transformed to Scalar::all(-DBL_MAX) for dilation.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga94756fad83d9d24d29c9bf478558c40a
func MorphologyDefaultBorderValue() Scalar {
	var scalar C.Scalar = C.MorphologyDefaultBorderValue()
	return NewScalar(float64(scalar.val1), float64(scalar.val2), float64(scalar.val3), float64(scalar.val4))
}

// MorphologyEx performs advanced morphological transformations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f
func MorphologyEx(src Mat, dst *Mat, op MorphType, kernel Mat) {
	C.MorphologyEx(src.p, dst.p, C.int(op), kernel.p)
}

// MorphologyExWithParams performs advanced morphological transformations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f
func MorphologyExWithParams(src Mat, dst *Mat, op MorphType, kernel Mat, iterations int, borderType BorderType) {
	pt := C.struct_Point{
		x: C.int(-1),
		y: C.int(-1),
	}
	C.MorphologyExWithParams(src.p, dst.p, C.int(op), kernel.p, pt, C.int(iterations), C.int(borderType))
}

// MorphShape is the shape of the structuring element used for Morphing operations.
type MorphShape int

const (
	// MorphRect is the rectangular morph shape.
	MorphRect MorphShape = 0

	// MorphCross is the cross morph shape.
	MorphCross MorphShape = 1

	// MorphEllipse is the ellipse morph shape.
	MorphEllipse MorphShape = 2
)

// GetStructuringElement returns a structuring element of the specified size
// and shape for morphological operations.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac342a1bb6eabf6f55c803b09268e36dc
func GetStructuringElement(shape MorphShape, ksize image.Point) Mat {
	sz := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	return newMat(C.GetStructuringElement(C.int(shape), sz))
}

// MorphType type of morphological operation.
type MorphType int

const (
	// MorphErode operation
	MorphErode MorphType = 0

	// MorphDilate operation
	MorphDilate MorphType = 1

	// MorphOpen operation
	MorphOpen MorphType = 2

	// MorphClose operation
	MorphClose MorphType = 3

	// MorphGradient operation
	MorphGradient MorphType = 4

	// MorphTophat operation
	MorphTophat MorphType = 5

	// MorphBlackhat operation
	MorphBlackhat MorphType = 6

	// MorphHitmiss operation
	MorphHitmiss MorphType = 7
)

// BorderType type of border.
type BorderType int

const (
	// BorderConstant border type
	BorderConstant BorderType = 0

	// BorderReplicate border type
	BorderReplicate BorderType = 1

	// BorderReflect border type
	BorderReflect BorderType = 2

	// BorderWrap border type
	BorderWrap BorderType = 3

	// BorderReflect101 border type
	BorderReflect101 BorderType = 4

	// BorderTransparent border type
	BorderTransparent BorderType = 5

	// BorderDefault border type
	BorderDefault = BorderReflect101

	// BorderIsolated border type
	BorderIsolated BorderType = 16
)

// GaussianBlur blurs an image Mat using a Gaussian filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
func GaussianBlur(src Mat, dst *Mat, ksize image.Point, sigmaX float64,
	sigmaY float64, borderType BorderType) {
	pSize := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	C.GaussianBlur(src.p, dst.p, pSize, C.double(sigmaX), C.double(sigmaY), C.int(borderType))
}

// GetGaussianKernel returns Gaussian filter coefficients.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa
func GetGaussianKernel(ksize int, sigma float64) Mat {
	return newMat(C.GetGaussianKernel(C.int(ksize), C.double(sigma), C.int(MatTypeCV64F)))
}

// GetGaussianKernelWithParams returns Gaussian filter coefficients.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa
func GetGaussianKernelWithParams(ksize int, sigma float64, ktype MatType) Mat {
	return newMat(C.GetGaussianKernel(C.int(ksize), C.double(sigma), C.int(ktype)))
}

// Sobel calculates the first, second, third, or mixed image derivatives using an extended Sobel operator
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gacea54f142e81b6758cb6f375ce782c8d
func Sobel(src Mat, dst *Mat, ddepth MatType, dx, dy, ksize int, scale, delta float64, borderType BorderType) {
	C.Sobel(src.p, dst.p, C.int(ddepth), C.int(dx), C.int(dy), C.int(ksize), C.double(scale), C.double(delta), C.int(borderType))
}

// SpatialGradient calculates the first order image derivative in both x and y using a Sobel operator.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga405d03b20c782b65a4daf54d233239a2
func SpatialGradient(src Mat, dx, dy *Mat, ksize MatType, borderType BorderType) {
	C.SpatialGradient(src.p, dx.p, dy.p, C.int(ksize), C.int(borderType))
}

// Laplacian calculates the Laplacian of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad78703e4c8fe703d479c1860d76429e6
func Laplacian(src Mat, dst *Mat, dDepth MatType, size int, scale float64,
	delta float64, borderType BorderType) {
	C.Laplacian(src.p, dst.p, C.int(dDepth), C.int(size), C.double(scale), C.double(delta), C.int(borderType))
}

// Scharr calculates the first x- or y- image derivative using Scharr operator.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaa13106761eedf14798f37aa2d60404c9
func Scharr(src Mat, dst *Mat, dDepth MatType, dx int, dy int, scale float64,
	delta float64, borderType BorderType) {
	C.Scharr(src.p, dst.p, C.int(dDepth), C.int(dx), C.int(dy), C.double(scale), C.double(delta), C.int(borderType))
}

// MedianBlur blurs an image using the median filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga564869aa33e58769b4469101aac458f9
func MedianBlur(src Mat, dst *Mat, ksize int) {
	C.MedianBlur(src.p, dst.p, C.int(ksize))
}

// Canny finds edges in an image using the Canny algorithm.
// The function finds edges in the input image image and marks
// them in the output map edges using the Canny algorithm.
// The smallest value between threshold1 and threshold2 is used
// for edge linking. The largest value is used to
// find initial segments of strong edges.
// See http://en.wikipedia.org/wiki/Canny_edge_detector
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga04723e007ed888ddf11d9ba04e2232de
func Canny(src Mat, edges *Mat, t1 float32, t2 float32) {
	C.Canny(src.p, edges.p, C.double(t1), C.double(t2))
}

// CornerSubPix Refines the corner locations. The function iterates to find
// the sub-pixel accurate location of corners or radial saddle points.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga354e0d7c86d0d9da75de9b9701a9a87e
func CornerSubPix(img Mat, corners *Mat, winSize image.Point, zeroZone image.Point, criteria TermCriteria) {
	winSz := C.struct_Size{
		width:  C.int(winSize.X),
		height: C.int(winSize.Y),
	}

	zeroSz := C.struct_Size{
		width:  C.int(zeroZone.X),
		height: C.int(zeroZone.Y),
	}

	C.CornerSubPix(img.p, corners.p, winSz, zeroSz, criteria.p)
	return
}

// GoodFeaturesToTrack determines strong corners on an image. The function
// finds the most prominent corners in the image or in the specified image region.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga1d6bb77486c8f92d79c8793ad995d541
func GoodFeaturesToTrack(img Mat, corners *Mat, maxCorners int, quality float64, minDist float64) {
	C.GoodFeaturesToTrack(img.p, corners.p, C.int(maxCorners), C.double(quality), C.double(minDist))
}

// GrabCutMode is the flag for GrabCut algorithm.
type GrabCutMode int

const (
	// GCInitWithRect makes the function initialize the state and the mask using the provided rectangle.
	// After that it runs the itercount iterations of the algorithm.
	GCInitWithRect GrabCutMode = 0
	// GCInitWithMask makes the function initialize the state using the provided mask.
	// GCInitWithMask and GCInitWithRect can be combined.
	// Then all the pixels outside of the ROI are automatically initialized with GC_BGD.
	GCInitWithMask GrabCutMode = 1
	// GCEval means that the algorithm should just resume.
	GCEval GrabCutMode = 2
	// GCEvalFreezeModel means that the algorithm should just run a single iteration of the GrabCut algorithm
	// with the fixed model
	GCEvalFreezeModel GrabCutMode = 3
)

// Grabcut runs the GrabCut algorithm.
// The function implements the GrabCut image segmentation algorithm.
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga909c1dda50efcbeaa3ce126be862b37f
func GrabCut(img Mat, mask *Mat, r image.Rectangle, bgdModel *Mat, fgdModel *Mat, iterCount int, mode GrabCutMode) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	C.GrabCut(img.p, mask.p, cRect, bgdModel.p, fgdModel.p, C.int(iterCount), C.int(mode))
}

// HoughMode is the type for Hough transform variants.
type HoughMode int

const (
	// HoughStandard is the classical or standard Hough transform.
	HoughStandard HoughMode = 0
	// HoughProbabilistic is the probabilistic Hough transform (more efficient
	// in case if the picture contains a few long linear segments).
	HoughProbabilistic HoughMode = 1
	// HoughMultiScale is the multi-scale variant of the classical Hough
	// transform.
	HoughMultiScale HoughMode = 2
	// HoughGradient is basically 21HT, described in: HK Yuen, John Princen,
	// John Illingworth, and Josef Kittler. Comparative study of hough
	// transform methods for circle finding. Image and Vision Computing,
	// 8(1):71–77, 1990.
	HoughGradient HoughMode = 3
)

// HoughCircles finds circles in a grayscale image using the Hough transform.
// The only "method" currently supported is HoughGradient. If you want to pass
// more parameters, please see `HoughCirclesWithParams`.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga47849c3be0d0406ad3ca45db65a25d2d
func HoughCircles(src Mat, circles *Mat, method HoughMode, dp, minDist float64) {
	C.HoughCircles(src.p, circles.p, C.int(method), C.double(dp), C.double(minDist))
}

// HoughCirclesWithParams finds circles in a grayscale image using the Hough
// transform. The only "method" currently supported is HoughGradient.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga47849c3be0d0406ad3ca45db65a25d2d
func HoughCirclesWithParams(src Mat, circles *Mat, method HoughMode, dp, minDist, param1, param2 float64, minRadius, maxRadius int) {
	C.HoughCirclesWithParams(src.p, circles.p, C.int(method), C.double(dp), C.double(minDist), C.double(param1), C.double(param2), C.int(minRadius), C.int(maxRadius))
}

// HoughLines implements the standard or standard multi-scale Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga46b4e588934f6c8dfd509cc6e0e4545a
func HoughLines(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
	C.HoughLines(src.p, lines.p, C.double(rho), C.double(theta), C.int(threshold))
}

// HoughLinesP implements the probabilistic Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga8618180a5948286384e3b7ca02f6feeb
func HoughLinesP(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
	C.HoughLinesP(src.p, lines.p, C.double(rho), C.double(theta), C.int(threshold))
}
func HoughLinesPWithParams(src Mat, lines *Mat, rho float32, theta float32, threshold int, minLineLength float32, maxLineGap float32) {
	C.HoughLinesPWithParams(src.p, lines.p, C.double(rho), C.double(theta), C.int(threshold), C.double(minLineLength), C.double(maxLineGap))
}

// HoughLinesPointSet implements the Hough transform algorithm for line
// detection on a set of points. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga2858ef61b4e47d1919facac2152a160e
func HoughLinesPointSet(points Mat, lines *Mat, linesMax int, threshold int,
	minRho float32, maxRho float32, rhoStep float32,
	minTheta float32, maxTheta float32, thetaStep float32) {
	C.HoughLinesPointSet(points.p, lines.p, C.int(linesMax), C.int(threshold),
		C.double(minRho), C.double(maxRho), C.double(rhoStep),
		C.double(minTheta), C.double(maxTheta), C.double(thetaStep))
}

// Integral calculates one or more integral images for the source image.
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga97b87bec26908237e8ba0f6e96d23e28
func Integral(src Mat, sum *Mat, sqsum *Mat, tilted *Mat) {
	C.Integral(src.p, sum.p, sqsum.p, tilted.p)
}

// ThresholdType type of threshold operation.
type ThresholdType int

const (
	// ThresholdBinary threshold type
	ThresholdBinary ThresholdType = 0

	// ThresholdBinaryInv threshold type
	ThresholdBinaryInv ThresholdType = 1

	// ThresholdTrunc threshold type
	ThresholdTrunc ThresholdType = 2

	// ThresholdToZero threshold type
	ThresholdToZero ThresholdType = 3

	// ThresholdToZeroInv threshold type
	ThresholdToZeroInv ThresholdType = 4

	// ThresholdMask threshold type
	ThresholdMask ThresholdType = 7

	// ThresholdOtsu threshold type
	ThresholdOtsu ThresholdType = 8

	// ThresholdTriangle threshold type
	ThresholdTriangle ThresholdType = 16
)

// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#gae8a4a146d1ca78c626a53577199e9c57
func Threshold(src Mat, dst *Mat, thresh float32, maxvalue float32, typ ThresholdType) (threshold float32) {
	return float32(C.Threshold(src.p, dst.p, C.double(thresh), C.double(maxvalue), C.int(typ)))
}

// AdaptiveThresholdType type of adaptive threshold operation.
type AdaptiveThresholdType int

const (
	// AdaptiveThresholdMean threshold type
	AdaptiveThresholdMean AdaptiveThresholdType = 0

	// AdaptiveThresholdGaussian threshold type
	AdaptiveThresholdGaussian AdaptiveThresholdType = 1
)

// AdaptiveThreshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga72b913f352e4a1b1b397736707afcde3
func AdaptiveThreshold(src Mat, dst *Mat, maxValue float32, adaptiveTyp AdaptiveThresholdType, typ ThresholdType, blockSize int, c float32) {
	C.AdaptiveThreshold(src.p, dst.p, C.double(maxValue), C.int(adaptiveTyp), C.int(typ), C.int(blockSize), C.double(c))
}

// ArrowedLine draws a arrow segment pointing from the first point
// to the second one.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga0a165a3ca093fd488ac709fdf10c05b2
func ArrowedLine(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	sp1 := C.struct_Point{
		x: C.int(pt1.X),
		y: C.int(pt1.Y),
	}

	sp2 := C.struct_Point{
		x: C.int(pt2.X),
		y: C.int(pt2.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.ArrowedLine(img.p, sp1, sp2, sColor, C.int(thickness))
}

// Circle draws a circle.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670
func Circle(img *Mat, center image.Point, radius int, c color.RGBA, thickness int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Circle(img.p, pc, C.int(radius), sColor, C.int(thickness))
}

// CircleWithParams draws a circle.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670
func CircleWithParams(img *Mat, center image.Point, radius int, c color.RGBA, thickness int, lineType LineType, shift int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.CircleWithParams(img.p, pc, C.int(radius), sColor, C.int(thickness), C.int(lineType), C.int(shift))
}

// Ellipse draws a simple or thick elliptic arc or fills an ellipse sector.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga28b2267d35786f5f890ca167236cbc69
func Ellipse(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	pa := C.struct_Point{
		x: C.int(axes.X),
		y: C.int(axes.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Ellipse(img.p, pc, pa, C.double(angle), C.double(startAngle), C.double(endAngle), sColor, C.int(thickness))
}

// Ellipse draws a simple or thick elliptic arc or fills an ellipse sector.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga28b2267d35786f5f890ca167236cbc69
func EllipseWithParams(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int, lineType LineType, shift int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	pa := C.struct_Point{
		x: C.int(axes.X),
		y: C.int(axes.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.EllipseWithParams(img.p, pc, pa, C.double(angle), C.double(startAngle), C.double(endAngle), sColor, C.int(thickness), C.int(lineType), C.int(shift))
}

// Line draws a line segment connecting two points.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga7078a9fae8c7e7d13d24dac2520ae4a2
func Line(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	sp1 := C.struct_Point{
		x: C.int(pt1.X),
		y: C.int(pt1.Y),
	}

	sp2 := C.struct_Point{
		x: C.int(pt2.X),
		y: C.int(pt2.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Line(img.p, sp1, sp2, sColor, C.int(thickness))
}

// Rectangle draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
func Rectangle(img *Mat, r image.Rectangle, c color.RGBA, thickness int) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Rectangle(img.p, cRect, sColor, C.int(thickness))
}

// RectangleWithParams draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
func RectangleWithParams(img *Mat, r image.Rectangle, c color.RGBA, thickness int, lineType LineType, shift int) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.RectangleWithParams(img.p, cRect, sColor, C.int(thickness), C.int(lineType), C.int(shift))
}

// FillPoly fills the area bounded by one or more polygons.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf30888828337aa4c6b56782b5dfbd4b7
func FillPoly(img *Mat, pts PointsVector, c color.RGBA) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.FillPoly(img.p, pts.p, sColor)
}

// FillPolyWithParams fills the area bounded by one or more polygons.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf30888828337aa4c6b56782b5dfbd4b7
func FillPolyWithParams(img *Mat, pts PointsVector, c color.RGBA, lineType LineType, shift int, offset image.Point) {
	offsetP := C.struct_Point{
		x: C.int(offset.X),
		y: C.int(offset.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.FillPolyWithParams(img.p, pts.p, sColor, C.int(lineType), C.int(shift), offsetP)
}

// Polylines draws several polygonal curves.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga1ea127ffbbb7e0bfc4fd6fd2eb64263c
func Polylines(img *Mat, pts PointsVector, isClosed bool, c color.RGBA, thickness int) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Polylines(img.p, pts.p, C.bool(isClosed), sColor, C.int(thickness))
}

// HersheyFont are the font libraries included in OpenCV.
// Only a subset of the available Hershey fonts are supported by OpenCV.
//
// For more information, see:
// http://sources.isc.org/utils/misc/hershey-font.txt
type HersheyFont int

const (
	// FontHersheySimplex is normal size sans-serif font.
	FontHersheySimplex HersheyFont = 0
	// FontHersheyPlain issmall size sans-serif font.
	FontHersheyPlain HersheyFont = 1
	// FontHersheyDuplex normal size sans-serif font
	// (more complex than FontHersheySIMPLEX).
	FontHersheyDuplex HersheyFont = 2
	// FontHersheyComplex i a normal size serif font.
	FontHersheyComplex HersheyFont = 3
	// FontHersheyTriplex is a normal size serif font
	// (more complex than FontHersheyCOMPLEX).
	FontHersheyTriplex HersheyFont = 4
	// FontHersheyComplexSmall is a smaller version of FontHersheyCOMPLEX.
	FontHersheyComplexSmall HersheyFont = 5
	// FontHersheyScriptSimplex is a hand-writing style font.
	FontHersheyScriptSimplex HersheyFont = 6
	// FontHersheyScriptComplex is a more complex variant of FontHersheyScriptSimplex.
	FontHersheyScriptComplex HersheyFont = 7
	// FontItalic is the flag for italic font.
	FontItalic HersheyFont = 16
)

// LineType are the line libraries included in OpenCV.
//
// For more information, see:
// https://vovkos.github.io/doxyrest-showcase/opencv/sphinx_rtd_theme/enum_cv_LineTypes.html
type LineType int

const (
	// Filled line
	Filled LineType = -1
	// Line4 4-connected line
	Line4 LineType = 4
	// Line8 8-connected line
	Line8 LineType = 8
	// LineAA antialiased line
	LineAA LineType = 16
)

// GetTextSize calculates the width and height of a text string.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
func GetTextSize(text string, fontFace HersheyFont, fontScale float64, thickness int) image.Point {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sz := C.GetTextSize(cText, C.int(fontFace), C.double(fontScale), C.int(thickness))
	return image.Pt(int(sz.width), int(sz.height))
}

// GetTextSizeWithBaseline calculates the width and height of a text string including the basline of the text.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness as well as its baseline.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
func GetTextSizeWithBaseline(text string, fontFace HersheyFont, fontScale float64, thickness int) (image.Point, int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cBaseline := C.int(0)

	sz := C.GetTextSizeWithBaseline(cText, C.int(fontFace), C.double(fontScale), C.int(thickness), &cBaseline)
	return image.Pt(int(sz.width), int(sz.height)), int(cBaseline)
}

// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
func PutText(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutText(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness))
	return
}

// PutTextWithParams draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
func PutTextWithParams(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int, lineType LineType, bottomLeftOrigin bool) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutTextWithParams(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness), C.int(lineType), C.bool(bottomLeftOrigin))
	return
}

// InterpolationFlags are bit flags that control the interpolation algorithm
// that is used.
type InterpolationFlags int

const (
	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear InterpolationFlags = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic InterpolationFlags = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea InterpolationFlags = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 InterpolationFlags = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax InterpolationFlags = 7

	// WarpFillOutliers fills all of the destination image pixels. If some of them correspond to outliers in the source image, they are set to zero.
	WarpFillOutliers = 8

	// WarpInverseMap, inverse transformation.
	WarpInverseMap = 16
)

// Resize resizes an image.
// It resizes the image src down to or up to the specified size, storing the
// result in dst. Note that src and dst may be the same image. If you wish to
// scale by factor, an empty sz may be passed and non-zero fx and fy. Likewise,
// if you wish to scale to an explicit size, a non-empty sz may be passed with
// zero for both fx and fy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga47a974309e9102f5f08231edc7e7529d
func Resize(src Mat, dst *Mat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.Resize(src.p, dst.p, pSize, C.double(fx), C.double(fy), C.int(interp))
	return
}

// GetRectSubPix retrieves a pixel rectangle from an image with sub-pixel accuracy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga77576d06075c1a4b6ba1a608850cd614
func GetRectSubPix(src Mat, patchSize image.Point, center image.Point, dst *Mat) {
	sz := C.struct_Size{
		width:  C.int(patchSize.X),
		height: C.int(patchSize.Y),
	}
	pt := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	C.GetRectSubPix(src.p, sz, pt, dst.p)
}

// GetRotationMatrix2D calculates an affine matrix of 2D rotation.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gafbbc470ce83812914a70abfb604f4326
func GetRotationMatrix2D(center image.Point, angle, scale float64) Mat {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	return newMat(C.GetRotationMatrix2D(pc, C.double(angle), C.double(scale)))
}

// WarpAffine applies an affine transformation to an image. For more parameters please check WarpAffineWithParams
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga0203d9ee5fcd28d40dbc4a1ea4451983
func WarpAffine(src Mat, dst *Mat, m Mat, sz image.Point) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.WarpAffine(src.p, dst.p, m.p, pSize)
}

// WarpAffineWithParams applies an affine transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga0203d9ee5fcd28d40dbc4a1ea4451983
func WarpAffineWithParams(src Mat, dst *Mat, m Mat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}
	C.WarpAffineWithParams(src.p, dst.p, m.p, pSize, C.int(flags), C.int(borderType), bv)
}

// WarpPerspective applies a perspective transformation to an image.
// For more parameters please check WarpPerspectiveWithParams.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaf73673a7e8e18ec6963e3774e6a94b87
func WarpPerspective(src Mat, dst *Mat, m Mat, sz image.Point) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.WarpPerspective(src.p, dst.p, m.p, pSize)
}

// WarpPerspectiveWithParams applies a perspective transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaf73673a7e8e18ec6963e3774e6a94b87
func WarpPerspectiveWithParams(src Mat, dst *Mat, m Mat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}
	C.WarpPerspectiveWithParams(src.p, dst.p, m.p, pSize, C.int(flags), C.int(borderType), bv)
}

// Watershed performs a marker-based image segmentation using the watershed algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga3267243e4d3f95165d55a618c65ac6e1
func Watershed(image Mat, markers *Mat) {
	C.Watershed(image.p, markers.p)
}

// ColormapTypes are the 12 GNU Octave/MATLAB equivalent colormaps.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html
type ColormapTypes int

// List of the available color maps
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#ga9a805d8262bcbe273f16be9ea2055a65
const (
	ColormapAutumn  ColormapTypes = 0
	ColormapBone    ColormapTypes = 1
	ColormapJet     ColormapTypes = 2
	ColormapWinter  ColormapTypes = 3
	ColormapRainbow ColormapTypes = 4
	ColormapOcean   ColormapTypes = 5
	ColormapSummer  ColormapTypes = 6
	ColormapSpring  ColormapTypes = 7
	ColormapCool    ColormapTypes = 8
	ColormapHsv     ColormapTypes = 9
	ColormapPink    ColormapTypes = 10
	ColormapHot     ColormapTypes = 11
	ColormapParula  ColormapTypes = 12
)

// ApplyColorMap applies a GNU Octave/MATLAB equivalent colormap on a given image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#gadf478a5e5ff49d8aa24e726ea6f65d15
func ApplyColorMap(src Mat, dst *Mat, colormapType ColormapTypes) {
	C.ApplyColorMap(src.p, dst.p, C.int(colormapType))
}

// ApplyCustomColorMap applies a custom defined colormap on a given image.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d50/group__imgproc__colormap.html#gacb22288ddccc55f9bd9e6d492b409cae
func ApplyCustomColorMap(src Mat, dst *Mat, customColormap Mat) {
	C.ApplyCustomColorMap(src.p, dst.p, customColormap.p)
}

// GetPerspectiveTransform returns 3x3 perspective transformation for the
// corresponding 4 point pairs as image.Point.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform(src, dst PointVector) Mat {
	return newMat(C.GetPerspectiveTransform(src.p, dst.p))
}

// GetPerspectiveTransform2f returns 3x3 perspective transformation for the
// corresponding 4 point pairs as gocv.Point2f.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform2f(src, dst Point2fVector) Mat {
	return newMat(C.GetPerspectiveTransform2f(src.p, dst.p))
}

// GetAffineTransform returns a 2x3 affine transformation matrix for the
// corresponding 3 point pairs as image.Point.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999
func GetAffineTransform(src, dst PointVector) Mat {
	return newMat(C.GetAffineTransform(src.p, dst.p))
}

// GetAffineTransform2f returns a 2x3 affine transformation matrix for the
// corresponding 3 point pairs as gocv.Point2f.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999
func GetAffineTransform2f(src, dst Point2fVector) Mat {
	return newMat(C.GetAffineTransform2f(src.p, dst.p))
}

type HomographyMethod int

const (
	HomograpyMethodAllPoints HomographyMethod = 0
	HomograpyMethodLMEDS     HomographyMethod = 4
	HomograpyMethodRANSAC    HomographyMethod = 8
)

// FindHomography finds an optimal homography matrix using 4 or more point pairs (as opposed to GetPerspectiveTransform, which uses exactly 4)
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga4abc2ece9fab9398f2e560d53c8c9780
func FindHomography(srcPoints Mat, dstPoints *Mat, method HomographyMethod, ransacReprojThreshold float64, mask *Mat, maxIters int, confidence float64) Mat {
	return newMat(C.FindHomography(srcPoints.Ptr(), dstPoints.Ptr(), C.int(method), C.double(ransacReprojThreshold), mask.Ptr(), C.int(maxIters), C.double(confidence)))
}

// DrawContours draws contours outlines or filled contours.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga746c0625f1781f1ffc9056259103edbc
func DrawContours(img *Mat, contours PointsVector, contourIdx int, c color.RGBA, thickness int) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.DrawContours(img.p, contours.p, C.int(contourIdx), sColor, C.int(thickness))
}

// DrawContoursWithParams draws contours outlines or filled contours.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga746c0625f1781f1ffc9056259103edbc
func DrawContoursWithParams(img *Mat, contours PointsVector, contourIdx int, c color.RGBA, thickness int, lineType LineType, hierarchy Mat, maxLevel int, offset image.Point) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}
	offsetP := C.struct_Point{
		x: C.int(offset.X),
		y: C.int(offset.Y),
	}

	C.DrawContoursWithParams(img.p, contours.p, C.int(contourIdx), sColor, C.int(thickness), C.int(lineType), hierarchy.p, C.int(maxLevel), offsetP)
}

// Remap applies a generic geometrical transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gab75ef31ce5cdfb5c44b6da5f3b908ea4
func Remap(src Mat, dst, map1, map2 *Mat, interpolation InterpolationFlags, borderMode BorderType, borderValue color.RGBA) {
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}
	C.Remap(src.p, dst.p, map1.p, map2.p, C.int(interpolation), C.int(borderMode), bv)
}

// Filter2D applies an arbitrary linear filter to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga27c049795ce870216ddfb366086b5a04
func Filter2D(src Mat, dst *Mat, ddepth MatType, kernel Mat, anchor image.Point, delta float64, borderType BorderType) {
	anchorP := C.struct_Point{
		x: C.int(anchor.X),
		y: C.int(anchor.Y),
	}
	C.Filter2D(src.p, dst.p, C.int(ddepth), kernel.p, anchorP, C.double(delta), C.int(borderType))
}

// SepFilter2D applies a separable linear filter to the image.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga910e29ff7d7b105057d1625a4bf6318d
func SepFilter2D(src Mat, dst *Mat, ddepth MatType, kernelX, kernelY Mat, anchor image.Point, delta float64, borderType BorderType) {
	anchorP := C.struct_Point{
		x: C.int(anchor.X),
		y: C.int(anchor.Y),
	}
	C.SepFilter2D(src.p, dst.p, C.int(ddepth), kernelX.p, kernelY.p, anchorP, C.double(delta), C.int(borderType))
}

// LogPolar remaps an image to semilog-polar coordinates space.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaec3a0b126a85b5ca2c667b16e0ae022d
func LogPolar(src Mat, dst *Mat, center image.Point, m float64, flags InterpolationFlags) {
	centerP := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	C.LogPolar(src.p, dst.p, centerP, C.double(m), C.int(flags))
}

// LinearPolar remaps an image to polar coordinates space.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaa38a6884ac8b6e0b9bed47939b5362f3
func LinearPolar(src Mat, dst *Mat, center image.Point, maxRadius float64, flags InterpolationFlags) {
	centerP := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	C.LinearPolar(src.p, dst.p, centerP, C.double(maxRadius), C.int(flags))
}

// DistanceTypes types for Distance Transform and M-estimatorss
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#gaa2bfbebbc5c320526897996aafa1d8eb
type DistanceTypes int

const (
	DistUser   DistanceTypes = 0
	DistL1     DistanceTypes = 1
	DistL2     DistanceTypes = 2
	DistC      DistanceTypes = 3
	DistL12    DistanceTypes = 4
	DistFair   DistanceTypes = 5
	DistWelsch DistanceTypes = 6
	DistHuber  DistanceTypes = 7
)

// FitLine fits a line to a 2D or 3D point set.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf849da1fdafa67ee84b1e9a23b93f91f
func FitLine(pts PointVector, line *Mat, distType DistanceTypes, param, reps, aeps float64) {
	C.FitLine(pts.p, line.p, C.int(distType), C.double(param), C.double(reps), C.double(aeps))
}

// Shape matching methods.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317
type ShapeMatchModes int

const (
	ContoursMatchI1 ShapeMatchModes = 1
	ContoursMatchI2 ShapeMatchModes = 2
	ContoursMatchI3 ShapeMatchModes = 3
)

// Compares two shapes.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317
func MatchShapes(contour1 PointVector, contour2 PointVector, method ShapeMatchModes, parameter float64) float64 {
	return float64(C.MatchShapes(contour1.p, contour2.p, C.int(method), C.double(parameter)))
}

// CLAHE is a wrapper around the cv::CLAHE algorithm.
type CLAHE struct {
	// C.CLAHE
	p unsafe.Pointer
}

// NewCLAHE returns a new CLAHE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html
func NewCLAHE() CLAHE {
	return CLAHE{p: unsafe.Pointer(C.CLAHE_Create())}
}

// NewCLAHEWithParams returns a new CLAHE algorithm
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html
func NewCLAHEWithParams(clipLimit float64, tileGridSize image.Point) CLAHE {
	pSize := C.struct_Size{
		width:  C.int(tileGridSize.X),
		height: C.int(tileGridSize.Y),
	}
	return CLAHE{p: unsafe.Pointer(C.CLAHE_CreateWithParams(C.double(clipLimit), pSize))}
}

// Close CLAHE.
func (c *CLAHE) Close() error {
	C.CLAHE_Close((C.CLAHE)(c.p))
	c.p = nil
	return nil
}

// Apply CLAHE.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/db6/classcv_1_1CLAHE.html#a4e92e0e427de21be8d1fae8dcd862c5e
func (c *CLAHE) Apply(src Mat, dst *Mat) {
	C.CLAHE_Apply((C.CLAHE)(c.p), src.p, dst.p)
}

func InvertAffineTransform(src Mat, dst *Mat) {
	C.InvertAffineTransform(src.p, dst.p)
}

// Apply phaseCorrelate.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga552420a2ace9ef3fb053cd630fdb4952
func PhaseCorrelate(src1, src2, window Mat) (phaseShift Point2f, response float64) {
	var responseDouble C.double
	result := C.PhaseCorrelate(src1.p, src2.p, window.p, &responseDouble)

	return Point2f{
		X: float32(result.x),
		Y: float32(result.y),
	}, float64(responseDouble)
}

// ToImage converts a Mat to a image.Image.
func (m *Mat) ToImage() (image.Image, error) {
	switch m.Type() {
	case MatTypeCV8UC1:
		img := image.NewGray(image.Rect(0, 0, m.Cols(), m.Rows()))
		data, err := m.DataPtrUint8()
		if err != nil {
			return nil, err
		}
		copy(img.Pix, data[0:])
		return img, nil

	case MatTypeCV8UC3:
		dst := NewMat()
		defer dst.Close()

		C.CvtColor(m.p, dst.p, C.int(ColorBGRToRGBA))

		img := image.NewRGBA(image.Rect(0, 0, m.Cols(), m.Rows()))
		data, err := dst.DataPtrUint8()
		if err != nil {
			return nil, err
		}

		copy(img.Pix, data[0:])
		return img, nil

	case MatTypeCV8UC4:
		dst := NewMat()
		defer dst.Close()

		C.CvtColor(m.p, dst.p, C.int(ColorBGRAToRGBA))

		img := image.NewNRGBA(image.Rect(0, 0, m.Cols(), m.Rows()))
		data, err := dst.DataPtrUint8()
		if err != nil {
			return nil, err
		}
		copy(img.Pix, data[0:])
		return img, nil

	default:
		return nil, errors.New("ToImage supports only MatType CV8UC1, CV8UC3 and CV8UC4")
	}
}

// ToImageYUV converts a Mat to a image.YCbCr using image.YCbCrSubsampleRatio420 as default subsampling param.
func (m *Mat) ToImageYUV() (*image.YCbCr, error) {
	img, err := m.ToImage()
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	converted := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio420)

	for row := 0; row < bounds.Max.Y; row++ {
		for col := 0; col < bounds.Max.X; col++ {
			r, g, b, _ := img.At(col, row).RGBA()
			y, cb, cr := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))

			converted.Y[converted.YOffset(col, row)] = y
			converted.Cb[converted.COffset(col, row)] = cb
			converted.Cr[converted.COffset(col, row)] = cr
		}
	}
	return converted, nil
}

// ToImageYUV converts a Mat to a image.YCbCr using provided YUV subsample ratio param.
func (m *Mat) ToImageYUVWithParams(ratio image.YCbCrSubsampleRatio) (*image.YCbCr, error) {
	img, err := m.ToImage()
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	converted := image.NewYCbCr(bounds, ratio)

	for row := 0; row < bounds.Max.Y; row++ {
		for col := 0; col < bounds.Max.X; col++ {
			r, g, b, _ := img.At(col, row).RGBA()
			y, cb, cr := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))

			converted.Y[converted.YOffset(col, row)] = y
			converted.Cb[converted.COffset(col, row)] = cb
			converted.Cr[converted.COffset(col, row)] = cr
		}
	}
	return converted, nil
}

// ImageToMatRGBA converts image.Image to gocv.Mat,
// which represents RGBA image having 8bit for each component.
// Type of Mat is gocv.MatTypeCV8UC4.
func ImageToMatRGBA(img image.Image) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()

	var data []uint8
	switch img.ColorModel() {
	case color.RGBAModel:
		m, res := img.(*image.RGBA)
		if !res {
			return NewMat(), errors.New("Image color format error")
		}
		data = m.Pix

	case color.NRGBAModel:
		m, res := img.(*image.NRGBA)
		if !res {
			return NewMat(), errors.New("Image color format error")
		}
		data = m.Pix

	default:
		data := make([]byte, 0, x*y*3)
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			for i := bounds.Min.X; i < bounds.Max.X; i++ {
				r, g, b, _ := img.At(i, j).RGBA()
				data = append(data, byte(b>>8), byte(g>>8), byte(r>>8))
			}
		}
		return NewMatFromBytes(y, x, MatTypeCV8UC3, data)
	}

	// speed up the conversion process of RGBA format
	cvt, err := NewMatFromBytes(y, x, MatTypeCV8UC4, data)
	if err != nil {
		return NewMat(), err
	}

	defer cvt.Close()

	dst := NewMat()
	C.CvtColor(cvt.p, dst.p, C.int(ColorBGRAToRGBA))
	return dst, nil
}

// ImageToMatRGB converts image.Image to gocv.Mat,
// which represents RGB image having 8bit for each component.
// Type of Mat is gocv.MatTypeCV8UC3.
func ImageToMatRGB(img image.Image) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()

	var data []uint8
	switch img.ColorModel() {
	case color.RGBAModel:
		m, res := img.(*image.RGBA)
		if true != res {
			return NewMat(), errors.New("Image color format error")
		}
		data = m.Pix
		// speed up the conversion process of RGBA format
		src, err := NewMatFromBytes(y, x, MatTypeCV8UC4, data)
		if err != nil {
			return NewMat(), err
		}
		defer src.Close()

		dst := NewMat()
		CvtColor(src, &dst, ColorRGBAToBGR)
		return dst, nil

	default:
		data := make([]byte, 0, x*y*3)
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			for i := bounds.Min.X; i < bounds.Max.X; i++ {
				r, g, b, _ := img.At(i, j).RGBA()
				data = append(data, byte(b>>8), byte(g>>8), byte(r>>8))
			}
		}
		return NewMatFromBytes(y, x, MatTypeCV8UC3, data)
	}
}

// ImageGrayToMatGray converts image.Gray to gocv.Mat,
// which represents grayscale image 8bit.
// Type of Mat is gocv.MatTypeCV8UC1.
func ImageGrayToMatGray(img *image.Gray) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	m, err := NewMatFromBytes(y, x, MatTypeCV8UC1, img.Pix)
	if err != nil {
		return NewMat(), err
	}
	return m, nil
}

// Adds the square of a source image to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga1a567a79901513811ff3b9976923b199
//

func Accumulate(src Mat, dst *Mat) {
	C.Mat_Accumulate(src.p, dst.p)
}

// Adds an image to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga1a567a79901513811ff3b9976923b199
func AccumulateWithMask(src Mat, dst *Mat, mask Mat) {
	C.Mat_AccumulateWithMask(src.p, dst.p, mask.p)
}

// Adds the square of a source image to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#gacb75e7ffb573227088cef9ceaf80be8c
func AccumulateSquare(src Mat, dst *Mat) {
	C.Mat_AccumulateSquare(src.p, dst.p)
}

// Adds the square of a source image to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#gacb75e7ffb573227088cef9ceaf80be8c
func AccumulateSquareWithMask(src Mat, dst *Mat, mask Mat) {
	C.Mat_AccumulateSquareWithMask(src.p, dst.p, mask.p)
}

// Adds the per-element product of two input images to the accumulator image.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga82518a940ecfda49460f66117ac82520
func AccumulateProduct(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_AccumulateProduct(src1.p, src2.p, dst.p)
}

// Adds the per-element product of two input images to the accumulator image with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga82518a940ecfda49460f66117ac82520
func AccumulateProductWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	C.Mat_AccumulateProductWithMask(src1.p, src2.p, dst.p, mask.p)
}

// Updates a running average.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga4f9552b541187f61f6818e8d2d826bc7
func AccumulatedWeighted(src Mat, dst *Mat, alpha float64) {
	C.Mat_AccumulatedWeighted(src.p, dst.p, C.double(alpha))
}

// Updates a running average with mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga4f9552b541187f61f6818e8d2d826bc7
func AccumulatedWeightedWithMask(src Mat, dst *Mat, alpha float64, mask Mat) {
	C.Mat_AccumulatedWeightedWithMask(src.p, dst.p, C.double(alpha), mask.p)
}
