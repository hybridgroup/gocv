#ifndef _OPENCV3_CUDA_H_
#define _OPENCV3_CUDA_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/core/cuda.hpp>

extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::cuda::GpuMat* GpuMat;
#else
typedef void* GpuMat;
#endif

GpuMat GpuMat_New();
GpuMat GpuMat_NewFromMat(Mat mat);
void GpuMat_Upload(GpuMat m,Mat data);
void GpuMat_Download(GpuMat m,Mat dst);
void GpuMat_Close(GpuMat m);
int GpuMat_Empty(GpuMat m);
void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type);

void PrintCudaDeviceInfo(int device);
void PrintShortCudaDeviceInfo(int device);
int GetCudaEnabledDeviceCount();

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_H_
