#include "photo.h"

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags) {
    cv::Point pt(p.x, p.y);
    cv::seamlessClone(*src, *dst, *mask, pt, *blend, flags);
}
