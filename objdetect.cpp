#include "objdetect.h"

// CascadeClassifier
CascadeClassifier CascadeClassifier_New() {
    return new cv::CascadeClassifier();
}
  
void CascadeClassifier_Delete(CascadeClassifier cs) {
    delete cs;
}
  
int CascadeClassifier_Load(CascadeClassifier cs, const char* name) {
    return cs->load(name);
}
  
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, Mat img) {
    std::vector<cv::Rect> faces;
    cs->detectMultiScale(*img, faces); // TODO control default parameter
    Rect* rects = new Rect[faces.size()];
    for (size_t i = 0; i < faces.size(); ++i) {
      Rect r = {faces[i].x, faces[i].y, faces[i].width, faces[i].height};
      rects[i] = r;
    }
    Rects ret = {rects, (int)faces.size()};
    return ret;
}
  