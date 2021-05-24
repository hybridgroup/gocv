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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
        kps[i] = k;
    }

    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

GFTTDetector GFTTDetector_Create() {
    // TODO: params
    return new cv::Ptr<cv::GFTTDetector>(cv::GFTTDetector::create());
}

void GFTTDetector_Close(GFTTDetector a) {
    delete a;
}

struct KeyPoints GFTTDetector_Detect(GFTTDetector a, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detect(*src, detected);

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

KAZE KAZE_Create() {
    // TODO: params
    return new cv::Ptr<cv::KAZE>(cv::KAZE::create());
}

void KAZE_Close(KAZE a) {
    delete a;
}

struct KeyPoints KAZE_Detect(KAZE a, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detect(*src, detected);

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

struct KeyPoints KAZE_DetectAndCompute(KAZE a, Mat src, Mat mask, Mat desc) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detectAndCompute(*src, *mask, detected, *desc);

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

MSER MSER_Create() {
    // TODO: params
    return new cv::Ptr<cv::MSER>(cv::MSER::create());
}

void MSER_Close(MSER a) {
    delete a;
}

struct KeyPoints MSER_Detect(MSER a, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*a)->detect(*src, detected);

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

FastFeatureDetector FastFeatureDetector_Create() {
    return new cv::Ptr<cv::FastFeatureDetector>(cv::FastFeatureDetector::create());
}

void FastFeatureDetector_Close(FastFeatureDetector f) {
    delete f;
}

FastFeatureDetector FastFeatureDetector_CreateWithParams(int threshold, bool nonmaxSuppression, int type) {
    return new cv::Ptr<cv::FastFeatureDetector>(cv::FastFeatureDetector::create(threshold,nonmaxSuppression,static_cast<cv::FastFeatureDetector::DetectorType>(type)));
}

struct KeyPoints FastFeatureDetector_Detect(FastFeatureDetector f, Mat src) {
    std::vector<cv::KeyPoint> detected;
    (*f)->detect(*src, detected);

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

ORB ORB_Create() {
    return new cv::Ptr<cv::ORB>(cv::ORB::create());
}

ORB ORB_CreateWithParams(int nfeatures, float scaleFactor, int nlevels, int edgeThreshold, int firstLevel, int WTA_K, int scoreType, int patchSize, int fastThreshold) {
    return new cv::Ptr<cv::ORB>(cv::ORB::create(nfeatures, scaleFactor, nlevels, edgeThreshold, firstLevel, WTA_K, static_cast<cv::ORB::ScoreType>(scoreType), patchSize, fastThreshold));
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
        kps[i] = k;
    }

    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

cv::SimpleBlobDetector::Params ConvertCParamsToCPPParams(SimpleBlobDetectorParams params) {
    cv::SimpleBlobDetector::Params converted;

    converted.blobColor = params.blobColor;
    converted.filterByArea = params.filterByArea;
    converted.filterByCircularity = params.filterByCircularity;
    converted.filterByColor = params.filterByColor;
    converted.filterByConvexity = params.filterByConvexity;
    converted.filterByInertia = params.filterByInertia;
    converted.maxArea = params.maxArea;
    converted.maxCircularity = params.maxCircularity;
    converted.maxConvexity = params.maxConvexity;
    converted.maxInertiaRatio = params.maxInertiaRatio;
    converted.maxThreshold = params.maxThreshold;
    converted.minArea = params.minArea;
    converted.minCircularity = params.minCircularity;
    converted.minConvexity = params.minConvexity;
    converted.minDistBetweenBlobs = params.minDistBetweenBlobs;
    converted.minInertiaRatio = params.minInertiaRatio;
    converted.minRepeatability = params.minRepeatability;
    converted.minThreshold = params.minThreshold;
    converted.thresholdStep = params.thresholdStep;

    return converted;
}

SimpleBlobDetectorParams ConvertCPPParamsToCParams(cv::SimpleBlobDetector::Params params) {
    SimpleBlobDetectorParams converted;

    converted.blobColor = params.blobColor;
    converted.filterByArea = params.filterByArea;
    converted.filterByCircularity = params.filterByCircularity;
    converted.filterByColor = params.filterByColor;
    converted.filterByConvexity = params.filterByConvexity;
    converted.filterByInertia = params.filterByInertia;
    converted.maxArea = params.maxArea;
    converted.maxCircularity = params.maxCircularity;
    converted.maxConvexity = params.maxConvexity;
    converted.maxInertiaRatio = params.maxInertiaRatio;
    converted.maxThreshold = params.maxThreshold;
    converted.minArea = params.minArea;
    converted.minCircularity = params.minCircularity;
    converted.minConvexity = params.minConvexity;
    converted.minDistBetweenBlobs = params.minDistBetweenBlobs;
    converted.minInertiaRatio = params.minInertiaRatio;
    converted.minRepeatability = params.minRepeatability;
    converted.minThreshold = params.minThreshold;
    converted.thresholdStep = params.thresholdStep;

    return converted;
}

SimpleBlobDetector SimpleBlobDetector_Create_WithParams(SimpleBlobDetectorParams params){
    cv::SimpleBlobDetector::Params actualParams;
    return new cv::Ptr<cv::SimpleBlobDetector>(cv::SimpleBlobDetector::create(ConvertCParamsToCPPParams(params)));
}

SimpleBlobDetector SimpleBlobDetector_Create() {
    return new cv::Ptr<cv::SimpleBlobDetector>(cv::SimpleBlobDetector::create());
}

SimpleBlobDetectorParams SimpleBlobDetectorParams_Create() {
    return ConvertCPPParamsToCParams(cv::SimpleBlobDetector::Params());
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
                      detected[i].response, detected[i].octave, detected[i].class_id
                     };
        kps[i] = k;
    }

    KeyPoints ret = {kps, (int)detected.size()};
    return ret;
}

BFMatcher BFMatcher_Create() {
    return new cv::Ptr<cv::BFMatcher>(cv::BFMatcher::create());
}

BFMatcher BFMatcher_CreateWithParams(int normType, bool crossCheck) {
    return new cv::Ptr<cv::BFMatcher>(cv::BFMatcher::create(normType, crossCheck));
}

void BFMatcher_Close(BFMatcher b) {
    delete b;
}

struct MultiDMatches BFMatcher_KnnMatch(BFMatcher b, Mat query, Mat train, int k) {
    std::vector< std::vector<cv::DMatch> > matches;
    (*b)->knnMatch(*query, *train, matches, k);

    DMatches *dms = new DMatches[matches.size()];
    for (size_t i = 0; i < matches.size(); ++i) {
        DMatch *dmatches = new DMatch[matches[i].size()];
        for (size_t j = 0; j < matches[i].size(); ++j) {
            DMatch dmatch = {matches[i][j].queryIdx, matches[i][j].trainIdx, matches[i][j].imgIdx,
                             matches[i][j].distance};
            dmatches[j] = dmatch;
        }
        dms[i] = {dmatches, (int) matches[i].size()};
    }
    MultiDMatches ret = {dms, (int) matches.size()};
    return ret;
}

struct MultiDMatches BFMatcher_KnnMatchWithParams(BFMatcher b, Mat query, Mat train, int k, Mat mask, bool compactResult) {
    std::vector< std::vector<cv::DMatch> > matches;
    (*b)->knnMatch(*query, *train, matches, k, *mask, compactResult);

    DMatches *dms = new DMatches[matches.size()];
    for (size_t i = 0; i < matches.size(); ++i) {
        DMatch *dmatches = new DMatch[matches[i].size()];
        for (size_t j = 0; j < matches[i].size(); ++j) {
            DMatch dmatch = {matches[i][j].queryIdx, matches[i][j].trainIdx, matches[i][j].imgIdx,
                             matches[i][j].distance};
            dmatches[j] = dmatch;
        }
        dms[i] = {dmatches, (int) matches[i].size()};
    }
    MultiDMatches ret = {dms, (int) matches.size()};
    return ret;
}

FlannBasedMatcher FlannBasedMatcher_Create() {
    return new cv::Ptr<cv::FlannBasedMatcher>(cv::FlannBasedMatcher::create());
}

void FlannBasedMatcher_Close(FlannBasedMatcher f) {
    delete f;
}

struct MultiDMatches FlannBasedMatcher_KnnMatch(FlannBasedMatcher f, Mat query, Mat train, int k) {
    std::vector< std::vector<cv::DMatch> > matches;
    (*f)->knnMatch(*query, *train, matches, k);

    DMatches *dms = new DMatches[matches.size()];
    for (size_t i = 0; i < matches.size(); ++i) {
        DMatch *dmatches = new DMatch[matches[i].size()];
        for (size_t j = 0; j < matches[i].size(); ++j) {
            DMatch dmatch = {matches[i][j].queryIdx, matches[i][j].trainIdx, matches[i][j].imgIdx,
                             matches[i][j].distance};
            dmatches[j] = dmatch;
        }
        dms[i] = {dmatches, (int) matches[i].size()};
    }
    MultiDMatches ret = {dms, (int) matches.size()};
    return ret;
}

struct MultiDMatches FlannBasedMatcher_KnnMatchWithParams(FlannBasedMatcher f, Mat query, Mat train, int k, Mat mask, bool compactResult) {
    std::vector< std::vector<cv::DMatch> > matches;
    (*f)->knnMatch(*query, *train, matches, k, *mask, compactResult);

    DMatches *dms = new DMatches[matches.size()];
    for (size_t i = 0; i < matches.size(); ++i) {
        DMatch *dmatches = new DMatch[matches[i].size()];
        for (size_t j = 0; j < matches[i].size(); ++j) {
            DMatch dmatch = {matches[i][j].queryIdx, matches[i][j].trainIdx, matches[i][j].imgIdx,
                             matches[i][j].distance};
            dmatches[j] = dmatch;
        }
        dms[i] = {dmatches, (int) matches[i].size()};
    }
    MultiDMatches ret = {dms, (int) matches.size()};
    return ret;
}

void DrawKeyPoints(Mat src, struct KeyPoints kp, Mat dst, Scalar s, int flags) {
        std::vector<cv::KeyPoint> keypts;
        cv::KeyPoint keypt;

        for (int i = 0; i < kp.length; ++i) {
                keypt = cv::KeyPoint(kp.keypoints[i].x, kp.keypoints[i].y,
                                kp.keypoints[i].size, kp.keypoints[i].angle, kp.keypoints[i].response,
                                kp.keypoints[i].octave, kp.keypoints[i].classID);
                keypts.push_back(keypt);
        }

        cv::Scalar color = cv::Scalar(s.val1, s.val2, s.val3, s.val4);

        cv::drawKeypoints(*src, keypts, *dst, color, static_cast<cv::DrawMatchesFlags>(flags));
}

SIFT SIFT_Create() {
    // TODO: params
    return new cv::Ptr<cv::SIFT>(cv::SIFT::create());
}

void SIFT_Close(SIFT d) {
    delete d;
}

struct KeyPoints SIFT_Detect(SIFT d, Mat src) {
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

struct KeyPoints SIFT_DetectAndCompute(SIFT d, Mat src, Mat mask, Mat desc) {
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

void DrawMatches(Mat img1, struct KeyPoints kp1, Mat img2, struct KeyPoints kp2, struct DMatches matches1to2, Mat outImg, const Scalar matchesColor, const Scalar pointColor, struct ByteArray matchesMask, int flags) {
    std::vector<cv::KeyPoint> kp1vec, kp2vec;
    cv::KeyPoint keypt;

    for (int i = 0; i < kp1.length; ++i) {
        keypt = cv::KeyPoint(kp1.keypoints[i].x, kp1.keypoints[i].y,
                            kp1.keypoints[i].size, kp1.keypoints[i].angle, kp1.keypoints[i].response,
                            kp1.keypoints[i].octave, kp1.keypoints[i].classID);
        kp1vec.push_back(keypt);
    }

    for (int i = 0; i < kp2.length; ++i) {
        keypt = cv::KeyPoint(kp2.keypoints[i].x, kp2.keypoints[i].y,
                            kp2.keypoints[i].size, kp2.keypoints[i].angle, kp2.keypoints[i].response,
                            kp2.keypoints[i].octave, kp2.keypoints[i].classID);
        kp2vec.push_back(keypt);
    }

    cv::Scalar cvmatchescolor = cv::Scalar(matchesColor.val1, matchesColor.val2, matchesColor.val3, matchesColor.val4);
    cv::Scalar cvpointcolor = cv::Scalar(pointColor.val1, pointColor.val2, pointColor.val3, pointColor.val4);
    
    std::vector<cv::DMatch> dmatchvec;
    cv::DMatch dm;

    for (int i = 0; i < matches1to2.length; i++) {
        dm = cv::DMatch(matches1to2.dmatches[i].queryIdx, matches1to2.dmatches[i].trainIdx,
                        matches1to2.dmatches[i].imgIdx, matches1to2.dmatches[i].distance);
        dmatchvec.push_back(dm);
    }

    std::vector<char> maskvec;

    for (int i = 0; i < matchesMask.length; i++) {
        maskvec.push_back(matchesMask.data[i]);
    }

    cv::drawMatches(*img1, kp1vec, *img2, kp2vec, dmatchvec, *outImg, cvmatchescolor, cvpointcolor, maskvec, static_cast<cv::DrawMatchesFlags>(flags));
}
