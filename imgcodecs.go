package opencv3

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import (
	"unsafe"
)

// IMRead flags
const (
	// IMReadUnchanged return the loaded image as is (with alpha channel, otherwise it gets cropped).
	IMReadUnchanged = -1

	// IMReadGrayScale always convert image to the single channel grayscale image.
	IMReadGrayScale = 0

	// If set, always convert image to the 3 channel BGR color image.
	IMReadColor = 1

	// If set, return 16-bit/32-bit image when the input has the corresponding depth, otherwise convert it to 8-bit.
	IMReadAnyDepth = 2

	// If set, the image is read in any possible color format.
	IMReadAnyColor = 4

	// If set, use the gdal driver for loading the image.
	IMReadLoadGDAL = 8

	// If set, always convert image to the single channel grayscale image and the image size reduced 1/2.
	IMReadReducedGrayscale2 = 16

	// If set, always convert image to the 3 channel BGR color image and the image size reduced 1/2.
	IMReadReducedColor2 = 17

	// If set, always convert image to the single channel grayscale image and the image size reduced 1/4.
	IMReadReducedGrayscale4 = 32

	// If set, always convert image to the 3 channel BGR color image and the image size reduced 1/4.
	IMReadReducedColor4 = 33

	// If set, always convert image to the single channel grayscale image and the image size reduced 1/8.
	IMReadReducedGrayscale8 = 64

	// If set, always convert image to the 3 channel BGR color image and the image size reduced 1/8.
	IMReadReducedColor8 = 65

	// If set, do not rotate the image according to EXIF's orientation flag.
	IMReadIgnoreOrientation = 128
)

// IMRead reads an image file into a Mat. The flags param is one of the IMRead flags.
func IMRead(name string, flags int) Mat {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return Mat{p: C.Image_IMRead(cName, C.int(flags))}
}

// IMWrite writes a Mat to an image file
func IMWrite(name string, img Mat) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Image_IMWrite(cName, img.p))
}
