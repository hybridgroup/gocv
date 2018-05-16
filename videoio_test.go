package gocv

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestVideoCaptureEmptyNumericalParameters(t *testing.T) {
	_, err := VideoWriterFile(
		"images/small.mp4", "MJPEG", 0, 0, 0, true)
	if err == nil {
		t.Error("Must fail due to an empty numerical parameters.")
	}
	if !strings.Contains(err.Error(), "one of the numerical parameters is equal to zero") {
		t.Errorf("Must fail due to an empty numerical "+
			"parameters, but have different error: %v", err)
	}
}

func TestVideoCaptureCodecString(t *testing.T) {
	vc, err := VideoCaptureFile("images/small.mp4")
	if err != nil {
		t.Errorf("TestVideoCaptureCodecString: error loading a file: %v", err)
	}
	if vc.CodecString() == "" {
		t.Fatal("TestVideoCaptureCodecString: empty codec string")
	}
}

func TestVideoCaptureFile(t *testing.T) {
	vc, _ := VideoCaptureFile("images/small.mp4")
	defer vc.Close()

	if !vc.IsOpened() {
		t.Error("Unable to open VideoCaptureFile")
	}

	if fw := vc.Get(VideoCaptureFrameWidth); int(fw) != 560 {
		t.Errorf("Expected frame width property of 560.0 got %f", fw)
	}
	if fh := vc.Get(VideoCaptureFrameHeight); int(fh) != 320 {
		t.Errorf("Expected frame height property of 320.0 got %f", fh)
	}

	vc.Set(VideoCaptureBrightness, 100.0)

	vc.Grab(10)

	img := NewMat()
	defer img.Close()

	vc.Read(&img)
	if img.Empty() {
		t.Error("Unable to read VideoCaptureFile")
	}
}

func TestVideoWriterFile(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gocvtests")
	tmpfn := filepath.Join(dir, "test.avi")

	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in VideoWriterFile test")
	}
	defer img.Close()

	vw, _ := VideoWriterFile(tmpfn, "MJPG", 25, img.Cols(), img.Rows(), true)
	defer vw.Close()

	if !vw.IsOpened() {
		t.Error("Unable to open VideoWriterFile")
	}

	err := vw.Write(img)
	if err != nil {
		t.Error("Invalid Write() in VideoWriter")
	}
}
