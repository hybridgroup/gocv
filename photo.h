#ifndef _OPENCV3_PHOTO_H_
#define _OPENCV3_PHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>

extern "C" {
#endif

#include "core.h"

void ColorChange(Mat src, Mat mask, Mat dst, float red_mul, float green_mul, float blue_mul);

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags);

void IlluminationChange(Mat src, Mat mask, Mat dst, float alpha, float beta);

void TextureFlattening(Mat src, Mat mask, Mat dst, float low_threshold, float high_threshold, int kernel_size);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PHOTO_H