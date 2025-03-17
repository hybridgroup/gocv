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

OpenCVResult ColorChange(Mat src, Mat mask, Mat dst, float red_mul, float green_mul, float blue_mul);

OpenCVResult SeamlessClone(Mat src, Mat dst, Mat mask, Point p, Mat blend, int flags);

OpenCVResult IlluminationChange(Mat src, Mat mask, Mat dst, float alpha, float beta);

OpenCVResult TextureFlattening(Mat src, Mat mask, Mat dst, float low_threshold, float high_threshold, int kernel_size);

OpenCVResult FastNlMeansDenoisingColoredMulti(struct Mats src, Mat dst, int imgToDenoiseIndex, int 	temporalWindowSize);
OpenCVResult FastNlMeansDenoisingColoredMultiWithParams(struct Mats src, Mat dst, int imgToDenoiseIndex, int 	temporalWindowSize, float 	h, float 	hColor, int 	templateWindowSize, int 	searchWindowSize );
OpenCVResult FastNlMeansDenoising(Mat src, Mat dst);
OpenCVResult FastNlMeansDenoisingWithParams(Mat src, Mat dst, float h, int templateWindowSize, int searchWindowSize);
OpenCVResult FastNlMeansDenoisingColored(Mat src, Mat dst);
OpenCVResult FastNlMeansDenoisingColoredWithParams(Mat src, Mat dst, float h, float hColor, int templateWindowSize, int searchWindowSize);

MergeMertens MergeMertens_Create();
MergeMertens MergeMertens_CreateWithParams(float contrast_weight, float saturation_weight, float exposure_weight);
OpenCVResult MergeMertens_Process(MergeMertens b, struct Mats src, Mat dst);
void MergeMertens_Close(MergeMertens b);

AlignMTB AlignMTB_Create();
AlignMTB AlignMTB_CreateWithParams(int max_bits, int exclude_range, bool cut);
OpenCVResult AlignMTB_Process(AlignMTB b, struct Mats src, struct Mats *dst);
void AlignMTB_Close(AlignMTB b);

OpenCVResult DetailEnhance(Mat src, Mat dst, float sigma_s, float sigma_r);
OpenCVResult EdgePreservingFilter(Mat src, Mat dst, int filter, float sigma_s, float sigma_r);
OpenCVResult PencilSketch(Mat src, Mat dst1, Mat dst2, float sigma_s, float sigma_r, float shade_factor);
OpenCVResult Stylization(Mat src, Mat dst, float sigma_s, float sigma_r);

OpenCVResult PhotoInpaint(Mat src, Mat mask, Mat dst, float inpaint_radius, int algorithm_type);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PHOTO_H
