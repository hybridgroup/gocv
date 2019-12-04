// +build openvino

package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"

type AsyncArray struct {
	p C.AsyncArray
}

// NewAsyncArray returns a new empty AsyncArray.
func NewAsyncArray() AsyncArray {
	return newAsyncArray(C.AsyncArray_New())
}

// Ptr returns the AsyncArray's underlying object pointer.
func (a *AsyncArray) Ptr() C.AsyncArray {
	return a.p
}

// Get async returns the Mat
func (m *AsyncArray) Get(mat *Mat) {
	C.AsyncArray_GetAsync(m.p, mat.p)
}

// newAsyncArray returns a new AsyncArray from a C AsyncArray
func newAsyncArray(p C.AsyncArray) AsyncArray {
	return AsyncArray{p: p}
}

// Close the AsyncArray object.
func (a *AsyncArray) Close() error {
	C.AsyncArray_Close(a.p)
	a.p = nil
	return nil
}
