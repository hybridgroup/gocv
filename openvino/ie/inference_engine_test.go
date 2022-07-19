package ie

import (
	"strings"
	"testing"
)

func TestInferenceEngineVersion(t *testing.T) {
	v := Version()
	expected := "2022.1"
	if v == "" || len(v) < len(expected) {
		t.Error("Invalid IE Version")
		return
	}

	// Different platforms can have different strings, only warn if expected
	// version string does not match.
	// eg. archlinux: 2.1.custom_makepkg_cdb9bec7210f8c24fde3e416c7ada820faaaa23e
	if !strings.Contains(v, expected) {
		t.Log("Unexpected IE Version: ", v)
	}
}
