package opencv3

/*
#include <stdlib.h>
#include "highgui.h"
*/
import "C"
import (
	"unsafe"
)

// Window is a bind of OpenCV's highgui windows
type Window struct {
	name string
	open bool
}

// NewWindow creates a new named cv window
func NewWindow(name string) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_New(cName, 1)

	return &Window{name: name, open: true}
}

// Close deletes a specific Window
func (w *Window) Close() {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_Close(cName)
	w.open = false
}

// IsOpen checks to see if the Window seems to be open.
func (w *Window) IsOpen() bool {
	return w.open
}

// IMShow takes an image Mat and displays it in the Window
func (w *Window) IMShow(img Mat) {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_IMShow(cName, img.p)
}

// WaitKey waits for keyboard input
func WaitKey(delay int) int {
	return int(C.Window_WaitKey(C.int(delay)))
}
