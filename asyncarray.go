// +build openvino

package gocv

import (
	"errors"
)

/*
#include <stdlib.h>
#include "dnn.h"
#include "asyncarray.h"
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
func (m *AsyncArray) Get(mat *Mat) error {
	result := C.AsyncArray_GetAsync(m.p, mat.p)
	err := C.GoString(result)

	if len(err) > 0 {
		return errors.New(err)
	}
	return nil
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
