package contrib

//#include <stdlib.h>
//#include "img_hash.h"
import "C"

import (
	"gocv.io/x/gocv"
)

type ImgHashBase interface {
	Compare(a, b gocv.Mat) float64
	Compute(inputArr gocv.Mat, outputArr *gocv.Mat)
}

type PHash struct{}

func (hash PHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.pHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

func (hash PHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.pHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

type AverageHash struct{}

func (hash AverageHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.averageHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

func (hash AverageHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.averageHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

type BlockMeanHash struct {
	Mode BlockMeanHashMode
}

type BlockMeanHashMode int

const (
	BlockMeanHashMode0 BlockMeanHashMode = iota
	BlockMeanHashMode1
	BlockMeanHashModeDefault = BlockMeanHashMode0
)

func (hash BlockMeanHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.blockMeanHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()), C.int(hash.Mode))
}

func (hash BlockMeanHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.blockMeanHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()), C.int(hash.Mode)))
}

// TODO: BlockMeanHash.GetMean isn't implemented, because it requires state from the last
// call to Compute, and there's no easy way to keep it.

type ColorMomentHash struct{}

func (hash ColorMomentHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.colorMomentHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()))
}

func (hash ColorMomentHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.colorMomentHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr())))
}

type MarrHildrethHash struct {
	Alpha float32
	Scale float32
}

func NewMarrHildrethHash() MarrHildrethHash {
	return MarrHildrethHash{2.0, 1.0}
}

func (hash MarrHildrethHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.marrHildrethHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()),
		C.float(hash.Alpha), C.float(hash.Scale))
}

func (hash MarrHildrethHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.marrHildrethHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()),
		C.float(hash.Alpha), C.float(hash.Scale)))
}

type RadialVarianceHash struct {
	Sigma          float64
	NumOfAngleLine int
}

func NewRadialVarianceHash() RadialVarianceHash {
	return RadialVarianceHash{1, 180}
}

func (hash RadialVarianceHash) Compute(input gocv.Mat, output *gocv.Mat) {
	C.radialVarianceHashCompute(C.Mat(input.Ptr()), C.Mat(output.Ptr()),
		C.double(hash.Sigma), C.int(hash.NumOfAngleLine))
}

func (hash RadialVarianceHash) Compare(a, b gocv.Mat) float64 {
	return float64(C.radialVarianceHashCompare(C.Mat(a.Ptr()), C.Mat(b.Ptr()),
		C.double(hash.Sigma), C.int(hash.NumOfAngleLine)))
}

// TODO: RadialVariance getFeatures, getHash, getPixPerLine, getProjection are not
// implemented here, because they're stateful.
