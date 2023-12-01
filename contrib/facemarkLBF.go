package contrib

/*
#include <stdlib.h>
#include "facemarkLBF.h"
#include "../core.h"
*/
import "C"
import (
	"gocv.io/x/gocv"
	"image"
	"unsafe"
)

type LBPHFaceMark struct {
	p C.LBPHFaceMark
}

func NewLBPHFaceMark() *LBPHFaceMark {
	return &LBPHFaceMark{p: C.CreateLBPHFaceMark()}
}

func (mark *LBPHFaceMark) LoadModel(model string) {
	cName := C.CString(model)
	defer C.free(unsafe.Pointer(cName))
	C.LBPHFaceMark_LoadModel(mark.p, cName)
}

// LBPHFaceMark_Fit(LBPHFaceMark fm, Mat frame, struct Rects faces, Points2fVector landmarks)
func (mark *LBPHFaceMark) Fit(img gocv.Mat, faceBox []image.Rectangle) (bool, [][]gocv.Point2f) {

	bboxesRectArr := []C.struct_Rect{}
	for _, v := range faceBox {
		bbox := C.struct_Rect{
			x:      C.int(v.Min.X),
			y:      C.int(v.Min.Y),
			width:  C.int(v.Size().X),
			height: C.int(v.Size().Y),
		}
		bboxesRectArr = append(bboxesRectArr, bbox)
	}

	bboxesRects := C.Rects{
		rects:  (*C.Rect)(&bboxesRectArr[0]),
		length: C.int(len(faceBox)),
	}

	points2fVector := gocv.NewPoints2fVector()
	result := bool(C.LBPHFaceMark_Fit(
		mark.p,
		C.Mat(img.Ptr()),
		bboxesRects,
		C.Points2fVector(points2fVector.P()),
	))

	return result, points2fVector.ToPoints()
}
