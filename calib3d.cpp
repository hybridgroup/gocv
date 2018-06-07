#include "calib3d.h"


void Fisheye_UndistortImage(Mat distorted, Mat undistorted, Mat k, Mat d) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d);
}

void Fisheye_UndistortImageWithKNewMat(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew) {
    cv::fisheye::undistortImage(*distorted, *undistorted, *k, *d, *knew);
}

