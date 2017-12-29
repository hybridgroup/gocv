package gocv

import (
	"image"
	"os"
	"testing"
)

func TestCaffe(t *testing.T) {
	path := os.Getenv("GOCV_CAFFE_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Caffe model files for tests")
	}

	net := ReadNetFromCaffe(path+"/bvlc_googlenet.prototxt", path+"/bvlc_googlenet.caffemodel")
	if net.Empty() {
		t.Errorf("Unable to load Caffe model")
	}

	img := IMRead("images/space_shuttle.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in Caffe test")
	}
	defer img.Close()

	blob := BlobFromImage(img, 1.0, image.Pt(224, 244), NewScalar(104, 117, 123, 0), false, false)
	if blob.Empty() {
		t.Error("Invalid blob in Caffe test")
	}
	defer blob.Close()

	net.SetInput(blob, "data")
	prob := net.Forward("prob")
	if prob.Empty() {
		t.Error("Invalid prob in Caffe test")
	}

	probMat := prob.Reshape(1, 1)
	_, maxVal, minLoc, maxLoc := MinMaxLoc(probMat)

	if round(float64(maxVal), 0.00005) != 0.9999 {
		t.Errorf("Caffe maxVal incorrect: %v\n", round(float64(maxVal), 0.00005))
	}

	if minLoc.X != 793 || minLoc.Y != 0 {
		t.Errorf("Caffe minLoc incorrect: %v\n", minLoc)
	}

	if maxLoc.X != 812 || maxLoc.Y != 0 {
		t.Errorf("Caffe maxLoc incorrect: %v\n", maxLoc)
	}
}
