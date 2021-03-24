#include "../core.h"
#include "filters.h"
#include <string.h>

GaussianFilter CreateGaussianFilter(int srcType, int dstType, Size ksize, double sigma1) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1));
}

GaussianFilter CreateGaussianFilterWithParams(int srcType, int dstType, Size ksize, double sigma1, double sigma2, int rowBorderMode, int columnBorderMode) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1, sigma2, rowBorderMode, columnBorderMode));
}

void GaussianFilter_Close(GaussianFilter gf) {
    delete gf;
}

GpuMat GaussianFilter_Apply(GaussianFilter gf, GpuMat img) {
    GpuMat dst = new cv::cuda::GpuMat();
    (*gf)->apply(*img, *dst);

    return dst;
}

SobelFilter CreateSobelFilter(int srcType, int dstType, int dx, int dy) {
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy));
}

SobelFilter CreateSobelFilterWithParams(int srcType, int dstType, int dx, int dy, int ksize, double scale, int rowBorderMode, int columnBorderMode) {
    return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy, ksize, rowBorderMode, columnBorderMode));
}

void SobelFilter_Close(SobelFilter sf) {
    delete sf;
}

GpuMat SobelFilter_Apply(SobelFilter sf, GpuMat img) {
    GpuMat dst = new cv::cuda::GpuMat();
    (*sf)->apply(*img, *dst);

    return dst;
}
