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
		t.Errorf("Mat clone incorrect row count: %v\n", mat.Rows())
	}

	if clone.Cols() != 102 {
		t.Errorf("Mat clone incorrect col count: %v\n", mat.Cols())
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

func TestScalar(t *testing.T) {
	s := NewScalar(127.0, 255.0, 64.0, 0.0)
	if s.Val1 != 127.0 || s.Val2 != 255.0 || s.Val3 != 64.0 || s.Val4 != 0.0 {
		t.Error("Scalar has invalid value")
	}
}
