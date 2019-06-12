#include "imgproc.h"

double ArcLength(Contour curve, bool is_closed) {
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < curve.length; i++) {
        pts.push_back(cv::Point(curve.points[i].x, curve.points[i].y));
    }

    return cv::arcLength(pts, is_closed);
}

Contour ApproxPolyDP(Contour curve, double epsilon, bool closed) {
    std::vector<cv::Point> curvePts;

    for (size_t i = 0; i < curve.length; i++) {
        curvePts.push_back(cv::Point(curve.points[i].x, curve.points[i].y));
    }

    std::vector<cv::Point> approxCurvePts;
    cv::approxPolyDP(curvePts, approxCurvePts, epsilon, closed);

    int length = approxCurvePts.size();
    Point* points = new Point[length];

    for (size_t i = 0; i < length; i++) {
        points[i] = (Point){approxCurvePts[i].x, approxCurvePts[i].y};
    }

    return (Contour){points, length};
}

void CvtColor(Mat src, Mat dst, int code) {
    cv::cvtColor(*src, *dst, code);
}

void EqualizeHist(Mat src, Mat dst) {
    cv::equalizeHist(*src, *dst);
}

void CalcHist(struct Mats mats, IntVector chans, Mat mask, Mat hist, IntVector sz, FloatVector rng, bool acc) {
        std::vector<cv::Mat> images;

        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }

        std::vector<int> channels;

        for (int i = 0, *v = chans.val; i < chans.length; ++v, ++i) {
            channels.push_back(*v);
        }

        std::vector<int> histSize;

        for (int i = 0, *v = sz.val; i < sz.length; ++v, ++i) {
            histSize.push_back(*v);
        }

        std::vector<float> ranges;

        float* f;
        int i;
        for (i = 0, f = rng.val; i < rng.length; ++f, ++i) {
            ranges.push_back(*f);
        }

        cv::calcHist(images, channels, *mask, *hist, histSize, ranges, acc);
}

void ConvexHull(Contour points, Mat hull, bool clockwise, bool returnPoints) {
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < points.length; i++) {
        pts.push_back(cv::Point(points.points[i].x, points.points[i].y));
    }

    cv::convexHull(pts, *hull, clockwise, returnPoints);
}

void ConvexityDefects(Contour points, Mat hull, Mat result) {
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < points.length; i++) {
        pts.push_back(cv::Point(points.points[i].x, points.points[i].y));
    }

    cv::convexityDefects(pts, *hull, *result);
}

void BilateralFilter(Mat src, Mat dst, int d, double sc, double ss) {
    cv::bilateralFilter(*src, *dst, d, sc, ss);
}

void Blur(Mat src, Mat dst, Size ps) {
    cv::Size sz(ps.width, ps.height);
    cv::blur(*src, *dst, sz);
}

void BoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    cv::Size sz(ps.width, ps.height);
    cv::boxFilter(*src, *dst, ddepth, sz);
}

void SqBoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    cv::Size sz(ps.width, ps.height);
    cv::sqrBoxFilter(*src, *dst, ddepth, sz);
}

void Dilate(Mat src, Mat dst, Mat kernel) {
    cv::dilate(*src, *dst, *kernel);
}

void Erode(Mat src, Mat dst, Mat kernel) {
    cv::erode(*src, *dst, *kernel);
}

void MatchTemplate(Mat image, Mat templ, Mat result, int method, Mat mask) {
    cv::matchTemplate(*image, *templ, *result, method, *mask);
}

struct Moment Moments(Mat src, bool binaryImage) {
    cv::Moments m = cv::moments(*src, binaryImage);
    Moment mom = {m.m00, m.m10, m.m01, m.m20, m.m11, m.m02, m.m30, m.m21, m.m12, m.m03,
                  m.mu20, m.mu11, m.mu02, m.mu30, m.mu21, m.mu12, m.mu03,
                  m.nu20, m.nu11, m.nu02, m.nu30, m.nu21, m.nu12, m.nu03
                 };
    return mom;
}

