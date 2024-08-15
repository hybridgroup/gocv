//go:build linux

package contrib

/*
#include <stdlib.h>
#include "freetype.h"
*/
import "C"
import (
	"image"
	"image/color"
	"unsafe"

	"gocv.io/x/gocv"
)

type FreeType2 struct {
	// C.FreeType2
	p unsafe.Pointer
}

// NewFreeType2 create instance to draw UTF-8 strings.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/dfc/group__freetype.html#ga0fd8f9c0ae69bb4d95c41af79473a756
func NewFreeType2() FreeType2 {
	return FreeType2{p: unsafe.Pointer(C.FreeType2_CreateFreeType2())}
}

// Close FreeType2.
func (f *FreeType2) Close() error {
	C.FreeType2_Close((C.FreeType2)(f.p))
	f.p = nil
	return nil
}

// LoadFontData loads font data.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/dfa/classcv_1_1freetype_1_1FreeType2.html#af059d49b806b916ffdd6380b9eb2f59a
func (f *FreeType2) LoadFontData(fontFileName string, id int) {
	cFontFileName := C.CString(fontFileName)
	defer C.free(unsafe.Pointer(cFontFileName))
	C.FreeType2_LoadFontData((C.FreeType2)(f.p), cFontFileName, C.int(id))
}

// SetSplitNumber set the number of split points from bezier-curve to line.
// If you want to draw large glyph, large is better.
// If you want to draw small glyph, small is better.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/dfa/classcv_1_1freetype_1_1FreeType2.html#a572143e6c68eab181387d9f4b3366f8b
func (f *FreeType2) SetSplitNumber(num int) {
	C.FreeType2_SetSplitNumber((C.FreeType2)(f.p), C.int(num))
}

// PutText draws a text string.
// It renders the specified text string in the image.
// Symbols that cannot be rendered using the specified font are replaced by "Tofu" or non-drawn.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/dfa/classcv_1_1freetype_1_1FreeType2.html#aba641f774c47a70eaeb76bf7aa865915
func (f *FreeType2) PutText(img *gocv.Mat, text string, org image.Point,
	fontHeight int, c color.RGBA, thickness int, lineType gocv.LineType, bottomLeftOrigin bool) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.FreeType2_PutText((C.FreeType2)(f.p), (C.Mat)(img.Ptr()), cText, sOrg, C.int(fontHeight), sColor, C.int(thickness), C.int(lineType), C.bool(bottomLeftOrigin))
}

// GetTextSize calculates the width and height of a text string.
// The function getTextSize calculates and returns the approximate size of a box that contains the specified text.
// That is, the following code renders some text, the tight box surrounding it, and the baseline.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/dfa/classcv_1_1freetype_1_1FreeType2.html#af135a132505125bdea74b378dda3bb5d
func (f *FreeType2) GetTextSize(text string, fontHeight int, thickness int) (image.Point, int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cBaseLine := C.int(0)

	sz := C.FreeType2_GetTextSize((C.FreeType2)(f.p), cText, C.int(fontHeight), C.int(thickness), &cBaseLine)
	return image.Point{
		X: int(sz.width),
		Y: int(sz.height),
	}, int(cBaseLine)
}
