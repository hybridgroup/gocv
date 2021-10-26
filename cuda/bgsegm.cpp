#include "bgsegm.h"

CudaBackgroundSubtractorMOG2 CudaBackgroundSubtractorMOG2_Create() {
    return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG2>(cv::cuda::createBackgroundSubtractorMOG2());
}

void CudaBackgroundSubtractorMOG2_Close(CudaBackgroundSubtractorMOG2 b) {
    delete b;
}

void CudaBackgroundSubtractorMOG2_Apply(CudaBackgroundSubtractorMOG2 b, GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*b)->apply(*src, *dst);
        return;
    }
    (*b)->apply(*src, *dst, -1.0, *s);
}

CudaBackgroundSubtractorMOG CudaBackgroundSubtractorMOG_Create() {
    return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG>(cv::cuda::createBackgroundSubtractorMOG());
}

void CudaBackgroundSubtractorMOG_Close(CudaBackgroundSubtractorMOG b) {
    delete b;
}

void CudaBackgroundSubtractorMOG_Apply(CudaBackgroundSubtractorMOG b, GpuMat src, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*b)->apply(*src, *dst);
        return;
    }
    (*b)->apply(*src, *dst, -1.0, *s);
}