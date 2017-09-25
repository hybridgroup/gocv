#include "pvl.h"

// FaceDetector
FaceDetector FaceDetector_New() 
{
    return cv::pvl::FaceDetector::create();
}

void FaceDetector_Delete(FaceDetector f) 
{
    delete f;
}

void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled)
{
    f->setTrackingModeEnabled(enabled);
    return;
}