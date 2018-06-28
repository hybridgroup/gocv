package ie

import (
	"testing"
)

func TestInferenceEngineVersion(t *testing.T) {
	if Version() == "" {
		t.Error("Invalid IE Version")
		return
	}
}
