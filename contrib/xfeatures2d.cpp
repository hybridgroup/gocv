#include "xfeatures2d.h"


SURF SURF_Create() {
    return new cv::Ptr<cv::xfeatures2d::SURF>(cv::xfeatures2d::SURF::create());
}

SURF SURF_CreateWithParams(double hessianThreshold, int nOctaves, int nOctaveLayers, bool extended, bool upright) {
    return new cv::Ptr<cv::xfeatures2d::SURF>(cv::xfeatures2d::SURF::create(hessianThreshold, nOctaves, nOctaveLayers, extended, upright));
}

void SURF_Close(SURF d) {
    delete d;
}

struct KeyPoints SURF_Detect(SURF d, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*d)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
        kps[i] = k;
    }

    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

struct KeyPoints SURF_DetectAndCompute(SURF d, Mat src, Mat mask, Mat desc) {
    std::vector<cv::KeyPoint> detected;
    (*d)->detectAndCompute(*src, *mask, detected, *desc);

    KeyPoint* kps = new KeyPoint[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
        kps[i] = k;
    }

    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}
