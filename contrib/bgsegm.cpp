#include "bgsegm.h"

BackgroundSubtractorCNT BackgroundSubtractorCNT_Create() {
    return new cv::Ptr<cv::bgsegm::BackgroundSubtractorCNT>(cv::bgsegm::createBackgroundSubtractorCNT());
}

void BackgroundSubtractorCNT_Close(BackgroundSubtractorCNT b) {
    delete b;
}

void BackgroundSubtractorCNT_Apply(BackgroundSubtractorCNT b, Mat src, Mat dst) {
    (*b)->apply(*src, *dst);
}
