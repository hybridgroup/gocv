package gocv

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestIMRead(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
}

func TestIMWrite(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gocvtests")
	tmpfn := filepath.Join(dir, "test.jpg")

	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in IMWrite test")
	}

	result := IMWrite(tmpfn, img)
	if !result {
		t.Error("Invalid write of Mat in IMWrite test")
	}
}

func TestIMWriteWithParams(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gocvtests")
	tmpfn := filepath.Join(dir, "test.jpg")

	img := IMRead("images/face-detect.jpg", IMReadColor)
	if img.Empty() {
		t.Error("Invalid read of Mat in IMWrite test")
	}

	result := IMWriteWithParams(tmpfn, img, []int{ImwriteJpegQuality, 60})
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
