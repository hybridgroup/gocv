package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"errors"
	"image"
	"image/color"
	"reflect"
	"unsafe"
)

const (
	// MatChannels1 is a single channel Mat.
	MatChannels1 = 0

	// MatChannels2 is 2 channel Mat.
	MatChannels2 = 8

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24
)

// MatType is the type for the various different kinds of Mat you can create.
type MatType int

const (
	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U MatType = 0

	// MatTypeCV8S is a Mat of 8-bit signed int
	MatTypeCV8S = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S = 3

	// MatTypeCV16SC2 is a Mat of 16-bit signed int with 2 channels
	MatTypeCV16SC2 = MatTypeCV16S + MatChannels2

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F = 6

	// MatTypeCV8UC1 is a Mat of 8-bit unsigned int with a single channel
	MatTypeCV8UC1 = MatTypeCV8U + MatChannels1

	// MatTypeCV8UC2 is a Mat of 8-bit unsigned int with 2 channels
	MatTypeCV8UC2 = MatTypeCV8U + MatChannels2

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4
)

// CompareType is used for Compare operations to indicate which kind of
// comparison to use.
type CompareType int

const (
	// CompareEQ src1 is equal to src2.
	CompareEQ CompareType = 0

	// CompareGT src1 is greater than src2.
	CompareGT = 1

	// CompareGE src1 is greater than or equal to src2.
	CompareGE = 2

	// CompareLT src1 is less than src2.
	CompareLT = 3

	// CompareLE src1 is less than or equal to src2.
	CompareLE = 4

	// CompareNE src1 is unequal to src2.
	CompareNE = 5
)

var ErrEmptyByteSlice = errors.New("empty byte array")

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html
//
type Mat struct {
	p C.Mat
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return newMat(C.Mat_New())
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return newMat(C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt)))
}

// NewMatFromScalar returns a new Mat for a specific Scalar value
func NewMatFromScalar(s Scalar, mt MatType) Mat {
	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	return newMat(C.Mat_NewFromScalar(sVal, C.int(mt)))
}

// NewMatWithSizeFromScalar returns a new Mat for a specific Scala value with a specific size and type
// This simplifies creation of specific color filters or creating Mats of specific colors and sizes
func NewMatWithSizeFromScalar(s Scalar, rows int, cols int, mt MatType) Mat {
	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	return newMat(C.Mat_NewWithSizeFromScalar(sVal, C.int(rows), C.int(cols), C.int(mt)))
}

// NewMatFromBytes returns a new Mat with a specific size and type, initialized from a []byte.
func NewMatFromBytes(rows int, cols int, mt MatType, data []byte) (Mat, error) {
	cBytes, err := toByteArray(data)
	if err != nil {
		return Mat{}, err
	}
	return newMat(C.Mat_NewFromBytes(C.int(rows), C.int(cols), C.int(mt), *cBytes)), nil
}

// FromPtr returns a new Mat with a specific size and type, initialized from a Mat Ptr.
func (m *Mat) FromPtr(rows int, cols int, mt MatType, prow int, pcol int) (Mat, error) {
	return newMat(C.Mat_FromPtr(m.p, C.int(rows), C.int(cols), C.int(mt), C.int(prow), C.int(pcol))), nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return m.p
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Clone returns a cloned full copy of the Mat.
func (m *Mat) Clone() Mat {
	return newMat(C.Mat_Clone(m.p))
}

// CopyTo copies Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a33fd5d125b4c302b0c9aa86980791a77
//
func (m *Mat) CopyTo(dst *Mat) {
	C.Mat_CopyTo(m.p, dst.p)
	return
}

// CopyToWithMask copies Mat into destination Mat after applying the mask Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a626fe5f96d02525e2604d2ad46dd574f
//
func (m *Mat) CopyToWithMask(dst *Mat, mask Mat) {
	C.Mat_CopyToWithMask(m.p, dst.p, mask.p)
	return
}

// ConvertTo converts Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b
//
func (m *Mat) ConvertTo(dst *Mat, mt MatType) {
	C.Mat_ConvertTo(m.p, dst.p, C.int(mt))
	return
}

// Total returns the total number of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa4d317d43fb0cba9c2503f3c61b866c8
//
func (m *Mat) Total() int {
	return int(C.Mat_Total(m.p))
}

// Size returns an array with one element for each dimension containing the size of that dimension for the Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa4d317d43fb0cba9c2503f3c61b866c8
//
func (m *Mat) Size() (dims []int) {
	cdims := C.IntVector{}
	C.Mat_Size(m.p, &cdims)

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cdims.val)),
		Len:  int(cdims.length),
		Cap:  int(cdims.length),
	}
	pdims := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(cdims.length); i++ {
		dims = append(dims, int(pdims[i]))
	}
	return
}

// ToBytes copies the underlying Mat data to a byte array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d3/d63/classcv_1_1Mat.html#a4d33bed1c850265370d2af0ff02e1564
func (m *Mat) ToBytes() []byte {
	b := C.Mat_DataPtr(m.p)
	return toGoBytes(b)
}

// DataPtrUint8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint8() []uint8 {
	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length),
		Cap:  int(p.length),
	}
	return *(*[]uint8)(unsafe.Pointer(h))
}

// DataPtrInt8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrInt8() []int8 {
	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length),
		Cap:  int(p.length),
	}
	return *(*[]int8)(unsafe.Pointer(h))
}

// DataPtrUint16 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint16() ([]uint16, error) {
	if m.Type()&MatTypeCV16U != MatTypeCV16U {
		return nil, errors.New("DataPtrUint16 only supports MatTypeCV16U")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length) / 2,
		Cap:  int(p.length) / 2,
	}
	return *(*[]uint16)(unsafe.Pointer(h)), nil
}

// DataPtrInt16 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrInt16() ([]int16, error) {
	if m.Type()&MatTypeCV16S != MatTypeCV16S {
		return nil, errors.New("DataPtrInt16 only supports MatTypeCV16S")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length) / 2,
		Cap:  int(p.length) / 2,
	}
	return *(*[]int16)(unsafe.Pointer(h)), nil
}

