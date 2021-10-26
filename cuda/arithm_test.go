package cuda

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestAbs(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Abs test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dest := gocv.NewMat()
	defer dest.Close()

	Abs(cimg, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Abs test")
	}
}

func TestAbsWithStream(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Abs test")
	}
	defer src.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, s)
	AbsWithStream(cimg, &dimg, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Abs test")
	}
}

func TestThreshold(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Threshold test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dest := gocv.NewMat()
	defer dest.Close()

	Threshold(cimg, &dimg, 25, 255, gocv.ThresholdBinary)
	dimg.Download(&dest)
	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Threshold test")
	}
}

func TestThresholdWithStream(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Threshold test")
	}
	defer src.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, s)
	ThresholdWithStream(cimg, &dimg, 25, 255, gocv.ThresholdBinary, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Threshold test")
	}
}

func TestFlip(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Flip test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dest := gocv.NewMat()
	defer dest.Close()

	Flip(cimg, &dimg, 0)
	dimg.Download(&dest)
	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Flip test")
	}
}

func TestFlipWithStream(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Flip test")
	}
	defer src.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.UploadWithStream(src, s)
	FlipWithStream(cimg, &dimg, 0, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src.Rows() != dest.Rows() || src.Cols() != dest.Cols() {
		t.Error("Invalid Flip test")
	}
}
