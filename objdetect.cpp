#include "objdetect.h"

// CascadeClassifier

CascadeClassifier CascadeClassifier_New() {
    return new cv::CascadeClassifier();
}
  
void CascadeClassifier_Close(CascadeClassifier cs) {
    delete cs;
}
  
int CascadeClassifier_Load(CascadeClassifier cs, const char* name) {
    return cs->load(name);
}
  
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, Mat img) {
    std::vector<cv::Rect> detected;
    cs->detectMultiScale(*img, detected); // TODO control default parameter
    Rect* rects = new Rect[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
      rects[i] = r;
    }
    Rects ret = {rects, (int)detected.size()};
    return ret;
}

// HOGDescriptor

HOGDescriptor HOGDescriptor_New() {
    return new cv::HOGDescriptor();
}

void HOGDescriptor_Close(HOGDescriptor hog) {
    delete hog;
}

int HOGDescriptor_Load(HOGDescriptor hog, const char* name) {
    return hog->load(name);
}

struct Rects HOGDescriptor_DetectMultiScale(HOGDescriptor hog, Mat img) {
    std::vector<cv::Rect> detected;
    hog->detectMultiScale(*img, detected); // TODO control default parameter
    Rect* rects = new Rect[detected.size()];
    for (size_t i = 0; i < detected.size(); ++i) {
      Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
      rects[i] = r;
    }
    Rects ret = {rects, (int)detected.size()};
    return ret;
}

Mat HOG_GetDefaultPeopleDetector() {
    return new cv::Mat(cv::HOGDescriptor::getDefaultPeopleDetector());
}

void HOGDescriptor_SetSVMDetector(HOGDescriptor hog, Mat det) {
    hog->setSVMDetector(*det);
}
