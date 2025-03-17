#include "../core.h"
#include "filters.h"
#include <string.h>

GaussianFilter CreateGaussianFilter(int srcType, int dstType, Size ksize, double sigma1) {
    try {
        cv::Size sz(ksize.width, ksize.height);
        return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1));    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

GaussianFilter CreateGaussianFilterWithParams(int srcType, int dstType, Size ksize, double sigma1, double sigma2, int rowBorderMode, int columnBorderMode) {
    try {
        cv::Size sz(ksize.width, ksize.height);
        return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createGaussianFilter(srcType, dstType, sz, sigma1, sigma2, rowBorderMode, columnBorderMode));    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void GaussianFilter_Close(GaussianFilter gf) {
    delete gf;
}

OpenCVResult GaussianFilter_Apply(GaussianFilter gf, GpuMat img, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*gf)->apply(*img, *dst);
        } else {
            (*gf)->apply(*img, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

SobelFilter CreateSobelFilter(int srcType, int dstType, int dx, int dy) {
    try {
        return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

SobelFilter CreateSobelFilterWithParams(int srcType, int dstType, int dx, int dy, int ksize, double scale, int rowBorderMode, int columnBorderMode) {
    try {
        return new cv::Ptr<cv::cuda::Filter>(cv::cuda::createSobelFilter(srcType, dstType, dx, dy, ksize, rowBorderMode, columnBorderMode));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void SobelFilter_Close(SobelFilter sf) {
    delete sf;
}

OpenCVResult SobelFilter_Apply(SobelFilter sf, GpuMat img, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*sf)->apply(*img, *dst);
        } else {
            (*sf)->apply(*img, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
