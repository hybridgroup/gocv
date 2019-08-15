#include "cudawarping.h"

void CudaResize(GpuMat src, GpuMat dst, Size dsize, double fx, double fy, int interp) {
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::resize(*src, *dst, sz, fx, fy, interp);
}

void CudaPyrDown(GpuMat src, GpuMat dst) {
    cv::cuda::pyrDown(*src, *dst);
}

void CudaPyrUp(GpuMat src, GpuMat dst) {
    cv::cuda::pyrUp(*src, *dst);
}

void CudaBuildWarpAffineMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap) {
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::buildWarpAffineMaps(*M, inverse, sz, *xmap, *ymap);
}

void CudaBuildWarpPerspectiveMaps(GpuMat M, bool inverse, Size dsize, GpuMat xmap, GpuMat ymap) {
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::buildWarpPerspectiveMaps(*M, inverse, sz, *xmap, *ymap);
}

void CudaRemap(GpuMat src, GpuMat dst, GpuMat xmap, GpuMat ymap, int interp, int borderMode, Scalar borderValue) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::cuda::remap(*src, *dst, *xmap, *ymap, interp, borderMode, c);
}

void CudaRotate(GpuMat src, GpuMat dst, Size dsize, double angle, double xShift, double yShift, int interp) {  
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::rotate(*src, *dst, sz, angle, xShift, yShift, interp);
}

void CudaWarpAffine(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::warpAffine(*src, *dst, *M, sz, flags, borderMode, c);
}

void CudaWarpPerspective(GpuMat src, GpuMat dst, GpuMat M, Size dsize, int flags, int borderMode, Scalar borderValue) {
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::Size sz(dsize.width, dsize.height);
    cv::cuda::warpPerspective(*src, *dst, *M, sz, flags, borderMode, c);
}