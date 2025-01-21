package contrib

/*
#include <stdlib.h>
#include "xobjdetect.h"
*/
import "C"
import (
	"image"
	"unsafe"

	"gocv.io/x/gocv"
)

// WBDetector is a wrapper around the cv::xobjdetect::WBDetector.
type WBDetector struct {
	p C.WBDetector
}

// NewWBDetector Creates a new WBDetector
//
// For further details, please see:
// https://docs.opencv.org/4.x/de/d0e/classcv_1_1xobjdetect_1_1WBDetector.html#a58377ae61694aac08ad842ac830972d9
func NewWBDetector() WBDetector {
	p := C.WBDetector_Create()
	return WBDetector{p: p}
}

// Close Releases WBDetector allocated resources.
func (det *WBDetector) Close() {
	C.WBDetector_Close(det.p)
}

// Detect Detect objects on image using WaldBoost detector
//
// img: Input image for detection
//
// Returns:
//
// Bounding boxes coordinates and
// Confidence values for bounding boxes output vector
//
// For further details, please see:
// https://docs.opencv.org/4.x/de/d0e/classcv_1_1xobjdetect_1_1WBDetector.html#ad19680e6545f49a9ca42dfc3457319e2
func (det *WBDetector) Detect(img gocv.Mat) ([]image.Rectangle, []float32) {

	result := C.WBDetector_Detect(det.p, C.Mat(img.Ptr()))

	defer C.free(unsafe.Pointer(result.bboxes.rects))
	defer C.free(unsafe.Pointer(result.confidences.val))

	cRects := unsafe.Slice(result.bboxes.rects, result.bboxes.length)
	cConfs := unsafe.Slice(result.confidences.val, result.confidences.length)

	goRects := make([]image.Rectangle, int(result.bboxes.length))
	goConfs := make([]float32, int(result.confidences.length))

	for i := 0; i < int(result.bboxes.length); i++ {
		r := image.Rect(int(cRects[i].x),
			int(cRects[i].y),
			int(cRects[i].width),
			int(cRects[i].height))

		goRects = append(goRects, r)
	}

	for i := 0; i < int(result.confidences.length); i++ {
		goConfs[i] = float32(cConfs[i])
	}

	return goRects, goConfs
}

// Read Read detector from gocv.FileNode
//
// For further details, please see:
// https://docs.opencv.org/4.x/de/d0e/classcv_1_1xobjdetect_1_1WBDetector.html#aef2df760f45d25aade518196986e139f
func (det *WBDetector) Read(node *gocv.FileNode) {
	C.WBDetector_Read(det.p, C.FileNode(node.Ptr()))
}

// Train WaldBoost detector.
//
// Parameters:
//
// posSamples: Path to directory with cropped positive samples
//
// negImgs: Path to directory with negative (background) images
//
// For further details, please see:
// https://docs.opencv.org/4.x/de/d0e/classcv_1_1xobjdetect_1_1WBDetector.html#a3720fb425a2d16f6cd0625a2d8bc563e
func (det *WBDetector) Train(posSamples string, negImgs string) {

	pos_samples := C.CString(posSamples)
	neg_imgs := C.CString(negImgs)
	defer C.free(unsafe.Pointer(pos_samples))
	defer C.free(unsafe.Pointer(neg_imgs))

	C.WBDetector_Train(det.p, pos_samples, neg_imgs)
}

// Write detector to gocv.FileStorage.
//
// For further details, please see:
// https://docs.opencv.org/4.x/de/d0e/classcv_1_1xobjdetect_1_1WBDetector.html#a7d85338895707904ae1ddb4374ec8dac
func (det *WBDetector) Write(fs *gocv.FileStorage) {
	C.WBDetector_Write(det.p, C.FileStorage(fs.Ptr()))
}
