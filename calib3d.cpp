#include "calib3d.h"

double Fisheye_Calibrate(Mat objectPoints, Mat imagePoints, Size size, Mat k, Mat d, Mat rvecs, Mat tvecs) {
    cv::Size sz(size.width, size.height);
    return cv::fisheye::calibrate(*objectPoints, *imagePoints, sz, *k, *d, *rvecs, *tvecs);
}

void Fisheye_UndistortPoints(Mat distorted, Mat undistorted, Mat k, Mat d) {
    cv::fisheye::undistortPoints(*distorted, *undistorted, *k, *d);
}

void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d);
}

void Fisheye_UndistortImageWithKNewMat(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d, *knew);
}

