package contrib

import (
	"gocv.io/x/gocv"
	"image"
	"testing"
)

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
		defer test.tracker.Close()

		if test.tracker == nil {
			t.Error("TestTracker " + test.name + " should not be nil")
		}

		img := gocv.IMRead("../images/face.jpg", 1)
		if img.Empty() {
			t.Error("TestTracker " + test.name + " input img failed to load")
		}
		defer img.Close()

		rect := image.Rect(250, 150, 250+200, 150+250)
		init := test.tracker.Init(img, rect)
		if !init {
			t.Error("TestTracker " + test.name + " failed in Init")
		}

		_, ok := test.tracker.Update(img)
		if !ok {
			t.Error("TestTracker " + test.name + " lost object in Update")
		}
	}
}
