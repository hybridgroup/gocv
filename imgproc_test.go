package gocv

import (
	"image"
	"image/color"
	"math"
	"reflect"
	"testing"
)

func TestApproxPolyDP(t *testing.T) {
	img := NewMatWithSize(100, 200, MatTypeCV8UC1)
	defer img.Close()

	white := color.RGBA{255, 255, 255, 255}
	// Draw triangle
	Line(&img, image.Pt(25, 25), image.Pt(25, 75), white, 1)
	Line(&img, image.Pt(25, 75), image.Pt(75, 50), white, 1)
	Line(&img, image.Pt(75, 50), image.Pt(25, 25), white, 1)
	// Draw rectangle
	Rectangle(&img, image.Rect(125, 25, 175, 75), white, 1)

	contours := FindContours(img, RetrievalExternal, ChainApproxSimple)

	trianglePerimeter := ArcLength(contours[0], true)
	triangleContour := ApproxPolyDP(contours[0], 0.04*trianglePerimeter, true)
	expectedTriangleContour := []image.Point{image.Pt(25, 25), image.Pt(25, 75), image.Pt(75, 50)}
	if !reflect.DeepEqual(triangleContour, expectedTriangleContour) {
		t.Errorf("Failed to approximate triangle.\nActual:%v\nExpect:%v", triangleContour, expectedTriangleContour)
	}

	rectPerimeter := ArcLength(contours[1], true)
	rectContour := ApproxPolyDP(contours[1], 0.04*rectPerimeter, true)
	expectedRectContour := []image.Point{image.Pt(125, 24), image.Pt(124, 75), image.Pt(175, 76), image.Pt(176, 25)}
	if !reflect.DeepEqual(rectContour, expectedRectContour) {
		t.Errorf("Failed to approximate rectangle.\nActual:%v\nExpect:%v", rectContour, expectedRectContour)
	}
}

func TestConvexity(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	res := FindContours(img, RetrievalExternal, ChainApproxSimple)
	if len(res) < 1 {
		t.Error("Invalid FindContours test")
	}

	area := ContourArea(res[0])
	if area != 127280.0 {
		t.Errorf("Invalid ContourArea test: %f", area)
	}

	hull := NewMat()
	defer hull.Close()

	ConvexHull(res[0], &hull, true, false)
	if hull.Empty() {
		t.Error("Invalid ConvexHull test")
	}

	defects := NewMat()
	defer defects.Close()

	ConvexityDefects(res[0], hull, &defects)
	if defects.Empty() {
		t.Error("Invalid ConvexityDefects test")
	}
}

func TestMinEnclosingCircle(t *testing.T) {
	pts := []image.Point{
		image.Pt(0, 2),
		image.Pt(2, 0),
		image.Pt(0, -2),
		image.Pt(-2, 0),
		image.Pt(1, -1),
	}
	x, y, radius := MinEnclosingCircle(pts)
	const epsilon = 0.001
	if math.Abs(float64(radius-2.0)) > epsilon ||
		math.Abs(float64(x-0.0)) > epsilon ||
		math.Abs(float64(y-0.0)) > epsilon {
		t.Error("Invalid circle returned in MinEnclosingCircle test")
	}
}

func TestCvtColor(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in CvtColor test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img, &dest, ColorBGRAToGray)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid convert in CvtColor test")
	}
}

func TestBilateralFilter(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in BilateralFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	BilateralFilter(img, &dest, 1, 2.0, 3.0)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid BilateralFilter test")
	}
}

func TestBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in GaussianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Blur(img, &dest, image.Pt(3, 3))
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
	}
}

func TestSobel(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in Sobel test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Sobel(img, &dest, MatTypeCV16S, 0, 1, 3, 1, 0, BorderDefault)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Sober test")
	}
}

func TestSpatialGradient(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in SpatialGradient test")
	}
	defer img.Close()

	dx := NewMat()
	defer dx.Close()

	dy := NewMat()
	defer dy.Close()

	SpatialGradient(img, &dx, &dy, MatTypeCV16S, BorderDefault)
	if dx.Empty() || dy.Empty() || img.Rows() != dx.Rows() || img.Rows() != dy.Rows() || img.Cols() != dx.Cols() || img.Cols() != dy.Cols() {
		t.Error("Invalid SpatialGradient test")
	}
}

func TestBoxFilter(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in BoxFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	BoxFilter(img, &dest, -1, image.Pt(3, 3))
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid BoxFilter test")
	}
}

func TestSqBoxFilter(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in SqBoxFilter test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	SqBoxFilter(img, &dest, -1, image.Pt(3, 3))
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid SqBoxFilter test")
	}
}

func TestDilate(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Dilate test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	Dilate(img, &dest, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Dilate test")
	}
}

func TestDistanceTransform(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in DistanceTransform test")
	}
	defer img.Close()

	gray := NewMat()
	defer gray.Close()
	CvtColor(img, &gray, ColorBGRToGray)

	threshImg := NewMat()
	defer threshImg.Close()
	Threshold(gray, &threshImg, 25, 255, ThresholdBinary)

	dest := NewMat()
	defer dest.Close()

	labels := NewMat()
	defer labels.Close()

	DistanceTransform(threshImg, &dest, &labels, DistL2, DistanceMask3, DistanceLabelCComp)
	if dest.Empty() || dest.Rows() != img.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid DistanceTransform test")
	}
}

func TestMatchTemplate(t *testing.T) {
	imgScene := IMRead("images/face.jpg", IMReadGrayScale)
	if imgScene.Empty() {
		t.Error("Invalid read of face.jpg in MatchTemplate test")
	}
	defer imgScene.Close()

	imgTemplate := IMRead("images/toy.jpg", IMReadGrayScale)
	if imgTemplate.Empty() {
		t.Error("Invalid read of toy.jpg in MatchTemplate test")
	}
	defer imgTemplate.Close()

	result := NewMat()
	defer result.Close()
	m := NewMat()
	MatchTemplate(imgScene, imgTemplate, &result, TmCcoeffNormed, m)
	m.Close()
	_, maxConfidence, _, _ := MinMaxLoc(result)
	if maxConfidence < 0.95 {
		t.Errorf("Max confidence of %f is too low. MatchTemplate could not find template in scene.", maxConfidence)
	}
}

func TestMoments(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in Moments test")
	}
	defer img.Close()

	result := Moments(img, true)
	if len(result) < 1 {
		t.Errorf("Invalid Moments test: %v", result)
	}
}

