package gocv

/*
#include <stdlib.h>
#include "dnn.h"
*/
import "C"
import (
	"image"
	"reflect"
	"unsafe"
)

// Net allows you to create and manipulate comprehensive artificial neural networks.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html
//
type Net struct {
	// C.Net
	p unsafe.Pointer
}

// NetBackendType is the type for the various different kinds of DNN backends.
type NetBackendType int

const (
	// NetBackendDefault is the default backend.
	NetBackendDefault NetBackendType = 0

	// NetBackendHalide is the Halide backend.
	NetBackendHalide NetBackendType = 1

	// NetBackendOpenVINO is the OpenVINO backend.
	NetBackendOpenVINO NetBackendType = 2

	// NetBackendOpenCV is the OpenCV backend.
	NetBackendOpenCV NetBackendType = 3

	// NetBackendVKCOM is the Vulkan backend.
	NetBackendVKCOM NetBackendType = 4
)

// ParseNetBackend returns a valid NetBackendType given a string. Valid values are:
// - halide
// - openvino
// - opencv
// - vulkan
// - default
func ParseNetBackend(backend string) NetBackendType {
	switch backend {
	case "halide":
		return NetBackendHalide
	case "openvino":
		return NetBackendOpenVINO
	case "opencv":
		return NetBackendOpenCV
	case "vulkan":
		return NetBackendVKCOM
	default:
		return NetBackendDefault
	}
}

// NetTargetType is the type for the various different kinds of DNN device targets.
type NetTargetType int

const (
	// NetTargetCPU is the default CPU device target.
	NetTargetCPU NetTargetType = 0

	// NetTargetFP32 is the 32-bit OpenCL target.
	NetTargetFP32 NetTargetType = 1

	// NetTargetFP16 is the 16-bit OpenCL target.
	NetTargetFP16 NetTargetType = 2

	// NetTargetVPU is the Movidius VPU target.
	NetTargetVPU NetTargetType = 3

	// NetTargetVulkan is the NVIDIA Vulkan target.
	NetTargetVulkan NetTargetType = 4

	// NetTargetFPGA is the FPGA target.
	NetTargetFPGA NetTargetType = 5
)

// ParseNetTarget returns a valid NetTargetType given a string. Valid values are:
// - cpu
// - fp32
// - fp16
// - vpu
// - vulkan
// - fpga
func ParseNetTarget(target string) NetTargetType {
	switch target {
	case "cpu":
		return NetTargetCPU
	case "fp32":
		return NetTargetFP32
	case "fp16":
		return NetTargetFP16
	case "vpu":
		return NetTargetVPU
	case "vulkan":
		return NetTargetVulkan
	case "fpga":
		return NetTargetFPGA
	default:
		return NetTargetCPU
	}
}

// Close Net
func (net *Net) Close() error {
	C.Net_Close((C.Net)(net.p))
	net.p = nil
	return nil
}

// Empty returns true if there are no layers in the network.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a6a5778787d5b8770deab5eda6968e66c
//
func (net *Net) Empty() bool {
	return bool(C.Net_Empty((C.Net)(net.p)))
}

// SetInput sets the new value for the layer output blob.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a672a08ae76444d75d05d7bfea3e4a328
//
func (net *Net) SetInput(blob Mat, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Net_SetInput((C.Net)(net.p), blob.p, cName)
}

// Forward runs forward pass to compute output of layer with name outputName.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a98ed94cb6ef7063d3697259566da310b
//
func (net *Net) Forward(outputName string) Mat {
	cName := C.CString(outputName)
	defer C.free(unsafe.Pointer(cName))

	return newMat(C.Net_Forward((C.Net)(net.p), cName))
}

// ForwardLayers forward pass to compute outputs of layers listed in outBlobNames.
//
// For further details, please see:
// https://docs.opencv.org/3.4.1/db/d30/classcv_1_1dnn_1_1Net.html#adb34d7650e555264c7da3b47d967311b
//
func (net *Net) ForwardLayers(outBlobNames []string) (blobs []Mat) {
	cMats := C.struct_Mats{}
	C.Net_ForwardLayers((C.Net)(net.p), &(cMats), toCStrings(outBlobNames))
	blobs = make([]Mat, cMats.length)
	for i := C.int(0); i < cMats.length; i++ {
		blobs[i].p = C.Mats_get(cMats, i)
	}
	return
}

