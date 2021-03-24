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
#else
typedef void* GaussianFilter;
#endif

// GaussianFilter
GaussianFilter CreateGaussianFilter(int srcType, int dstType, Size ksize, double sigma1);
GaussianFilter CreateGaussianFilterWithParams(int srcType, int dstType, Size ksize, double sigma1, double sigma2, int rowBorderMode, int columnBorderMode);

void GaussianFilter_Close(GaussianFilter gf);
GpuMat GaussianFilter_Apply(GaussianFilter gf, GpuMat img);

#ifdef __cplusplus
}
#endif

#endif //_GOCV_CUDA_FILTERS_H_
