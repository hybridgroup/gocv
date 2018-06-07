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
void Fisheye_UndistortImageWithKNewMat(Mat distorted, Mat undistorted, Mat k, Mat d, Mat knew);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CALIB_H
