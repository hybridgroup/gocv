package cuda

/*
#include <stdlib.h>
#include "../core.h"
#include "core.h"
#include "arithm.h"
*/
import "C"

import "gocv.io/x/gocv"

// Abs computes an absolute value of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga54a72bd772494ab34d05406fd76df2b6
//
func Abs(src GpuMat, dst *GpuMat) {
	C.GpuAbs(src.p, dst.p, nil)
}

// AbsWithStream computes an absolute value of each matrix element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga54a72bd772494ab34d05406fd76df2b6
//
func AbsWithStream(src GpuMat, dst *GpuMat, stream Stream) {
	C.GpuAbs(src.p, dst.p, stream.p)
}

// AbsDiff computes per-element absolute difference of two matrices
// (or of a matrix and scalar) using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac062b283cf46ee90f74a773d3382ab54
//
func AbsDiff(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuAbsDiff(src1.p, src2.p, dst.p, nil)
}

// AbsDiffWithStream computes an absolute value of each matrix element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac062b283cf46ee90f74a773d3382ab54
//
func AbsDiffWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuAbsDiff(src1.p, src2.p, dst.p, s.p)
}

// Add computes a matrix-matrix or matrix-scalar sum.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5d9794bde97ed23d1c1485249074a8b1
//
func Add(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuAdd(src1.p, src2.p, dst.p, nil)
}

// AddWithStream computes a matrix-matrix or matrix-scalar sum
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5d9794bde97ed23d1c1485249074a8b1
//
func AddWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuAdd(src1.p, src2.p, dst.p, s.p)
}

// BitwiseAnd performs a per-element bitwise conjunction of two matrices
// (or of matrix and scalar).
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga78d7c1a013877abd4237fbfc4e13bd76
//
func BitwiseAnd(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuBitwiseAnd(src1.p, src2.p, dst.p, nil)
}

// BitwiseAndWithStream performs a per-element bitwise conjunction of two matrices
// (or of matrix and scalar) using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga78d7c1a013877abd4237fbfc4e13bd76
//
func BitwiseAndWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuBitwiseAnd(src1.p, src2.p, dst.p, s.p)
}

// BitwiseNot performs a per-element bitwise inversion.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gae58159a2259ae1acc76b531c171cf06a
//
func BitwiseNot(src GpuMat, dst *GpuMat) {
	C.GpuBitwiseNot(src.p, dst.p, nil)
}

// BitwiseNotWithStream performs a per-element bitwise inversion
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gae58159a2259ae1acc76b531c171cf06a
//
func BitwiseNotWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.GpuBitwiseNot(src.p, dst.p, s.p)
}

// BitwiseOr performs a per-element bitwise disjunction of two matrices
// (or of matrix and scalar).
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd098ee3e51c68daa793999c1da3dfb7
//
func BitwiseOr(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuBitwiseOr(src1.p, src2.p, dst.p, nil)
}

// BitwiseOrWithStream performs a per-element bitwise disjunction of two matrices
// (or of matrix and scalar) using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd098ee3e51c68daa793999c1da3dfb7
//
func BitwiseOrWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuBitwiseXor(src1.p, src2.p, dst.p, s.p)
}

// BitwiseXor performs a per-element exclusive or of two matrices
// (or of matrix and scalar).
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d95d4faafb099aacf18e8b915a4ad8d
//
func BitwiseXor(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuBitwiseXor(src1.p, src2.p, dst.p, nil)
}

// BitwiseXorWithStream performs a per-element exclusive or of two matrices
// (or of matrix and scalar) using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d95d4faafb099aacf18e8b915a4ad8d
//
func BitwiseXorWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuBitwiseXor(src1.p, src2.p, dst.p, s.p)
}

// Divide computes a matrix-matrix or matrix-scalar division.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga124315aa226260841e25cc0b9ea99dc3
//
func Divide(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuDivide(src1.p, src2.p, dst.p, nil)
}

// DivideWithStream computes a matrix-matrix or matrix-scalar division
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga124315aa226260841e25cc0b9ea99dc3
//
func DivideWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuDivide(src1.p, src2.p, dst.p, s.p)
}

// Exp computes an exponent of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61
//
func Exp(src GpuMat, dst *GpuMat) {
	C.GpuExp(src.p, dst.p, nil)
}

// ExpWithStream computes an exponent of each matrix element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61
//
func ExpWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.GpuExp(src.p, dst.p, s.p)
}

