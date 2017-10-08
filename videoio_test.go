package gocv

import (
	"testing"
)

func TestVideoCaptureFile(t *testing.T) {
	vc, _ := VideoCaptureFile("images/small.mp4")
	defer vc.Close()

	if !vc.IsOpened() {
		t.Error("Unable to open VideoCaptureFile")
	}

	vc.Set(VideoCaptureBrightness, 100)
	vc.Grab(10)

	img := NewMat()
	defer img.Close()

	vc.Read(img)
	if img.Empty() {
		t.Error("Unable to read VideoCaptureFile")
	}
}

func TestVideoWriterFile(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in VideoWriterFile test")
	}
	defer img.Close()

	vw, _ := VideoWriterFile("/tmp/test.avi", "MJPG", 25, img.Cols(), img.Rows())
	defer vw.Close()

	if !vw.IsOpened() {
		t.Error("Unable to open VideoWriterFile")
	}

	err := vw.Write(img)
	if err != nil {
		t.Error("Invalid Write() in VideoWriter")
	}
}
