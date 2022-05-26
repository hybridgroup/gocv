#ifndef _OPENCV3_XIMGPROC_H_
#define _OPENCV3_XIMGPROC_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/ximgproc.hpp>
extern "C" {
#endif

#include "../core.h"

void anisotropicDiffusion(Mat src, Mat dst, float alpha, float K, int niters);
void edgePreservingFilter(Mat src, Mat dst, int d, float threshold);
void niBlackThreshold(Mat src, Mat dst, float maxValue, int type, int blockSize, float k, int binarizationMethod, float r);
void PeiLinNormalization(Mat src, Mat dst);
void thinning(Mat src, Mat dst, int typ);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XIMGPROC_H
