//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_cuda_objdetect)

#include "../core.h"
#include "cuda.h"
#include "objdetect.h"

// CascadeClassifier_GPU

CascadeClassifier_GPU CascadeClassifier_GPU_Create(const char*  cascade_name) {
    try {
        return new cv::Ptr<cv::cuda::CascadeClassifier>(cv::cuda::CascadeClassifier::create(cascade_name));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

struct Rects CascadeClassifier_GPU_DetectMultiScale(CascadeClassifier_GPU cs, GpuMat img) {
    try {
        std::vector<cv::Rect> detected;
        cv::cuda::GpuMat objbuf;
        
        (*cs)->detectMultiScale(*img, objbuf); // uses all default parameters
        (*cs)->convert(objbuf, detected);
        
        Rect* rects = new Rect[detected.size()];
    
        for (size_t i = 0; i < detected.size(); ++i) {
            Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
            rects[i] = r;
        }
    
        Rects ret = {rects, (int)detected.size()};
        return ret;    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rects ret = {NULL, 0};
        return ret;
    }
}

// HOG

HOG HOG_Create() {
    try {
        return new cv::Ptr<cv::cuda::HOG>(cv::cuda::HOG::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

HOG HOG_CreateWithParams(Size winSize, Size blockSize, Size blockStride, Size cellSize, int nbins) {
    try {
        cv::Size winSz(winSize.width, winSize.height);
        cv::Size blockSz(blockSize.width, blockSize.height);
        cv::Size blockSt(blockStride.width, blockStride.height);
        cv::Size cellSz(cellSize.width, cellSize.height);
    
        return new cv::Ptr<cv::cuda::HOG>(cv::cuda::HOG::create(winSz, blockSz, blockSt, cellSz, nbins));    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

struct Rects HOG_DetectMultiScale(HOG hog, GpuMat img) {    
    try {
        std::vector<cv::Rect> detected;    
        (*hog)->detectMultiScale(*img, detected);
    
        Rect* rects = new Rect[detected.size()];
        for (size_t i = 0; i < detected.size(); ++i) {
            Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
            rects[i] = r;
        }
    
        Rects ret = {rects, (int)detected.size()};
        return ret;    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rects ret = {NULL, 0};
        return ret;
    }
}

GpuMat HOG_Compute(HOG hog, GpuMat img) {    
    try {
        GpuMat dst = new cv::cuda::GpuMat();
        (*hog)->compute(*img, *dst);
    
        return dst;    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Mat HOG_GetPeopleDetector(HOG hog) {
    try {
        return new cv::Mat((*hog)->getDefaultPeopleDetector());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void HOG_SetSVMDetector(HOG hog, Mat det) {
    try {
        (*hog)->setSVMDetector(*det);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int HOG_GetDescriptorFormat(HOG hog) {
    try {
        return int((*hog)->getDescriptorFormat());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

size_t HOG_GetBlockHistogramSize(HOG hog) {
    try {
        return size_t((*hog)->getBlockHistogramSize());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

size_t HOG_GetDescriptorSize(HOG hog) {
    try {
        return size_t((*hog)->getDescriptorSize());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

bool HOG_GetGammaCorrection(HOG hog) {
    try {
        return bool((*hog)->getGammaCorrection());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

int HOG_GetGroupThreshold(HOG hog) {
    try {
        return int((*hog)->getGroupThreshold());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

double HOG_GetHitThreshold(HOG hog) {
    try {
        return double((*hog)->getHitThreshold());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

double HOG_GetL2HysThreshold(HOG hog) {
    try {
        return double((*hog)->getL2HysThreshold());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

int HOG_GetNumLevels(HOG hog) {
    try {
        return int((*hog)->getNumLevels());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

double HOG_GetScaleFactor(HOG hog) {
    try {
        return double((*hog)->getScaleFactor());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

double HOG_GetWinSigma(HOG hog) {
    try {
        return double((*hog)->getWinSigma());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

struct Size HOG_GetWinStride(HOG hog) {
    try {
        cv::Size sz = (*hog)->getWinStride();
        Size size = {sz.width, sz.height};
        return size;    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Size size = {0, 0};
        return size;
    }
}

void HOG_SetDescriptorFormat(HOG hog, int descrFormat) {
    try {
        auto df = static_cast<cv::HOGDescriptor::DescriptorStorageFormat>(descrFormat); 
        (*hog)->setDescriptorFormat(df);   
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetGammaCorrection(HOG hog, bool gammaCorrection) {
    try {
        (*hog)->setGammaCorrection(gammaCorrection);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetGroupThreshold(HOG hog, int groupThreshold) {
    try {
        (*hog)->setGroupThreshold(groupThreshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetHitThreshold(HOG hog, double hitThreshold) {
    try {
        (*hog)->setHitThreshold(hitThreshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetL2HysThreshold(HOG hog, double thresholdL2hys) {
    try {
        (*hog)->setL2HysThreshold(thresholdL2hys);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetNumLevels(HOG hog, int nlevels) {
    try {
        (*hog)->setNumLevels(nlevels);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetScaleFactor(HOG hog, double scale0) {
    try {
        (*hog)->setScaleFactor(scale0);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetWinSigma(HOG hog, double winSigma) {
    try {
        (*hog)->setWinSigma(winSigma);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HOG_SetWinStride(HOG hog, Size dsize) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        (*hog)->setWinStride(sz);    
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
