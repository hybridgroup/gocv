package gocv

import (
	"testing"
)

func TestRect(t *testing.T) {
	t.Skip("Test needed")
}
func TestGetTextSize(t *testing.T) {
	size := GetTextSize("test", FontHersheySimplex, 1.2, 1)
	if size.X != 72 {
		t.Error("Invalid text size width")
	}

	if size.Y != 26 {
		t.Error("Invalid text size height")
	}
}
func TestPutText(t *testing.T) {
	t.Skip("Test needed")
}
