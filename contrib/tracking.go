package contrib

/*
#include "tracking.h"
*/
import "C"
import (
	"image"

	"gocv.io/x/gocv"
)

// TrackerKCF is a Tracker based on KCF, which is a novel tracking framework that
// utilizes properties of circulant matrix to enhance the processing speed.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/dff/classcv_1_1TrackerKCF.html
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

// TrackerCSRT is an implementation of Discriminative Correlation Filter Tracker
// with Channel and Spatial Reliability.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/da2/classcv_1_1TrackerCSRT.html
//
type TrackerCSRT struct {
	p C.TrackerCSRT
}

// NewTrackerCSRT returns a new TrackerCSRT.
func NewTrackerCSRT() gocv.Tracker {
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

func trackerInit(trk C.Tracker, img gocv.Mat, boundingBox image.Rectangle) bool {
	cBox := C.struct_Rect{
		x:      C.int(boundingBox.Min.X),
		y:      C.int(boundingBox.Min.Y),
		width:  C.int(boundingBox.Size().X),
		height: C.int(boundingBox.Size().Y),
	}

	ret := C.TrackerSubclass_Init(trk, C.Mat(img.Ptr()), cBox)
	return bool(ret)
}

func trackerUpdate(trk C.Tracker, img gocv.Mat) (image.Rectangle, bool) {
	cBox := C.struct_Rect{}

	ret := C.TrackerSubclass_Update(trk, C.Mat(img.Ptr()), &cBox)

	rect := image.Rect(int(cBox.x), int(cBox.y), int(cBox.x+cBox.width), int(cBox.y+cBox.height))
	return rect, bool(ret)
}
