#include "video.h"

BackgroundSubtractorMOG2 BackgroundSubtractorMOG2_Create() {
    return cv::createBackgroundSubtractorMOG2();
}

BackgroundSubtractorKNN BackgroundSubtractorKNN_Create() {
    return cv::createBackgroundSubtractorKNN();
}

void BackgroundSubtractorMOG2_Close(BackgroundSubtractorMOG2 b) {
    delete b;
}

void BackgroundSubtractorMOG2_Apply(BackgroundSubtractorMOG2 b, Mat src, Mat dst) {
    b->apply(*src, *dst);
}

void BackgroundSubtractorKNN_Close(BackgroundSubtractorKNN k) {
    delete k;
}

void BackgroundSubtractorKNN_Apply(BackgroundSubtractorKNN k, Mat src, Mat dst) {
    k->apply(*src, *dst);
}
