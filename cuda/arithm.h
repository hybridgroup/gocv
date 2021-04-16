#ifndef _OPENCV3_CUDA_ARITHM_H_
#define _OPENCV3_CUDA_ARITHM_H_

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudaarithm.hpp>
extern "C" {
#endif
#include "cuda.h"

void GpuAbs(GpuMat src, GpuMat dst);
void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ);
void GpuFlip(GpuMat src, GpuMat dst, int flipCode);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_ARITHM_H_
