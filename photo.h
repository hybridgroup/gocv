#ifndef _OPENCV3_PHOTO_H_
#define _OPENCV3_PHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>

extern "C" {
#endif

#include "core.h"

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PHOTO_H