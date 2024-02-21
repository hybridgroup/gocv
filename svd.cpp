#include "svd.h"

void SVD_Compute(Mat src, Mat w, Mat u, Mat vt, int flags) {
    cv::SVD::compute(*src, *w, *u, *vt, flags);
}