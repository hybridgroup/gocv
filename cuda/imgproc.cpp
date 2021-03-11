#include "../core.h"
#include "imgproc.h"
#include <string.h>

void GpuCvtColor(GpuMat src, GpuMat dst, int code) {
    cv::cuda::cvtColor(*src, *dst, code);
}

CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh) {
    return new cv::Ptr<cv::cuda::CannyEdgeDetector>(cv::cuda::createCannyEdgeDetector(lowThresh, highThresh));
}

CannyEdgeDetector CreateCannyEdgeDetectorWithParams(double lowThresh, double highThresh, int appertureSize, bool L2gradient) {
    return new cv::Ptr<cv::cuda::CannyEdgeDetector>(cv::cuda::createCannyEdgeDetector(lowThresh, highThresh, appertureSize, L2gradient));
}

void CannyEdgeDetector_Close(CannyEdgeDetector det) {
    delete det;
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

HoughLinesDetector HoughLinesDetector_Create(double rho, double theta, int threshold) {
    return new cv::Ptr<cv::cuda::HoughLinesDetector>(cv::cuda::createHoughLinesDetector(rho, theta, threshold));
}

HoughLinesDetector HoughLinesDetector_CreateWithParams(double rho, double theta, int threshold, bool sort, int maxlines) {
    return new cv::Ptr<cv::cuda::HoughLinesDetector>(cv::cuda::createHoughLinesDetector(rho, theta, threshold, sort, maxlines));
}

void HoughLinesDetector_Close(HoughLinesDetector hld) {
    delete hld;
}

GpuMat HoughLinesDetector_Detect(HoughLinesDetector hld, GpuMat img) {
    GpuMat dst = new cv::cuda::GpuMat();
    (*hld)->detect(*img, *dst);

    return dst;
}

HoughSegmentDetector HoughSegmentDetector_Create(double rho, double theta, int minLineLength, int maxLineGap) {
    return new cv::Ptr<cv::cuda::HoughSegmentDetector>(cv::cuda::createHoughSegmentDetector(rho, theta, minLineLength, maxLineGap));
}

void HoughSegmentDetector_Close(HoughSegmentDetector hsd) {
    delete hsd;
}

GpuMat HoughSegmentDetector_Detect(HoughSegmentDetector hsd, GpuMat img) {
    GpuMat dst = new cv::cuda::GpuMat();
    (*hsd)->detect(*img, *dst);

    return dst;
}
