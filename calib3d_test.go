package gocv

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestFisheyeUndistorImage(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	FisheyeUndistortImage(img, &dest, k, d)

	if dest.Empty() {
		t.Error("final image is empty")
		return
	}
	// IMWrite("images/fisheye_sample-u.jpg", dest)
}

func TestFisheyeUndistorImageWithParams(t *testing.T) {
	img := IMRead("images/distortion.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 689.21)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 1295.56)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 690.48)
	k.SetDoubleAt(1, 2, 942.17)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 4, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, 0)
	d.SetDoubleAt(0, 1, 0)
	d.SetDoubleAt(0, 2, 0)
	d.SetDoubleAt(0, 3, 0)

	knew := NewMat()
	defer knew.Close()

	k.CopyTo(&knew)

	knew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	size := image.Point{dest.Rows(), dest.Cols()}
	FisheyeUndistortImageWithParams(img, &dest, k, d, knew, size)

	if dest.Empty() {
		t.Error("final image is empty")
		return
	}
	// IMWrite("images/fisheye_sample-up.jpg", dest)
}

func TestInitUndistortRectifyMap(t *testing.T) {
	img := IMRead("images/distortion.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
		return
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	k := NewMatWithSize(3, 3, MatTypeCV64F)
	defer k.Close()

	k.SetDoubleAt(0, 0, 842.0261028)
	k.SetDoubleAt(0, 1, 0)
	k.SetDoubleAt(0, 2, 667.7569792)

	k.SetDoubleAt(1, 0, 0)
	k.SetDoubleAt(1, 1, 707.3668897)
	k.SetDoubleAt(1, 2, 385.56476464)

	k.SetDoubleAt(2, 0, 0)
	k.SetDoubleAt(2, 1, 0)
	k.SetDoubleAt(2, 2, 1)

	d := NewMatWithSize(1, 5, MatTypeCV64F)
	defer d.Close()

	d.SetDoubleAt(0, 0, -3.65584802e-01)
	d.SetDoubleAt(0, 1, 1.41555815e-01)
	d.SetDoubleAt(0, 2, -2.62985819e-03)
	d.SetDoubleAt(0, 3, 2.05841873e-04)
	d.SetDoubleAt(0, 4, -2.35021914e-02)
	//FisheyeUndistortImage(img, &dest, k, d)
	//img.Reshape()
	newC, roi := GetOptimalNewCameraMatrixWithParams(k, d, image.Point{X: img.Cols(), Y: img.Rows()}, (float64)(1), image.Point{X: img.Cols(), Y: img.Rows()}, false)
	if newC.Empty() {
		t.Error("final image is empty")
		return
	}
	fmt.Printf("roi:%+v\n", roi)
	defer newC.Close()
	r := NewMat()
	defer r.Close()
	mapx := NewMat()
	defer mapx.Close()
	mapy := NewMat()
	defer mapy.Close()
	//dest := NewMat()
	InitUndistortRectifyMap(k, d, r, newC, image.Point{X: img.Cols(), Y: img.Rows()}, 5, mapx, mapy)

	Remap(img, &dest, &mapx, &mapy, InterpolationDefault, BorderConstant, color.RGBA{0, 0, 0, 0})
	flg := IMWrite("images/distortion-correct.jpg", dest)
	if !flg {
		t.Error("IMWrite failed")
	}
}
