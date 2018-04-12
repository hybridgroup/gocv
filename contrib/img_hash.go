package contrib

//#include <stdlib.h>
//#include "img_hash.h"
import "C"

import (
	"gocv.io/x/gocv"
)

// ImgHashBase defines the interface used for all of the img_hash algorithms.
type ImgHashBase interface {
	Compare(a, b gocv.Mat) float64
	Compute(inputArr gocv.Mat, outputArr *gocv.Mat)
}

// PHash is implementation of the PHash algorithm.
//
type PHash struct{}

// Compute computes hash of the input image using PHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash PHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.pHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

// Compare compares the hash value between a and b using PHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash PHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.pHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

// AverageHash is implementation of the AverageHash algorithm.
//
type AverageHash struct{}

// Compute computes hash of the input image using AverageHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash AverageHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.averageHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

// Compare compares the hash value between a and b using AverageHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash AverageHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.averageHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

// BlockMeanHash is implementation of the BlockMeanHash algorithm.
//
type BlockMeanHash struct {
	Mode BlockMeanHashMode
}

type BlockMeanHashMode int

const (
	BlockMeanHashMode0 BlockMeanHashMode = iota
	BlockMeanHashMode1
	BlockMeanHashModeDefault = BlockMeanHashMode0
)

// Compute computes hash of the input image using BlockMeanHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash BlockMeanHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.blockMeanHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()), C.int(hash.Mode))
}

// Compare compares the hash value between a and b using BlockMeanHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash BlockMeanHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.blockMeanHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()), C.int(hash.Mode)))
}

// TODO: BlockMeanHash.GetMean isn't implemented, because it requires state from the last
// call to Compute, and there's no easy way to keep it.

// ColorMomentHash is implementation of the ColorMomentHash algorithm.
//
type ColorMomentHash struct{}

// Compute computes hash of the input image using ColorMomentHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash ColorMomentHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.colorMomentHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

// Compare compares the hash value between a and b using ColorMomentHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash ColorMomentHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.colorMomentHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

// MarrHildrethHash is implementation of the MarrHildrethHash algorithm.
//
type MarrHildrethHash struct {
	Alpha float32
	Scale float32
}

func NewMarrHildrethHash() MarrHildrethHash {
	return MarrHildrethHash{2.0, 1.0}
}

// Compute computes hash of the input image using MarrHildrethHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash MarrHildrethHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.marrHildrethHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()),
		C.float(hash.Alpha), C.float(hash.Scale))
}

// Compare compares the hash value between a and b using MarrHildrethHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash MarrHildrethHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.marrHildrethHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()),
		C.float(hash.Alpha), C.float(hash.Scale)))
}

// RadialVarianceHash is implementation of the RadialVarianceHash algorithm.
//
type RadialVarianceHash struct {
	Sigma          float64
	NumOfAngleLine int
}

func NewRadialVarianceHash() RadialVarianceHash {
	return RadialVarianceHash{1, 180}
}

// Compute computes hash of the input image using RadialVarianceHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#ae2d9288db370089dfd8aab85d5e0b0f3
//
func (hash RadialVarianceHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.radialVarianceHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()),
		C.double(hash.Sigma), C.int(hash.NumOfAngleLine))
}

// Compare compares the hash value between a and b using RadialVarianceHash.
//
// For further information, see:
// https://docs.opencv.org/master/de/d29/classcv_1_1img__hash_1_1ImgHashBase.html#a444a3e9ec792cf029385809393f84ad5
//
func (hash RadialVarianceHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.radialVarianceHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()),
		C.double(hash.Sigma), C.int(hash.NumOfAngleLine)))
}

// TODO: RadialVariance getFeatures, getHash, getPixPerLine, getProjection are not
// implemented here, because they're stateful.
