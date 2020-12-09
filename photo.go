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