// DataPtrFloat32 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrFloat32() ([]float32, error) {
	if m.Type()&MatTypeCV32F != MatTypeCV32F {
		return nil, errors.New("DataPtrFloat32 only supports MatTypeCV32F")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length) / 4,
		Cap:  int(p.length) / 4,
	}
	return *(*[]float32)(unsafe.Pointer(h)), nil
}

// DataPtrFloat64 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrFloat64() ([]float64, error) {
	if m.Type()&MatTypeCV64F != MatTypeCV64F {
		return nil, errors.New("DataPtrFloat64 only supports MatTypeCV64F")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length) / 8,
		Cap:  int(p.length) / 8,
	}
	return *(*[]float64)(unsafe.Pointer(h)), nil
}

// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, since they are pointers to the underlying
// OpenCV Mat object.
func (m *Mat) Region(rio image.Rectangle) Mat {
	cRect := C.struct_Rect{
		x:      C.int(rio.Min.X),
		y:      C.int(rio.Min.Y),
		width:  C.int(rio.Size().X),
		height: C.int(rio.Size().Y),
	}

	return newMat(C.Mat_Region(m.p, cRect))
}

// Reshape changes the shape and/or the number of channels of a 2D matrix without copying the data.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8
//
func (m *Mat) Reshape(cn int, rows int) Mat {
	return newMat(C.Mat_Reshape(m.p, C.int(cn), C.int(rows)))
}

// ConvertFp16 converts a Mat to half-precision floating point.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9c25d9ef44a2a48ecc3774b30cb80082
//
func (m *Mat) ConvertFp16() Mat {
	return newMat(C.Mat_ConvertFp16(m.p))
}

// Mean calculates the mean value M of array elements, independently for each channel, and return it as Scalar
// TODO pass second paramter with mask
func (m *Mat) Mean() Scalar {
	s := C.Mat_Mean(m.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// Sqrt calculates a square root of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga186222c3919657890f88df5a1f64a7d7
//
func (m *Mat) Sqrt() Mat {
	return newMat(C.Mat_Sqrt(m.p))
}

// Sum calculates the per-channel pixel sum of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga716e10a2dd9e228e4d3c95818f106722
//
func (m *Mat) Sum() Scalar {
	s := C.Mat_Sum(m.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// PatchNaNs converts NaN's to zeros.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga62286befb7cde3568ff8c7d14d5079da
//
func (m *Mat) PatchNaNs() {
	C.Mat_PatchNaNs(m.p)
}

// LUT performs a look-up table transform of an array.
//
// The function LUT fills the output array with values from the look-up table.
// Indices of the entries are taken from the input array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab55b8d062b7f5587720ede032d34156f
func LUT(src, wbLUT Mat, dst *Mat) {
	C.LUT(src.p, wbLUT.p, dst.p)
}

// Rows returns the number of rows for this Mat.
func (m *Mat) Rows() int {
	return int(C.Mat_Rows(m.p))
}

// Cols returns the number of columns for this Mat.
func (m *Mat) Cols() int {
	return int(C.Mat_Cols(m.p))
}

// Channels returns the number of channels for this Mat.
func (m *Mat) Channels() int {
	return int(C.Mat_Channels(m.p))
}

// Type returns the type for this Mat.
func (m *Mat) Type() MatType {
	return MatType(C.Mat_Type(m.p))
}

// Step returns the number of bytes each matrix row occupies.
func (m *Mat) Step() int {
	return int(C.Mat_Step(m.p))
}

// GetUCharAt returns a value from a specific row/col
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt(row int, col int) uint8 {
	return uint8(C.Mat_GetUChar(m.p, C.int(row), C.int(col)))
}

// GetUCharAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt3(x, y, z int) uint8 {
	return uint8(C.Mat_GetUChar3(m.p, C.int(x), C.int(y), C.int(z)))
}

// GetSCharAt returns a value from a specific row/col
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) GetSCharAt(row int, col int) int8 {
	return int8(C.Mat_GetSChar(m.p, C.int(row), C.int(col)))
}

// GetSCharAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) GetSCharAt3(x, y, z int) int8 {
	return int8(C.Mat_GetSChar3(m.p, C.int(x), C.int(y), C.int(z)))
}

// GetShortAt returns a value from a specific row/col
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) GetShortAt(row int, col int) int16 {
	return int16(C.Mat_GetShort(m.p, C.int(row), C.int(col)))
}

// GetShortAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) GetShortAt3(x, y, z int) int16 {
	return int16(C.Mat_GetShort3(m.p, C.int(x), C.int(y), C.int(z)))
}

// GetIntAt returns a value from a specific row/col
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) GetIntAt(row int, col int) int32 {
	return int32(C.Mat_GetInt(m.p, C.int(row), C.int(col)))
}

// GetIntAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) GetIntAt3(x, y, z int) int32 {
	return int32(C.Mat_GetInt3(m.p, C.int(x), C.int(y), C.int(z)))
}

// GetFloatAt returns a value from a specific row/col
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) GetFloatAt(row int, col int) float32 {
	return float32(C.Mat_GetFloat(m.p, C.int(row), C.int(col)))
}

// GetFloatAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) GetFloatAt3(x, y, z int) float32 {
	return float32(C.Mat_GetFloat3(m.p, C.int(x), C.int(y), C.int(z)))
}

// GetDoubleAt returns a value from a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt(row int, col int) float64 {
	return float64(C.Mat_GetDouble(m.p, C.int(row), C.int(col)))
}

// GetDoubleAt3 returns a value from a specific x, y, z coordinate location
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt3(x, y, z int) float64 {
	return float64(C.Mat_GetDouble3(m.p, C.int(x), C.int(y), C.int(z)))
}

// SetTo sets all or some of the array elements to the specified scalar value.
func (m *Mat) SetTo(s Scalar) {
	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	C.Mat_SetTo(m.p, sVal)
}

// SetUCharAt sets a value at a specific row/col
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) SetUCharAt(row int, col int, val uint8) {
	C.Mat_SetUChar(m.p, C.int(row), C.int(col), C.uint8_t(val))
}

// SetUCharAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) SetUCharAt3(x, y, z int, val uint8) {
	C.Mat_SetUChar3(m.p, C.int(x), C.int(y), C.int(z), C.uint8_t(val))
}

