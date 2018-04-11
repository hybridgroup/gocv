package gocv

import (
	"image"
	"testing"
)

func TestFisheyeCalibrate(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
	}
	defer img.Close()

	obj := NewMat()
	defer obj.Close()

	obj.SetDoubleAt(0, 0, 1)
	obj.SetDoubleAt(0, 1, 0)
	obj.SetDoubleAt(0, 0, 1)
	obj.SetDoubleAt(0, 0, 1)

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

	k.CopyTo(knew)

	knew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	rvec := NewMatWithSize(10, 10, MatTypeCV64F)
	defer rvec.Close()

	tvec := NewMatWithSize(10, 10, MatTypeCV64F)
	defer tvec.Close()

	FisheyeCalibrate(obj, img, k, d, &rvec, &tvec, image.Point{X: 0, Y: 0})

}

func TestFisheyeUndistorImage(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
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
	}
}

func TestFisheyeUndistorImageWithKNewMat(t *testing.T) {
	img := IMRead("images/fisheye_sample.jpg", IMReadUnchanged)
	if img.Empty() {
		t.Error("Invalid read of Mat test")
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

	k.CopyTo(knew)

	knew.SetDoubleAt(0, 0, 0.4*k.GetDoubleAt(0, 0))
	knew.SetDoubleAt(1, 1, 0.4*k.GetDoubleAt(1, 1))

	FisheyeUndistortImageWithKNewMat(img, &dest, k, d, knew)

	if dest.Empty() {
		t.Error("final image is empty")
	}
}
