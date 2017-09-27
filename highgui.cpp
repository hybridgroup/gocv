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

int Window_WaitKey(int delay = 0) {
    return cv::waitKey(delay);
}