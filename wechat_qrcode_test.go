package gocv

import (
	"os"
	"reflect"
	"testing"
)

func TestNewWeChatQRCode(t *testing.T) {
	tests := []struct {
		name    string
		notWant *WeChatQRCode
	}{
		{"testNewWeChatQRCode", nil},
	}

	path := os.Getenv("GOCV_CAFFE_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Caffe model files for tests")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewWeChatQRCode(path+"/detect.prototxt", path+"/detect.caffemodel",
				path+"/sr.prototxt", path+"/sr.caffemodel"); reflect.DeepEqual(got, tt.notWant) {
				t.Errorf("NewWeChatQRCode() = %v, want %v", got, tt.notWant)
			}
		})
	}
}

func TestWeChatQRCode_DetectAndDecode(t *testing.T) {
	mat := IMRead("images/qrcode.png", IMReadColor)
	mats := make([]Mat, 0)
	defer mat.Close()

	type args struct {
		img   Mat
		point *[]Mat
	}
	tests := []struct {
		name     string
		args     args
		want     []string
		qrCounts int
	}{
		{"TestDetectAndDecode", args{point: &mats, img: mat}, []string{"Hello World!"}, 1},
	}

	path := os.Getenv("GOCV_CAFFE_TEST_FILES")
	if path == "" {
		t.Skip("Unable to locate Caffe model files for tests")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wq := NewWeChatQRCode(path+"/detect.prototxt", path+"/detect.caffemodel",
				path+"/sr.prototxt", path+"/sr.caffemodel")
			if got := wq.DetectAndDecode(tt.args.img, tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectAndDecode() = %v, want %v", got, tt.want)
			}
			if len(mats) != tt.qrCounts {
				t.Errorf("DetectAndDecode() = %v, want qrcode counts %v", tt.qrCounts, len(mats))
			}
		})
	}
}
