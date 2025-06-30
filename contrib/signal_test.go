package contrib

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestResampleSignal(t *testing.T) {

	img := gocv.IMRead("../images/face.jpg", gocv.IMReadAnyColor)
	if img.Empty() {
		t.Error("xobjdetect: cannot read image")
	}

	resampledMat := ResampleSignal(img, 1000, 1000)
	defer resampledMat.Close()

	if resampledMat.Empty() {
		t.Error("contrib.signal.ResampleSignal: empty output")
	}

}
