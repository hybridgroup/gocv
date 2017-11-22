#ifndef _OPENCV3_FEATURES2D_H_
#define _OPENCV3_FEATURES2D_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::AgastFeatureDetector>* AgastFeatureDetector;
typedef cv::Ptr<cv::SimpleBlobDetector>* SimpleBlobDetector;
#else
typedef void* AgastFeatureDetector;
typedef void* SimpleBlobDetector;
#endif

SimpleBlobDetector SimpleBlobDetector_Create();
void SimpleBlobDetector_Close(SimpleBlobDetector b);
struct KeyPoints SimpleBlobDetector_Detect(SimpleBlobDetector b, Mat src);

AgastFeatureDetector AgastFeatureDetector_Create();
void AgastFeatureDetector_Close(AgastFeatureDetector a);
struct KeyPoints AgastFeatureDetector_Detect(AgastFeatureDetector a, Mat src);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FEATURES2D_H_
