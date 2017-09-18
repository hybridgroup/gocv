#include "videowriter.h"

VideoWriter VideoWriter_New() {
    return new cv::VideoWriter();
}
  
void VideoWriter_Delete(VideoWriter vw) {
    delete vw;
}
  
void VideoWriter_Open(VideoWriter vw, const char* name, double fps, int width,
      int height) {
    vw->open(name, CV_FOURCC('M', 'J', 'P', 'G'), fps, cv::Size(width, height), true);
}
  
void VideoWriter_OpenWithMat(VideoWriter vw, const char* name, double fps,
      MatVec3b img) {
    vw->open(name, CV_FOURCC('M', 'J', 'P', 'G'), fps, img->size(), true);
}
  
int VideoWriter_IsOpened(VideoWriter vw) {
    return vw->isOpened();
}
  
void VideoWriter_Write(VideoWriter vw, MatVec3b img) {
    *vw << *img;
}
