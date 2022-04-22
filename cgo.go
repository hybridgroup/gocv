//go:build !customenv && !static
// +build !customenv,!static

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows CXXFLAGS: -I/usr/local/include  -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_imgcodecs -lopencv_imgproc -lopencv_core -lz -ljpeg -lpng -lgif -ldl -lm -lpthread -lrt -lquadmath
*/
import "C"
