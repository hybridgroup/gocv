package cuda

/*
#include <stdlib.h>
#include "../core.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Converts a OpenCVResult struct to an error.
func OpenCVResult(result C.OpenCVResult) error {
	if result.Code == 0 {
		return nil
	}
	if result.Message == nil {
		return errors.New("unknown openCV error")
	}
	defer C.free(unsafe.Pointer(result.Message))
	return errors.New(C.GoString(result.Message))
}
