package cuda

import (
	"image"
	"testing"

	"gocv.io/x/gocv"
)

func TestGaussianFilter_Apply(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in GaussianFilter test")
	}
	defer src.Close()

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	filter := NewGaussianFilter(src.Type(), src.Type(), image.Pt(23, 23), 30)
	defer filter.Close()

	dimg := filter.Apply(cimg)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty GaussianFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid GaussianFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid GaussianFilter test cols")
	}
}

func TestSobelFilter_Apply(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in SobelFilter test")
	}
	defer src.Close()

	cimg := NewGpuMat()
	defer cimg.Close()

	cimg.Upload(src)

	filter := NewSobelFilter(src.Type(), src.Type(), 0, 1)
	defer filter.Close()

	dimg := filter.Apply(cimg)
	defer dimg.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	dimg.Download(&dest)

	if dest.Empty() {
		t.Error("Empty SobelFilter test")
	}
	if src.Rows() != dest.Rows() {
		t.Error("Invalid SobelFilter test rows")
	}
	if src.Cols() != dest.Cols() {
		t.Error("Invalid SobelFilter test cols")
	}
}
