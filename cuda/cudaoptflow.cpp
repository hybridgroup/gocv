#include "cudaoptflow.h"

CudaSparsePyrLKOpticalFlow CudaSparsePyrLKOpticalFlow_Create() {
    return new cv::Ptr<cv::cuda::SparsePyrLKOpticalFlow>(cv::cuda::SparsePyrLKOpticalFlow::create());
}

void CudaSparsePyrLKOpticalFlow_Calc(CudaSparsePyrLKOpticalFlow p, GpuMat prevImg, GpuMat nextImg, GpuMat prevPts, GpuMat nextPts, GpuMat status){
    (*p)->calc(*prevImg,*nextImg,*prevPts,*nextPts,*status);
}