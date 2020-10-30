package cuda

/*
#include <stdlib.h>
#include "cudawarping.h"
*/
import "C"
import (
	"image"
	"image/color"
)

// InterpolationFlags are bit flags that control the interpolation algorithm
// that is used.
type InterpolationFlags int

const (
	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear InterpolationFlags = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic InterpolationFlags = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea InterpolationFlags = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 InterpolationFlags = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax InterpolationFlags = 7
)

// BorderType type of border.
type BorderType int

const (
	// BorderConstant border type
	BorderConstant BorderType = 0

	// BorderReplicate border type
	BorderReplicate BorderType = 1

	// BorderReflect border type
	BorderReflect BorderType = 2

	// BorderWrap border type
	BorderWrap BorderType = 3

	// BorderReflect101 border type
	BorderReflect101 BorderType = 4

	// BorderTransparent border type
	BorderTransparent BorderType = 5

	// BorderDefault border type
	BorderDefault = BorderReflect101

	// BorderIsolated border type
	BorderIsolated BorderType = 16
)

// Resize resizes an image.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga4f5fa0770d1c9efbadb9be1b92a6452a
func Resize(src GpuMat, dst *GpuMat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.CudaResize(src.p, dst.p, pSize, C.double(fx), C.double(fy), C.int(interp))
}

// Rotate rotates an image around the origin (0,0) and then shifts it.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga55d958eceb0f871e04b1be0adc6ef1b5
func Rotate(src GpuMat, dst *GpuMat, sz image.Point, angle, xShift, yShift float64, interp InterpolationFlags) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	C.CudaRotate(src.p, dst.p, pSize, C.double(angle), C.double(xShift), C.double(yShift), C.int(interp))
}

// Remap applies a generic geometrical transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga0ece6c76e8efa3171adb8432d842beb0
func Remap(src GpuMat, dst, xmap, ymap *GpuMat, interpolation InterpolationFlags, borderMode BorderType, borderValue color.RGBA) {
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}
	C.CudaRemap(src.p, dst.p, xmap.p, ymap.p, C.int(interpolation), C.int(borderMode), bv)
}

// PyrDown blurs an image and downsamples it.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga9c8456de9792d96431e065f407c7a91b
func PyrDown(src GpuMat, dst *GpuMat) {
	C.CudaPyrDown(src.p, dst.p)
}

// PyrUp upsamples an image and then blurs it.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga2048da0dfdb9e4a726232c5cef7e5747
func PyrUp(src GpuMat, dst *GpuMat) {
	C.CudaPyrUp(src.p, dst.p)
}

// WarpPerspective applies a perspective transformation to an image.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga7a6cf95065536712de6b155f3440ccff
func WarpPerspective(src GpuMat, dst *GpuMat, m GpuMat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}

	C.CudaWarpPerspective(src.p, dst.p, m.p, pSize, C.int(flags), C.int(borderType), bv)
}

// WarpAffine applies an affine transformation to an image. For more parameters please check WarpAffineWithParams
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga9e8dd9e73b96bdc8e27d85c0e83f1130
func WarpAffine(src GpuMat, dst *GpuMat, m GpuMat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}
	bv := C.struct_Scalar{
		val1: C.double(borderValue.B),
		val2: C.double(borderValue.G),
		val3: C.double(borderValue.R),
		val4: C.double(borderValue.A),
	}

	C.CudaWarpAffine(src.p, dst.p, m.p, pSize, C.int(flags), C.int(borderType), bv)
}

// BuildWarpAffineMaps builds transformation maps for affine transformation.
//
// For further details. please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga63504590a96e4cc702d994281d17bc1c
func BuildWarpAffineMaps(M GpuMat, inverse bool, sz image.Point, xmap, ymap *GpuMat) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.CudaBuildWarpAffineMaps(M.p, C.bool(inverse), pSize, xmap.p, ymap.p)
}

// BuildWarpPerspectiveMaps builds transformation maps for perspective transformation.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d29/group__cudawarping.html#ga8d16e3003703bd3b89cca98c913ef864
func BuildWarpPerspectiveMaps(M GpuMat, inverse bool, sz image.Point, xmap, ymap *GpuMat) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.CudaBuildWarpPerspectiveMaps(M.p, C.bool(inverse), pSize, xmap.p, ymap.p)
}