// SetSCharAt sets a value at a specific row/col
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) SetSCharAt(row int, col int, val int8) {
	C.Mat_SetSChar(m.p, C.int(row), C.int(col), C.int8_t(val))
}

// SetSCharAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) SetSCharAt3(x, y, z int, val int8) {
	C.Mat_SetSChar3(m.p, C.int(x), C.int(y), C.int(z), C.int8_t(val))
}

// SetShortAt sets a value at a specific row/col
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) SetShortAt(row int, col int, val int16) {
	C.Mat_SetShort(m.p, C.int(row), C.int(col), C.int16_t(val))
}

// SetShortAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) SetShortAt3(x, y, z int, val int16) {
	C.Mat_SetShort3(m.p, C.int(x), C.int(y), C.int(z), C.int16_t(val))
}

// SetIntAt sets a value at a specific row/col
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) SetIntAt(row int, col int, val int32) {
	C.Mat_SetInt(m.p, C.int(row), C.int(col), C.int32_t(val))
}

// SetIntAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) SetIntAt3(x, y, z int, val int32) {
	C.Mat_SetInt3(m.p, C.int(x), C.int(y), C.int(z), C.int32_t(val))
}

// SetFloatAt sets a value at a specific row/col
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) SetFloatAt(row int, col int, val float32) {
	C.Mat_SetFloat(m.p, C.int(row), C.int(col), C.float(val))
}

// SetFloatAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) SetFloatAt3(x, y, z int, val float32) {
	C.Mat_SetFloat3(m.p, C.int(x), C.int(y), C.int(z), C.float(val))
}

// SetDoubleAt sets a value at a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt(row int, col int, val float64) {
	C.Mat_SetDouble(m.p, C.int(row), C.int(col), C.double(val))
}

// SetDoubleAt3 sets a value at a specific x, y, z coordinate location
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt3(x, y, z int, val float64) {
	C.Mat_SetDouble3(m.p, C.int(x), C.int(y), C.int(z), C.double(val))
}

// AddUChar adds a uchar value to each element in the Mat. Performs a
// mat += val operation.
func (m *Mat) AddUChar(val uint8) {
	C.Mat_AddUChar(m.p, C.uint8_t(val))
}

// SubtractUChar subtracts a uchar value from each element in the Mat. Performs a
// mat -= val operation.
func (m *Mat) SubtractUChar(val uint8) {
	C.Mat_SubtractUChar(m.p, C.uint8_t(val))
}

// MultiplyUChar multiplies each element in the Mat by a uint value. Performs a
// mat *= val operation.
func (m *Mat) MultiplyUChar(val uint8) {
	C.Mat_MultiplyUChar(m.p, C.uint8_t(val))
}

// DivideUChar divides each element in the Mat by a uint value. Performs a
// mat /= val operation.
func (m *Mat) DivideUChar(val uint8) {
	C.Mat_DivideUChar(m.p, C.uint8_t(val))
}

// AddFloat adds a float value to each element in the Mat. Performs a
// mat += val operation.
func (m *Mat) AddFloat(val float32) {
	C.Mat_AddFloat(m.p, C.float(val))
}

// SubtractFloat subtracts a float value from each element in the Mat. Performs a
// mat -= val operation.
func (m *Mat) SubtractFloat(val float32) {
	C.Mat_SubtractFloat(m.p, C.float(val))
}

// MultiplyFloat multiplies each element in the Mat by a float value. Performs a
// mat *= val operation.
func (m *Mat) MultiplyFloat(val float32) {
	C.Mat_MultiplyFloat(m.p, C.float(val))
}

// DivideFloat divides each element in the Mat by a float value. Performs a
// mat /= val operation.
func (m *Mat) DivideFloat(val float32) {
	C.Mat_DivideFloat(m.p, C.float(val))
}

// ToImage converts a Mat to a image.Image.
func (m *Mat) ToImage() (image.Image, error) {
	t := m.Type()
	if t != MatTypeCV8UC1 && t != MatTypeCV8UC3 && t != MatTypeCV8UC4 {
		return nil, errors.New("ToImage supports only MatType CV8UC1, CV8UC3 and CV8UC4")
	}

	width := m.Cols()
	height := m.Rows()
	step := m.Step()
	data := m.ToBytes()
	channels := m.Channels()

	if t == MatTypeCV8UC1 {
		img := image.NewGray(image.Rect(0, 0, width, height))
		c := color.Gray{Y: uint8(0)}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				c.Y = uint8(data[y*step+x])
				img.SetGray(x, y, c)
			}
		}

		return img, nil
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	c := color.RGBA{
		R: uint8(0),
		G: uint8(0),
		B: uint8(0),
		A: uint8(255),
	}

	for y := 0; y < height; y++ {
		for x := 0; x < step; x = x + channels {
			c.B = uint8(data[y*step+x])
			c.G = uint8(data[y*step+x+1])
			c.R = uint8(data[y*step+x+2])
			if channels == 4 {
				c.A = uint8(data[y*step+x+3])
			}
			img.SetRGBA(int(x/channels), y, c)
		}
	}

	return img, nil
}

//ImageToMatRGBA converts image.Image to gocv.Mat,
//which represents RGBA image having 8bit for each component.
//Type of Mat is gocv.MatTypeCV8UC4.
func ImageToMatRGBA(img image.Image) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	data := make([]byte, 0, x*y*4)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, a := img.At(i, j).RGBA()
			data = append(data, byte(b>>8), byte(g>>8), byte(r>>8), byte(a>>8))
		}
	}
	return NewMatFromBytes(y, x, MatTypeCV8UC4, data)
}

//ImageToMatRGB converts image.Image to gocv.Mat,
//which represents RGB image having 8bit for each component.
//Type of Mat is gocv.MatTypeCV8UC3.
func ImageToMatRGB(img image.Image) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	data := make([]byte, 0, x*y*3)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			data = append(data, byte(b>>8), byte(g>>8), byte(r>>8))
		}
	}
	return NewMatFromBytes(y, x, MatTypeCV8UC3, data)
}

