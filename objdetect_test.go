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

func TestHOGDescriptor(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in HOGDescriptor test")
	}
	defer img.Close()

	// load HOGDescriptor to recognize people
	hog := NewHOGDescriptor()
	defer hog.Close()

	hog.SetSVMDetector(HOGDefaultPeopleDetector())

	rects := hog.DetectMultiScale(img)
	if len(rects) != 1 {
		t.Errorf("Error in TestHOGDescriptor test: %d", len(rects))
	}
}
