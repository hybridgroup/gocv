// +build !customenv,!openvino

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
*/
import "C"