//ImageGrayToMatGray converts image.Gray to gocv.Mat,
//which represents grayscale image 8bit.
//Type of Mat is gocv.MatTypeCV8UC1.
func ImageGrayToMatGray(img *image.Gray) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	data := make([]byte, 0, x*y)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			data = append(data, img.GrayAt(i, j).Y)
		}
	}
	return NewMatFromBytes(y, x, MatTypeCV8UC1, data)
}

// AbsDiff calculates the per-element absolute difference between two arrays
// or between an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6fef31bc8c4071cbc114a758a2b79c14
//
func AbsDiff(src1, src2 Mat, dst *Mat) {
	C.Mat_AbsDiff(src1.p, src2.p, dst.p)
}

// Add calculates the per-element sum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga10ac1bfb180e2cfda1701d06c24fdbd6
//
func Add(src1, src2 Mat, dst *Mat) {
	C.Mat_Add(src1.p, src2.p, dst.p)
}

// AddWeighted calculates the weighted sum of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gafafb2513349db3bcff51f54ee5592a19
//
func AddWeighted(src1 Mat, alpha float64, src2 Mat, beta float64, gamma float64, dst *Mat) {
	C.Mat_AddWeighted(src1.p, C.double(alpha),
		src2.p, C.double(beta), C.double(gamma), dst.p)
}

// BitwiseAnd computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
//
func BitwiseAnd(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_BitwiseAnd(src1.p, src2.p, dst.p)
}

// BitwiseAndWithMask computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
//
func BitwiseAndWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	C.Mat_BitwiseAndWithMask(src1.p, src2.p, dst.p, mask.p)
}

// BitwiseNot inverts every bit of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
//
func BitwiseNot(src1 Mat, dst *Mat) {
	C.Mat_BitwiseNot(src1.p, dst.p)
}

// BitwiseNotWithMask inverts every bit of an array. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
//
func BitwiseNotWithMask(src1 Mat, dst *Mat, mask Mat) {
	C.Mat_BitwiseNotWithMask(src1.p, dst.p, mask.p)
}

// BitwiseOr calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
//
func BitwiseOr(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_BitwiseOr(src1.p, src2.p, dst.p)
}

// BitwiseOrWithMask calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
//
func BitwiseOrWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	C.Mat_BitwiseOrWithMask(src1.p, src2.p, dst.p, mask.p)
}

// BitwiseXor calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
//
func BitwiseXor(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_BitwiseXor(src1.p, src2.p, dst.p)
}

// BitwiseXorWithMask calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar. It has an additional parameter for a mask.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
//
func BitwiseXorWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
	C.Mat_BitwiseXorWithMask(src1.p, src2.p, dst.p, mask.p)
}

// BatchDistance is a naive nearest neighbor finder.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4ba778a1c57f83233b1d851c83f5a622
//
func BatchDistance(src1 Mat, src2 Mat, dist Mat, dtype int, nidx Mat, normType int, K int, mask Mat, update int, crosscheck bool) {
	C.Mat_BatchDistance(src1.p, src2.p, dist.p, C.int(dtype), nidx.p, C.int(normType), C.int(K), mask.p, C.int(update), C.bool(crosscheck))
}

// BorderInterpolate computes the source location of an extrapolated pixel.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga247f571aa6244827d3d798f13892da58
//
func BorderInterpolate(p int, len int, borderType CovarFlags) int {
	ret := C.Mat_BorderInterpolate(C.int(p), C.int(len), C.int(borderType))
	return int(ret)
}

// CovarFlags are the covariation flags used by functions such as BorderInterpolate.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/de1/group__core.html#ga719ebd4a73f30f4fab258ab7616d0f0f
//
type CovarFlags int

const (
	// CovarScrambled indicates to scramble the results.
	CovarScrambled CovarFlags = 0

	// CovarNormal indicates to use normal covariation.
	CovarNormal = 1

	// CovarUseAvg indicates to use average covariation.
	CovarUseAvg = 2

	// CovarScale indicates to use scaled covariation.
	CovarScale = 4

	// CovarRows indicates to use covariation on rows.
	CovarRows = 8

	// CovarCols indicates to use covariation on columns.
	CovarCols = 16
)

// CalcCovarMatrix calculates the covariance matrix of a set of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga017122d912af19d7d0d2cccc2d63819f
//
func CalcCovarMatrix(samples Mat, covar *Mat, mean *Mat, flags CovarFlags, ctype int) {
	C.Mat_CalcCovarMatrix(samples.p, covar.p, mean.p, C.int(flags), C.int(ctype))
}

// CartToPolar calculates the magnitude and angle of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gac5f92f48ec32cacf5275969c33ee837d
//
func CartToPolar(x Mat, y Mat, magnitude *Mat, angle *Mat, angleInDegrees bool) {
	C.Mat_CartToPolar(x.p, y.p, magnitude.p, angle.p, C.bool(angleInDegrees))
}

// CheckRange checks every element of an input array for invalid values.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga2bd19d89cae59361416736f87e3c7a64
//
func CheckRange(src Mat) bool {
	return bool(C.Mat_CheckRange(src.p))
}

// Compare performs the per-element comparison of two arrays
// or an array and scalar value.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga303cfb72acf8cbb36d884650c09a3a97
//
func Compare(src1 Mat, src2 Mat, dst *Mat, ct CompareType) {
	C.Mat_Compare(src1.p, src2.p, dst.p, C.int(ct))
}

// CountNonZero counts non-zero array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa4b89393263bb4d604e0fe5986723914
//
func CountNonZero(src Mat) int {
	return int(C.Mat_CountNonZero(src.p))
}

// CompleteSymm copies the lower or the upper half of a square matrix to its another half.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa9d88dcd0e54b6d1af38d41f2a3e3d25
//
func CompleteSymm(m Mat, lowerToUpper bool) {
	C.Mat_CompleteSymm(m.p, C.bool(lowerToUpper))
}

// ConvertScaleAbs scales, calculates absolute values, and converts the result to 8-bit.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3460e9c9f37b563ab9dd550c4d8c4e7d
//
func ConvertScaleAbs(src Mat, dst *Mat, alpha float64, beta float64) {
	C.Mat_ConvertScaleAbs(src.p, dst.p, C.double(alpha), C.double(beta))
}

