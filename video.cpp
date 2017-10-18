#include "video.h"

BackgroundSubtractor BackgroundSubtractor_CreateMOG2() {
    return cv::createBackgroundSubtractorMOG2();
}

BackgroundSubtractor BackgroundSubtractor_CreateKNN() {
    return cv::createBackgroundSubtractorKNN();
}

void BackgroundSubtractor_Close(BackgroundSubtractor b) {
    delete b;
}

void BackgroundSubtractor_Apply(BackgroundSubtractor b, Mat src, Mat dst) {
    b->apply(*src, *dst);
}
