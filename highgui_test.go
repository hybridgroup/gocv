package opencv3

import (
	"testing"
)

func TestWindow(t *testing.T) {
	window := NewWindow("test")
	if window == nil {
		t.Error("Unable to create Window")
	}
	if window.name != "test" {
		t.Error("Invalid Window name")
	}
	val := WaitKey(1)
	if val != -1 {
		t.Error("Invalid WaitKey")
	}
}
