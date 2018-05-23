#include "inference_engine.h"

// InferencePlugin
InferenceEnginePluginPtr InferenceEnginePluginPtr_New(const char* libpath) {
    return InferenceEnginePluginPtr(InferenceEngine::PluginDispatcher({libpath}).getPluginByDevice("CPU"));
}

void InferenceEnginePluginPtr_Close(InferenceEnginePluginPtr p) {
    //p->Release();
}
