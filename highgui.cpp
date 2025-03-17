#include "highgui_gocv.h"

void Window_SetMouseCallback(char* winname, mouse_callback on_mouse) {
    try {
        cv::setMouseCallback(winname, on_mouse, (void*)winname);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

// Window
void Window_New(const char* winname, int flags) {
    try {
        cv::namedWindow(winname, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Window_Close(const char* winname) {
    try {
        cv::destroyWindow(winname);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

OpenCVResult Window_IMShow(const char* winname, Mat mat) {
    try {
        cv::imshow(winname, *mat);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

double Window_GetProperty(const char* winname, int flag) {
    try {
        return cv::getWindowProperty(winname, flag);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

OpenCVResult Window_SetProperty(const char* winname, int flag, double value) {
    try {
        cv::setWindowProperty(winname, flag, value);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Window_SetTitle(const char* winname, const char* title) {
    try {
        cv::setWindowTitle(winname, title);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

int Window_WaitKey(int delay = 0) {
    try {
        return cv::waitKey(delay);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

int Window_WaitKeyEx(int delay = 0) {
    try {
        return cv::waitKeyEx(delay);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

int Window_PollKey(void) {
    try {
        return cv::pollKey();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

OpenCVResult Window_Move(const char* winname, int x, int y) {
    try {
        cv::moveWindow(winname, x, y);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Window_Resize(const char* winname, int width, int height) {
    try {
        cv::resizeWindow(winname, width, height);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

struct Rect Window_SelectROI(const char* winname, Mat img) {
    try {
        cv::Rect bRect = cv::selectROI(winname, *img);
        Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
        return r;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rect r = {0, 0, 0, 0};
        return r;
    }
}

struct Rects Window_SelectROIs(const char* winname, Mat img) {
    try {
        std::vector<cv::Rect> rois;
        cv::selectROIs(winname, *img, rois);
        Rect* rects = new Rect[rois.size()];
    
        for (size_t i = 0; i < rois.size(); ++i) {
            Rect r = {rois[i].x, rois[i].y, rois[i].width, rois[i].height};
            rects[i] = r;
        }
    
        Rects ret = {rects, (int)rois.size()};
        return ret;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rects ret = {NULL, 0};
        return ret;
    }
}

// Trackbar
void Trackbar_Create(const char* winname, const char* trackname, int max) {
    try {
        cv::createTrackbar(trackname, winname, NULL, max);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Trackbar_CreateWithValue(const char* winname, const char* trackname, int* value, int max) {
    try {
        cv::createTrackbar(trackname, winname, value, max);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int Trackbar_GetPos(const char* winname, const char* trackname) {
    try {
        return cv::getTrackbarPos(trackname, winname);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

void Trackbar_SetPos(const char* winname, const char* trackname, int pos) {
    try {
        cv::setTrackbarPos(trackname, winname, pos);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Trackbar_SetMin(const char* winname, const char* trackname, int pos) {
    try {
        cv::setTrackbarMin(trackname, winname, pos);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Trackbar_SetMax(const char* winname, const char* trackname, int pos) {
    try {
        cv::setTrackbarMax(trackname, winname, pos);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
