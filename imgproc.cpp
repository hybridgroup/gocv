#include "imgproc.h"

void CvtColor(Mat src, Mat dst, int code)
{
    cv::cvtColor(*src, *dst, code);
}

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt)
{
    cv:GaussianBlur(src, dst, ps, sX, sY, bt);
}

void Rectangle(Mat img, Rect r, Scalar color) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
        c, 3, CV_AA);
}

struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness)
{
    cv::Size sz = cv::getTextSize(text, fontFace, fontScale, thickness, NULL);
    Size size = {sz.width, sz.height};
    return size;
}

void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale, 
    Scalar color, int thickness)
{
    cv::Point pt(org.x, org.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::putText(*img, text, pt, fontFace, fontScale, c, thickness);
}
