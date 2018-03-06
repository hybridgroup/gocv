package contrib

import (
	"gocv.io/x/gocv"
	"image"
	"testing"
)

func BaseTestTracker(t *testing.T, tracker Tracker, name string) {
	if tracker == nil {
		t.Error("TestTracker " + name + " should not be nil")
	}
	img := gocv.IMRead("../images/face.jpg", 1)
	if img.Empty() {
		t.Error("TestTracker " + name + " input img failed to load")
	}
	rect := image.Rect(250, 150, 250+200, 150+250)
	init := tracker.Init(img, rect)
	if !init {
		t.Error("TestTracker " + name + " failed in Init")
	}
	_, ok := tracker.Update(img)
	if !ok {
		t.Error("TestTracker " + name + " lost object in Update")
	}
	img.Close()
}

func TestTrackerMil(t *testing.T) {
	tracker := NewTrackerMil()
	BaseTestTracker(t, tracker, "Mil")
	tracker.Close()
}
func TestTrackerBoosting(t *testing.T) {
	tracker := NewTrackerBoosting()
	BaseTestTracker(t, tracker, "Boosting")
	tracker.Close()
}
func TestTrackerMedianFlow(t *testing.T) {
	tracker := NewTrackerMedianFlow()
	BaseTestTracker(t, tracker, "MedianFlow")
	tracker.Close()
}
func TestTrackerTld(t *testing.T) {
	tracker := NewTrackerTld()
	BaseTestTracker(t, tracker, "Tld")
	tracker.Close()
}
func TestTrackerKcf(t *testing.T) {
	tracker := NewTrackerKcf()
	BaseTestTracker(t, tracker, "Kcf")
	tracker.Close()
}

func TestTrackerMosse(t *testing.T) {
	tracker := NewTrackerMosse()
	BaseTestTracker(t, tracker, "Mosse")
	tracker.Close()
}

func TestTrackerCsrt(t *testing.T) {
	tracker := NewTrackerCsrt()
	BaseTestTracker(t, tracker, "Csrt")
	tracker.Close()
}
