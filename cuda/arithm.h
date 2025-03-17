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

#ifdef __cplusplus
typedef cv::Ptr< cv::cuda::LookUpTable >* LookUpTable;
#else
typedef void* LookUpTable;
#endif

OpenCVResult GpuAbs(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuAbsDiff(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuAdd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuBitwiseAnd(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuBitwiseNot(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuBitwiseOr(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuBitwiseXor(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuDivide(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuExp(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuLog(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuMax(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuMin(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuMultiply(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuSqr(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuSqrt(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuSubtract(GpuMat src1, GpuMat src2, GpuMat dst, Stream s);
OpenCVResult GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ, Stream s);
OpenCVResult GpuFlip(GpuMat src, GpuMat dst, int flipCode, Stream s);
OpenCVResult GpuMerge(struct GpuMats mats, GpuMat dst, Stream s);
OpenCVResult GpuTranspose(GpuMat src, GpuMat dst, Stream s);
OpenCVResult GpuAddWeighted(GpuMat src1, double alpha, GpuMat src2, double beta, double gamma, GpuMat dst, int dType, Stream s);
OpenCVResult GpuCopyMakeBorder(GpuMat src, GpuMat dst, int top, int bottom, int left, int right, int borderType, Scalar value, Stream s);

//LookUpTable
LookUpTable Cuda_Create_LookUpTable(GpuMat lut);
void Cuda_LookUpTable_Close(LookUpTable lt);
OpenCVResult Cuda_LookUpTable_Transform(LookUpTable lt, GpuMat src, GpuMat dst, Stream s);

bool Cuda_LookUpTable_Empty(LookUpTable lut);

OpenCVResult Cuda_Split(GpuMat src, GpuMats dst, Stream s);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDA_ARITHM_H_
