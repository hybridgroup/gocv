#include "features2d.h"

SimpleBlobDetector SimpleBlobDetector_Create() {
    // TODO: params
    return cv::SimpleBlobDetector::create();
}

void SimpleBlobDetector_Close(SimpleBlobDetector b) {
    delete b;
}

void SimpleBlobDetector_Detect(SimpleBlobDetector b, Mat src) {
    // TODO: return the detected keypoints
    std::vector<cv::KeyPoint> keypoints;
    b->detect(*src, keypoints);
}
