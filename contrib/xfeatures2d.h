#ifndef _OPENCV3_XFEATURES2D_H_
#define _OPENCV3_XFEATURES2D_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xfeatures2d.hpp>
extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::xfeatures2d::SURF>* SURF;
typedef cv::Ptr<cv::xfeatures2d::BEBLID>* BeblidDescriptorExtractor;
#else
typedef void* SURF;
typedef void* BeblidDescriptorExtractor;
#endif

SURF SURF_Create();
void SURF_Close(SURF f);
struct KeyPoints SURF_Detect(SURF f, Mat src);
struct KeyPoints SURF_DetectAndCompute(SURF f, Mat src, Mat mask, Mat desc);

BeblidDescriptorExtractor BeblidDescriptorExtractor_Create(float scaleFactor, int size);
void BeblidDescriptorExtractor_Close(BeblidDescriptorExtractor b);
void BeblidDescriptorExtractor_Compute(BeblidDescriptorExtractor b, Mat src, struct KeyPoints kp, Mat desc);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XFEATURES2D_H_
