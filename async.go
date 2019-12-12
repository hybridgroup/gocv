// +build openvino

package gocv

import (
	"unsafe"
)

/*
#include <stdlib.h>
#include "dnn.h"
#include "asyncarray.h"
*/
import "C"

// ForwardAsync runs forward pass to compute output of layer with name outputName.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a814890154ea9e10b132fec00b6f6ba30
//
func (net *Net) ForwardAsync(outputName string) AsyncArray {
	cName := C.CString(outputName)
	defer C.free(unsafe.Pointer(cName))

	return newAsyncArray(C.Net_forwardAsync((C.Net)(net.p), cName))
}
