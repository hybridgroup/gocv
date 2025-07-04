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

OpenCVResult CvtColor(Mat src, Mat dst, int code) {
    try {
        cv::cvtColor(*src, *dst, code);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Demosaicing(Mat src, Mat dst, int code) {
    try {
        cv::demosaicing(*src, *dst, code);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult EqualizeHist(Mat src, Mat dst) {
    try {
        cv::equalizeHist(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CalcHist(struct Mats mats, IntVector chans, Mat mask, Mat hist, IntVector sz, FloatVector rng, bool acc) {
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
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CalcBackProject(struct Mats mats, IntVector chans, Mat hist, Mat backProject, FloatVector rng, bool uniform){
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
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult ConvexHull(PointVector points, Mat hull, bool clockwise, bool returnPoints) {
    try {
        cv::convexHull(*points, *hull, clockwise, returnPoints);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ConvexityDefects(PointVector points, Mat hull, Mat result) {
    try {
        cv::convexityDefects(*points, *hull, *result);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult BilateralFilter(Mat src, Mat dst, int d, double sc, double ss) {
    try {
        cv::bilateralFilter(*src, *dst, d, sc, ss);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Blur(Mat src, Mat dst, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::blur(*src, *dst, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult BoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::boxFilter(*src, *dst, ddepth, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult SqBoxFilter(Mat src, Mat dst, int ddepth, Size ps) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::sqrBoxFilter(*src, *dst, ddepth, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Dilate(Mat src, Mat dst, Mat kernel) {
    try {
        cv::dilate(*src, *dst, *kernel);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult DilateWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue) {
    try {
        cv::Point pt1(anchor.x, anchor.y);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);

        cv::dilate(*src, *dst, *kernel, pt1, iterations, borderType, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult DistanceTransform(Mat src, Mat dst, Mat labels, int distanceType, int maskSize, int labelType) {
    try {
        cv::distanceTransform(*src, *dst, *labels, distanceType, maskSize, labelType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Erode(Mat src, Mat dst, Mat kernel) {
    try {
        cv::erode(*src, *dst, *kernel);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ErodeWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType) {
    try {
        cv::Point pt1(anchor.x, anchor.y);

        cv::erode(*src, *dst, *kernel, pt1, iterations, borderType, cv::morphologyDefaultBorderValue());
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ErodeWithParamsAndBorderValue(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue) {
    try {
        cv::Point pt1(anchor.x, anchor.y);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);

        cv::erode(*src, *dst, *kernel, pt1, iterations, borderType, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult MatchTemplate(Mat image, Mat templ, Mat result, int method, Mat mask) {
    try {
        cv::matchTemplate(*image, *templ, *result, method, *mask);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult PyrDown(Mat src, Mat dst, Size size, int borderType) {
    try {
        cv::Size cvSize(size.width, size.height);
        cv::pyrDown(*src, *dst, cvSize, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult PyrUp(Mat src, Mat dst, Size size, int borderType) {
    try {
        cv::Size cvSize(size.width, size.height);
        cv::pyrUp(*src, *dst, cvSize, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult BoxPoints(RotatedRect rect, Mat boxPts){
    try {
        cv::Point2f centerPt(rect.center.x , rect.center.y);
        cv::Size2f rSize(rect.size.width, rect.size.height);
        cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
        cv::boxPoints(rotatedRectangle, *boxPts);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult BoxPoints2f(RotatedRect2f rect, Mat boxPts){
    try {
        cv::Point2f centerPt(rect.center.x , rect.center.y);
        cv::Size2f rSize(rect.size.width, rect.size.height);
        cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
        cv::boxPoints(rotatedRectangle, *boxPts);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult MinEnclosingCircle(PointVector pts, Point2f* center, float* radius){
    try {
        cv::Point2f center2f;
        cv::minEnclosingCircle(*pts, center2f, *radius);
        center->x = center2f.x;
        center->y = center2f.y;
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult MorphologyEx(Mat src, Mat dst, int op, Mat kernel) {
    try {
        cv::morphologyEx(*src, *dst, op, *kernel);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult MorphologyExWithParams(Mat src, Mat dst, int op, Mat kernel, Point pt, int iterations, int borderType) {
    try {
        cv::Point pt1(pt.x, pt.y);
        cv::morphologyEx(*src, *dst, op, *kernel, pt1, iterations, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt) {
    try {
        cv::Size sz(ps.width, ps.height);
        cv::GaussianBlur(*src, *dst, sz, sX, sY, bt);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta, int borderType) {
    try {
        cv::Laplacian(*src, *dst, dDepth, kSize, scale, delta, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta, int borderType) {
    try {
        cv::Scharr(*src, *dst, dDepth, dx, dy, scale, delta, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult MedianBlur(Mat src, Mat dst, int ksize) {
    try {
        cv::medianBlur(*src, *dst, ksize);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Canny(Mat src, Mat edges, double t1, double t2) {
    try {
        cv::Canny(*src, *edges, t1, t2);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CornerSubPix(Mat img, Mat corners, Size winSize, Size zeroZone, TermCriteria criteria) {
    try {
        cv::Size wsz(winSize.width, winSize.height);
        cv::Size zsz(zeroZone.width, zeroZone.height);
        cv::cornerSubPix(*img, *corners, wsz, zsz, *criteria);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GoodFeaturesToTrack(Mat img, Mat corners, int maxCorners, double quality, double minDist) {
    try {
        cv::goodFeaturesToTrack(*img, *corners, maxCorners, quality, minDist);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GrabCut(Mat img, Mat mask, Rect r, Mat bgdModel, Mat fgdModel, int iterCount, int mode) {
    try {
        cv::Rect cvRect = cv::Rect(r.x, r.y, r.width, r.height);
        cv::grabCut(*img, *mask, cvRect, *bgdModel, *fgdModel, iterCount, mode);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughCircles(Mat src, Mat circles, int method, double dp, double minDist) {
    try {
        cv::HoughCircles(*src, *circles, method, dp, minDist);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughCirclesWithParams(Mat src, Mat circles, int method, double dp, double minDist,
                            double param1, double param2, int minRadius, int maxRadius) {
    try {
        cv::HoughCircles(*src, *circles, method, dp, minDist, param1, param2, minRadius, maxRadius);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughLines(Mat src, Mat lines, double rho, double theta, int threshold) {
    try {
        cv::HoughLines(*src, *lines, rho, theta, threshold);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold) {
    try {
        cv::HoughLinesP(*src, *lines, rho, theta, threshold);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughLinesPWithParams(Mat src, Mat lines, double rho, double theta, int threshold, double minLineLength, double maxLineGap) {
    try {
        cv::HoughLinesP(*src, *lines, rho, theta, threshold, minLineLength, maxLineGap);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult HoughLinesPointSet(Mat points, Mat lines, int linesMax, int threshold,
                        double minRho, double  maxRho, double rhoStep,
                        double minTheta, double maxTheta, double thetaStep) {
    try {
        cv::HoughLinesPointSet(*points, *lines, linesMax, threshold, minRho, maxRho, rhoStep, minTheta, maxTheta, thetaStep );
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Integral(Mat src, Mat sum, Mat sqsum, Mat tilted) {
    try {
        cv::integral(*src, *sum, *sqsum, *tilted);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult AdaptiveThreshold(Mat src, Mat dst, double maxValue, int adaptiveMethod, int thresholdType, int blockSize, double c) {
    try {
        cv::adaptiveThreshold(*src, *dst, maxValue, adaptiveMethod, thresholdType, blockSize, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    try {
        cv::Point p1(pt1.x, pt1.y);
        cv::Point p2(pt2.x, pt2.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::arrowedLine(*img, p1, p2, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult Circle(Mat img, Point center, int radius, Scalar color, int thickness) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::circle(*img, p1, radius, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult CircleWithParams(Mat img, Point center, int radius, Scalar color, int thickness, int lineType, int shift) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::circle(*img, p1, radius, c, thickness, lineType, shift);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Ellipse(Mat img, Point center, Point axes, double angle, double startAngle, double endAngle, Scalar color, int thickness) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Point p2(axes.x, axes.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult EllipseWithParams(Mat img, Point center, Point axes, double angle, double startAngle, double endAngle, Scalar color, int thickness, int lineType, int shift) {
    try {
        cv::Point p1(center.x, center.y);
        cv::Point p2(axes.x, axes.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness, lineType, shift);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness) {
    try {
        cv::Point p1(pt1.x, pt1.y);
        cv::Point p2(pt2.x, pt2.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::line(*img, p1, p2, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Rectangle(Mat img, Rect r, Scalar color, int thickness) {
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
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult RectangleWithParams(Mat img, Rect r, Scalar color, int thickness, int lineType, int shift) {
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
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult FillPoly(Mat img, PointsVector pts, Scalar color) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::fillPoly(*img, *pts, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult FillPolyWithParams(Mat img, PointsVector pts, Scalar color, int lineType, int shift, Point offset) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::fillPoly(*img, *pts, c, lineType, shift, cv::Point(offset.x, offset.y));
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Polylines(Mat img, PointsVector pts, bool isClosed, Scalar color,int thickness) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

        cv::polylines(*img, *pts, isClosed, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult PutText(Mat img, const char* text, Point org, int fontFace, double fontScale, Scalar color, int thickness) {
    try {
        cv::Point pt(org.x, org.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::putText(*img, text, pt, fontFace, fontScale, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult PutTextWithParams(Mat img, const char* text, Point org, int fontFace, double fontScale,
                       Scalar color, int thickness, int lineType, bool bottomLeftOrigin) {
    try {
        cv::Point pt(org.x, org.y);
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::putText(*img, text, pt, fontFace, fontScale, c, thickness, lineType, bottomLeftOrigin);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Resize(Mat src, Mat dst, Size dsize, double fx, double fy, int interp) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::resize(*src, *dst, sz, fx, fy, interp);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult GetRectSubPix(Mat src, Size patchSize, Point center, Mat dst) {
    try {
        cv::Size sz(patchSize.width, patchSize.height);
        cv::Point pt(center.x, center.y);
        cv::getRectSubPix(*src, sz, pt, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult WarpAffine(Mat src, Mat dst, Mat m, Size dsize) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::warpAffine(*src, *dst, *m, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult WarpAffineWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode, Scalar borderValue) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::warpAffine(*src, *dst, *rot_mat, sz, flags, borderMode, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult WarpPerspective(Mat src, Mat dst, Mat m, Size dsize) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::warpPerspective(*src, *dst, *m, sz);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult WarpPerspectiveWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode, Scalar borderValue) {
    try {
        cv::Size sz(dsize.width, dsize.height);
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::warpPerspective(*src, *dst, *rot_mat, sz, flags, borderMode, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Watershed(Mat image, Mat markers) {
    try {
        cv::watershed(*image, *markers);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ApplyColorMap(Mat src, Mat dst, int colormap) {
    try {
        cv::applyColorMap(*src, *dst, colormap);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult ApplyCustomColorMap(Mat src, Mat dst, Mat colormap) {
    try {
        cv::applyColorMap(*src, *dst, *colormap);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult DrawContours(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness) {
    try {
        cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
        cv::drawContours(*src, *contours, contourIdx, c, thickness);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult DrawContoursWithParams(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness, int lineType, Mat hierarchy, int maxLevel, Point offset) {
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
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Sobel(Mat src, Mat dst, int ddepth, int dx, int dy, int ksize, double scale, double delta, int borderType) {
    try {
        cv::Sobel(*src, *dst, ddepth, dx, dy, ksize, scale, delta, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult SpatialGradient(Mat src, Mat dx, Mat dy, int ksize, int borderType) {
    try {
        cv::spatialGradient(*src, *dx, *dy, ksize, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Remap(Mat src, Mat dst, Mat map1, Mat map2, int interpolation, int borderMode, Scalar borderValue) {
    try {
        cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::remap(*src, *dst, *map1, *map2, interpolation, borderMode, c);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Filter2D(Mat src, Mat dst, int ddepth, Mat kernel, Point anchor, double delta, int borderType) {
    try {
        cv::Point anchorPt(anchor.x, anchor.y);
        cv::filter2D(*src, *dst, ddepth, *kernel, anchorPt, delta, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult SepFilter2D(Mat src, Mat dst, int ddepth, Mat kernelX, Mat kernelY, Point anchor, double delta, int borderType) {
    try {
        cv::Point anchorPt(anchor.x, anchor.y);
        cv::sepFilter2D(*src, *dst, ddepth, *kernelX, *kernelY, anchorPt, delta, borderType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult LogPolar(Mat src, Mat dst, Point center, double m, int flags) {
    try {
        cv::Point2f centerPt(center.x, center.y);
        cv::logPolar(*src, *dst, centerPt, m, flags);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult FitLine(PointVector pts, Mat line, int distType, double param, double reps, double aeps) {
    try {
        cv::fitLine(*pts, *line, distType, param, reps, aeps);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult LinearPolar(Mat src, Mat dst, Point center, double maxRadius, int flags) {
    try {
        cv::Point2f centerPt(center.x, center.y);
        cv::linearPolar(*src, *dst, centerPt, maxRadius, flags);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult CLAHE_Apply(CLAHE c, Mat src, Mat dst) {
    try {
        (*c)->apply(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult InvertAffineTransform(Mat src, Mat dst) {
    try {
        cv::invertAffineTransform(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
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

OpenCVResult CreateHanningWindow(Mat dst, Size size, int typ) {
    try {
        cv::Size sz(size.width, size.height);
        cv::createHanningWindow(*dst, sz, typ);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_Accumulate(Mat src, Mat dst) {
    try {
        cv::accumulate(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulateWithMask(Mat src, Mat dst, Mat mask) {
    try {
        cv::accumulate(*src, *dst, *mask);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulateSquare(Mat src, Mat dst) {
    try {
        cv::accumulateSquare(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulateSquareWithMask(Mat src, Mat dst, Mat mask) {
    try {
        cv::accumulateSquare(*src, *dst, *mask);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulateProduct(Mat src1, Mat src2, Mat dst) {
    try {
        cv::accumulateProduct(*src1, *src2, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulateProductWithMask(Mat src1, Mat src2, Mat dst, Mat mask) {
    try {
        cv::accumulateProduct(*src1, *src2, *dst, *mask);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulatedWeighted(Mat src, Mat dst, double alpha) {
    try {
        cv::accumulateWeighted(*src, *dst, alpha);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Mat_AccumulatedWeightedWithMask(Mat src, Mat dst, double alpha, Mat mask) {
    try {
        cv::accumulateWeighted(*src, *dst, alpha, *mask);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
