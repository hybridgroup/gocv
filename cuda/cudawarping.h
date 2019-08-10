#ifndef _OPENCV3_CUDAWARPING_H_
#define _OPENCV3_CUDAWARPING_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudawarping.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"

void  CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDAWARPING_H_
