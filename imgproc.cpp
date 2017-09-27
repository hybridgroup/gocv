#include "imgproc.h"

void CvtColor(Mat src, Mat dst, int code)
{
    cv::cvtColor(*src, *dst, code);
}

void Rectangle(Mat img, Rect r) {
    cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
        cv::Scalar(0, 200, 0), 3, CV_AA);
}
