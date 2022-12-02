#include "video.h"

BackgroundSubtractorMOG2 BackgroundSubtractorMOG2_Create() {
    return new cv::Ptr<cv::BackgroundSubtractorMOG2>(cv::createBackgroundSubtractorMOG2());
}

BackgroundSubtractorMOG2 BackgroundSubtractorMOG2_CreateWithParams(int history, double varThreshold, bool detectShadows) {
    return new cv::Ptr<cv::BackgroundSubtractorMOG2>(cv::createBackgroundSubtractorMOG2(history,varThreshold,detectShadows));
}

BackgroundSubtractorKNN BackgroundSubtractorKNN_Create() {
    return new cv::Ptr<cv::BackgroundSubtractorKNN>(cv::createBackgroundSubtractorKNN());
}

BackgroundSubtractorKNN BackgroundSubtractorKNN_CreateWithParams(int history, double dist2Threshold, bool detectShadows) {
    return new cv::Ptr<cv::BackgroundSubtractorKNN>(cv::createBackgroundSubtractorKNN(history,dist2Threshold,detectShadows));
}

void BackgroundSubtractorMOG2_Close(BackgroundSubtractorMOG2 b) {
    delete b;
}

void BackgroundSubtractorMOG2_Apply(BackgroundSubtractorMOG2 b, Mat src, Mat dst) {
    (*b)->apply(*src, *dst);
}

void BackgroundSubtractorKNN_Close(BackgroundSubtractorKNN k) {
    delete k;
}

void BackgroundSubtractorKNN_Apply(BackgroundSubtractorKNN k, Mat src, Mat dst) {
    (*k)->apply(*src, *dst);
}

void CalcOpticalFlowFarneback(Mat prevImg, Mat nextImg, Mat flow, double scale, int levels,
                              int winsize, int iterations, int polyN, double polySigma, int flags) {
    cv::calcOpticalFlowFarneback(*prevImg, *nextImg, *flow, scale, levels, winsize, iterations, polyN,
                                 polySigma, flags);
}

void CalcOpticalFlowPyrLK(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err) {
    cv::calcOpticalFlowPyrLK(*prevImg, *nextImg, *prevPts, *nextPts, *status, *err);
}

void CalcOpticalFlowPyrLKWithParams(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err, Size winSize, int maxLevel, TermCriteria criteria, int flags, double minEigThreshold){
    cv::Size sz(winSize.width, winSize.height);
    cv::calcOpticalFlowPyrLK(*prevImg, *nextImg, *prevPts, *nextPts, *status, *err, sz, maxLevel, *criteria, flags, minEigThreshold);
}

double FindTransformECC(Mat templateImage, Mat inputImage, Mat warpMatrix, int motionType, TermCriteria criteria, Mat inputMask, int gaussFiltSize){
    return cv::findTransformECC(*templateImage, *inputImage, *warpMatrix, motionType, *criteria, *inputMask, gaussFiltSize);
}