void PyrDown(Mat src, Mat dst, Size size, int borderType) {
    cv::Size cvSize(size.width, size.height);
    cv::pyrDown(*src, *dst, cvSize, borderType);
}

void PyrUp(Mat src, Mat dst, Size size, int borderType) {
    cv::Size cvSize(size.width, size.height);
    cv::pyrUp(*src, *dst, cvSize, borderType);
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

void BoxPoints(RotatedRect rect, Mat boxPts){
    cv::Point2f centerPt(rect.center.x , rect.center.y);
    cv::Size2f rSize(rect.size.width, rect.size.height);
    cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
     cv::boxPoints(rotatedRectangle, *boxPts);
}

double ContourArea(Contour con) {
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < con.length; i++) {
        pts.push_back(cv::Point(con.points[i].x, con.points[i].y));
    }

    return cv::contourArea(pts);
}

struct RotatedRect MinAreaRect(Points points){
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < points.length; i++) {
        pts.push_back(cv::Point(points.points[i].x, points.points[i].y));
    }

    cv::RotatedRect cvrect = cv::minAreaRect(pts);

    Point* rpts = new Point[4];
    cv::Point2f* pts4 = new cv::Point2f[4];
    cvrect.points(pts4);

    for (size_t j = 0; j < 4; j++) {
        Point pt = {int(lroundf(pts4[j].x)), int(lroundf(pts4[j].y))};
        rpts[j] = pt;
    }

    delete[] pts4;

    cv::Rect bRect = cvrect.boundingRect();
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    Point centrpt = {int(lroundf(cvrect.center.x)), int(lroundf(cvrect.center.y))};
    Size szsz = {int(lroundf(cvrect.size.width)), int(lroundf(cvrect.size.height))};

    RotatedRect retrect = {(Contour){rpts, 4}, r, centrpt, szsz, cvrect.angle};
    return retrect;
}

void MinEnclosingCircle(Points points, Point2f* center, float* radius){
    std::vector<cv::Point> pts;

    for (size_t i = 0; i < points.length; i++) {
        pts.push_back(cv::Point(points.points[i].x, points.points[i].y));
    }

    cv::Point2f center2f;
    cv::minEnclosingCircle(pts, center2f, *radius);
    center->x = center2f.x;
    center->y = center2f.y;
}

struct Contours FindContours(Mat src, int mode, int method) {
    std::vector<std::vector<cv::Point> > contours;
    cv::findContours(*src, contours, mode, method);

    Contour* points = new Contour[contours.size()];

    for (size_t i = 0; i < contours.size(); i++) {
        Point* pts = new Point[contours[i].size()];

        for (size_t j = 0; j < contours[i].size(); j++) {
            Point pt = {contours[i][j].x, contours[i][j].y};
            pts[j] = pt;
        }

        points[i] = (Contour){pts, (int)contours[i].size()};
    }

    Contours cons = {points, (int)contours.size()};
    return cons;
}

int ConnectedComponents(Mat src, Mat labels, int connectivity, int ltype, int ccltype){
    return cv::connectedComponents(*src, *labels, connectivity, ltype, ccltype);
}


int ConnectedComponentsWithStats(Mat src, Mat labels, Mat stats, Mat centroids,
    int connectivity, int ltype, int ccltype){
    return cv::connectedComponentsWithStats(*src, *labels, *stats, *centroids, connectivity, ltype, ccltype);
}

Mat GetStructuringElement(int shape, Size ksize) {
    cv::Size sz(ksize.width, ksize.height);
    return new cv::Mat(cv::getStructuringElement(shape, sz));
}

Scalar MorphologyDefaultBorderValue(){
    cv::Scalar cs = cv::morphologyDefaultBorderValue();
    return (Scalar){cs[0],cs[1],cs[2],cs[3]};
}

