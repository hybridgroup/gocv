package gocv

import (
	"image"
	"image/color"
	"math"
	"testing"
)

func TestMOG2(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in MOG2 test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	mog2.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestMOG2 test")
	}
}

func TestMOG2WithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in MOG2 test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2WithParams(250, 8, false)
	defer mog2.Close()

	mog2.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestMOG2WithParams test")
	}
}

func TestKNN(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in KNN test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	knn := NewBackgroundSubtractorKNN()
	defer knn.Close()

	knn.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestKNN test")
	}
}

func TestKNNWithParams(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in KNN test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	knn := NewBackgroundSubtractorKNNWithParams(250, 200, false)
	defer knn.Close()

	knn.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestKNNWithParams test")
	}
}

func TestCalcOpticalFlowFarneback(t *testing.T) {
	img1 := IMRead("images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowFarneback test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	flow := NewMat()
	defer flow.Close()

	CalcOpticalFlowFarneback(dest, img2, &flow, 0.4, 1, 12, 2, 8, 1.2, 0)

	if flow.Empty() {
		t.Error("Error in CalcOpticalFlowFarneback test")
	}
	if flow.Rows() != 480 {
		t.Errorf("Invalid CalcOpticalFlowFarneback test rows: %v", flow.Rows())
	}
	if flow.Cols() != 640 {
		t.Errorf("Invalid CalcOpticalFlowFarneback test cols: %v", flow.Cols())
	}
}

func TestCalcOpticalFlowPyrLK(t *testing.T) {
	img1 := IMRead("images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowPyrLK test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	prevPts := NewMat()
	defer prevPts.Close()

	nextPts := NewMat()
	defer nextPts.Close()

	status := NewMat()
	defer status.Close()

	err := NewMat()
	defer err.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(dest, &corners, 500, 0.01, 10)
	tc := NewTermCriteria(Count|EPS, 20, 0.03)
	CornerSubPix(dest, &corners, image.Pt(10, 10), image.Pt(-1, -1), tc)

	CalcOpticalFlowPyrLK(dest, img2, corners, nextPts, &status, &err)

	if status.Empty() {
		t.Error("Error in CalcOpticalFlowPyrLK test")
	}
	if status.Rows() != 323 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test rows: %v", status.Rows())
	}
	if status.Cols() != 1 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test cols: %v", status.Cols())
	}
}

func TestCalcOpticalFlowPyrLKWithParams(t *testing.T) {
	img1 := IMRead("images/face.jpg", IMReadColor)
	if img1.Empty() {
		t.Error("Invalid Mat in CalcOpticalFlowPyrLK test")
	}
	defer img1.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img1, &dest, ColorBGRAToGray)

	img2 := dest.Clone()
	defer img2.Close()

	prevPts := NewMat()
	defer prevPts.Close()

	nextPts := NewMat()
	defer nextPts.Close()

	status := NewMat()
	defer status.Close()

	err := NewMat()
	defer err.Close()

	corners := NewMat()
	defer corners.Close()

	GoodFeaturesToTrack(dest, &corners, 500, 0.01, 10)
	tc := NewTermCriteria(Count|EPS, 30, 0.03)
	CornerSubPix(dest, &corners, image.Pt(10, 10), image.Pt(-1, -1), tc)

	CalcOpticalFlowPyrLKWithParams(dest, img2, corners, nextPts, &status, &err, image.Pt(21, 21), 3, tc, 0, 0.0001)

	if status.Empty() {
		t.Error("Error in CalcOpticalFlowPyrLK test")
	}
	if status.Rows() != 323 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test rows: %v", status.Rows())
	}
	if status.Cols() != 1 {
		t.Errorf("Invalid CalcOpticalFlowPyrLK test cols: %v", status.Cols())
	}
}

func computeRMS(mat1 Mat, mat2 Mat) float64 {
	var rms float64
	for y := 0; y < mat1.Rows(); y++ {
		for x := 0; x < mat1.Cols(); x++ {
			diff := float64(mat1.GetFloatAt(y, x) - mat2.GetFloatAt(y, x))
			rms += diff * diff
		}
	}

	rms /= float64(mat1.Rows() * mat1.Cols())
	return math.Sqrt(rms)
}

func TestFindTransformECC(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in FindTransformECC test")
	}
	defer img.Close()
	testImg := NewMat()
	defer testImg.Close()
	Resize(img, &testImg, image.Point{216, 216}, 0, 0, InterpolationLinear)

	translationGround := Eye(2, 3, MatTypeCV32F)
	defer translationGround.Close()
	translationGround.SetFloatAt(0, 2, 11.4159)
	translationGround.SetFloatAt(1, 2, 17.1828)

	warpedImage := NewMat()
	defer warpedImage.Close()
	WarpAffineWithParams(testImg, &warpedImage, translationGround, image.Point{200, 200}, InterpolationLinear+WarpInverseMap, BorderConstant, color.RGBA{})

	mapTranslation := Eye(2, 3, MatTypeCV32F)
	defer mapTranslation.Close()
	eecIterations := 50
	// Negative value means that ECC_Iterations will be executed.
	var eecEpsilon float64 = -1
	criteria := NewTermCriteria(Count+EPS, eecIterations, eecEpsilon)
	inputMask := NewMat()
	defer inputMask.Close()
	gaussFiltSize := 5
	FindTransformECC(warpedImage, testImg, &mapTranslation, MotionTranslation, criteria, inputMask, gaussFiltSize)

	maxRMSECC := 0.1
	rms := computeRMS(mapTranslation, translationGround)
	if rms > maxRMSECC {
		t.Errorf("FindTransformECC RMS = %f", rms)
	}
}

func BaseTestTracker(t *testing.T, tracker Tracker, name string) {
	if tracker == nil {
		t.Error("TestTracker " + name + " should not be nil")
	}

	img := IMRead("./images/face.jpg", 1)
	if img.Empty() {
		t.Error("TestTracker " + name + " input img failed to load")
	}
	defer img.Close()

	rect := image.Rect(250, 150, 250+200, 150+250)
	init := tracker.Init(img, rect)
	if !init {
		t.Error("TestTracker " + name + " failed in Init")
	}

	_, ok := tracker.Update(img)
	if !ok {
		t.Error("TestTracker " + name + " lost object in Update")
	}
}

func TestSingleTrackers(t *testing.T) {
	tab := []struct {
		name    string
		tracker Tracker
	}{
		{"MIL", NewTrackerMIL()},
		// {"GOTURN", NewTrackerGOTURN()},
	}

	for _, test := range tab {
		func() {
			defer test.tracker.Close()
			BaseTestTracker(t, test.tracker, test.name)
		}()
	}
}
