package gocv

import (
	"strings"
	"testing"
)

func TestVersions(t *testing.T) {
	ocvv := OpenCVVersion()

	// TODO CV5: we should probably remove the 4.x checks
	if !(strings.Contains(ocvv, "4.11") ||
		strings.Contains(ocvv, "4.12") ||
		strings.Contains(ocvv, "4.13") ||
		strings.HasPrefix(ocvv, "5.")) {
		t.Error("Wrong version of OpenCV:", ocvv)
	}

	v := Version()

	if v != GoCVVersion {
		t.Error("Wrong version of GoCV")
	}
}
