#ifndef _OPENCV3_NMS_H_
#define _OPENCV3_NMS_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/dnn/dnn.hpp>
extern "C" {
#endif

#include "core.h"

void NMSBoxes(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector indices);
void NMSBoxesWithParams(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector indices, const float eta, const int top_k);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_NMS_H