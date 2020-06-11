#ifndef _OPENCV3_CALIB_H_
#define _OPENCV3_CALIB_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/calib3d.hpp>


extern "C" {
#endif

#include "core.h"

//Calib
void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d);
void Fisheye_UndistortImageWithParams(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew, Size size);
void Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat R, Mat P);
void Fisheye_EstimateNewCameraMatrixForUndistortRectify(Mat k, Mat d, Size imgSize, Mat r, Mat p, double balance, Size newSize, double fovScale);

void InitUndistortRectifyMap(Mat cameraMatrix,Mat distCoeffs,Mat r,Mat newCameraMatrix,Size size,int m1type,Mat map1,Mat map2);
Mat GetOptimalNewCameraMatrixWithParams(Mat cameraMatrix,Mat distCoeffs,Size size,double alpha,Size newImgSize,Rect* validPixROI,bool centerPrincipalPoint);
void Undistort(Mat src, Mat dst, Mat cameraMatrix, Mat distCoeffs, Mat newCameraMatrix);
void UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p);
bool FindChessboardCorners(Mat image, Size patternSize, Mat corners, int flags);
void DrawChessboardCorners(Mat image, Size patternSize, Mat corners, bool patternWasFound);
Mat EstimateAffinePartial2D(Contour2f from, Contour2f to);
#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CALIB_H