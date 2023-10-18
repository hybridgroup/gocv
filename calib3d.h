#ifndef _OPENCV3_CALIB_H_
#define _OPENCV3_CALIB_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/calib3d.hpp>


extern "C" {
#endif

#include "core.h"

//Calib
double Fisheye_Calibrate(Points3fVector objectPoints, Points2fVector imagePoints, Size size, Mat k, Mat d, Mat rvecs, Mat tvecs, int flags);
void Fisheye_DistortPoints(Mat undistorted, Mat distorted, Mat k, Mat d);
void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d);
void Fisheye_UndistortImageWithParams(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew, Size size);
void Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat R, Mat P);
void Fisheye_EstimateNewCameraMatrixForUndistortRectify(Mat k, Mat d, Size imgSize, Mat r, Mat p, double balance, Size newSize, double fovScale);

void InitUndistortRectifyMap(Mat cameraMatrix,Mat distCoeffs,Mat r,Mat newCameraMatrix,Size size,int m1type,Mat map1,Mat map2);
Mat GetOptimalNewCameraMatrixWithParams(Mat cameraMatrix,Mat distCoeffs,Size size,double alpha,Size newImgSize,Rect* validPixROI,bool centerPrincipalPoint);
double CalibrateCamera(Points3fVector objectPoints, Points2fVector imagePoints, Size imageSize, Mat cameraMatrix, Mat distCoeffs, Mat rvecs, Mat tvecs, int flag);
void Undistort(Mat src, Mat dst, Mat cameraMatrix, Mat distCoeffs, Mat newCameraMatrix);
void UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p);
bool CheckChessboard(Mat image, Size sz);
bool FindChessboardCorners(Mat image, Size patternSize, Mat corners, int flags);
bool FindChessboardCornersSB(Mat image, Size patternSize, Mat corners, int flags);
bool FindChessboardCornersSBWithMeta(Mat image, Size patternSize, Mat corners, int flags, Mat meta);
void DrawChessboardCorners(Mat image, Size patternSize, Mat corners, bool patternWasFound);
Mat EstimateAffinePartial2D(Point2fVector from, Point2fVector to);
Mat EstimateAffinePartial2DWithParams(Point2fVector from, Point2fVector to, Mat inliers, int method, double ransacReprojThreshold, size_t maxIters, double confidence, size_t refineIters);
Mat EstimateAffine2D(Point2fVector from, Point2fVector to);
Mat EstimateAffine2DWithParams(Point2fVector from, Point2fVector to, Mat inliers, int method, double ransacReprojThreshold, size_t maxIters, double confidence, size_t refineIters);
void TriangulatePoints(Mat projMatr1, Mat projMatr2, Point2fVector projPoints1, Point2fVector projPoints2, Mat points4D);
void ConvertPointsFromHomogeneous(Mat src, Mat dst);
void Rodrigues(Mat src, Mat dst);
bool SolvePnP(Point3fVector objectPoints, Point2fVector imagePoints, Mat cameraMatrix, Mat distCoeffs, Mat rvec, Mat tvec, bool useExtrinsicGuess, int flags);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CALIB_H