func TestPyrDown(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in PyrDown test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	PyrDown(img, &dest, image.Point{X: dest.Cols(), Y: dest.Rows()}, BorderDefault)
	if dest.Empty() && math.Abs(float64(img.Cols()-2*dest.Cols())) < 2.0 && math.Abs(float64(img.Rows()-2*dest.Rows())) < 2.0 {
		t.Error("Invalid PyrDown test")
	}
}

func TestPyrUp(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in PyrUp test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	PyrUp(img, &dest, image.Point{X: dest.Cols(), Y: dest.Rows()}, BorderDefault)
	if dest.Empty() && math.Abs(float64(2*img.Cols()-dest.Cols())) < 2.0 && math.Abs(float64(2*img.Rows()-dest.Rows())) < 2.0 {
		t.Error("Invalid PyrUp test")
	}
}

func TestBoxPoints(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in BoxPoints test")
	}
	defer img.Close()

	threshImg := NewMat()
	defer threshImg.Close()

	Threshold(img, &threshImg, 25, 255, ThresholdBinary)

	contours := FindContours(threshImg, RetrievalExternal, ChainApproxSimple)
	contour := contours[0]

	hull := NewMat()
	defer hull.Close()
	ConvexHull(contour, &hull, false, false)
	hullPoints := []image.Point{}
	for i := 0; i < hull.Cols(); i++ {
		for j := 0; j < hull.Rows(); j++ {
			p := hull.GetIntAt(j, i)
			hullPoints = append(hullPoints, contour[p])
		}
	}
	rect := MinAreaRect(hullPoints)
	pts := NewMat()
	defer pts.Close()
	BoxPoints(rect, &pts)

	if pts.Empty() || pts.Rows() != 4 || pts.Cols() != 2 {
		t.Error("Invalid BoxPoints test")
	}
}

func TestMinAreaRect(t *testing.T) {
	src := []image.Point{
		image.Pt(0, 2),
		image.Pt(2, 0),
		image.Pt(4, 2),
		image.Pt(2, 4),
	}
	m := MinAreaRect(src)

	if m.Center.X != 2 {
		t.Errorf("TestMinAreaRect(): unexpected center.X = %v, want = %v", m.Center.X, 2)
	}
	if m.Center.Y != 2 {
		t.Errorf("TestMinAreaRect(): unexpected center.Y = %v, want = %v", m.Center.Y, 2)
	}
	if m.Angle != -45.0 {
		t.Errorf("TestMinAreaRect(): unexpected angle = %v, want = %v", m.Angle, -45.0)
	}
}

func TestFitEllipse(t *testing.T) {
	src := []image.Point{
		image.Pt(1, 1),
		image.Pt(0, 1),
		image.Pt(0, 2),
		image.Pt(1, 3),
		image.Pt(2, 3),
		image.Pt(4, 2),
		image.Pt(4, 1),
		image.Pt(0, 3),
		image.Pt(0, 2),
	}
	rect := FitEllipse(src)
	if rect.Center.X != 2 {
		t.Errorf("TestFitEllipse(): unexpected center.X = %v, want = %v", rect.Center.X, 2)
	}
	if rect.Center.Y != 2 {
		t.Errorf("TestFitEllipse(): unexpected center.Y = %v, want = %v", rect.Center.Y, 2)
	}
	if rect.Angle != 78.60807800292969 {
		t.Errorf("TestFitEllipse(): unexpected angle = %v, want = %v", rect.Angle, 78.60807800292969)
	}
}

func TestFindContours(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	res := FindContours(img, RetrievalExternal, ChainApproxSimple)
	if len(res) < 1 {
		t.Error("Invalid FindContours test")
	}

	area := ContourArea(res[0])
	if area != 127280.0 {
		t.Errorf("Invalid ContourArea test: %f", area)
	}

	r := BoundingRect(res[0])
	if !r.Eq(image.Rect(0, 0, 400, 320)) {
		t.Errorf("Invalid BoundingRect test: %v", r)
	}

	length := ArcLength(res[0], true)
	if int(length) != 1436 {
		t.Errorf("Invalid ArcLength test: %f", length)
	}

	length = ArcLength(res[0], false)
	if int(length) != 1037 {
		t.Errorf("Invalid ArcLength test: %f", length)
	}
}

func TestConnectedComponents(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	labels := NewMat()
	defer labels.Close()
	res := ConnectedComponents(img, &labels)
	if res < 1 || labels.Empty() {
		t.Error("Invalid ConnectedComponents test")
	}
}

func TestConnectedComponentsWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	labels := NewMat()
	defer labels.Close()
	res := ConnectedComponentsWithParams(img, &labels, 8, MatTypeCV32S, CCL_DEFAULT)
	if res < 1 || labels.Empty() {
		t.Error("Invalid ConnectedComponentsWithParams test")
	}
}

func TestConnectedComponentsWithStats(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	labels := NewMat()
	defer labels.Close()

	stats := NewMat()
	defer stats.Close()

	centroids := NewMat()
	defer centroids.Close()

	res := ConnectedComponentsWithStats(img, &labels, &stats, &centroids)
	if res < 1 || labels.Empty() || stats.Empty() || centroids.Empty() {
		t.Error("Invalid ConnectedComponentsWithStats test")
	}
}

func TestConnectedComponentsWithStatsWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in FindContours test")
	}
	defer img.Close()

	labels := NewMat()
	defer labels.Close()

	stats := NewMat()
	defer stats.Close()

	centroids := NewMat()
	defer centroids.Close()

	res := ConnectedComponentsWithStatsWithParams(img, &labels, &stats, &centroids,
		8, MatTypeCV32S, CCL_DEFAULT)
	if res < 1 || labels.Empty() || stats.Empty() || centroids.Empty() {
		t.Error("Invalid ConnectedComponentsWithStatsWithParams test")
	}
}

func TestErode(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Erode test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	Erode(img, &dest, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Erode test")
	}
}

func TestErodeWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in ErodeWithParams test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	ErodeWithParams(img, &dest, kernel, image.Pt(-1, -1), 3, 0)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid ErodeWithParams test")
	}
}

func TestMorphologyDefaultBorderValue(t *testing.T) {
	zeroScalar := Scalar{}
	morphologyDefaultBorderValue := MorphologyDefaultBorderValue()

	if reflect.DeepEqual(zeroScalar, morphologyDefaultBorderValue) {
		t.Error("Got zero valued scalar")
	}

}

