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
typedef cv::Ptr<cv::cuda::TemplateMatching>* TemplateMatching;
#else
typedef void* CannyEdgeDetector;
typedef void* HoughLinesDetector;
typedef void* HoughSegmentDetector;
typedef void* TemplateMatching;
#endif

// standalone functions
void GpuCvtColor(GpuMat src, GpuMat dst, int code, Stream s);
void GpuDemosaicing(GpuMat src, GpuMat dst, int code, Stream s);
void AlphaComp(GpuMat img1, GpuMat img2, GpuMat dst, int alpha_op, Stream s);
void GammaCorrection(GpuMat src, GpuMat dst, bool forward, Stream s);
void SwapChannels(GpuMat image, int dstOrder[4], Stream s);
void Cuda_CalcHist(GpuMat src, GpuMat dst, Stream s);
void Cuda_CalcHist_WithParams(GpuMat src, GpuMat mask, GpuMat dst, Stream s);
void Cuda_EqualizeHist(GpuMat src, GpuMat dst, Stream s);
void Cuda_EvenLevels(GpuMat levels, int nLevels, int lowerLevel, int upperLevel, Stream s);
void Cuda_HistEven(GpuMat src, GpuMat hist, int histSize, int lowerLevel, int upperLevel, Stream s);
void Cuda_HistRange(GpuMat src, GpuMat hist, GpuMat levels, Stream s);
void Cuda_BilateralFilter(GpuMat src, GpuMat dst, int kernel_size, float sigma_color, float sigma_spatial, int borderMode, Stream s);
void Cuda_BlendLinear(GpuMat img1, GpuMat img2, GpuMat weights1, GpuMat weights2, GpuMat result, Stream s);
void Cuda_MeanShiftFiltering(GpuMat src, GpuMat dst, int sp, int sr, TermCriteria criteria, Stream s);
void Cuda_MeanShiftProc(GpuMat src, GpuMat dstr, GpuMat dstsp, int sp, int sr, TermCriteria criteria, Stream s);
void Cuda_MeanShiftSegmentation(GpuMat src, GpuMat dst, int sp, int sr, int minSize, TermCriteria criteria, Stream s);

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

// TemplateMatching
TemplateMatching TemplateMatching_Create(int srcType, int method);
void TemplateMatching_Close(TemplateMatching tm);
void TemplateMatching_Match(TemplateMatching tm, GpuMat img, GpuMat tmpl, GpuMat dst, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_IMGPROC_H_
