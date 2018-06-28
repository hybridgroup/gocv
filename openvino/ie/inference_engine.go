package ie

/*
#include <stdlib.h>
#include "inference_engine.h"
*/
import (
	"C"
)

// Version returns the current Inference Engine library version
func Version() string {
	return C.GoString(C.OpenVinoVersion())
}