func TestMorphologyEx(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in MorphologyEx test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	MorphologyEx(img, &dest, MorphOpen, kernel)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid MorphologyEx test")
	}
}

func TestMorphologyExWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in MorphologyEx test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	MorphologyExWithParams(img, &dest, MorphOpen, kernel, 2, BorderConstant)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid MorphologyExWithParams test")
	}
}

func TestGaussianBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in GaussianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	GaussianBlur(img, &dest, image.Pt(23, 23), 30, 50, 4)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
	}
}

func TestGetGaussianKernel(t *testing.T) {
	kernel := GetGaussianKernel(1, 0.5)
	defer kernel.Close()
	if kernel.Empty() {
		t.Error("Invalid GetGaussianKernel test")
	}

}

func TestGetGaussianKernelWithParams(t *testing.T) {
	kernel := GetGaussianKernelWithParams(1, 0.5, MatTypeCV64F)
	defer kernel.Close()
	if kernel.Empty() {
		t.Error("Invalid GetGaussianKernel test")
	}

}

func TestLaplacian(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Laplacian test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Laplacian(img, &dest, MatTypeCV16S, 1, 1, 0, BorderDefault)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Laplacian test")
	}
}

func TestScharr(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Scharr test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Scharr(img, &dest, MatTypeCV16S, 1, 0, 0, 0, BorderDefault)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Scharr test")
	}
}

func TestMedianBlur(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in MedianBlur test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	MedianBlur(img, &dest, 3)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid MedianBlur test")
	}
}

func TestCanny(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in Canny test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Canny(img, &dest, 50, 150)
	if dest.Empty() {
		t.Error("Empty Canny test")
	}
	if img.Rows() != dest.Rows() {
		t.Error("Invalid Canny test rows")
	}
	if img.Cols() != dest.Cols() {
		t.Error("Invalid Canny test cols")
	}
}

func TestGoodFeaturesToTrackAndCornerSubPix(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in GoodFeaturesToTrack test")
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(img, &corners, 500, 0.01, 10)
	if corners.Empty() {
		t.Error("Empty GoodFeaturesToTrack test")
	}
	if corners.Rows() != 205 {
		t.Errorf("Invalid GoodFeaturesToTrack test rows: %v", corners.Rows())
	}
	if corners.Cols() != 1 {
		t.Errorf("Invalid GoodFeaturesToTrack test cols: %v", corners.Cols())
	}

	tc := NewTermCriteria(Count|EPS, 20, 0.03)

	CornerSubPix(img, &corners, image.Pt(10, 10), image.Pt(-1, -1), tc)
	if corners.Empty() {
		t.Error("Empty CornerSubPix test")
	}
	if corners.Rows() != 205 {
		t.Errorf("Invalid CornerSubPix test rows: %v", corners.Rows())
	}
	if corners.Cols() != 1 {
		t.Errorf("Invalid CornerSubPix test cols: %v", corners.Cols())
	}
}

func TestGrabCut(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in GrabCut test")
	}
	defer img.Close()

	src := NewMat()
	defer src.Close()
	CvtColor(img, &img, ColorRGBAToBGR)
	img.ConvertTo(&src, MatTypeCV8UC3)

	mask := NewMatWithSize(img.Rows(), img.Cols(), MatTypeCV8U)
	defer mask.Close()

	bgdModel := NewMat()
	defer bgdModel.Close()
	fgdModel := NewMat()
	defer fgdModel.Close()

	r := image.Rect(0, 0, 50, 50)

	GrabCut(src, &mask, r, &bgdModel, &fgdModel, 1, GCEval)
	if bgdModel.Empty() {
		t.Error("Empty bgdmodel")
	} else if fgdModel.Empty() {
		t.Error("Empty fgdmodel")
	}
}

func TestHoughCircles(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughCircles test")
	}
	defer img.Close()

	circles := NewMat()
	defer circles.Close()

	HoughCircles(img, &circles, 3, 5.0, 5.0)
	if circles.Empty() {
		t.Error("Empty HoughCircles test")
	}
	if circles.Rows() != 1 {
		t.Errorf("Invalid HoughCircles test rows: %v", circles.Rows())
	}
	if circles.Cols() < 317 || circles.Cols() > 334 {
		t.Errorf("Invalid HoughCircles test cols: %v", circles.Cols())
	}
}

func TestHoughCirclesWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughCircles test")
	}
	defer img.Close()

	circles := NewMat()
	defer circles.Close()

	HoughCirclesWithParams(img, &circles, 3, 5.0, 5.0, 100, 100, 0, 0)
	if circles.Empty() {
		t.Error("Empty HoughCirclesWithParams test")
	}
	if circles.Rows() != 1 {
		t.Errorf("Invalid HoughCirclesWithParams test rows: %v", circles.Rows())
	}
	if circles.Cols() < 317 || circles.Cols() > 334 {
		t.Errorf("Invalid HoughCirclesWithParams test cols: %v", circles.Cols())
	}
}

func TestHoughLines(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLines test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	HoughLines(img, &dest, math.Pi/180, 1, 1)
	if dest.Empty() {
		t.Error("Empty HoughLines test")
	}
	rowsDiff := dest.Rows() - 10411
	if rowsDiff > 2 || rowsDiff < 0 {
		t.Errorf("Invalid HoughLines test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLines test cols: %v", dest.Cols())
	}

	if dest.GetFloatAt(0, 0) != 0 && dest.GetFloatAt(0, 1) != 0 {
		t.Errorf("Invalid HoughLines first test element: %v, %v", dest.GetFloatAt(0, 0), dest.GetFloatAt(0, 1))
	}

	if dest.GetFloatAt(1, 0) != 0.99483764 && dest.GetFloatAt(1, 1) != 0 {
		t.Errorf("Invalid HoughLines second test element: %v, %v", dest.GetFloatAt(1, 0), dest.GetFloatAt(1, 1))
	}

	if dest.GetFloatAt(10409, 0) != -118.246056 && dest.GetFloatAt(10409, 1) != 2 {
		t.Errorf("Invalid HoughLines penultimate test element: %v, %v", dest.GetFloatAt(10409, 0), dest.GetFloatAt(10409, 1))
	}

	if dest.GetFloatAt(10410, 0) != -118.246056 && dest.GetFloatAt(10410, 1) != 2 {
		t.Errorf("Invalid HoughLines last test element: %v, %v", dest.GetFloatAt(10410, 0), dest.GetFloatAt(10410, 1))
	}
}