// CopyMakeBorder forms a border around an image (applies padding).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga2ac1049c2c3dd25c2b41bffe17658a36
//
func CopyMakeBorder(src Mat, dst *Mat, top int, bottom int, left int, right int, bt BorderType, value color.RGBA) {

	cValue := C.struct_Scalar{
		val1: C.double(value.B),
		val2: C.double(value.G),
		val3: C.double(value.R),
		val4: C.double(value.A),
	}

	C.Mat_CopyMakeBorder(src.p, dst.p, C.int(top), C.int(bottom), C.int(left), C.int(right), C.int(bt), cValue)
}

// DftFlags represents a DFT or DCT flag.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf4dde112b483b38175621befedda1f1c
//
type DftFlags int

const (
	// DftForward performs forward 1D or 2D dft or dct.
	DftForward DftFlags = 0

	// DftInverse performs an inverse 1D or 2D transform.
	DftInverse = 1

	// DftScale scales the result: divide it by the number of array elements. Normally, it is combined with DFT_INVERSE.
	DftScale = 2

	// DftRows performs a forward or inverse transform of every individual row of the input matrix.
	DftRows = 4

	// DftComplexOutput performs a forward transformation of 1D or 2D real array; the result, though being a complex array, has complex-conjugate symmetry
	DftComplexOutput = 16

	// DftRealOutput performs an inverse transformation of a 1D or 2D complex array; the result is normally a complex array of the same size,
	// however, if the input array has conjugate-complex symmetry (for example, it is a result of forward transformation with DFT_COMPLEX_OUTPUT flag),
	// the output is a real array.
	DftRealOutput = 32

	// DftComplexInput specifies that input is complex input. If this flag is set, the input must have 2 channels.
	DftComplexInput = 64

	// DctInverse performs an inverse 1D or 2D dct transform.
	DctInverse = DftInverse

	// DctRows performs a forward or inverse dct transform of every individual row of the input matrix.
	DctRows = DftRows
)

// DCT performs a forward or inverse discrete Cosine transform of 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga85aad4d668c01fbd64825f589e3696d4
//
func DCT(src Mat, dst *Mat, flags DftFlags) {
	C.Mat_DCT(src.p, dst.p, C.int(flags))
}

// Determinant returns the determinant of a square floating-point matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf802bd9ca3e07b8b6170645ef0611d0c
//
func Determinant(src Mat) float64 {
	return float64(C.Mat_Determinant(src.p))
}

// DFT performs a forward or inverse Discrete Fourier Transform (DFT)
// of a 1D or 2D floating-point array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gadd6cf9baf2b8b704a11b5f04aaf4f39d
//
func DFT(src Mat, dst *Mat, flags DftFlags) {
	C.Mat_DFT(src.p, dst.p, C.int(flags))
}

// Divide performs the per-element division
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6db555d30115642fedae0cda05604874
//
func Divide(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_Divide(src1.p, src2.p, dst.p)
}

// Eigen calculates eigenvalues and eigenvectors of a symmetric matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9fa0d58657f60eaa6c71f6fbb40456e3
//
func Eigen(src Mat, eigenvalues *Mat, eigenvectors *Mat) bool {
	ret := C.Mat_Eigen(src.p, eigenvalues.p, eigenvectors.p)
	return bool(ret)
}

// EigenNonSymmetric calculates eigenvalues and eigenvectors of a non-symmetric matrix (real eigenvalues only).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf51987e03cac8d171fbd2b327cf966f6
//
func EigenNonSymmetric(src Mat, eigenvalues *Mat, eigenvectors *Mat) {
	C.Mat_EigenNonSymmetric(src.p, eigenvalues.p, eigenvectors.p)
}

// Exp calculates the exponent of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3e10108e2162c338f1b848af619f39e5
//
func Exp(src Mat, dst *Mat) {
	C.Mat_Exp(src.p, dst.p)
}

// ExtractChannel extracts a single channel from src (coi is 0-based index).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacc6158574aa1f0281878c955bcf35642
//
func ExtractChannel(src Mat, dst *Mat, coi int) {
	C.Mat_ExtractChannel(src.p, dst.p, C.int(coi))
}

// FindNonZero returns the list of locations of non-zero pixels.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaed7df59a3539b4cc0fe5c9c8d7586190
//
func FindNonZero(src Mat, idx *Mat) {
	C.Mat_FindNonZero(src.p, idx.p)
}

// Flip flips a 2D array around horizontal(0), vertical(1), or both axes(-1).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaca7be533e3dac7feb70fc60635adf441
//
func Flip(src Mat, dst *Mat, flipCode int) {
	C.Mat_Flip(src.p, dst.p, C.int(flipCode))
}

// Gemm performs generalized matrix multiplication.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacb6e64071dffe36434e1e7ee79e7cb35
//
func Gemm(src1, src2 Mat, alpha float64, src3 Mat, beta float64, dst *Mat, flags int) {
	C.Mat_Gemm(src1.p, src2.p, C.double(alpha), src3.p, C.double(beta), dst.p, C.int(flags))
}

// GetOptimalDFTSize returns the optimal Discrete Fourier Transform (DFT) size
// for a given vector size.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6577a2e59968936ae02eb2edde5de299
//
func GetOptimalDFTSize(vecsize int) int {
	return int(C.Mat_GetOptimalDFTSize(C.int(vecsize)))
}

// Hconcat applies horizontal concatenation to given matrices.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaab5ceee39e0580f879df645a872c6bf7
//
func Hconcat(src1, src2 Mat, dst *Mat) {
	C.Mat_Hconcat(src1.p, src2.p, dst.p)
}

// Vconcat applies vertical concatenation to given matrices.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaab5ceee39e0580f879df645a872c6bf7
//
func Vconcat(src1, src2 Mat, dst *Mat) {
	C.Mat_Vconcat(src1.p, src2.p, dst.p)
}

// RotateFlag for image rotation
//
//
// For further details please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6f45d55c0b1cc9d97f5353a7c8a7aac2
type RotateFlag int

const (
	// Rotate90Clockwise allows to rotate image 90 degrees clockwise
	Rotate90Clockwise RotateFlag = 0
	// Rotate180Clockwise allows to rotate image 180 degrees clockwise
	Rotate180Clockwise = 1
	// Rotate90CounterClockwise allows to rotate 270 degrees clockwise
	Rotate90CounterClockwise = 2
)

