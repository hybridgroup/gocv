package cuda

/*
#include <stdlib.h>
#include "cudawarping.h"
*/
import "C"
import "image"

// InterpolationFlags are bit flags that control the interpolation algorithm
// that is used.
type InterpolationFlags int

const (
	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax = 7
)

// Resize resizes an image.
// It resizes the image src down to or up to the specified size, storing the
// result in dst. Note that src and dst may be the same image. If you wish to
// scale by factor, an empty sz may be passed and non-zero fx and fy. Likewise,
// if you wish to scale to an explicit size, a non-empty sz may be passed with
// zero for both fx and fy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga47a974309e9102f5f08231edc7e7529d
func Resize(src GpuMat, dst *GpuMat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.CudaResize(src.p, dst.p, pSize, C.double(fx), C.double(fy), C.int(interp))
	return
}