func TestHoughLinesP(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLinesP test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	HoughLinesP(img, &dest, math.Pi/180, 1, 1)
	if dest.Empty() {
		t.Error("Empty HoughLinesP test")
	}
	if dest.Rows() != 435 {
		t.Errorf("Invalid HoughLinesP test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLinesP test cols: %v", dest.Cols())
	}

	if dest.GetIntAt(0, 0) != 197 && dest.GetIntAt(0, 1) != 319 && dest.GetIntAt(0, 2) != 197 && dest.GetIntAt(0, 3) != 197 {
		t.Errorf("Invalid HoughLinesP first test element: %v, %v, %v, %v", dest.GetIntAt(0, 0), dest.GetIntAt(0, 1), dest.GetIntAt(0, 2), dest.GetIntAt(0, 3))
	}

	if dest.GetIntAt(1, 0) != 62 && dest.GetIntAt(1, 1) != 319 && dest.GetIntAt(1, 2) != 197 && dest.GetIntAt(1, 3) != 197 {
		t.Errorf("Invalid HoughLinesP second test element: %v, %v, %v, %v", dest.GetIntAt(1, 0), dest.GetIntAt(1, 1), dest.GetIntAt(1, 2), dest.GetIntAt(1, 3))
	}

	if dest.GetIntAt(433, 0) != 357 && dest.GetIntAt(433, 1) != 316 && dest.GetIntAt(433, 2) != 357 && dest.GetIntAt(433, 3) != 316 {
		t.Errorf("Invalid HoughLinesP penultimate test element: %v, %v, %v, %v", dest.GetIntAt(433, 0), dest.GetIntAt(433, 1), dest.GetIntAt(433, 2), dest.GetIntAt(433, 3))
	}

	if dest.GetIntAt(434, 0) != 43 && dest.GetIntAt(434, 1) != 316 && dest.GetIntAt(434, 2) != 43 && dest.GetIntAt(434, 3) != 316 {
		t.Errorf("Invalid HoughLinesP last test element: %v, %v, %v, %v", dest.GetIntAt(434, 0), dest.GetIntAt(434, 1), dest.GetIntAt(434, 2), dest.GetIntAt(434, 3))
	}
}

func TestHoughLinesPWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in HoughLinesPWithParams test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	HoughLinesPWithParams(img, &dest, math.Pi/180, 1, 1, 0, 0)
	if dest.Empty() {
		t.Error("Empty HoughLinesPWithParams test")
	}
	if dest.Rows() != 435 {
		t.Errorf("Invalid HoughLinesPWithParams test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLinesPWithParams test cols: %v", dest.Cols())
	}

	if dest.GetIntAt(0, 0) != 197 && dest.GetIntAt(0, 1) != 319 && dest.GetIntAt(0, 2) != 197 && dest.GetIntAt(0, 3) != 197 {
		t.Errorf("Invalid HoughLinesPWithParams first test element: %v, %v, %v, %v", dest.GetIntAt(0, 0), dest.GetIntAt(0, 1), dest.GetIntAt(0, 2), dest.GetIntAt(0, 3))
	}

	if dest.GetIntAt(1, 0) != 62 && dest.GetIntAt(1, 1) != 319 && dest.GetIntAt(1, 2) != 197 && dest.GetIntAt(1, 3) != 197 {
		t.Errorf("Invalid HoughLinesPWithParams second test element: %v, %v, %v, %v", dest.GetIntAt(1, 0), dest.GetIntAt(1, 1), dest.GetIntAt(1, 2), dest.GetIntAt(1, 3))
	}

	if dest.GetIntAt(433, 0) != 357 && dest.GetIntAt(433, 1) != 316 && dest.GetIntAt(433, 2) != 357 && dest.GetIntAt(433, 3) != 316 {
		t.Errorf("Invalid HoughLinesPWithParams penultimate test element: %v, %v, %v, %v", dest.GetIntAt(433, 0), dest.GetIntAt(433, 1), dest.GetIntAt(433, 2), dest.GetIntAt(433, 3))
	}

	if dest.GetIntAt(434, 0) != 43 && dest.GetIntAt(434, 1) != 316 && dest.GetIntAt(434, 2) != 43 && dest.GetIntAt(434, 3) != 316 {
		t.Errorf("Invalid HoughLinesPWithParams last test element: %v, %v, %v, %v", dest.GetIntAt(434, 0), dest.GetIntAt(434, 1), dest.GetIntAt(434, 2), dest.GetIntAt(434, 3))
	}
}

func TestHoughLinesPointSet(t *testing.T) {

	points := [][2]int{
		{0, 369}, {10, 364}, {20, 358}, {30, 352},
		{40, 346}, {50, 341}, {60, 335}, {70, 329},
		{80, 323}, {90, 318}, {100, 312}, {110, 306},
		{120, 300}, {130, 295}, {140, 289}, {150, 284},
		{160, 277}, {170, 271}, {180, 266}, {190, 260},
	}

	img := NewMatWithSize(len(points), 1, MatTypeCV32F+MatChannels2)
	defer img.Close()
	for i, p := range points {
		img.SetFloatAt(i, 0, float32(p[0]))
		img.SetFloatAt(i, 1, float32(p[1]))
	}

	dest := NewMat()
	defer dest.Close()

	rhoMin, rhoMax, rhoStep := float32(0), float32(360), float32(1)
	thetaMin, thetaMax, thetaStep := float32(0), float32(math.Pi/2), float32(math.Pi/180)

	HoughLinesPointSet(img, &dest, 20, 1,
		rhoMin, rhoMax, rhoStep,
		thetaMin, thetaMax, thetaStep)

	if dest.Empty() {
		t.Error("Empty HoughLinesPointSet test")
	}
	if dest.Rows() != 20 {
		t.Errorf("Invalid HoughLinesPointSet test rows: %v", dest.Rows())
	}
	if dest.Cols() != 1 {
		t.Errorf("Invalid HoughLinesPointSet test cols: %v", dest.Cols())
	}

	if dest.GetDoubleAt(0, 0) != 19 && dest.GetDoubleAt(0, 1) != 320 && dest.GetDoubleAt(0, 2) != 1.0471975803375244 {
		t.Errorf("Invalid HoughLinesPointSet first test element: %v, %v, %v", dest.GetDoubleAt(0, 0), dest.GetDoubleAt(0, 1), dest.GetDoubleAt(0, 2))
	}

	if dest.GetDoubleAt(1, 0) != 7 && dest.GetDoubleAt(1, 1) != 321 && dest.GetDoubleAt(1, 2) != 1.0646508932113647 {
		t.Errorf("Invalid HoughLinesPointSet second test element: %v, %v, %v", dest.GetDoubleAt(1, 0), dest.GetDoubleAt(1, 1), dest.GetDoubleAt(1, 2))
	}

	if dest.GetDoubleAt(18, 0) != 2 && dest.GetDoubleAt(18, 1) != 317 && dest.GetDoubleAt(18, 2) != 0 {
		t.Errorf("Invalid HoughLinesPointSet penultimate test element: %v, %v, %v", dest.GetDoubleAt(18, 0), dest.GetDoubleAt(18, 1), dest.GetDoubleAt(18, 2))
	}

	if dest.GetDoubleAt(19, 0) != 2 && dest.GetDoubleAt(19, 1) != 330 && dest.GetDoubleAt(19, 2) != 0 {
		t.Errorf("Invalid HoughLinesPointSet last test element: %v, %v, %v", dest.GetDoubleAt(19, 0), dest.GetDoubleAt(19, 1), dest.GetDoubleAt(19, 1))
	}
}

