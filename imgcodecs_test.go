package gocv

import (
	"testing"
)

func TestIMRead(t *testing.T) {
	img := IMRead("images/face-detect.png", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
}

func TestIMWrite(t *testing.T) {
	t.Skip("Tests needed")
}
