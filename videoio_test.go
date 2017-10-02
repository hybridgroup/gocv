package opencv3

import (
	"testing"
)

func TestVideoWriterFile(t *testing.T) {
	vw, _ := VideoWriterFile("/tmp/test.mp4", 25, 800, 600)
	if !vw.IsOpened() {
		t.Error("Unable to open VideoWriterFile")
	}
}
