#ifndef _OPENCV3_PHOTO_H_
#define _OPENCV3_PHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

void FastNlMeansDenoising(Mat src, Mat dst);
void FastNlMeansDenoisingWithParams(Mat src, Mat dst, float h, int templateWindowSize, int searchWindowSize);
void FastNlMeansDenoisingColored(Mat src, Mat dst);
void FastNlMeansDenoisingColoredWithParams(Mat src, Mat dst, float h, float hColor, int templateWindowSize, int searchWindowSize);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PHOTO_H_