#ifndef _OPENCV_CUDAOPTFLOW_HPP_
#define _OPENCV_CUDAOPTFLOW_HPP_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/core/cuda.hpp>
#include <opencv2/cudaoptflow.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::cuda::SparsePyrLKOpticalFlow>* CudaSparsePyrLKOpticalFlow;
#else
typedef void* CudaSparsePyrLKOpticalFlow;
#endif

CudaSparsePyrLKOpticalFlow CudaSparsePyrLKOpticalFlow_Create();
void CudaSparsePyrLKOpticalFlow_Calc(CudaSparsePyrLKOpticalFlow p, GpuMat prevImg, GpuMat nextImg, GpuMat prevPts, GpuMat nextPts, GpuMat status);

#ifdef __cplusplus
}
#endif

#endif // _OPENCV_CUDAOPTFLOW_HPP_