#ifndef _GOCV_CUDA_FILTERS_H_
#define _GOCV_CUDA_FILTERS_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudafilters.hpp>
extern "C" {
#endif
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::Filter>* GaussianFilter;
typedef cv::Ptr<cv::cuda::Filter>* SobelFilter;
#else
typedef void* GaussianFilter;
typedef void* SobelFilter;
#endif

// GaussianFilter
GaussianFilter CreateGaussianFilter(int srcType, int dstType, Size ksize, double sigma1);
GaussianFilter CreateGaussianFilterWithParams(int srcType, int dstType, Size ksize, double sigma1, double sigma2, int rowBorderMode, int columnBorderMode);
void GaussianFilter_Close(GaussianFilter gf);
void GaussianFilter_Apply(GaussianFilter gf, GpuMat img, GpuMat dst, Stream s);

// SobelFilter
SobelFilter CreateSobelFilter(int srcType, int dstType, int dx, int dy);
SobelFilter CreateSobelFilterWithParams(int srcType, int dstType, int dx, int dy, int ksize, double scale, int rowBorderMode, int columnBorderMode);
void SobelFilter_Close(SobelFilter sf);
void SobelFilter_Apply(SobelFilter sf, GpuMat img, GpuMat dst, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_GOCV_CUDA_FILTERS_H_