void MorphologyEx(Mat src, Mat dst, int op, Mat kernel) {
    cv::morphologyEx(*src, *dst, op, *kernel);
}

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt) {
    cv::Size sz(ps.width, ps.height);
    cv::GaussianBlur(*src, *dst, sz, sX, sY, bt);
}

void Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta,
               int borderType) {
    cv::Laplacian(*src, *dst, dDepth, kSize, scale, delta, borderType);
}

void Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta,
            int borderType) {
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

void HoughCirclesWithParams(Mat src, Mat circles, int method, double dp, double minDist,
                            double param1, double param2, int minRadius, int maxRadius) {
    cv::HoughCircles(*src, *circles, method, dp, minDist, param1, param2, minRadius, maxRadius);
}

void HoughLines(Mat src, Mat lines, double rho, double theta, int threshold) {
    cv::HoughLines(*src, *lines, rho, theta, threshold);
}

void HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold) {
    cv::HoughLinesP(*src, *lines, rho, theta, threshold);
}

void HoughLinesPWithParams(Mat src, Mat lines, double rho, double theta, int threshold, double minLineLength, double maxLineGap) {
    cv::HoughLinesP(*src, *lines, rho, theta, threshold, minLineLength, maxLineGap);
}

void HoughLinesPointSet(Mat points, Mat lines, int linesMax, int threshold,
                        double minRho, double  maxRho, double rhoStep,
                        double minTheta, double maxTheta, double thetaStep) {
    cv::HoughLinesPointSet(*points, *lines, linesMax, threshold,
                           minRho, maxRho, rhoStep, minTheta, maxTheta, thetaStep );
}

void Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ) {
    cv::threshold(*src, *dst, thresh, maxvalue, typ);
}

void AdaptiveThreshold(Mat src, Mat dst, double maxValue, int adaptiveMethod, int thresholdType,
                       int blockSize, double c) {
    cv::adaptiveThreshold(*src, *dst, maxValue, adaptiveMethod, thresholdType, blockSize, c);
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

void Ellipse(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness) {
    cv::Point p1(center.x, center.y);
    cv::Point p2(axes.x, axes.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness);
}

void Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    cv::Point p1(pt1.x, pt1.y);
    cv::Point p2(pt2.x, pt2.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::line(*img, p1, p2, c, thickness);
}

void Rectangle(Mat img, Rect r, Scalar color, int thickness) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::rectangle(
        *img,
        cv::Point(r.x, r.y),
        cv::Point(r.x + r.width, r.y + r.height),
        c,
        thickness,
        cv::LINE_AA
    );
}

void FillPoly(Mat img, Contours points, Scalar color) {
    std::vector<std::vector<cv::Point> > pts;

    for (size_t i = 0; i < points.length; i++) {
        Contour contour = points.contours[i];

        std::vector<cv::Point> cntr;

        for (size_t i = 0; i < contour.length; i++) {
            cntr.push_back(cv::Point(contour.points[i].x, contour.points[i].y));
        }

        pts.push_back(cntr);
    }

    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::fillPoly(*img, pts, c);
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

Mat GetRotationMatrix2D(Point center, double angle, double scale) {
    cv::Point pt(center.x, center.y);
    return new  cv::Mat(cv::getRotationMatrix2D(pt, angle, scale));
}

void WarpAffine(Mat src, Mat dst, Mat m, Size dsize) {
    cv::Size sz(dsize.width, dsize.height);
    cv::warpAffine(*src, *dst, *m, sz);
}

void WarpAffineWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                          Scalar borderValue) {
    cv::Size sz(dsize.width, dsize.height);
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::warpAffine(*src, *dst, *rot_mat, sz, flags, borderMode, c);
}

void WarpPerspective(Mat src, Mat dst, Mat m, Size dsize) {
    cv::Size sz(dsize.width, dsize.height);
    cv::warpPerspective(*src, *dst, *m, sz);
}

