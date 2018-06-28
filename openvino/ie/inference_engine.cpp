#include "inference_engine.h"

const char* OpenVinoVersion() {
    std::ostringstream res;
    res << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.major) 
        << "." 
        << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.minor)
        << "." 
        << InferenceEngine::GetInferenceEngineVersion()->buildNumber;
    return res.str().c_str();
}