// Rotate rotates a 2D array in multiples of 90 degrees
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4ad01c0978b0ce64baa246811deeac24
func Rotate(src Mat, dst *Mat, code RotateFlag) {
	C.Rotate(src.p, dst.p, C.int(code))
}

// IDCT calculates the inverse Discrete Cosine Transform of a 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga77b168d84e564c50228b69730a227ef2
//
func IDCT(src Mat, dst *Mat, flags int) {
	C.Mat_Idct(src.p, dst.p, C.int(flags))
}

// IDFT calculates the inverse Discrete Fourier Transform of a 1D or 2D array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa708aa2d2e57a508f968eb0f69aa5ff1
//
func IDFT(src Mat, dst *Mat, flags, nonzeroRows int) {
	C.Mat_Idft(src.p, dst.p, C.int(flags), C.int(nonzeroRows))
}

// InRange checks if array elements lie between the elements of two Mat arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
//
func InRange(src, lb, ub Mat, dst *Mat) {
	C.Mat_InRange(src.p, lb.p, ub.p, dst.p)
}

// InRangeWithScalar checks if array elements lie between the elements of two Scalars
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
//
func InRangeWithScalar(src Mat, lb, ub Scalar, dst *Mat) {
	lbVal := C.struct_Scalar{
		val1: C.double(lb.Val1),
		val2: C.double(lb.Val2),
		val3: C.double(lb.Val3),
		val4: C.double(lb.Val4),
	}

	ubVal := C.struct_Scalar{
		val1: C.double(ub.Val1),
		val2: C.double(ub.Val2),
		val3: C.double(ub.Val3),
		val4: C.double(ub.Val4),
	}

	C.Mat_InRangeWithScalar(src.p, lbVal, ubVal, dst.p)
}

// InsertChannel inserts a single channel to dst (coi is 0-based index)
// (it replaces channel i with another in dst).
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1d4bd886d35b00ec0b764cb4ce6eb515
//
func InsertChannel(src Mat, dst *Mat, coi int) {
	C.Mat_InsertChannel(src.p, dst.p, C.int(coi))
}

// Invert finds the inverse or pseudo-inverse of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad278044679d4ecf20f7622cc151aaaa2
//
func Invert(src Mat, dst *Mat, flags int) float64 {
	ret := C.Mat_Invert(src.p, dst.p, C.int(flags))
	return float64(ret)
}

// Log calculates the natural logarithm of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga937ecdce4679a77168730830a955bea7
//
func Log(src Mat, dst *Mat) {
	C.Mat_Log(src.p, dst.p)
}

// Magnitude calculates the magnitude of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6d3b097586bca4409873d64a90fe64c3
//
func Magnitude(x, y Mat, magnitude *Mat) {
	C.Mat_Magnitude(x.p, y.p, magnitude.p)
}

// Max calculates per-element maximum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gacc40fa15eac0fb83f8ca70b7cc0b588d
//
func Max(src1, src2 Mat, dst *Mat) {
	C.Mat_Max(src1.p, src2.p, dst.p)
}

// MeanStdDev calculates a mean and standard deviation of array elements.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga846c858f4004d59493d7c6a4354b301d
//
func MeanStdDev(src Mat, dst *Mat, dstStdDev *Mat) {
	C.Mat_MeanStdDev(src.p, dst.p, dstStdDev.p)
}

// Merge creates one multi-channel array out of several single-channel ones.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7d7b4d6c6ee504b30a20b1680029c7b4
//
func Merge(mv []Mat, dst *Mat) {
	cMatArray := make([]C.Mat, len(mv))
	for i, r := range mv {
		cMatArray[i] = r.p
	}
	cMats := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(mv)),
	}

	C.Mat_Merge(cMats, dst.p)
}

// Min calculates per-element minimum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9af368f182ee76d0463d0d8d5330b764
//
func Min(src1, src2 Mat, dst *Mat) {
	C.Mat_Min(src1.p, src2.p, dst.p)
}

// MinMaxIdx finds the global minimum and maximum in an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7622c466c628a75d9ed008b42250a73f
//
func MinMaxIdx(input Mat) (minVal, maxVal float32, minIdx, maxIdx int) {
	var cMinVal C.double
	var cMaxVal C.double
	var cMinIdx C.int
	var cMaxIdx C.int

	C.Mat_MinMaxIdx(input.p, &cMinVal, &cMaxVal, &cMinIdx, &cMaxIdx)

	return float32(cMinVal), float32(cMaxVal), int(minIdx), int(maxIdx)
}

// MinMaxLoc finds the global minimum and maximum in an array.
//
// For further details, please see:
// https://docs.opencv.org/trunk/d2/de8/group__core__array.html#gab473bf2eb6d14ff97e89b355dac20707
//
func MinMaxLoc(input Mat) (minVal, maxVal float32, minLoc, maxLoc image.Point) {
	var cMinVal C.double
	var cMaxVal C.double
	var cMinLoc C.struct_Point
	var cMaxLoc C.struct_Point

	C.Mat_MinMaxLoc(input.p, &cMinVal, &cMaxVal, &cMinLoc, &cMaxLoc)

	minLoc = image.Pt(int(cMinLoc.x), int(cMinLoc.y))
	maxLoc = image.Pt(int(cMaxLoc.x), int(cMaxLoc.y))

	return float32(cMinVal), float32(cMaxVal), minLoc, maxLoc
}

//Mulspectrums performs the per-element multiplication of two Fourier spectrums.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3ab38646463c59bf0ce962a9d51db64f
//
func MulSpectrums(a Mat, b Mat, dst *Mat, flags DftFlags) {
	C.Mat_MulSpectrums(a.p, b.p, dst.p, C.int(flags))
}

// Multiply calculates the per-element scaled product of two arrays.
// Both input arrays must be of the same size and the same type.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga979d898a58d7f61c53003e162e7ad89f
//
func Multiply(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_Multiply(src1.p, src2.p, dst.p)
}

// NormType for normalization operations.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad12cefbcb5291cf958a85b4b67b6149f
//
type NormType int

