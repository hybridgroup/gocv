package opencv3

/*
#cgo linux pkg-config: opencv
#cgo darwin pkg-config: opencv
#include <stdlib.h>
#include "videowriter.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// VideoWriter is a bind of `cv::VideoWriter`.
type VideoWriter struct {
	mu sync.RWMutex
	p  C.VideoWriter
}

// NewVideoWriter returns a new video writer.
func NewVideoWriter() VideoWriter {
	return VideoWriter{p: C.VideoWriter_New()}
}

// Delete object.
func (vw *VideoWriter) Delete() {
	C.VideoWriter_Delete(vw.p)
	vw.p = nil
}

// Open a video writer.
func (vw *VideoWriter) Open(name string, fps float64, width int, height int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.VideoWriter_Open(vw.p, cName, C.double(fps), C.int(width), C.int(height))
}

// OpenWithMat opens video writer.
func (vw *VideoWriter) OpenWithMat(name string, fps float64, img MatVec3b) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.VideoWriter_OpenWithMat(vw.p, cName, C.double(fps), img.p)
}

// IsOpened returns the video writer opens a file or not.
func (vw *VideoWriter) IsOpened() bool {
	isOpend := C.VideoWriter_IsOpened(vw.p)
	return isOpend != 0
}

// Write the image to file.
func (vw *VideoWriter) Write(img MatVec3b) {
	vw.mu.Lock()
	defer vw.mu.Unlock()
	C.VideoWriter_Write(vw.p, img.p)
}
