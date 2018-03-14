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

type trackerMIL struct {
	p C.TrackerMIL
}

//
// The MIL algorithm trains a classifier in an online manner to separate the object from the background.
// Multiple Instance Learning avoids the drift problem for a robust tracking.
//
// see: https://docs.opencv.org/master/d0/d26/classcv_1_1TrackerMIL.html
//
func NewTrackerMIL() Tracker {
	return trackerMIL{p: C.TrackerMIL_Create()}
}

func (self trackerMIL) Close() error {
	C.TrackerMIL_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerMIL) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerMIL) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerBoosting struct {
	p C.TrackerBoosting
}

//
// This is a real-time object tracker based on a novel on-line version of the AdaBoost algorithm.
//
// see: https://docs.opencv.org/master/d1/d1a/classcv_1_1TrackerBoosting.html
//
func NewTrackerBoosting() Tracker {
	return trackerBoosting{p: C.TrackerBoosting_Create()}
}

func (self trackerBoosting) Close() error {
	C.TrackerBoosting_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerBoosting) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerBoosting) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerMedianFlow struct {
	p C.TrackerMedianFlow
}

//
// This Tracker implementation is suitable for very smooth and predictable movements when the object is visible throughout the whole sequence
//
// see: https://docs.opencv.org/master/d7/d86/classcv_1_1TrackerMedianFlow.html
//
func NewTrackerMedianFlow() Tracker {
	return trackerMedianFlow{p: C.TrackerMedianFlow_Create()}
}

func (self trackerMedianFlow) Close() error {
	C.TrackerMedianFlow_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerMedianFlow) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerMedianFlow) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerTLD struct {
	p C.TrackerTLD
}

//
// This is a novel tracking framework that explicitly decomposes the long-term tracking task into tracking, learning and detection.
//
// see: https://docs.opencv.org/master/dc/d1c/classcv_1_1TrackerTLD.html
//
func NewTrackerTLD() Tracker {
	return trackerTLD{p: C.TrackerTLD_Create()}
}

func (self trackerTLD) Close() error {
	C.TrackerTLD_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerTLD) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerTLD) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerKCF struct {
	p C.TrackerKCF
}

//
// KCF is a novel tracking framework that utilizes properties of circulant matrix to enhance the processing speed.
//
// see: https://docs.opencv.org/master/d2/dff/classcv_1_1TrackerKCF.html
//
func NewTrackerKCF() Tracker {
	return trackerKCF{p: C.TrackerKCF_Create()}
}

func (self trackerKCF) Close() error {
	C.TrackerKCF_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerKCF) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerKCF) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerMOSSE struct {
	p C.TrackerMOSSE
}

//
// Based on: Visual Object Tracking using Adaptive Correlation Filters.
// Note, that this tracker is working on graysccale images.
//
// see: https://docs.opencv.org/master/d0/d02/classcv_1_1TrackerMOSSE.html
//
func NewTrackerMOSSE() Tracker {
	return trackerMOSSE{p: C.TrackerMOSSE_Create()}
}

func (self trackerMOSSE) Close() error {
	C.TrackerMOSSE_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerMOSSE) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerMOSSE) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

type trackerCSRT struct {
	p C.TrackerCSRT
}

//
// An implementation of:
// Discriminative Correlation Filter Tracker with Channel and Spatial Reliability.
//
// see: https://docs.opencv.org/master/d2/da2/classcv_1_1TrackerCSRT.html
//
func NewTrackerCSRT() Tracker {
	return trackerCSRT{p: C.TrackerCSRT_Create()}
}

func (self trackerCSRT) Close() error {
	C.TrackerCSRT_Close(self.p)
	self.p = nil
	return nil
}

func (self trackerCSRT) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self trackerCSRT) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}
