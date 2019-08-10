#include "cudawarping.h"

void CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp) {
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::resize(*src, *dst, sz, fx, fy, interp);
}