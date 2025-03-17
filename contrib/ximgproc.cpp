#include "ximgproc.h"

void anisotropicDiffusion(Mat src, Mat dst, float alpha, float K, int niters) {
    try {
        cv::ximgproc::anisotropicDiffusion(*src, *dst, alpha, K, niters);	
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void edgePreservingFilter(Mat src, Mat dst, int d, float threshold) {
    try {
        cv::ximgproc::edgePreservingFilter(*src, *dst, d, threshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void niBlackThreshold(Mat src, Mat dst, float maxValue, int type, int blockSize, float k, int binarizationMethod, float r) {
    try {
        cv::ximgproc::niBlackThreshold(*src, *dst, maxValue, type, blockSize, k, binarizationMethod, r);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void PeiLinNormalization(Mat src, Mat dst) {
    try {
        cv::ximgproc::PeiLinNormalization(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void thinning(Mat src, Mat dst, int typ) {
    try {
        cv::ximgproc::thinning(*src, *dst, typ);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
