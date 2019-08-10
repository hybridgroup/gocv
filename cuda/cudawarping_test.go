package cuda

import (
	"image"
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
