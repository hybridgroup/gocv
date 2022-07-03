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

func TestFastNlMeansDenoisingColoredMultiWithParams(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	dst := NewMat()
	defer dst.Close()

	FastNlMeansDenoisingColoredMultiWithParams([]Mat{src[0], src[1], src[2]}, &dst, 1, 1, 3, 3, 7, 21)

	if dst.Empty() || dst.Rows() != src[0].Rows() || dst.Cols() != src[0].Cols() {
		t.Error("Invalid FastNlMeansDenoisingColoredMultiWithParams test")
	}
}

func TestMergeMertens(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	dst := NewMat()
	defer dst.Close()

	mertens := NewMergeMertens()
	defer mertens.Close()

	mertens.Process([]Mat{src[0], src[1], src[2]}, &dst)

	if dst.Empty() || dst.Rows() != src[0].Rows() || dst.Cols() != src[0].Cols() {
		t.Error("Invalid TestMergeMertens test")
	}
}

func TestNewAlignMTB(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	alignwtb := NewAlignMTB()
	defer alignwtb.Close()

	var dst []Mat
	alignwtb.Process([]Mat{src[0], src[1], src[2]}, &dst)

	sizedst := len(dst)
	t.Logf(" Size Dst slice : %d ", sizedst)
	if sizedst > 0 {
		if dst[0].Empty() || dst[0].Rows() != src[0].Rows() || dst[0].Cols() != src[0].Cols() {
			t.Error("Invalid TestNewAlignMTB test")
		}
	}
	if sizedst <= 0 {
		t.Error("Invalid TestNewAlignMTB test : empty result")
	}
}

func TestFastNlMeansDenoising(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in TestFastNlMeansDenoising test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoising(img, &dest)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in TestFastNlMeansDenoising test")
	}
}

func TestFastNlMeansDenoisingWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadGrayScale)
	if img.Empty() {
		t.Error("Invalid read of Mat in TestFastNlMeansDenoising test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingWithParams(img, &dest, 3, 7, 21)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in TestFastNlMeansDenoising test")
	}
}

func TestFastNlMeansDenoisingColored(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in FastNlMeansDenoisingColored test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingColored(img, &dest)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in FastNlMeansDenoisingColored test")
	}
}

func TestFastNlMeansDenoisingColoredWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in FastNlMeansDenoisingColored test")
	}
	defer img.Close()

	dest := NewMat()
	defer dest.Close()

	FastNlMeansDenoisingColoredWithParams(img, &dest, 3, 3, 7, 21)
	if dest.Empty() || img.Rows() != dest.Rows() || img.Cols() != dest.Cols() {
		t.Error("Error in FastNlMeansDenoisingColored test")
	}
}

func TestDetailEnhance(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	DetailEnhance(src, &dst, 100.0, .5)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid DetailEnhance test")
	}
}

func TestEdgePreservingFilter(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	EdgePreservingFilter(src, &dst, RecursFilter, 100.0, .5)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invalid EdgePreservingFilter test")
	}
}

func TestStylization(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()

	Stylization(src, &dst, 100.0, .5)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid Stylization test")
	}
}

func TestPencilSketch(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst1 := NewMat()
	defer dst1.Close()
	dst2 := NewMat()
	defer dst2.Close()

	PencilSketch(src, &dst1, &dst2, 100.0, .5, 0.05)
	if dst1.Empty() || dst1.Rows() != src.Rows() || dst1.Cols() != src.Cols() {
		t.Error("Invlalid PencilSketch test")
	}
	if dst2.Empty() || dst2.Rows() != src.Rows() || dst2.Cols() != src.Cols() {
		t.Error("Invlalid PencilSketch test")
	}
} 

func TestInpaint(t *testing.T) {
	src := IMRead("images/inpaint-src.jpg", IMReadColor)
	defer src.Close()
	mask := IMRead("images/inpaint-mask.jpg", IMReadGrayScale)
	defer mask.Close()
	dst := NewMatWithSize(src.Rows(), src.Cols(), MatTypeCV8U)
	defer dst.Close()

	Inpaint(src, mask, &dst, 10, Telea)
	if dst.Channels() == 1 {
		t.Error("Invalid inpaint test")
	}

	if sum := dst.Sum(); sum.Val1 == 0 || sum.Val2 == 0 || sum.Val3 == 0 {
		t.Error("Invalid inpaint test")
	}
}
