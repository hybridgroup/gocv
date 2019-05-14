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
	defer img.Close()

	rect := image.Rect(250, 150, 250+200, 150+250)
	init := tracker.Init(img, rect)
	if !init {
		t.Error("TestTracker " + name + " failed in Init")
	}

	_, ok := tracker.Update(img)
	if !ok {
		t.Error("TestTracker " + name + " lost object in Update")
	}
}

func TestSingleTrackers(t *testing.T) {
	tab := []struct {
		name    string
		tracker Tracker
	}{
		{"MIL", NewTrackerMIL()},
		{"Boosting", NewTrackerBoosting()},
		{"MedianFlow", NewTrackerMedianFlow()},
		{"TLD", NewTrackerTLD()},
		{"KCF", NewTrackerKCF()},
		{"MOSSE", NewTrackerMOSSE()},
		{"CSRT", NewTrackerCSRT()},
	}

	for _, test := range tab {
		func() {
			defer test.tracker.Close()
			BaseTestTracker(t, test.tracker, test.name)
		}()
	}
}
