#include "imgproc.h"

double ArcLength(PointVector curve, bool is_closed) {
    try {
        return cv::arcLength(*curve, is_closed);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

PointVector ApproxPolyDP(PointVector curve, double epsilon, bool closed) {
    try {
        PointVector approxCurvePts = new std::vector<cv::Point>;
        cv::approxPolyDP(*curve, *approxCurvePts, epsilon, closed);

        return approxCurvePts;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CvtColor(Mat src, Mat dst, int code) {
    try {
        cv::cvtColor(*src, *dst, code);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void Demosaicing(Mat src, Mat dst, int code) {
    try {
        cv::demosaicing(*src, *dst, code);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void EqualizeHist(Mat src, Mat dst) {
    try {
        cv::equalizeHist(*src, *dst);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
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

    try {
        cv::calcHist(images, channels, *mask, *hist, histSize, ranges, acc);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CalcBackProject(struct Mats mats, IntVector chans, Mat hist, Mat backProject, FloatVector rng, bool uniform){
    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    std::vector<int> channels;
    for (int i = 0, *v = chans.val; i < chans.length; ++v, ++i) {
        channels.push_back(*v);
    }

    std::vector<float> ranges;

    float* f;
    int i;
    for (i = 0, f = rng.val; i < rng.length; ++f, ++i) {
        ranges.push_back(*f);
    }

    try {
        cv::calcBackProject(images, channels, *hist, *backProject, ranges, uniform);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

double CompareHist(Mat hist1, Mat hist2, int method) {
    try {
        return cv::compareHist(*hist1, *hist2, method);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

float EMD(Mat sig1, Mat sig2, int distType) {
    try {
        return cv::EMD(*sig1, *sig2, distType);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

struct RotatedRect FitEllipse(PointVector pts)
{
    cv::RotatedRect bRect;
    try {
        bRect = cv::fitEllipse(*pts);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());

        RotatedRect emptyRect;
        return emptyRect;
    }

    Rect r = {bRect.boundingRect().x, bRect.boundingRect().y, bRect.boundingRect().width, bRect.boundingRect().height};
    Point centrpt = {int(lroundf(bRect.center.x)), int(lroundf(bRect.center.y))};
    Size szsz = {int(lroundf(bRect.size.width)), int(lroundf(bRect.size.height))};

    cv::Point2f* pts4 = new cv::Point2f[4];
    bRect.points(pts4);
    Point* rpts = new Point[4];
    for (size_t j = 0; j < 4; j++) {
        Point pt = {int(lroundf(pts4[j].x)), int(lroundf(pts4[j].y))};
        rpts[j] = pt;
    }

    delete[] pts4;

    RotatedRect rotRect = {Points{rpts, 4}, r, centrpt, szsz, bRect.angle};
    return rotRect;
}

void ConvexHull(PointVector points, Mat hull, bool clockwise, bool returnPoints) {
    try {
        cv::convexHull(*points, *hull, clockwise, returnPoints);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void ConvexityDefects(PointVector points, Mat hull, Mat result) {
    try {
        cv::convexityDefects(*points, *hull, *result);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void BilateralFilter(Mat src, Mat dst, int d, double sc, double ss) {
    try {
        cv::bilateralFilter(*src, *dst, d, sc, ss);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void Blur(Mat src, Mat dst, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::blur(*src, *dst, sz);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void BoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::boxFilter(*src, *dst, ddepth, sz);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void SqBoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::sqrBoxFilter(*src, *dst, ddepth, sz);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void Dilate(Mat src, Mat dst, Mat kernel) {
    try {
        cv::dilate(*src, *dst, *kernel);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void DilateWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue) {
    try {
        cv::Point pt1(anchor.x, anchor.y);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);

        cv::dilate(*src, *dst, *kernel, pt1, iterations, borderType, c);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void DistanceTransform(Mat src, Mat dst, Mat labels, int distanceType, int maskSize, int labelType) {
    try {
        cv::distanceTransform(*src, *dst, *labels, distanceType, maskSize, labelType);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void Erode(Mat src, Mat dst, Mat kernel) {
    try {
        cv::erode(*src, *dst, *kernel);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void ErodeWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType) {
    try {
        cv::Point pt1(anchor.x, anchor.y);

        cv::erode(*src, *dst, *kernel, pt1, iterations, borderType, cv::morphologyDefaultBorderValue());
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void ErodeWithParamsAndBorderValue(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue) {
    try {
        cv::Point pt1(anchor.x, anchor.y);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);

        cv::erode(*src, *dst, *kernel, pt1, iterations, borderType, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void MatchTemplate(Mat image, Mat templ, Mat result, int method, Mat mask) {
    try {
        cv::matchTemplate(*image, *templ, *result, method, *mask);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

struct Moment Moments(Mat src, bool binaryImage) {
    try {
        cv::Moments m = cv::moments(*src, binaryImage);
        Moment mom = {m.m00, m.m10, m.m01, m.m20, m.m11, m.m02, m.m30, m.m21, m.m12, m.m03,
                      m.mu20, m.mu11, m.mu02, m.mu30, m.mu21, m.mu12, m.mu03,
                      m.nu20, m.nu11, m.nu02, m.nu30, m.nu21, m.nu12, m.nu03
                     };
        return mom;
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
        Moment mom;
        return mom;
    }
}

void PyrDown(Mat src, Mat dst, Size size, int borderType) {
    try {
        cv::Size cvSize(size.width, size.height);
        cv::pyrDown(*src, *dst, cvSize, borderType);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

void PyrUp(Mat src, Mat dst, Size size, int borderType) {
    try {
        cv::Size cvSize(size.width, size.height);
        cv::pyrUp(*src, *dst, cvSize, borderType);
    } catch(const cv::Exception& e) {
        setExceptionInfo(e.code, e.what());
    }
}

struct Rect BoundingRect(PointVector pts) {
    cv::Rect bRect;
    try {
        bRect = cv::boundingRect(*pts);
        Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
        return r;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rect r;
        return r;
    }
}

void BoxPoints(RotatedRect rect, Mat boxPts){
    try {
        cv::Point2f centerPt(rect.center.x , rect.center.y);
        cv::Size2f rSize(rect.size.width, rect.size.height);
        cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
        cv::boxPoints(rotatedRectangle, *boxPts);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void BoxPoints2f(RotatedRect2f rect, Mat boxPts){
    try {
        cv::Point2f centerPt(rect.center.x , rect.center.y);
        cv::Size2f rSize(rect.size.width, rect.size.height);
        cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
        cv::boxPoints(rotatedRectangle, *boxPts);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double ContourArea(PointVector pts) {
    try {
        return cv::contourArea(*pts);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

struct RotatedRect MinAreaRect(PointVector pts){
    try {
        cv::RotatedRect cvrect = cv::minAreaRect(*pts);

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
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        RotatedRect retrect;
        return retrect;
    }
}

struct RotatedRect2f MinAreaRect2f(PointVector pts){
    try {
        cv::RotatedRect cvrect = cv::minAreaRect(*pts);

        Point2f* rpts = new Point2f[4];
        cv::Point2f* pts4 = new cv::Point2f[4];
        cvrect.points(pts4);

        for (size_t j = 0; j < 4; j++) {
            Point2f pt = {pts4[j].x, pts4[j].y};
            rpts[j] = pt;
        }

        delete[] pts4;

        cv::Rect bRect = cvrect.boundingRect();
        Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
        Point2f centrpt = {cvrect.center.x, cvrect.center.y};
        Size2f szsz = {cvrect.size.width, cvrect.size.height};

        RotatedRect2f retrect = {(Contour2f){rpts, 4}, r, centrpt, szsz, cvrect.angle};
        return retrect;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        RotatedRect2f retrect;
        return retrect;
    }
}

void MinEnclosingCircle(PointVector pts, Point2f* center, float* radius){
    try {
        cv::Point2f center2f;
        cv::minEnclosingCircle(*pts, center2f, *radius);
        center->x = center2f.x;
        center->y = center2f.y;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

PointsVector FindContours(Mat src, Mat hierarchy, int mode, int method) {
    try {
        PointsVector contours = new std::vector<std::vector<cv::Point> >;
        cv::findContours(*src, *contours, *hierarchy, mode, method);

        return contours;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

double PointPolygonTest(PointVector pts, Point pt, bool measureDist) {
    try {
        cv::Point2f pt1(pt.x, pt.y);

        return cv::pointPolygonTest(*pts, pt1, measureDist);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

int ConnectedComponents(Mat src, Mat labels, int connectivity, int ltype, int ccltype){
    try {
        return cv::connectedComponents(*src, *labels, connectivity, ltype, ccltype);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}


int ConnectedComponentsWithStats(Mat src, Mat labels, Mat stats, Mat centroids,
    int connectivity, int ltype, int ccltype) {
    try {
        return cv::connectedComponentsWithStats(*src, *labels, *stats, *centroids, connectivity, ltype, ccltype);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

Mat GetStructuringElement(int shape, Size ksize) {
    try {
        cv::Size sz(ksize.width, ksize.height);
        return new cv::Mat(cv::getStructuringElement(shape, sz));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Scalar MorphologyDefaultBorderValue(){
    try {
        cv::Scalar cs = cv::morphologyDefaultBorderValue();
        return (Scalar){cs[0],cs[1],cs[2],cs[3]};
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Scalar scal = Scalar();
        return scal;
    }
}

void MorphologyEx(Mat src, Mat dst, int op, Mat kernel) {
    try {
        cv::morphologyEx(*src, *dst, op, *kernel);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void MorphologyExWithParams(Mat src, Mat dst, int op, Mat kernel, Point pt, int iterations, int borderType) {
    try {
        cv::Point pt1(pt.x, pt.y);
        cv::morphologyEx(*src, *dst, op, *kernel, pt1, iterations, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::GaussianBlur(*src, *dst, sz, sX, sY, bt);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Mat GetGaussianKernel(int ksize, double sigma, int ktype){
    try {
        return new cv::Mat(cv::getGaussianKernel(ksize, sigma, ktype));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta, int borderType) {
    try {
        cv::Laplacian(*src, *dst, dDepth, kSize, scale, delta, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta, int borderType) {
    try {
        cv::Scharr(*src, *dst, dDepth, dx, dy, scale, delta, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void MedianBlur(Mat src, Mat dst, int ksize) {
    try {
        cv::medianBlur(*src, *dst, ksize);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Canny(Mat src, Mat edges, double t1, double t2) {
    try {
        cv::Canny(*src, *edges, t1, t2);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CornerSubPix(Mat img, Mat corners, Size winSize, Size zeroZone, TermCriteria criteria) {
    try {
        cv::Size wsz(winSize.width, winSize.height);
        cv::Size zsz(zeroZone.width, zeroZone.height);
        cv::cornerSubPix(*img, *corners, wsz, zsz, *criteria);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GoodFeaturesToTrack(Mat img, Mat corners, int maxCorners, double quality, double minDist) {
    try {
        cv::goodFeaturesToTrack(*img, *corners, maxCorners, quality, minDist);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GrabCut(Mat img, Mat mask, Rect r, Mat bgdModel, Mat fgdModel, int iterCount, int mode) {
    try {
        cv::Rect cvRect = cv::Rect(r.x, r.y, r.width, r.height);
        cv::grabCut(*img, *mask, cvRect, *bgdModel, *fgdModel, iterCount, mode);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughCircles(Mat src, Mat circles, int method, double dp, double minDist) {
    try {
        cv::HoughCircles(*src, *circles, method, dp, minDist);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughCirclesWithParams(Mat src, Mat circles, int method, double dp, double minDist,
                            double param1, double param2, int minRadius, int maxRadius) {
    try {
        cv::HoughCircles(*src, *circles, method, dp, minDist, param1, param2, minRadius, maxRadius);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughLines(Mat src, Mat lines, double rho, double theta, int threshold) {
    try {
        cv::HoughLines(*src, *lines, rho, theta, threshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold) {
    try {
        cv::HoughLinesP(*src, *lines, rho, theta, threshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughLinesPWithParams(Mat src, Mat lines, double rho, double theta, int threshold, double minLineLength, double maxLineGap) {
    try {
        cv::HoughLinesP(*src, *lines, rho, theta, threshold, minLineLength, maxLineGap);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void HoughLinesPointSet(Mat points, Mat lines, int linesMax, int threshold,
                        double minRho, double  maxRho, double rhoStep,
                        double minTheta, double maxTheta, double thetaStep) {
    try {
        cv::HoughLinesPointSet(*points, *lines, linesMax, threshold,
            minRho, maxRho, rhoStep, minTheta, maxTheta, thetaStep );
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Integral(Mat src, Mat sum, Mat sqsum, Mat tilted) {
    try {
        cv::integral(*src, *sum, *sqsum, *tilted);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ) {
    try {
        return cv::threshold(*src, *dst, thresh, maxvalue, typ);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void AdaptiveThreshold(Mat src, Mat dst, double maxValue, int adaptiveMethod, int thresholdType, int blockSize, double c) {
    try {
        cv::adaptiveThreshold(*src, *dst, maxValue, adaptiveMethod, thresholdType, blockSize, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    try {
        cv::Point p1(pt1.x, pt1.y);
        cv::Point p2(pt2.x, pt2.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::arrowedLine(*img, p1, p2, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

bool ClipLine(Size imgSize, Point pt1, Point pt2) {
    try {
        cv::Size sz(imgSize.width, imgSize.height);
        cv::Point p1(pt1.x, pt1.y);
        cv::Point p2(pt2.x, pt2.y);

        return cv::clipLine(sz, p1, p2);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

void Circle(Mat img, Point center, int radius, Scalar color, int thickness) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::circle(*img, p1, radius, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void CircleWithParams(Mat img, Point center, int radius, Scalar color, int thickness, int lineType, int shift) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::circle(*img, p1, radius, c, thickness, lineType, shift);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Ellipse(Mat img, Point center, Point axes, double angle, double startAngle, double endAngle, Scalar color, int thickness) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Point p2(axes.x, axes.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void EllipseWithParams(Mat img, Point center, Point axes, double angle, double startAngle, double endAngle, Scalar color, int thickness, int lineType, int shift) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Point p2(axes.x, axes.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness, lineType, shift);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    try {
        cv::Point p1(pt1.x, pt1.y);
        cv::Point p2(pt2.x, pt2.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::line(*img, p1, p2, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Rectangle(Mat img, Rect r, Scalar color, int thickness) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::rectangle(
            *img,
            cv::Point(r.x, r.y),
            cv::Point(r.x + r.width, r.y + r.height),
            c,
            thickness,
            cv::LINE_AA
        );
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void RectangleWithParams(Mat img, Rect r, Scalar color, int thickness, int lineType, int shift) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::rectangle(
            *img,
            cv::Point(r.x, r.y),
            cv::Point(r.x + r.width, r.y + r.height),
            c,
            thickness,
            lineType,
            shift
        );
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FillPoly(Mat img, PointsVector pts, Scalar color) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::fillPoly(*img, *pts, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FillPolyWithParams(Mat img, PointsVector pts, Scalar color, int lineType, int shift, Point offset) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::fillPoly(*img, *pts, c, lineType, shift, cv::Point(offset.x, offset.y));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Polylines(Mat img, PointsVector pts, bool isClosed, Scalar color,int thickness) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::polylines(*img, *pts, isClosed, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness) {
    try {
        return GetTextSizeWithBaseline(text, fontFace, fontScale, thickness, NULL);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Size size = {0, 0};
        return size;
    }
}

struct Size GetTextSizeWithBaseline(const char* text, int fontFace, double fontScale, int thickness, int* baesline) {
    try {
        cv::Size sz = cv::getTextSize(text, fontFace, fontScale, thickness, baesline);
        Size size = {sz.width, sz.height};
        return size;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Size size = {0, 0};
        return size;
    }
}

void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale, Scalar color, int thickness) {
    try {
        cv::Point pt(org.x, org.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::putText(*img, text, pt, fontFace, fontScale, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void PutTextWithParams(Mat img, const char* text, Point org, int fontFace, double fontScale,
                       Scalar color, int thickness, int lineType, bool bottomLeftOrigin) {
    try {
        cv::Point pt(org.x, org.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::putText(*img, text, pt, fontFace, fontScale, c, thickness, lineType, bottomLeftOrigin);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Resize(Mat src, Mat dst, Size dsize, double fx, double fy, int interp) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::resize(*src, *dst, sz, fx, fy, interp);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void GetRectSubPix(Mat src, Size patchSize, Point center, Mat dst) {
    try {
        cv::Size sz(patchSize.width, patchSize.height);
        cv::Point pt(center.x, center.y);
        cv::getRectSubPix(*src, sz, pt, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Mat GetRotationMatrix2D(Point center, double angle, double scale) {
    try {
        cv::Point pt(center.x, center.y);
        return new cv::Mat(cv::getRotationMatrix2D(pt, angle, scale));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void WarpAffine(Mat src, Mat dst, Mat m, Size dsize) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::warpAffine(*src, *dst, *m, sz);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void WarpAffineWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                          Scalar borderValue) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::warpAffine(*src, *dst, *rot_mat, sz, flags, borderMode, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void WarpPerspective(Mat src, Mat dst, Mat m, Size dsize) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::warpPerspective(*src, *dst, *m, sz);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void WarpPerspectiveWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                               Scalar borderValue) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::warpPerspective(*src, *dst, *rot_mat, sz, flags, borderMode, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Watershed(Mat image, Mat markers) {
    try {
        cv::watershed(*image, *markers);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void ApplyColorMap(Mat src, Mat dst, int colormap) {
    try {
        cv::applyColorMap(*src, *dst, colormap);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void ApplyCustomColorMap(Mat src, Mat dst, Mat colormap) {
    try {
        cv::applyColorMap(*src, *dst, *colormap);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Mat GetPerspectiveTransform(PointVector src, PointVector dst) {
    try {
        std::vector<cv::Point2f> src_pts;
        copyPointVectorToPoint2fVector(src, &src_pts);

        std::vector<cv::Point2f> dst_pts;
        copyPointVectorToPoint2fVector(dst, &dst_pts);

        return new cv::Mat(cv::getPerspectiveTransform(src_pts, dst_pts));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat GetPerspectiveTransform2f(Point2fVector src, Point2fVector dst) {
    try {
        return new cv::Mat(cv::getPerspectiveTransform(*src, *dst));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat GetAffineTransform(PointVector src, PointVector dst) {
    try {
        std::vector<cv::Point2f> src_pts;
        copyPointVectorToPoint2fVector(src, &src_pts);

        std::vector<cv::Point2f> dst_pts;
        copyPointVectorToPoint2fVector(dst, &dst_pts);

        return new cv::Mat(cv::getAffineTransform(src_pts, dst_pts));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat GetAffineTransform2f(Point2fVector src, Point2fVector dst) {
    try {
        return new cv::Mat(cv::getAffineTransform(*src, *dst));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat FindHomography(Mat src, Mat dst, int method, double ransacReprojThreshold, Mat mask, const int maxIters, const double confidence) {
    try {
        return new cv::Mat(cv::findHomography(*src, *dst, method, ransacReprojThreshold, *mask, maxIters, confidence));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void DrawContours(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::drawContours(*src, *contours, contourIdx, c, thickness);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void DrawContoursWithParams(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness, int lineType, Mat hierarchy, int maxLevel, Point offset) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::Point offsetPt(offset.x, offset.y);

        std::vector<cv::Vec4i> vecHierarchy;
        if (hierarchy->empty() == 0) {
            for (int j = 0; j < hierarchy->cols; ++j) {
                vecHierarchy.push_back(hierarchy->at<cv::Vec4i>(0, j));
            }
        }
        cv::drawContours(*src, *contours, contourIdx, c, thickness, lineType, vecHierarchy, maxLevel, offsetPt);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Sobel(Mat src, Mat dst, int ddepth, int dx, int dy, int ksize, double scale, double delta, int borderType) {
    try {
        cv::Sobel(*src, *dst, ddepth, dx, dy, ksize, scale, delta, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void SpatialGradient(Mat src, Mat dx, Mat dy, int ksize, int borderType) {
    try {
        cv::spatialGradient(*src, *dx, *dy, ksize, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}


void Remap(Mat src, Mat dst, Mat map1, Mat map2, int interpolation, int borderMode, Scalar borderValue) {
    try {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::remap(*src, *dst, *map1, *map2, interpolation, borderMode, c);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Filter2D(Mat src, Mat dst, int ddepth, Mat kernel, Point anchor, double delta, int borderType) {
    try {
        cv::Point anchorPt(anchor.x, anchor.y);
        cv::filter2D(*src, *dst, ddepth, *kernel, anchorPt, delta, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void SepFilter2D(Mat src, Mat dst, int ddepth, Mat kernelX, Mat kernelY, Point anchor, double delta, int borderType) {
    try {
        cv::Point anchorPt(anchor.x, anchor.y);
        cv::sepFilter2D(*src, *dst, ddepth, *kernelX, *kernelY, anchorPt, delta, borderType);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LogPolar(Mat src, Mat dst, Point center, double m, int flags) {
    try {
        cv::Point2f centerPt(center.x, center.y);
        cv::logPolar(*src, *dst, centerPt, m, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FitLine(PointVector pts, Mat line, int distType, double param, double reps, double aeps) {
    try {
        cv::fitLine(*pts, *line, distType, param, reps, aeps);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LinearPolar(Mat src, Mat dst, Point center, double maxRadius, int flags) {
    try {
        cv::Point2f centerPt(center.x, center.y);
        cv::linearPolar(*src, *dst, centerPt, maxRadius, flags);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

double MatchShapes(PointVector contour1, PointVector contour2, int method, double parameter) {
    try {
        return cv::matchShapes(*contour1, *contour2, method, parameter);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

CLAHE CLAHE_Create() {
    try {
        return new cv::Ptr<cv::CLAHE>(cv::createCLAHE());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

CLAHE CLAHE_CreateWithParams(double clipLimit, Size tileGridSize) {
    try {
        cv::Size sz(tileGridSize.width, tileGridSize.height);
        return new cv::Ptr<cv::CLAHE>(cv::createCLAHE(clipLimit, sz));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void CLAHE_Close(CLAHE c) {
    delete c;
}

void CLAHE_Apply(CLAHE c, Mat src, Mat dst) {
    try {
        (*c)->apply(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void InvertAffineTransform(Mat src, Mat dst) {
    try {
        cv::invertAffineTransform(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Point2f PhaseCorrelate(Mat src1, Mat src2, Mat window, double* response) {
    try {
        cv::Point2d result = cv::phaseCorrelate(*src1, *src2, *window, response);

        Point2f result2f = {
            .x = float(result.x),
            .y = float(result.y),
        };
        return result2f;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Point2f result2f = {
            .x = 0.0,
            .y = 0.0,
        };
        return result2f;
    }
}

void CreateHanningWindow(Mat dst, Size size, int typ) {
    try {
        cv::Size sz(size.width, size.height);
        cv::createHanningWindow(*dst, sz, typ);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_Accumulate(Mat src, Mat dst) {
    try {
        cv::accumulate(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
void Mat_AccumulateWithMask(Mat src, Mat dst, Mat mask) {
    try {
        cv::accumulate(*src, *dst, *mask);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulateSquare(Mat src, Mat dst) {
    try {
        cv::accumulateSquare(*src, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulateSquareWithMask(Mat src, Mat dst, Mat mask) {
    try {
        cv::accumulateSquare(*src, *dst, *mask);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulateProduct(Mat src1, Mat src2, Mat dst) {
    try {
        cv::accumulateProduct(*src1, *src2, *dst);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulateProductWithMask(Mat src1, Mat src2, Mat dst, Mat mask) {
    try {
        cv::accumulateProduct(*src1, *src2, *dst, *mask);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulatedWeighted(Mat src, Mat dst, double alpha) {
    try {
        cv::accumulateWeighted(*src, *dst, alpha);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Mat_AccumulatedWeightedWithMask(Mat src, Mat dst, double alpha, Mat mask) {
    try {
        cv::accumulateWeighted(*src, *dst, alpha, *mask);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
