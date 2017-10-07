package gocv

/*
#include <stdlib.h>
#include "videoio.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// VideoCapture is a wrapper around the OpenCV VideoCapture class.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d8/dfe/classcv_1_1VideoCapture.html
//
type VideoCapture struct {
	p C.VideoCapture
}

// VideoCaptureFile opens a VideoCapture from a file and prepares
// to start capturing.
func VideoCaptureFile(uri string) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}

	cURI := C.CString(uri)
	defer C.free(unsafe.Pointer(cURI))

	C.VideoCapture_Open(vc.p, cURI)
	return
}

// VideoCaptureDevice opens a VideoCapture from a device and prepares
// to start capturing.
func VideoCaptureDevice(device int) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}
	C.VideoCapture_OpenDevice(vc.p, C.int(device))
	return
}

// Close VideoCapture object.
func (v *VideoCapture) Close() error {
	C.VideoCapture_Close(v.p)
	v.p = nil
	return nil
}

// Set parameter with property (=key).
func (v *VideoCapture) Set(prop int, param int) {
	C.VideoCapture_Set(v.p, C.int(prop), C.int(param))
}

// IsOpened returns if the VideoCapture has been opened to read from
// a file or capture device.
func (v *VideoCapture) IsOpened() bool {
	isOpened := C.VideoCapture_IsOpened(v.p)
	return isOpened != 0
}

// Read read the next frame from the VideoCapture to the Mat passed in
// as the parem. It returns false if the VideoCapture cannot read frame.
func (v *VideoCapture) Read(m Mat) bool {
	return C.VideoCapture_Read(v.p, m.p) != 0
}

// Grab skips a specific number of frames.
func (v *VideoCapture) Grab(skip int) {
	C.VideoCapture_Grab(v.p, C.int(skip))
}

// VideoWriter is a wrapper around the OpenCV VideoWriter`class.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d9e/classcv_1_1VideoWriter.html
//
type VideoWriter struct {
	mu *sync.RWMutex
	p  C.VideoWriter
}

// VideoWriterFile opens a VideoWriter with a specific output file.
func VideoWriterFile(name string, fps float64, width int, height int) (vw *VideoWriter, err error) {
	vw = &VideoWriter{
		p:  C.VideoWriter_New(),
		mu: &sync.RWMutex{},
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.VideoWriter_Open(vw.p, cName, C.double(fps), C.int(width), C.int(height))
	return
}

// Close VideoWriter object.
func (vw *VideoWriter) Close() error {
	C.VideoWriter_Close(vw.p)
	vw.p = nil
	return nil
}

// IsOpened checks if the VideoWriter is open and ready to be written to.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d9e/classcv_1_1VideoWriter.html#a9a40803e5f671968ac9efa877c984d75
//
func (vw *VideoWriter) IsOpened() bool {
	isOpend := C.VideoWriter_IsOpened(vw.p)
	return isOpend != 0
}

// Write the next video frame from the Mat image to the open VideoWriter.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d9e/classcv_1_1VideoWriter.html#a3115b679d612a6a0b5864a0c88ed4b39
//
func (vw *VideoWriter) Write(img Mat) error {
	vw.mu.Lock()
	defer vw.mu.Unlock()
	C.VideoWriter_Write(vw.p, img.p)
	return nil
}
