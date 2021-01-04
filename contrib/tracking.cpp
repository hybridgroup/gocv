#include "tracking.h"
#include <opencv2/opencv.hpp>


bool TrackerSubclass_Init(Tracker self, Mat image, Rect boundingBox) {
    cv::Rect bb(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

    (*self)->init(*image, bb);
    return true;
}

bool TrackerSubclass_Update(Tracker self, Mat image, Rect* boundingBox) {
    cv::Rect bb;
    bool ret = (*self)->update(*image, bb);
    boundingBox->x = int(bb.x);
    boundingBox->y = int(bb.y);
    boundingBox->width = int(bb.width);
    boundingBox->height = int(bb.height);
    return ret;
}

TrackerKCF TrackerKCF_Create() {
    return new cv::Ptr<cv::TrackerKCF>(cv::TrackerKCF::create());
}

void TrackerKCF_Close(TrackerKCF self) {
    delete self;
}

TrackerCSRT TrackerCSRT_Create() {
    return new cv::Ptr<cv::TrackerCSRT>(cv::TrackerCSRT::create());
}

void TrackerCSRT_Close(TrackerCSRT self) {
    delete self;
}
