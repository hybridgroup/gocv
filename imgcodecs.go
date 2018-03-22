package gocv

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import (
	"errors"
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

	// IMWriteJPEGQuality sets JPEG quality from 0 to 100 (the higher is the better). Default value is 95.
	IMWriteJPEGQuality = 1

	// IMWriteJPEGProgressive enables JPEG features, 0 or 1, default is False.
	IMWriteJPEGProgressive = 2

	// IMWriteJPEGOptimize enables JPEG features, 0 or 1, default is False.
	IMWriteJPEGOptimize = 3

	// IMWriteJPEGRstInterval sets JPEG restart interval, 0 - 65535, default is 0 - no restart.
	IMWriteJPEGRstInterval = 4

	// IMWriteJPEGLumaQuality sets separate luma quality level, 0 - 100, default is 0 - don't use.
	IMWriteJPEGLumaQuality = 5

	// IMWriteJPEGChromaQuality sets separate chroma quality level, 0 - 100, default is 0 - don't use.
	IMWriteJPEGChromaQuality = 6

	// IMWritePNGCompression for PNG, it can be the compression level from 0 to 9. A higher value
	// means a smaller size and longer compression time.
	// If specified, strategy is changed to IMWRITE_PNG_STRATEGY_DEFAULT (Z_DEFAULT_STRATEGY).
	// Default value is 1 (best speed setting).
	IMWritePngCompression = 16

	// IMWritePNGStrategy is one of IMWritePNGFlags, default is IMWRITE_PNG_STRATEGY_RLE.
	IMWritePNGStrategy = 17

	// IMWritePNGBilevel sets binary level PNG, 0 or 1, default is 0.
	IMWritePNGBilevel = 18

	// IMWritePxMBinary sets for PPM, PGM, or PBM, it can be a binary format flag, 0 or 1. Default value is 1.
	IMWritePxMBinary = 32

	// IMWriteWEBPQuality sets WEBP quality, it can be a quality from 1 to 100 (the higher is the better).
	// By default (without any parameter) and for quality above 100 the lossless compression is used.
	IMWriteWEBPQuality = 64

	// IMWritePAMTupletype for PAM, sets the TUPLETYPE field to the corresponding string value
	// that is defined for the format.
	IMWritePAMTupletype = 128

	// IMWritePNGStrategyDefault means use this value for normal data.
	IMWritePNGStrategyDefault = 0

	// IMWritePNGStrategyFiltered means use this value for data produced by a filter (or predictor).
	// Filtered data consists mostly of small values with a somewhat random distribution.
	// In this case, the compression algorithm is tuned to compress them better.
	IMWritePNGStrategyFiltered = 1

	// IMWritePNGStrategyHuffmanOnly means use this value to force Huffman encoding only (no string match).
	IMWritePNGStrategyHuffmanOnly = 2

	// IMWritePNGStrategyRle means use this value to limit match distances to one (run-length encoding).
	IMWritePNGStrategyRle = 3

	// IMWritePNGStrategyFixed means using this value prevents the use of dynamic Huffman codes,
	// allowing for a simpler decoder for special applications.
	IMWritePNGStrategyFixed = 4
)

// IMRead reads an image from a file into a Mat.
// The flags param is one of the IMReadFlag flags.
// If the image cannot be read (because of missing file, improper permissions,
// unsupported or invalid format), the function returns an empty Mat.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga288b8b3da0892bd651fce07b3bbd3a56
//
func IMRead(name string, flags IMReadFlag) (img Mat, err error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	img = Mat{p: C.Image_IMRead(cName, C.int(flags))}
	if img.Empty() {
		err = errors.New("IMRead error")
	}
	return
}

// IMWrite writes a Mat to an image file using default compression parameters.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#gabbc7ef1aa2edfaa87772f1202d67e0ce
//
func IMWrite(name string, img Mat) (err error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	if !bool(C.Image_IMWrite(cName, img.p)) {
		err = errors.New("IMWrite write error")
	}
	return
}

// IMWriteWithParams writes a Mat to an image file using custom compression parameters.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#gabbc7ef1aa2edfaa87772f1202d67e0ce
//
func IMWriteWithParams(name string, img Mat, params []int) (err error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cparams := []C.int{}

	for _, v := range params {
		cparams = append(cparams, C.int(v))
	}

	paramsVector := C.struct_IntVector{}
	paramsVector.val = (*C.int)(&cparams[0])
	paramsVector.length = (C.int)(len(cparams))

	if !bool(C.Image_IMWrite_WithParams(cName, img.p, paramsVector)) {
		err = errors.New("IMWrite write error")
	}
	return
}

type FileExt string

const (
	PNGFileExt  FileExt = ".png"
	JPEGFileExt FileExt = ".jpg"
	GIFFileExt  FileExt = ".gif"
)

// IMEncode encodes an image Mat into a memory buffer.
// This function compresses the image and stores it in the returned memory buffer,
// using the image format passed in in the form of a file extension string.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga461f9ac09887e47797a54567df3b8b63
//
func IMEncode(fileExt FileExt, img Mat) (buf []byte, err error) {
	cfileExt := C.CString(string(fileExt))
	defer C.free(unsafe.Pointer(cfileExt))

	b := C.Image_IMEncode(cfileExt, img.Ptr())
	defer C.ByteArray_Release(b)
	return toGoBytes(b), nil
}

// IMDecode reads an image from a buffer in memory.
// The function IMDecode reads an image from the specified buffer in memory.
// If the buffer is too short or contains invalid data, the function
// returns an empty matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga26a67788faa58ade337f8d28ba0eb19e
//
func IMDecode(buf []byte, flags IMReadFlag) (img Mat, err error) {
	data := toByteArray(buf)
	img = Mat{p: C.Image_IMDecode(data, C.int(flags))}
	if img.Empty() {
		err = errors.New("IMDecode error")
	}
	return
}
