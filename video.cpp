#include "video.h"

BackgroundSubtractor BackgroundSubtractor_CreateMOG2() {
    cv::ocl::setUseOpenCL(false);
    return cv::createBackgroundSubtractorMOG2();
}

void BackgroundSubtractor_Close(BackgroundSubtractor b) {
    delete b;
}

void BackgroundSubtractor_Apply(BackgroundSubtractor b, Mat src, Mat dst) {
    b->apply(*src, *dst);
}
