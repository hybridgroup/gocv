package gocv

import (
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

	mog2.Apply(img, dst)

	if dst.Empty() {
		t.Error("Error in TestMOG2 test")
	}
}
