#ifndef _OPENCV3_XOBJDETECT_H_
#define _OPENCV3_XOBJDETECT_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xobjdetect.hpp>
extern "C" {
#endif

#include "../core.h"
#include "../persistence.h"

typedef struct WBDetector_detect_result_t {
    Rects       bboxes;
    FloatVector confidences;
} WBDetector_detect_result;

#ifdef __cplusplus
typedef cv::Ptr<cv::xobjdetect::WBDetector>* WBDetector;
#else
typedef void* WBDetector;
#endif

WBDetector WBDetector_Create();
void WBDetector_Close(WBDetector det);
WBDetector_detect_result WBDetector_Detect(WBDetector det, Mat img);
void WBDetector_Read(WBDetector det, FileNode node);
void WBDetector_Train(WBDetector det, const char* pos_samples, const char* neg_imgs);
void WBDetector_Write(WBDetector det, FileStorage fs);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XOBJDETECT_H_
