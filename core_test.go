package gocv

import (
	"image"
	"testing"
)

func TestMat(t *testing.T) {
	mat := NewMat()
	if !mat.Empty() {
		t.Error("New Mat should be empty")
	}
}

func TestMatWithSize(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	if mat.Empty() {
		t.Error("NewMatWithSize should not be empty")
	}

	if mat.Rows() != 101 {
		t.Errorf("NewMatWithSize incorrect row count: %v\n", mat.Rows())
	}

	if mat.Cols() != 102 {
		t.Errorf("NewMatWithSize incorrect col count: %v\n", mat.Cols())
	}
}

func TestMatClone(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	clone := mat.Clone()
	if clone.Rows() != 101 {
		t.Errorf("Mat clone incorrect row count: %v\n", clone.Rows())
	}

	if clone.Cols() != 102 {
		t.Errorf("Mat clone incorrect col count: %v\n", clone.Cols())
	}
}

func TestMatCopyTo(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	copy := NewMat()
	mat.CopyTo(copy)
	if copy.Rows() != 101 {
		t.Errorf("Mat copy incorrect row count: %v\n", copy.Rows())
	}

	if copy.Cols() != 102 {
		t.Errorf("Mat copy incorrect col count: %v\n", copy.Cols())
	}
}

func TestMatRegion(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8U)
	region := mat.Region(image.Rect(20, 25, 80, 75))
	if region.Rows() != 50 {
		t.Errorf("Mat region incorrect row count: %v\n", region.Rows())
	}

	if region.Cols() != 60 {
		t.Errorf("Mat region incorrect col count: %v\n", region.Cols())
	}
}

func TestMatMean(t *testing.T) {
	mat := NewMatWithSize(100, 100, MatTypeCV8U)
	mean := mat.Mean()
	if mean.Val1 != 0 {
		t.Errorf("Mat Mean incorrect Val1")
	}
}

func TestLUT(t *testing.T) {
	src := IMRead("images/gocvlogo.jpg", IMReadColor)
	if src.Empty() {
		t.Error("Invalid read of Source Mat in LUT test")
	}
	defer src.Close()

	lut := IMRead("images/lut.png", IMReadColor)
	if lut.Empty() {
		t.Error("Invalid read of LUT Mat in LUT test")
	}
	defer lut.Close()

	dst := NewMat()
	defer dst.Close()

	LUT(src, lut, dst)
	if dst.Cols() != 400 || dst.Rows() != 343 {
		t.Errorf("Expected dst size of 200x172 got %dx%d", dst.Cols(), dst.Rows())
	}
}

func TestMatAccessors(t *testing.T) {
	mat := NewMatWithSize(101, 102, MatTypeCV8U)
	if mat.GetUCharAt(50, 50) != 0 {
		t.Errorf("GetUCharAt incorrect value: %v\n", mat.GetUCharAt(50, 50))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV8S)
	if mat.GetSCharAt(50, 50) != 0 {
		t.Errorf("GetSCharAt incorrect value: %v\n", mat.GetSCharAt(50, 50))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV16S)
	if mat.GetShortAt(50, 50) != 0 {
		t.Errorf("GetShortAt incorrect value: %v\n", mat.GetShortAt(50, 50))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV32S)
	if mat.GetIntAt(50, 50) != 0 {
		t.Errorf("GetIntAt incorrect value: %v\n", mat.GetIntAt(50, 50))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV32F)
	if mat.GetFloatAt(50, 50) != 0.0 {
		t.Errorf("GetFloatAt incorrect value: %v\n", mat.GetFloatAt(50, 50))
	}
	mat.Close()

	mat = NewMatWithSize(101, 102, MatTypeCV64F)
	if mat.GetDoubleAt(50, 50) != 0.0 {
		t.Errorf("GetDoubleAt incorrect value: %v\n", mat.GetDoubleAt(50, 50))
	}
	mat.Close()
}

func TestMatAbsDiff(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat3 := NewMat()
	AbsDiff(mat1, mat2, mat3)
	if mat3.Empty() {
		t.Error("TestMatAbsDiff dest mat3 should not be empty.")
	}
}

func TestMatAdd(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat3 := NewMat()
	Add(mat1, mat2, mat3)
	if mat3.Empty() {
		t.Error("TestMatAdd dest mat3 should not be empty.")
	}
}

func TestMatAddWeighted(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat3 := NewMat()
	AddWeighted(mat1, 2.0, mat2, 3.0, 4.0, mat3)
	if mat3.Empty() {
		t.Error("TestMatAddWeighted dest mat3 should not be empty.")
	}
}

func TestMatBitwiseOperations(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat2 := NewMatWithSize(101, 102, MatTypeCV8U)
	mat3 := NewMat()
	BitwiseAnd(mat1, mat2, mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseAnd dest mat3 should not be empty.")
	}

	BitwiseOr(mat1, mat2, mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseOr dest mat3 should not be empty.")
	}

	BitwiseXor(mat1, mat2, mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseXor dest mat3 should not be empty.")
	}

	BitwiseNot(mat1, mat3)
	if mat3.Empty() {
		t.Error("TestMatBitwiseNot dest mat3 should not be empty.")
	}

}

func TestMatInRange(t *testing.T) {
	mat1 := NewMatWithSize(101, 102, MatTypeCV8U)
	lb := NewMatFromScalar(NewScalar(20.0, 100.0, 100.0, 0.0), MatTypeCV8U)
	ub := NewMatFromScalar(NewScalar(20.0, 100.0, 100.0, 0.0), MatTypeCV8U)
	dst := NewMat()
	InRange(mat1, lb, ub, dst)
	if dst.Empty() {
		t.Error("TestMatAddWeighted dest mat3 should not be empty.")
	}
}

func TestMatDFT(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV32F)
	dst := NewMat()

	m := GetOptimalDFTSize(101)
	n := GetOptimalDFTSize(102)
	if m != 108 {
		t.Errorf("TestMatOptimalDFT dst error: %d", m)
	}

	if n != 108 {
		t.Errorf("TestMatOptimalDFT dst error: %d", n)
	}

	DFT(src, dst)
	if dst.Empty() {
		t.Error("TestMatDFT dst should not be empty.")
	}
}

func TestMatMerge(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	dst := NewMat()
	Merge(src, 1, dst)
	if dst.Empty() {
		t.Error("TestMatMerge dst should not be empty.")
	}
}
func TestMatNormalize(t *testing.T) {
	src := NewMatWithSize(101, 102, MatTypeCV8U)
	dst := NewMat()
	Normalize(src, dst, 0.0, 255.0, NormMixMax)
	if dst.Empty() {
		t.Error("TestMatNormalize dst should not be empty.")
	}
}

func TestTermCriteria(t *testing.T) {
	tc := NewTermCriteria(Count, 50, 2.0)
	if tc.p == nil {
		t.Error("TermCriteria has invalid value")
	}
}

func TestScalar(t *testing.T) {
	s := NewScalar(127.0, 255.0, 64.0, 0.0)
	if s.Val1 != 127.0 || s.Val2 != 255.0 || s.Val3 != 64.0 || s.Val4 != 0.0 {
		t.Error("Scalar has invalid value")
	}
}
