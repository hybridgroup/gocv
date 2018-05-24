#ifndef _OPENVINO_IE_H_
#define _OPENVINO_IE_H_

#ifdef __cplusplus
#include <inference_engine.hpp>
extern "C" {
#endif

//#include "../../core.h"

#ifdef __cplusplus
typedef InferenceEngine::InferenceEnginePluginPtr InferenceEnginePluginPtr;
typedef InferenceEngine::CNNNetReader CNNNetReader;
typedef InferenceEngine::CNNNetwork CNNNetwork;
#else
typedef void* InferenceEnginePluginPtr;
typedef void* CNNNetReader;
typedef void* CNNNetwork;
#endif

const char* IEVersion();

// InferencePlugin
InferenceEnginePluginPtr InferenceEnginePluginPtr_New(const char* libpath);
void InferenceEnginePluginPtr_Close(InferenceEnginePluginPtr p);

CNNNetReader* CNNNetReader_New();
void CNNNetReader_ReadNetwork(CNNNetReader* r, const char* cModelFile);
void CNNNetReader_ReadWeights(CNNNetReader* r, const char* cWeightsFile);
CNNNetwork* CNNNetReader_GetNetwork(CNNNetReader* r);

//CNNNetwork CNNNetwork_New();

#ifdef __cplusplus
}
#endif

#endif //_OPENVINO_IE_H_
