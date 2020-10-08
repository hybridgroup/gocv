#ifndef _OPENCV3_CUDA_IMGPROC_H_
#define _OPENCV3_CUDA_IMGPROC_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudaimgproc.hpp>
#include <opencv2/cudaarithm.hpp>
extern "C" {
#endif
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::CannyEdgeDetector>* CannyEdgeDetector;
#else
typedef void* CannyEdgeDetector;
#endif

void GpuCvtColor(GpuMat src, GpuMat dst, int code);
void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ);
CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh, int appertureSize, bool L2gradient);
GpuMat CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img);
int CannyEdgeDetector_GetAppertureSize(CannyEdgeDetector det);
double CannyEdgeDetector_GetHighThreshold(CannyEdgeDetector det);
bool CannyEdgeDetector_GetL2Gradient(CannyEdgeDetector det);
double CannyEdgeDetector_GetLowThreshold(CannyEdgeDetector det);
void CannyEdgeDetector_SetAppertureSize(CannyEdgeDetector det, int appertureSize);
void CannyEdgeDetector_SetHighThreshold(CannyEdgeDetector det, double highThresh);
void CannyEdgeDetector_SetL2Gradient(CannyEdgeDetector det, bool L2gradient);
void CannyEdgeDetector_SetLowThreshold(CannyEdgeDetector det, double lowThresh);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_IMGPROC_H_
