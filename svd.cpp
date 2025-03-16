#include "svd.h"

void SVD_Compute(Mat src, Mat w, Mat u, Mat vt) {
    try {
        cv::SVD::compute(*src, *w, *u, *vt, 0);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}