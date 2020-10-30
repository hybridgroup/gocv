package gocv

import (
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
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
	vc, err := OpenVideoCapture("images/small.mp4")
	if err != nil {
		t.Errorf("TestVideoCaptureCodecString: error loading a file: %v", err)
	}
	if vc.CodecString() == "" {
		t.Fatal("TestVideoCaptureCodecString: empty codec string")
	}
}

func TestVideoCaptureCodecConversion(t *testing.T) {
	vc, err := OpenVideoCapture("images/small.mp4")
	if err != nil {
		t.Errorf("TestVideoCaptureCodecConversion: error loading a file: %v", err)
	}
	if vc.CodecString() == "" {
		t.Fatal("TestVideoCaptureCodecConversion: empty codec string")
	}
	if int64(vc.ToCodec(vc.CodecString())) != int64(vc.Get(VideoCaptureFOURCC)) {
		t.Fatal("TestVideoCaptureCodecConversion: codec conversion failed")
	}
}

func TestVideoCaptureCodecConversionBadInput(t *testing.T) {
	vc, err := OpenVideoCapture("images/small.mp4")
	if err != nil {
		t.Errorf("TestVideoCaptureCodecString: error loading a file: %v", err)
	}
	codec := vc.ToCodec("BAD CODEC")
	if int64(codec) != -1 {
		t.Fatal("TestVideoCaptureCodecConversionBadInput: input validation failed")
	}
}

func TestVideoCaptureInvalid(t *testing.T) {
	_, err := OpenVideoCapture(1.1)
	if err == nil {
		t.Errorf("Should return error with invalid param")
	}
}

func TestVideoCaptureWithAPI(t *testing.T) {
	t.Run("video capture file with api", func(t *testing.T) {
		vc, err := OpenVideoCaptureWithAPI("images/small.mp4", VideoCaptureAny)
		if err != nil {
			t.Errorf("error loading a file: %v", err)
		}
		backend := vc.Get(VideoCaptureBackend)
		if backend == float64(VideoCaptureAny) {
			t.Errorf("video capture backend api did not select a backend")
		}
	})

	t.Run("video capture unknown device with api", func(t *testing.T) {
		_, err := OpenVideoCaptureWithAPI(math.MaxInt32, VideoCaptureAny)
		if err == nil {
			t.Errorf("should return error opening device")
		}
	})

	t.Run("video capture invalid with api", func(t *testing.T) {
		_, err := OpenVideoCaptureWithAPI(1.1, VideoCaptureAny)
		if err == nil {
			t.Errorf("should return error with invalid param")
		}
	})

	t.Run("video capture valid int string with api", func(t *testing.T) {
		vc5, err := OpenVideoCaptureWithAPI("1", VideoCaptureAny)
		defer vc5.Close()
		if err == nil {
			t.Errorf("should return error opening device")
		}
	})
}

func TestVideoCaptureFile(t *testing.T) {
	vc, err := VideoCaptureFile("images/small.mp4")
	defer vc.Close()

	if err != nil {
		t.Errorf("%s", err)
	}

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

	// video capture file with non-existent video
	vc2, err := VideoCaptureFile("nonexistent.mp4")
	defer vc2.Close()

	if err == nil {
		t.Errorf("Expected error when opening invalid file")
	}

	t.Run(" video capture file with api", func(t *testing.T) {
		vc3, err := VideoCaptureFileWithAPI("images/small.mp4", VideoCaptureAny)
		defer vc3.Close()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("video capture non-existent video with api", func(t *testing.T) {
		vc4, err := VideoCaptureFileWithAPI("nonexistent.mp4", VideoCaptureAny)
		defer vc4.Close()
		if err == nil {
			t.Errorf("Expected error when opening invalid file")
		}
	})

	t.Run("video capture invalid int", func(t *testing.T) {
		vc5, err := OpenVideoCapture(math.MaxInt32)
		defer vc5.Close()
		if err == nil {
			t.Errorf("should return error opening device")
		}
	})

	t.Run("video capture invalid string", func(t *testing.T) {
		vc5, err := OpenVideoCapture("test-device")
		defer vc5.Close()
		if err == nil {
			t.Errorf("should return error opening device")
		}
	})

	t.Run("video capture valid string", func(t *testing.T) {
		vc5, err := OpenVideoCapture(strconv.Itoa(math.MaxInt32))
		defer vc5.Close()
		if err == nil {
			t.Errorf("should return error opening device")
		}
	})
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
