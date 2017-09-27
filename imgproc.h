#ifndef _OPENCV3_IMGPROC_H_
#define _OPENCV3_IMGPROC_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

void CvtColor(Mat src, Mat dst, int code);
void Rectangle(Mat img, Rect rect);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGPROC_H_
