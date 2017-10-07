#include "videoio.h"

// VideoWriter
VideoCapture VideoCapture_New() {
    return new cv::VideoCapture();
}
  
void VideoCapture_Close(VideoCapture v) {
    delete v;
}
  
int VideoCapture_Open(VideoCapture v, const char* uri) {
    return v->open(uri);
}
  
int VideoCapture_OpenDevice(VideoCapture v, int device) {
    return v->open(device);
}
  
void VideoCapture_Set(VideoCapture v, int prop, int param) {
    v->set(prop, param);
}
  
int VideoCapture_IsOpened(VideoCapture v) {
    return v->isOpened();
}
  
int VideoCapture_Read(VideoCapture v, Mat buf) {
    return v->read(*buf);
}
  
void VideoCapture_Grab(VideoCapture v, int skip) {
    for (int i =0; i < skip; i++) {
        v->grab();
    }
}

// VideoWriter
VideoWriter VideoWriter_New() {
    return new cv::VideoWriter();
}
  
void VideoWriter_Close(VideoWriter vw) {
    delete vw;
}
  
void VideoWriter_Open(VideoWriter vw, const char* name, double fps, int width,
      int height) {
    vw->open(name, CV_FOURCC('M', 'J', 'P', 'G'), fps, cv::Size(width, height), true);
}

int VideoWriter_IsOpened(VideoWriter vw) {
    return vw->isOpened();
}
  
void VideoWriter_Write(VideoWriter vw, Mat img) {
    *vw << *img;
}
