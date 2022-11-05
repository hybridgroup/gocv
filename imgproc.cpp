#include "imgproc.h"

double ArcLength(PointVector curve, bool is_closed) {
    return cv::arcLength(*curve, is_closed);
}

PointVector ApproxPolyDP(PointVector curve, double epsilon, bool closed) {
    PointVector approxCurvePts = new std::vector<cv::Point>;
    cv::approxPolyDP(*curve, *approxCurvePts, epsilon, closed);

    return approxCurvePts;
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

        cv::calcBackProject(images, channels, *hist, *backProject, ranges, uniform);
}

double CompareHist(Mat hist1, Mat hist2, int method) {
    return cv::compareHist(*hist1, *hist2, method);
}

struct RotatedRect FitEllipse(PointVector pts)
{
    cv::RotatedRect bRect = cv::fitEllipse(*pts);

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
    cv::convexHull(*points, *hull, clockwise, returnPoints);
}

void ConvexityDefects(PointVector points, Mat hull, Mat result) {
    cv::convexityDefects(*points, *hull, *result);
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

void DilateWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue) {
    cv::Point pt1(anchor.x, anchor.y);
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);

    cv::dilate(*src, *dst, *kernel, pt1, iterations, borderType, c);
}

void DistanceTransform(Mat src, Mat dst, Mat labels, int distanceType, int maskSize, int labelType) {
    cv::distanceTransform(*src, *dst, *labels, distanceType, maskSize, labelType);
}

void Erode(Mat src, Mat dst, Mat kernel) {
    cv::erode(*src, *dst, *kernel);
}

void ErodeWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType) {
    cv::Point pt1(anchor.x, anchor.y);

    cv::erode(*src, *dst, *kernel, pt1, iterations, borderType, cv::morphologyDefaultBorderValue());
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

struct Rect BoundingRect(PointVector pts) {
    cv::Rect bRect = cv::boundingRect(*pts);
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    return r;
}

void BoxPoints(RotatedRect rect, Mat boxPts){
    cv::Point2f centerPt(rect.center.x , rect.center.y);
    cv::Size2f rSize(rect.size.width, rect.size.height);
    cv::RotatedRect rotatedRectangle(centerPt, rSize, rect.angle);
    cv::boxPoints(rotatedRectangle, *boxPts);
}

double ContourArea(PointVector pts) {
    return cv::contourArea(*pts);
}

struct RotatedRect MinAreaRect(PointVector pts){
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
}

void MinEnclosingCircle(PointVector pts, Point2f* center, float* radius){
    cv::Point2f center2f;
    cv::minEnclosingCircle(*pts, center2f, *radius);
    center->x = center2f.x;
    center->y = center2f.y;
}

PointsVector FindContours(Mat src, Mat hierarchy, int mode, int method) {
    PointsVector contours = new std::vector<std::vector<cv::Point> >;
    cv::findContours(*src, *contours, *hierarchy, mode, method);

    return contours;
}

double PointPolygonTest(PointVector pts, Point pt, bool measureDist) {
	cv::Point2f pt1(pt.x, pt.y);

  return cv::pointPolygonTest(*pts, pt1, measureDist);
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

void MorphologyExWithParams(Mat src, Mat dst, int op, Mat kernel, Point pt, int iterations, int borderType) {
    cv::Point pt1(pt.x, pt.y);
    cv::morphologyEx(*src, *dst, op, *kernel, pt1, iterations, borderType);
}

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt) {
    cv::Size sz(ps.width, ps.height);
    cv::GaussianBlur(*src, *dst, sz, sX, sY, bt);
}

Mat GetGaussianKernel(int ksize, double sigma, int ktype){
    return new cv::Mat(cv::getGaussianKernel(ksize, sigma, ktype));
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

void GrabCut(Mat img, Mat mask, Rect r, Mat bgdModel, Mat fgdModel, int iterCount, int mode) {
    cv::Rect cvRect = cv::Rect(r.x, r.y, r.width, r.height);
    cv::grabCut(*img, *mask, cvRect, *bgdModel, *fgdModel, iterCount, mode);
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

void Integral(Mat src, Mat sum, Mat sqsum, Mat tilted) {
    cv::integral(*src, *sum, *sqsum, *tilted);
}

double Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ) {
    return cv::threshold(*src, *dst, thresh, maxvalue, typ);
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

bool ClipLine(Size imgSize, Point pt1, Point pt2) {
	cv::Size sz(imgSize.width, imgSize.height);
	cv::Point p1(pt1.x, pt1.y);
	cv::Point p2(pt2.x, pt2.y);

	return	cv::clipLine(sz, p1, p2);
}

void Circle(Mat img, Point center, int radius, Scalar color, int thickness) {
    cv::Point p1(center.x, center.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::circle(*img, p1, radius, c, thickness);
}

void CircleWithParams(Mat img, Point center, int radius, Scalar color, int thickness, int lineType, int shift) {
    cv::Point p1(center.x, center.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::circle(*img, p1, radius, c, thickness, lineType, shift);
}

void Ellipse(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness) {
    cv::Point p1(center.x, center.y);
    cv::Point p2(axes.x, axes.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness);
}

void EllipseWithParams(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness, int lineType, int shift) {
    cv::Point p1(center.x, center.y);
    cv::Point p2(axes.x, axes.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::ellipse(*img, p1, p2, angle, startAngle, endAngle, c, thickness, lineType, shift);
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

void RectangleWithParams(Mat img, Rect r, Scalar color, int thickness, int lineType, int shift) {
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
}

void FillPoly(Mat img, PointsVector pts, Scalar color) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::fillPoly(*img, *pts, c);
}

void FillPolyWithParams(Mat img, PointsVector pts, Scalar color, int lineType, int shift, Point offset) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::fillPoly(*img, *pts, c, lineType, shift, cv::Point(offset.x, offset.y));
}

void Polylines(Mat img, PointsVector pts, bool isClosed, Scalar color,int thickness) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);

    cv::polylines(*img, *pts, isClosed, c, thickness);
}

struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness) {
    return GetTextSizeWithBaseline(text, fontFace, fontScale, thickness, NULL);
}

struct Size GetTextSizeWithBaseline(const char* text, int fontFace, double fontScale, int thickness, int* baesline) {
    cv::Size sz = cv::getTextSize(text, fontFace, fontScale, thickness, baesline);
    Size size = {sz.width, sz.height};
    return size;
}

void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale,
             Scalar color, int thickness) {
    cv::Point pt(org.x, org.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::putText(*img, text, pt, fontFace, fontScale, c, thickness);
}

void PutTextWithParams(Mat img, const char* text, Point org, int fontFace, double fontScale,
                       Scalar color, int thickness, int lineType, bool bottomLeftOrigin) {
    cv::Point pt(org.x, org.y);
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::putText(*img, text, pt, fontFace, fontScale, c, thickness, lineType, bottomLeftOrigin);
}

void Resize(Mat src, Mat dst, Size dsize, double fx, double fy, int interp) {
    cv::Size sz(dsize.width, dsize.height);
    cv::resize(*src, *dst, sz, fx, fy, interp);
}

void GetRectSubPix(Mat src, Size patchSize, Point center, Mat dst) {
    cv::Size sz(patchSize.width, patchSize.height);
    cv::Point pt(center.x, center.y);
    cv::getRectSubPix(*src, sz, pt, *dst);
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

void WarpPerspectiveWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                               Scalar borderValue) {
    cv::Size sz(dsize.width, dsize.height);
    cv::Scalar c = cv::Scalar(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
    cv::warpPerspective(*src, *dst, *rot_mat, sz, flags, borderMode, c);
}

void Watershed(Mat image, Mat markers) {
    cv::watershed(*image, *markers);
}

void ApplyColorMap(Mat src, Mat dst, int colormap) {
    cv::applyColorMap(*src, *dst, colormap);
}

void ApplyCustomColorMap(Mat src, Mat dst, Mat colormap) {
    cv::applyColorMap(*src, *dst, *colormap);
}

Mat GetPerspectiveTransform(PointVector src, PointVector dst) {
    std::vector<cv::Point2f> src_pts;
    copyPointVectorToPoint2fVector(src, &src_pts);

    std::vector<cv::Point2f> dst_pts;
    copyPointVectorToPoint2fVector(dst, &dst_pts);

    return new cv::Mat(cv::getPerspectiveTransform(src_pts, dst_pts));
}

Mat GetPerspectiveTransform2f(Point2fVector src, Point2fVector dst) {
    return new cv::Mat(cv::getPerspectiveTransform(*src, *dst));
}

Mat GetAffineTransform(PointVector src, PointVector dst) {
    std::vector<cv::Point2f> src_pts;
    copyPointVectorToPoint2fVector(src, &src_pts);

    std::vector<cv::Point2f> dst_pts;
    copyPointVectorToPoint2fVector(dst, &dst_pts);

    return new cv::Mat(cv::getAffineTransform(src_pts, dst_pts));
}

Mat GetAffineTransform2f(Point2fVector src, Point2fVector dst) {
    return new cv::Mat(cv::getAffineTransform(*src, *dst));
}

Mat FindHomography(Mat src, Mat dst, int method, double ransacReprojThreshold, Mat mask, const int maxIters, const double confidence) {
    return new cv::Mat(cv::findHomography(*src, *dst, method, ransacReprojThreshold, *mask, maxIters, confidence));
}

void DrawContours(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::drawContours(*src, *contours, contourIdx, c, thickness);
}

void DrawContoursWithParams(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness, int lineType, Mat hierarchy, int maxLevel, Point offset) {
    cv::Scalar c = cv::Scalar(color.val1, color.val2, color.val3, color.val4);
    cv::Point offsetPt(offset.x, offset.y);

    std::vector<cv::Vec4i> vecHierarchy;
    if (hierarchy->empty() == 0) {
        for (int j = 0; j < hierarchy->cols; ++j) {
            vecHierarchy.push_back(hierarchy->at<cv::Vec4i>(0, j));
        }
    }
    cv::drawContours(*src, *contours, contourIdx, c, thickness, lineType, vecHierarchy, maxLevel, offsetPt);
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

void FitLine(PointVector pts, Mat line, int distType, double param, double reps, double aeps) {
	cv::fitLine(*pts, *line, distType, param, reps, aeps);
}

void LinearPolar(Mat src, Mat dst, Point center, double maxRadius, int flags) {
	cv::Point2f centerPt(center.x, center.y);
	cv::linearPolar(*src, *dst, centerPt, maxRadius, flags);
}

double MatchShapes(PointVector contour1, PointVector contour2, int method, double parameter) {
    return cv::matchShapes(*contour1, *contour2, method, parameter);
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

void InvertAffineTransform(Mat src, Mat dst) {
	cv::invertAffineTransform(*src, *dst);
}

Point2f PhaseCorrelate(Mat src1, Mat src2, Mat window, double* response) {
    cv::Point2d result = cv::phaseCorrelate(*src1, *src2, *window, response);

    Point2f result2f = {
        .x = float(result.x),
        .y = float(result.y),
    };
    return result2f;
}

void Mat_Accumulate(Mat src, Mat dst) {
    cv::accumulate(*src, *dst);
}
void Mat_AccumulateWithMask(Mat src, Mat dst, Mat mask) {
    cv::accumulate(*src, *dst, *mask);
}

void Mat_AccumulateSquare(Mat src, Mat dst) {
    cv::accumulateSquare(*src, *dst);
}

void Mat_AccumulateSquareWithMask(Mat src, Mat dst, Mat mask) {
    cv::accumulateSquare(*src, *dst, *mask);
}

void Mat_AccumulateProduct(Mat src1, Mat src2, Mat dst) {
    cv::accumulateProduct(*src1, *src2, *dst);
}

void Mat_AccumulateProductWithMask(Mat src1, Mat src2, Mat dst, Mat mask) {
    cv::accumulateProduct(*src1, *src2, *dst, *mask);
}

void Mat_AccumulatedWeighted(Mat src, Mat dst, double alpha) {
    cv::accumulateWeighted(*src, *dst, alpha);
}

void Mat_AccumulatedWeightedWithMask(Mat src, Mat dst, double alpha, Mat mask) {
    cv::accumulateWeighted(*src, *dst, alpha, *mask);
}
