package cuda

import (
	"testing"
)

func TestNewGpuMat(t *testing.T) {
	mat := NewGpuMat()
	defer mat.Close()

	if !mat.Empty() {
		t.Error("New Mat should be empty")
	}
}

func TestGetCudaEnabledDeviceCount(t *testing.T) {
	if GetCudaEnabledDeviceCount() < 1 {
		t.Fatal("expected atleast one cuda enabled device")
	}
}
