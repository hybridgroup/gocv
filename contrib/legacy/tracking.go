package legacy

/*
#include "tracking.h"
*/
import "C"
import (
	"image"

	"gocv.io/x/gocv"
)

// TrackerBoosting is a real-time object tracker based
// on a novel on-line version of the AdaBoost algorithm.
//
// For further details, please see:
// https://docs.opencv.org/master/db/df1/classcv_1_1legacy_1_1TrackerBoosting.html
//
type TrackerBoosting struct {
	p C.TrackerBoosting
}

// NewTrackerBoosting returns a new TrackerBoosting.
func NewTrackerBoosting() gocv.Tracker {
	return TrackerBoosting{p: C.TrackerBoosting_Create()}
}

// Close closes the TrackerBoosting.
func (trk TrackerBoosting) Close() error {
	C.TrackerBoosting_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes the Tracker.
func (trk TrackerBoosting) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates the Tracker.
func (trk TrackerBoosting) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerCSRT is an implementation of Discriminative Correlation Filter Tracker
// with Channel and Spatial Reliability.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d8f/classcv_1_1legacy_1_1TrackerCSRT.html
//
type TrackerCSRT struct {
	p C.TrackerCSRT
}

// NewTrackerCSRT returns a new TrackerCSRT.
func NewTrackerCSRT() Tracker {
	return TrackerCSRT{p: C.TrackerCSRT_Create()}
}

// Close closes this Tracker.
func (trk TrackerCSRT) Close() error {
	C.TrackerCSRT_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes this Tracker.
func (trk TrackerCSRT) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates this Tracker.
func (trk TrackerCSRT) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerKCF is a Tracker based on KCF, which is a novel tracking framework that
// utilizes properties of circulant matrix to enhance the processing speed.
//
// For further details, please see:
// https://docs.opencv.org/master/d5/d6a/classcv_1_1legacy_1_1TrackerKCF.html
//
type TrackerKCF struct {
	p C.TrackerKCF
}

// NewTrackerKCF returns a new TrackerKCF.
func NewTrackerKCF() gocv.Tracker {
	return TrackerKCF{p: C.TrackerKCF_Create()}
}

// Close closes this Tracker.
func (trk TrackerKCF) Close() error {
	C.TrackerKCF_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes this Tracker.
func (trk TrackerKCF) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates this Tracker.
func (trk TrackerKCF) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerMedianFlow is a Tracker implementation suitable for very smooth and predictable movements
// when the object is visible throughout the whole sequence.
//
// For further details, please see:
// https://docs.opencv.org/master/dd/d94/classcv_1_1legacy_1_1TrackerMedianFlow.html
//
type TrackerMedianFlow struct {
	p C.TrackerMedianFlow
}

// NewTrackerMedianFlow returns a new TrackerMedianFlow.
func NewTrackerMedianFlow() gocv.Tracker {
	return TrackerMedianFlow{p: C.TrackerMedianFlow_Create()}
}

// Close closes the Tracker.
func (trk TrackerMedianFlow) Close() error {
	C.TrackerMedianFlow_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes the Tracker.
func (trk TrackerMedianFlow) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates the Tracker.
func (trk TrackerMedianFlow) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerMIL is a Tracker that uses the MIL algorithm. MIL trains a classifier in an online manner
// to separate the object from the background.
// Multiple Instance Learning avoids the drift problem for a robust tracking.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/dbc/classcv_1_1legacy_1_1TrackerMIL.html
//
type TrackerMIL struct {
	p C.TrackerMIL
}

// NewTrackerMIL returns a new TrackerMIL.
func NewTrackerMIL() gocv.Tracker {
	return TrackerMIL{p: C.TrackerMIL_Create()}
}

// Close closes the TrackerMIL.
func (trk TrackerMIL) Close() error {
	C.TrackerMIL_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes the TrackerMIL.
func (trk TrackerMIL) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates the TrackerMIL.
func (trk TrackerMIL) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerMOSSE uses Visual Object Tracking using Adaptive Correlation Filters.
// Note, that this tracker only works on graysccale images.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d20/classcv_1_1legacy_1_1TrackerMOSSE.html
//
type TrackerMOSSE struct {
	p C.TrackerMOSSE
}

// NewTrackerMOSSE returns a new TrackerMOSSE.
func NewTrackerMOSSE() gocv.Tracker {
	return TrackerMOSSE{p: C.TrackerMOSSE_Create()}
}

// Close closes this Tracker.
func (trk TrackerMOSSE) Close() error {
	C.TrackerMOSSE_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes this Tracker.
func (trk TrackerMOSSE) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates this Tracker.
func (trk TrackerMOSSE) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}

// TrackerTLD is a novel tracking framework that explicitly decomposes
// the long-term tracking task into tracking, learning and detection.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/da6/classcv_1_1legacy_1_1TrackerTLD.html
//
type TrackerTLD struct {
	p C.TrackerTLD
}

// NewTrackerTLD returns a new TrackerTLD.
func NewTrackerTLD() gocv.Tracker {
	return TrackerTLD{p: C.TrackerTLD_Create()}
}

// Close closes this Tracker.
func (trk TrackerTLD) Close() error {
	C.TrackerTLD_Close(trk.p)
	trk.p = nil
	return nil
}

// Init initializes this Tracker.
func (trk TrackerTLD) Init(img gocv.Mat, boundingBox image.Rectangle) bool {
	return trackerInit(C.Tracker(trk.p), img, boundingBox)
}

// Update updates this Tracker.
func (trk TrackerTLD) Update(img gocv.Mat) (image.Rectangle, bool) {
	return trackerUpdate(C.Tracker(trk.p), img)
}
