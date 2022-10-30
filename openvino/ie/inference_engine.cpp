#include "inference_engine.h"

const char* OpenVinoVersion() {
    std::ostringstream buf;
    buf << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.major) 
        << "." 
        << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.minor)
        << "." 
        << InferenceEngine::GetInferenceEngineVersion()->buildNumber;
    auto version = buf.str();

    size_t resLen = version.size() + 1;
    auto res = (char*)malloc(resLen);

    memset(res ,0, resLen);
    memcpy(res, version.c_str(), resLen);
    return res;
}
