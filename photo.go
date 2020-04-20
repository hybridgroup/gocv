package gocv

/*
#include <stdlib.h>
#include "photo.h"
*/
import "C"

func FastNlMeansDenoising(src Mat, dst *Mat) {
	C.FastNlMeansDenoising(src.p, dst.p)
}

func FastNlMeansDenoisingWithParams(src Mat, dst *Mat, h float32, templateWindowSize int, searchWindowSize int) {
	C.FastNlMeansDenoisingWithParams(src.p, dst.p, C.float(h), C.int(templateWindowSize), C.int(searchWindowSize))
}

func FastNlMeansDenoisingColored(src Mat, dst *Mat) {
	C.FastNlMeansDenoisingColored(src.p, dst.p)
}

func FastNlMeansDenoisingColoredWithParams(src Mat, dst *Mat, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
	C.FastNlMeansDenoisingColoredWithParams(src.p, dst.p, C.float(h), C.float(hColor), C.int(templateWindowSize), C.int(searchWindowSize))
}
