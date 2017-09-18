#include "videocapture.h"

VideoCapture VideoCapture_New() {
    return new cv::VideoCapture();
}
  
void VideoCapture_Delete(VideoCapture v) {
    delete v;
}
  
int VideoCapture_Open(VideoCapture v, const char* uri) {
    return v->open(uri);
}
  
int VideoCapture_OpenDevice(VideoCapture v, int device) {
    return v->open(device);
}
  
void VideoCapture_Release(VideoCapture v) {
    v->release();
}
  
void VideoCapture_Set(VideoCapture v, int prop, int param) {
    v->set(prop, param);
}
  
int VideoCapture_IsOpened(VideoCapture v) {
    return v->isOpened();
}
  
int VideoCapture_Read(VideoCapture v, MatVec3b buf) {
    return v->read(*buf);
}
  
void VideoCapture_Grab(VideoCapture v, int skip) {
    for (int i =0; i < skip; i++) {
        v->grab();
    }
}
