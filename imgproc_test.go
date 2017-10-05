package gocv

import (
	"image"
	"image/color"
	"testing"
)

func TestCvtColor(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in CvtColor test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	CvtColor(img, dest, ColorBGRAToGray)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid convert in CvtColor test")
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

	Blur(img, dest, image.Pt(3, 3))
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
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

	GaussianBlur(img, dest, image.Pt(23, 23), 30, 50, 4)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid Blur test")
	}
}

func TestHoughLines(t *testing.T) {
	t.Skip("Test needed")
}

func TestHoughLinesP(t *testing.T) {
	t.Skip("Test needed")
}

func TestRectangle(t *testing.T) {
	t.Skip("Test needed")
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
	PutText(img, "Testing", pt, FontHersheyPlain, 1.2, color.RGBA{255, 255, 255, 0}, 2)

	if img.Empty() {
		t.Error("Error in PutText test")
	}
}
