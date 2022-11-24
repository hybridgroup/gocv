package gocv

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"testing"
)

func TestFisheyeUndistorImage(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	FisheyeUndistortImage(img, &dest, k, d)

	if dest.Empty() {
		t.Error("final image is empty")
		return
	}
	// IMWrite("images/fisheye_sample-u.jpg", dest)
}

func TestFisheyeUndistorImageWithParams(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	knew := NewMat()
	defer knew.Close()

	k.CopyTo(&knew)

	knew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	size := image.Point{dest.Rows(), dest.Cols()}
	FisheyeUndistortImageWithParams(img, &dest, k, d, knew, size)

	if dest.Empty() {
		t.Error("final image is empty")
		return
	}
	// IMWrite("images/fisheye_sample-up.jpg", dest)
}

func TestInitUndistortRectifyMap(t *testing.T) {
	img := IMRead("images/distortion.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 842.0261028)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 667.7569792)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 707.3668897)
	k.SetDoubleAt(1, 2, 385.56476464)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 5, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, -3.65584802e-01)
	d.SetDoubleAt(0, 1, 1.41555815e-01)
	d.SetDoubleAt(0, 2, -2.62985819e-03)
	d.SetDoubleAt(0, 3, 2.05841873e-04)
	d.SetDoubleAt(0, 4, -2.35021914e-02)
	//FisheyeUndistortImage(img, &dest, k, d)
	//img.Reshape()
	newC, roi := GetOptimalNewCameraMatrixWithParams(k, d, image.Point{X: img.Cols(), Y: img.Rows()}, (float64)(1), image.Point{X: img.Cols(), Y: img.Rows()}, false)
	if newC.Empty() {
		t.Error("final image is empty")
		return
	}
	fmt.Printf("roi:%+v\n", roi)
	defer newC.Close()
	r := NewMat()
	defer r.Close()
	mapx := NewMat()
	defer mapx.Close()
	mapy := NewMat()
	defer mapy.Close()
	//dest := NewMat()
	InitUndistortRectifyMap(k, d, r, newC, image.Point{X: img.Cols(), Y: img.Rows()}, 5, mapx, mapy)

	Remap(img, &dest, &mapx, &mapy, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0})
	flg := IMWrite("images/distortion-correct.jpg", dest)
	if !flg {
		t.Error("IMWrite failed")
	}
}

func TestUndistort(t *testing.T) {
	img := IMRead("images/distortion.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := img.Clone()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	knew := NewMat()
	defer knew.Close()

	k.CopyTo(&knew)

	knew.SetDoubleAt(0, 0, 0.5*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.5*k.GetDoubleAt(1, 1))

	Undistort(img, &dest, k, d, knew)

	if dest.Empty() {
		t.Error("final image is empty")
		return
	}
	//IMWrite("images/distortion_up.jpg", dest)
}

func TestUndistortPoint(t *testing.T) {
	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 1094.7249578198823)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 959.4907612030962)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 1094.9945708128778)
	k.SetDoubleAt(1, 2, 536.4566143451868)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, -0.05207412392075069)
	d.SetDoubleAt(0, 1, -0.089168300192224)
	d.SetDoubleAt(0, 2, 0.10465607695792184)
	d.SetDoubleAt(0, 3, -0.045693446831115585)

	r := NewMat()
	defer r.Close()

	// transform 3 points in one go
	src := NewMatWithSize(3, 1, MatTypeCV64FC2)
	defer src.Close()
	dst := NewMatWithSize(3, 1, MatTypeCV64FC2)
	defer dst.Close()

	// This camera matrix is 1920x1080. Points where x < 960 and y < 540 should move toward the top left (x and y get smaller)
	// The centre point should be mostly unchanged
	// Points where x > 960 and y > 540 should move toward the bottom right (x and y get bigger)

	// The index being used for col here is actually the channel (i.e. the point's x/y dimensions)
	// (since there's only 1 column so the formula: (colNumber * numChannels + channelNumber) reduces to
	// (0 * 2) + channelNumber
	// so col = 0 is the x coordinate and col = 1 is the y coordinate

	src.SetDoubleAt(0, 0, 480)
	src.SetDoubleAt(0, 1, 270)

	src.SetDoubleAt(1, 0, 960)
	src.SetDoubleAt(1, 1, 540)

	src.SetDoubleAt(2, 0, 1920)
	src.SetDoubleAt(2, 1, 1080)

	UndistortPoints(src, &dst, k, d, r, k)

	if dst.GetDoubleAt(0, 0) >= 480 || dst.GetDoubleAt(0, 1) >= 270 {
		t.Error("undistortion expected top left point to move further up and left")
		return
	}

	if math.Round(dst.GetDoubleAt(1, 0)) != 960 || math.Round(dst.GetDoubleAt(1, 1)) != 540 {
		t.Error("undistortion expected centre point to be nearly unchanged")
		return
	}

	if dst.GetDoubleAt(2, 0) != 1920 || dst.GetDoubleAt(2, 1) != 1080 {
		t.Error("undistortion expected bottom right corner to be unchanged")
		return
	}

}

