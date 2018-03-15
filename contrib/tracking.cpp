#include "tracking.h"
#include <opencv2/opencv.hpp>


bool Tracker_Init(Tracker self, Mat image, Rect boundingBox) {
    cv::Rect2d bb(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

    bool ret = (*self)->init(*image, bb);
    return ret;
}

bool Tracker_Update(Tracker self, Mat image, Rect* boundingBox) {
    cv::Rect2d bb;
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

TrackerBoosting TrackerBoosting_Create() {
    return new cv::Ptr<cv::TrackerBoosting>(cv::TrackerBoosting::create());
}

void TrackerBoosting_Close(TrackerBoosting self) {
    delete self;
}

TrackerMedianFlow TrackerMedianFlow_Create() {
    return new cv::Ptr<cv::TrackerMedianFlow>(cv::TrackerMedianFlow::create());
}

void TrackerMedianFlow_Close(TrackerMedianFlow self) {
    delete self;
}

TrackerTLD TrackerTLD_Create() {
    return new cv::Ptr<cv::TrackerTLD>(cv::TrackerTLD::create());
}

void TrackerTLD_Close(TrackerTLD self) {
    delete self;
}

TrackerKCF TrackerKCF_Create() {
    return new cv::Ptr<cv::TrackerKCF>(cv::TrackerKCF::create());
}

void TrackerKCF_Close(TrackerKCF self) {
    delete self;
}

TrackerMOSSE TrackerMOSSE_Create() {
    return new cv::Ptr<cv::TrackerMOSSE>(cv::TrackerMOSSE::create());
}

void TrackerMOSSE_Close(TrackerMOSSE self) {
    delete self;
}

TrackerCSRT TrackerCSRT_Create() {
    return new cv::Ptr<cv::TrackerCSRT>(cv::TrackerCSRT::create());
}

void TrackerCSRT_Close(TrackerCSRT self) {
    delete self;
}
