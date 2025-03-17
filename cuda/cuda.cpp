#include "cuda.h"

GpuMat GpuMat_New() {
    try {
        return new cv::cuda::GpuMat();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

GpuMat GpuMat_NewFromMat(Mat mat) {
    try {
        return new cv::cuda::GpuMat(*mat);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

GpuMat GpuMat_NewWithSize(int rows, int cols, int type) {
    try {
        return new cv::cuda::GpuMat(rows, cols, type);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void GpuMat_Upload(GpuMat m, Mat data, Stream s){
    try {
        if (s == NULL) {
            m->upload(*data);
            return;
        }
        m->upload(*data, *s);    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GpuMat_Download(GpuMat m, Mat dst, Stream s){
    try {
        if (s == NULL) {
            m->download(*dst);
            return;
        }
        m->download(*dst, *s);    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int GpuMat_Empty(GpuMat m){
    try {
        return m->empty();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return true;
    }
}

void GpuMat_Close(GpuMat m){
    delete m;
}

void PrintCudaDeviceInfo(int device){
    try {
        cv::cuda::printCudaDeviceInfo(device);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void PrintShortCudaDeviceInfo(int device){
    try {
        cv::cuda::printShortCudaDeviceInfo(device);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int GetCudaEnabledDeviceCount(){
    try {
        return cv::cuda::getCudaEnabledDeviceCount();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

int GetCudaDevice() {
    try {
        return cv::cuda::getDevice();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

void SetCudaDevice(int device) {
    try {
        cv::cuda::setDevice(device);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void ResetCudaDevice(){
    try {
        cv::cuda::resetDevice();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

bool CudaDeviceSupports(int features) {
    try {
        return cv::cuda::deviceSupports(cv::cuda::FeatureSet(features));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

void GpuMat_ConvertTo(GpuMat m, GpuMat dst, int type, Stream s) {
    try {
        if (s == NULL) {
            m->convertTo(*dst, type);
            return;
        }
        m->convertTo(*dst, type, *s);    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GpuMat_ConvertFp16(GpuMat m, GpuMat dst) {
    try {
        cv::cuda::convertFp16(*m, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GpuMat_CopyTo(GpuMat m, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            m->copyTo(*dst);
            return;
        }
        m->copyTo(*dst, *s);    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

GpuMat GpuMat_Reshape(GpuMat m, int cn, int rows) {
    try {
        return new cv::cuda::GpuMat(m->reshape(cn, rows));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

int GpuMat_Cols(GpuMat m) {
    return m->cols;
}

int GpuMat_Rows(GpuMat m) {
    return m->rows;
}

int GpuMat_Channels(GpuMat m) {
    return m->channels();
}

int GpuMat_Type(GpuMat m) {
    return m->type();
}

Stream Stream_New() {
    try {
        return new cv::cuda::Stream();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void Stream_Close(Stream s){
    delete s;
}

bool Stream_QueryIfComplete(Stream s) {
    try {
        return s->queryIfComplete();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

void Stream_WaitForCompletion(Stream s) {
    try {
        s->waitForCompletion();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
