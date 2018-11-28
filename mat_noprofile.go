// +build !matprofile

package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"

// newMat returns a new Mat from a C Mat
func newMat(p C.Mat) Mat {
	return Mat{p: p}
}

// Close the Mat object.
func (m *Mat) Close() error {
	C.Mat_Close(m.p)
	m.p = nil
	return nil
}
