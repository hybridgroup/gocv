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

BriefDescriptorExtractor BriefDescriptorExtractor_Create() {
    return new cv::Ptr<cv::xfeatures2d::BriefDescriptorExtractor>(cv::xfeatures2d::BriefDescriptorExtractor::create());
}

BriefDescriptorExtractor BriefDescriptorExtractor_CreateWithParams(int bytes, bool useOrientation) {
    return new cv::Ptr<cv::xfeatures2d::BriefDescriptorExtractor>(cv::xfeatures2d::BriefDescriptorExtractor::create(bytes, useOrientation));
}

void BriefDescriptorExtractor_Close(BriefDescriptorExtractor b) {
    delete b;
}

void BriefDescriptorExtractor_Compute(BriefDescriptorExtractor b, Mat src, struct KeyPoints kp, Mat desc) {
    std::vector<cv::KeyPoint> keypts;
    keypts.reserve(kp.length);
    cv::KeyPoint keypt;

    for (int i = 0; i < kp.length; ++i) {
        keypt = cv::KeyPoint(kp.keypoints[i].x, kp.keypoints[i].y,
                        kp.keypoints[i].size, kp.keypoints[i].angle, kp.keypoints[i].response,
                        kp.keypoints[i].octave, kp.keypoints[i].classID);
        keypts.push_back(keypt);
    }

    (*b)->compute(*src, keypts, *desc);
}
