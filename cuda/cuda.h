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
typedef cv::cuda::Stream* Stream;
#else
typedef void* GpuMat;
typedef void* Stream;
#endif

GpuMat GpuMat_New();
GpuMat GpuMat_NewFromMat(Mat mat);
GpuMat GpuMat_NewWithSize(int rows, int cols, int type);
void GpuMat_Upload(GpuMat m, Mat data, Stream s);
void GpuMat_Download(GpuMat m, Mat dst, Stream s);
void GpuMat_Close(GpuMat m);
int GpuMat_Empty(GpuMat m);
void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type, Stream s);
void GpuMat_CopyTo(GpuMat m, GpuMat dst, Stream s);
GpuMat GpuMat_Reshape(GpuMat m, int cn, int rows);
int GpuMat_Cols(GpuMat m);
int GpuMat_Rows(GpuMat m);
int GpuMat_Channels(GpuMat m);
int GpuMat_Type(GpuMat m);

void PrintCudaDeviceInfo(int device);
void PrintShortCudaDeviceInfo(int device);
int GetCudaEnabledDeviceCount();
int GetCudaDevice();
void SetCudaDevice(int device);
void ResetCudaDevice();

Stream Stream_New();
void Stream_Close(Stream s);
bool Stream_QueryIfComplete(Stream s);
void Stream_WaitForCompletion(Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_H_
