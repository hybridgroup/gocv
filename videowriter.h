#ifndef _OPENCV3_VIDEOWRITER_H_
#define _OPENCV3_VIDEOWRITER_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "opencv3.h"

#ifdef __cplusplus
typedef cv::VideoWriter* VideoWriter;
#else
typedef void* VideoWriter;
#endif

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

#endif //_OPENCV3_VIDEOWRITER_H_