const (
	// NormInf indicates use infinite normalization.
	NormInf NormType = 1

	// NormL1 indicates use L1 normalization.
	NormL1 = 2

	// NormL2 indicates use L2 normalization.
	NormL2 = 4

	// NormL2Sqr indicates use L2 squared normalization.
	NormL2Sqr = 5

	// NormHamming indicates use Hamming normalization.
	NormHamming = 6

	// NormHamming2 indicates use Hamming 2-bit normalization.
	NormHamming2 = 7

	// NormTypeMask indicates use type mask for normalization.
	NormTypeMask = 7

	// NormRelative indicates use relative normalization.
	NormRelative = 8

	// NormMinMax indicates use min/max normalization.
	NormMinMax = 32
)

// Normalize normalizes the norm or value range of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga87eef7ee3970f86906d69a92cbf064bd
//
func Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType) {
	C.Mat_Normalize(src.p, dst.p, C.double(alpha), C.double(beta), C.int(typ))
}

// Norm calculates the absolute norm of an array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33
//
func Norm(src1 Mat, normType NormType) float64 {
	return float64(C.Norm(src1.p, C.int(normType)))
}

// PerspectiveTransform performs the perspective matrix transformation of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad327659ac03e5fd6894b90025e6900a7
//
func PerspectiveTransform(src Mat, dst *Mat, tm Mat) {
	C.Mat_PerspectiveTransform(src.p, dst.p, tm.p)
}

// TermCriteriaType for TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d5d/classcv_1_1TermCriteria.html#a56fecdc291ccaba8aad27d67ccf72c57
//
type TermCriteriaType int

const (
	// Count is the maximum number of iterations or elements to compute.
	Count TermCriteriaType = 1

	// MaxIter is the maximum number of iterations or elements to compute.
	MaxIter = 1

	// EPS is the desired accuracy or change in parameters at which the
	// iterative algorithm stops.
	EPS = 2
)

type SolveDecompositionFlags int

const (
	// Gaussian elimination with the optimal pivot element chosen.
	SolveDecompositionLu = 0

	// Singular value decomposition (SVD) method. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionSvd = 1

	// Eigenvalue decomposition. The matrix src1 must be symmetrical.
	SolveDecompositionEing = 2

	// Cholesky LL^T factorization. The matrix src1 must be symmetrical and positively defined.
	SolveDecompositionCholesky = 3

	// QR factorization. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionQr = 4

	// While all the previous flags are mutually exclusive, this flag can be used together with any of the previous.
	// It means that the normal equations 𝚜𝚛𝚌𝟷^T⋅𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟷^T𝚜𝚛𝚌𝟸 are solved instead of the original system
	// 𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟸.
	SolveDecompositionNormal = 5
)

// Solve solves one or more linear systems or least-squares problems.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga12b43690dbd31fed96f213eefead2373
//
func Solve(src1 Mat, src2 Mat, dst *Mat, flags SolveDecompositionFlags) bool {
	return bool(C.Mat_Solve(src1.p, src2.p, dst.p, C.int(flags)))
}

// SolveCubic finds the real roots of a cubic equation.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1c3b0b925b085b6e96931ee309e6a1da
//
func SolveCubic(coeffs Mat, roots *Mat) int {
	return int(C.Mat_SolveCubic(coeffs.p, roots.p))
}

// SolvePoly finds the real or complex roots of a polynomial equation.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gac2f5e953016fabcdf793d762f4ec5dce
//
func SolvePoly(coeffs Mat, roots *Mat, maxIters int) float64 {
	return float64(C.Mat_SolvePoly(coeffs.p, roots.p, C.int(maxIters)))
}

type ReduceTypes int

const (
	// The output is the sum of all rows/columns of the matrix.
	ReduceSum ReduceTypes = 0

	// The output is the mean vector of all rows/columns of the matrix.
	ReduceAvg ReduceTypes = 1

	// The output is the maximum (column/row-wise) of all rows/columns of the matrix.
	ReduceMax ReduceTypes = 2

	// The output is the minimum (column/row-wise) of all rows/columns of the matrix.
	ReduceMin ReduceTypes = 3
)

// Reduce reduces a matrix to a vector.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4b78072a303f29d9031d56e5638da78e
//
func Reduce(src Mat, dst *Mat, dim int, rType ReduceTypes, dType int) {
	C.Mat_Reduce(src.p, dst.p, C.int(dim), C.int(rType), C.int(dType))
}

// Repeat fills the output array with repeated copies of the input array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga496c3860f3ac44c40b48811333cfda2d
//
func Repeat(src Mat, nY int, nX int, dst *Mat) {
	C.Mat_Repeat(src.p, C.int(nY), C.int(nX), dst.p)
}

// Calculates the sum of a scaled array and another array.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9e0845db4135f55dcf20227402f00d98
//
func ScaleAdd(src1 Mat, alpha float64, src2 Mat, dst *Mat) {
	C.Mat_ScaleAdd(src1.p, C.double(alpha), src2.p, dst.p)
}

type SortFlags int

const (
	// Each matrix row is sorted independently
	SortEveryRow SortFlags = 0

	// Each matrix column is sorted independently; this flag and the previous one are mutually exclusive.
	SortEveryColumn SortFlags = 1

	// Each matrix row is sorted in the ascending order.
	SortAscending SortFlags = 0

	// Each matrix row is sorted in the descending order; this flag and the previous one are also mutually exclusive.
	SortDescending SortFlags = 16
)

// Sort sorts each row or each column of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga45dd56da289494ce874be2324856898f
//
func Sort(src Mat, dst *Mat, flags SortFlags) {
	C.Mat_Sort(src.p, dst.p, C.int(flags))
}

// SortIdx sorts each row or each column of a matrix.
// Instead of reordering the elements themselves, it stores the indices of sorted elements in the output array
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gadf35157cbf97f3cb85a545380e383506
//
func SortIdx(src Mat, dst *Mat, flags SortFlags) {
	C.Mat_SortIdx(src.p, dst.p, C.int(flags))
}

// Split creates an array of single channel images from a multi-channel image
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0547c7fed86152d7e9d0096029c8518a
//
func Split(src Mat) (mv []Mat) {
	cMats := C.struct_Mats{}
	C.Mat_Split(src.p, &(cMats))
	mv = make([]Mat, cMats.length)
	for i := C.int(0); i < cMats.length; i++ {
		mv[i].p = C.Mats_get(cMats, i)
	}
	return
}

