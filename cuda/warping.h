#ifndef _OPENCV3_CUDAWARPING_H_
#define _OPENCV3_CUDAWARPING_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/cudawarping.hpp>

extern "C" {
#endif

#include "../core.h"
#include "cuda.h"

OpenCVResult CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp, Stream s);
OpenCVResult CudaPyrDown(GpuMat src, GpuMat dst, Stream s);
OpenCVResult CudaPyrUp(GpuMat src, GpuMat dst, Stream s);
OpenCVResult CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s);
OpenCVResult CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap, Stream s);
OpenCVResult CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue, Stream s);
OpenCVResult CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp, Stream s);
OpenCVResult CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s);
OpenCVResult CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue, Stream s);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CUDAWARPING_H_
