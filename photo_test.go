package gocv

import (
	"testing"
)

func TestFastNlMeansDenoising(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in TestFastNlMeansDenoising test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoising(img, &dest)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in TestFastNlMeansDenoising test")
	}
}

func TestFastNlMeansDenoisingWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in TestFastNlMeansDenoising test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingWithParams(img, &dest, 3, 7, 21)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in TestFastNlMeansDenoising test")
	}
}

func TestFastNlMeansDenoisingColored(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in FastNlMeansDenoisingColored test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingColored(img, &dest)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in FastNlMeansDenoisingColored test")
	}
}

func TestFastNlMeansDenoisingColoredWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in FastNlMeansDenoisingColored test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingColoredWithParams(img, &dest, 3, 3, 7, 21)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in FastNlMeansDenoisingColored test")
	}
}
