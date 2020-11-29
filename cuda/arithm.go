package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "core.h"
#include "arithm.h"
*/
import "C"

// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga40f1c94ae9a9456df3cad48e3cb008e1
//
func Threshold(src GpuMat, dst *GpuMat, thresh, maxval float64, typ int) {
	C.GpuThreshold(src.p, dst.p, C.double(thresh), C.double(maxval), C.int(typ))
}
