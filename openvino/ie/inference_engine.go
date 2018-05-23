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

// InferenceEnginePluginPtr is a wrapper around InferenceEngine::InferenceEnginePluginPtr.
type InferenceEnginePluginPtr struct {
	p C.InferenceEnginePluginPtr
}

// NewInferenceEnginePluginPtr returns a new OpenVINO InferencePlugin.
func NewInferenceEnginePluginPtr() InferenceEnginePluginPtr {
	libPath := os.Getenv("INTEL_CVSDK_DIR") + "/deployment_tools/inference_engine/lib/ubuntu_16.04/intel64"
	cLibPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cLibPath))

	return InferenceEnginePluginPtr{p: C.InferenceEnginePluginPtr_New(cLibPath)}
}

// Close FaceDetector.
func (pd InferenceEnginePluginPtr) Close() error {
	//C.FaceDetector_Close((C.FaceDetector)(f.p))
	pd.p = nil
	return nil
}
