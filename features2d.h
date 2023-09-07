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
typedef cv::Ptr<cv::GFTTDetector>* GFTTDetector;
typedef cv::Ptr<cv::KAZE>* KAZE;
typedef cv::Ptr<cv::MSER>* MSER;
typedef cv::Ptr<cv::ORB>* ORB;
typedef cv::Ptr<cv::SimpleBlobDetector>* SimpleBlobDetector;
typedef cv::Ptr<cv::BFMatcher>* BFMatcher;
typedef cv::Ptr<cv::FlannBasedMatcher>* FlannBasedMatcher;
typedef cv::Ptr<cv::SIFT>* SIFT;
#else
typedef void* AKAZE;
typedef void* AgastFeatureDetector;
typedef void* BRISK;
typedef void* FastFeatureDetector;
typedef void* GFTTDetector;
typedef void* KAZE;
typedef void* MSER;
typedef void* ORB;
typedef void* SimpleBlobDetector;
typedef void* BFMatcher;
typedef void* FlannBasedMatcher;
typedef void* SIFT;
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
FastFeatureDetector FastFeatureDetector_CreateWithParams(int threshold, bool nonmaxSuppression, int type);
void FastFeatureDetector_Close(FastFeatureDetector f);
struct KeyPoints FastFeatureDetector_Detect(FastFeatureDetector f, Mat src);

GFTTDetector GFTTDetector_Create();
void GFTTDetector_Close(GFTTDetector a);
struct KeyPoints GFTTDetector_Detect(GFTTDetector a, Mat src);

KAZE KAZE_Create();
void KAZE_Close(KAZE a);
struct KeyPoints KAZE_Detect(KAZE a, Mat src);
struct KeyPoints KAZE_DetectAndCompute(KAZE a, Mat src, Mat mask, Mat desc);

MSER MSER_Create();
void MSER_Close(MSER a);
struct KeyPoints MSER_Detect(MSER a, Mat src);

ORB ORB_Create();
ORB ORB_CreateWithParams(int nfeatures, float scaleFactor, int nlevels, int edgeThreshold, int firstLevel, int WTA_K, int scoreType, int patchSize, int fastThreshold);
void ORB_Close(ORB o);
struct KeyPoints ORB_Detect(ORB o, Mat src);
struct KeyPoints ORB_DetectAndCompute(ORB o, Mat src, Mat mask, Mat desc);

SimpleBlobDetector SimpleBlobDetector_Create();
SimpleBlobDetector SimpleBlobDetector_Create_WithParams(SimpleBlobDetectorParams params);
void SimpleBlobDetector_Close(SimpleBlobDetector b);
struct KeyPoints SimpleBlobDetector_Detect(SimpleBlobDetector b, Mat src);
SimpleBlobDetectorParams SimpleBlobDetectorParams_Create();

BFMatcher BFMatcher_Create();
BFMatcher BFMatcher_CreateWithParams(int normType, bool crossCheck);
void BFMatcher_Close(BFMatcher b);
struct DMatches BFMatcher_Match(BFMatcher b, Mat query, Mat train);
struct MultiDMatches BFMatcher_KnnMatch(BFMatcher b, Mat query, Mat train, int k);

FlannBasedMatcher FlannBasedMatcher_Create();
void FlannBasedMatcher_Close(FlannBasedMatcher f);
struct MultiDMatches FlannBasedMatcher_KnnMatch(FlannBasedMatcher f, Mat query, Mat train, int k);

void DrawKeyPoints(Mat src, struct KeyPoints kp, Mat dst, const Scalar s, int flags);

SIFT SIFT_Create();
void SIFT_Close(SIFT f);
struct KeyPoints SIFT_Detect(SIFT f, Mat src);
struct KeyPoints SIFT_DetectAndCompute(SIFT f, Mat src, Mat mask, Mat desc);

void DrawMatches(Mat img1, struct KeyPoints kp1, Mat img2, struct KeyPoints kp2, struct DMatches matches1to2, Mat outImg, const Scalar matchesColor, const Scalar pointColor, struct ByteArray matchesMask, int flags);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FEATURES2D_H_
