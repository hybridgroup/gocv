#include "bgsegm.h"

CudaBackgroundSubtractorMOG2 CudaBackgroundSubtractorMOG2_Create() {
    try {
        return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG2>(cv::cuda::createBackgroundSubtractorMOG2());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CudaBackgroundSubtractorMOG2_Close(CudaBackgroundSubtractorMOG2 b) {
    delete b;
}

OpenCVResult CudaBackgroundSubtractorMOG2_Apply(CudaBackgroundSubtractorMOG2 b, GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*b)->apply(*src, *dst);
        } else {
            (*b)->apply(*src, *dst, -1.0, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

CudaBackgroundSubtractorMOG CudaBackgroundSubtractorMOG_Create() {
    try {
        return new cv::Ptr<cv::cuda::BackgroundSubtractorMOG>(cv::cuda::createBackgroundSubtractorMOG());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CudaBackgroundSubtractorMOG_Close(CudaBackgroundSubtractorMOG b) {
    delete b;
}

OpenCVResult CudaBackgroundSubtractorMOG_Apply(CudaBackgroundSubtractorMOG b, GpuMat src, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*b)->apply(*src, *dst);
        } else {
            (*b)->apply(*src, *dst, -1.0, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}