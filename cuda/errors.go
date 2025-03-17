package cuda

/*
#include "../core.h"
*/
import "C"
import "errors"

// Converts a OpenCVResult struct to an error.
func OpenCVResult(result C.OpenCVResult) error {
	if result.Code == 0 {
		return nil
	}
	return errors.New(C.GoString(result.Message))
}
