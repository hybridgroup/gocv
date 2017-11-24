#ifndef _OPENCV3_VIDEO_H_
#define _OPENCV3_VIDEO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::BackgroundSubtractor> BackgroundSubtractor;
#else
typedef void* BackgroundSubtractor;
#endif

BackgroundSubtractor BackgroundSubtractor_CreateMOG2();
BackgroundSubtractor BackgroundSubtractor_CreateKNN();
void BackgroundSubtractor_Close(BackgroundSubtractor b);
void BackgroundSubtractor_Apply(BackgroundSubtractor b, Mat src, Mat dst);

void CalcOpticalFlowPyrLK(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err);
void CalcOpticalFlowFarneback(Mat prevImg, Mat nextImg, Mat flow, double pyrScale, int levels, int winsize,
	int iterations, int polyN, double polySigma, int flags);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_VIDEO_H_
