#include "features2d.h"

AKAZE AKAZE_Create() {
    // TODO: params
    return new cv::Ptr<cv::AKAZE>(cv::AKAZE::create());
}

void AKAZE_Close(AKAZE a) {
    delete a;
}

struct KeyPoints AKAZE_Detect(AKAZE a, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

struct KeyPoints AKAZE_DetectAndCompute(AKAZE a, Mat src, Mat mask, Mat desc) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detectAndCompute(*src, *mask, detected, *desc);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

AgastFeatureDetector AgastFeatureDetector_Create() {
    // TODO: params
    return new cv::Ptr<cv::AgastFeatureDetector>(cv::AgastFeatureDetector::create());
}

void AgastFeatureDetector_Close(AgastFeatureDetector a) {
    delete a;
}

struct KeyPoints AgastFeatureDetector_Detect(AgastFeatureDetector a, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

BRISK BRISK_Create() {
    // TODO: params
    return new cv::Ptr<cv::BRISK>(cv::BRISK::create());
}

void BRISK_Close(BRISK b) {
    delete b;
}

struct KeyPoints BRISK_Detect(BRISK b, Mat src) {
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

struct KeyPoints BRISK_DetectAndCompute(BRISK b, Mat src, Mat mask, Mat desc) {
    std::vector<cv::KeyPoint> detected;
    (*b)->detectAndCompute(*src, *mask, detected, *desc);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

FastFeatureDetector FastFeatureDetector_Create() {
    // TODO: params
    return new cv::Ptr<cv::FastFeatureDetector>(cv::FastFeatureDetector::create());
}

void FastFeatureDetector_Close(FastFeatureDetector f) {
    delete f;
}

struct KeyPoints FastFeatureDetector_Detect(FastFeatureDetector f, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*f)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

ORB ORB_Create() {
    // TODO: params
    return new cv::Ptr<cv::ORB>(cv::ORB::create());
}

void ORB_Close(ORB o) {
    delete o;
}

struct KeyPoints ORB_Detect(ORB o, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*o)->detect(*src, detected);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

struct KeyPoints ORB_DetectAndCompute(ORB o, Mat src, Mat mask, Mat desc) {
    std::vector<cv::KeyPoint> detected;
    (*o)->detectAndCompute(*src, *mask, detected, *desc);

    KeyPoint* kps = new KeyPoint[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      KeyPoint k = {detected[i].pt.x, detected[i].pt.y, detected[i].size, detected[i].angle,
        detected[i].response, detected[i].octave, detected[i].class_id};
      kps[i] = k;
    }
    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

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