// SetPreferableBackend ask network to use specific computation backend.
//
// For further details, please see:
// https://docs.opencv.org/3.4/db/d30/classcv_1_1dnn_1_1Net.html#a7f767df11386d39374db49cd8df8f59e
//
func (net *Net) SetPreferableBackend(backend NetBackendType) error {
	C.Net_SetPreferableBackend((C.Net)(net.p), C.int(backend))
	return nil
}

// SetPreferableTarget ask network to make computations on specific target device.
//
// For further details, please see:
// https://docs.opencv.org/3.4/db/d30/classcv_1_1dnn_1_1Net.html#a9dddbefbc7f3defbe3eeb5dc3d3483f4
//
func (net *Net) SetPreferableTarget(target NetTargetType) error {
	C.Net_SetPreferableTarget((C.Net)(net.p), C.int(target))
	return nil
}

// ReadNet reads a deep learning network represented in one of the supported formats.
//
// For further details, please see:
// https://docs.opencv.org/3.4/d6/d0f/group__dnn.html#ga3b34fe7a29494a6a4295c169a7d32422
//
func ReadNet(model string, config string) Net {
	cModel := C.CString(model)
	defer C.free(unsafe.Pointer(cModel))

	cConfig := C.CString(config)
	defer C.free(unsafe.Pointer(cConfig))
	return Net{p: unsafe.Pointer(C.Net_ReadNet(cModel, cConfig))}
}

// ReadNetBytes reads a deep learning network represented in one of the supported formats.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga138439da76f26266fdefec9723f6c5cd
//
func ReadNetBytes(framework string, model []byte, config []byte) (Net, error) {
	cFramework := C.CString(framework)
	defer C.free(unsafe.Pointer(cFramework))
	bModel, err := toByteArray(model)
	if err != nil {
		return Net{}, err
	}
	bConfig, err := toByteArray(config)
	if err != nil {
		return Net{}, err
	}
	return Net{p: unsafe.Pointer(C.Net_ReadNetBytes(cFramework, *bModel, *bConfig))}, nil
}

// ReadNetFromCaffe reads a network model stored in Caffe framework's format.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga29d0ea5e52b1d1a6c2681e3f7d68473a
//
func ReadNetFromCaffe(prototxt string, caffeModel string) Net {
	cprototxt := C.CString(prototxt)
	defer C.free(unsafe.Pointer(cprototxt))

	cmodel := C.CString(caffeModel)
	defer C.free(unsafe.Pointer(cmodel))
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromCaffe(cprototxt, cmodel))}
}

// ReadNetFromCaffeBytes reads a network model stored in Caffe model in memory.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga946b342af1355185a7107640f868b64a
//
func ReadNetFromCaffeBytes(prototxt []byte, caffeModel []byte) (Net, error) {
	bPrototxt, err := toByteArray(prototxt)
	if err != nil {
		return Net{}, err
	}
	bCaffeModel, err := toByteArray(caffeModel)
	if err != nil {
		return Net{}, err
	}
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromCaffeBytes(*bPrototxt, *bCaffeModel))}, nil
}

// ReadNetFromTensorflow reads a network model stored in Tensorflow framework's format.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#gad820b280978d06773234ba6841e77e8d
//
func ReadNetFromTensorflow(model string) Net {
	cmodel := C.CString(model)
	defer C.free(unsafe.Pointer(cmodel))
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromTensorflow(cmodel))}
}

// ReadNetFromTensorflowBytes reads a network model stored in Tensorflow framework's format.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#gacdba30a7c20db2788efbf5bb16a7884d
//
func ReadNetFromTensorflowBytes(model []byte) (Net, error) {
	bModel, err := toByteArray(model)
	if err != nil {
		return Net{}, err
	}
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromTensorflowBytes(*bModel))}, nil
}

