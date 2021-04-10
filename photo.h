#ifndef _OPENCV3_PHOTO_H_
#define _OPENCV3_PHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/photo.hpp>

extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
// see : https://docs.opencv.org/3.4/d7/dd6/classcv_1_1MergeMertens.html
typedef cv::Ptr<cv::MergeMertens> *MergeMertens;
// see : https://docs.opencv.org/master/d7/db6/classcv_1_1AlignMTB.html
typedef cv::Ptr<cv::AlignMTB> *AlignMTB;
#else
typedef void *MergeMertens;
typedef void *AlignMTB;
#endif

void ColorChange(Mat src, Mat mask, Mat dst, float red_mul, float green_mul, float blue_mul);

void SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags);

void IlluminationChange(Mat src, Mat mask, Mat dst, float alpha, float beta);

void TextureFlattening(Mat src, Mat mask, Mat dst, float low_threshold, float high_threshold, int kernel_size);

MergeMertens MergeMertens_Create();
MergeMertens MergeMertens_CreateWithParams(float contrast_weight, float saturation_weight, float exposure_weight);
void MergeMertens_Close(MergeMertens b);
void MergeMertens_Process(MergeMertens b, struct Mats src, Mat dst);

AlignMTB AlignMTB_Create();
AlignMTB AlignMTB_CreateWithParams(int max_bits, int exclude_range, bool cut);
void AlignMTB_Close(AlignMTB b);
void AlignMTB_Process(AlignMTB b, struct Mats src, struct Mats dst);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PHOTO_H
