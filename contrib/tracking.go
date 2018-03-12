package contrib

/*
#include "tracking.h"
*/
import "C"
import (
	"image"

	"gocv.io/x/gocv"
)

// This is the base interface for object tracking.
//
// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html
//
type Tracker interface {
	// Trackers need to be Closed manually.
	//
	Close() error

	// Init initializes the tracker with a known bounding box that surrounded the target.
	// Note: this can only be called once. If you lose the object, you have to Close() the instance,
	// create a new one, and call Init() on it again.
	//
	// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a4d285747589b1bdd16d2e4f00c3255dc
	//
	Init(image gocv.Mat, boundingBox image.Rectangle) bool

	// Update updates the tracker, returns a new bounding box and a boolean determining whether the tracker lost the target.
	//
	// see: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a549159bd0553e6a8de356f3866df1f18
	//
	Update(image gocv.Mat) (image.Rectangle, bool)
}

func tracker_Init(trk C.Tracker, img gocv.Mat, boundingBox image.Rectangle) bool {
	cBox := C.struct_Rect{
		x:      C.int(boundingBox.Min.X),
		y:      C.int(boundingBox.Min.Y),
		width:  C.int(boundingBox.Size().X),
		height: C.int(boundingBox.Size().Y),
	}

	ret := C.Tracker_Init(trk, C.Mat(img.Ptr()), cBox)
	return bool(ret)
}

func tracker_Update(trk C.Tracker, img gocv.Mat) (image.Rectangle, bool) {
	cBox := C.struct_Rect{}

	ret := C.Tracker_Update(trk, C.Mat(img.Ptr()), &cBox)

	rect := image.Rect(int(cBox.x), int(cBox.y), int(cBox.x+cBox.width), int(cBox.y+cBox.height))
	return rect, bool(ret)
}

//
// The MIL algorithm trains a classifier in an online manner to separate the object from the background.
// Multiple Instance Learning avoids the drift problem for a robust tracking.
//
// see: https://docs.opencv.org/master/d0/d26/classcv_1_1TrackerMIL.html
//
type TrackerMIL struct {
	p C.TrackerMIL
}

func NewTrackerMIL() TrackerMIL {
	return TrackerMIL{p: C.TrackerMIL_Create()}
}

func (self TrackerMIL) Close() error {
	C.TrackerMIL_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerMIL) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerMIL) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// This is a real-time object tracker based on a novel on-line version of the AdaBoost algorithm.
//
// see: https://docs.opencv.org/master/d1/d1a/classcv_1_1TrackerBoosting.html
//
type TrackerBoosting struct {
	p C.TrackerBoosting
}

func NewTrackerBoosting() TrackerBoosting {
	return TrackerBoosting{p: C.TrackerBoosting_Create()}
}

func (self TrackerBoosting) Close() error {
	C.TrackerBoosting_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerBoosting) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerBoosting) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// This Tracker implementation is suitable for very smooth and predictable movements when the object is visible throughout the whole sequence
//
// see: https://docs.opencv.org/master/d7/d86/classcv_1_1TrackerMedianFlow.html
//
type TrackerMedianFlow struct {
	p C.TrackerMedianFlow
}

func NewTrackerMedianFlow() TrackerMedianFlow {
	return TrackerMedianFlow{p: C.TrackerMedianFlow_Create()}
}

func (self TrackerMedianFlow) Close() error {
	C.TrackerMedianFlow_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerMedianFlow) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerMedianFlow) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// This is a novel tracking framework that explicitly decomposes the long-term tracking task into tracking, learning and detection.
//
// see: https://docs.opencv.org/master/dc/d1c/classcv_1_1TrackerTLD.html
//
type TrackerTLD struct {
	p C.TrackerTLD
}

func NewTrackerTLD() TrackerTLD {
	return TrackerTLD{p: C.TrackerTLD_Create()}
}

func (self TrackerTLD) Close() error {
	C.TrackerTLD_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerTLD) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerTLD) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// KCF is a novel tracking framework that utilizes properties of circulant matrix to enhance the processing speed.
//
// see: https://docs.opencv.org/master/d2/dff/classcv_1_1TrackerKCF.html
//
type TrackerKCF struct {
	p C.TrackerKCF
}

func NewTrackerKCF() TrackerKCF {
	return TrackerKCF{p: C.TrackerKCF_Create()}
}

func (self TrackerKCF) Close() error {
	C.TrackerKCF_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerKCF) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerKCF) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// Based on: Visual Object Tracking using Adaptive Correlation Filters.
// Note, that this tracker is working on graysccale images.
//
// see: https://docs.opencv.org/master/d0/d02/classcv_1_1TrackerMOSSE.html
//
type TrackerMOSSE struct {
	p C.TrackerMOSSE
}

func NewTrackerMOSSE() TrackerMOSSE {
	return TrackerMOSSE{p: C.TrackerMOSSE_Create()}
}

func (self TrackerMOSSE) Close() error {
	C.TrackerMOSSE_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerMOSSE) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerMOSSE) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// An implementation of:
// Discriminative Correlation Filter Tracker with Channel and Spatial Reliability.
//
// see: https://docs.opencv.org/master/d2/da2/classcv_1_1TrackerCSRT.html
//
type TrackerCSRT struct {
	p C.TrackerCSRT
}

func NewTrackerCSRT() TrackerCSRT {
	return TrackerCSRT{p: C.TrackerCSRT_Create()}
}

func (self TrackerCSRT) Close() error {
	C.TrackerCSRT_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerCSRT) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerCSRT) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}