func TestIntegral(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Integral test")
	}
	defer img.Close()

	sum := NewMat()
	defer sum.Close()
	sqSum := NewMat()
	defer sqSum.Close()
	tilted := NewMat()
	defer tilted.Close()

	Integral(img, &sum, &sqSum, &tilted)
	if sum.Empty() || sqSum.Empty() || tilted.Empty() {
		t.Error("Invalid Integral test")
	}
}

func TestThreshold(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in Threshold test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	Threshold(img, &dest, 25, 255, ThresholdBinary)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Threshold test")
	}
}
func TestAdaptiveThreshold(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in AdaptiveThreshold test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	AdaptiveThreshold(img, &dest, 255, AdaptiveThresholdMean, ThresholdBinary, 11, 2)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Threshold test")
	}
}

func TestEqualizeHist(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in EqualizeHist test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	EqualizeHist(img, &dest)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid EqualizeHist test")
	}
}

func TestCalcHist(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in CalcHist test")
	}
	defer img.Close()

	hist := NewMat()
	defer hist.Close()

	mask := NewMat()
	defer mask.Close()

	CalcHist([]Mat{img}, []int{0}, mask, &hist, []int{256}, []float64{0.0, 256.0}, false)
	if hist.Empty() || hist.Rows() != 256 || hist.Cols() != 1 {
		t.Error("Invalid CalcHist test")
	}
}

func TestCalcBackProject(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in CalcHist test")
	}
	defer img.Close()

	hist := NewMat()
	defer hist.Close()

	backProject := NewMat()
	defer backProject.Close()

	mask := NewMat()
	defer mask.Close()

	CalcHist([]Mat{img}, []int{0}, mask, &hist, []int{256}, []float64{0.0, 256.0}, false)
	CalcBackProject([]Mat{img}, []int{0}, hist, &backProject, []float64{0.0, 256.0}, false)
	if backProject.Empty() {
		t.Error("Invalid CalcBackProject test")
	}
}

func TestCompareHist(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in CompareHist test")
	}
	defer img.Close()

	hist1 := NewMat()
	defer hist1.Close()

	hist2 := NewMat()
	defer hist2.Close()

	mask := NewMat()
	defer mask.Close()

	CalcHist([]Mat{img}, []int{0}, mask, &hist1, []int{256}, []float64{0.0, 256.0}, false)
	CalcHist([]Mat{img}, []int{0}, mask, &hist2, []int{256}, []float64{0.0, 256.0}, false)
	dist := CompareHist(hist1, hist2, HistCmpCorrel)
	if dist != 1 {
		t.Error("Invalid CompareHist test")
	}

}

func TestDrawing(t *testing.T) {
	img := NewMatWithSize(150, 150, MatTypeCV8U)
	if img.Empty() {
		t.Error("Invalid Mat in Rectangle")
	}
	defer img.Close()

	ArrowedLine(&img, image.Pt(50, 50), image.Pt(75, 75), color.RGBA{0, 0, 255, 0}, 3)
	Circle(&img, image.Pt(60, 60), 20, color.RGBA{0, 0, 255, 0}, 3)
	Rectangle(&img, image.Rect(50, 50, 75, 75), color.RGBA{0, 0, 255, 0}, 3)
	Line(&img, image.Pt(50, 50), image.Pt(75, 75), color.RGBA{0, 0, 255, 0}, 3)

	if img.Empty() {
		t.Error("Error in Rectangle test")
	}
}

func TestGetTextSize(t *testing.T) {
	size := GetTextSize("test", FontHersheySimplex, 1.2, 1)
	if size.X != 72 {
		t.Error("Invalid text size width")
	}

	if size.Y != 26 {
		t.Error("Invalid text size height")
	}
}
func TestPutText(t *testing.T) {
	img := NewMatWithSize(150, 150, MatTypeCV8U)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
	defer img.Close()

	pt := image.Pt(10, 10)
	PutText(&img, "Testing", pt, FontHersheyPlain, 1.2, color.RGBA{255, 255, 255, 0}, 2)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}
func TestPutTextWithParams(t *testing.T) {
	img := NewMatWithSize(150, 150, MatTypeCV8U)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
	defer img.Close()

	pt := image.Pt(10, 10)
	PutTextWithParams(&img, "Testing", pt, FontHersheyPlain, 1.2, color.RGBA{255, 255, 255, 0}, 2, LineAA, false)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}