func TestFisheyeUndistortPoint(t *testing.T) {
	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 1094.7249578198823)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 959.4907612030962)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 1094.9945708128778)
	k.SetDoubleAt(1, 2, 536.4566143451868)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, -0.05207412392075069)
	d.SetDoubleAt(0, 1, -0.089168300192224)
	d.SetDoubleAt(0, 2, 0.10465607695792184)
	d.SetDoubleAt(0, 3, -0.045693446831115585)

	r := NewMat()
	defer r.Close()

	// transform 3 points in one go (X and Y values of points go in each channel)
	src := NewMatWithSize(3, 1, MatTypeCV64FC2)
	defer src.Close()
	dst := NewMatWithSize(3, 1, MatTypeCV64FC2)
	defer dst.Close()

	// This camera matrix is 1920x1080. Points where x < 960 and y < 540 should move toward the top left (x and y get smaller)
	// The centre point should be mostly unchanged
	// Points where x > 960 and y > 540 should move toward the bottom right (x and y get bigger)

	// The index being used for col here is actually the channel (i.e. the point's x/y dimensions)
	// (since there's only 1 column so the formula: (colNumber * numChannels + channelNumber) reduces to
	// (0 * 2) + channelNumber
	// so col = 0 is the x coordinate and col = 1 is the y coordinate

	src.SetDoubleAt(0, 0, 480)
	src.SetDoubleAt(0, 1, 270)

	src.SetDoubleAt(1, 0, 960)
	src.SetDoubleAt(1, 1, 540)

	src.SetDoubleAt(2, 0, 1440)
	src.SetDoubleAt(2, 1, 810)

	kNew := NewMat()
	defer kNew.Close()

	k.CopyTo(&kNew)

	kNew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	kNew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	imgSize := image.Point{X: 1920, Y: 1080}

	EstimateNewCameraMatrixForUndistortRectify(k, d, imgSize, r, &kNew, 1, imgSize, 1)

	FisheyeUndistortPoints(src, &dst, k, d, r, kNew)

	if dst.GetDoubleAt(0, 0) == 0 {
		t.Error("expected destination Mat to be populated")
	}

}
func TestFindAndDrawChessboard(t *testing.T) {
	img := IMRead("images/chessboard_4x6.png", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of chessboard image")
		return
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	found := FindChessboardCorners(img, image.Point{X: 4, Y: 6}, &corners, 0)
	if found == false {
		t.Error("chessboard pattern not found")
		return
	}
	if corners.Empty() {
		t.Error("chessboard pattern not found")
		return
	}

	img2 := NewMatWithSize(150, 150, MatTypeCV8U)
	defer img2.Close()

	DrawChessboardCorners(&img2, image.Pt(4, 6), corners, true)
	if img2.Empty() {
		t.Error("Error in DrawChessboardCorners test")
	}
}

func TestFindAndDrawChessboardSB(t *testing.T) {
	img := IMRead("images/chessboard_4x6.png", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of chessboard image")
		return
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	found := FindChessboardCornersSB(img, image.Point{X: 4, Y: 6}, &corners, 0)
	if found == false {
		t.Error("chessboard pattern not found")
		return
	}
	if corners.Empty() {
		t.Error("chessboard pattern not found")
		return
	}

	img2 := NewMatWithSize(150, 150, MatTypeCV8U)
	defer img2.Close()

	DrawChessboardCorners(&img2, image.Pt(4, 6), corners, true)
	if img2.Empty() {
		t.Error("Error in DrawChessboardCorners test")
	}
}

func TestFindChessboardCornersSBWithMeta(t *testing.T) {
	img := IMRead("images/chessboard_4x6.png", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of chessboard image")
		return
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	meta := NewMat()
	defer meta.Close()

	found := FindChessboardCornersSBWithMeta(img, image.Point{X: 4, Y: 6}, &corners, 0, &meta)
	if found == false {
		t.Error("chessboard pattern not found")
		return
	}
	if corners.Empty() {
		t.Error("chessboard pattern not found")
		return
	}

	img2 := NewMatWithSize(150, 150, MatTypeCV8U)
	defer img2.Close()

	DrawChessboardCorners(&img2, image.Pt(4, 6), corners, true)
	if img2.Empty() {
		t.Error("Error in DrawChessboardCorners test")
	}
}

func TestCalibrateCamera(t *testing.T) {
	img := IMRead("images/chessboard_4x6_distort.png", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of chessboard image")
		return
	}
	defer img.Close()

	corners := NewMat()
	defer corners.Close()

	size := image.Pt(4, 6)
	found := FindChessboardCorners(img, size, &corners, 0)
	if !found {
		t.Error("chessboard pattern not found")
		return
	}
	if corners.Empty() {
		t.Error("chessboard pattern not found")
		return
	}

	imagePoints := NewPoint2fVectorFromMat(corners)
	defer imagePoints.Close()

	objectPoints := NewPoint3fVector()
	defer objectPoints.Close()

	for j := 0; j < size.Y; j++ {
		for i := 0; i < size.X; i++ {
			objectPoints.Append(Point3f{
				X: float32(100 * i),
				Y: float32(100 * j),
				Z: 0,
			})
		}
	}

	cameraMatrix := NewMat()
	defer cameraMatrix.Close()
	distCoeffs := NewMat()
	defer distCoeffs.Close()
	rvecs := NewMat()
	defer rvecs.Close()
	tvecs := NewMat()
	defer tvecs.Close()

	objectPointsVector := NewPoints3fVector()
	objectPointsVector.Append(objectPoints)
	defer objectPointsVector.Close()

	imagePointsVector := NewPoints2fVector()
	imagePointsVector.Append(imagePoints)
	defer imagePointsVector.Close()

	CalibrateCamera(
		objectPointsVector, imagePointsVector, image.Pt(img.Cols(), img.Rows()),
		&cameraMatrix, &distCoeffs, &rvecs, &tvecs, 0,
	)

	dest := NewMat()
	defer dest.Close()
	Undistort(img, &dest, cameraMatrix, distCoeffs, cameraMatrix)

	target := IMRead("images/chessboard_4x6_distort_correct.png", IMReadGrayScale)
	defer target.Close()

	xor := NewMat()
	defer xor.Close()

	// The method for compare is ugly : different pix number < 0.5%
	BitwiseXor(dest, target, &xor)
	differentPixelsNumber := xor.Sum().Val1
	maxDifferentPixelsNumber := float64(img.Cols()*img.Rows()) * 0.005
	if differentPixelsNumber > maxDifferentPixelsNumber {
		t.Error("the undisorted image not equal the target one:", differentPixelsNumber, "bigger than", maxDifferentPixelsNumber)
	}
}

func TestEstimateAffinePartial2D(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10, 5},
		{10, 10},
		{5, 10},
	}

	dst := []Point2f{
		{0, 0},
		{10, 0},
		{10, 10},
		{0, 10},
	}

	pvsrc := NewPoint2fVectorFromPoints(src)
	defer pvsrc.Close()

	pvdst := NewPoint2fVectorFromPoints(dst)
	defer pvdst.Close()

	m := EstimateAffinePartial2D(pvsrc, pvdst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestEstimateAffinePartial2D(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestEstimateAffinePartial2D(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}

func TestEstimateAffinePartial2DWithParams(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10, 5},
		{10, 10},
		{5, 10},
	}

	dst := []Point2f{
		{0, 0},
		{10, 0},
		{10, 10},
		{0, 10},
	}

	pvsrc := NewPoint2fVectorFromPoints(src)
	defer pvsrc.Close()

	pvdst := NewPoint2fVectorFromPoints(dst)
	defer pvdst.Close()

	inliers := NewMat()
	defer inliers.Close()
	method := 8
	ransacProjThreshold := 3.0
	maxiters := uint(2000)
	confidence := 0.99
	refineIters := uint(10)

	m := EstimateAffinePartial2DWithParams(pvsrc, pvdst, inliers, method, ransacProjThreshold, maxiters, confidence, refineIters)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestEstimateAffinePartial2D(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestEstimateAffinePartial2D(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}

func TestEstimateAffine2D(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10, 5},
		{10, 10},
		{5, 10},
	}

	dst := []Point2f{
		{0, 0},
		{10, 0},
		{10, 10},
		{0, 10},
	}

	pvsrc := NewPoint2fVectorFromPoints(src)
	defer pvsrc.Close()

	pvdst := NewPoint2fVectorFromPoints(dst)
	defer pvdst.Close()

	m := EstimateAffine2D(pvsrc, pvdst)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestEstimateAffine2D(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestEstimateAffine2D(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}

func TestEstimateAffine2DWithParams(t *testing.T) {
	src := []Point2f{
		{0, 0},
		{10, 5},
		{10, 10},
		{5, 10},
	}

	dst := []Point2f{
		{0, 0},
		{10, 0},
		{10, 10},
		{0, 10},
	}

	pvsrc := NewPoint2fVectorFromPoints(src)
	defer pvsrc.Close()

	pvdst := NewPoint2fVectorFromPoints(dst)
	defer pvdst.Close()

	inliers := NewMat()
	defer inliers.Close()
	method := 8
	ransacProjThreshold := 3.0
	maxiters := uint(2000)
	confidence := 0.99
	refineIters := uint(10)

	m := EstimateAffine2DWithParams(pvsrc, pvdst, inliers, method, ransacProjThreshold, maxiters, confidence, refineIters)
	defer m.Close()

	if m.Cols() != 3 {
		t.Errorf("TestEstimateAffine2DWithParams(): unexpected cols = %v, want = %v", m.Cols(), 3)
	}
	if m.Rows() != 2 {
		t.Errorf("TestEstimateAffine2DWithParams(): unexpected rows = %v, want = %v", m.Rows(), 2)
	}
}