// Log computes natural logarithm of absolute value of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61
//
func Log(src GpuMat, dst *GpuMat) {
	C.GpuLog(src.p, dst.p, nil)
}

// LogWithStream computes natural logarithm of absolute value of each matrix element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61
//
func LogWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.GpuLog(src.p, dst.p, s.p)
}

// Max computes the per-element maximum of two matrices (or a matrix and a scalar).
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gadb5dd3d870f10c0866035755b929b1e7
//
func Max(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuMax(src1.p, src2.p, dst.p, nil)
}

// MaxWithStream computes the per-element maximum of two matrices (or a matrix and a scalar).
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gadb5dd3d870f10c0866035755b929b1e7
//
func MaxWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuMax(src1.p, src2.p, dst.p, s.p)
}

// Min computes the per-element minimum of two matrices (or a matrix and a scalar).
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga74f0b05a65b3d949c237abb5e6c60867
//
func Min(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuMin(src1.p, src2.p, dst.p, nil)
}

// MinWithStream computes the per-element minimum of two matrices (or a matrix and a scalar).
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga74f0b05a65b3d949c237abb5e6c60867
//
func MinWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuMin(src1.p, src2.p, dst.p, s.p)
}

// Multiply computes a matrix-matrix or matrix-scalar multiplication.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga124315aa226260841e25cc0b9ea99dc3
//
func Multiply(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuMultiply(src1.p, src2.p, dst.p, nil)
}

// Sqr computes a square value of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga8aae233da90ce0ffe309ab8004342acb
//
func Sqr(src GpuMat, dst *GpuMat) {
	C.GpuSqr(src.p, dst.p, nil)
}

// SqrWithStream computes a square value of each matrix element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga8aae233da90ce0ffe309ab8004342acb
//
func SqrWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.GpuSqr(src.p, dst.p, s.p)
}

// Sqrt computes a square root of each matrix element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga09303680cb1a5521a922b6d392028d8c
//
func Sqrt(src GpuMat, dst *GpuMat) {
	C.GpuSqrt(src.p, dst.p, nil)
}

// SqrtWithStream computes a square root of each matrix element.
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga09303680cb1a5521a922b6d392028d8c
//
func SqrtWithStream(src GpuMat, dst *GpuMat, s Stream) {
	C.GpuSqrt(src.p, dst.p, s.p)
}

// Subtract computes a matrix-matrix or matrix-scalar difference.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga6eab60fc250059e2fda79c5636bd067f
//
func Subtract(src1, src2 GpuMat, dst *GpuMat) {
	C.GpuSubtract(src1.p, src2.p, dst.p, nil)
}

// SubtractWithStream computes a matrix-matrix or matrix-scalar difference
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga6eab60fc250059e2fda79c5636bd067f
//
func SubtractWithStream(src1, src2 GpuMat, dst *GpuMat, s Stream) {
	C.GpuSubtract(src1.p, src2.p, dst.p, s.p)
}

// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga40f1c94ae9a9456df3cad48e3cb008e1
//
func Threshold(src GpuMat, dst *GpuMat, thresh, maxval float64, typ gocv.ThresholdType) {
	C.GpuThreshold(src.p, dst.p, C.double(thresh), C.double(maxval), C.int(typ), nil)
}

// ThresholdWithStream applies a fixed-level threshold to each array element
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga40f1c94ae9a9456df3cad48e3cb008e1
//
func ThresholdWithStream(src GpuMat, dst *GpuMat, thresh, maxval float64, typ gocv.ThresholdType, s Stream) {
	C.GpuThreshold(src.p, dst.p, C.double(thresh), C.double(maxval), C.int(typ), s.p)
}

// Flip flips a 2D matrix around vertical, horizontal, or both axes.
//
// For further details, please see:
// https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga4d0a3f2b46e8f0f1ec2b5ac178dcd871
//
func Flip(src GpuMat, dst *GpuMat, flipCode int) {
	C.GpuFlip(src.p, dst.p, C.int(flipCode), nil)
}

// FlipWithStream flips a 2D matrix around vertical, horizontal, or both axes
// using a Stream for concurrency.
//
// For further details, please see:
// https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga4d0a3f2b46e8f0f1ec2b5ac178dcd871
//
func FlipWithStream(src GpuMat, dst *GpuMat, flipCode int, stream Stream) {
	C.GpuFlip(src.p, dst.p, C.int(flipCode), stream.p)
}
