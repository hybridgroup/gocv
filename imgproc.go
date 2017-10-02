package opencv3

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import (
	"image"
	"image/color"
	"unsafe"
)

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
//
func CvtColor(src Mat, dst Mat, code ColorConversionCode) {
	C.CvtColor(src.p, dst.p, C.int(code))
}

// GaussianBlur blurs an image Mat using a Gaussian filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
//
func GaussianBlur(src Mat, dst Mat, ksize image.Point, sigmaX float64,
	sigmaY float64, borderType int) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}

	C.GaussianBlur(src.p, dst.p, pSize, C.double(sigmaX), C.double(sigmaY), C.int(borderType))
}

// Rectangle draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
//
func Rectangle(img Mat, r image.Rectangle, c color.RGBA) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Rectangle(img.p, cRect, sColor)
}

// HersheyFont is based on the enum HersheyFonts
// Only a subset of the available Hershey fonts
// <http://sources.isc.org/utils/misc/hershey-font.txt> are supported
type HersheyFont int

const (
	// normal size sans-serif font
	FontHersheySimplex HersheyFont = 0
	// small size sans-serif font
	FontHersheyPlain = 1
	// normal size sans-serif font (more complex than FontHersheySIMPLEX)
	FontHersheyDuplex = 2
	// normal size serif font
	FontHersheyComplex = 3
	// normal size serif font (more complex than FontHersheyCOMPLEX)
	FontHersheyTriplex = 4
	// smaller version of FontHersheyCOMPLEX
	FontHersheyComplexSmall = 5
	// hand-writing style font
	FontHersheyScriptSimplex = 6
	// more complex variant of FontHersheySCRIPT_SIMPLEX
	FontHersheyScriptComplex = 7
	// flag for italic font
	FontItalic = 16
)

// GetTextSize calculates the width and height of a text string.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
//
func GetTextSize(text string, fontFace HersheyFont, fontScale float64,
	thickness int) image.Point {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sz := C.GetTextSize(cText, C.int(fontFace), C.double(fontScale), C.int(thickness))
	return image.Pt(int(sz.width), int(sz.height))
}

// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
//
func PutText(img Mat, text string, org image.Point, fontFace HersheyFont,
	fontScale float64, c color.RGBA, thickness int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutText(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness))
	return
}
