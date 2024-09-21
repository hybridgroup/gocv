//go:build !customenv && static && darwin

package gocv

// Changes here should be mirrored in contrib/cgo_static_darwin.go and cuda/cgo_static_darwin.go.

/*
#cgo CXXFLAGS: --std=c++11
#cgo pkg-config: --static opencv4
#cgo LDFLAGS: -static
*/
import "C"
