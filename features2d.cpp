#include "features2d.h"

SimpleBlobDetector SimpleBlobDetector_Create() {
    // TODO: params
    return new cv::Ptr<cv::SimpleBlobDetector>(cv::SimpleBlobDetector::create());
}

void SimpleBlobDetector_Close(SimpleBlobDetector b) {
    delete b;
}

struct KeyPoints SimpleBlobDetector_Detect(SimpleBlobDetector b, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*b)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}
