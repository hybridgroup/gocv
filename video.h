#ifndef _OPENCV3_VIDEO_H_
#define _OPENCV3_VIDEO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/core/ocl.hpp>
#include <opencv2/bgsegm.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::BackgroundSubtractor* BackgroundSubtractor;
#else
typedef void* BackgroundSubtractor;
#endif

BackgroundSubtractor BackgroundSubtractor_CreateMOG2();
void BackgroundSubtractor_Close(BackgroundSubtractor b);
void BackgroundSubtractor_Apply(BackgroundSubtractor b, Mat src, Mat dst);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEO_H_
