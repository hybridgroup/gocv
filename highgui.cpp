#include "highgui.h"

// Window
void Window_New(const char* winname, int flags) {
    cv::namedWindow(winname, flags);
}

void Window_Close(const char* winname) {
    cv::destroyWindow(winname);
}

void Window_IMShow(const char* winname, Mat mat) {
    cv::imshow(winname, *mat);
}

void Window_SetWindowProperty(const char* winname, int flag, double value) {
    cv::setWindowProperty(winname, flag, value);
}

int Window_WaitKey(int delay = 0) {
    return cv::waitKey(delay);
}

// Trackbar
void Trackbar_Create(const char* winname, const char* trackname, int max) {
    cv::createTrackbar(trackname, winname, NULL, max);
}

int Trackbar_GetPos(const char* winname, const char* trackname) {
    return cv::getTrackbarPos(trackname, winname);
}

void Trackbar_SetPos(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarPos(trackname, winname, pos);
}

void Trackbar_SetMin(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarMin(trackname, winname, pos);
}

void Trackbar_SetMax(const char* winname, const char* trackname, int pos) {
    cv::setTrackbarMax(trackname, winname, pos);
}
