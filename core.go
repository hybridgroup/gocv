package opencv3

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"unsafe"
)

const (
	// CvCapPropFrameWidth is OpenCV parameter of Frame Width
	CvCapPropFrameWidth = 3
	// CvCapPropFrameHeight is OpenCV parameter of Frame Height
	CvCapPropFrameHeight = 4
	// CvCapPropFps is OpenCV parameter of FPS
	CvCapPropFps = 5
)

// CMat is an alias for C pointer.
type CMat C.Mat

// Mat is a bind of `cv::Mat
type Mat struct {
	p C.Mat
}

// NewMat returns a new Mat.
func NewMat() Mat {
	return Mat{p: C.Mat_New()}
}

// Delete object.
func (m *Mat) Delete() {
	C.Mat_Delete(m.p)
	m.p = nil
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Rect represents rectangle. X and Y is a start point of Width and Height.
type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

// DrawRectsToImage draws rectangle information to target image Mat.
func DrawRectsToImage(img Mat, rects []Rect) {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.X),
			y:      C.int(r.Y),
			width:  C.int(r.Width),
			height: C.int(r.Height),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}
	C.DrawRectsToImage(img.p, cRects)
}

// CMatVec3b is an alias for C pointer.
type CMatVec3b C.MatVec3b

// MatVec3b is a bind of `cv::Mat_<cv::Vec3b>`
type MatVec3b struct {
	p C.MatVec3b
}

// GetCPointer returns C pointer of MatVec3b.
func (m *MatVec3b) GetCPointer() C.MatVec3b {
	return m.p
}

// NewMatVec3b returns a new MatVec3b.
func NewMatVec3b() MatVec3b {
	return MatVec3b{p: C.MatVec3b_New()}
}

// NewMatVec3bWithCPointer return a new MatVec3b with argument C pointer.
func NewMatVec3bWithCPointer(p CMatVec3b) MatVec3b {
	return MatVec3b{p: C.MatVec3b(p)}
}

// ToJpegData convert to JPEG data.
func (m *MatVec3b) ToJpegData(quality int) []byte {
	b := C.MatVec3b_ToJpegData(m.p, C.int(quality))
	defer C.ByteArray_Release(b)
	return toGoBytes(b)
}

// Delete object.
func (m *MatVec3b) Delete() {
	C.MatVec3b_Delete(m.p)
	m.p = nil
}

// CopyTo copies MatVec3b.
func (m *MatVec3b) CopyTo(dst *MatVec3b) {
	C.MatVec3b_CopyTo(m.p, dst.p)
}

// Empty returns the MatVec3b is empty or not.
func (m *MatVec3b) Empty() bool {
	isEmpty := C.MatVec3b_Empty(m.p)
	return isEmpty != 0
}

// ToRawData converts MatVec3b to RawData.
func (m *MatVec3b) ToRawData() (int, int, []byte) {
	r := C.MatVec3b_ToRawData(m.p)
	return int(r.width), int(r.height), toGoBytes(r.data)
}

// ToMatVec3b converts RawData to MatVec3b. Returned MatVec3b is required to
// delete after using.
func ToMatVec3b(width int, height int, data []byte) MatVec3b {
	cr := C.struct_RawData{
		width:  C.int(width),
		height: C.int(height),
		data:   toByteArray(data),
	}
	return MatVec3b{p: C.RawData_ToMatVec3b(cr)}
}

// MatVec4b is a bind of `cv::Mat_<cv::Vec4b>`
type MatVec4b struct {
	p C.MatVec4b
}

// Delete object.
func (m *MatVec4b) Delete() {
	C.MatVec4b_Delete(m.p)
	m.p = nil
}

// ToRawData converts MatVec4b to RawData.
func (m *MatVec4b) ToRawData() (int, int, []byte) {
	r := C.MatVec4b_ToRawData(m.p)
	return int(r.width), int(r.height), toGoBytes(r.data)
}

// ToMatVec4b converts RawData to MatVec4b. Returned MatVec4b is required to
// delete after using.
func ToMatVec4b(width int, height int, data []byte) MatVec4b {
	cr := C.struct_RawData{
		width:  C.int(width),
		height: C.int(height),
		data:   toByteArray(data),
	}
	return MatVec4b{p: C.RawData_ToMatVec4b(cr)}
}

// LoadAlphaImage loads RGBA type image.
func LoadAlphaImage(name string) MatVec4b {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return MatVec4b{p: C.LoadAlphaImg(cName)}
}

// MountAlphaImage draws img on back leading to rects. img is required RGBA,
// TODO should be check file type.
func MountAlphaImage(img MatVec4b, back MatVec3b, rects []Rect) {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.X),
			y:      C.int(r.Y),
			width:  C.int(r.Width),
			height: C.int(r.Height),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}
	C.MountAlphaImage(img.p, back.p, cRects)
}

func toByteArray(b []byte) C.struct_ByteArray {
	return C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}
}

// toGoBytes returns binary data. Serializing is depends on C/C++ implementation.
func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}
