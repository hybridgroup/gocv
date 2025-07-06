//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_contrib_ximgproc)

#include "ximgproc.h"

OpenCVResult anisotropicDiffusion(Mat src, Mat dst, float alpha, float K, int niters) {
    try {
        cv::ximgproc::anisotropicDiffusion(*src, *dst, alpha, K, niters);	
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult edgePreservingFilter(Mat src, Mat dst, int d, float threshold) {
    try {
        cv::ximgproc::edgePreservingFilter(*src, *dst, d, threshold);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult niBlackThreshold(Mat src, Mat dst, float maxValue, int type, int blockSize, float k, int binarizationMethod, float r) {
    try {
        cv::ximgproc::niBlackThreshold(*src, *dst, maxValue, type, blockSize, k, binarizationMethod, r);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult PeiLinNormalization(Mat src, Mat dst) {
    try {
        cv::ximgproc::PeiLinNormalization(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult thinning(Mat src, Mat dst, int typ) {
    try {
        cv::ximgproc::thinning(*src, *dst, typ);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
