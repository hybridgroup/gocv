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

InferenceEnginePluginDispatcher* InferenceEnginePluginDispatcher_New(const char* libpath) {
    return new InferenceEnginePluginDispatcher(InferenceEngine::PluginDispatcher({libpath, ""}));
}

void InferenceEnginePluginDispatcher_Close(InferenceEnginePluginDispatcher p) {
}

InferenceEnginePluginPtr* InferenceEnginePluginDispatcher_GetPluginByDevice(InferenceEnginePluginDispatcher* p, 
                                                                            const char* device) {
    return new InferenceEnginePluginPtr(p->getPluginByDevice(device));
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
size_t CNNNetwork_Size(CNNNetwork* n) {
    return n->size();
}
