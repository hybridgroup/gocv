#ifndef _OPENCV3_FEATURES2D_H_
#define _OPENCV3_FEATURES2D_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::AKAZE>* AKAZE;
typedef cv::Ptr<cv::AgastFeatureDetector>* AgastFeatureDetector;
typedef cv::Ptr<cv::BRISK>* BRISK;
typedef cv::Ptr<cv::FastFeatureDetector>* FastFeatureDetector;
typedef cv::Ptr<cv::ORB>* ORB;
typedef cv::Ptr<cv::SimpleBlobDetector>* SimpleBlobDetector;
#else
typedef void* AKAZE;
typedef void* AgastFeatureDetector;
typedef void* BRISK;
typedef void* FastFeatureDetector;
typedef void* ORB;
typedef void* SimpleBlobDetector;
#endif

AKAZE AKAZE_Create();
void AKAZE_Close(AKAZE a);
struct KeyPoints AKAZE_Detect(AKAZE a, Mat src);
struct KeyPoints AKAZE_DetectAndCompute(AKAZE a, Mat src, Mat mask, Mat desc);

AgastFeatureDetector AgastFeatureDetector_Create();
void AgastFeatureDetector_Close(AgastFeatureDetector a);
struct KeyPoints AgastFeatureDetector_Detect(AgastFeatureDetector a, Mat src);

BRISK BRISK_Create();
void BRISK_Close(BRISK b);
struct KeyPoints BRISK_Detect(BRISK b, Mat src);
struct KeyPoints BRISK_DetectAndCompute(BRISK b, Mat src, Mat mask, Mat desc);

FastFeatureDetector FastFeatureDetector_Create();
void FastFeatureDetector_Close(FastFeatureDetector f);
struct KeyPoints FastFeatureDetector_Detect(FastFeatureDetector f, Mat src);

ORB ORB_Create();
void ORB_Close(ORB o);
struct KeyPoints ORB_Detect(ORB o, Mat src);
struct KeyPoints ORB_DetectAndCompute(ORB o, Mat src, Mat mask, Mat desc);

SimpleBlobDetector SimpleBlobDetector_Create();
void SimpleBlobDetector_Close(SimpleBlobDetector b);
struct KeyPoints SimpleBlobDetector_Detect(SimpleBlobDetector b, Mat src);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FEATURES2D_H_
