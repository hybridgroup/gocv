// Package cuda is the GoCV wrapper around OpenCV cuda.
//
// For further details, please see:
// https://github.com/opencv/opencv
//
// import "gocv.io/x/gocv/cuda"
package cuda

/*
#include <stdlib.h>
#include "cuda.h"
*/
import "C"
import "gocv.io/x/gocv"

// GpuMat is the GPU version of a Mat
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html
type GpuMat struct {
	p C.GpuMat
}

// Upload performs data upload to GpuMat (Blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a00ef5bfe18d14623dcf578a35e40a46b
//
func (g *GpuMat) Upload(data gocv.Mat) {
	C.GpuMat_Upload(g.p, C.Mat(data.Ptr()), nil)
}

// UploadWithStream performs data upload to GpuMat (non-blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a00ef5bfe18d14623dcf578a35e40a46b
//
func (g *GpuMat) UploadWithStream(data gocv.Mat, s Stream) {
	C.GpuMat_Upload(g.p, C.Mat(data.Ptr()), s.p)
}

// Download performs data download from GpuMat (Blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a027e74e4364ddfd9687b58aa5db8d4e8
func (g *GpuMat) Download(dst *gocv.Mat) {
	C.GpuMat_Download(g.p, C.Mat(dst.Ptr()), nil)
}

// DownloadWithStream performs data download from GpuMat (non-blocking call)
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a027e74e4364ddfd9687b58aa5db8d4e8
func (g *GpuMat) DownloadWithStream(dst *gocv.Mat, s Stream) {
	C.GpuMat_Download(g.p, C.Mat(dst.Ptr()), s.p)
}

// Empty returns true if GpuMat is empty
func (g *GpuMat) Empty() bool {
	return C.GpuMat_Empty(g.p) != 0
}

// Close the GpuMat object
func (g *GpuMat) Close() error {
	C.GpuMat_Close(g.p)
	g.p = nil
	return nil
}

// NewGpuMat returns a new empty GpuMat
func NewGpuMat() GpuMat {
	return newGpuMat(C.GpuMat_New())
}

// NewGpuMatFromMat returns a new GpuMat based on a Mat
func NewGpuMatFromMat(mat gocv.Mat) GpuMat {
	return newGpuMat(C.GpuMat_NewFromMat(C.Mat(mat.Ptr())))
}

// NewGpuMatWithSize returns a new GpuMat with a specific size and type.
func NewGpuMatWithSize(rows int, cols int, mt gocv.MatType) GpuMat {
	return newGpuMat(C.GpuMat_NewWithSize(C.int(rows), C.int(cols), C.int(mt)))
}

func newGpuMat(p C.GpuMat) GpuMat {
	return GpuMat{p: p}
}

// PrintCudaDeviceInfo prints extensive cuda device information
func PrintCudaDeviceInfo(device int) {
	C.PrintCudaDeviceInfo(C.int(device))
}

// PrintShortCudaDeviceInfo prints a small amount of cuda device information
func PrintShortCudaDeviceInfo(device int) {
	C.PrintShortCudaDeviceInfo(C.int(device))
}

// GetCudaEnabledDeviceCount returns the number of cuda enabled devices on the
// system
func GetCudaEnabledDeviceCount() int {
	return int(C.GetCudaEnabledDeviceCount())
}

// GetDevice returns the current device index.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#ga6ded4ed8e4fc483a9863d31f34ec9c0e
//
func GetDevice() int {
	return int(C.GetCudaDevice())
}

// SetDevice sets a device and initializes it for the current thread.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#gaefa34186b185de47851836dba537828b
//
func SetDevice(device int) {
	C.SetCudaDevice(C.int(device))
}

// ResetDevice explicitly destroys and cleans up all resources associated
// with the current device in the current process.
//
// Any subsequent API call to this device will reinitialize the device.
//
// For further details, please see:
// https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#ga6153b6f461101374e655a54fc77e725e
//
func ResetDevice() {
	C.ResetCudaDevice()
}

// ConvertTo converts GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a3a1b076e54d8a8503014e27a5440d98a
//
func (m *GpuMat) ConvertTo(dst *GpuMat, mt gocv.MatType) {
	C.GpuMat_ConvertTo(m.p, dst.p, C.int(mt), nil)
	return
}

// ConvertToWithStream converts GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a3a1b076e54d8a8503014e27a5440d98a
//
func (m *GpuMat) ConvertToWithStream(dst *GpuMat, mt gocv.MatType, s Stream) {
	C.GpuMat_ConvertTo(m.p, dst.p, C.int(mt), s.p)
	return
}

// CopyTo copies GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a948c562ee340c0678a44884bde1f5a3e
//
func (m *GpuMat) CopyTo(dst *GpuMat) {
	C.GpuMat_CopyTo(m.p, dst.p, nil)
	return
}

// CopyToWithStream copies GpuMat into destination GpuMat.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a948c562ee340c0678a44884bde1f5a3e
//
func (m *GpuMat) CopyToWithStream(dst *GpuMat, s Stream) {
	C.GpuMat_CopyTo(m.p, dst.p, s.p)
	return
}

// Rows returns the number of rows for this GpuMat.
func (m *GpuMat) Rows() int {
	return int(C.GpuMat_Rows(m.p))
}

// Cols returns the number of columns for this GpuMat.
func (m *GpuMat) Cols() int {
	return int(C.GpuMat_Cols(m.p))
}

// Channels returns the number of channels for this GpuMat.
func (m *GpuMat) Channels() int {
	return int(C.GpuMat_Channels(m.p))
}

// Type returns the type for this GpuMat.
func (m *GpuMat) Type() gocv.MatType {
	return gocv.MatType(C.GpuMat_Type(m.p))
}

// Reshape creates a new GpuMat with the same data
// but with a different number of channels and/or different number of rows.
//
// For further details, please see:
// https://docs.opencv.org/master/d0/d60/classcv_1_1cuda_1_1GpuMat.html#a408e22ed824d1ddf59f58bda895017a8
//
func (m *GpuMat) Reshape(cn int, rows int) GpuMat {
	return newGpuMat(C.GpuMat_Reshape(m.p, C.int(cn), C.int(rows)))
}

// Stream asynchronous stream used for CUDA operations.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/df3/classcv_1_1cuda_1_1Stream.html#aa6434e2f5f29bd81406732b39951c246
type Stream struct {
	p C.Stream
}

// NewStream returns a new empty Stream.
func NewStream() Stream {
	return Stream{p: C.Stream_New()}
}

// Close the Stream.
func (s *Stream) Close() error {
	C.Stream_Close(s.p)
	s.p = nil
	return nil
}

// QueryIfComplete returns true if the current stream queue is finished
//
// For further details, please see:
// https://docs.opencv.org/master/d9/df3/classcv_1_1cuda_1_1Stream.html#a9fab618395d42fa31987506e42fab1b4
//
func (s *Stream) QueryIfComplete() bool {
	return bool(C.Stream_QueryIfComplete(s.p))
}

// WaitForCompletion blocks the current CPU thread until all operations in the stream are complete.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/df3/classcv_1_1cuda_1_1Stream.html#a0e1d939503e8faad741ab584b720bca6
//
func (s *Stream) WaitForCompletion() {
	C.Stream_WaitForCompletion(s.p)
}
