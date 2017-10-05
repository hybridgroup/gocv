package gocv

import (
	"testing"
)

func TestCascadeClassifier(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CascadeClassifier test")
	}
	defer img.Close()

	// load classifier to recognize faces
	classifier := NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("data/haarcascade_frontalface_default.xml")

	rects := classifier.DetectMultiScale(img)
	if len(rects) != 1 {
		t.Error("Error in TestCascadeClassifier test")
	}
}
