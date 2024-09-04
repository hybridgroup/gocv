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

func TestAbsDiff(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in AbsDiff test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	AbsDiff(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid AbsDiff test")
	}
}

func TestAdd(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in AbsDiff test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Add(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Add test")
	}
}

func TestBitwiseAnd(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in AbsDiff test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	BitwiseAnd(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid BitwiseAnd test")
	}
}

func TestBitwiseNot(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in AbsDiff test")
	}
	defer src1.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	BitwiseNot(cimg, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid BitwiseNot test")
	}
}

func TestBitwiseOr(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in BitwiseOr test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	BitwiseOr(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid BitwiseOr test")
	}
}

func TestBitwiseXor(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in BitwiseXor test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	BitwiseXor(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid BitwiseXor test")
	}
}

func TestDivide(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Divide test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Divide(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Divide test")
	}
}

func TestDivideWithStream(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Divide test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	var s = NewStream()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()
	defer s.Close()

	cimg1.UploadWithStream(src1, s)
	cimg2.UploadWithStream(src1, s)

	dest := gocv.NewMat()
	defer dest.Close()

	DivideWithStream(cimg1, cimg2, &dimg, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Divide test")
	}
}

func TestExp(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Exp test")
	}
	defer src1.Close()

	var cimg1, dimg = NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer dimg.Close()

	cimg1.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Exp(cimg1, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Exp test")
	}
}

func TestLog(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Log test")
	}
	defer src1.Close()

	var cimg1, dimg = NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer dimg.Close()

	cimg1.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Log(cimg1, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Log test")
	}
}

func TestMax(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Max test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Max(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Max test")
	}
}

func TestMin(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Min test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Min(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Min test")
	}
}

func TestMultiply(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Multiply test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()

	cimg1.Upload(src1)
	cimg2.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	Multiply(cimg1, cimg2, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Multiply test")
	}
}

func TestMultiplyWithStream(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in Multiply test")
	}
	defer src1.Close()

	var cimg1, cimg2, dimg = NewGpuMat(), NewGpuMat(), NewGpuMat()
	var s = NewStream()
	defer cimg1.Close()
	defer cimg2.Close()
	defer dimg.Close()
	defer s.Close()

	cimg1.UploadWithStream(src1, s)
	cimg2.UploadWithStream(src1, s)

	dest := gocv.NewMat()
	defer dest.Close()

	MultiplyWithStream(cimg1, cimg2, &dimg, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid Multiply test")
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

func TestMerge(t *testing.T) {
	src := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src.Close()
	src2 := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src2.Close()
	src3 := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src3.Close()

	dstGPU := NewGpuMat()
	defer dstGPU.Close()

	Merge([]GpuMat{src, src2, src3}, &dstGPU)
	if dstGPU.Empty() {
		t.Error("TestMerge dst should not be empty.")
	}
}

func TestMergeWithStream(t *testing.T) {
	src := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src.Close()
	src2 := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src2.Close()
	src3 := NewGpuMatWithSize(101, 102, gocv.MatTypeCV8U)
	defer src3.Close()
	s := NewStream()
	defer s.Close()

	dstGPU := NewGpuMat()
	defer dstGPU.Close()

	MergeWithStream([]GpuMat{src, src2, src3}, &dstGPU, s)

	s.WaitForCompletion()
	if dstGPU.Empty() {
		t.Error("TestMergeWithStream dst should not be empty.")
	}
}

func TestTranspose(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in Transpose test")
	}
	defer src.Close()

	var cimg, dimg = NewGpuMat(), NewGpuMat()
	defer cimg.Close()
	defer dimg.Close()

	cimg.Upload(src)

	dest := gocv.NewMat()
	defer dest.Close()

	Transpose(cimg, &dimg)
	dimg.Download(&dest)
	if dest.Empty() || src.Rows() != dest.Cols() || src.Cols() != dest.Rows() {
		t.Error("Invalid Transpose test")
	}
}

func TestTransposeWithStream(t *testing.T) {
	src := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadGrayScale)
	if src.Empty() {
		t.Error("Invalid read of Mat in TransposeWithStream test")
	}
	defer src.Close()

	var cimg, dimg, s = NewGpuMat(), NewGpuMat(), NewStream()
	defer cimg.Close()
	defer dimg.Close()
	defer s.Close()

	cimg.Upload(src)

	dest := gocv.NewMat()
	defer dest.Close()

	TransposeWithStream(cimg, &dimg, s)
	dimg.DownloadWithStream(&dest, s)

	s.WaitForCompletion()

	if dest.Empty() || src.Rows() != dest.Cols() || src.Cols() != dest.Rows() {
		t.Error("Invalid TransposeWithStream test")
	}
}
