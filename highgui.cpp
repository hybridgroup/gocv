#include "highgui.h"

// Window
void Window_New(const char* winname, int flags) {
    cv::namedWindow(winname, flags);
}

// void Window_Delete(const char* winname) {
//     cv::destroyWindow(winname);
// }

int Window_WaitKey(int delay = 0) {
    return cv::waitKey(delay);
}