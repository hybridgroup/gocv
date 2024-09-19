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

void GpuDemosaicing(GpuMat src, GpuMat dst, int code, Stream s) {
    if (s == NULL) {
        cv::cuda::demosaicing(*src, *dst, code);
        return;
    }
    cv::cuda::demosaicing(*src, *dst, code, -1, *s);
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

TemplateMatching TemplateMatching_Create(int srcType, int method) {
    return new cv::Ptr<cv::cuda::TemplateMatching>(cv::cuda::createTemplateMatching(srcType, method));
}

void TemplateMatching_Close(TemplateMatching tm) {
    delete tm;
}

void TemplateMatching_Match(TemplateMatching tm, GpuMat img, GpuMat tmpl, GpuMat dst, Stream s) {
    if (s == NULL) {
        (*tm)->match(*img, *tmpl, *dst);
    } else {
        (*tm)->match(*img, *tmpl, *dst, *s);
    }
    return;
}

void AlphaComp(GpuMat img1, GpuMat img2, GpuMat dst, int alpha_op, Stream s) {
    if(s == NULL) {
        cv::cuda::alphaComp(*img1, *img2, *dst, alpha_op);
    } else {
        cv::cuda::alphaComp(*img1, *img2, *dst, alpha_op, *s);
    }
}

void GammaCorrection(GpuMat src, GpuMat dst, bool forward, Stream s) {
    if(s == NULL) {
        cv::cuda::gammaCorrection(*src, *dst, forward);
    } else {
        cv::cuda::gammaCorrection(*src, *dst, forward, *s);
    }
}

void SwapChannels(GpuMat image, int dstOrder[4], Stream s) {
    if(s == NULL) {
        cv::cuda::swapChannels(*image, dstOrder);
    } else {
        cv::cuda::swapChannels(*image, dstOrder, *s);
    }
}

void Cuda_CalcHist(GpuMat src, GpuMat dst, Stream s) {
    if(s == NULL) {
        cv::cuda::calcHist(*src, *dst);
    }else{
        cv::cuda::calcHist(*src, *dst, *s);
    }
}

void Cuda_CalcHist_WithParams(GpuMat src, GpuMat mask, GpuMat dst, Stream s) {
    if(s == NULL) {
        cv::cuda::calcHist(*src, *mask, *dst);
    }else{
        cv::cuda::calcHist(*src, *mask, *dst, *s);
    }
}

void Cuda_EqualizeHist(GpuMat src, GpuMat dst, Stream s) {
    if(s == NULL) {
        cv::cuda::equalizeHist(*src, *dst);
    }else{
        cv::cuda::equalizeHist(*src, *dst, *s);
    }
}

void Cuda_EvenLevels(GpuMat levels, int nLevels, int lowerLevel, int upperLevel, Stream s) {
    if(s == NULL) {
        cv::cuda::evenLevels(*levels, nLevels, lowerLevel, upperLevel);
    }else{
        cv::cuda::evenLevels(*levels, nLevels, lowerLevel, upperLevel, *s);
    }
}

void Cuda_HistEven(GpuMat src, GpuMat hist, int histSize, int lowerLevel, int upperLevel, Stream s) {
    if(s == NULL) {
        cv::cuda::histEven(*src, *hist, histSize, lowerLevel, upperLevel);
    }else{
        cv::cuda::histEven(*src, *hist, histSize, lowerLevel, upperLevel, *s);
    }
}


void Cuda_HistRange(GpuMat src, GpuMat hist, GpuMat levels, Stream s){
    if(s == NULL) {
        cv::cuda::histRange(*src, *hist, *levels);
    } else {
        cv::cuda::histRange(*src, *hist, *levels, *s);
    }
} 

void Cuda_BilateralFilter(GpuMat src, GpuMat dst, int kernel_size, float sigma_color, float sigma_spatial, int borderMode, Stream s) {
    if(s == NULL) {
        cv::cuda::bilateralFilter(*src, *dst, kernel_size, sigma_color, sigma_spatial, borderMode);
    } else {
        cv::cuda::bilateralFilter(*src, *dst, kernel_size, sigma_color, sigma_spatial, borderMode, *s);
    }
}

void Cuda_BlendLinear(GpuMat img1, GpuMat img2, GpuMat weights1, GpuMat weights2, GpuMat result, Stream s) {
    if(s == NULL) {
        cv::cuda::blendLinear(*img1, *img2, *weights1, *weights2, *result);
    } else {
        cv::cuda::blendLinear(*img1, *img2, *weights1, *weights2, *result, *s);
    }
}

void Cuda_MeanShiftFiltering(GpuMat src, GpuMat dst, int sp, int sr, TermCriteria criteria, Stream s) {
    if(s == NULL) {
        cv::cuda::meanShiftFiltering(*src, *dst, sp, sr, *criteria);
    } else {
        cv::cuda::meanShiftFiltering(*src, *dst, sp, sr, *criteria, *s);
    }
}

void Cuda_MeanShiftProc(GpuMat src, GpuMat dstr, GpuMat dstsp, int sp, int sr, TermCriteria criteria, Stream s) {
    if(s == NULL) {
        cv::cuda::meanShiftProc(*src, *dstr, *dstsp, sp, sr, *criteria);
    } else {
        cv::cuda::meanShiftProc(*src, *dstr, *dstsp, sp, sr, *criteria, *s);
    }
}


void Cuda_MeanShiftSegmentation(GpuMat src, GpuMat dst, int sp, int sr, int minSize, TermCriteria criteria, Stream s) {
    if(s == NULL) {
        cv::cuda::meanShiftSegmentation(*src, *dst, sp, sr, minSize, *criteria);
    } else {
        cv::cuda::meanShiftSegmentation(*src, *dst, sp, sr, minSize, *criteria, *s);
    }
}