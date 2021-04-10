#ifndef _OPENCV3_XPHOTO_H_
#define _OPENCV3_XPHOTO_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xphoto/white_balance.hpp>
extern "C" {
#endif

#include "core.h"
    
#ifdef __cplusplus
typedef cv::Ptr<cv::xphoto::GrayworldWB>* GrayworldWB;
#else
typedef void* GrayworldWB;
#endif

GrayworldWB GrayworldWB_Create();
void GrayworldWB_Close(GrayworldWB b);
void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold);
float GrayworldWB_GetSaturationThreshold(GrayworldWB b);
void GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst); 

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XPHOTO_H
