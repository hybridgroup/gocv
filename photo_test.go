package gocv

import (
	"image"
	"testing"
)

func TestColorChange(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	ColorChange(src, mask, &dst, 1.5, .5, .5)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid ColorChange test")
	}
}

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

func TestIlluminationChange(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	IlluminationChange(src, mask, &dst, 0.2, 0.4)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid IlluminationChange test")
	}
}

func TestTextureFlattening(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	TextureFlattening(src, mask, &dst, 30, 45, 3)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid TextureFlattening test")
	}
}
