package contrib

import (
	"gocv.io/x/gocv"
	"testing"
)

func TestAnisotropicDiffusion(t *testing.T) {
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC3)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	AnisotropicDiffusion(src, &dst, 0.5, 0.5, 100)

	if src.Empty() || dst.Rows() != src.Rows() {
		t.Error("invalid AnisotropicDiffusion test")
	}
}

func TestEdgePreservingFilter(t *testing.T) {
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC3)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	EdgePreservingFilter(src, &dst, 3, 0.5)

	if src.Empty() || dst.Rows() != src.Rows() {
		t.Error("invalid EdgePreservingFilter test")
	}
}

func TestNiblackThreshold(t *testing.T) {
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC1)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	NiblackThreshold(src, &dst, 127.0, gocv.ThresholdBinary, 3, 0.5, BinarizationNiblack, 128)

	if src.Empty() || dst.Rows() != src.Rows() {
		t.Error("invalid NiblackThreshold test")
	}
}

func TestPeiLinNormalization(t *testing.T) {
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC1)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	PeiLinNormalization(src, &dst)

	if src.Empty() || dst.Rows() != 2 || dst.Cols() != 3 {
		t.Error("invalid PeiLinNormalization test", dst.Rows(), dst.Cols())
	}
}

func TestThinning(t *testing.T) {
	src := gocv.NewMatWithSize(200, 200, gocv.MatTypeCV8UC1)
	defer src.Close()
	dst := gocv.NewMat()
	defer dst.Close()

	Thinning(src, &dst, ThinningZhangSuen)

	if src.Empty() || dst.Rows() != src.Rows() {
		t.Error("invalid Thinning test")
	}
}
