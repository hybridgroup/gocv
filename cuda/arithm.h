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

void GpuAbs(GpuMat src, GpuMat dst, Stream s);
void GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s);
void GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuExp(GpuMat src, GpuMat dst, Stream s);
void GpuLog(GpuMat src, GpuMat dst, Stream s);
void GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuSqr(GpuMat src, GpuMat dst, Stream s);
void GpuSqrt(GpuMat src, GpuMat dst, Stream s);
void GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s);
void GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_ARITHM_H_
