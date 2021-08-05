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
	MatTypeCV8S MatType = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U MatType = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S MatType = 3

	// MatTypeCV16SC2 is a Mat of 16-bit signed int with 2 channels
	MatTypeCV16SC2 = MatTypeCV16S + MatChannels2

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S MatType = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F MatType = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F MatType = 6

	// MatTypeCV8UC1 is a Mat of 8-bit unsigned int with a single channel
	MatTypeCV8UC1 = MatTypeCV8U + MatChannels1

	// MatTypeCV8UC2 is a Mat of 8-bit unsigned int with 2 channels
	MatTypeCV8UC2 = MatTypeCV8U + MatChannels2

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4

	// MatTypeCV8SC1 is a Mat of 8-bit signed int with a single channel
	MatTypeCV8SC1 = MatTypeCV8S + MatChannels1

	// MatTypeCV8SC2 is a Mat of 8-bit signed int with 2 channels
	MatTypeCV8SC2 = MatTypeCV8S + MatChannels2

	// MatTypeCV8SC3 is a Mat of 8-bit signed int with 3 channels
	MatTypeCV8SC3 = MatTypeCV8S + MatChannels3

	// MatTypeCV8SC4 is a Mat of 8-bit signed int with 4 channels
	MatTypeCV8SC4 = MatTypeCV8S + MatChannels4

	// MatTypeCV16UC1 is a Mat of 16-bit unsigned int with a single channel
	MatTypeCV16UC1 = MatTypeCV16U + MatChannels1

	// MatTypeCV16UC2 is a Mat of 16-bit unsigned int with 2 channels
	MatTypeCV16UC2 = MatTypeCV16U + MatChannels2

	// MatTypeCV16UC3 is a Mat of 16-bit unsigned int with 3 channels
	MatTypeCV16UC3 = MatTypeCV16U + MatChannels3

	// MatTypeCV16UC4 is a Mat of 16-bit unsigned int with 4 channels
	MatTypeCV16UC4 = MatTypeCV16U + MatChannels4

	// MatTypeCV16SC1 is a Mat of 16-bit signed int with a single channel
	MatTypeCV16SC1 = MatTypeCV16S + MatChannels1

	// MatTypeCV16SC3 is a Mat of 16-bit signed int with 3 channels
	MatTypeCV16SC3 = MatTypeCV16S + MatChannels3

	// MatTypeCV16SC4 is a Mat of 16-bit signed int with 4 channels
	MatTypeCV16SC4 = MatTypeCV16S + MatChannels4

	// MatTypeCV32SC1 is a Mat of 32-bit signed int with a single channel
	MatTypeCV32SC1 = MatTypeCV32S + MatChannels1

	// MatTypeCV32SC2 is a Mat of 32-bit signed int with 2 channels
	MatTypeCV32SC2 = MatTypeCV32S + MatChannels2

	// MatTypeCV32SC3 is a Mat of 32-bit signed int with 3 channels
	MatTypeCV32SC3 = MatTypeCV32S + MatChannels3

	// MatTypeCV32SC4 is a Mat of 32-bit signed int with 4 channels
	MatTypeCV32SC4 = MatTypeCV32S + MatChannels4

	// MatTypeCV32FC1 is a Mat of 32-bit float int with a single channel
	MatTypeCV32FC1 = MatTypeCV32F + MatChannels1

	// MatTypeCV32FC2 is a Mat of 32-bit float int with 2 channels
	MatTypeCV32FC2 = MatTypeCV32F + MatChannels2

	// MatTypeCV32FC3 is a Mat of 32-bit float int with 3 channels
	MatTypeCV32FC3 = MatTypeCV32F + MatChannels3

	// MatTypeCV32FC4 is a Mat of 32-bit float int with 4 channels
	MatTypeCV32FC4 = MatTypeCV32F + MatChannels4

	// MatTypeCV64FC1 is a Mat of 64-bit float int with a single channel
	MatTypeCV64FC1 = MatTypeCV64F + MatChannels1

	// MatTypeCV64FC2 is a Mat of 64-bit float int with 2 channels
	MatTypeCV64FC2 = MatTypeCV64F + MatChannels2

	// MatTypeCV64FC3 is a Mat of 64-bit float int with 3 channels
	MatTypeCV64FC3 = MatTypeCV64F + MatChannels3

	// MatTypeCV64FC4 is a Mat of 64-bit float int with 4 channels
	MatTypeCV64FC4 = MatTypeCV64F + MatChannels4
)

// CompareType is used for Compare operations to indicate which kind of
// comparison to use.
type CompareType int

const (
	// CompareEQ src1 is equal to src2.
	CompareEQ CompareType = 0

	// CompareGT src1 is greater than src2.
	CompareGT CompareType = 1

	// CompareGE src1 is greater than or equal to src2.
	CompareGE CompareType = 2

	// CompareLT src1 is less than src2.
	CompareLT CompareType = 3

	// CompareLE src1 is less than or equal to src2.
	CompareLE CompareType = 4

	// CompareNE src1 is unequal to src2.
	CompareNE CompareType = 5
)

type Point2f struct {
	X float32
	Y float32
}

