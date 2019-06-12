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

double ArcLength(Contour curve, bool is_closed);
Contour ApproxPolyDP(Contour curve, double epsilon, bool closed);
void CvtColor(Mat src, Mat dst, int code);
void EqualizeHist(Mat src, Mat dst);
void CalcHist(struct Mats mats, IntVector chans, Mat mask, Mat hist, IntVector sz, FloatVector rng, bool acc);
void ConvexHull(Contour points, Mat hull, bool clockwise, bool returnPoints);
void ConvexityDefects(Contour points, Mat hull, Mat result);
void BilateralFilter(Mat src, Mat dst, int d, double sc, double ss);
void Blur(Mat src, Mat dst, Size ps);
void BoxFilter(Mat src, Mat dst, int ddepth, Size ps);
void SqBoxFilter(Mat src, Mat dst, int ddepth, Size ps);
void Dilate(Mat src, Mat dst, Mat kernel);
void Erode(Mat src, Mat dst, Mat kernel);
void MatchTemplate(Mat image, Mat templ, Mat result, int method, Mat mask);
struct Moment Moments(Mat src, bool binaryImage);
void PyrDown(Mat src, Mat dst, Size dstsize, int borderType);
void PyrUp(Mat src, Mat dst, Size dstsize, int borderType);
struct Rect BoundingRect(Contour con);
void BoxPoints(RotatedRect rect, Mat boxPts);
double ContourArea(Contour con);
struct RotatedRect MinAreaRect(Points points);
void MinEnclosingCircle(Points points, Point2f* center, float* radius);
struct Contours FindContours(Mat src, int mode, int method);
int ConnectedComponents(Mat src, Mat dst, int connectivity, int ltype, int ccltype);
int ConnectedComponentsWithStats(Mat src, Mat labels, Mat stats, Mat centroids, int connectivity, int ltype, int ccltype);

void GaussianBlur(Mat src, Mat dst, Size ps, double sX, double sY, int bt);
void Laplacian(Mat src, Mat dst, int dDepth, int kSize, double scale, double delta, int borderType);
void Scharr(Mat src, Mat dst, int dDepth, int dx, int dy, double scale, double delta,
            int borderType);
Mat GetStructuringElement(int shape, Size ksize);
Scalar MorphologyDefaultBorderValue();
void MorphologyEx(Mat src, Mat dst, int op, Mat kernel);
void MedianBlur(Mat src, Mat dst, int ksize);

void Canny(Mat src, Mat edges, double t1, double t2);
void CornerSubPix(Mat img, Mat corners, Size winSize, Size zeroZone, TermCriteria criteria);
void GoodFeaturesToTrack(Mat img, Mat corners, int maxCorners, double quality, double minDist);
void HoughCircles(Mat src, Mat circles, int method, double dp, double minDist);
void HoughCirclesWithParams(Mat src, Mat circles, int method, double dp, double minDist,
                            double param1, double param2, int minRadius, int maxRadius);
void HoughLines(Mat src, Mat lines, double rho, double theta, int threshold);
void HoughLinesP(Mat src, Mat lines, double rho, double theta, int threshold);
void HoughLinesPWithParams(Mat src, Mat lines, double rho, double theta, int threshold, double minLineLength, double maxLineGap);
void HoughLinesPointSet(Mat points, Mat lines, int lines_max, int threshold,
                        double min_rho, double  max_rho, double rho_step,
                        double min_theta, double max_theta, double theta_step);
void Threshold(Mat src, Mat dst, double thresh, double maxvalue, int typ);
void AdaptiveThreshold(Mat src, Mat dst, double maxValue, int adaptiveTyp, int typ, int blockSize,
                       double c);

void ArrowedLine(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
void Circle(Mat img, Point center, int radius, Scalar color, int thickness);
void Ellipse(Mat img, Point center, Point axes, double angle, double
             startAngle, double endAngle, Scalar color, int thickness);
void Line(Mat img, Point pt1, Point pt2, Scalar color, int thickness);
void Rectangle(Mat img, Rect rect, Scalar color, int thickness);
void FillPoly(Mat img, Contours points, Scalar color);
struct Size GetTextSize(const char* text, int fontFace, double fontScale, int thickness);
void PutText(Mat img, const char* text, Point org, int fontFace, double fontScale,
             Scalar color, int thickness);
void Resize(Mat src, Mat dst, Size sz, double fx, double fy, int interp);
Mat GetRotationMatrix2D(Point center, double angle, double scale);
void WarpAffine(Mat src, Mat dst, Mat rot_mat, Size dsize);
void WarpAffineWithParams(Mat src, Mat dst, Mat rot_mat, Size dsize, int flags, int borderMode,
                          Scalar borderValue);
void WarpPerspective(Mat src, Mat dst, Mat m, Size dsize);
void ApplyColorMap(Mat src, Mat dst, int colormap);
void ApplyCustomColorMap(Mat src, Mat dst, Mat colormap);
Mat GetPerspectiveTransform(Contour src, Contour dst);
void DrawContours(Mat src, Contours contours, int contourIdx, Scalar color, int thickness);
void Sobel(Mat src, Mat dst, int ddepth, int dx, int dy, int ksize, double scale, double delta, int borderType);
void SpatialGradient(Mat src, Mat dx, Mat dy, int ksize, int borderType);
void Remap(Mat src, Mat dst, Mat map1, Mat map2, int interpolation, int borderMode, Scalar borderValue);
void Filter2D(Mat src, Mat dst, int ddepth, Mat kernel, Point anchor, double delta, int borderType);
void SepFilter2D(Mat src, Mat dst, int ddepth, Mat kernelX, Mat kernelY, Point anchor, double delta, int borderType);
void LogPolar(Mat src, Mat dst, Point center, double m, int flags);
void FitLine(Contour points, Mat line, int distType, double param, double reps, double aeps);
CLAHE CLAHE_Create();
CLAHE CLAHE_CreateWithParams(double clipLimit, Size tileGridSize);
void CLAHE_Close(CLAHE c);
void CLAHE_Apply(CLAHE c, Mat src, Mat dst);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGPROC_H_
