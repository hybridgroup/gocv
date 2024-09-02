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

type IFaceRecognizer interface {
	Empty() bool
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

type IBasicFaceRecognizer interface {
	GetEigenValues() gocv.Mat
	GetEigenVectors() gocv.Mat
	GetLabels() gocv.Mat
	GetMean() gocv.Mat
	GetNumComponents() int
	SetNumComponents(val int)
	GetProjections() []gocv.Mat
	SaveFile(fname string)
	LoadFile(fname string)
}

type FaceRecognizer struct {
	p unsafe.Pointer
}

type BasicFaceRecognizer struct {
	FaceRecognizer
}

func faceRecognizer_Empty(fr C.FaceRecognizer) bool {
	b := C.FaceRecognizer_Empty(fr)
	return bool(b)
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

func basicFaceRecognizer_Train(fr C.BasicFaceRecognizer, images []gocv.Mat, labels []int) {
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

	C.BasicFaceRecognizer_Train(fr, matsVector, labelsVector)
}

func basicFaceRecognizer_Update(fr C.BasicFaceRecognizer, newImages []gocv.Mat, newLabels []int) {
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

	C.BasicFaceRecognizer_Update(fr, matsVector, labelsVector)
}

func basicFaceRecognizer_GetEigenValues(fr C.BasicFaceRecognizer) gocv.Mat {
	c_mat := C.BasicFaceRecognizer_getEigenValues(fr)
	return gocv.NewMatFromCMat(unsafe.Pointer(c_mat))
}

func basicFaceRecognizer_GetEigenVectors(fr C.BasicFaceRecognizer) gocv.Mat {
	c_mat := C.BasicFaceRecognizer_getEigenVectors(fr)
	return gocv.NewMatFromCMat(unsafe.Pointer(c_mat))
}

func basicFaceRecognizer_GetLabels(fr C.BasicFaceRecognizer) gocv.Mat {
	c_mat := C.BasicFaceRecognizer_getLabels(fr)
	return gocv.NewMatFromCMat(unsafe.Pointer(c_mat))
}

func basicFaceRecognizer_GetMean(fr C.BasicFaceRecognizer) gocv.Mat {
	c_mat := C.BasicFaceRecognizer_getMean(fr)
	return gocv.NewMatFromCMat(unsafe.Pointer(c_mat))
}

func basicFaceRecognizer_GetNumComponents(fr C.BasicFaceRecognizer) int {
	i := C.BasicFaceRecognizer_getNumComponents(fr)
	return int(i)
}

func basicFaceRecognizer_SetNumComponents(fr C.BasicFaceRecognizer, val int) {
	C.BasicFaceRecognizer_setNumComponents(fr, C.int(val))
}

func basicFaceRecognizer_GetProjections(fr C.BasicFaceRecognizer) []gocv.Mat {

	c_mats := C.BasicFaceRecognizer_getProjections(fr)
	defer C.Mats_Close(c_mats)

	mats := make([]gocv.Mat, 0, c_mats.length)

	for i := 0; i < int(c_mats.length); i++ {
		c_mat := C.Mats_get(c_mats, C.int(i))
		mat := gocv.NewMatFromCMat(unsafe.Pointer(c_mat))
		mats = append(mats, mat)
	}

	return mats
}

func basicFaceRecognizer_SaveFile(fr C.BasicFaceRecognizer, fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.BasicFaceRecognizer_SaveFile(fr, cName)
}

func basicFaceRecognizer_LoadFile(fr C.BasicFaceRecognizer, fname string) {
	cName := C.CString(fname)
	defer C.free(unsafe.Pointer(cName))
	C.BasicFaceRecognizer_LoadFile(fr, cName)
}
