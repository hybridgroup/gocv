package ie

/*
#include <stdlib.h>
#include "inference_engine.h"
*/
import (
	"C"
)
import (
	"os"
	"unsafe"
)

// Version returns the current Inference Engine library version
func Version() string {
	return C.GoString(C.IEVersion())
}

// InferenceEnginePlugin is a wrapper around InferenceEngine::InferenceEnginePluginPtr.
type InferenceEnginePlugin struct {
	p C.InferenceEnginePluginPtr
}

// NewInferenceEnginePlugin returns a new OpenVINO InferencePlugin.
func NewInferenceEnginePlugin() InferenceEnginePlugin {
	libPath := os.Getenv("INTEL_CVSDK_DIR") + "/deployment_tools/inference_engine/lib/ubuntu_16.04/intel64"
	cLibPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cLibPath))

	return InferenceEnginePlugin{p: C.InferenceEnginePluginPtr_New(cLibPath)}
}

// Close InferenceEnginePlugin.
func (pd InferenceEnginePlugin) Close() error {
	//C.FaceDetector_Close((C.FaceDetector)(f.p))
	pd.p = nil
	return nil
}

// CNNNetReader is a wrapper around InferenceEngine::CNNNetReader.
type CNNNetReader struct {
	// C.CNNNetReader
	p unsafe.Pointer
}

// NewCNNNetReader returns a new OpenVINO CNNNetReader.
func NewCNNNetReader() CNNNetReader {
	return CNNNetReader{p: unsafe.Pointer(C.CNNNetReader_New())}
}

// Close CNNNetReader.
func (r *CNNNetReader) Close() error {
	r.p = nil
	return nil
}

// ReadNetwork reads a IR model into a CNNNetReader.
func (r *CNNNetReader) ReadNetwork(modelFile string) error {
	cModelFile := C.CString(modelFile)
	defer C.free(unsafe.Pointer(cModelFile))

	C.CNNNetReader_ReadNetwork((*C.CNNNetReader)(r.p), cModelFile)
	return nil
}

// ReadWeights reads IR bin file into a CNNNetReader.
func (r *CNNNetReader) ReadWeights(weightsFile string) error {
	cWeightsFile := C.CString(weightsFile)
	defer C.free(unsafe.Pointer(cWeightsFile))

	C.CNNNetReader_ReadWeights((*C.CNNNetReader)(r.p), cWeightsFile)
	return nil
}

// GetNetwork returns a the CNNNetwork that has been loaded up using the CNNNetReader.
func (r *CNNNetReader) GetNetwork() CNNNetwork {
	return CNNNetwork{p: unsafe.Pointer(C.CNNNetReader_GetNetwork((*C.CNNNetReader)(r.p)))}
}

// CNNNetwork is a wrapper around InferenceEngine::CNNNetwork.
type CNNNetwork struct {
	// C.CNNNetwork
	p unsafe.Pointer
}

// NewCNNNetwork returns a new OpenVINO CNNNetwork.
func NewCNNNetwork() CNNNetwork {
	return CNNNetwork{} //CNNNetwork{p: unsafe.Pointer(C.CNNNetwork_New(cModelFile, cWeightsFile))}
}

// Close CNNNetwork.
func (n CNNNetwork) Close() error {
	n.p = nil
	return nil
}
