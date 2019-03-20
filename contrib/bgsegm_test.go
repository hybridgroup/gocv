package contrib

import (
	"testing"
	"gocv.io/x/gocv"
	v "gocv.io/x/gocv/contrib"
)

func TestCNT(t *testing.T) {
	img := gocv.IMRead("../images/face.jpg", gocv.IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in CNT test")
	}
	defer img.Close()

	dst := gocv.NewMat()
	defer dst.Close()

	cnt := v.NewBackgroundSubtractorCNT()
	defer cnt.Close()

	cnt.Apply(img, &dst)

	if dst.Empty() {
		t.Error("Error in TestCNT test")
	}
}

