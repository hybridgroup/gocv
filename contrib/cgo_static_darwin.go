//go:build !customenv && opencvstatic && darwin

package contrib

// Changes here should be mirrored in contrib/cgo_static_darwin.go and cuda/cgo_static_darwin.go.

/*
#cgo CXXFLAGS: --std=c++11 -DNDEBUG
#cgo pkg-config: --static opencv4
*/
import "C"
