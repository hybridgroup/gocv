#include "bgsegm.h"

BackgroundSubtractorCNT BackgroundSubtractorCNT_Create() {
    try {
        return new cv::Ptr<cv::bgsegm::BackgroundSubtractorCNT>(cv::bgsegm::createBackgroundSubtractorCNT());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void BackgroundSubtractorCNT_Close(BackgroundSubtractorCNT b) {
    delete b;
}

OpenCVResult BackgroundSubtractorCNT_Apply(BackgroundSubtractorCNT b, Mat src, Mat dst) {
    try {
        (*b)->apply(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