func NewPoint2f(x, y float32) Point2f {
	return Point2f{x, y}
}

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

	// Non-nil if Mat was created with a []byte (using NewMatFromBytes()). Nil otherwise.
	d []byte
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return newMat(C.Mat_New())
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return newMat(C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt)))
}

// NewMatWithSizes returns a new multidimensional Mat with a specific size and type.
func NewMatWithSizes(sizes []int, mt MatType) Mat {
	sizesArray := make([]C.int, len(sizes))
	for i, s := range sizes {
		sizesArray[i] = C.int(s)
	}

	sizesIntVector := C.IntVector{
		val:    (*C.int)(&sizesArray[0]),
		length: C.int(len(sizes)),
	}
	return newMat(C.Mat_NewWithSizes(sizesIntVector, C.int(mt)))
}

// NewMatWithSizesWithScalar returns a new multidimensional Mat with a specific size, type and scalar value.
func NewMatWithSizesWithScalar(sizes []int, mt MatType, s Scalar) Mat {
	csizes := []C.int{}
	for _, v := range sizes {
		csizes = append(csizes, C.int(v))
	}
	sizesVector := C.struct_IntVector{}
	sizesVector.val = (*C.int)(&csizes[0])
	sizesVector.length = (C.int)(len(csizes))

	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	return newMat(C.Mat_NewWithSizesFromScalar(sizesVector, C.int(mt), sVal))
}

// NewMatWithSizesWithScalar returns a new multidimensional Mat with a specific size, type and preexisting data.
func NewMatWithSizesFromBytes(sizes []int, mt MatType, data []byte) (Mat, error) {
	cBytes, err := toByteArray(data)
	if err != nil {
		return Mat{}, err
	}

	csizes := []C.int{}
	for _, v := range sizes {
		csizes = append(csizes, C.int(v))
	}
	sizesVector := C.struct_IntVector{}
	sizesVector.val = (*C.int)(&csizes[0])
	sizesVector.length = (C.int)(len(csizes))

	return newMat(C.Mat_NewWithSizesFromBytes(sizesVector, C.int(mt), *cBytes)), nil
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
	mat := newMat(C.Mat_NewFromBytes(C.int(rows), C.int(cols), C.int(mt), *cBytes))

	// Store a reference to the backing data slice. This is needed because we pass the backing
	// array directly to C code and without keeping a Go reference to it, it might end up
	// garbage collected which would result in crashes.
	//
	// TODO(bga): This could live in newMat() but I wanted to reduce the change surface.
	// TODO(bga): Code that needs access to the array from Go could use this directly.
	mat.d = data

	return mat, nil
}

// Returns an identity matrix of the specified size and type.
//
// The method returns a Matlab-style identity matrix initializer, similarly to Mat::zeros. Similarly to Mat::ones.
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a2cf9b9acde7a9852542bbc20ef851ed2
func Eye(rows int, cols int, mt MatType) Mat {
	return newMat(C.Eye(C.int(rows), C.int(cols), C.int(mt)))
}

// Returns a zero array of the specified size and type.
//
// The method returns a Matlab-style zero array initializer.
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a0b57b6a326c8876d944d188a46e0f556
func Zeros(rows int, cols int, mt MatType) Mat {
	return newMat(C.Zeros(C.int(rows), C.int(cols), C.int(mt)))
}

// Returns an array of all 1's of the specified size and type.
//
// The method returns a Matlab-style 1's array initializer
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a69ae0402d116fc9c71908d8508dc2f09
func Ones(rows int, cols int, mt MatType) Mat {
	return newMat(C.Ones(C.int(rows), C.int(cols), C.int(mt)))
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

// IsContinuous determines if the Mat is continuous.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa90cea495029c7d1ee0a41361ccecdf3
//
func (m *Mat) IsContinuous() bool {
	return bool(C.Mat_IsContinuous(m.p))
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

func (m *Mat) ConvertToWithParams(dst *Mat, mt MatType, alpha, beta float32) {
	C.Mat_ConvertToWithParams(m.p, dst.p, C.int(mt), C.float(alpha), C.float(beta))
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
	defer C.IntVector_Close(cdims)

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
func (m *Mat) DataPtrUint8() ([]uint8, error) {
	if !m.IsContinuous() {
		return nil, errors.New("DataPtrUint8 requires continuous Mat")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length),
		Cap:  int(p.length),
	}
	return *(*[]uint8)(unsafe.Pointer(h)), nil
}

// DataPtrInt8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrInt8() ([]int8, error) {
	if m.Type()&MatTypeCV8S != MatTypeCV8S {
		return nil, errors.New("DataPtrInt8 only supports MatTypeCV8S")
	}

	if !m.IsContinuous() {
		return nil, errors.New("DataPtrInt8 requires continuous Mat")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length),
		Cap:  int(p.length),
	}
	return *(*[]int8)(unsafe.Pointer(h)), nil
}

