#ifndef _OPENCV3_PVL_H_
#define _OPENCV3_PVL_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/pvl.hpp>
extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::pvl::Face* Face;
typedef cv::Ptr<cv::pvl::FaceDetector> FaceDetector;
#else
typedef void* Face;
typedef void* FaceDetector;
#endif

// Face
Face Face_New();
void Face_Delete(Face f);
Rect Face_GetRect(Face f);

// FaceDetector
FaceDetector FaceDetector_New();
void FaceDetector_Delete(FaceDetector f);
void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PVL_H_