// Subtract calculates the per-element subtraction of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaa0f00d98b4b5edeaeb7b8333b2de353b
//
func Subtract(src1 Mat, src2 Mat, dst *Mat) {
	C.Mat_Subtract(src1.p, src2.p, dst.p)
}

// Trace returns the trace of a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3419ac19c7dcd2be4bd552a23e147dd8
//
func Trace(src Mat) Scalar {
	s := C.Mat_Trace(src.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// Transform performs the matrix transformation of every array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga393164aa54bb9169ce0a8cc44e08ff22
//
func Transform(src Mat, dst *Mat, tm Mat) {
	C.Mat_Transform(src.p, dst.p, tm.p)
}

// Transpose transposes a matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga46630ed6c0ea6254a35f447289bd7404
//
func Transpose(src Mat, dst *Mat) {
	C.Mat_Transpose(src.p, dst.p)
}

// Pow raises every array element to a power.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf0d056b5bd1dc92500d6f6cf6bac41ef
//
func Pow(src Mat, power float64, dst *Mat) {
	C.Mat_Pow(src.p, C.double(power), dst.p)
}

// PolatToCart calculates x and y coordinates of 2D vectors from their magnitude and angle.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga581ff9d44201de2dd1b40a50db93d665
//
func PolarToCart(magnitude Mat, degree Mat, x *Mat, y *Mat, angleInDegrees bool) {
	C.Mat_PolarToCart(magnitude.p, degree.p, x.p, y.p, C.bool(angleInDegrees))
}

// Phase calculates the rotation angle of 2D vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9db9ca9b4d81c3bde5677b8f64dc0137
//
func Phase(x, y Mat, angle *Mat, angleInDegrees bool) {
	C.Mat_Phase(x.p, y.p, angle.p, C.bool(angleInDegrees))
}

// TermCriteria is the criteria for iterative algorithms.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d5d/classcv_1_1TermCriteria.html
//
type TermCriteria struct {
	p C.TermCriteria
}

// NewTermCriteria returns a new TermCriteria.
func NewTermCriteria(typ TermCriteriaType, maxCount int, epsilon float64) TermCriteria {
	return TermCriteria{p: C.TermCriteria_New(C.int(typ), C.int(maxCount), C.double(epsilon))}
}

// Scalar is a 4-element vector widely used in OpenCV to pass pixel values.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/da0/classcv_1_1Scalar__.html
//
type Scalar struct {
	Val1 float64
	Val2 float64
	Val3 float64
	Val4 float64
}

// NewScalar returns a new Scalar. These are usually colors typically being in BGR order.
func NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
	s := Scalar{Val1: v1, Val2: v2, Val3: v3, Val4: v4}
	return s
}

// KeyPoint is data structure for salient point detectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/d29/classcv_1_1KeyPoint.html
//
type KeyPoint struct {
	X, Y                  float64
	Size, Angle, Response float64
	Octave, ClassID       int
}

// DMatch is data structure for matching keypoint descriptors.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/de0/classcv_1_1DMatch.html#a546ddb9a87898f06e510e015a6de596e
//
type DMatch struct {
	QueryIdx int
	TrainIdx int
	ImgIdx   int
	Distance float64
}

// Vecf is a generic vector of floats.
type Vecf []float32

// GetVecfAt returns a vector of floats. Its size corresponds to the number of
// channels of the Mat.
func (m *Mat) GetVecfAt(row int, col int) Vecf {
	ch := m.Channels()
	v := make(Vecf, ch)

	for c := 0; c < ch; c++ {
		v[c] = m.GetFloatAt(row, col*ch+c)
	}

	return v
}

// Veci is a generic vector of integers.
type Veci []int32

// GetVeciAt returns a vector of integers. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVeciAt(row int, col int) Veci {
	ch := m.Channels()
	v := make(Veci, ch)

	for c := 0; c < ch; c++ {
		v[c] = m.GetIntAt(row, col*ch+c)
	}

	return v
}

// GetTickCount returns the number of ticks.
//
// For further details, please see:
// https://docs.opencv.org/master/db/de0/group__core__utils.html#gae73f58000611a1af25dd36d496bf4487
//
func GetTickCount() float64 {
	return float64(C.GetCVTickCount())
}

// GetTickFrequency returns the number of ticks per second.
//
// For further details, please see:
// https://docs.opencv.org/master/db/de0/group__core__utils.html#ga705441a9ef01f47acdc55d87fbe5090c
//
func GetTickFrequency() float64 {
	return float64(C.GetTickFrequency())
}

func toByteArray(b []byte) (*C.struct_ByteArray, error) {
	if len(b) == 0 {
		return nil, ErrEmptyByteSlice
	}
	return &C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}, nil
}

func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}

func toRectangles(ret C.Rects) []image.Rectangle {
	cArray := ret.rects
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.Rect)(unsafe.Pointer(&hdr))

	rects := make([]image.Rectangle, length)
	for i, r := range s {
		rects[i] = image.Rect(int(r.x), int(r.y), int(r.x+r.width), int(r.y+r.height))
	}
	return rects
}

func toRect(rect C.Rect) image.Rectangle {
	return image.Rect(int(rect.x), int(rect.y), int(rect.x+rect.width), int(rect.y+rect.height))
}

func toCPoints(points []image.Point) C.struct_Points {
	cPointSlice := make([]C.struct_Point, len(points))
	for i, point := range points {
		cPointSlice[i] = C.struct_Point{
			x: C.int(point.X),
			y: C.int(point.Y),
		}
	}

	return C.struct_Points{
		points: (*C.Point)(&cPointSlice[0]),
		length: C.int(len(points)),
	}
}

func toCStrings(strs []string) C.struct_CStrings {
	cStringsSlice := make([]*C.char, len(strs))
	for i, s := range strs {
		cStringsSlice[i] = C.CString(s)
	}

	return C.struct_CStrings{
		strs:   (**C.char)(&cStringsSlice[0]),
		length: C.int(len(strs)),
	}
}
