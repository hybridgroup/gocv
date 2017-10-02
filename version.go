package gocv

/*
#include <stdlib.h>
#include "version.h"
*/
import "C"

const version = "0.0.1"

// Version returns the current golang package version
func Version() string {
	return version
}

// OpenCVVersion returns the current OpenCV lib version
func OpenCVVersion() string {
	return C.GoString(C.openCVVersion())
}
