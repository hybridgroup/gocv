package cuda

import (
	"gocv.io/x/gocv"
	"testing"
)

func TestCudaMOG2(t *testing.T) {
	img := gocv.IMRead("../images/face.jpg", gocv.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG2 test")
	}
	defer img.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(img)

	dst := gocv.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	mog2.Apply(cimg, &dimg)

	dimg.Download(&dst)

	if dst.Empty() {
		t.Error("Error in TestCudaMOG2 test")
	}
}

func TestCudaMOG(t *testing.T) {
	img := gocv.IMRead("../images/face.jpg", gocv.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CudaMOG test")
	}
	defer img.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(img)

	dst := gocv.NewMat()
	defer dst.Close()

	mog2 := NewBackgroundSubtractorMOG()
	defer mog2.Close()

	mog2.Apply(cimg, &dimg)

	dimg.Download(&dst)

	if dst.Empty() {
		t.Error("Error in TestCudaMOG test")
	}
}
