#include "../core.h"
#include "cuda.h"
#include "objdetect.h"

// CascadeClassifier_GPU

CascadeClassifier_GPU CascadeClassifier_GPU_Create(const char*  cascade_name) {
    return new cv::Ptr<cv::cuda::CascadeClassifier>(cv::cuda::CascadeClassifier::create(cascade_name));
}

struct Rects CascadeClassifier_GPU_DetectMultiScale(CascadeClassifier_GPU cs, GpuMat img) {
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
}

// HOG

HOG HOG_Create() {
    return new cv::Ptr<cv::cuda::HOG>(cv::cuda::HOG::create());
}

HOG HOG_CreateWithParams(Size winSize, Size blockSize, Size blockStride, Size cellSize, int nbins) {
    cv::Size winSz(winSize.width, winSize.height);
    cv::Size blockSz(blockSize.width, blockSize.height);
    cv::Size blockSt(blockStride.width, blockStride.height);
    cv::Size cellSz(cellSize.width, cellSize.height);

    return new cv::Ptr<cv::cuda::HOG>(cv::cuda::HOG::create(winSz, blockSz, blockSt, cellSz, nbins));
}

struct Rects HOG_DetectMultiScale(HOG hog, GpuMat img) {    
    std::vector<cv::Rect> detected;    
    (*hog)->detectMultiScale(*img, detected);

    Rect* rects = new Rect[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
        Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)detected.size()};
    return ret;
}

GpuMat HOG_Compute(HOG hog, GpuMat img) {    
    GpuMat dst = new cv::cuda::GpuMat();
    (*hog)->compute(*img, *dst);

    return dst;
}

Mat HOG_GetPeopleDetector(HOG hog) {
    return new cv::Mat((*hog)->getDefaultPeopleDetector());
}

void HOG_SetSVMDetector(HOG hog, Mat det) {
    (*hog)->setSVMDetector(*det);
}

int HOG_GetDescriptorFormat(HOG hog) {
    return int((*hog)->getDescriptorFormat());
}

size_t HOG_GetBlockHistogramSize(HOG hog) {
    return size_t((*hog)->getBlockHistogramSize());
}

size_t HOG_GetDescriptorSize(HOG hog) {
    return size_t((*hog)->getDescriptorSize());
}

bool HOG_GetGammaCorrection(HOG hog) {
    return bool((*hog)->getGammaCorrection());
}

int HOG_GetGroupThreshold(HOG hog) {
    return int((*hog)->getGroupThreshold());
}

double HOG_GetHitThreshold(HOG hog) {
    return double((*hog)->getHitThreshold());
}

double HOG_GetL2HysThreshold(HOG hog) {
    return double((*hog)->getL2HysThreshold());
}

int HOG_GetNumLevels(HOG hog) {
    return int((*hog)->getNumLevels());
}

double HOG_GetScaleFactor(HOG hog) {
    return double((*hog)->getScaleFactor());
}

double HOG_GetWinSigma(HOG hog) {
    return double((*hog)->getWinSigma());
}

struct Size HOG_GetWinStride(HOG hog) {
    cv::Size sz = (*hog)->getWinStride();
    Size size = {sz.width, sz.height};
    return size;
}

void HOG_SetDescriptorFormat(HOG hog, int descrFormat) {
    auto df = static_cast<cv::HOGDescriptor::DescriptorStorageFormat>(descrFormat); 
     (*hog)->setDescriptorFormat(df);
}

void HOG_SetGammaCorrection(HOG hog, bool gammaCorrection) {
     (*hog)->setGammaCorrection(gammaCorrection);
}

void HOG_SetGroupThreshold(HOG hog, int groupThreshold) {
     (*hog)->setGroupThreshold(groupThreshold);
}

void HOG_SetHitThreshold(HOG hog, double hitThreshold) {
     (*hog)->setHitThreshold(hitThreshold);
}

void HOG_SetL2HysThreshold(HOG hog, double thresholdL2hys) {
     (*hog)->setL2HysThreshold(thresholdL2hys);
}

void HOG_SetNumLevels(HOG hog, int nlevels) {
     (*hog)->setNumLevels(nlevels);
}

void HOG_SetScaleFactor(HOG hog, double scale0) {
     (*hog)->setScaleFactor(scale0);
}

void HOG_SetWinSigma(HOG hog, double winSigma) {
     (*hog)->setWinSigma(winSigma);
}

void HOG_SetWinStride(HOG hog, Size dsize) {
    cv::Size sz(dsize.width, dsize.height);
    (*hog)->setWinStride(sz);
}
