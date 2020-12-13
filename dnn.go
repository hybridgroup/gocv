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

	// NetBackendCUDA is the Cuda backend.
	NetBackendCUDA NetBackendType = 5
)

// ParseNetBackend returns a valid NetBackendType given a string. Valid values are:
// - halide
// - openvino
// - opencv
// - vulkan
// - cuda
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
	case "cuda":
		return NetBackendCUDA
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

	// NetTargetCUDA is the CUDA target.
	NetTargetCUDA NetTargetType = 6

	// NetTargetCUDAFP16 is the CUDA target.
	NetTargetCUDAFP16 NetTargetType = 7
)

// ParseNetTarget returns a valid NetTargetType given a string. Valid values are:
// - cpu
// - fp32
// - fp16
// - vpu
// - vulkan
// - fpga
// - cuda
// - cudafp16
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
	case "cuda":
		return NetTargetCUDA
	case "cudafp16":
		return NetTargetCUDAFP16
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
		addMatToProfile(blobs[i].p)
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

// ReadNetFromTorch reads a network model stored in Torch framework's format (t7).
//   check net.Empty() for read failure
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#gaaaed8c8530e9e92fe6647700c13d961e
//
func ReadNetFromTorch(model string) Net {
	cmodel := C.CString(model)
	defer C.free(unsafe.Pointer(cmodel))
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromTorch(cmodel))}
}

// ReadNetFromONNX reads a network model stored in ONNX framework's format.
//   check net.Empty() for read failure
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga7faea56041d10c71dbbd6746ca854197
//
func ReadNetFromONNX(model string) Net {
	cmodel := C.CString(model)
	defer C.free(unsafe.Pointer(cmodel))
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromONNX(cmodel))}
}

// ReadNetFromONNXBytes reads a network model stored in ONNX framework's format.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga9198ecaac7c32ddf0aa7a1bcbd359567
//
func ReadNetFromONNXBytes(model []byte) (Net, error) {
	bModel, err := toByteArray(model)
	if err != nil {
		return Net{}, err
	}
	return Net{p: unsafe.Pointer(C.Net_ReadNetFromONNXBytes(*bModel))}, nil
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

// BlobFromImages Creates 4-dimensional blob from series of images.
// Optionally resizes and crops images from center, subtract mean values,
// scales values by scalefactor, swap Blue and Red channels.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga2b89ed84432e4395f5a1412c2926293c
//
func BlobFromImages(imgs []Mat, blob *Mat, scaleFactor float64, size image.Point, mean Scalar,
	swapRB bool, crop bool, ddepth MatType) {

	cMatArray := make([]C.Mat, len(imgs))
	for i, r := range imgs {
		cMatArray[i] = r.p
	}

	cMats := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(imgs)),
	}

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

	C.Net_BlobFromImages(cMats, blob.p, C.double(scaleFactor), sz, sMean, C.bool(swapRB), C.bool(crop), C.int(ddepth))
}

// ImagesFromBlob Parse a 4D blob and output the images it contains as
// 2D arrays through a simpler data structure (std::vector<cv::Mat>).
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga4051b5fa2ed5f54b76c059a8625df9f5
//
func ImagesFromBlob(blob Mat, imgs []Mat) {
	cMats := C.struct_Mats{}
	C.Net_ImagesFromBlob(blob.p, &(cMats))
	// mv = make([]Mat, cMats.length)
	for i := C.int(0); i < cMats.length; i++ {
		imgs[i].p = C.Mats_get(cMats, i)
	}
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
	defer C.free(unsafe.Pointer(cids.val))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cids.val)),
		Len:  int(cids.length),
		Cap:  int(cids.length),
	}
	pcids := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(cids.length); i++ {
		ids = append(ids, int(pcids[i]))
	}
	return
}

// GetLayerNames returns all layer names.
//
// For furtherdetails, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#ae8be9806024a0d1d41aba687cce99e6b
//
func (net *Net) GetLayerNames() (names []string) {
	cstrs := C.CStrings{}
	defer C.CStrings_Close(cstrs)
	C.Net_GetLayerNames((C.Net)(net.p), &cstrs)
	return toGoStrings(cstrs)
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

// NMSBoxes performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
func NMSBoxes(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int) {
	bboxesRectArr := []C.struct_Rect{}
	for _, v := range bboxes {
		bbox := C.struct_Rect{
			x:      C.int(v.Min.X),
			y:      C.int(v.Min.Y),
			width:  C.int(v.Size().X),
			height: C.int(v.Size().Y),
		}
		bboxesRectArr = append(bboxesRectArr, bbox)
	}

	bboxesRects := C.Rects{
		rects:  (*C.Rect)(&bboxesRectArr[0]),
		length: C.int(len(bboxes)),
	}

	scoresFloats := []C.float{}
	for _, v := range scores {
		scoresFloats = append(scoresFloats, C.float(v))
	}
	scoresVector := C.struct_FloatVector{}
	scoresVector.val = (*C.float)(&scoresFloats[0])
	scoresVector.length = (C.int)(len(scoresFloats))

	indicesVector := C.IntVector{}

	C.NMSBoxes(bboxesRects, scoresVector, C.float(scoreThreshold), C.float(nmsThreshold), &indicesVector)
	defer C.free(unsafe.Pointer(indicesVector.val))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(indicesVector.val)),
		Len:  int(indicesVector.length),
		Cap:  int(indicesVector.length),
	}

	ptr := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(indicesVector.length); i++ {
		indices[i] = int(ptr[i])
	}
	return
}

// NMSBoxesWithParams performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
func NMSBoxesWithParams(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int, eta float32, topK int) {
	bboxesRectArr := []C.struct_Rect{}
	for _, v := range bboxes {
		bbox := C.struct_Rect{
			x:      C.int(v.Min.X),
			y:      C.int(v.Min.Y),
			width:  C.int(v.Size().X),
			height: C.int(v.Size().Y),
		}
		bboxesRectArr = append(bboxesRectArr, bbox)
	}

	bboxesRects := C.Rects{
		rects:  (*C.Rect)(&bboxesRectArr[0]),
		length: C.int(len(bboxes)),
	}

	scoresFloats := []C.float{}
	for _, v := range scores {
		scoresFloats = append(scoresFloats, C.float(v))
	}
	scoresVector := C.struct_FloatVector{}
	scoresVector.val = (*C.float)(&scoresFloats[0])
	scoresVector.length = (C.int)(len(scoresFloats))

	indicesVector := C.IntVector{}

	C.NMSBoxesWithParams(bboxesRects, scoresVector, C.float(scoreThreshold), C.float(nmsThreshold), &indicesVector, C.float(eta), C.int(topK))
	defer C.free(unsafe.Pointer(indicesVector.val))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(indicesVector.val)),
		Len:  int(indicesVector.length),
		Cap:  int(indicesVector.length),
	}

	ptr := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(indicesVector.length); i++ {
		indices[i] = int(ptr[i])
	}
	return
}
