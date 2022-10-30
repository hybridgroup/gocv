package cuda

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestNewGpuMat(t *testing.T) {
	mat := NewGpuMat()
	defer mat.Close()

	if !mat.Empty() {
		t.Error("New GpuMat should be empty")
	}
}

func TestNewGpuMatFromMat(t *testing.T) {
	mat := gocv.NewMat()
	defer mat.Close()

	gpumat := NewGpuMatFromMat(mat)
	defer gpumat.Close()

	if !gpumat.Empty() {
		t.Error("New GpuMat should be empty")
	}
}

func TestNewGpuMatFromMatWithSize(t *testing.T) {
	mat := gocv.NewMatWithSize(100, 200, gocv.MatTypeCV32FC4)
	defer mat.Close()

	gpumat := NewGpuMatFromMat(mat)
	defer gpumat.Close()

	if gpumat.Empty() {
		t.Error("New GpuMat should be not empty")
	}

	if gpumat.Rows() != 100 {
		t.Error("incorrect number of rows for GpuMat")
	}

	if gpumat.Cols() != 200 {
		t.Error("incorrect number of cols for GpuMat")
	}

	if gpumat.Type() != gocv.MatTypeCV32FC4 {
		t.Error("incorrect type for GpuMat")
	}
}

func TestNewGpuMatWithSize(t *testing.T) {
	gpumat := NewGpuMatWithSize(100, 200, gocv.MatTypeCV32FC4)
	defer gpumat.Close()

	if gpumat.Empty() {
		t.Error("New GpuMat should be not empty")
	}

	if gpumat.Rows() != 100 {
		t.Error("incorrect number of rows for GpuMat")
	}

	if gpumat.Cols() != 200 {
		t.Error("incorrect number of cols for GpuMat")
	}

	if gpumat.Type() != gocv.MatTypeCV32FC4 {
		t.Error("incorrect type for GpuMat")
	}
}

func TestConvertToWithParams(t *testing.T) {
	src1 := gocv.IMRead("../images/gocvlogo.jpg", gocv.IMReadColor)
	if src1.Empty() {
		t.Error("Invalid read of Mat in ConvertToWithParams test")
	}
	defer src1.Close()

	var cimg1, dimg = NewGpuMat(), NewGpuMat()
	defer cimg1.Close()
	defer dimg.Close()

	cimg1.Upload(src1)

	dest := gocv.NewMat()
	defer dest.Close()

	cimg.ConvertToWithParams(&dimg, cimg.Type(), 1.0, 10.0)
	dimg.Download(&dest)
	if dest.Empty() || src1.Rows() != dest.Rows() || src1.Cols() != dest.Cols() {
		t.Error("Invalid ConvertToWithParams test")
	}
}

func TestGetCudaEnabledDeviceCount(t *testing.T) {
	if GetCudaEnabledDeviceCount() < 1 {
		t.Fatal("expected atleast one cuda enabled device")
	}
}
