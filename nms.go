package gocv

/*
#include <stdlib.h>
#include "nms.h"
*/
import "C"
import (
	"image"
	"reflect"
	"unsafe"
)

// NMSBoxes performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
func NMSBoxes(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int) {
	bboxesRectArr := []C.struct_Rect{}
	for _, v := range bboxes {
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
		length: C.int(len(bboxes)),
	}

	scoresFloats := []C.float{}
	for _, v := range scores {
		scoresFloats = append(scoresFloats, C.float(v))
	}
	scoresVector := C.struct_FloatVector{}
	scoresVector.val = (*C.float)(&scoresFloats[0])
	scoresVector.length = (C.int)(len(scoresFloats))

	indicesVector := C.IntVector{}

	C.NMSBoxes(bboxesRects, scoresVector, C.float(scoreThreshold), C.float(nmsThreshold), &indicesVector)
	defer C.free(unsafe.Pointer(indicesVector.val))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(indicesVector.val)),
		Len:  int(indicesVector.length),
		Cap:  int(indicesVector.length),
	}

	ptr := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(indicesVector.length); i++ {
		indices[i] = int(ptr[i])
	}
	return
}

// NMSBoxesWithParams performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
func NMSBoxesWithParams(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int, eta float32, topK int) {
	bboxesRectArr := []C.struct_Rect{}
	for _, v := range bboxes {
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
		length: C.int(len(bboxes)),
	}

	scoresFloats := []C.float{}
	for _, v := range scores {
		scoresFloats = append(scoresFloats, C.float(v))
	}
	scoresVector := C.struct_FloatVector{}
	scoresVector.val = (*C.float)(&scoresFloats[0])
	scoresVector.length = (C.int)(len(scoresFloats))

	indicesVector := C.IntVector{}

	C.NMSBoxesWithParams(bboxesRects, scoresVector, C.float(scoreThreshold), C.float(nmsThreshold), &indicesVector, C.float(eta), C.int(topK))
	defer C.free(unsafe.Pointer(indicesVector.val))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(indicesVector.val)),
		Len:  int(indicesVector.length),
		Cap:  int(indicesVector.length),
	}

	ptr := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(indicesVector.length); i++ {
		indices[i] = int(ptr[i])
	}
	return
}
