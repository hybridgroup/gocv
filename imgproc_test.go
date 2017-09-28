package opencv3

import (
	"testing"
)

func TestRect(t *testing.T) {
	t.Skip("Test needed")
}
func TestGetTextSize(t *testing.T) {
	size := GetTextSize("test", FontHersheySimplex, 1.2, 1)
	if size.Width != 72 {
		t.Error("Invalid text size width")
	}

	if size.Height != 26 {
		t.Error("Invalid text size height")
	}
}
func TestPutText(t *testing.T) {
	t.Skip("Test needed")
}
