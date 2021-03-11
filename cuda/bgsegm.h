#ifndef _OPENCV3_CUDABGSEGM_H_
#define _OPENCV3_CUDABGSEGM_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudabgsegm.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::BackgroundSubtractorMOG2>* CudaBackgroundSubtractorMOG2;
typedef cv::Ptr<cv::cuda::BackgroundSubtractorMOG>* CudaBackgroundSubtractorMOG;
#else
typedef void* CudaBackgroundSubtractorMOG2;
typedef void* CudaBackgroundSubtractorMOG;
#endif

CudaBackgroundSubtractorMOG2 CudaBackgroundSubtractorMOG2_Create();
void CudaBackgroundSubtractorMOG2_Close(CudaBackgroundSubtractorMOG2 b);
void CudaBackgroundSubtractorMOG2_Apply(CudaBackgroundSubtractorMOG2 b, GpuMat src, GpuMat dst);

CudaBackgroundSubtractorMOG CudaBackgroundSubtractorMOG_Create();
void CudaBackgroundSubtractorMOG_Close(CudaBackgroundSubtractorMOG b);
void CudaBackgroundSubtractorMOG_Apply(CudaBackgroundSubtractorMOG b, GpuMat src, GpuMat dst);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDABGSEGM_H_
