#ifndef _OPENCV3_IMGPROC_H_
#define _OPENCV3_IMGPROC_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#ifdef __cplusplus
typedef cv::Ptr<cv::CLAHE>* CLAHE;
#else
typedef void* CLAHE;
#endif

#include "core.h"

double ArcLength(PointVector curve, bool is_closed);
PointVector ApproxPolyDP(PointVector curve, double epsilon, bool closed);
OpenCVResult CvtColor(Mat src, Mat dst, int code);
OpenCVResult Demosaicing(Mat src, Mat dst, int code);
OpenCVResult EqualizeHist(Mat src, Mat dst);
OpenCVResult CalcHist(struct Mats mats, IntVector chans, Mat mask, Mat hist, IntVector sz, FloatVector rng, bool acc);
OpenCVResult CalcBackProject(struct Mats mats, IntVector chans, Mat hist, Mat backProject, FloatVector rng, bool uniform);
double CompareHist(Mat hist1, Mat hist2, int method);
float EMD(Mat sig1, Mat sig2, int distType);
OpenCVResult ConvexHull(PointVector points, Mat hull, bool clockwise, bool returnPoints);
OpenCVResult ConvexityDefects(PointVector points, Mat hull, Mat result);
OpenCVResult BilateralFilter(Mat src, Mat dst, int d, double sc, double ss);
OpenCVResult Blur(Mat src, Mat dst, Size ps);
OpenCVResult BoxFilter(Mat src, Mat dst, int ddepth, Size ps);
OpenCVResult SqBoxFilter(Mat src, Mat dst, int ddepth, Size ps);
OpenCVResult Dilate(Mat src, Mat dst, Mat kernel);
OpenCVResult DilateWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue);
OpenCVResult DistanceTransform(Mat src, Mat dst, Mat labels, int distanceType, int maskSize, int labelType);
OpenCVResult Erode(Mat src, Mat dst, Mat kernel);
OpenCVResult ErodeWithParams(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType);
OpenCVResult ErodeWithParamsAndBorderValue(Mat src, Mat dst, Mat kernel, Point anchor, int iterations, int borderType, Scalar borderValue);
OpenCVResult MatchTemplate(Mat image, Mat templ, Mat result, int method, Mat mask);
struct Moment Moments(Mat src, bool binaryImage);
OpenCVResult PyrDown(Mat src, Mat dst, Size dstsize, int borderType);
OpenCVResult PyrUp(Mat src, Mat dst, Size dstsize, int borderType);
struct Rect BoundingRect(PointVector pts);
OpenCVResult BoxPoints(RotatedRect rect, Mat boxPts);
OpenCVResult BoxPoints2f(RotatedRect2f rect, Mat boxPts);
double ContourArea(PointVector pts);
struct RotatedRect MinAreaRect(PointVector pts);
struct RotatedRect2f MinAreaRect2f(PointVector pts);
struct RotatedRect FitEllipse(PointVector pts);
OpenCVResult MinEnclosingCircle(PointVector pts, Point2f* center, float* radius);
PointsVector FindContours(Mat src, Mat hierarchy, int mode, int method);
double PointPolygonTest(PointVector pts, Point pt, bool measureDist);
int ConnectedComponents(Mat src, Mat dst, int connectivity, int ltype, int ccltype);
int ConnectedComponentsWithStats(Mat src, Mat labels, Mat stats, Mat centroids, int connectivity, int ltype, int ccltype);

OpenCVResult GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt);
Mat GetGaussianKernel(int ksize, double sigma, int ktype);
OpenCVResult Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta, int borderType);
OpenCVResult Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta,
            int borderType);
Mat GetStructuringElement(int shape, Size ksize);
Scalar MorphologyDefaultBorderValue();
OpenCVResult MorphologyEx(Mat src, Mat dst, int op, Mat kernel);
OpenCVResult MorphologyExWithParams(Mat src, Mat dst, int op, Mat kernel, Point pt, int iterations, int borderType);
OpenCVResult MedianBlur(Mat src, Mat dst, int ksize);

OpenCVResult Canny(Mat src, Mat edges, double t1, double t2);
OpenCVResult CornerSubPix(Mat img, Mat corners, Size winSize, Size zeroZone, TermCriteria criteria);
OpenCVResult GoodFeaturesToTrack(Mat img, Mat corners, int maxCorners, double quality, double minDist);
OpenCVResult GrabCut(Mat img, Mat mask, Rect rect, Mat bgdModel, Mat fgdModel, int iterCount, int mode);
OpenCVResult HoughCircles(Mat src, Mat circles, int method, double dp, double minDist);
OpenCVResult HoughCirclesWithParams(Mat src, Mat circles, int method, double dp, double minDist,
                            double param1, double param2, int minRadius, int maxRadius);
OpenCVResult HoughLines(Mat src, Mat lines, double rho, double theta, int threshold);
OpenCVResult HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold);
OpenCVResult HoughLinesPWithParams(Mat src, Mat lines, double rho, double theta, int threshold, double minLineLength, double maxLineGap);
OpenCVResult HoughLinesPointSet(Mat points, Mat lines, int lines_max, int threshold,
                        double min_rho, double  max_rho, double rho_step,
                        double min_theta, double max_theta, double theta_step);
OpenCVResult Integral(Mat src, Mat sum, Mat sqsum, Mat tilted);
double Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ);
OpenCVResult AdaptiveThreshold(Mat src, Mat dst, double maxValue, int adaptiveTyp, int typ, int blockSize,
                       double c);

