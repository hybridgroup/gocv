#include "inference_engine.h"

const char* IEVersion() {
    std::ostringstream res;
    res << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.major) 
        << "." 
        << std::to_string(InferenceEngine::GetInferenceEngineVersion()->apiVersion.minor)
        << "." 
        << InferenceEngine::GetInferenceEngineVersion()->buildNumber;
    return res.str().c_str();
}

// InferencePlugin
InferenceEnginePluginPtr InferenceEnginePluginPtr_New(const char* libpath) {
    return InferenceEnginePluginPtr(InferenceEngine::PluginDispatcher({libpath, ""}).getPluginByDevice("CPU"));
}

void InferenceEnginePluginPtr_Close(InferenceEnginePluginPtr p) {
    //p->Release();
}

// CNNNetReader
CNNNetReader* CNNNetReader_New() {
    return new InferenceEngine::CNNNetReader();
}

void CNNNetReader_ReadNetwork(CNNNetReader* r, const char* cModelFile) {
    r->ReadNetwork(cModelFile);
}

void CNNNetReader_ReadWeights(CNNNetReader* r, const char* cWeightsFile) {
    r->ReadWeights(cWeightsFile);
}

CNNNetwork* CNNNetReader_GetNetwork(CNNNetReader* r) {
    return new CNNNetwork(r->getNetwork());
}

// CNNNetwork

// CNNNetwork_New(const char* cModelFile, const char* cWeightsFile) {
//     printf("%s\n", cModelFile);
//     printf("%s\n", cWeightsFile);

//     InferenceEngine::CNNNetReader r;
    

//     r.getNetwork().setBatchSize(1);

//     //r.ReadWeights(cWeightsFile);
//     return r    
// }
