#include "../core.h"
#include "imgproc.h"
#include <string.h>

void GpuCvtColor(GpuMat src, GpuMat dst, int code, Stream s) {
    if (s == NULL) {
        cv::cuda::cvtColor(*src, *dst, code);
        return;
    }
    cv::cuda::cvtColor(*src, *dst, code, 0, *s);
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

void CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*det)->detect(*img, *dst);
    } else {
        (*det)->detect(*img, *dst, *s);
    }
    return;
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

void HoughLinesDetector_Detect(HoughLinesDetector hld, GpuMat img, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*hld)->detect(*img, *dst);
    } else {
        (*hld)->detect(*img, *dst, *s);
    }
    return;
}

HoughSegmentDetector HoughSegmentDetector_Create(double rho, double theta, int minLineLength, int maxLineGap) {
    return new cv::Ptr<cv::cuda::HoughSegmentDetector>(cv::cuda::createHoughSegmentDetector(rho, theta, minLineLength, maxLineGap));
}

void HoughSegmentDetector_Close(HoughSegmentDetector hsd) {
    delete hsd;
}

void HoughSegmentDetector_Detect(HoughSegmentDetector hsd, GpuMat img, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*hsd)->detect(*img, *dst);
    } else {
        (*hsd)->detect(*img, *dst, *s);
    }
    return;
}
