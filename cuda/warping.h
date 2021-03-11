#ifndef _OPENCV3_CUDAWARPING_H_
#define _OPENCV3_CUDAWARPING_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudawarping.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"

void CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp);
void CudaPyrDown(GpuMat src, GpuMat dst);
void CudaPyrUp(GpuMat src, GpuMat dst);
void CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap);
void CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap);
void CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue);
void CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp);
void CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue);
void CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDAWARPING_H_