bool Tracker_Init(Tracker self, Mat image, Rect boundingBox) {
    cv::Rect bb(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

    (*self)->init(*image, bb);
    return true;
}

bool Tracker_Update(Tracker self, Mat image, Rect* boundingBox) {
    cv::Rect bb;
    bool ret = (*self)->update(*image, bb);
    boundingBox->x = int(bb.x);
    boundingBox->y = int(bb.y);
    boundingBox->width = int(bb.width);
    boundingBox->height = int(bb.height);
    return ret;
}

TrackerMIL TrackerMIL_Create() {
    return new cv::Ptr<cv::TrackerMIL>(cv::TrackerMIL::create());
}

void TrackerMIL_Close(TrackerMIL self) {
    delete self;
}

KalmanFilter KalmanFilter_New(int dynamParams, int measureParams) {
    return new cv::KalmanFilter(dynamParams, measureParams, 0, CV_32F);
}

KalmanFilter KalmanFilter_NewWithParams(int dynamParams, int measureParams, int controlParams, int type) {
    return new cv::KalmanFilter(dynamParams, measureParams, controlParams, type);
}

void KalmanFilter_Init(KalmanFilter kf, int dynamParams, int measureParams) {
  kf->init(dynamParams, measureParams, 0, CV_32F);
}

void KalmanFilter_InitWithParams(KalmanFilter kf, int dynamParams, int measureParams, int controlParams, int type) {
  kf->init(dynamParams, measureParams, controlParams, type);
}

void KalmanFilter_Close(KalmanFilter kf) {
    delete kf;
}

Mat KalmanFilter_Predict(KalmanFilter kf) {
 return new cv::Mat(kf->predict());
}

Mat KalmanFilter_PredictWithParams(KalmanFilter kf, Mat control) {
 return new cv::Mat(kf->predict(*control));
}

Mat KalmanFilter_Correct(KalmanFilter kf, Mat measurement) {
  return new cv::Mat(kf->correct(*measurement));
}

Mat KalmanFilter_GetStatePre(KalmanFilter kf) {
  return new cv::Mat(kf->statePre);
}

Mat KalmanFilter_GetStatePost(KalmanFilter kf) {
  return new cv::Mat(kf->statePost);
}

Mat KalmanFilter_GetTransitionMatrix(KalmanFilter kf) {
  return new cv::Mat(kf->transitionMatrix);
}

Mat KalmanFilter_GetControlMatrix(KalmanFilter kf) {
  return new cv::Mat(kf->controlMatrix);
}

Mat KalmanFilter_GetMeasurementMatrix(KalmanFilter kf) {
  return new cv::Mat(kf->measurementMatrix);
}

Mat KalmanFilter_GetProcessNoiseCov(KalmanFilter kf) {
  return new cv::Mat(kf->processNoiseCov);
}

Mat KalmanFilter_GetMeasurementNoiseCov(KalmanFilter kf) {
  return new cv::Mat(kf->measurementNoiseCov);
}

Mat KalmanFilter_GetErrorCovPre(KalmanFilter kf) {
  return new cv::Mat(kf->errorCovPre);
}

Mat KalmanFilter_GetGain(KalmanFilter kf) {
  return new cv::Mat(kf->gain);
}

Mat KalmanFilter_GetErrorCovPost(KalmanFilter kf) {
  return new cv::Mat(kf->errorCovPost);
}

Mat KalmanFilter_GetTemp1(KalmanFilter kf) {
  return new cv::Mat(kf->temp1);
}

Mat KalmanFilter_GetTemp2(KalmanFilter kf) {
  return new cv::Mat(kf->temp2);
}

Mat KalmanFilter_GetTemp3(KalmanFilter kf) {
  return new cv::Mat(kf->temp3);
}

Mat KalmanFilter_GetTemp4(KalmanFilter kf) {
  return new cv::Mat(kf->temp4);
}

Mat KalmanFilter_GetTemp5(KalmanFilter kf) {
  return new cv::Mat(kf->temp5);
}

void KalmanFilter_SetStatePre(KalmanFilter kf, Mat statePre) {
  kf->statePre = *statePre;
}

void KalmanFilter_SetStatePost(KalmanFilter kf, Mat statePost) {
  kf->statePost = *statePost;
}

void KalmanFilter_SetTransitionMatrix(KalmanFilter kf, Mat transitionMatrix) {
  kf->transitionMatrix = *transitionMatrix;
}

void KalmanFilter_SetControlMatrix(KalmanFilter kf, Mat controlMatrix) {
  kf->controlMatrix = *controlMatrix;
}

void KalmanFilter_SetMeasurementMatrix(KalmanFilter kf, Mat measurementMatrix) {
  kf->measurementMatrix = *measurementMatrix;
}

void KalmanFilter_SetProcessNoiseCov(KalmanFilter kf, Mat processNoiseCov) {
  kf->processNoiseCov = *processNoiseCov;
}

void KalmanFilter_SetMeasurementNoiseCov(KalmanFilter kf, Mat measurementNoiseCov) {
  kf->measurementNoiseCov = *measurementNoiseCov;
}

void KalmanFilter_SetErrorCovPre(KalmanFilter kf, Mat errorCovPre) {
  kf->errorCovPre = *errorCovPre;
}

void KalmanFilter_SetGain(KalmanFilter kf, Mat gain) {
  kf->gain = *gain;
}

void KalmanFilter_SetErrorCovPost(KalmanFilter kf, Mat errorCovPost) {
  kf->errorCovPost = *errorCovPost;
}
