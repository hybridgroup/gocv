#ifndef _OPENCV3_FACE_H_
#define _OPENCV3_FACE_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/face.hpp>

extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::face::Facemark>* LBPHFaceMark;
#else
typedef void* LBPHFaceMark;
#endif

struct PredictResponse {
    int label;
    double confidence;
};

LBPHFaceMark CreateLBPHFaceMark();
void LBPHFaceMark_LoadModel(LBPHFaceMark fm, const char*  model); //Points2fVector
bool LBPHFaceMark_Fit(LBPHFaceMark fm, Mat frame, struct Rects faces, Points2fVector landmarks);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_FACE_H_