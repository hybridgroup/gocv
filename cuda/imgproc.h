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
typedef cv::Ptr<cv::cuda::HoughLinesDetector>* HoughLinesDetector;
typedef cv::Ptr<cv::cuda::HoughSegmentDetector>* HoughSegmentDetector;
#else
typedef void* CannyEdgeDetector;
typedef void* HoughLinesDetector;
typedef void* HoughSegmentDetector;
#endif

// standalone functions
void GpuCvtColor(GpuMat src, GpuMat dst, int code, Stream s);

// CannyEdgeDetector
CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh);
CannyEdgeDetector CreateCannyEdgeDetectorWithParams(double lowThresh, double highThresh, int appertureSize, bool L2gradient);
void CannyEdgeDetector_Close(CannyEdgeDetector det);
void CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img, GpuMat dst, Stream s);
int CannyEdgeDetector_GetAppertureSize(CannyEdgeDetector det);
double CannyEdgeDetector_GetHighThreshold(CannyEdgeDetector det);
bool CannyEdgeDetector_GetL2Gradient(CannyEdgeDetector det);
double CannyEdgeDetector_GetLowThreshold(CannyEdgeDetector det);
void CannyEdgeDetector_SetAppertureSize(CannyEdgeDetector det, int appertureSize);
void CannyEdgeDetector_SetHighThreshold(CannyEdgeDetector det, double highThresh);
void CannyEdgeDetector_SetL2Gradient(CannyEdgeDetector det, bool L2gradient);
void CannyEdgeDetector_SetLowThreshold(CannyEdgeDetector det, double lowThresh);

// HoughLinesDetector
HoughLinesDetector HoughLinesDetector_Create(double rho, double theta, int threshold);
HoughLinesDetector HoughLinesDetector_CreateWithParams(double rho, double theta, int threshold, bool sort, int maxlines);
void HoughLinesDetector_Close(HoughLinesDetector hld);
void HoughLinesDetector_Detect(HoughLinesDetector hld, GpuMat img, GpuMat dst, Stream s);

// HoughSegmentDetector
HoughSegmentDetector HoughSegmentDetector_Create(double rho, double theta, int minLineLength, int maxLineGap);
void HoughSegmentDetector_Close(HoughSegmentDetector hsd);
void HoughSegmentDetector_Detect(HoughSegmentDetector hsd, GpuMat img, GpuMat dst, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_IMGPROC_H_
