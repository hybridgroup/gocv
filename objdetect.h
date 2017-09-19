#ifndef _OPENCV3_OBJDETECT_H_
#define _OPENCV3_OBJDETECT_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::CascadeClassifier* CascadeClassifier;
#else
typedef void* CascadeClassifier;
#endif

// CascadeClassifier
CascadeClassifier CascadeClassifier_New();
void CascadeClassifier_Delete(CascadeClassifier cs);
int CascadeClassifier_Load(CascadeClassifier cs, const char* name);
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, MatVec3b img);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_OBJDETECT_H_
