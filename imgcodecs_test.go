package gocv

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
)

func TestIMRead(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	defer img.Close()
	if img.Empty() {
		t.Error("Invalid Mat in IMRead")
	}
}

func TestIMWrite(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gocvtests")
	tmpfn := filepath.Join(dir, "test.jpg")

	img := IMRead("images/face-detect.jpg", IMReadColor)
	defer img.Close()
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
	defer img.Close()
	if img.Empty() {
		t.Error("Invalid read of Mat in IMWrite test")
	}

	result := IMWriteWithParams(tmpfn, img, []int{IMWriteJpegQuality, 60})
	if !result {
		t.Error("Invalid write of Mat in IMWrite test")
	}
}

func TestIMEncode(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	defer img.Close()
	if img.Empty() {
		t.Error("Invalid Mat in IMEncode test")
	}

	buf, err := IMEncode(PNGFileExt, img)
	if err != nil {
		t.Error(err)
	}
	if len(buf) < 43000 {
		t.Errorf("Wrong buffer size in IMEncode test. Should have been %v\n", len(buf))
	}
}

func ExampleIMEncodeWithParams() {
	img := IMRead(path.Join(os.Getenv("GOPATH"), "src/gocv.io/x/gocv/images/face-detect.jpg"), IMReadColor)
	if img.Empty() {
		log.Fatal("Invalid Mat")
	}

	imgHandler := func(w http.ResponseWriter, req *http.Request) {
		quality := 75
		if q, err := strconv.Atoi(req.URL.Query().Get("q")); err == nil {
			quality = q
		}
		buffer, err := IMEncodeWithParams(JPEGFileExt, img, []int{IMWriteJpegQuality, quality})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(buffer)
	}

	http.HandleFunc("/img", imgHandler)
	fmt.Println("Open in browser http://127.0.0.1:8080/img?q=10 where q is a JPEG quality parameter")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func TestIMEncodeWithParams(t *testing.T) {
	img := IMRead("images/face-detect.jpg", IMReadColor)
	defer img.Close()
	if img.Empty() {
		t.Error("Invalid Mat in IMEncode test")
	}

	buf, err := IMEncodeWithParams(JPEGFileExt, img, []int{IMWriteJpegQuality, 75})
	if err != nil {
		t.Error(err)
	}
	if len(buf) < 18000 {
		t.Errorf("Wrong buffer size in IMEncode test. Should have been %v\n", len(buf))
	}

	buf2, err := IMEncodeWithParams(JPEGFileExt, img, []int{IMWriteJpegQuality, 100})
	if err != nil {
		t.Error(err)
	}
	if len(buf2) < 18000 {
		t.Errorf("Wrong buffer size in IMEncode test. Should have been %v\n", len(buf))
	}

	if len(buf) >= len(buf2) {
		t.Errorf("Jpeg quality parameter does not work correctly\n")
	}
}

func TestIMDecode(t *testing.T) {
	content, err := ioutil.ReadFile("images/face-detect.jpg")
	if err != nil {
		t.Error("Invalid ReadFile in IMDecode")
	}

	dec, err := IMDecode(content, IMReadColor)
	if err != nil {
		t.Error(err.Error())
	}
	if dec.Empty() {
		t.Error("Invalid Mat in IMDecode")
	}
	dec.Close()

	dec, err = IMDecode([]byte{}, IMReadColor)
	if err == nil {
		t.Error("Should not decode empty array")
	}
}
func TestIMDecodeWebp(t *testing.T) {
	content, err := ioutil.ReadFile("images/sample.webp")
	if err != nil {
		t.Error("Invalid ReadFile in IMDecodeWebp")
	}

	dec, err := IMDecode(content, IMReadColor)
	if err != nil {
		t.Error(err.Error())
	}
	if dec.Empty() {
		t.Error("Invalid Mat in IMDecodeWebp")
	}
	dec.Close()

}
