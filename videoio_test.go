package gocv

import (
	"testing"
)

func TestVideoWriterFile(t *testing.T) {
	vw, _ := VideoWriterFile("/tmp/test.avi", 25, 800, 600)
	if !vw.IsOpened() {
		t.Error("Unable to open VideoWriterFile")
	}
}
