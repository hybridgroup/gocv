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

// InferenceEnginePluginDispatcher is a wrapper around InferenceEngine::InferenceEnginePluginDispatcher.
type InferenceEnginePluginDispatcher struct {
	// C.InferenceEnginePluginDispatcher
	p unsafe.Pointer
}

// DefaultLibPath returns the usual lib path for Ubuntu
func DefaultLibPath() string {
	return os.Getenv("INTEL_CVSDK_DIR") + "/deployment_tools/inference_engine/lib/ubuntu_16.04/intel64"
}

// NewInferenceEnginePluginDispatcher returns a new OpenVINO InferenceEnginePluginDispatcher.
func NewInferenceEnginePluginDispatcher(libpath string) InferenceEnginePluginDispatcher {
	cLibPath := C.CString(libpath)
	defer C.free(unsafe.Pointer(cLibPath))

	pd := C.InferenceEnginePluginDispatcher_New(cLibPath)
	return InferenceEnginePluginDispatcher{p: unsafe.Pointer(pd)}
}

// Close InferenceEnginePluginDispatcher.
func (pd InferenceEnginePluginDispatcher) Close() error {
	pd.p = nil
	return nil
}

// GetPluginByDevice from InferenceEnginePluginDispatcher.
func (pd InferenceEnginePluginDispatcher) GetPluginByDevice(device string) InferenceEnginePlugin {
	cDevice := C.CString(device)
	defer C.free(unsafe.Pointer(cDevice))

	pu := C.InferenceEnginePluginDispatcher_GetPluginByDevice((*C.InferenceEnginePluginDispatcher)(pd.p), cDevice)
	return InferenceEnginePlugin{p: unsafe.Pointer(pu)}
}

// InferenceEnginePlugin is a wrapper around InferenceEngine::InferenceEnginePluginPtr.
type InferenceEnginePlugin struct {
	// C.InferenceEnginePluginPtr
	p unsafe.Pointer
}

// Close InferenceEnginePlugin.
func (pu InferenceEnginePlugin) Close() error {
	pu.p = nil
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

// Size returns the number of nodes in this CNNNetwork.
func (n CNNNetwork) Size() uint32 {
	return uint32(C.CNNNetwork_Size((*C.CNNNetwork)(n.p)))
}
