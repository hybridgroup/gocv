package gocv

import (
	"image"
	"testing"
)

func TestSeamlessClone(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMatWithSize(30, 30, MatTypeCV8UC3)
	defer dst.Close()
	blend := NewMatWithSize(dst.Rows(), dst.Cols(), dst.Type())
	defer blend.Close()
	mask := src.Clone()
	defer mask.Close()

	center := image.Point{dst.Cols() / 2, dst.Rows() / 2}
	SeamlessClone(src, dst, mask, center, &blend, NormalClone)
	if blend.Empty() || dst.Rows() != blend.Rows() || dst.Cols() != blend.Cols() {
		t.Error("Invlalid SeamlessClone test")
	}
}
