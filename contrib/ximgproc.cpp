#include "ximgproc.h"

void edgePreservingFilter(Mat src, Mat dst, int d, float threshold) {
    cv::ximgproc::edgePreservingFilter(*src, *dst, d, threshold);
}

void niBlackThreshold(Mat src, Mat dst, float maxValue, int type, int blockSize, float k, int binarizationMethod, float r) {
    cv::ximgproc::niBlackThreshold(*src, *dst, maxValue, type, blockSize, k, binarizationMethod, r);
}

void thinning(Mat src, Mat dst, int typ) {
    cv::ximgproc::thinning(*src, *dst, typ);
}
