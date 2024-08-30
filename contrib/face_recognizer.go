package contrib

/*
#include <stdlib.h>
#include "face.h"
*/
import "C"
import (
	"unsafe"

	"gocv.io/x/gocv"
)

type FaceRecognizer interface {
	Train(images []gocv.Mat, labels []int)
	Update(newImages []gocv.Mat, newLabels []int)
	Predict(sample gocv.Mat) int
	PredictExtendedResponse(sample gocv.Mat) PredictResponse
	GetThreshold() float32
	SetThreshold(threshold float32)
	SaveFile(fname string)
	LoadFile(fname string)
	Close() error
}

func faceRecognizer_Train(fr C.FaceRecognizer, images []gocv.Mat, labels []int) {
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

	C.FaceRecognizer_Train(fr, matsVector, labelsVector)
}

func faceRecognizer_Update(fr C.FaceRecognizer, newImages []gocv.Mat, newLabels []int) {
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

	C.FaceRecognizer_Update(fr, matsVector, labelsVector)
}

func faceRecognizer_Predict(fr C.FaceRecognizer, sample gocv.Mat) int {
	label := C.FaceRecognizer_Predict(fr, (C.Mat)(sample.Ptr()))

	return int(label)
}

func faceRecognizer_PredictExtendedResponse(fr C.FaceRecognizer, sample gocv.Mat) PredictResponse {
	respp := C.FaceRecognizer_PredictExtended(fr, (C.Mat)(sample.Ptr()))
	resp := PredictResponse{
		Label:      int32(respp.label),
		Confidence: float32(respp.confidence),
	}

	return resp
}

func faceRecognizer_GetThreshold(fr C.FaceRecognizer) float32 {
	t := C.FaceRecognizer_GetThreshold(fr)
	return float32(t)
}

func faceRecognizer_SetThreshold(fr C.FaceRecognizer, threshold float32) {
	C.FaceRecognizer_SetThreshold(fr, (C.double)(threshold))
}

func faceRecognizer_SaveFile(fr C.FaceRecognizer, fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.FaceRecognizer_SaveFile(fr, cName)
}

func faceRecognizer_LoadFile(fr C.FaceRecognizer, fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.FaceRecognizer_LoadFile(fr, cName)
}
