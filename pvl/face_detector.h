#ifndef _OPENCV3_FACE_DETECTOR_H_
#define _OPENCV3_FACE_DETECTOR_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/pvl.hpp>
extern "C" {
#endif

#include "../core.h"
#include "face.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::pvl::FaceDetector>* FaceDetector;
#else
typedef void* FaceDetector;
#endif

// FaceDetector
FaceDetector FaceDetector_New();
void FaceDetector_Close(FaceDetector f);
void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled);
struct Faces FaceDetector_DetectFaceRect(FaceDetector f, Mat img);
void FaceDetector_DetectEye(FaceDetector f, Mat img, Face face);
void FaceDetector_DetectMouth(FaceDetector f, Mat img, Face face);
void FaceDetector_DetectSmile(FaceDetector f, Mat img, Face face);
void FaceDetector_DetectBlink(FaceDetector f, Mat img, Face face);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FACE_DETECTOR_H_