OpenCVResult ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
OpenCVResult Circle(Mat img, Point center, int radius, Scalar color, int thickness);
OpenCVResult CircleWithParams(Mat img, Point center, int radius, Scalar color, int thickness, int lineType, int shift);
OpenCVResult Ellipse(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness);
OpenCVResult EllipseWithParams(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness, int lineType, int shift);
OpenCVResult Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
OpenCVResult Rectangle(Mat img, Rect rect, Scalar color, int thickness);
OpenCVResult RectangleWithParams(Mat img, Rect rect, Scalar color, int thickness, int lineType, int shift);
OpenCVResult FillPoly(Mat img, PointsVector points, Scalar color);
OpenCVResult FillPolyWithParams(Mat img, PointsVector points, Scalar color, int lineType, int shift, Point offset);
OpenCVResult Polylines(Mat img, PointsVector points, bool isClosed, Scalar color, int thickness);
struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness);
struct Size GetTextSizeWithBaseline(const char* text, int fontFace, double fontScale, int thickness, int* baseline);
OpenCVResult PutText(Mat img, const char* text, Point org, int fontFace, double fontScale,
            Scalar color, int thickness);
OpenCVResult PutTextWithParams(Mat img, const char* text, Point org, int fontFace, double fontScale,
            Scalar color, int thickness, int lineType, bool bottomLeftOrigin);
OpenCVResult Resize(Mat src, Mat dst, Size sz, double fx, double fy, int interp);
OpenCVResult GetRectSubPix(Mat src, Size patchSize, Point center, Mat dst);
Mat GetRotationMatrix2D(Point center, double angle, double scale);
OpenCVResult WarpAffine(Mat src, Mat dst, Mat rot_mat, Size dsize);
OpenCVResult WarpAffineWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode, Scalar borderValue);
OpenCVResult WarpPerspective(Mat src, Mat dst, Mat m, Size dsize);
OpenCVResult WarpPerspectiveWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                               Scalar borderValue);
OpenCVResult Watershed(Mat image, Mat markers);
OpenCVResult ApplyColorMap(Mat src, Mat dst, int colormap);
OpenCVResult ApplyCustomColorMap(Mat src, Mat dst, Mat colormap);
Mat GetPerspectiveTransform(PointVector src, PointVector dst);
Mat GetPerspectiveTransform2f(Point2fVector src, Point2fVector dst);
Mat GetAffineTransform(PointVector src, PointVector dst);
Mat GetAffineTransform2f(Point2fVector src, Point2fVector dst);
Mat FindHomography(Mat src, Mat dst, int method, double ransacReprojThreshold, Mat mask, const int maxIters, const double confidence) ;
OpenCVResult DrawContours(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness);
OpenCVResult DrawContoursWithParams(Mat src, PointsVector contours, int contourIdx, Scalar color, int thickness, int lineType, Mat hierarchy, int maxLevel, Point offset);
OpenCVResult Sobel(Mat src, Mat dst, int ddepth, int dx, int dy, int ksize, double scale, double delta, int borderType);
OpenCVResult SpatialGradient(Mat src, Mat dx, Mat dy, int ksize, int borderType);
OpenCVResult Remap(Mat src, Mat dst, Mat map1, Mat map2, int interpolation, int borderMode, Scalar borderValue);
OpenCVResult Filter2D(Mat src, Mat dst, int ddepth, Mat kernel, Point anchor, double delta, int borderType);
OpenCVResult SepFilter2D(Mat src, Mat dst, int ddepth, Mat kernelX, Mat kernelY, Point anchor, double delta, int borderType);
OpenCVResult LogPolar(Mat src, Mat dst, Point center, double m, int flags);
OpenCVResult FitLine(PointVector pts, Mat line, int distType, double param, double reps, double aeps);
OpenCVResult LinearPolar(Mat src, Mat dst, Point center, double maxRadius, int flags);
double MatchShapes(PointVector contour1, PointVector contour2, int method, double parameter);
bool ClipLine(Size imgSize, Point pt1, Point pt2);
CLAHE CLAHE_Create();
CLAHE CLAHE_CreateWithParams(double clipLimit, Size tileGridSize);
void CLAHE_Close(CLAHE c);
OpenCVResult CLAHE_Apply(CLAHE c, Mat src, Mat dst);
OpenCVResult InvertAffineTransform(Mat src, Mat dst);
Point2f PhaseCorrelate(Mat src1, Mat src2, Mat window, double* response);
OpenCVResult CreateHanningWindow(Mat dst, Size size, int typ);
OpenCVResult Mat_Accumulate(Mat src, Mat dst);
OpenCVResult Mat_AccumulateWithMask(Mat src, Mat dst, Mat mask);
OpenCVResult Mat_AccumulateSquare(Mat src, Mat dst);
OpenCVResult Mat_AccumulateSquareWithMask(Mat src, Mat dst, Mat mask);
OpenCVResult Mat_AccumulateProduct(Mat src1, Mat src2, Mat dst);
OpenCVResult Mat_AccumulateProductWithMask(Mat src1, Mat src2, Mat dst, Mat mask);
OpenCVResult Mat_AccumulatedWeighted(Mat src, Mat dst, double alpha);
OpenCVResult Mat_AccumulatedWeightedWithMask(Mat src, Mat dst, double alpha, Mat mask);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGPROC_H_
