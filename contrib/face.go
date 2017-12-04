package contrib

/*
#cgo pkg-config: opencv
#include <stdlib.h>
#include "face.h"
*/
import "C"
import (
	"gocv.io/x/gocv"
	"unsafe"
)

type PredictResponse struct {
	Label int32 `json:"label"`
	Confidence float32 `json:"confidence"`
}

type LBPHFaceRecognizer struct {
	p C.LBPHFaceRecognizer
}

func NewLBPHFaceRecognizer() *LBPHFaceRecognizer {
	return &LBPHFaceRecognizer{p: C.CreateLBPHFaceRecognizer()}
}


func (fr *LBPHFaceRecognizer) Train(images []gocv.Mat,labels []int){
	cparams := []C.int{}
	for _, v := range labels {
		cparams = append(cparams, C.int(v))
	}
	labelsVector := C.struct_IntVector{}
	labelsVector.val = (*C.int)(&cparams[0])
	labelsVector.length = (C.int)(len(cparams))

	cMatArray := make([]C.Mat, len(images))
	for i, r := range images {
		cMatArray[i] = (C.Mat)(r.Ptr())
	}
	matsVector := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(images)),
	}

	C.LBPHFaceRecognizer_Train(fr.p, matsVector, labelsVector)
}

func (fr *LBPHFaceRecognizer) Update(newImages []gocv.Mat,newLabels []int){
	cparams := []C.int{}
	for _, v := range newLabels {
		cparams = append(cparams, C.int(v))
	}
	labelsVector := C.struct_IntVector{}
	labelsVector.val = (*C.int)(&cparams[0])
	labelsVector.length = (C.int)(len(cparams))

	cMatArray := make([]C.Mat, len(newImages))
	for i, r := range newImages {
		cMatArray[i] = (C.Mat)(r.Ptr())
	}
	matsVector := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(newImages)),
	}

	C.LBPHFaceRecognizer_Update(fr.p, matsVector, labelsVector)
}


func (fr *LBPHFaceRecognizer) Predict(sample gocv.Mat) (int){
	label := C.LBPHFaceRecognizer_Predict(fr.p, (C.Mat)(sample.Ptr()));

	return int(label);
}

func (fr *LBPHFaceRecognizer) PredictExtendedResponse(sample gocv.Mat) (PredictResponse){
	respp := C.LBPHFaceRecognizer_PredictExtended(fr.p, (C.Mat)(sample.Ptr()));
	resp := PredictResponse{
		Label: int32(respp.label),
		Confidence: float32(respp.confidence),
	}

	return resp;
}

func (fr *LBPHFaceRecognizer) SetThreshold(threshold float32) {
	C.LBPHFaceRecognizer_SetThreshold(fr.p, (C.double)(threshold));
}

func (fr *LBPHFaceRecognizer) SetNeighbors(neighbors int) {
	C.LBPHFaceRecognizer_SetNeighbors(fr.p, (C.int)(neighbors));
}

func (fr *LBPHFaceRecognizer) SetRadius(radius int) {
	C.LBPHFaceRecognizer_SetRadius(fr.p, (C.int)(radius));
}

func (fr *LBPHFaceRecognizer) SaveFile(fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.LBPHFaceRecognizer_SaveFile(fr.p, cName);
}

func (fr *LBPHFaceRecognizer) LoadFile(fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.LBPHFaceRecognizer_LoadFile(fr.p, cName);
}
