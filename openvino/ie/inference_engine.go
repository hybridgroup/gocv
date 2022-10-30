package ie

/*
#include <stdlib.h>
#include "inference_engine.h"
*/
import (
	"C"
)
import "unsafe"

// Version returns the current Inference Engine library version
func Version() string {
	v := C.OpenVinoVersion()
	defer C.free(unsafe.Pointer(v))
	return C.GoString(v)
}
