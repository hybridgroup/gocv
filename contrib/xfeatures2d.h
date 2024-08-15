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
typedef cv::Ptr<cv::xfeatures2d::BriefDescriptorExtractor>* BriefDescriptorExtractor;
#else
typedef void* SURF;
typedef void* BriefDescriptorExtractor;
#endif

SURF SURF_Create();
SURF SURF_CreateWithParams(double hessianThreshold, int nOctaves, int nOctaveLayers, bool extended, bool upright);
void SURF_Close(SURF f);
struct KeyPoints SURF_Detect(SURF f, Mat src);
struct KeyPoints SURF_Compute(SURF f, Mat src, struct KeyPoints kp, Mat desc);
struct KeyPoints SURF_DetectAndCompute(SURF f, Mat src, Mat mask, Mat desc);
BriefDescriptorExtractor BriefDescriptorExtractor_Create();
BriefDescriptorExtractor BriefDescriptorExtractor_CreateWithParams(int bytes, bool useOrientation);
void BriefDescriptorExtractor_Close(BriefDescriptorExtractor b);
void BriefDescriptorExtractor_Compute(BriefDescriptorExtractor b, Mat src, struct KeyPoints kp, Mat desc);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XFEATURES2D_H_
