#ifndef _GOCVOPENVINO_IE_H_
#define _GOCVOPENVINO_IE_H_

#ifdef __cplusplus
#include <inference_engine.hpp>
extern "C" {
#endif

//#include "../../core.h"

#ifdef __cplusplus
typedef InferenceEngine::PluginDispatcher InferenceEnginePluginDispatcher;
typedef InferenceEngine::InferenceEnginePluginPtr InferenceEnginePluginPtr;
typedef InferenceEngine::CNNNetReader CNNNetReader;
typedef InferenceEngine::CNNNetwork CNNNetwork;
#else
typedef void* InferenceEnginePluginDispatcher;
typedef void* InferenceEnginePluginPtr;
typedef void* CNNNetReader;
typedef void* CNNNetwork;
#endif

const char* IEVersion();

// InferenceEnginePluginDispatcher
InferenceEnginePluginDispatcher* InferenceEnginePluginDispatcher_New(const char* libpath);
void InferenceEnginePluginDispatcher_Close(InferenceEnginePluginDispatcher p);
InferenceEnginePluginPtr* InferenceEnginePluginDispatcher_GetPluginByDevice(InferenceEnginePluginDispatcher* p, const char* device);

// InferencePlugin
void InferenceEnginePluginPtr_Close(InferenceEnginePluginPtr p);

CNNNetReader* CNNNetReader_New();
void CNNNetReader_ReadNetwork(CNNNetReader* r, const char* cModelFile);
void CNNNetReader_ReadWeights(CNNNetReader* r, const char* cWeightsFile);
CNNNetwork* CNNNetReader_GetNetwork(CNNNetReader* r);

size_t CNNNetwork_Size(CNNNetwork* n);

#ifdef __cplusplus
}
#endif

#endif //_GOCVOPENVINO_IE_H_
