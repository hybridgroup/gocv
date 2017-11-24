#include "imgproc.h"

void CvtColor(Mat src, Mat dst, int code) {
    cv::cvtColor(*src, *dst, code);
}

void BilateralFilter(Mat src, Mat dst, int d, double sc, double ss) {
    cv::bilateralFilter(*src, *dst, d, sc, ss);
}

void Blur(Mat src, Mat dst, Size ps) {
    cv::Size sz(ps.width, ps.height);
    cv::blur(*src, *dst, sz);
}

void Dilate(Mat src, Mat dst, Mat kernel) {
    cv::dilate(*src, *dst, *kernel);
}

void Erode(Mat src, Mat dst, Mat kernel) {
    cv::erode(*src, *dst, *kernel);
}

struct Moment Moments(Mat src, bool binaryImage) {
    cv::Moments m = cv::moments(*src, binaryImage);
    Moment mom = {m.m00, m.m10, m.m01, m.m20, m.m11, m.m02, m.m30, m.m21, m.m12, m.m03,
        m.mu20, m.mu11, m.mu02, m.mu30, m.mu21, m.mu12, m.mu03,
        m.nu20, m.nu11, m.nu02, m.nu30, m.nu21, m.nu12, m.nu03};
    return mom;
}

struct Rect BoundingRect(Contour con) {
    std::vector<cv::Point> pts;
    for (size_t i = 0; i < con.length; i++) {
        pts.push_back(cv::Point(con.points[i].x, con.points[i].y));
    }
    cv::Rect bRect = cv::boundingRect(pts);
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    return r;
}

double ContourArea(Contour con) {
    std::vector<cv::Point> pts;
    for (size_t i = 0; i < con.length; i++) {
        pts.push_back(cv::Point(con.points[i].x, con.points[i].y));
    }
    return cv::contourArea(pts);
}

struct Contours FindContours(Mat src, int mode, int method) {
    std::vector<std::vector<cv::Point> > contours;
    cv::findContours(*src, contours, mode, method);

    Contour* points = new Contour[contours.size()];
    for (size_t i = 0; i < contours.size(); i++) {
        Point *pts = new Point[contours[i].size()];
        for (size_t j = 0; j < contours[i].size(); j++) {
            Point pt = {contours[i][j].x, contours[i][j].y};
            pts[j] = pt;
        }
        points[i] = Contour{pts, (int)contours[i].size()};
    }

    Contours cons = {points, (int)contours.size()};
    return cons;
}

Mat GetStructuringElement(int shape, Size ksize) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Mat(cv::getStructuringElement(shape, sz));
}

void MorphologyEx(Mat src, Mat dst, int op, Mat kernel) {
    cv::morphologyEx(*src, *dst, op, *kernel);
}

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt) {
    cv::Size sz(ps.width, ps.height);
    cv::GaussianBlur(*src, *dst, sz, sX, sY, bt);
}

void Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta, int borderType) {
    cv::Laplacian(*src, *dst, dDepth, kSize, scale, delta, borderType);
}

void Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta, int borderType) {
    cv::Scharr(*src, *dst, dDepth, dx, dy, scale, delta, borderType);
}

void MedianBlur(Mat src, Mat dst, int ksize) {
    cv::medianBlur(*src, *dst, ksize);
}

void Canny(Mat src, Mat edges, double t1, double t2) {
    cv::Canny(*src, *edges, t1, t2);
}

void CornerSubPix(Mat img, Mat corners, Size winSize, Size zeroZone, TermCriteria criteria) {
    cv::Size wsz(winSize.width, winSize.height);
    cv::Size zsz(zeroZone.width, zeroZone.height);
    cv::cornerSubPix(*img, *corners, wsz, zsz, *criteria);
}

void GoodFeaturesToTrack(Mat img, Mat corners, int maxCorners, double quality, double minDist) {
    cv::goodFeaturesToTrack(*img, *corners, maxCorners, quality, minDist);
}

void HoughCircles(Mat src, Mat circles, int method, double dp, double minDist) {
    cv::HoughCircles(*src, *circles, method, dp, minDist);
}

void HoughLines(Mat src, Mat lines, double rho, double theta, int threshold) {
    cv::HoughLines(*src, *lines, rho, theta, threshold);
}

void HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold) {
    cv::HoughLinesP(*src, *lines, rho, theta, threshold);
}

void Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ) {
    cv:threshold(*src, *dst, thresh, maxvalue, typ);
}

void ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    cv::Point p1(pt1.x, pt1.y);
    cv::Point p2(pt2.x, pt2.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::arrowedLine(*img, p1, p2, c, thickness);
}

void Circle(Mat img, Point center, int radius, Scalar color, int thickness) {
    cv::Point p1(center.x, center.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::circle(*img, p1, radius, c, thickness);
}

void Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    cv::Point p1(pt1.x, pt1.y);
    cv::Point p2(pt2.x, pt2.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::line(*img, p1, p2, c, thickness);
}

void Rectangle(Mat img, Rect r, Scalar color, int thickness) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
        c, thickness, CV_AA);
}

struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness) {
    cv::Size sz = cv::getTextSize(text, fontFace, fontScale, thickness, NULL);
    Size size = {sz.width, sz.height};
    return size;
}

void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale,
    Scalar color, int thickness) {
    cv::Point pt(org.x, org.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::putText(*img, text, pt, fontFace, fontScale, c, thickness);
}

void Resize(Mat src, Mat dst, Size dsize, double fx, double fy, int interp) {
  cv::Size sz(dsize.width, dsize.height);
  cv::resize(*src, *dst, sz, fx, fy, interp);
}
