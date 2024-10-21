#ifndef _OPENCV3_VIDEOIO_H_
#define _OPENCV3_VIDEOIO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/videoio/registry.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::VideoCapture* VideoCapture;
typedef cv::VideoWriter* VideoWriter;
#else
typedef void* VideoCapture;
typedef void* VideoWriter;
#endif

// VideoCapture
VideoCapture VideoCapture_New();
void VideoCapture_Close(VideoCapture v);
bool VideoCapture_Open(VideoCapture v, const char* uri);
bool VideoCapture_OpenWithAPI(VideoCapture v, const char* uri, int apiPreference);
bool VideoCapture_OpenWithAPIParams(VideoCapture v, const char* uri, int apiPreference, int *paramsv, int paramsc);
bool VideoCapture_OpenDevice(VideoCapture v, int device);
bool VideoCapture_OpenDeviceWithAPI(VideoCapture v, int device, int apiPreference);
bool VideoCapture_OpenDeviceWithAPIParams(VideoCapture v, int device, int apiPreference, int *paramsv, int paramsc);
void VideoCapture_Set(VideoCapture v, int prop, double param);
double VideoCapture_Get(VideoCapture v, int prop);
int VideoCapture_IsOpened(VideoCapture v);
int VideoCapture_Read(VideoCapture v, Mat buf);
void VideoCapture_Grab(VideoCapture v, int skip);
int VideoCapture_Retrieve(VideoCapture v, Mat buf);

// VideoWriter
VideoWriter VideoWriter_New();
void VideoWriter_Close(VideoWriter vw);
void VideoWriter_Open(VideoWriter vw, const char* name, const char* codec, double fps, int width,
                      int height, bool isColor);
void VideoWriter_OpenWithAPI(VideoWriter vw, const char* name, int apiPreference, const char* codec, double fps,
                      int width, int height, bool isColor);
void VideoWriter_OpenWithAPIParams(VideoWriter vw, const char* name, int apiPreference, const char* codec, double fps,
                      int width, int height, IntVector params);

int VideoWriter_IsOpened(VideoWriter vw);
void VideoWriter_Write(VideoWriter vw, Mat img);

//Videoio Query I/O API backends registry
char* Videoio_Registry_GetBackendName(int api);
IntVector Videio_Registry_GetBackends();
char* Videoio_Registry_GetCameraBackendPluginVersion(int api, int* version_ABI, int* version_API);
IntVector Videoio_Registry_GetCameraBackends();
char* Videoio_Registry_GetStreamBackendPluginVersion(int api, int* version_ABI, int* version_API);
IntVector Videoio_Registry_GetStreamBackends();
char* Videoio_Registry_GetWriterBackendPluginVersion(int api, int* version_ABI, int* version_API);
IntVector Videoio_Registry_GetWriterBackends();
bool Videoio_Registry_HasBackend(int api);
bool Videoio_Registry_IsBackendBuiltIn(int api);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEOIO_H_
