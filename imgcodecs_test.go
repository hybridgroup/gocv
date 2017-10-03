package gocv

import (
	"testing"
)

func TestIMRead(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
}

func TestIMWrite(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in IMWrite test")
	}

	result := IMWrite("/tmp/test.jpg", img)
	if !result {
		t.Error("Invalid write of Mat in IMWrite test")
	}
}

func TestIMEncode(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in IMEncode test")
	}

	buf, err := IMEncode(".jpg", img)
	if err != nil {
		t.Error(err)
	}
	if len(buf) < 43000 {
		t.Errorf("Wrong buffer size in IMEncode test. Should have been %v\n", len(buf))
	}
}
