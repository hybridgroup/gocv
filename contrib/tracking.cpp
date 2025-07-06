//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_contrib_tracking)

#include "tracking.h"
#include <opencv2/opencv.hpp>

bool TrackerSubclass_Init(Tracker self, Mat image, Rect boundingBox) {
    try {
        cv::Rect bb(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

        (*self)->init(*image, bb);
        return true;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool TrackerSubclass_Update(Tracker self, Mat image, Rect* boundingBox) {
    try {
        cv::Rect bb;
        bool ret = (*self)->update(*image, bb);
        boundingBox->x = int(bb.x);
        boundingBox->y = int(bb.y);
        boundingBox->width = int(bb.width);
        boundingBox->height = int(bb.height);
        return ret;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

TrackerKCF TrackerKCF_Create() {
    try {
        return new cv::Ptr<cv::TrackerKCF>(cv::TrackerKCF::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void TrackerKCF_Close(TrackerKCF self) {
    delete self;
}

TrackerCSRT TrackerCSRT_Create() {
    try {
        return new cv::Ptr<cv::TrackerCSRT>(cv::TrackerCSRT::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void TrackerCSRT_Close(TrackerCSRT self) {
    delete self;
}
