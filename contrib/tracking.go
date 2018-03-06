package contrib

/*
#include "tracking.h"
*/
import "C"
import (
	"image"

	"gocv.io/x/gocv"
)

// Tracker - the base interface for object tracking
//
// s.a: https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html
type Tracker interface {
	// Init initializes the tracker with a known bounding box that surrounded the target
	// note: this can only be called once. if you loose the object, you have to Close() the instance,
	// create a new one, and call Init on it again
	//
	// https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a4d285747589b1bdd16d2e4f00c3255dc
	Init(image gocv.Mat, boundingBox image.Rectangle) bool

	// Update the tracker, return new bounding box and a bool, if it lost it.
	//
	// https://docs.opencv.org/master/d0/d0a/classcv_1_1Tracker.html#a549159bd0553e6a8de356f3866df1f18
	Update(image gocv.Mat) (image.Rectangle, bool)
}

//
// available implementations in gocv are:
//
// TrackerMil;
// TrackerBoosting;
// TrackerMedianFlow;
// TrackerTld;
// TrackerKcf;
// TrackerMosse;
// TrackerCsrt;
//

// (private) implementation helpers for init and update (saves some boilerplate code)
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
// Mil
// The MIL algorithm trains a classifier in an online manner to separate the object from the background.
// Multiple Instance Learning avoids the drift problem for a robust tracking.
//
// s.a: https://docs.opencv.org/master/d0/d26/classcv_1_1TrackerMIL.html
type TrackerMil struct {
	p C.TrackerMil
}

func NewTrackerMil() TrackerMil {
	return TrackerMil{p: C.TrackerMil_Create()}
}

func (self *TrackerMil) Close() error {
	C.TrackerMil_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerMil) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerMil) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// Boosting
// This is a real-time object tracking based on a novel on-line version of the AdaBoost algorithm.
//
// s.a: https://docs.opencv.org/master/d1/d1a/classcv_1_1TrackerBoosting.html
type TrackerBoosting struct {
	p C.TrackerBoosting
}

func NewTrackerBoosting() TrackerBoosting {
	return TrackerBoosting{p: C.TrackerBoosting_Create()}
}

func (self *TrackerBoosting) Close() error {
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
// MedianFlow
// The tracker is suitable for very smooth and predictable movements when object is visible throughout the whole sequence
//
// s.a: https://docs.opencv.org/master/d7/d86/classcv_1_1TrackerMedianFlow.html
type TrackerMedianFlow struct {
	p C.TrackerMedianFlow
}

func NewTrackerMedianFlow() TrackerMedianFlow {
	return TrackerMedianFlow{p: C.TrackerMedianFlow_Create()}
}

func (self *TrackerMedianFlow) Close() error {
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
// Tld
// TLD is a novel tracking framework that explicitly decomposes the long-term tracking task into tracking, learning and detection.
//
// s.a: https://docs.opencv.org/master/dc/d1c/classcv_1_1TrackerTLD.html
type TrackerTld struct {
	p C.TrackerTld
}

func NewTrackerTld() TrackerTld {
	return TrackerTld{p: C.TrackerTld_Create()}
}

func (self *TrackerTld) Close() error {
	C.TrackerTld_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerTld) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerTld) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// Kcf
// KCF is a novel tracking framework that utilizes properties of circulant matrix to enhance the processing speed.
//
// s.a: https://docs.opencv.org/master/d2/dff/classcv_1_1TrackerKCF.html
//
type TrackerKcf struct {
	p C.TrackerKcf
}

func NewTrackerKcf() TrackerKcf {
	return TrackerKcf{p: C.TrackerKcf_Create()}
}

func (self *TrackerKcf) Close() error {
	C.TrackerKcf_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerKcf) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerKcf) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// Mosse
// based on: Visual Object Tracking using Adaptive Correlation Filters
// note, that this tracker is working on graysccale images
//
// s.a: https://docs.opencv.org/master/d0/d02/classcv_1_1TrackerMOSSE.html
//
type TrackerMosse struct {
	p C.TrackerMosse
}

func NewTrackerMosse() TrackerMosse {
	return TrackerMosse{p: C.TrackerMosse_Create()}
}

func (self *TrackerMosse) Close() error {
	C.TrackerMosse_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerMosse) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerMosse) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}

//
// Csrt
// Discriminative Correlation Filter Tracker with Channel and Spatial Reliability.
//
// s.a: https://docs.opencv.org/master/d2/da2/classcv_1_1TrackerCSRT.html
//
type TrackerCsrt struct {
	p C.TrackerCsrt
}

func NewTrackerCsrt() TrackerCsrt {
	return TrackerCsrt{p: C.TrackerCsrt_Create()}
}

func (self *TrackerCsrt) Close() error {
	C.TrackerCsrt_Close(self.p)
	self.p = nil
	return nil
}

func (self TrackerCsrt) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return tracker_Init(C.Tracker(self.p), img, boundingBox)
}

func (self TrackerCsrt) Update(img gocv.Mat) (image.Rectangle, bool) {
	return tracker_Update(C.Tracker(self.p), img)
}
