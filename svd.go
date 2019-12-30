package gocv

/*
#include <stdlib.h>
#include "svd.h"
*/
import "C"

// SVDCompute decomposes matrix and stores the results to user-provided matrices
//
// https://docs.opencv.org/4.1.2/df/df7/classcv_1_1SVD.html#a76f0b2044df458160292045a3d3714c6
func SVDCompute(src Mat, w, u, vt *Mat) {
	C.SVD_Compute(src.Ptr(), w.Ptr(), u.Ptr(), vt.Ptr())
}
