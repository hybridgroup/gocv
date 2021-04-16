package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "core.h"
#include "arithm.h"
*/
import "C"

import "gocv.io/x/gocv"

// Abs computes an absolute value of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga54a72bd772494ab34d05406fd76df2b6
//
func Abs(src GpuMat, dst *GpuMat) {
	C.GpuAbs(src.p, dst.p)
}

// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga40f1c94ae9a9456df3cad48e3cb008e1
//
func Threshold(src GpuMat, dst *GpuMat, thresh, maxval float64, typ gocv.ThresholdType) {
	C.GpuThreshold(src.p, dst.p, C.double(thresh), C.double(maxval), C.int(typ))
}

// Flip flips a 2D matrix around vertical, horizontal, or both axes.
//
// For further details, please see:
// https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga4d0a3f2b46e8f0f1ec2b5ac178dcd871
//
func Flip(src GpuMat, dst *GpuMat, flipCode int) {
	C.GpuFlip(src.p, dst.p, C.int(flipCode))
}
