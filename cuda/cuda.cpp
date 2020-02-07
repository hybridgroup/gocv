#include "cuda.h"

GpuMat GpuMat_New() {
    return new cv::cuda::GpuMat();
}

GpuMat GpuMat_NewFromMat(Mat mat) {
    return new cv::cuda::GpuMat(*mat);
}

void GpuMat_Upload(GpuMat m,Mat data){
    m->upload(*data);
}

void GpuMat_Download(GpuMat m,Mat dst){
    m->download(*dst);
}

int GpuMat_Empty(GpuMat m){
    return m->empty();
}

void GpuMat_Close(GpuMat m){
    delete m;
}

void PrintCudaDeviceInfo(int device){
    cv::cuda::printCudaDeviceInfo(device);
}

void PrintShortCudaDeviceInfo(int device){
    cv::cuda::printShortCudaDeviceInfo(device);
}

int GetCudaEnabledDeviceCount(){
    return cv::cuda::getCudaEnabledDeviceCount();
}

void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type) {
    m->convertTo(*dst, type);
}
