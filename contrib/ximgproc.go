package contrib

/*
#include <stdlib.h>
#include "ximgproc.h"
*/
import "C"
import (
	"gocv.io/x/gocv"
)

// AnisotropicDiffusion performs anisotropic diffusion on an image.
//
// The function applies Perona-Malik anisotropic diffusion to an image.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d2d/group__ximgproc.html#gaffedd976e0a8efb5938107acab185ec2
//
func AnisotropicDiffusion(src gocv.Mat, dst *gocv.Mat, alpha float32, k float32, niters int) {
	C.anisotropicDiffusion(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.float(alpha), C.float(k), C.int(niters))
}

// EdgePreservingFilter smoothes an image using the Edge-Preserving filter.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d2d/group__ximgproc.html#ga86fcda65ced0aafa2741088d82e9161c
//
func EdgePreservingFilter(src gocv.Mat, dst *gocv.Mat, d int, threshold float32) {
	C.edgePreservingFilter(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.int(d), C.float(threshold))
}

type BinarizationMethod int

const (
	BinarizationNiblack BinarizationMethod = iota
	BinarizationSauvola
	BinarizationWolf
	BinarizationNICK
)

// NiblackThreshold performs thresholding on input images using Niblack's technique 
// or some of the popular variations it inspired.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d2d/group__ximgproc.html#gab042a5032bbb85275f1fd3e04e7c7660
//
func NiblackThreshold(src gocv.Mat, dst *gocv.Mat, maxValue float32,
	typ gocv.ThresholdType, blockSize int, k float32, binarizationMethod BinarizationMethod, r float32) {
	C.niBlackThreshold(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.float(maxValue), C.int(typ),
		C.int(blockSize), C.float(k), C.int(binarizationMethod), C.float(r))
}

// PeiLinNormalization calculates an affine transformation that normalize
// given image using Pei&Lin Normalization.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d2d/group__ximgproc.html#ga50d064b92f63916f4162474eea22d656
//
func PeiLinNormalization(src gocv.Mat, dst *gocv.Mat) {
	C.PeiLinNormalization(C.Mat(src.Ptr()), C.Mat(dst.Ptr()))
}

type ThinningType int

const (
	ThinningZhangSuen ThinningType = iota
	ThinningGuoHall
)

// Thinning applies a binary blob thinning operation, to achieve a skeletization 
// of the input image.
//
// The function transforms a binary blob image into a skeletized form using the
// technique of Zhang-Suen.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/d2d/group__ximgproc.html#ga37002c6ca80c978edb6ead5d6b39740c
//
func Thinning(src gocv.Mat, dst *gocv.Mat, typ ThinningType) {
	C.thinning(C.Mat(src.Ptr()), C.Mat(dst.Ptr()), C.int(typ))
}
