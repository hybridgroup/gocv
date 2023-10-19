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

// NewSURFWithParams returns a new SURF algorithm algorithm with parameters
//
// For further details, please see:
// https://docs.opencv.org/master/d5/df7/classcv_1_1xfeatures2d_1_1SURF.html#a436553ca44d9a2238761ddbee5b395e5
func NewSURFWithParams(hessianThreshold float64, nOctaves int, nOctaveLayers int, extended bool, upright bool) SURF {
	return SURF{p: unsafe.Pointer(C.SURF_CreateWithParams(C.double(hessianThreshold), C.int(nOctaves), C.int(nOctaveLayers), C.bool(extended), C.bool(upright)))}
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

// Compute keypoints in an image using SURF.
//
// For further details, please see:
// https://docs.opencv.org/3.4/d9/d37/classcv_1_1xfeatures2d_1_1DAISY.html#a12744f1611a374fb06ba251d9d2fec86
func (d *SURF) Compute(src gocv.Mat, mask gocv.Mat, kps []gocv.KeyPoint) ([]gocv.KeyPoint, gocv.Mat) {
	desc := gocv.NewMat()
	kp2arr := make([]C.struct_KeyPoint, len(kps))
	for i, kp := range kps {
		kp2arr[i].x = C.double(kp.X)
		kp2arr[i].y = C.double(kp.Y)
		kp2arr[i].size = C.double(kp.Size)
		kp2arr[i].angle = C.double(kp.Angle)
		kp2arr[i].response = C.double(kp.Response)
		kp2arr[i].octave = C.int(kp.Octave)
		kp2arr[i].classID = C.int(kp.ClassID)
	}
	cKeyPoints := C.struct_KeyPoints{
		keypoints: (*C.struct_KeyPoint)(&kp2arr[0]),
		length:    (C.int)(len(kps)),
	}

	ret := C.SURF_Compute((C.SURF)(d.p), C.Mat(src.Ptr()), cKeyPoints, C.Mat(desc.Ptr()))
	defer C.KeyPoints_Close(ret)

	return getKeyPoints(ret), desc
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

// BriefDescriptorExtractor is a wrapper around the cv::BriefDescriptorExtractor algorithm.
type BriefDescriptorExtractor struct {
	// C.BriefDescriptorExtractor
	p unsafe.Pointer
}

// NewBriefDescriptorExtractor returns a new BriefDescriptorExtractor algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d93/classcv_1_1xfeatures2d_1_1BriefDescriptorExtractor.html
func NewBriefDescriptorExtractor() BriefDescriptorExtractor {
	return BriefDescriptorExtractor{p: unsafe.Pointer(C.BriefDescriptorExtractor_Create())}
}

// NewBriefDescriptorExtractorWithParams returns a new BriefDescriptorExtractor algorithm algorithm with parameters
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d93/classcv_1_1xfeatures2d_1_1BriefDescriptorExtractor.html#ae3bc52666010fb137ab6f0d32de51f60
func NewBriefDescriptorExtractorWithParams(bytes int, useOrientation bool) BriefDescriptorExtractor {
	return BriefDescriptorExtractor{p: unsafe.Pointer(C.BriefDescriptorExtractor_CreateWithParams(C.int(bytes), C.bool(useOrientation)))}
}

// Close BriefDescriptorExtractor.
func (d *BriefDescriptorExtractor) Close() error {
	C.BriefDescriptorExtractor_Close((C.BriefDescriptorExtractor)(d.p))
	d.p = nil
	return nil
}

// Compute descriptors with given keypoints using BriefDescriptorExtractor
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0
func (b *BriefDescriptorExtractor) Compute(keyPoints []gocv.KeyPoint, src gocv.Mat) gocv.Mat {
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

	C.BriefDescriptorExtractor_Compute((C.BriefDescriptorExtractor)(b.p), C.Mat(src.Ptr()), cKeyPoints, C.Mat(desc.Ptr()))
	return desc
}
