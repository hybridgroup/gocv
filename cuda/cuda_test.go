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

func TestGetCudaEnabledDeviceCount(t *testing.T) {
	if GetCudaEnabledDeviceCount() < 1 {
		t.Fatal("expected atleast one cuda enabled device")
	}
}
