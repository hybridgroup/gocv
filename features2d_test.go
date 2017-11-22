package gocv

import (
	"testing"
)

func TestAgastFeatureDetector(t *testing.T) {
	img := IMRead("images/face.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in AgastFeatureDetector test")
	}
	defer img.Close()

	dst := NewMat()
	defer dst.Close()

	ad := NewAgastFeatureDetector()
	defer ad.Close()

	kp := ad.Detect(img)
	if len(kp) != 2800 {
		t.Errorf("Invalid KeyPoint array in AgastFeatureDetector test: %d", len(kp))
	}
}

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

	kp := bd.Detect(img)
	if len(kp) != 2 {
		t.Errorf("Invalid KeyPoint array in SimpleBlobDetector test: %d", len(kp))
	}
}
