package contrib

/*
#include <stdlib.h>
#include "signal.h"
*/
import "C"
import "gocv.io/x/gocv"

// ResampleSignal Signal resampling.
//
// Parameters:
//
// inputSignal: gocv.Mat with input signal.
//
// inFreq: Input signal frequency.
//
// outFreq: Output signal frequency. Signal resampling implemented a cubic interpolation function
// and a filtering function based on Kaiser window and Bessel function, used to construct a FIR filter.
// Result is similar to scipy.signal.resample.
//
// Returns a gocv.Mat with output signal.
//
// For further details, please see:
// https://docs.opencv.org/4.11.0/d2/d1f/group__signal.html#gad32480ac5c9832126cf0c6bde026686d
func ResampleSignal(inputSignal gocv.Mat, inFreq int, outFreq int) gocv.Mat {
	out := gocv.NewMat()

	C.Signal_ResampleSignal(C.Mat(inputSignal.Ptr()), C.Mat(out.Ptr()), C.int(inFreq), C.int(outFreq))

	return out
}
