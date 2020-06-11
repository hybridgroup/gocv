#include "calib3d.h"


void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d);
}

void Fisheye_UndistortImageWithParams(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew, Size size) {
    cv::Size sz(size.width, size.height);
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d, *knew, sz);
}

void Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    cv::fisheye::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
}

void Fisheye_EstimateNewCameraMatrixForUndistortRectify(Mat k, Mat d, Size imgSize, Mat r, Mat p, double balance, Size newSize, double fovScale) {
    cv::Size newSz(newSize.width, newSize.height);
    cv::Size imgSz(imgSize.width, imgSize.height);
    cv::fisheye::estimateNewCameraMatrixForUndistortRectify(*k, *d, imgSz, *r, *p, balance, newSz, fovScale);
}

void InitUndistortRectifyMap(Mat cameraMatrix,Mat distCoeffs,Mat r,Mat newCameraMatrix,Size size,int m1type,Mat map1,Mat map2) {
    cv::Size sz(size.width, size.height);
    cv::initUndistortRectifyMap(*cameraMatrix,*distCoeffs,*r,*newCameraMatrix,sz,m1type,*map1,*map2);
}

Mat GetOptimalNewCameraMatrixWithParams(Mat cameraMatrix,Mat distCoeffs,Size size,double alpha,Size newImgSize,Rect* validPixROI,bool centerPrincipalPoint) {
    cv::Size sz(size.width, size.height);
    cv::Size newSize(newImgSize.width, newImgSize.height);
    cv::Rect rect(validPixROI->x,validPixROI->y,validPixROI->width,validPixROI->height);
    cv::Mat* mat = new cv::Mat(cv::getOptimalNewCameraMatrix(*cameraMatrix,*distCoeffs,sz,alpha,newSize,&rect,centerPrincipalPoint));
    validPixROI->x = rect.x;
    validPixROI->y = rect.y;
    validPixROI->width = rect.width;
    validPixROI->height = rect.height;
    return mat;
}

void Undistort(Mat src, Mat dst, Mat cameraMatrix, Mat distCoeffs, Mat newCameraMatrix) {
    cv::undistort(*src, *dst, *cameraMatrix, *distCoeffs, *newCameraMatrix);
}

void UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d, Mat r, Mat p) {
    cv::undistortPoints(*distorted, *undistorted, *k, *d, *r, *p);
}

bool FindChessboardCorners(Mat image, Size patternSize, Mat corners, int flags) {
    cv::Size sz(patternSize.width, patternSize.height);
    return cv::findChessboardCorners(*image, sz, *corners, flags);
}

void DrawChessboardCorners(Mat image, Size patternSize, Mat corners, bool patternWasFound) {
    cv::Size sz(patternSize.width, patternSize.height);
    cv::drawChessboardCorners(*image, sz, *corners, patternWasFound);
}

Mat EstimateAffinePartial2D(Contour2f from, Contour2f to) {
  std::vector<cv::Point2f> from_pts;
  for (size_t i = 0; i < from.length; i++) {
    from_pts.push_back(cv::Point2f(from.points[i].x, from.points[i].y));
  }
  std::vector<cv::Point2f> to_pts;
  for (size_t i = 0; i < to.length; i++) {
    to_pts.push_back(cv::Point2f(to.points[i].x, to.points[i].y));
  }

  return new cv::Mat(cv::estimateAffinePartial2D(from_pts, to_pts));
}