func TestResize(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	Resize(src, &dst, image.Point{}, 0.5, 0.5, InterpolationDefault)
	if dst.Cols() != 200 || dst.Rows() != 172 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}

	Resize(src, &dst, image.Pt(440, 377), 0, 0, InterpolationCubic)
	if dst.Cols() != 440 || dst.Rows() != 377 {
		t.Errorf("Expected dst size of 440x377 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestGetRectSubPix(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	GetRectSubPix(src, image.Point{20, 30}, image.Point{200, 172}, &dst)
	if dst.Cols() != 20 || dst.Rows() != 30 {
		t.Errorf("Expected dst size of 20x30 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestGetRotationMatrix2D(t *testing.T) {
	type args struct {
		center image.Point
		angle  float64
		scale  float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "90",
			args: args{image.Point{0, 0}, 90.0, 1.0},
			want: [][]float64{
				{6.123233995736766e-17, 1, 0},
				{-1, 6.123233995736766e-17, 0},
			},
		},
		{
			name: "45",
			args: args{image.Point{0, 0}, 45.0, 1.0},
			want: [][]float64{
				{0.7071067811865476, 0.7071067811865475, 0},
				{-0.7071067811865475, 0.7071067811865476, 0},
			},
		},
		{
			name: "0",
			args: args{image.Point{0, 0}, 0.0, 1.0},
			want: [][]float64{
				{1, 0, 0},
				{-0, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRotationMatrix2D(tt.args.center, tt.args.angle, tt.args.scale)
			for row := 0; row < got.Rows(); row++ {
				for col := 0; col < got.Cols(); col++ {
					if !floatEquals(got.GetDoubleAt(row, col), tt.want[row][col]) {
						t.Errorf("GetRotationMatrix2D() = %v, want %v at row:%v col:%v", got.GetDoubleAt(row, col), tt.want[row][col], row, col)
					}
				}
			}
			got.Close()
		})
	}
}

func TestWarpAffine(t *testing.T) {
	src := NewMatWithSize(256, 256, MatTypeCV8UC1)
	defer src.Close()
	rot := GetRotationMatrix2D(image.Point{0, 0}, 1.0, 1.0)
	defer rot.Close()
	dst := src.Clone()
	defer dst.Close()

	WarpAffine(src, &dst, rot, image.Point{256, 256})
	result := Norm(dst, NormL2)
	if result != 0.0 {
		t.Errorf("WarpAffine() = %v, want %v", result, 0.0)
	}
}

func TestWarpAffineGocvLogo(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()
	rot := GetRotationMatrix2D(image.Point{0, 0}, 1.0, 1.0)
	defer rot.Close()
	dst := src.Clone()
	defer dst.Close()
	WarpAffine(src, &dst, rot, image.Point{343, 400})
	result := Norm(dst, NormL2)

	if !floatEquals(round(result, 0.05), round(111111.05, 0.05)) {
		t.Errorf("WarpAffine() = %v, want %v", round(result, 0.05), round(111111.05, 0.05))
	}
}

func TestWarpAffineWithParams(t *testing.T) {
	src := NewMatWithSize(256, 256, MatTypeCV8UC1)
	defer src.Close()
	rot := GetRotationMatrix2D(image.Point{0, 0}, 1.0, 1.0)
	defer rot.Close()
	dst := src.Clone()
	defer dst.Close()

	WarpAffineWithParams(src, &dst, rot, image.Point{256, 256}, InterpolationLinear, BorderConstant, color.RGBA{0, 0, 0, 0})
	result := Norm(dst, NormL2)
	if !floatEquals(result, 0.0) {
		t.Errorf("WarpAffineWithParams() = %v, want %v", result, 0.0)
	}
}

func TestWarpAffineWithParamsGocvLogo(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()
	rot := GetRotationMatrix2D(image.Point{0, 0}, 1.0, 1.0)
	defer rot.Close()
	dst := src.Clone()
	defer dst.Close()
	WarpAffineWithParams(src, &dst, rot, image.Point{343, 400}, InterpolationLinear, BorderConstant, color.RGBA{0, 0, 0, 0})
	result := Norm(dst, NormL2)
	if !floatEquals(round(result, 0.05), round(111111.05, 0.05)) {
		t.Errorf("WarpAffine() = %v, want %v", round(result, 0.05), round(111111.05, 0.05))
	}
}

func TestClipLine(t *testing.T) {

	if ok := ClipLine(image.Point{20, 20}, image.Point{5, 5}, image.Point{5, 5}); !ok {
		t.Error("ClipLine(): is false")
	}
}

func TestWatershed(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	if src.Empty() {
		t.Error("Invalid read of Mat in Watershed test")
	}
	defer src.Close()

	gray := NewMat()
	defer gray.Close()
	CvtColor(src, &gray, ColorBGRToGray)

	imgThresh := NewMat()
	defer imgThresh.Close()
	Threshold(gray, &imgThresh, 5, 50, ThresholdOtsu+ThresholdBinary)

	markers := NewMat()
	defer markers.Close()
	_ = ConnectedComponents(imgThresh, &markers)

	Watershed(src, &markers)
	if markers.Empty() || src.Cols() != markers.Cols() || src.Rows() != markers.Rows() {
		t.Error("Invalid Watershed test")
	}
}

func TestApplyColorMap(t *testing.T) {
	type args struct {
		colormapType ColormapTypes
		want         float64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "COLORMAP_AUTUMN", args: args{colormapType: ColormapAutumn, want: 118090.29593069873}},
		{name: "COLORMAP_BONE", args: args{colormapType: ColormapBone, want: 122067.44213343704}},
		{name: "COLORMAP_JET", args: args{colormapType: ColormapJet, want: 98220.64722857409}},
		{name: "COLORMAP_WINTER", args: args{colormapType: ColormapWinter, want: 94279.52859449394}},
		{name: "COLORMAP_RAINBOW", args: args{colormapType: ColormapRainbow, want: 92591.40608069411}},
		{name: "COLORMAP_OCEAN", args: args{colormapType: ColormapOcean, want: 106444.16919681415}},
		{name: "COLORMAP_SUMMER", args: args{colormapType: ColormapSummer, want: 114434.44957703952}},
		{name: "COLORMAP_SPRING", args: args{colormapType: ColormapSpring, want: 123557.60209715953}},
		{name: "COLORMAP_COOL", args: args{colormapType: ColormapCool, want: 123557.60209715953}},
		{name: "COLORMAP_HSV", args: args{colormapType: ColormapHsv, want: 107679.25179903508}},
		{name: "COLORMAP_PINK", args: args{colormapType: ColormapPink, want: 136043.97287274434}},
		{name: "COLORMAP_HOT", args: args{colormapType: ColormapHot, want: 124941.02475968412}},
		{name: "COLORMAP_PARULA", args: args{colormapType: ColormapParula, want: 111483.33555738274}},
	}
	src := IMRead("images/gocvlogo.jpg", IMReadGrayScale)
	defer src.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := src.Clone()
			defer dst.Close()
			ApplyColorMap(src, &dst, tt.args.colormapType)
			result := Norm(dst, NormL2)
			if !floatEquals(result, tt.args.want) {
				t.Errorf("TestApplyColorMap() = %v, want %v", result, tt.args.want)
			}
		})
	}
}

