#include "../core.h"
#include "imgproc.h"
#include <string.h>

void GpuCvtColor(GpuMat src, GpuMat dst, int code) {
    cv::cuda::cvtColor(*src, *dst, code);
}

void GpuThreshold(GpuMat src, GpuMat dst, double thresh, double maxval, int typ) {
    cv::cuda::threshold(*src, *dst, thresh, maxval, typ);
}

CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh, int appertureSize, bool L2gradient) {
    return new cv::Ptr<cv::cuda::CannyEdgeDetector>(cv::cuda::createCannyEdgeDetector(lowThresh,highThresh,appertureSize,L2gradient));
}

GpuMat CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img) {    
    GpuMat dst = new cv::cuda::GpuMat();
    (*det)->detect(*img, *dst);

    return dst;
}

int CannyEdgeDetector_GetAppertureSize(CannyEdgeDetector det) {
    return int((*det)->getAppertureSize());
}

double CannyEdgeDetector_GetHighThreshold(CannyEdgeDetector det) {
    return double((*det)->getHighThreshold());
}

bool CannyEdgeDetector_GetL2Gradient(CannyEdgeDetector det) {
    return bool((*det)->getL2Gradient());
}

double CannyEdgeDetector_GetLowThreshold(CannyEdgeDetector det) {
    return double((*det)->getLowThreshold());
}

void CannyEdgeDetector_SetAppertureSize(CannyEdgeDetector det, int appertureSize) {
     (*det)->setAppertureSize(appertureSize);
}

void CannyEdgeDetector_SetHighThreshold(CannyEdgeDetector det, double highThresh) {
     (*det)->setHighThreshold(highThresh);
}

void CannyEdgeDetector_SetL2Gradient(CannyEdgeDetector det, bool L2gradient) {
     (*det)->setL2Gradient(L2gradient);
}

void CannyEdgeDetector_SetLowThreshold(CannyEdgeDetector det, double lowThresh) {
     (*det)->setLowThreshold(lowThresh);
}
