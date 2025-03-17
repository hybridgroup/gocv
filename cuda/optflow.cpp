#include "optflow.h"

CudaSparsePyrLKOpticalFlow CudaSparsePyrLKOpticalFlow_Create() {
    try {
        return new cv::Ptr<cv::cuda::SparsePyrLKOpticalFlow>(cv::cuda::SparsePyrLKOpticalFlow::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CudaSparsePyrLKOpticalFlow_Calc(CudaSparsePyrLKOpticalFlow p, GpuMat prevImg, GpuMat nextImg, GpuMat prevPts, GpuMat nextPts, GpuMat status){
    try {
        (*p)->calc(*prevImg,*nextImg,*prevPts,*nextPts,*status);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}