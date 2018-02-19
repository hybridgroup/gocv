package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"image"
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

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/3.4.0/d3/d63/classcv_1_1Mat.html
//
type Mat struct {
	p C.Mat
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return Mat{p: C.Mat_New()}
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return Mat{p: C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt))}
}

// NewMatFromScalar returns a new Mat for a specific Scalar value
func NewMatFromScalar(s Scalar, mt MatType) Mat {
	sVal := C.struct_Scalar{
		val1: C.double(s.Val1),
		val2: C.double(s.Val2),
		val3: C.double(s.Val3),
		val4: C.double(s.Val4),
	}

	return Mat{p: C.Mat_NewFromScalar(sVal, C.int(mt))}
}

// NewMatFromBytes returns a new Mat with a specific size and type, initialized from a []byte.
func NewMatFromBytes(rows int, cols int, mt MatType, data []byte) Mat {
	return Mat{p: C.Mat_NewFromBytes(C.int(rows), C.int(cols), C.int(mt), toByteArray(data))}
}

// Close the Mat object.
func (m *Mat) Close() error {
	C.Mat_Close(m.p)
	m.p = nil
	return nil
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
	return Mat{p: C.Mat_Clone(m.p)}
}

// CopyTo copies Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d3/d63/classcv_1_1Mat.html#a33fd5d125b4c302b0c9aa86980791a77
//
func (m *Mat) CopyTo(dst Mat) {
	C.Mat_CopyTo(m.p, dst.p)
	return
}

// CopyToWithMask copies Mat into destination Mat after applying the mask Mat.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d3/d63/classcv_1_1Mat.html#a626fe5f96d02525e2604d2ad46dd574f
//
func (m *Mat) CopyToWithMask(dst Mat, mask Mat) {
	C.Mat_CopyToWithMask(m.p, dst.p, mask.p)
	return
}

// ConvertTo converts Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b
//
func (m *Mat) ConvertTo(dst Mat, mt MatType) {
	C.Mat_ConvertTo(m.p, dst.p, C.int(mt))
	return
}

// ToBytes copies the underlying Mat data to a byte array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d3/d63/classcv_1_1Mat.html#a4d33bed1c850265370d2af0ff02e1564
func (m *Mat) ToBytes() []byte {
	b := C.Mat_ToBytes(m.p)
	defer C.ByteArray_Release(b)
	return toGoBytes(b)
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

	return Mat{p: C.Mat_Region(m.p, cRect)}
}

// Reshape changes the shape and/or the number of channels of a 2D matrix without copying the data.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8
//
func (m *Mat) Reshape(cn int, rows int) Mat {
	return Mat{p: C.Mat_Reshape(m.p, C.int(cn), C.int(rows))}
}

// ConvertFp16 converts a Mat to half-precision floating point.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga9c25d9ef44a2a48ecc3774b30cb80082
//
func (m *Mat) ConvertFp16() Mat {
	return Mat{p: C.Mat_ConvertFp16(m.p)}
}

