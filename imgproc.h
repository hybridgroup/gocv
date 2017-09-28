#ifndef _OPENCV3_IMGPROC_H_
#define _OPENCV3_IMGPROC_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

void CvtColor(Mat src, Mat dst, int code);
void Rectangle(Mat img, Rect rect, Scalar color);
struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness);
void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale, 
             Scalar color, int thickness);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGPROC_H_
