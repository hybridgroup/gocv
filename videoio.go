package opencv3

/*
#include <stdlib.h>
#include "videoio.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// VideoCapture is a wrapper around the cv::VideoCapture class.
type VideoCapture struct {
	p C.VideoCapture
}

// VideoCaptureFile opens a VideoCapture from a file and prepares to start capturing
func VideoCaptureFile(uri string) (vc VideoCapture, err error) {
	vc = VideoCapture{p: C.VideoCapture_New()}

	cURI := C.CString(uri)
	defer C.free(unsafe.Pointer(cURI))

	C.VideoCapture_Open(vc.p, cURI)
	return vc, nil
}

// VideoCaptureDevice opens a VideoCapture from a device and prepares to start capturing
func VideoCaptureDevice(device int) (vc VideoCapture, err error) {
	vc = VideoCapture{p: C.VideoCapture_New()}
	C.VideoCapture_OpenDevice(vc.p, C.int(device))
	return vc, nil
}

// Close VideoCapture object.
func (v *VideoCapture) Close() error {
	C.VideoCapture_Close(v.p)
	v.p = nil
	return nil
}

// Release video capture object.
func (v *VideoCapture) Release() {
	C.VideoCapture_Release(v.p)
}

// Set parameter with property (=key).
func (v *VideoCapture) Set(prop int, param int) {
	C.VideoCapture_Set(v.p, C.int(prop), C.int(param))
}

// IsOpened returns the video capture opens a file(or device) or not.
func (v *VideoCapture) IsOpened() bool {
	isOpened := C.VideoCapture_IsOpened(v.p)
	return isOpened != 0
}

// Read set frame to argument Mat, returns `false` when the video capture
// cannot read frame.
func (v *VideoCapture) Read(m Mat) bool {
	return C.VideoCapture_Read(v.p, m.p) != 0
}

// Grab `skip` count frames.
func (v *VideoCapture) Grab(skip int) {
	C.VideoCapture_Grab(v.p, C.int(skip))
}

// VideoWriter is a bind of `cv::VideoWriter`.
type VideoWriter struct {
	mu sync.RWMutex
	p  C.VideoWriter
}

// VideoWriterFile opens a VideoWriter with a specific output file.
func VideoWriterFile(name string, fps float64, width int, height int) (vw VideoWriter, err error) {
	vw = VideoWriter{p: C.VideoWriter_New()}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.VideoWriter_Open(vw.p, cName, C.double(fps), C.int(width), C.int(height))
	return vw, nil
}

// VideoWriterFileMat opens a VideoWriter with a specific output file,
// using the dimensions from a specific Mat.
func VideoWriterFileMat(name string, fps float64, img Mat) (vw VideoWriter, err error) {
	vw = VideoWriter{p: C.VideoWriter_New()}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.VideoWriter_OpenWithMat(vw.p, cName, C.double(fps), img.p)
	return vw, nil
}

// Close VideoWriter object.
func (vw *VideoWriter) Close() error {
	C.VideoWriter_Close(vw.p)
	vw.p = nil
	return nil
}

// IsOpened checks if the VideoWriter is open and ready to be written to.
func (vw *VideoWriter) IsOpened() bool {
	isOpend := C.VideoWriter_IsOpened(vw.p)
	return isOpend != 0
}

// Write a single Mat image to the open VideoWriter.
func (vw *VideoWriter) Write(img Mat) error {
	vw.mu.Lock()
	defer vw.mu.Unlock()
	C.VideoWriter_Write(vw.p, img.p)
	return nil
}
