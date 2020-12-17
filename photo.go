package gocv

/*
#include <stdlib.h>
#include "photo.h"
*/
import "C"
import "image"

//SeamlessCloneFlags seamlessClone algorithm flags
type SeamlessCloneFlags int

const (
	// NormalClone The power of the method is fully expressed when inserting objects with complex outlines into a new background.
	NormalClone SeamlessCloneFlags = iota

	// MixedClone The classic method, color-based selection and alpha masking might be time consuming and often leaves an undesirable halo. Seamless cloning, even averaged with the original image, is not effective. Mixed seamless cloning based on a loose selection proves effective.
	MixedClone

	// MonochromeTransfer Monochrome transfer allows the user to easily replace certain features of one object by alternative features.
	MonochromeTransfer
)

// ColorChange mix two differently colored versions of an image seamlessly.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#ga6684f35dc669ff6196a7c340dc73b98e
//
func ColorChange(src, mask Mat, dst *Mat, red_mul, green_mul, blue_mul float32) {
	C.ColorChange(src.p, mask.p, dst.p, C.float(red_mul), C.float(green_mul), C.float(blue_mul))
}

// SeamlessClone blend two image by Poisson Blending.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#ga2bf426e4c93a6b1f21705513dfeca49d
//
func SeamlessClone(src, dst, mask Mat, p image.Point, blend *Mat, flags SeamlessCloneFlags) {
	cp := C.struct_Point{
		x: C.int(p.X),
		y: C.int(p.Y),
	}

	C.SeamlessClone(src.p, dst.p, mask.p, cp, blend.p, C.int(flags))
}

// IlluminationChange modifies locally the apparent illumination of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#gac5025767cf2febd8029d474278e886c7
//
func IlluminationChange(src, mask Mat, dst *Mat, alpha, beta float32) {
	C.IlluminationChange(src.p, mask.p, dst.p, C.float(alpha), C.float(beta))
}

// TextureFlattening washes out the texture of the selected region, giving its contents a flat aspect.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#gad55df6aa53797365fa7cc23959a54004
//
func TextureFlattening(src, mask Mat, dst *Mat, lowThreshold, highThreshold float32, kernelSize int) {
	C.TextureFlattening(src.p, mask.p, dst.p, C.float(lowThreshold), C.float(highThreshold), C.int(kernelSize))
}
