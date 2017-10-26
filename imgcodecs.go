package gocv

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import (
	"unsafe"
)

// IMReadFlag is one of the valid flags to use for the IMRead function.
type IMReadFlag int

const (
	// IMReadUnchanged return the loaded image as is (with alpha channel,
	// otherwise it gets cropped).
	IMReadUnchanged IMReadFlag = -1

	// IMReadGrayScale always convert image to the single channel
	// grayscale image.
	IMReadGrayScale = 0

	// IMReadColor always converts image to the 3 channel BGR color image.
	IMReadColor = 1

	// IMReadAnyDepth returns 16-bit/32-bit image when the input has the corresponding
	// depth, otherwise convert it to 8-bit.
	IMReadAnyDepth = 2

	// IMReadAnyColor the image is read in any possible color format.
	IMReadAnyColor = 4

	// IMReadLoadGDAL uses the gdal driver for loading the image.
	IMReadLoadGDAL = 8

	// IMReadReducedGrayscale2 always converts image to the single channel grayscale image
	// and the image size reduced 1/2.
	IMReadReducedGrayscale2 = 16

	// IMReadReducedColor2 always converts image to the 3 channel BGR color image and the
	// image size reduced 1/2.
	IMReadReducedColor2 = 17

	// IMReadReducedGrayscale4 always converts image to the single channel grayscale image and
	// the image size reduced 1/4.
	IMReadReducedGrayscale4 = 32

	// IMReadReducedColor4 always converts image to the 3 channel BGR color image and
	// the image size reduced 1/4.
	IMReadReducedColor4 = 33

	// IMReadReducedGrayscale8 always convert image to the single channel grayscale image and
	// the image size reduced 1/8.
	IMReadReducedGrayscale8 = 64

	// IMReadReducedColor8 always convert image to the 3 channel BGR color image and the
	// image size reduced 1/8.
	IMReadReducedColor8 = 65

	// IMReadIgnoreOrientation do not rotate the image according to EXIF's orientation flag.
	IMReadIgnoreOrientation = 128
)

// IMRead reads an image from a file into a Mat.
// The flags param is one of the IMReadFlag flags.
// If the image cannot be read (because of missing file, improper permissions,
// unsupported or invalid format), the function returns an empty Mat.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d4/da8/group__imgcodecs.html#ga288b8b3da0892bd651fce07b3bbd3a56
//
func IMRead(name string, flags IMReadFlag) Mat {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return Mat{p: C.Image_IMRead(cName, C.int(flags))}
}

// IMWrite writes a Mat to an image file.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d4/da8/group__imgcodecs.html#gabbc7ef1aa2edfaa87772f1202d67e0ce
//
func IMWrite(name string, img Mat) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Image_IMWrite(cName, img.p))
}

// IMEncode encodes an image Mat into a memory buffer.
// This function compresses the image and stores it in the returned memory buffer,
// using the image format passed in in the form of a file extension string.
//
// For further details, please see:
// http://docs.opencv.org/3.3.1/d4/da8/group__imgcodecs.html#ga461f9ac09887e47797a54567df3b8b63
//
func IMEncode(fileExt string, img Mat) (buf []byte, err error) {
	cfileExt := C.CString(fileExt)
	defer C.free(unsafe.Pointer(cfileExt))

	b := C.Image_IMEncode(cfileExt, img.Ptr())
	defer C.ByteArray_Release(b)
	return toGoBytes(b), nil
}
