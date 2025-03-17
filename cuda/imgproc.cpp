#include "../core.h"
#include "imgproc.h"
#include <string.h>

OpenCVResult GpuCvtColor(GpuMat src, GpuMat dst, int code, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::cvtColor(*src, *dst, code);
        } else {
            cv::cuda::cvtColor(*src, *dst, code, 0, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GpuDemosaicing(GpuMat src, GpuMat dst, int code, Stream s) {
    try {
        if (s == NULL) {
            cv::cuda::demosaicing(*src, *dst, code);
        } else {
            cv::cuda::demosaicing(*src, *dst, code, -1, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

CannyEdgeDetector CreateCannyEdgeDetector(double lowThresh, double highThresh) {
    try {
        return new cv::Ptr<cv::cuda::CannyEdgeDetector>(cv::cuda::createCannyEdgeDetector(lowThresh, highThresh));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

CannyEdgeDetector CreateCannyEdgeDetectorWithParams(double lowThresh, double highThresh, int appertureSize, bool L2gradient) {
    try {
        return new cv::Ptr<cv::cuda::CannyEdgeDetector>(cv::cuda::createCannyEdgeDetector(lowThresh, highThresh, appertureSize, L2gradient));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CannyEdgeDetector_Close(CannyEdgeDetector det) {
    delete det;
}

OpenCVResult CannyEdgeDetector_Detect(CannyEdgeDetector det, GpuMat img, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*det)->detect(*img, *dst);
        } else {
            (*det)->detect(*img, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

int CannyEdgeDetector_GetAppertureSize(CannyEdgeDetector det) {
    try {
        return int((*det)->getAppertureSize());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

double CannyEdgeDetector_GetHighThreshold(CannyEdgeDetector det) {
    try {
        return double((*det)->getHighThreshold());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

bool CannyEdgeDetector_GetL2Gradient(CannyEdgeDetector det) {
    try {
        return bool((*det)->getL2Gradient());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

double CannyEdgeDetector_GetLowThreshold(CannyEdgeDetector det) {
    try {
        return double((*det)->getLowThreshold());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void CannyEdgeDetector_SetAppertureSize(CannyEdgeDetector det, int appertureSize) {
    try {
        (*det)->setAppertureSize(appertureSize);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CannyEdgeDetector_SetHighThreshold(CannyEdgeDetector det, double highThresh) {
    try {
        (*det)->setHighThreshold(highThresh);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CannyEdgeDetector_SetL2Gradient(CannyEdgeDetector det, bool L2gradient) {
    try {
        (*det)->setL2Gradient(L2gradient);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CannyEdgeDetector_SetLowThreshold(CannyEdgeDetector det, double lowThresh) {
    try {
        (*det)->setLowThreshold(lowThresh);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

HoughLinesDetector HoughLinesDetector_Create(double rho, double theta, int threshold) {
    try {
        return new cv::Ptr<cv::cuda::HoughLinesDetector>(cv::cuda::createHoughLinesDetector(rho, theta, threshold));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

HoughLinesDetector HoughLinesDetector_CreateWithParams(double rho, double theta, int threshold, bool sort, int maxlines) {
    try {
        return new cv::Ptr<cv::cuda::HoughLinesDetector>(cv::cuda::createHoughLinesDetector(rho, theta, threshold, sort, maxlines));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void HoughLinesDetector_Close(HoughLinesDetector hld) {
    delete hld;
}

OpenCVResult HoughLinesDetector_Detect(HoughLinesDetector hld, GpuMat img, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*hld)->detect(*img, *dst);
        } else {
            (*hld)->detect(*img, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

HoughSegmentDetector HoughSegmentDetector_Create(double rho, double theta, int minLineLength, int maxLineGap) {
    try {
        return new cv::Ptr<cv::cuda::HoughSegmentDetector>(cv::cuda::createHoughSegmentDetector(rho, theta, minLineLength, maxLineGap));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void HoughSegmentDetector_Close(HoughSegmentDetector hsd) {
    delete hsd;
}

OpenCVResult HoughSegmentDetector_Detect(HoughSegmentDetector hsd, GpuMat img, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*hsd)->detect(*img, *dst);
        } else {
            (*hsd)->detect(*img, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

TemplateMatching TemplateMatching_Create(int srcType, int method) {
    try {
        return new cv::Ptr<cv::cuda::TemplateMatching>(cv::cuda::createTemplateMatching(srcType, method));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void TemplateMatching_Close(TemplateMatching tm) {
    delete tm;
}

OpenCVResult TemplateMatching_Match(TemplateMatching tm, GpuMat img, GpuMat tmpl, GpuMat dst, Stream s) {
    try {
        if (s == NULL) {
            (*tm)->match(*img, *tmpl, *dst);
        } else {
            (*tm)->match(*img, *tmpl, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult AlphaComp(GpuMat img1, GpuMat img2, GpuMat dst, int alpha_op, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::alphaComp(*img1, *img2, *dst, alpha_op);
        } else {
            cv::cuda::alphaComp(*img1, *img2, *dst, alpha_op, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GammaCorrection(GpuMat src, GpuMat dst, bool forward, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::gammaCorrection(*src, *dst, forward);
        } else {
            cv::cuda::gammaCorrection(*src, *dst, forward, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }

}

OpenCVResult SwapChannels(GpuMat image, int dstOrder[4], Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::swapChannels(*image, dstOrder);
        } else {
            cv::cuda::swapChannels(*image, dstOrder, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_CalcHist(GpuMat src, GpuMat dst, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::calcHist(*src, *dst);
        } else {
            cv::cuda::calcHist(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_CalcHist_WithParams(GpuMat src, GpuMat mask, GpuMat dst, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::calcHist(*src, *mask, *dst);
        } else {
            cv::cuda::calcHist(*src, *mask, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_EqualizeHist(GpuMat src, GpuMat dst, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::equalizeHist(*src, *dst);
        } else {
            cv::cuda::equalizeHist(*src, *dst, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_EvenLevels(GpuMat levels, int nLevels, int lowerLevel, int upperLevel, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::evenLevels(*levels, nLevels, lowerLevel, upperLevel);
        } else {
            cv::cuda::evenLevels(*levels, nLevels, lowerLevel, upperLevel, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_HistEven(GpuMat src, GpuMat hist, int histSize, int lowerLevel, int upperLevel, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::histEven(*src, *hist, histSize, lowerLevel, upperLevel);
        } else {
            cv::cuda::histEven(*src, *hist, histSize, lowerLevel, upperLevel, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_HistRange(GpuMat src, GpuMat hist, GpuMat levels, Stream s){
    try {
        if(s == NULL) {
            cv::cuda::histRange(*src, *hist, *levels);
        } else {
            cv::cuda::histRange(*src, *hist, *levels, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
} 

OpenCVResult Cuda_BilateralFilter(GpuMat src, GpuMat dst, int kernel_size, float sigma_color, float sigma_spatial, int borderMode, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::bilateralFilter(*src, *dst, kernel_size, sigma_color, sigma_spatial, borderMode);
        } else {
            cv::cuda::bilateralFilter(*src, *dst, kernel_size, sigma_color, sigma_spatial, borderMode, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_BlendLinear(GpuMat img1, GpuMat img2, GpuMat weights1, GpuMat weights2, GpuMat result, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::blendLinear(*img1, *img2, *weights1, *weights2, *result);
        } else {
            cv::cuda::blendLinear(*img1, *img2, *weights1, *weights2, *result, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_MeanShiftFiltering(GpuMat src, GpuMat dst, int sp, int sr, TermCriteria criteria, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::meanShiftFiltering(*src, *dst, sp, sr, *criteria);
        } else {
            cv::cuda::meanShiftFiltering(*src, *dst, sp, sr, *criteria, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_MeanShiftProc(GpuMat src, GpuMat dstr, GpuMat dstsp, int sp, int sr, TermCriteria criteria, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::meanShiftProc(*src, *dstr, *dstsp, sp, sr, *criteria);
        } else {
            cv::cuda::meanShiftProc(*src, *dstr, *dstsp, sp, sr, *criteria, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Cuda_MeanShiftSegmentation(GpuMat src, GpuMat dst, int sp, int sr, int minSize, TermCriteria criteria, Stream s) {
    try {
        if(s == NULL) {
            cv::cuda::meanShiftSegmentation(*src, *dst, sp, sr, minSize, *criteria);
        } else {
            cv::cuda::meanShiftSegmentation(*src, *dst, sp, sr, minSize, *criteria, *s);
        }
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
