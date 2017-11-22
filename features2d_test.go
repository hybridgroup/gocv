package gocv

import (
	"testing"
)

func TestSimpleBlobDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in SimpleBlobDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	bd := NewSimpleBlobDetector()
	defer bd.Close()

	bd.Detect(img)
}
