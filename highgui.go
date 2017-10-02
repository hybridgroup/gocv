package gocv

/*
#include <stdlib.h>
#include "highgui.h"
*/
import "C"
import (
	"unsafe"
)

// Window is a wrapper around OpenCV's "HighGUI" named windows.
// While OpenCV was designed for use in full-scale applications and can be used
// within functionally rich UI frameworks (such as Qt*, WinForms*, or Cocoa*)
// or without any UI at all, sometimes there it is required to try functionality
// quickly and visualize the results. This is what the HighGUI module has been designed for.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/dfc/group__highgui.html
//
type Window struct {
	name string
	open bool
}

// NewWindow creates a new named OpenCV window
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/dfc/group__highgui.html#ga5afdf8410934fd099df85c75b2e0888b
//
func NewWindow(name string) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_New(cName, 1)

	return &Window{name: name, open: true}
}

// Close closes and deletes a named OpenCV Window.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/dfc/group__highgui.html#ga851ccdd6961022d1d5b4c4f255dbab34
//
func (w *Window) Close() error {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_Close(cName)
	w.open = false
	return nil
}

// IsOpen checks to see if the Window seems to be open.
func (w *Window) IsOpen() bool {
	return w.open
}

// IMShow displays an image Mat in the specified window.
// This function should be followed by the WaitKey function which displays
// the image for specified milliseconds. Otherwise, it won't display the image.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/dfc/group__highgui.html#ga453d42fe4cb60e5723281a89973ee563
//
func (w *Window) IMShow(img Mat) {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_IMShow(cName, img.p)
}

// WaitKey waits for a pressed key.
// This function is the only method in OpenCV's HighGUI that can fetch
// and handle events, so it needs to be called periodically
// for normal event processing
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/dfc/group__highgui.html#ga5628525ad33f52eab17feebcfba38bd7
//
func WaitKey(delay int) int {
	return int(C.Window_WaitKey(C.int(delay)))
}
