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
}

// NewWindow creates a new named cv window
func NewWindow(name string) *Window {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_New(cName, 1)

	return &Window{name: name}
}

// Delete a specific Window
func (w *Window) Delete() {
	cName := C.CString(w.name)
	defer C.free(unsafe.Pointer(cName))

	C.Window_Delete(cName)
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
