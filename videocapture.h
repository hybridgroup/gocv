#ifndef _OPENCV3_VIDEOCAPTURE_H_
#define _OPENCV3_VIDEOCAPTURE_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "opencv3.h"

#ifdef __cplusplus
typedef cv::VideoCapture* VideoCapture;
typedef cv::VideoWriter* VideoWriter;
#else
typedef void* VideoCapture;
typedef void* VideoWriter;
#endif

// VideoCapture
VideoCapture VideoCapture_New();
void VideoCapture_Delete(VideoCapture v);
int VideoCapture_Open(VideoCapture v, const char* uri);
int VideoCapture_OpenDevice(VideoCapture v, int device);
void VideoCapture_Release(VideoCapture v);
void VideoCapture_Set(VideoCapture v, int prop, int param);
int VideoCapture_IsOpened(VideoCapture v);
int VideoCapture_Read(VideoCapture v, MatVec3b buf);
void VideoCapture_Grab(VideoCapture v, int skip);

// VideoWriter
VideoWriter VideoWriter_New();
void VideoWriter_Delete(VideoWriter vw);
void VideoWriter_Open(VideoWriter vw, const char* name, double fps, int width,
  int height);
void VideoWriter_OpenWithMat(VideoWriter vw, const char* name, double fps,
  MatVec3b img);
int VideoWriter_IsOpened(VideoWriter vw);
void VideoWriter_Write(VideoWriter vw, MatVec3b img);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEOCAPTURE_H_