void ApplyColorMap(Mat src, Mat dst, int colormap) {
    cv::applyColorMap(*src, *dst, colormap);
}

void ApplyCustomColorMap(Mat src, Mat dst, Mat colormap) {
    cv::applyColorMap(*src, *dst, *colormap);
}

Mat GetPerspectiveTransform(Contour src, Contour dst) {
    std::vector<cv::Point2f> src_pts;

    for (size_t i = 0; i < src.length; i++) {
        src_pts.push_back(cv::Point2f(src.points[i].x, src.points[i].y));
    }

    std::vector<cv::Point2f> dst_pts;

    for (size_t i = 0; i < dst.length; i++) {
        dst_pts.push_back(cv::Point2f(dst.points[i].x, dst.points[i].y));
    }

    return new cv::Mat(cv::getPerspectiveTransform(src_pts, dst_pts));
}

void DrawContours(Mat src, Contours contours, int contourIdx, Scalar color, int thickness) {
    std::vector<std::vector<cv::Point> > cntrs;

    for (size_t i = 0; i < contours.length; i++) {
        Contour contour = contours.contours[i];

        std::vector<cv::Point> cntr;

        for (size_t i = 0; i < contour.length; i++) {
            cntr.push_back(cv::Point(contour.points[i].x, contour.points[i].y));
        }

        cntrs.push_back(cntr);
    }

    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::drawContours(*src, cntrs, contourIdx, c, thickness);
}

void Sobel(Mat src, Mat dst, int ddepth, int dx, int dy, int ksize, double scale, double delta, int borderType) {
	cv::Sobel(*src, *dst, ddepth, dx, dy, ksize, scale, delta, borderType);
}

void SpatialGradient(Mat src, Mat dx, Mat dy, int ksize, int borderType) {
	cv::spatialGradient(*src, *dx, *dy, ksize, borderType);
}


void Remap(Mat src, Mat dst, Mat map1, Mat map2, int interpolation, int borderMode, Scalar borderValue) {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::remap(*src, *dst, *map1, *map2, interpolation, borderMode, c);
}

void Filter2D(Mat src, Mat dst, int ddepth, Mat kernel, Point anchor, double delta, int borderType) {
        cv::Point anchorPt(anchor.x, anchor.y);
        cv::filter2D(*src, *dst, ddepth, *kernel, anchorPt, delta, borderType);
}

void SepFilter2D(Mat src, Mat dst, int ddepth, Mat kernelX, Mat kernelY, Point anchor, double delta, int borderType) {
	cv::Point anchorPt(anchor.x, anchor.y);
	cv::sepFilter2D(*src, *dst, ddepth, *kernelX, *kernelY, anchorPt, delta, borderType);
}

void LogPolar(Mat src, Mat dst, Point center, double m, int flags) {
	cv::Point2f centerPt(center.x, center.y);
	cv::logPolar(*src, *dst, centerPt, m, flags);
}

void FitLine(Contour points, Mat line, int distType, double param, double reps, double aeps) {
	std::vector<cv::Point> pts;
	for (size_t i = 0; i < points.length; i++) {
		pts.push_back(cv::Point(points.points[i].x, points.points[i].y));
	}
	cv::fitLine(pts, *line, distType, param, reps, aeps);
}

CLAHE CLAHE_Create() {
    return new cv::Ptr<cv::CLAHE>(cv::createCLAHE());
}

CLAHE CLAHE_CreateWithParams(double clipLimit, Size tileGridSize) {
    cv::Size sz(tileGridSize.width, tileGridSize.height);
    return new cv::Ptr<cv::CLAHE>(cv::createCLAHE(clipLimit, sz));
}

void CLAHE_Close(CLAHE c) {
    delete c;
}

void CLAHE_Apply(CLAHE c, Mat src, Mat dst) {
    (*c)->apply(*src, *dst);
}
