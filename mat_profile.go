// +build matprofile

package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import (
	"C"
)

import (
	"runtime/pprof"
)

// MatProfile a pprof.Profile that contains stack traces that led to (currently)
// unclosed Mat's creations.  Every time a Mat is created, the stack trace is
// added to this profile and every time the Mat is closed the trace is removed.
// In a program that is not leaking, this profile's count should not
// continuously increase and ideally when a program is terminated the count
// should be zero.  You can get the count at any time with:
//
//	gocv.MatProfile.Count()
//
// and you can display the current entries with:
//
// 	var b bytes.Buffer
//	gocv.MatProfile.WriteTo(&b, 1)
//	fmt.Print(b.String())
//
// This will display stack traces of where the unclosed Mats were instantiated.
// For example, the results could look something like this:
//
//	1 @ 0x4146a0c 0x4146a57 0x4119666 0x40bb18f 0x405a841
//	#	0x4146a0b	gocv.io/x/gocv.newMat+0x4b	/go/src/gocv.io/x/gocv/core.go:120
//	#	0x4146a56	gocv.io/x/gocv.NewMat+0x26	/go/src/gocv.io/x/gocv/core.go:126
//	#	0x4119665	gocv.io/x/gocv.TestMat+0x25	/go/src/gocv.io/x/gocv/core_test.go:29
//	#	0x40bb18e	testing.tRunner+0xbe		/usr/local/Cellar/go/1.11/libexec/src/testing/testing.go:827
//
// Furthermore, if the program is a long running process or if gocv is being used on a
// web server, it may be helpful to install the HTTP interface using:
//
//	import _ "net/http/pprof"
//
// In order to include the MatProfile custom profiler, you MUST build or run your application
// or tests using the following build tag:
// -tags matprofile
//
// For more information, see the runtime/pprof package documentation.
var MatProfile *pprof.Profile

func init() {
	profName := "gocv.io/x/gocv.Mat"
	MatProfile = pprof.Lookup(profName)
	if MatProfile == nil {
		MatProfile = pprof.NewProfile(profName)
	}
}

// newMat returns a new Mat from a C Mat and records it to the MatProfile.
func newMat(p C.Mat) Mat {
	m := Mat{p: p}
	MatProfile.Add(p, 1)
	return m
}

// Close the Mat object.
func (m *Mat) Close() error {
	C.Mat_Close(m.p)
	MatProfile.Remove(m.p)
	m.p = nil
	return nil
}
