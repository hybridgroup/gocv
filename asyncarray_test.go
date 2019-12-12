// +build openvino

package gocv

import (
	"testing"
)

func TestAsyncArray(t *testing.T) {
	asyncarray := NewAsyncArray()
	defer asyncarray.Close()

	if asyncarray.Ptr() == nil {
		t.Error("New AsyncArray should not be nil")
	}
}
