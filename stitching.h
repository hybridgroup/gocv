//
// Created by rainy on 2024/3/3.
//

#ifndef OPENCV_DART_LIBRARY_STITCHING_H
#define OPENCV_DART_LIBRARY_STITCHING_H

#include "core.h"
#include "exception.h"

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C"
{
#endif

    enum
    {
        STITCHING_PANORAMA = 0,
        STITCHING_SCANS = 1
    };
    enum
    {
        STITCHING_OK = 0,
        STITCHING_ERR_NEED_MORE_IMGS = 1,
        STITCHING_ERR_HOMOGRAPHY_EST_FAIL = 2,
        STITCHING_ERR_CAMERA_PARAMS_ADJUST_FAIL = 3
    };

#ifdef __cplusplus
    typedef cv::Stitcher *Stitcher;
#else
typedef void *Stitcher;
#endif

    CvStatus Stitcher_Create(int mode, Stitcher *rval);
    void Stitcher_Close(Stitcher stitcher);
    // CvStatus Stitcher_Get(Stitcher* stitcher);

#pragma region getter/setter
    CvStatus Stitcher_GetRegistrationResol(Stitcher stitcher, double *rval);
    CvStatus Stitcher_SetRegistrationResol(Stitcher stitcher, double inval);

    CvStatus Stitcher_GetSeamEstimationResol(Stitcher stitcher, double *rval);
    CvStatus Stitcher_SetSeamEstimationResol(Stitcher stitcher, double inval);

    CvStatus Stitcher_GetCompositingResol(Stitcher stitcher, double *rval);
    CvStatus Stitcher_SetCompositingResol(Stitcher stitcher, double inval);

    CvStatus Stitcher_GetPanoConfidenceThresh(Stitcher stitcher, double *rval);
    CvStatus Stitcher_SetPanoConfidenceThresh(Stitcher stitcher, double inval);

    CvStatus Stitcher_GetWaveCorrection(Stitcher stitcher, bool *rval);
    CvStatus Stitcher_SetWaveCorrection(Stitcher stitcher, bool inval);

    CvStatus Stitcher_GetInterpolationFlags(Stitcher stitcher, int *rval);
    CvStatus Stitcher_SetInterpolationFlags(Stitcher stitcher, int inval);

    CvStatus Stitcher_GetWaveCorrectKind(Stitcher stitcher, int *rval);
    CvStatus Stitcher_SetWaveCorrectKind(Stitcher stitcher, int inval);
#pragma endregion

#pragma region functions
    CvStatus Stitcher_EstimateTransform(Stitcher stitcher, Mats mats, Rects masks, int *rval);

    CvStatus Stitcher_ComposePanorama(Stitcher stitcher, Mat rpano, int *rval);
    CvStatus Stitcher_ComposePanorama_1(Stitcher stitcher, Mats mats, Mat rpano, int *rval);

    CvStatus Stitcher_Stitch(Stitcher stitcher, Mats mats, Mat rpano, int *rval);
    CvStatus Stitcher_Stitch_1(Stitcher stitcher, Mats mats, Rects masks, Mat rpano, int *rval);

    CvStatus Stitcher_Component(Stitcher stitcher, IntVector *rval);
#pragma endregion

#ifdef __cplusplus
}
#endif

#endif // OPENCV_DART_LIBRARY_STITCHING_H