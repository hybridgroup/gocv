#include "facemarkLBF.h"

LBPHFaceMark CreateLBPHFaceMark() {
    return new cv::Ptr<cv::face::Facemark>(cv::face::FacemarkLBF::create());
}

void LBPHFaceMark_LoadModel(LBPHFaceMark fm, const char*  model) {
    (*fm)->loadModel(model);
}

bool LBPHFaceMark_Fit(LBPHFaceMark fm, Mat frame, struct Rects faces, Points2fVector landmarks) {
    std::vector<cv::Rect> _faces;

    for (int i = 0; i < faces.length; ++i) {
        _faces.push_back(cv::Rect(
            faces.rects[i].x,
            faces.rects[i].y,
            faces.rects[i].width,
            faces.rects[i].height
        ));
    }
    return (*fm)->fit(*frame, _faces, *landmarks);
}