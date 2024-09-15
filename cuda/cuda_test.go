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

func TestGpuMatClosed(t *testing.T) {
	mat := NewGpuMat()
	mat.Close()

	if !mat.Closed() {
		t.Error("Closed GpuMat should be closed")
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

func TestConvertFp16(t *testing.T) {
	gpumat := NewGpuMatWithSize(100, 200, gocv.MatTypeCV32FC4)
	defer gpumat.Close()

	fp16mat := NewGpuMatWithSize(100, 200, gocv.MatTypeCV32FC4)
	defer fp16mat.Close()

	gpumat.ConvertFp16(&fp16mat)

	if fp16mat.Empty() {
		t.Error("New fp16mat should be not empty")
	}

	if fp16mat.Rows() != 100 {
		t.Error("incorrect number of rows for fp16mat")
	}

	if fp16mat.Cols() != 200 {
		t.Error("incorrect number of cols for fp16mat")
	}

	if fp16mat.Type() != gocv.MatTypeCV16SC4 {
		t.Error("incorrect type for fp16mat", fp16mat.Type())
	}
}

func TestCudaDeviceSupports(t *testing.T) {
	if !DeviceSupports(FeatureSetCompute10) {
		t.Fatal("expected FeatureSetCompute10 on cuda enabled device")
	}
}
