#include "tracking.h"
#include <opencv2/opencv.hpp>


bool Tracker_Init(Tracker self, Mat image, Rect boundingBox) {
    cv::Rect2d c_boundingBox(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

    bool ret = (*self)->init(*image, c_boundingBox);
    return ret;
}


bool Tracker_Update(Tracker self, Mat image, Rect *boundingBox) {
    cv::Rect2d cBox;
    bool ret = (*self)->update(*image, cBox);
    boundingBox->x = int(cBox.x);
    boundingBox->y = int(cBox.y);
    boundingBox->width = int(cBox.width);
    boundingBox->height = int(cBox.height);
    return ret;
}




TrackerMil TrackerMil_Create() {
    return new cv::Ptr<cv::TrackerMIL>(cv::TrackerMIL::create());
}

void TrackerMil_Close(TrackerMil self){
    delete self;
}


TrackerBoosting TrackerBoosting_Create() {
    return new cv::Ptr<cv::TrackerBoosting>(cv::TrackerBoosting::create());
}

void TrackerBoosting_Close(TrackerBoosting self){
	delete self;
}


TrackerMedianFlow TrackerMedianFlow_Create() {
    return new cv::Ptr<cv::TrackerMedianFlow>(cv::TrackerMedianFlow::create());
}

void TrackerMedianFlow_Close(TrackerMedianFlow self){
	delete self;
}


TrackerTld TrackerTld_Create() {
    return new cv::Ptr<cv::TrackerTLD>(cv::TrackerTLD::create());
}

void TrackerTld_Close(TrackerTld self){
    delete self;
}


TrackerKcf TrackerKcf_Create() {
    return new cv::Ptr<cv::TrackerKCF>(cv::TrackerKCF::create());
}

void TrackerKcf_Close(TrackerKcf self){
    delete self;
}


TrackerMosse TrackerMosse_Create() {
    return new cv::Ptr<cv::TrackerMOSSE>(cv::TrackerMOSSE::create());
}

void TrackerMosse_Close(TrackerMosse self){
    delete self;
}



TrackerCsrt TrackerCsrt_Create() {
    return new cv::Ptr<cv::TrackerCSRT>(cv::TrackerCSRT::create());
}

void TrackerCsrt_Close(TrackerCsrt self){
    delete self;
}
