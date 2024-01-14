package contrib

/*
#include <stdlib.h>
#include "xfeatures2d.h"
*/
import "C"

import (
	"reflect"
	"unsafe"

	"gocv.io/x/gocv"
)

// SURF is a wrapper around the cv::SURF algorithm.
// Due to being a patented algorithm you must set the OpenCV contrib build flag OPENCV_ENABLE_NONFREE=1
// in order to use it.
type SURF struct {
	// C.SURF
	p unsafe.Pointer
}

// NewSURF returns a new SURF algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/df7/classcv_1_1xfeatures2d_1_1SURF.html
func NewSURF() SURF {
	return SURF{p: unsafe.Pointer(C.SURF_Create())}
}

// Close SURF.
func (d *SURF) Close() error {
	C.SURF_Close((C.SURF)(d.p))
	d.p = nil
	return nil
}

// Detect keypoints in an image using SURF.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887
func (d *SURF) Detect(src gocv.Mat) []gocv.KeyPoint {
	ret := C.SURF_Detect((C.SURF)(d.p), C.Mat(src.Ptr()))

	return getKeyPoints(ret)
}

// DetectAndCompute detects and computes keypoints in an image using SURF.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677
func (d *SURF) DetectAndCompute(src gocv.Mat, mask gocv.Mat) ([]gocv.KeyPoint, gocv.Mat) {
	desc := gocv.NewMat()
	ret := C.SURF_DetectAndCompute((C.SURF)(d.p), C.Mat(src.Ptr()), C.Mat(mask.Ptr()),
		C.Mat(desc.Ptr()))

	return getKeyPoints(ret), desc
}

func getKeyPoints(ret C.KeyPoints) []gocv.KeyPoint {
	cArray := ret.keypoints
	defer C.free(unsafe.Pointer(cArray))
	length := int(ret.length)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cArray)),
		Len:  length,
		Cap:  length,
	}
	s := *(*[]C.KeyPoint)(unsafe.Pointer(&hdr))

	keys := make([]gocv.KeyPoint, length)
	for i, r := range s {
		keys[i] = gocv.KeyPoint{X: float64(r.x), Y: float64(r.y), Size: float64(r.size), Angle: float64(r.angle),
			Response: float64(r.response), Octave: int(r.octave), ClassID: int(r.classID)}
	}
	return keys
}

// BeblidDescriptorExtractor is a wrapper around the cv::BeblidDescriptorExtractor descriptor algorithm.
type BeblidDescriptorExtractor struct {
	// C.BeblidDescriptorExtractor
	p unsafe.Pointer
}

type BeblidDescriptorExtractorSize = int

const (
	BEBLID_SIZE_256_BITS BeblidDescriptorExtractorSize = 101
	BEBLID_SIZE_512_BITS BeblidDescriptorExtractorSize = 100
)

// NewBeblidDescriptorExtractor returns a new BEBLID descriptor algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/df7/classcv_1_1xfeatures2d_1_1SURF.html
func NewBeblidDescriptorExtractor(scaleFactor float32, size BeblidDescriptorExtractorSize) BeblidDescriptorExtractor {
	return BeblidDescriptorExtractor{p: unsafe.Pointer(C.BeblidDescriptorExtractor_Create(C.float(scaleFactor), C.int(size)))}
}

// Close BEBLID.
func (d *BeblidDescriptorExtractor) Close() error {
	C.BeblidDescriptorExtractor_Close((C.BeblidDescriptorExtractor)(d.p))
	d.p = nil
	return nil
}

// Detect describes keypoints in an image using BEBLID
//
// For further details, please see:
// https://docs.opencv.org/4.9.0/d7/d99/classcv_1_1xfeatures2d_1_1BEBLID.html
func (b *BeblidDescriptorExtractor) Compute(keyPoints []gocv.KeyPoint, src gocv.Mat) gocv.Mat {
	desc := gocv.NewMat()
	cKeyPointArray := make([]C.struct_KeyPoint, len(keyPoints))

	for i, kp := range keyPoints {
		cKeyPointArray[i].x = C.double(kp.X)
		cKeyPointArray[i].y = C.double(kp.Y)
		cKeyPointArray[i].size = C.double(kp.Size)
		cKeyPointArray[i].angle = C.double(kp.Angle)
		cKeyPointArray[i].response = C.double(kp.Response)
		cKeyPointArray[i].octave = C.int(kp.Octave)
		cKeyPointArray[i].classID = C.int(kp.ClassID)
	}

	cKeyPoints := C.struct_KeyPoints{
		keypoints: (*C.struct_KeyPoint)(&cKeyPointArray[0]),
		length:    (C.int)(len(keyPoints)),
	}

	C.BeblidDescriptorExtractor_Compute((C.BeblidDescriptorExtractor)(b.p), C.Mat(src.Ptr()), cKeyPoints, C.Mat(desc.Ptr()))
	return desc
}
