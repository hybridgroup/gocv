package cuda

import (
	"image"
	"image/color"
	"math"
	"testing"

	"gocv.io/x/gocv"
)

func TestResize(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in Resize test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dst := gocv.NewMat()
	defer dst.Close()

	Resize(cimg, &dimg, image.Point{}, 0.5, 0.5, InterpolationDefault)
	dimg.Download(&dst)
	if dst.Cols() != 200 || dst.Rows() != 172 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}

	Resize(cimg, &dimg, image.Pt(440, 377), 0, 0, InterpolationCubic)
	dimg.Download(&dst)
	if dst.Cols() != 440 || dst.Rows() != 377 {
		t.Errorf("Expected dst size of 440x377 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestPyrDown(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrDown test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := gocv.NewMat()
	defer dst.Close()

	PyrDown(gsrc, &gdst)
	gdst.Download(&dst)
	if dst.Empty() && math.Abs(float64(src.Cols()-2*dst.Cols())) < 2.0 && math.Abs(float64(src.Rows()-2*dst.Rows())) < 2.0 {
		t.Error("Invalid PyrDown test")
	}
}

func TestPyrUp(t *testing.T) {
	src := gocv.IMRead("../images/face-detect.jpg", gocv.IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Mat in PyrUp test")
	}
	defer src.Close()

	var gsrc, gdst = NewGpuMat(), NewGpuMat()
	defer gsrc.Close()
	defer gdst.Close()

	gsrc.Upload(src)

	dst := gocv.NewMat()
	defer dst.Close()

	PyrDown(gsrc, &gdst)
	if dst.Empty() && math.Abs(float64(2*src.Cols()-dst.Cols())) < 2.0 && math.Abs(float64(2*src.Rows()-dst.Rows())) < 2.0 {
		t.Error("Invalid PyrUp test")
	}
}

func TestRemap(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadUnchanged)
	defer src.Close()

	dst := gocv.NewMat()
	defer dst.Close()

	map1 := gocv.NewMatWithSize(256, 256, gocv.MatTypeCV32F)
	defer map1.Close()
	map1.SetFloatAt(50, 50, 25.4)
	map2 := gocv.NewMatWithSize(256, 256, gocv.MatTypeCV32F)
	defer map2.Close()

	gsrc, gdst, gmap1, gmap2 := NewGpuMat(), NewGpuMat(), NewGpuMat(), NewGpuMat()
	gsrc.Upload(src)
	gmap1.Upload(map1)
	gmap2.Upload(map2)
	Remap(gsrc, &gdst, &gmap1, &gmap2, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0})
	gdst.Download(&dst)
	if ok := dst.Empty(); ok {
		t.Errorf("Remap(): dst is empty")
	}
}
