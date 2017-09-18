#ifndef _OPENCV3_VIDEOCAPTURE_H_
#define _OPENCV3_VIDEOCAPTURE_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "opencv3.h"

#ifdef __cplusplus
typedef cv::VideoCapture* VideoCapture;
#else
typedef void* VideoCapture;
#endif

VideoCapture VideoCapture_New();
void VideoCapture_Delete(VideoCapture v);
int VideoCapture_Open(VideoCapture v, const char* uri);
int VideoCapture_OpenDevice(VideoCapture v, int device);
void VideoCapture_Release(VideoCapture v);
void VideoCapture_Set(VideoCapture v, int prop, int param);
int VideoCapture_IsOpened(VideoCapture v);
int VideoCapture_Read(VideoCapture v, MatVec3b buf);
void VideoCapture_Grab(VideoCapture v, int skip);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEOCAPTURE_H_