// DataPtrUint16 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint16() ([]uint16, error) {
	if m.Type()&MatTypeCV16U != MatTypeCV16U {
		return nil, errors.New("DataPtrUint16 only supports MatTypeCV16U")
	}

	if !m.IsContinuous() {
		return nil, errors.New("DataPtrUint16 requires continuous Mat")
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

	if !m.IsContinuous() {
		return nil, errors.New("DataPtrInt16 requires continuous Mat")
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

	if !m.IsContinuous() {
		return nil, errors.New("DataPtrFloat32 requires continuous Mat")
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

	if !m.IsContinuous() {
		return nil, errors.New("DataPtrFloat64 requires continuous Mat")
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
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga191389f8a0e58180bb13a727782cd461
//
func (m *Mat) Mean() Scalar {
	s := C.Mat_Mean(m.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// MeanWithMask calculates the mean value M of array elements,independently for each channel,
// and returns it as Scalar vector while applying the mask.
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga191389f8a0e58180bb13a727782cd461
//
func (m *Mat) MeanWithMask(mask Mat) Scalar {
	s := C.Mat_MeanWithMask(m.p, mask.p)
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

// MultiplyMatrix multiplies matrix (m*x)
func (m *Mat) MultiplyMatrix(x Mat) Mat {
	return newMat(C.Mat_MultiplyMatrix(m.p, x.p))
}

// T  transpose matrix
// https://docs.opencv.org/4.1.2/d3/d63/classcv_1_1Mat.html#aaa428c60ccb6d8ea5de18f63dfac8e11
func (m *Mat) T() Mat {
	return newMat(C.Mat_T(m.p))
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
func BatchDistance(src1 Mat, src2 Mat, dist Mat, dtype MatType, nidx Mat, normType NormType, K int, mask Mat, update int, crosscheck bool) {
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
	CovarNormal CovarFlags = 1

	// CovarUseAvg indicates to use average covariation.
	CovarUseAvg CovarFlags = 2

	// CovarScale indicates to use scaled covariation.
	CovarScale CovarFlags = 4

	// CovarRows indicates to use covariation on rows.
	CovarRows CovarFlags = 8

	// CovarCols indicates to use covariation on columns.
	CovarCols CovarFlags = 16
)

// CalcCovarMatrix calculates the covariance matrix of a set of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga017122d912af19d7d0d2cccc2d63819f
//
func CalcCovarMatrix(samples Mat, covar *Mat, mean *Mat, flags CovarFlags, ctype MatType) {
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
	DftInverse DftFlags = 1

	// DftScale scales the result: divide it by the number of array elements. Normally, it is combined with DFT_INVERSE.
	DftScale DftFlags = 2

	// DftRows performs a forward or inverse transform of every individual row of the input matrix.
	DftRows DftFlags = 4

	// DftComplexOutput performs a forward transformation of 1D or 2D real array; the result, though being a complex array, has complex-conjugate symmetry
	DftComplexOutput DftFlags = 16

	// DftRealOutput performs an inverse transformation of a 1D or 2D complex array; the result is normally a complex array of the same size,
	// however, if the input array has conjugate-complex symmetry (for example, it is a result of forward transformation with DFT_COMPLEX_OUTPUT flag),
	// the output is a real array.
	DftRealOutput DftFlags = 32

	// DftComplexInput specifies that input is complex input. If this flag is set, the input must have 2 channels.
	DftComplexInput DftFlags = 64

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
	Rotate180Clockwise RotateFlag = 1
	// Rotate90CounterClockwise allows to rotate 270 degrees clockwise
	Rotate90CounterClockwise RotateFlag = 2
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
func Invert(src Mat, dst *Mat, flags SolveDecompositionFlags) float64 {
	ret := C.Mat_Invert(src.p, dst.p, C.int(flags))
	return float64(ret)
}

// KMeansFlags for kmeans center selection
//
// For further details, please see:
// https://docs.opencv.org/master/d0/de1/group__core.html#ga276000efe55ee2756e0c471c7b270949
type KMeansFlags int

const (
	// KMeansRandomCenters selects random initial centers in each attempt.
	KMeansRandomCenters KMeansFlags = 0
	// KMeansPPCenters uses kmeans++ center initialization by Arthur and Vassilvitskii [Arthur2007].
	KMeansPPCenters KMeansFlags = 1
	// KMeansUseInitialLabels uses the user-supplied lables during the first (and possibly the only) attempt
	// instead of computing them from the initial centers. For the second and further attempts, use the random or semi-random     // centers. Use one of KMEANS_*_CENTERS flag to specify the exact method.
	KMeansUseInitialLabels KMeansFlags = 2
)

// KMeans finds centers of clusters and groups input samples around the clusters.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga9a34dc06c6ec9460e90860f15bcd2f88
//
func KMeans(data Mat, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
	ret := C.KMeans(data.p, C.int(k), bestLabels.p, criteria.p, C.int(attempts), C.int(flags), centers.p)
	return float64(ret)
}

// KMeansPoints finds centers of clusters and groups input samples around the clusters.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga9a34dc06c6ec9460e90860f15bcd2f88
//
func KMeansPoints(points PointVector, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
	ret := C.KMeansPoints(points.p, C.int(k), bestLabels.p, criteria.p, C.int(attempts), C.int(flags), centers.p)
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

// Copies specified channels from input arrays to the specified channels of output arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga51d768c270a1cdd3497255017c4504be
//
func MixChannels(src []Mat, dst []Mat, fromTo []int) {
	cSrcArray := make([]C.Mat, len(src))
	for i, r := range src {
		cSrcArray[i] = r.p
	}
	cSrcMats := C.struct_Mats{
		mats:   (*C.Mat)(&cSrcArray[0]),
		length: C.int(len(src)),
	}

	cDstArray := make([]C.Mat, len(dst))
	for i, r := range dst {
		cDstArray[i] = r.p
	}
	cDstMats := C.struct_Mats{
		mats:   (*C.Mat)(&cDstArray[0]),
		length: C.int(len(dst)),
	}

	cFromToArray := make([]C.int, len(fromTo))
	for i, ft := range fromTo {
		cFromToArray[i] = C.int(ft)
	}

	cFromToIntVector := C.IntVector{
		val:    (*C.int)(&cFromToArray[0]),
		length: C.int(len(fromTo)),
	}

	C.Mat_MixChannels(cSrcMats, cDstMats, cFromToIntVector)

	for i := C.int(0); i < cDstMats.length; i++ {
		dst[i].p = C.Mats_get(cDstMats, i)
	}
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

// MultiplyWithParams calculates the per-element scaled product of two arrays.
// Both input arrays must be of the same size and the same type.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga979d898a58d7f61c53003e162e7ad89f
//
func MultiplyWithParams(src1 Mat, src2 Mat, dst *Mat, scale float64, dtype MatType) {
	C.Mat_MultiplyWithParams(src1.p, src2.p, dst.p, C.double(scale), C.int(dtype))
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
	NormL1 NormType = 2

	// NormL2 indicates use L2 normalization.
	NormL2 NormType = 4

	// NormL2Sqr indicates use L2 squared normalization.
	NormL2Sqr NormType = 5

	// NormHamming indicates use Hamming normalization.
	NormHamming NormType = 6

	// NormHamming2 indicates use Hamming 2-bit normalization.
	NormHamming2 NormType = 7

	// NormTypeMask indicates use type mask for normalization.
	NormTypeMask NormType = 7

	// NormRelative indicates use relative normalization.
	NormRelative NormType = 8

	// NormMinMax indicates use min/max normalization.
	NormMinMax NormType = 32
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

// Norm calculates the absolute difference/relative norm of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33
//
func NormWithMats(src1 Mat, src2 Mat, normType NormType) float64 {
	return float64(C.NormWithMats(src1.p, src2.p, C.int(normType)))
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
	MaxIter TermCriteriaType = 1

	// EPS is the desired accuracy or change in parameters at which the
	// iterative algorithm stops.
	EPS TermCriteriaType = 2
)

type SolveDecompositionFlags int

const (
	// Gaussian elimination with the optimal pivot element chosen.
	SolveDecompositionLu SolveDecompositionFlags = 0

	// Singular value decomposition (SVD) method. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionSvd SolveDecompositionFlags = 1

	// Eigenvalue decomposition. The matrix src1 must be symmetrical.
	SolveDecompositionEing SolveDecompositionFlags = 2

	// Cholesky LL^T factorization. The matrix src1 must be symmetrical and positively defined.
	SolveDecompositionCholesky SolveDecompositionFlags = 3

	// QR factorization. The system can be over-defined and/or the matrix src1 can be singular.
	SolveDecompositionQr SolveDecompositionFlags = 4

	// While all the previous flags are mutually exclusive, this flag can be used together with any of the previous.
	// It means that the normal equations 𝚜𝚛𝚌𝟷^T⋅𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟷^T𝚜𝚛𝚌𝟸 are solved instead of the original system
	// 𝚜𝚛𝚌𝟷⋅𝚍𝚜𝚝=𝚜𝚛𝚌𝟸.
	SolveDecompositionNormal SolveDecompositionFlags = 5
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
func Reduce(src Mat, dst *Mat, dim int, rType ReduceTypes, dType MatType) {
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

// SetIdentity initializes a scaled identity matrix.
// For further details, please see:
//  https://docs.opencv.org/master/d2/de8/group__core__array.html#ga388d7575224a4a277ceb98ccaa327c99
//
func SetIdentity(src Mat, scalar float64) {
	C.Mat_SetIdentity(src.p, C.double(scalar))
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
// Created images should be closed manualy to avoid memory leaks.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0547c7fed86152d7e9d0096029c8518a
//
func Split(src Mat) (mv []Mat) {
	cMats := C.struct_Mats{}
	C.Mat_Split(src.p, &(cMats))
	defer C.Mats_Close(cMats)
	mv = make([]Mat, cMats.length)
	for i := C.int(0); i < cMats.length; i++ {
		mv[i].p = C.Mats_get(cMats, i)
		addMatToProfile(mv[i].p)
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

// Vecb is a generic vector of bytes.
type Vecb []uint8

// GetVecbAt returns a vector of bytes. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVecbAt(row int, col int) Vecb {
	ch := m.Channels()
	v := make(Vecb, ch)

	for c := 0; c < ch; c++ {
		v[c] = m.GetUCharAt(row, col*ch+c)
	}

	return v
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

// Vecd is a generic vector of float64/doubles.
type Vecd []float64

// GetVecdAt returns a vector of float64s. Its size corresponds to the number
// of channels of the Mat.
func (m *Mat) GetVecdAt(row int, col int) Vecd {
	ch := m.Channels()
	v := make(Vecd, ch)

	for c := 0; c < ch; c++ {
		v[c] = m.GetDoubleAt(row, col*ch+c)
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

// PointVector is a wrapper around a std::vector< cv::Point >*
// This is needed anytime that you need to pass or receive a collection of points.
type PointVector struct {
	p C.PointVector
}

// NewPointVector returns a new empty PointVector.
func NewPointVector() PointVector {
	return PointVector{p: C.PointVector_New()}
}

// NewPointVectorFromPoints returns a new PointVector that has been
// initialized to a slice of image.Point.
func NewPointVectorFromPoints(pts []image.Point) PointVector {
	p := (*C.struct_Point)(C.malloc(C.size_t(C.sizeof_struct_Point * len(pts))))
	defer C.free(unsafe.Pointer(p))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  len(pts),
		Cap:  len(pts),
	}
	pa := *(*[]C.Point)(unsafe.Pointer(h))

	for j, point := range pts {
		pa[j] = C.struct_Point{
			x: C.int(point.X),
			y: C.int(point.Y),
		}
	}

	cpoints := C.struct_Points{
		points: (*C.Point)(p),
		length: C.int(len(pts)),
	}

	return PointVector{p: C.PointVector_NewFromPoints(cpoints)}
}

// NewPointVectorFromMat returns a new PointVector that has been
// wrapped around a Mat of type CV_32SC2 with a single columm.
func NewPointVectorFromMat(mat Mat) PointVector {
	return PointVector{p: C.PointVector_NewFromMat(mat.p)}
}

// IsNil checks the CGo pointer in the PointVector.
func (pv PointVector) IsNil() bool {
	return pv.p == nil
}

// Size returns how many Point are in the PointVector.
func (pv PointVector) Size() int {
	return int(C.PointVector_Size(pv.p))
}

// At returns the image.Point
func (pv PointVector) At(idx int) image.Point {
	if idx > pv.Size() {
		return image.Point{}
	}

	cp := C.PointVector_At(pv.p, C.int(idx))
	return image.Pt(int(cp.x), int(cp.y))
}

// Append appends an image.Point at end of the PointVector.
func (pv PointVector) Append(point image.Point) {
	p := C.struct_Point{
		x: C.int(point.X),
		y: C.int(point.Y),
	}

	C.PointVector_Append(pv.p, p)

	return
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pv PointVector) ToPoints() []image.Point {
	points := make([]image.Point, pv.Size())

	for j := 0; j < pv.Size(); j++ {
		points[j] = pv.At(j)
	}
	return points
}

// Close closes and frees memory for this PointVector.
func (pv PointVector) Close() {
	C.PointVector_Close(pv.p)
}

// PointsVector is a wrapper around a std::vector< std::vector< cv::Point > >*
type PointsVector struct {
	p C.PointsVector
}

// NewPointsVector returns a new empty PointsVector.
func NewPointsVector() PointsVector {
	return PointsVector{p: C.PointsVector_New()}
}

// NewPointsVectorFromPoints returns a new PointsVector that has been
// initialized to a slice of slices of image.Point.
func NewPointsVectorFromPoints(pts [][]image.Point) PointsVector {
	points := make([]C.struct_Points, len(pts))

	for i, pt := range pts {
		p := (*C.struct_Point)(C.malloc(C.size_t(C.sizeof_struct_Point * len(pt))))
		defer C.free(unsafe.Pointer(p))

		h := &reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(p)),
			Len:  len(pt),
			Cap:  len(pt),
		}
		pa := *(*[]C.Point)(unsafe.Pointer(h))

		for j, point := range pt {
			pa[j] = C.struct_Point{
				x: C.int(point.X),
				y: C.int(point.Y),
			}
		}

		points[i] = C.struct_Points{
			points: (*C.Point)(p),
			length: C.int(len(pt)),
		}
	}

	cPoints := C.struct_Contours{
		contours: (*C.struct_Points)(&points[0]),
		length:   C.int(len(pts)),
	}

	return PointsVector{p: C.PointsVector_NewFromPoints(cPoints)}
}

func (pvs PointsVector) P() C.PointsVector {
	return pvs.p
}

// ToPoints returns a slice of slices of image.Point for the data in this PointsVector.
func (pvs PointsVector) ToPoints() [][]image.Point {
	ppoints := make([][]image.Point, pvs.Size())
	for i := 0; i < pvs.Size(); i++ {
		pts := pvs.At(i)
		points := make([]image.Point, pts.Size())

		for j := 0; j < pts.Size(); j++ {
			points[j] = pts.At(j)
		}
		ppoints[i] = points
	}

	return ppoints
}

// IsNil checks the CGo pointer in the PointsVector.
func (pvs PointsVector) IsNil() bool {
	return pvs.p == nil
}

// Size returns how many vectors of Points are in the PointsVector.
func (pvs PointsVector) Size() int {
	return int(C.PointsVector_Size(pvs.p))
}

// At returns the PointVector at that index of the PointsVector.
func (pvs PointsVector) At(idx int) PointVector {
	if idx > pvs.Size() {
		return PointVector{}
	}

	return PointVector{p: C.PointsVector_At(pvs.p, C.int(idx))}
}

// Append appends a PointVector at end of the PointsVector.
func (pvs PointsVector) Append(pv PointVector) {
	if !pv.IsNil() {
		C.PointsVector_Append(pvs.p, pv.p)
	}

	return
}

// Close closes and frees memory for this PointsVector.
func (pvs PointsVector) Close() {
	C.PointsVector_Close(pvs.p)
}

// Point2fVector is a wrapper around a std::vector< cv::Point2f >*
// This is needed anytime that you need to pass or receive a collection of points.
type Point2fVector struct {
	p C.Point2fVector
}

// NewPoint2fVector returns a new empty Point2fVector.
func NewPoint2fVector() Point2fVector {
	return Point2fVector{p: C.Point2fVector_New()}
}

// NewPoint2fVectorFromPoints returns a new Point2fVector that has been
// initialized to a slice of image.Point.
func NewPoint2fVectorFromPoints(pts []Point2f) Point2fVector {
	p := (*C.struct_Point2f)(C.malloc(C.size_t(C.sizeof_struct_Point2f * len(pts))))
	defer C.free(unsafe.Pointer(p))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  len(pts),
		Cap:  len(pts),
	}
	pa := *(*[]C.Point2f)(unsafe.Pointer(h))

	for j, point := range pts {
		pa[j] = C.struct_Point2f{
			x: C.float(point.X),
			y: C.float(point.Y),
		}
	}

	cpoints := C.struct_Points2f{
		points: (*C.Point2f)(p),
		length: C.int(len(pts)),
	}

	return Point2fVector{p: C.Point2fVector_NewFromPoints(cpoints)}
}

// NewPoint2fVectorFromMat returns a new Point2fVector that has been
// wrapped around a Mat of type CV_32FC2 with a single columm.
func NewPoint2fVectorFromMat(mat Mat) Point2fVector {
	return Point2fVector{p: C.Point2fVector_NewFromMat(mat.p)}
}

// IsNil checks the CGo pointer in the Point2fVector.
func (pfv Point2fVector) IsNil() bool {
	return pfv.p == nil
}

// Size returns how many Point are in the PointVector.
func (pfv Point2fVector) Size() int {
	return int(C.Point2fVector_Size(pfv.p))
}

// At returns the image.Point
func (pfv Point2fVector) At(idx int) Point2f {
	if idx > pfv.Size() {
		return Point2f{}
	}

	cp := C.Point2fVector_At(pfv.p, C.int(idx))
	return Point2f{float32(cp.x), float32(cp.y)}
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pfv Point2fVector) ToPoints() []Point2f {
	points := make([]Point2f, pfv.Size())

	for j := 0; j < pfv.Size(); j++ {
		points[j] = pfv.At(j)
	}
	return points
}

// Close closes and frees memory for this Point2fVector.
func (pfv Point2fVector) Close() {
	C.Point2fVector_Close(pfv.p)
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

// Converts CStrings to a slice of Go strings even when the C strings are not contiguous in memory
func toGoStrings(strs C.CStrings) []string {
	length := int(strs.length)
	tmpslice := (*[1 << 20]*C.char)(unsafe.Pointer(strs.strs))[:length:length]
	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
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

func toCPoints2f(points []Point2f) C.struct_Points2f {
	cPointSlice := make([]C.struct_Point2f, len(points))
	for i, point := range points {
		cPointSlice[i] = C.struct_Point2f{
			x: C.float(point.X),
			y: C.float(point.Y),
		}
	}

	return C.struct_Points2f{
		points: (*C.Point2f)(&cPointSlice[0]),
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

// RowRange creates a matrix header for the specified row span.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa6542193430356ad631a9beabc624107
//
func (m *Mat) RowRange(start, end int) Mat {
	return newMat(C.Mat_rowRange(m.p, C.int(start), C.int(end)))
}

// ColRange creates a matrix header for the specified column span.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aadc8f9210fe4dec50513746c246fa8d9
//
func (m *Mat) ColRange(start, end int) Mat {
	return newMat(C.Mat_colRange(m.p, C.int(start), C.int(end)))
}

// RNG Random Number Generator.
// It encapsulates the state (currently, a 64-bit integer) and
// has methods to return scalar random values and to fill arrays
// with random values
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html
//
type RNG struct {
	p C.RNG
}

type RNGDistType int

const (
	// Uniform distribution
	RNGDistUniform RNGDistType = 0
	// Normal distribution
	RNGDistNormal RNGDistType = 1
)

// TheRNG Returns the default random number generator.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722
//
func TheRNG() RNG {
	return RNG{
		p: C.TheRNG(),
	}
}

// TheRNG Sets state of default random number generator.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f
//
func SetRNGSeed(seed int) {
	C.SetRNGSeed(C.int(seed))
}

// Fill Fills arrays with random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#ad26f2b09d9868cf108e84c9814aa682d
//
func (r *RNG) Fill(mat *Mat, distType RNGDistType, a, b float64, saturateRange bool) {
	C.RNG_Fill(r.p, mat.p, C.int(distType), C.double(a), C.double(b), C.bool(saturateRange))
}

// Gaussian Returns the next random number sampled from
// the Gaussian distribution.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#a8df8ce4dc7d15916cee743e5a884639d
//
func (r *RNG) Gaussian(sigma float64) float64 {
	return float64(C.RNG_Gaussian(r.p, C.double(sigma)))
}

// Next The method updates the state using the MWC algorithm
// and returns the next 32-bit random number.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/dd6/classcv_1_1RNG.html#a8df8ce4dc7d15916cee743e5a884639d
//
func (r *RNG) Next() uint {
	return uint(C.RNG_Next(r.p))
}

// RandN Fills the array with normally distributed random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808
//
func RandN(mat *Mat, mean, stddev Scalar) {
	meanVal := C.struct_Scalar{
		val1: C.double(mean.Val1),
		val2: C.double(mean.Val2),
		val3: C.double(mean.Val3),
		val4: C.double(mean.Val4),
	}
	stddevVal := C.struct_Scalar{
		val1: C.double(stddev.Val1),
		val2: C.double(stddev.Val2),
		val3: C.double(stddev.Val3),
		val4: C.double(stddev.Val4),
	}

	C.RandN(mat.p, meanVal, stddevVal)
}

// RandShuffle Shuffles the array elements randomly.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763
//
func RandShuffle(mat *Mat) {
	C.RandShuffle(mat.p)
}

// RandShuffleWithParams Shuffles the array elements randomly.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763
//
func RandShuffleWithParams(mat *Mat, iterFactor float64, rng RNG) {
	C.RandShuffleWithParams(mat.p, C.double(iterFactor), rng.p)
}

// RandU Generates a single uniformly-distributed random
// number or an array of random numbers.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0
//
func RandU(mat *Mat, low, high Scalar) {
	lowVal := C.struct_Scalar{
		val1: C.double(low.Val1),
		val2: C.double(low.Val2),
		val3: C.double(low.Val3),
		val4: C.double(low.Val4),
	}
	highVal := C.struct_Scalar{
		val1: C.double(high.Val1),
		val2: C.double(high.Val2),
		val3: C.double(high.Val3),
		val4: C.double(high.Val4),
	}

	C.RandU(mat.p, lowVal, highVal)
}

type NativeByteBuffer struct {
	// std::vector is build of 3 pointers And this will not change ever.
	stdVectorOpaq [3]uintptr
}

func newNativeByteBuffer() *NativeByteBuffer {
	buffer := &NativeByteBuffer{}
	C.StdByteVectorInitialize(buffer.nativePointer())
	return buffer
}

func (buffer *NativeByteBuffer) nativePointer() unsafe.Pointer {
	return unsafe.Pointer(&buffer.stdVectorOpaq[0])
}

func (buffer *NativeByteBuffer) dataPointer() unsafe.Pointer {
	return unsafe.Pointer(C.StdByteVectorData(buffer.nativePointer()))
}

// GetBytes returns slice of bytes backed by native buffer
func (buffer *NativeByteBuffer) GetBytes() []byte {
	var result []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	vectorLen := int(C.StdByteVectorLen(buffer.nativePointer()))
	sliceHeader.Cap = vectorLen
	sliceHeader.Len = vectorLen
	sliceHeader.Data = uintptr(buffer.dataPointer())
	return result
}

// Len - returns length in bytes of underlying buffer
func (buffer *NativeByteBuffer) Len() int {
	return int(C.StdByteVectorLen(buffer.nativePointer()))
}

// Close the buffer releasing all its resources
func (buffer *NativeByteBuffer) Close() {
	C.StdByteVectorFree(buffer.nativePointer())
}
// Points2fVector is a wrapper around a std::vector< std::vector< cv::Point2f > >*
type Points2fVector struct {
	p C.Points2fVector
}

// NewPoints2fVector returns a new empty Points2fVector.
func NewPoints2fVector() Points2fVector {
	return Points2fVector{p: C.Points2fVector_New()}
}

// NewPoints2fVectorFromPoints returns a new Points2fVector that has been
// initialized to a slice of slices of Point2f.
func NewPoints2fVectorFromPoints(pts [][]Point2f) Points2fVector {
	pvf := NewPoints2fVector()
	for j := 0;j<len(pts);j++{
		pv := NewPoint2fVectorFromPoints(pts[j])
		pvf.Append(pv)
		pv.Close()
	}
	return pvf
}

func (pvs Points2fVector) P() C.Points2fVector {
	return pvs.p
}

// ToPoints returns a slice of slices of Point2f for the data in this Points2fVector.
func (pvs Points2fVector) ToPoints() [][]Point2f {
	ppoints := make([][]Point2f, pvs.Size())
	for j := 0;j < pvs.Size();j++{
		pts := pvs.At(j)
		points := pts.ToPoints()
		ppoints[j] = points
	}
	return ppoints
}

// IsNil checks the CGo pointer in the Points2fVector.
func (pvs Points2fVector) IsNil() bool {
	return pvs.p == nil
}

// Size returns how many vectors of Points are in the Points2fVector.
func (pvs Points2fVector) Size() int {
	return int(C.Points2fVector_Size(pvs.p))
}

// At returns the Point2fVector at that index of the Points2fVector.
func (pvs Points2fVector) At(idx int) Point2fVector {
	if idx > pvs.Size() {
		return Point2fVector{}
	}
	return Point2fVector{p : C.Points2fVector_At(pvs.p, C.int(idx))}
}

// Append appends a Point2fVector at end of the Points2fVector.
func (pvs Points2fVector) Append(pv Point2fVector) {
	if !pv.IsNil() {
		C.Points2fVector_Append(pvs.p, pv.p)
	}
}

// Close closes and frees memory for this Points2fVector.
func (pvs Points2fVector) Close() {
	C.Points2fVector_Close(pvs.p)
}

type Point3f struct {
	X float32
	Y float32
	Z float32
}

func NewPoint3f(x, y, z float32) Point3f {
	return Point3f{x, y, z}
}

// Point3fVector is a wrapper around a std::vector< cv::Point3f >*
type Point3fVector struct {
	p C.Point3fVector
}

// NewPoint3fVector returns a new empty Point3fVector.
func NewPoint3fVector() Point3fVector {
	return Point3fVector{p: C.Point3fVector_New()}
}

// NewPoint3fVectorFromPoints returns a new Point3fVector that has been
// initialized to a slice of image.Point.
func NewPoint3fVectorFromPoints(pts []Point3f) Point3fVector {
	p := (*C.struct_Point3f)(C.malloc(C.size_t(C.sizeof_struct_Point3f * len(pts))))
	defer C.free(unsafe.Pointer(p))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  len(pts),
		Cap:  len(pts),
	}
	pa := *(*[]C.Point3f)(unsafe.Pointer(h))

	for j, point := range pts {
		pa[j] = C.struct_Point3f{
			x: C.float(point.X),
			y: C.float(point.Y),
			z: C.float(point.Z),
		}
	}

	cPoints := C.struct_Points3f{
		points: (*C.Point3f)(p),
		length: C.int(len(pts)),
	}

	return Point3fVector{p: C.Point3fVector_NewFromPoints(cPoints)}
}

// NewPoint3fVectorFromMat returns a new Point3fVector that has been
// wrapped around a Mat of type CV_32FC3 with a single columm.
func NewPoint3fVectorFromMat(mat Mat) Point3fVector {
	return Point3fVector{p: C.Point3fVector_NewFromMat(mat.p)}
}

// IsNil checks the CGo pointer in the Point3fVector.
func (pfv Point3fVector) IsNil() bool {
	return pfv.p == nil
}

// Size returns how many Point are in the Point3fVector.
func (pfv Point3fVector) Size() int {
	return int(C.Point3fVector_Size(pfv.p))
}

// At returns the Point3f
func (pfv Point3fVector) At(idx int) Point3f {
	if idx > pfv.Size() {
		return Point3f{}
	}
	cp := C.Point3fVector_At(pfv.p, C.int(idx))
	return Point3f{X: float32(cp.x), Y: float32(cp.y), Z: float32(cp.z)}
}

func (pfv Point3fVector) Append(point Point3f) {
	C.Point3fVector_Append(pfv.p, C.Point3f{
		x: C.float(point.X),
		y: C.float(point.Y),
		z: C.float(point.Z),
	});
}

// ToPoints returns a slice of Point3f for the data in this Point3fVector.
func (pfv Point3fVector) ToPoints() []Point3f {
	points := make([]Point3f, pfv.Size())
	for j := 0; j < pfv.Size(); j++ {
		points[j] = pfv.At(j)
	}
	return points
}

// Close closes and frees memory for this Point3fVector.
func (pfv Point3fVector) Close() {
	C.Point3fVector_Close(pfv.p)
}

// Points3fVector is a wrapper around a std::vector< std::vector< cv::Point3f > >*
type Points3fVector struct {
	p C.Points3fVector
}

// NewPoints3fVector returns a new empty Points3fVector.
func NewPoints3fVector() Points3fVector {
	return Points3fVector{p: C.Points3fVector_New()}
}

// NewPoints3fVectorFromPoints returns a new Points3fVector that has been
// initialized to a slice of slices of Point3f.
func NewPoints3fVectorFromPoints(pts [][]Point3f) Points3fVector {
	pvf := NewPoints3fVector()
	for j := 0;j<len(pts);j++{
		pv := NewPoint3fVectorFromPoints(pts[j])
		pvf.Append(pv)
		pv.Close()
	}
	return pvf
}

// ToPoints returns a slice of slices of Point3f for the data in this Points3fVector.
func (pvs Points3fVector) ToPoints() [][]Point3f {
	ppoints := make([][]Point3f, pvs.Size())
	for j := 0;j < pvs.Size();j++{
		pts := pvs.At(j)
		points := pts.ToPoints()
		ppoints[j] = points
	}
	return ppoints
}

// IsNil checks the CGo pointer in the Points3fVector.
func (pvs Points3fVector) IsNil() bool {
	return pvs.p == nil
}

// Size returns how many vectors of Points are in the Points3fVector.
func (pvs Points3fVector) Size() int {
	return int(C.Points3fVector_Size(pvs.p))
}

// At returns the Point3fVector at that index of the Points3fVector.
func (pvs Points3fVector) At(idx int) Point3fVector {
	if idx > pvs.Size() {
		return Point3fVector{}
	}
	return Point3fVector{p : C.Points3fVector_At(pvs.p, C.int(idx))}
}

// Append appends a Point3fVector at end of the Points3fVector.
func (pvs Points3fVector) Append(pv Point3fVector) {
	if !pv.IsNil() {
		C.Points3fVector_Append(pvs.p, pv.p)
	}
}

// Close closes and frees memory for this Points3fVector.
func (pvs Points3fVector) Close() {
	C.Points3fVector_Close(pvs.p)
}