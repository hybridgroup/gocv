package contrib

import (
	"os"
	"testing"

	"gocv.io/x/gocv"
)

func TestSURF(t *testing.T) {
	testNonFree := os.Getenv("OPENCV_ENABLE_NONFREE")
	if testNonFree == "" {
		t.Skip("Skipping SURF test since OPENCV_ENABLE_NONFREE was not set")
	}

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in SURF test")
	}
	defer img.Close()

	dst := gocv.NewMat()
	defer dst.Close()

	si := NewSURF()
	defer si.Close()

	kp := si.Detect(img)
	if len(kp) == 512 {
		t.Errorf("Invalid KeyPoint array in SURF Detect: %d", len(kp))
	}

	mask := gocv.NewMat()
	defer mask.Close()

	kp2, desc := si.DetectAndCompute(img, mask)
	if len(kp2) == 512 {
		t.Errorf("Invalid KeyPoint array in SURF DetectAndCompute: %d", len(kp2))
	}

	if desc.Empty() {
		t.Error("Invalid Mat desc in SURF DetectAndCompute")
	}
}
