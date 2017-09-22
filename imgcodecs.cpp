#include "imgcodecs.h"

// Image
Mat Image_IMRead(const char* filename, int flags) {
    cv::Mat img = cv::imread(filename, flags);
    return new cv::Mat(img);
}

bool Image_IMWrite(const char* filename, Mat img) {
    return cv::imwrite(filename, *img);
}