func TestApplyCustomColorMap(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadGrayScale)
	defer src.Close()
	customColorMap := NewMatWithSize(256, 1, MatTypeCV8UC1)
	defer customColorMap.Close()

	dst := src.Clone()
	defer dst.Close()
	ApplyCustomColorMap(src, &dst, customColorMap)
	result := Norm(dst, NormL2)
	if !floatEquals(result, 0.0) {
		t.Errorf("TestApplyCustomColorMap() = %v, want %v", result, 0.0)
	}
}

func TestGetPerspectiveTransform(t *testing.T) {
	src := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 5),
		image.Pt(10, 10),
		image.Pt(5, 10),
	}
	dst := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 0),
		image.Pt(10, 10),
		image.Pt(0, 10),
	}

	m := GetPerspectiveTransform(src, dst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestGetPerspectiveTransform(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 3 {
		t.Errorf("TestGetPerspectiveTransform(): unexpected rows = %v, want = %v", m.Rows(), 3)
	}
}

func TestGetPerspectiveTransform2f(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10.5, 5.5},
		{10.5, 10.5},
		{5.5, 10.5},
	}
	dst := []Point2f{
		{0, 0},
		{590.20, 24.12},
		{100.12, 150.21},
		{0, 10},
	}

	m := GetPerspectiveTransform2f(src, dst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestGetPerspectiveTransform2f(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 3 {
		t.Errorf("TestGetPerspectiveTransform2f(): unexpected rows = %v, want = %v", m.Rows(), 3)
	}
}

func TestGetAffineTransform(t *testing.T) {
	src := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 5),
		image.Pt(10, 10),
	}
	dst := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 0),
		image.Pt(10, 10),
	}

	m := GetAffineTransform(src, dst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestGetAffineTransform(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestGetAffineTransform(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}

func TestGetAffineTransform2f(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10.5, 5.5},
		{10.5, 10.5},
	}
	dst := []Point2f{
		{0, 0},
		{590.20, 24.12},
		{100.12, 150.21},
	}

	m := GetAffineTransform2f(src, dst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestGetAffineTransform2f(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestGetAffineTransform2f(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}

func TestFindHomography(t *testing.T) {
	src := NewMatWithSize(4, 1, MatTypeCV64FC2)
	defer src.Close()
	dst := NewMatWithSize(4, 1, MatTypeCV64FC2)
	defer dst.Close()

	srcPoints := []Point2f{
		{193, 932},
		{191, 378},
		{1497, 183},
		{1889, 681},
	}
	dstPoints := []Point2f{
		{51.51206544281359, -0.10425475260813055},
		{51.51211051314331, -0.10437947532732306},
		{51.512222354139325, -0.10437679311830816},
		{51.51214828037607, -0.1042212249954444},
	}

	for i, point := range srcPoints {
		src.SetDoubleAt(i, 0, float64(point.X))
		src.SetDoubleAt(i, 1, float64(point.Y))
	}

	for i, point := range dstPoints {
		dst.SetDoubleAt(i, 0, float64(point.X))
		dst.SetDoubleAt(i, 1, float64(point.Y))
	}

	mask := NewMat()
	defer mask.Close()

	m := FindHomography(src, &dst, HomograpyMethodAllPoints, 3, &mask, 2000, 0.995)
	defer m.Close()

	m2 := GetPerspectiveTransform2f(srcPoints, dstPoints)
	defer m2.Close()

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if math.Abs(m.GetDoubleAt(row, col)-m2.GetDoubleAt(row, col)) > 0.002 {
				t.Errorf("expected little difference between GetPerspectiveTransform2f and FindHomography results, got %f for row %d col %d", math.Abs(m.GetDoubleAt(row, col)-m2.GetDoubleAt(row, col)), row, col)
			}
		}
	}
}

func TestWarpPerspective(t *testing.T) {
	img := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer img.Close()

	w := img.Cols()
	h := img.Rows()

	s := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 5),
		image.Pt(10, 10),
		image.Pt(5, 10),
	}
	d := []image.Point{
		image.Pt(0, 0),
		image.Pt(10, 0),
		image.Pt(10, 10),
		image.Pt(0, 10),
	}
	m := GetPerspectiveTransform(s, d)
	defer m.Close()

	dst := NewMat()
	defer dst.Close()

	WarpPerspective(img, &dst, m, image.Pt(w, h))

	if dst.Cols() != w {
		t.Errorf("TestWarpPerspective(): unexpected cols = %v, want = %v", dst.Cols(), w)
	}

	if dst.Rows() != h {
		t.Errorf("TestWarpPerspective(): unexpected rows = %v, want = %v", dst.Rows(), h)
	}
}

func TestDrawContours(t *testing.T) {
	img := NewMatWithSize(100, 200, MatTypeCV8UC1)
	defer img.Close()

	// Draw rectangle
	white := color.RGBA{255, 255, 255, 255}
	Rectangle(&img, image.Rect(125, 25, 175, 75), white, 1)

	contours := FindContours(img, RetrievalExternal, ChainApproxSimple)

	if v := img.GetUCharAt(23, 123); v != 0 {
		t.Errorf("TestDrawContours(): wrong pixel value = %v, want = %v", v, 0)
	}
	if v := img.GetUCharAt(25, 125); v != 206 {
		t.Errorf("TestDrawContours(): wrong pixel value = %v, want = %v", v, 206)
	}

	DrawContours(&img, contours, -1, white, 2)

	// contour should be drawn with thickness = 2
	if v := img.GetUCharAt(24, 124); v != 255 {
		t.Errorf("TestDrawContours(): contour has not been drawn (value = %v, want = %v)", v, 255)
	}
	if v := img.GetUCharAt(25, 125); v != 255 {
		t.Errorf("TestDrawContours(): contour has not been drawn (value = %v, want = %v)", v, 255)
	}
}

func TestEllipse(t *testing.T) {
	img := NewMatWithSize(100, 100, MatTypeCV8UC1)
	defer img.Close()

	white := color.RGBA{255, 255, 255, 0}
	Ellipse(&img, image.Pt(50., 50.), image.Pt(25., 25.), 0., 0, 360, white, 2)

	if v := img.GetUCharAt(24, 50); v != 255 {
		t.Errorf("TestEllipse(): wrong pixel value = %v, want = %v", v, 255)
	}
}

