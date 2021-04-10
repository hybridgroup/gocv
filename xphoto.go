package gocv

/*
#include <stdlib.h>
#include "xphoto.h"
*/
import "C"
import (
	"unsafe"
)

// GrayworldWB is a wrapper around the cv::GrayworldWB.
type GrayworldWB struct {
	// C.GrayworldWB
	p unsafe.Pointer
}

// NewGrayworldWBWithParams returns a new Gray-world white balance algorithm.
// of type GrayworldWB with customized parameters. GrayworldWB algorithm scales the values
// of pixels based on a gray-world assumption which states that the average of all
// channels should result in a gray image.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
//
func NewGrayworldWB() GrayworldWB {
	return GrayworldWB{p: unsafe.Pointer(C.GrayworldWB_Create())}
}

// SetSaturationThreshold set a Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) SetSaturationThreshold(saturationThreshold float32) {
	C.GrayworldWB_SetSaturationThreshold((C.GrayworldWB)(b.p), C.float(saturationThreshold))
	return
}

// GetSaturationThreshold return the Maximum saturation for a pixel to be included in the gray-world assumption.
//
// For further details, please see:
// https://docs.opencv.org/master/de/daa/group__xphoto.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#ac6e17766e394adc15588b8522202cc71
//
func (b *GrayworldWB) GetSaturationThreshold() float32 {
	return float32(C.GrayworldWB_GetSaturationThreshold((C.GrayworldWB)(b.p)))
}

// Close GrayworldWB.
func (b *GrayworldWB) Close() error {
	C.GrayworldWB_Close((C.GrayworldWB)(b.p))
	b.p = nil
	return nil
}

// BalanceWhite computes a Gray-world white balance using the current GrayworldWB.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d71/classcv_1_1xphoto_1_1GrayworldWB.html#details
//
func (b *GrayworldWB) BalanceWhite(src Mat, dst *Mat) {
	C.GrayworldWB_BalanceWhite((C.GrayworldWB)(b.p), src.p, dst.p)
	return
}
