#ifndef _OPENVINO_IE_H_
#define _OPENVINO_IE_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <inference_engine.hpp>
extern "C" {
#endif

//#include "../../core.h"

#ifdef __cplusplus
typedef InferenceEngine::InferenceEnginePluginPtr InferenceEnginePluginPtr;
#else
typedef void* InferenceEnginePluginPtr;
#endif

// InferencePlugin
InferenceEnginePluginPtr InferenceEnginePluginPtr_New(const char* libpath);
void InferenceEnginePluginPtr_Close(InferenceEnginePluginPtr p);

#ifdef __cplusplus
}
#endif

#endif //_OPENVINO_IE_H_