func TestFillPoly(t *testing.T) {
	img := NewMatWithSize(100, 100, MatTypeCV8UC1)
	defer img.Close()

	white := color.RGBA{255, 255, 255, 0}
	pts := [][]image.Point{
		{
			image.Pt(10, 10),
			image.Pt(10, 20),
			image.Pt(20, 20),
			image.Pt(20, 10),
		},
	}
	FillPoly(&img, pts, white)

	if v := img.GetUCharAt(10, 10); v != 255 {
		t.Errorf("TestFillPoly(): wrong pixel value = %v, want = %v", v, 255)
	}
}

func TestPolylines(t *testing.T) {
	img := NewMatWithSize(100, 100, MatTypeCV8UC1)
	defer img.Close()

	white := color.RGBA{255, 255, 255, 0}
	pts := [][]image.Point{
		{
			image.Pt(10, 10),
			image.Pt(10, 20),
			image.Pt(20, 20),
			image.Pt(20, 10),
		},
	}
	Polylines(&img, pts, true, white, 1)

	if v := img.GetUCharAt(10, 10); v != 255 {
		t.Errorf("TestPolylines(): wrong pixel value = %v, want = %v", v, 255)
	}
}

func TestRemap(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()

	dst := NewMat()
	defer dst.Close()

	map1 := NewMatWithSize(256, 256, MatTypeCV16SC2)
	defer map1.Close()
	map1.SetFloatAt(50, 50, 25.4)
	map2 := NewMat()
	defer map2.Close()

	Remap(src, &dst, &map1, &map2, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0})

	if ok := dst.Empty(); ok {
		t.Errorf("Remap(): dst is empty")
	}
}

func TestFilter2D(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()

	dst := src.Clone()
	defer dst.Close()

	kernel := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernel.Close()

	Filter2D(src, &dst, -1, kernel, image.Pt(-1, -1), 0, BorderDefault)

	if ok := dst.Empty(); ok {
		t.Errorf("Filter2D(): dst is empty")
	}
}

func TestSepFilter2D(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()

	dst := src.Clone()
	defer dst.Close()

	kernelX := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernelX.Close()
	kernelY := GetStructuringElement(MorphRect, image.Pt(1, 1))
	defer kernelY.Close()

	SepFilter2D(src, &dst, -1, kernelX, kernelY, image.Pt(-1, -1), 0, BorderDefault)

	if ok := dst.Empty(); ok {
		t.Errorf("Filter2D(): dst is empty")
	}
}

func TestLogPolar(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()

	dst := src.Clone()
	defer dst.Close()

	LogPolar(src, &dst, image.Pt(22, 22), 1, InterpolationDefault)

	if ok := dst.Empty(); ok {
		t.Errorf("LogPolar(): dst is empty")
	}
}

func TestLinearPolar(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadUnchanged)
	defer src.Close()

	dst := src.Clone()
	defer dst.Close()

	LinearPolar(src, &dst, image.Pt(22, 22), 1, InterpolationDefault)

	if ok := dst.Empty(); ok {
		t.Errorf("LinearPolar(): dst is empty")
	}
}

func TestFitLine(t *testing.T) {
	points := []image.Point{image.Pt(125, 24), image.Pt(124, 75), image.Pt(175, 76), image.Pt(176, 25)}

	line := NewMat()
	defer line.Close()

	FitLine(points, &line, DistL2, 0, 0.01, 0.01)

	if ok := line.Empty(); ok {
		t.Errorf("FitLine(): line is empty")
	}
}

func TestInvertAffineTransform(t *testing.T) {
	src := NewMatWithSize(2, 3, MatTypeCV32F)
	defer src.Close()

	dst := NewMatWithSize(2, 3, MatTypeCV32F)
	defer dst.Close()

	InvertAffineTransform(src, &dst)

	if ok := dst.Empty(); ok {
		t.Errorf("InvertAffineTransform(): dst is empty")
	}
}

func TestCLAHE(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in NewCLAHE test")
	}
	defer img.Close()

	src := NewMat()
	defer src.Close()
	img.ConvertTo(&src, MatTypeCV8UC1)

	dst := NewMat()
	defer dst.Close()

	c := NewCLAHE()
	defer c.Close()
	c.Apply(src, &dst)
	if dst.Empty() || img.Rows() != dst.Rows() || img.Cols() != dst.Cols() {
		t.Error("Invalid NewCLAHE test")
	}
}

func TestCLAHEWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in CLAHEWithParams test")
	}
	defer img.Close()

	src := NewMat()
	defer src.Close()
	img.ConvertTo(&src, MatTypeCV8UC1)

	dst := NewMat()
	defer dst.Close()

	c := NewCLAHEWithParams(2.0, image.Pt(10, 10))
	defer c.Close()
	c.Apply(src, &dst)
	if dst.Empty() || img.Rows() != dst.Rows() || img.Cols() != dst.Cols() {
		t.Error("Invalid NewCLAHEWithParams test")
	}
}

func TestPhaseCorrelate(t *testing.T) {
	template := IMRead("images/simple.jpg", IMReadGrayScale)
	matched := IMRead("images/simple-translated.jpg", IMReadGrayScale)
	notMatchedOrig := IMRead("images/space_shuttle.jpg", IMReadGrayScale)
	notMatched := NewMat()

	defer template.Close()
	defer matched.Close()
	defer notMatchedOrig.Close()
	defer notMatched.Close()

	Resize(notMatchedOrig, &notMatched, image.Point{X: matched.Size()[0], Y: matched.Size()[1]}, 0, 0, InterpolationLinear)

	template32FC1 := NewMat()
	matched32FC1 := NewMat()
	notMatched32FC1 := NewMat()

	defer template32FC1.Close()
	defer matched32FC1.Close()
	defer notMatched32FC1.Close()

	template.ConvertTo(&template32FC1, MatTypeCV32FC1)
	matched.ConvertTo(&matched32FC1, MatTypeCV32FC1)
	notMatched.ConvertTo(&notMatched32FC1, MatTypeCV32FC1)

	window := NewMat()
	defer window.Close()

	shiftTranslated, responseTranslated := PhaseCorrelate(template32FC1, matched32FC1, window)
	_, responseDifferent := PhaseCorrelate(template32FC1, notMatched32FC1, window)

	if !(shiftTranslated.X < 15) || !(shiftTranslated.Y < 15) {
		t.Errorf("expected shift to be > 15 pixels, got %v", shiftTranslated)
	}

	if responseTranslated < 0.85 {
		t.Errorf("expected response for translated image to be > 0.85, got %f", responseTranslated)
	}

	if responseDifferent > 0.05 {
		t.Errorf("expected response for different image to be < 0.05, but got %f", responseDifferent)
	}
}
