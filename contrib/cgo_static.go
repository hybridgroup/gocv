//go:build !customenv && opencvstatic && linux

package contrib

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS: --std=c++17 -DNDEBUG
#cgo pkg-config: --static opencv5
*/
import "C"