// Mean calculates the mean value M of array elements, independently for each channel, and return it as Scalar
// TODO pass second paramter with mask
func (m *Mat) Mean() Scalar {
	s := C.Mat_Mean(m.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// LUT performs a look-up table transform of an array.
//
// The function LUT fills the output array with values from the look-up table.
// Indices of the entries are taken from the input array.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gab55b8d062b7f5587720ede032d34156f
func LUT(src, wbLUT, dst Mat) {
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
func (m *Mat) Type() int {
	return int(C.Mat_Type(m.p))
}

// GetUCharAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) GetUCharAt(row int, col int, channel ...interface{}) uint8 {
	if len(channel) > 0 {
		return uint8(C.Mat_GetUCharChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return uint8(C.Mat_GetUChar(m.p, C.int(row), C.int(col)))
}

// GetSCharAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) GetSCharAt(row int, col int, channel ...interface{}) int8 {
	if len(channel) > 0 {
		return int8(C.Mat_GetSCharChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return int8(C.Mat_GetSChar(m.p, C.int(row), C.int(col)))
}

// GetShortAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) GetShortAt(row int, col int, channel ...interface{}) int16 {
	if len(channel) > 0 {
		return int16(C.Mat_GetShortChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return int16(C.Mat_GetShort(m.p, C.int(row), C.int(col)))
}

// GetIntAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) GetIntAt(row int, col int, channel ...interface{}) int32 {
	if len(channel) > 0 {
		return int32(C.Mat_GetIntChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return int32(C.Mat_GetInt(m.p, C.int(row), C.int(col)))
}

// GetFloatAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) GetFloatAt(row int, col int, channel ...interface{}) float32 {
	if len(channel) > 0 {
		return float32(C.Mat_GetFloatChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return float32(C.Mat_GetFloat(m.p, C.int(row), C.int(col)))
}

// GetDoubleAt returns a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt(row int, col int, channel ...interface{}) float64 {
	if len(channel) > 0 {
		return float64(C.Mat_GetDoubleChannel(m.p, C.int(row), C.int(col), C.int(channel[0].(int))))
	}
	return float64(C.Mat_GetDouble(m.p, C.int(row), C.int(col)))
}

// SetUCharAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type uchar aka CV_8U.
func (m *Mat) SetUCharAt(row int, col int, args ...interface{}) {
	if len(args) > 1 {
		val := uint8(args[1].(int))
		C.Mat_SetUCharChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.uint8_t(val))
		return
	}
	val := uint8(args[0].(int))
	C.Mat_SetUChar(m.p, C.int(row), C.int(col), C.uint8_t(val))
}

// SetSCharAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type schar aka CV_8S.
func (m *Mat) SetSCharAt(row int, col int, args ...interface{}) {
	if len(args) > 1 {
		val := int8(args[1].(int))
		C.Mat_SetSCharChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.int8_t(val))
		return
	}
	val := uint8(args[0].(int))
	C.Mat_SetSChar(m.p, C.int(row), C.int(col), C.int8_t(val))
}

// SetShortAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type short aka CV_16S.
func (m *Mat) SetShortAt(row int, col int, args ...interface{}) {
	if len(args) > 1 {
		val := int16(args[1].(int))
		C.Mat_SetShortChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.int16_t(val))
		return
	}
	val := int16(args[0].(int))
	C.Mat_SetShort(m.p, C.int(row), C.int(col), C.int16_t(val))
}

// SetIntAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type int aka CV_32S.
func (m *Mat) SetIntAt(row int, col int, args ...interface{}) {
	if len(args) > 1 {
		val := int32(args[1].(int))
		C.Mat_SetIntChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.int32_t(val))
		return
	}
	val := int32(args[0].(int))
	C.Mat_SetInt(m.p, C.int(row), C.int(col), C.int32_t(val))
}

// SetFloatAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type float aka CV_32F.
func (m *Mat) SetFloatAt(row int, col int, args ...interface{}) {
	var val float32
	if len(args) > 1 {
		val = toFloat32(args[1])
		C.Mat_SetFloatChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.float(val))
		return
	}
	val = toFloat32(args[0])
	C.Mat_SetFloat(m.p, C.int(row), C.int(col), C.float(val))
}

// SetDoubleAt sets a value from a specific row/col (and optionally channel)
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt(row int, col int, args ...interface{}) {
	var val float64
	if len(args) > 1 {
		val = toFloat64(args[1])
		C.Mat_SetDoubleChannel(m.p, C.int(row), C.int(col), C.int(args[0].(int)), C.double(val))
		return
	}
	val = toFloat64(args[0])
	C.Mat_SetDouble(m.p, C.int(row), C.int(col), C.double(val))
}

// AbsDiff calculates the per-element absolute difference between two arrays
// or between an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga6fef31bc8c4071cbc114a758a2b79c14
//
func AbsDiff(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_AbsDiff(src1.p, src2.p, dst.p)
}

// Add calculates the per-element sum of two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga10ac1bfb180e2cfda1701d06c24fdbd6
//
func Add(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_Add(src1.p, src2.p, dst.p)
}

// AddWeighted calculates the weighted sum of two arrays.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gafafb2513349db3bcff51f54ee5592a19
//
func AddWeighted(src1 Mat, alpha float64, src2 Mat, beta float64, gamma float64, dst Mat) {
	C.Mat_AddWeighted(src1.p, C.double(alpha),
		src2.p, C.double(beta), C.double(gamma), dst.p)
}

// BitwiseAnd computes bitwise conjunction of the two arrays (dst = src1 & src2).
// Calculates the per-element bit-wise conjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga60b4d04b251ba5eb1392c34425497e14
//
func BitwiseAnd(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseAnd(src1.p, src2.p, dst.p)
}

// BitwiseNot inverts every bit of an array.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga0002cf8b418479f4cb49a75442baee2f
//
func BitwiseNot(src1 Mat, dst Mat) {
	C.Mat_BitwiseNot(src1.p, dst.p)
}

// BitwiseOr calculates the per-element bit-wise disjunction of two arrays
// or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gab85523db362a4e26ff0c703793a719b4
//
func BitwiseOr(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseOr(src1.p, src2.p, dst.p)
}

// BitwiseXor calculates the per-element bit-wise "exclusive or" operation
// on two arrays or an array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga84b2d8188ce506593dcc3f8cd00e8e2c
//
func BitwiseXor(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_BitwiseXor(src1.p, src2.p, dst.p)
}

// InRange checks if array elements lie between the elements of two Mat arrays.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga48af0ab51e36436c5d04340e036ce981
//
func InRange(src Mat, lb Mat, ub Mat, dst Mat) {
	C.Mat_InRange(src.p, lb.p, ub.p, dst.p)
}

// GetOptimalDFTSize returns the optimal Discrete Fourier Transform (DFT) size
// for a given vector size.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga6577a2e59968936ae02eb2edde5de299
//
func GetOptimalDFTSize(vecsize int) int {
	return int(C.Mat_GetOptimalDFTSize(C.int(vecsize)))
}

// DFT performs a forward or inverse Discrete Fourier Transform (DFT)
// of a 1D or 2D floating-point array.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gadd6cf9baf2b8b704a11b5f04aaf4f39d
//
func DFT(src Mat, dst Mat) {
	C.Mat_DFT(src.p, dst.p)
}

// Merge creates one multi-channel array out of several single-channel ones.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga7d7b4d6c6ee504b30a20b1680029c7b4
//
func Merge(mv []Mat, dst Mat) {
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

// Subtract Calculates the per-element difference between two arrays or array and a scalar.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gaa0f00d98b4b5edeaeb7b8333b2de353b
//
func Subtract(src1 Mat, src2 Mat, dst Mat) {
	C.Mat_Subtract(src1.p, src2.p, dst.p)
}

// NormType for normalization operations.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#gad12cefbcb5291cf958a85b4b67b6149f
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

	// NormMixMax indicates use min/max normalization.
	NormMixMax = 32
)

// Normalize normalizes the norm or value range of an array.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga87eef7ee3970f86906d69a92cbf064bd
//
func Normalize(src Mat, dst Mat, alpha float64, beta float64, typ NormType) {
	C.Mat_Normalize(src.p, dst.p, C.double(alpha), C.double(beta), C.int(typ))
}

// Norm calculates the absolute norm of an array.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33
//
func Norm(src1 Mat, normType NormType) float64 {
	return float64(C.Norm(src1.p, C.int(normType)))
}

// TermCriteriaType for TermCriteria.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d9/d5d/classcv_1_1TermCriteria.html#a56fecdc291ccaba8aad27d67ccf72c57
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

// TermCriteria is the criteria for iterative algorithms.
//
// For further details, please see:
// https://docs.opencv.org/3.4.0/d9/d5d/classcv_1_1TermCriteria.html
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
// http://docs.opencv.org/3.4.0/d1/da0/classcv_1_1Scalar__.html
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
// https://docs.opencv.org/3.4.0/d2/d29/classcv_1_1KeyPoint.html
//
type KeyPoint struct {
	X, Y                  float64
	Size, Angle, Response float64
	Octave, ClassID       int
}

func toByteArray(b []byte) C.struct_ByteArray {
	return C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}
}

func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}

func toFloat32(arg interface{}) float32 {
	switch arg.(type) {
	case int:
		return float32(arg.(int))
	case float64:
		return float32(arg.(float64))
	default:
		return 0
	}
}

func toFloat64(arg interface{}) float64 {
	switch arg.(type) {
	case int:
		return float64(arg.(int))
	case float64:
		return float64(arg.(float64))
	default:
		return 0
	}
}
