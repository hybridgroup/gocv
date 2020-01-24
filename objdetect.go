package gocv

/*
#include <stdlib.h>
#include "objdetect.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

// CascadeClassifier is a cascade classifier class for object detection.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html
//
type CascadeClassifier struct {
	p C.CascadeClassifier
}

// NewCascadeClassifier returns a new CascadeClassifier.
func NewCascadeClassifier() CascadeClassifier {
	return CascadeClassifier{p: C.CascadeClassifier_New()}
}

// Close deletes the CascadeClassifier's pointer.
func (c *CascadeClassifier) Close() error {
	C.CascadeClassifier_Close(c.p)
	c.p = nil
	return nil
}

// Load cascade classifier from a file.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#a1a5884c8cc749422f9eb77c2471958bc
//
func (c *CascadeClassifier) Load(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.CascadeClassifier_Load(c.p, cName) != 0
}

// DetectMultiScale detects objects of different sizes in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
//
func (c *CascadeClassifier) DetectMultiScale(img Mat) []image.Rectangle {
	ret := C.CascadeClassifier_DetectMultiScale(c.p, img.p)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// DetectMultiScaleWithParams calls DetectMultiScale but allows setting parameters
// to values other than just the defaults.
//
// For further details, please see:
// http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498
//
func (c *CascadeClassifier) DetectMultiScaleWithParams(img Mat, scale float64,
	minNeighbors, flags int, minSize, maxSize image.Point) []image.Rectangle {

	minSz := C.struct_Size{
		width:  C.int(minSize.X),
		height: C.int(minSize.Y),
	}

	maxSz := C.struct_Size{
		width:  C.int(maxSize.X),
		height: C.int(maxSize.Y),
	}

	ret := C.CascadeClassifier_DetectMultiScaleWithParams(c.p, img.p, C.double(scale),
		C.int(minNeighbors), C.int(flags), minSz, maxSz)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// HOGDescriptor is a Histogram Of Gradiants (HOG) for object detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a723b95b709cfd3f95cf9e616de988fc8
//
type HOGDescriptor struct {
	p C.HOGDescriptor
}

// NewHOGDescriptor returns a new HOGDescriptor.
func NewHOGDescriptor() HOGDescriptor {
	return HOGDescriptor{p: C.HOGDescriptor_New()}
}

// Close deletes the HOGDescriptor's pointer.
func (h *HOGDescriptor) Close() error {
	C.HOGDescriptor_Close(h.p)
	h.p = nil
	return nil
}

// DetectMultiScale detects objects in the input Mat image.
// The detected objects are returned as a slice of image.Rectangle structs.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
//
func (h *HOGDescriptor) DetectMultiScale(img Mat) []image.Rectangle {
	ret := C.HOGDescriptor_DetectMultiScale(h.p, img.p)
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// DetectMultiScaleWithParams calls DetectMultiScale but allows setting parameters
// to values other than just the defaults.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
//
func (h *HOGDescriptor) DetectMultiScaleWithParams(img Mat, hitThresh float64,
	winStride, padding image.Point, scale, finalThreshold float64, useMeanshiftGrouping bool) []image.Rectangle {
	wSz := C.struct_Size{
		width:  C.int(winStride.X),
		height: C.int(winStride.Y),
	}

	pSz := C.struct_Size{
		width:  C.int(padding.X),
		height: C.int(padding.Y),
	}

	ret := C.HOGDescriptor_DetectMultiScaleWithParams(h.p, img.p, C.double(hitThresh),
		wSz, pSz, C.double(scale), C.double(finalThreshold), C.bool(useMeanshiftGrouping))
	defer C.Rects_Close(ret)

	return toRectangles(ret)
}

// HOGDefaultPeopleDetector returns a new Mat with the HOG DefaultPeopleDetector.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948
//
func HOGDefaultPeopleDetector() Mat {
	return newMat(C.HOG_GetDefaultPeopleDetector())
}

// SetSVMDetector sets the data for the HOGDescriptor.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a09e354ad701f56f9c550dc0385dc36f1
//
func (h *HOGDescriptor) SetSVMDetector(det Mat) error {
	C.HOGDescriptor_SetSVMDetector(h.p, det.p)
	return nil
}

// GroupRectangles groups the object candidate rectangles.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d54/group__objdetect.html#ga3dba897ade8aa8227edda66508e16ab9
//
func GroupRectangles(rects []image.Rectangle, groupThreshold int, eps float64) []image.Rectangle {
	cRectArray := make([]C.struct_Rect, len(rects))
	for i, r := range rects {
		cRect := C.struct_Rect{
			x:      C.int(r.Min.X),
			y:      C.int(r.Min.Y),
			width:  C.int(r.Size().X),
			height: C.int(r.Size().Y),
		}
		cRectArray[i] = cRect
	}
	cRects := C.struct_Rects{
		rects:  (*C.Rect)(&cRectArray[0]),
		length: C.int(len(rects)),
	}

	ret := C.GroupRectangles(cRects, C.int(groupThreshold), C.double(eps))

	return toRectangles(ret)
}

// QRCodeDetector groups the object candidate rectangles.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html
//
type QRCodeDetector struct {
	p C.QRCodeDetector
}

// newQRCodeDetector returns a new QRCodeDetector from a C QRCodeDetector
func newQRCodeDetector(p C.QRCodeDetector) QRCodeDetector {
	return QRCodeDetector{p: p}
}

func NewQRCodeDetector() QRCodeDetector {
	return newQRCodeDetector(C.QRCodeDetector_New())
}

func (a *QRCodeDetector) Close() error {
	C.QRCodeDetector_Close(a.p)
	a.p = nil
	return nil
}

// DetectAndDecode Both detects and decodes QR code.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a7290bd6a5d59b14a37979c3a14fbf394
//
func (a *QRCodeDetector) DetectAndDecode(input Mat, points *Mat, straight_qrcode *Mat) string {
	goResult := C.GoString(C.QRCodeDetector_DetectAndDecode(a.p, input.p, points.p, straight_qrcode.p))
	return string(goResult)
}

// Detect detects QR code in image and returns the quadrangle containing the code.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a64373f7d877d27473f64fe04bb57d22b
//
func (a *QRCodeDetector) Detect(input Mat, points *Mat) bool {
	result := C.QRCodeDetector_Detect(a.p, input.p, points.p)
	return bool(result)
}

// Decode decodes QR code in image once it's found by the detect() method. Returns UTF8-encoded output string or empty string if the code cannot be decoded.
//
// For further details, please see:
// https://docs.opencv.org/master/de/dc3/classcv_1_1QRCodeDetector.html#a4172c2eb4825c844fb1b0ae67202d329
//
func (a *QRCodeDetector) Decode(input Mat, points Mat, straight_qrcode *Mat) string {
	goResult := C.GoString(C.QRCodeDetector_DetectAndDecode(a.p, input.p, points.p, straight_qrcode.p))
	return string(goResult)
}
