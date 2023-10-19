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
	desc := gocv.NewMat()
	defer desc.Close()

	kpc := si.Compute(img, mask, kp, desc)
	if len(kpc) < 512 {
		t.Errorf("Invalid KeyPoint array in SURF Compute: %d", len(kpc))
	}
	if desc.Empty() {
		t.Error("Invalid Mat desc in SURF Compute")
	}

	kpdc, desc2 := si.DetectAndCompute(img, mask)
	defer desc2.Close()
	if len(kpdc) < 512 {
		t.Errorf("Invalid KeyPoint array in SURF DetectAndCompute: %d", len(kpdc))
	}
	if desc2.Empty() {
		t.Error("Invalid Mat desc in SURF DetectAndCompute")
	}
}

func TestSURFWithParams(t *testing.T) {
	testNonFree := os.Getenv("OPENCV_ENABLE_NONFREE")
	if testNonFree == "" {
		t.Skip("Skipping SURFWithParams test since OPENCV_ENABLE_NONFREE was not set")
	}

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in SURF test")
	}
	defer img.Close()

	dst := gocv.NewMat()
	defer dst.Close()

	si := NewSURFWithParams(100, 4, 3, false, false)
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

func TestBriefDescriptorExtractor(t *testing.T) {
	testNonFree := os.Getenv("OPENCV_ENABLE_NONFREE")
	if testNonFree == "" {
		t.Skip("Skipping BriefDescriptorExtractor test since OPENCV_ENABLE_NONFREE was not set")
	}

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in BriefDescriptorExtractor test")
	}
	defer img.Close()

	fast := gocv.NewFastFeatureDetector()
	defer fast.Close()

	b := NewBriefDescriptorExtractor()
	defer b.Close()

	kp := fast.Detect(img)

	desc := b.Compute(kp, img)

	if desc.Empty() {
		t.Error("Invalid Mat desc in BriefDescriptorExtractor Compute")
	}
}

func TestBriefDescriptorExtractorWithParams(t *testing.T) {
	testNonFree := os.Getenv("OPENCV_ENABLE_NONFREE")
	if testNonFree == "" {
		t.Skip("Skipping BriefDescriptorExtractorWithParams test since OPENCV_ENABLE_NONFREE was not set")
	}

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid Mat in BriefDescriptorExtractorWithParams test")
	}
	defer img.Close()

	fast := gocv.NewFastFeatureDetector()
	defer fast.Close()

	b := NewBriefDescriptorExtractorWithParams(32, false)
	defer b.Close()

	kp := fast.Detect(img)

	desc := b.Compute(kp, img)

	if desc.Empty() {
		t.Error("Invalid Mat desc in BriefDescriptorExtractorWithParams Compute")
	}
}
