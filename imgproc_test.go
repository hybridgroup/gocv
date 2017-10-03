package gocv

import (
	"testing"
)

func TestCvtColor(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in CvtColor test")
	}

	dest := NewMat()
	defer dest.Close()

	CvtColor(img, dest, ColorBGRAToGray)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Invalid convert in CvtColor test")
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
	t.Skip("Test needed")
}