// BlobFromImage creates 4-dimensional blob from image. Optionally resizes and crops
// image from center, subtract mean values, scales values by scalefactor,
// swap Blue and Red channels.
//
// For further details, please see:
// https://docs.opencv.org/trunk/d6/d0f/group__dnn.html#ga152367f253c81b53fe6862b299f5c5cd
//
func BlobFromImage(img Mat, scaleFactor float64, size image.Point, mean Scalar,
	swapRB bool, crop bool) Mat {

	sz := C.struct_Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}

	sMean := C.struct_Scalar{
		val1: C.double(mean.Val1),
		val2: C.double(mean.Val2),
		val3: C.double(mean.Val3),
		val4: C.double(mean.Val4),
	}

	return newMat(C.Net_BlobFromImage(img.p, C.double(scaleFactor), sz, sMean, C.bool(swapRB), C.bool(crop)))
}

// GetBlobChannel extracts a single (2d)channel from a 4 dimensional blob structure
// (this might e.g. contain the results of a SSD or YOLO detection,
//  a bones structure from pose detection, or a color plane from Colorization)
//
func GetBlobChannel(blob Mat, imgidx int, chnidx int) Mat {
	return newMat(C.Net_GetBlobChannel(blob.p, C.int(imgidx), C.int(chnidx)))
}

// GetBlobSize retrieves the 4 dimensional size information in (N,C,H,W) order
//
func GetBlobSize(blob Mat) Scalar {
	s := C.Net_GetBlobSize(blob.p)
	return NewScalar(float64(s.val1), float64(s.val2), float64(s.val3), float64(s.val4))
}

// Layer is a wrapper around the cv::dnn::Layer algorithm.
type Layer struct {
	// C.Layer
	p unsafe.Pointer
}

// GetLayer returns pointer to layer with specified id from the network.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a70aec7f768f38c32b1ee25f3a56526df
//
func (net *Net) GetLayer(layer int) Layer {
	return Layer{p: unsafe.Pointer(C.Net_GetLayer((C.Net)(net.p), C.int(layer)))}
}

// GetPerfProfile returns overall time for inference and timings (in ticks) for layers
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a06ce946f675f75d1c020c5ddbc78aedc
//
func (net *Net) GetPerfProfile() float64 {
	return float64(C.Net_GetPerfProfile((C.Net)(net.p)))
}

// GetUnconnectedOutLayers returns indexes of layers with unconnected outputs.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#ae62a73984f62c49fd3e8e689405b056a
//
func (net *Net) GetUnconnectedOutLayers() (ids []int) {
	cids := C.IntVector{}
	C.Net_GetUnconnectedOutLayers((C.Net)(net.p), &cids)

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cids.val)),
		Len:  int(cids.length),
		Cap:  int(cids.length),
	}
	pcids := *(*[]int)(unsafe.Pointer(h))

	for i := 0; i < int(cids.length); i++ {
		ids = append(ids, int(pcids[i]))
	}
	return
}

// Close Layer
func (l *Layer) Close() error {
	C.Layer_Close((C.Layer)(l.p))
	l.p = nil
	return nil
}

// GetName returns name for this layer.
func (l *Layer) GetName() string {
	return C.GoString(C.Layer_GetName((C.Layer)(l.p)))
}

// GetType returns type for this layer.
func (l *Layer) GetType() string {
	return C.GoString(C.Layer_GetType((C.Layer)(l.p)))
}

// InputNameToIndex returns index of input blob in input array.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d6c/classcv_1_1dnn_1_1Layer.html#a60ffc8238f3fa26cd3f49daa7ac0884b
//
func (l *Layer) InputNameToIndex(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.Layer_InputNameToIndex((C.Layer)(l.p), cName))
}

// OutputNameToIndex returns index of output blob in output array.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d6c/classcv_1_1dnn_1_1Layer.html#a60ffc8238f3fa26cd3f49daa7ac0884b
//
func (l *Layer) OutputNameToIndex(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.Layer_OutputNameToIndex((C.Layer)(l.p), cName))
}
