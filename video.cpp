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

void CalcOpticalFlowFarneback(Mat prevImg, Mat nextImg, Mat flow, double scale, int levels, int winsize,
	int iterations, int polyN, double polySigma, int flags) {
        cv::calcOpticalFlowFarneback(*prevImg, *nextImg, *flow, scale, levels, winsize, iterations, polyN, polySigma, flags);
}

void CalcOpticalFlowPyrLK(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err) {
    cv::calcOpticalFlowPyrLK(*prevImg, *nextImg, *prevPts, *nextPts, *status, *err);
}
